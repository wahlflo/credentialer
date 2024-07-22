package llms

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strings"
	"sync"
	"time"
)

type ollamaConnector struct {
	addressOfService        string
	mutex                   *sync.Mutex
	currentNumberOfSessions int
	maximumNumberOfSessions int
}

func NewOllamaConnector(addressOfService string, maximumNumberOfSessions int) LlmConnector {
	return &ollamaConnector{
		addressOfService:        addressOfService,
		mutex:                   &sync.Mutex{},
		currentNumberOfSessions: 0,
		maximumNumberOfSessions: maximumNumberOfSessions,
	}
}

func (connector *ollamaConnector) CheckConnection() error {
	response, err := http.Get(connector.addressOfService)
	if err != nil {
		return err
	}

	if response.StatusCode != 200 {
		return errors.New("response status code " + fmt.Sprint(response.StatusCode))
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return errors.New("error reading response body: " + err.Error())
	}

	if string(responseBody) == "Ollama is running" {
		return nil
	}
	return errors.New("received unexpected response from ollama service: " + string(responseBody))
}

func (connector *ollamaConnector) GetResponseOutputModifier() string {
	return " Respond to the Question only with 'yes' or 'no'. "
}

func (connector *ollamaConnector) GetBooleanResponse(prompt string) (bool, error) {
	connector.blockUntilSessionIsFree()
	defer connector.freeSession()

	slog.Debug("LLM is asked a boolean question", "prompt", prompt)

	start := time.Now()

	body := map[string]interface{}{
		"model":      "llama3",
		"prompt":     prompt,
		"stream":     false,
		"keep_alive": "1m",
	}

	encodedBody, err := json.Marshal(body)
	if err != nil {
		panic(err.Error())
	}

	response, err := http.Post(connector.addressOfService+"/api/generate", "application/json", bytes.NewBuffer(encodedBody))
	if err != nil {
		slog.Warn("error: " + err.Error())
		return false, err
	}

	if response.StatusCode != 200 {
		return false, errors.New("response status code not 200. is " + fmt.Sprint(response.StatusCode))
	}

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		slog.Warn("error: " + err.Error())
		return false, errors.New("error reading response body: " + err.Error())
	}

	parsedBody := make(map[string]interface{})
	err = json.Unmarshal(responseBody, &parsedBody)
	if err != nil {
		slog.Warn("error: " + err.Error())
		return false, errors.New("error parsing body: " + err.Error())
	}

	responseMessage := parsedBody["response"].(string)

	responseMessage = strings.Replace(responseMessage, "'", "", -1)
	responseMessage = strings.Replace(responseMessage, ".", "", -1)
	responseMessage = strings.ToLower(responseMessage)

	slog.Debug("response from LLM", "response", responseMessage, "time", time.Now().Sub(start).Seconds())

	switch responseMessage {
	case "yes":
		return true, nil
	case "no":
		return false, nil
	default:
		return false, errors.New("unexpected response from LLM: " + responseMessage)
	}
}

func (connector *ollamaConnector) blockUntilSessionIsFree() {
	for {
		connector.mutex.Lock()
		if connector.currentNumberOfSessions < connector.maximumNumberOfSessions {
			connector.currentNumberOfSessions += 1
			connector.mutex.Unlock()
			return
		} else {
			connector.mutex.Unlock()
			time.Sleep(1 * time.Second)
		}
	}
}

func (connector *ollamaConnector) freeSession() {
	connector.mutex.Lock()
	defer connector.mutex.Unlock()
	connector.currentNumberOfSessions -= 1
}
