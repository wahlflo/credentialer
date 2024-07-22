package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/fatih/color"
	"github.com/wahlflo/credentialer/llms"
	"github.com/wahlflo/credentialer/pkg"
	"github.com/wahlflo/credentialer/pkg/detectors"
	"github.com/wahlflo/credentialer/pkg/interfaces"
	"github.com/wahlflo/credentialer/pkg/output_formatters"
	"io/fs"
	"log"
	"log/slog"
	"math"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
)

type appOptions struct {
	pathToDirectory                      string
	optionOutputFormat                   string
	parsedOutputFormat                   interfaces.OutputFormatter
	findingDestination                   string
	parsedDestination                    *os.File
	showHelp                             bool
	silent                               bool
	noColor                              bool
	debug                                bool
	optionResumeFromFile                 string
	optionLogScannedFiles                string
	optionLogFilesWhichCouldNotBeScanned string
	largeLanguageModel                   string
	ollamaUrl                            string
	ollamaModel                          string
	ollamaSessions                       int
}

func (o *appOptions) parse() error {
	flag.StringVar(&o.pathToDirectory, "i", "", "path to the directory which should be scanned")
	flag.StringVar(&o.optionOutputFormat, "format", "text", "type of output [text, json, csv], default is text")
	flag.StringVar(&o.findingDestination, "o", "", "output file where findings will be stored, default is standard output")
	flag.BoolVar(&o.showHelp, "help", false, "shows help information")
	flag.BoolVar(&o.silent, "s", false, "suppress info messages for a clearer output")
	flag.BoolVar(&o.noColor, "no-color", false, "don't use ANSI escape color codes in output")
	flag.StringVar(&o.optionResumeFromFile, "resume", "", "resume scan based on file containing already scanned files")
	flag.StringVar(&o.optionLogScannedFiles, "success-logMessage", "", "logMessage scanned files to a file, needed when you want to resume a paused scan")
	flag.StringVar(&o.optionLogFilesWhichCouldNotBeScanned, "failed-logMessage", "", "logMessage files, which could not be scanned, to a file")
	flag.StringVar(&o.largeLanguageModel, "llm", "none", "which large language model should be used for fine tuning, possible options: [none, ollama], default is none")
	flag.StringVar(&o.ollamaUrl, "ollama-url", "http://127.0.0.1:11434", "URL of the ollama LLM, default value is http://127.0.0.1:11434")
	flag.IntVar(&o.ollamaSessions, "ollama-sessions", 1, "number of concurrent sessions, default is 1")
	flag.StringVar(&o.ollamaModel, "ollama-model", "llama3", "name of the model which should be used, default is llama3")
	flag.BoolVar(&o.debug, "debug", false, "displays debug information")
	flag.Parse()

	outputFile, err := o.parseFindingDestination()
	if err != nil {
		return err
	}

	if err := o.parseOutputFormatOption(outputFile); err != nil {
		return err
	}
	if err := o.ensureThatOptionPathToDirectoryIsValid(); err != nil {
		return err
	}

	if o.largeLanguageModel != "ollama" {
		if o.ollamaUrl != "http://127.0.0.1:11434" {
			return errors.New("[!] the option -ollama-url can only be specified if the ollama LLM is used")
		}
		if o.ollamaSessions != 1 {
			return errors.New("[!] the option -ollama-sessions can only be specified if the ollama LLM is used")
		}
		if o.ollamaModel != "llama3" {
			return errors.New("[!] the option -ollama-model can only be specified if the ollama LLM is used")
		}
	}

	return nil
}

func (o *appOptions) parseOutputFormatOption(file *os.File) error {
	switch o.optionOutputFormat {
	case "text":
		o.parsedOutputFormat = output_formatters.NewOutputText(file, !o.noColor)
	case "json":
		o.parsedOutputFormat = output_formatters.NewOutputJson(file, !o.noColor)
	case "csv":
		o.parsedOutputFormat = output_formatters.NewOutputCsv(file, !o.noColor)
		break
	default:
		return errors.New("[!] invalid parameter for output: \"" + o.optionOutputFormat + "\"")
	}
	return nil
}

func (o *appOptions) ensureThatOptionPathToDirectoryIsValid() error {
	if o.pathToDirectory == "" {
		return errors.New("[!] path to the directory which should be scanned has to be specified")
	}

	if _, err := os.Stat(o.pathToDirectory); os.IsNotExist(err) {
		return errors.New("[!] path to the directory is not valid")
	}
	return nil
}

func (o *appOptions) parseFindingDestination() (*os.File, error) {
	if o.findingDestination == "" {
		return os.Stdout, nil
	}

	return os.OpenFile(o.findingDestination, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0640)
}

func printBanner() {
	color.Green("  ____              _            _   _       _")
	color.Green(" / ___|_ __ ___  __| | ___ _ __ | |_(_) __ _| | ___ _ __")
	color.Green("| |   | '__/ _ \\/ _` |/ _ \\ '_ \\| __| |/ _` | |/ _ \\ '__|")
	color.Green("| |___| | |  __/ (_| |  __/ | | | |_| | (_| | |  __/ |")
	color.Green(" \\____|_|  \\___|\\__,_|\\___|_| |_|\\__|_|\\__,_|_|\\___|_|")
	println()
}

func exitWithError(err error) {
	printBanner()

	color.Red(err.Error())

	println()
	flag.Usage()
	os.Exit(1)
}

func logMessage(silent bool, message string) {
	if !silent {
		fmt.Println("[+] " + message)
	}
}

func loadAlreadyPreviousScannedFiles(optionResumeFromFile string) map[string]struct{} {
	resultMap := make(map[string]struct{})

	fileContent, err := os.ReadFile(optionResumeFromFile)
	if err != nil {
		color.Red("[!] fatal error while creating file: \"" + optionResumeFromFile + "\"" + err.Error())
		color.Red("[!] exiting due error")
		os.Exit(1)
	}

	for _, line := range strings.Split(string(fileContent), "\n") {
		resultMap[line] = struct{}{}
	}
	return resultMap
}

func logScannedFiles(scanner *pkg.Scanner, optionLogScannedFiles string) {
	channel := scanner.GetScannedFiles()

	outputFile, err := os.OpenFile(optionLogScannedFiles, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0640)
	if err != nil {
		color.Red("[!] fatal error while opening file: \"" + optionLogScannedFiles + "\" - error: " + err.Error())
		color.Red("[!] exiting due error")
		os.Exit(1)
	}
	defer outputFile.Close()

	for {
		scannedFile, channelOpen := <-channel
		if !channelOpen {
			return
		}
		if _, err := outputFile.WriteString(scannedFile + "\n"); err != nil {
			color.Red("[!] fatal error while writing to file: \"" + optionLogScannedFiles + "\" - error: " + err.Error())
			color.Red("[!] exiting due error")
			os.Exit(1)
		}
	}
}

func logFilesWithError(scanner *pkg.Scanner, optionLogFilesWhichCouldNotBeScanned string) {
	channel := scanner.GetFilesWhichCouldNotBeScanned()

	outputFile, err := os.OpenFile(optionLogFilesWhichCouldNotBeScanned, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0640)
	if err != nil {
		color.Red("[!] fatal error while opening file: \"" + optionLogFilesWhichCouldNotBeScanned + "\" - error: " + err.Error())
		color.Red("[!] exiting due error")
		os.Exit(1)
	}
	defer outputFile.Close()

	for {
		scannedFile, channelOpen := <-channel
		if !channelOpen {
			return
		}
		if _, err := outputFile.WriteString(scannedFile + "\n"); err != nil {
			color.Red("[!] fatal error while writing to file: \"" + optionLogFilesWhichCouldNotBeScanned + "\" - error: " + err.Error())
			color.Red("[!] exiting due error")
			os.Exit(1)
		}
	}
}

type file struct {
	filename string
	filepath string
}

func newFile(filename string, filepath string) interfaces.File {
	return &file{
		filename: filename,
		filepath: filepath,
	}
}

func (receiver file) GetFilename() string {
	return receiver.filename
}

func (receiver file) GetFilepath() string {
	return receiver.filepath
}

func createLlmConnector(options *appOptions) (interfaces.LlmConnector, error) {
	if options.largeLanguageModel == "none" {
		return nil, nil
	}

	if options.largeLanguageModel == "ollama" {
		connector := llms.NewOllamaConnector(options.ollamaUrl, options.ollamaSessions, options.ollamaModel)
		return connector, connector.CheckConnection()
	}

	return nil, errors.New("invalid value for -llm, see help for possible values")
}

func main() {
	options := appOptions{}
	if err := options.parse(); err != nil {
		exitWithError(err)
	}

	if options.debug {
		logMessage(options.silent, "debug messages will be displayed")
		log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
		slog.SetLogLoggerLevel(slog.LevelDebug)
	} else if options.silent {
		slog.SetLogLoggerLevel(slog.LevelError)
	} else {
		slog.SetLogLoggerLevel(slog.LevelWarn)
	}

	llmConnector, err := createLlmConnector(&options)
	if err != nil {
		color.Red("[!] fatal error while setting up connector to LLM: " + err.Error())
		color.Red("[!] exiting due error")
		os.Exit(1)
	}

	if llmConnector != nil {
		logMessage(options.silent, "a Large Language Model is used for pre-filtering the findings")
	}

	outputFormat := options.parsedOutputFormat

	alreadyPreviousScannedFiles := make(map[string]struct{})
	if options.optionResumeFromFile != "" {
		logMessage(options.silent, "load previous scanned files from file: "+options.optionResumeFromFile)
		alreadyPreviousScannedFiles = loadAlreadyPreviousScannedFiles(options.optionResumeFromFile)
		logMessage(options.silent, "loaded "+fmt.Sprint(len(alreadyPreviousScannedFiles))+" previous scanned files")
	}

	fileQueue := make(chan interfaces.File, 15000000)
	go func() {
		logMessage(options.silent, "start loading files which should be scanned")
		err := filepath.WalkDir(options.pathToDirectory, func(s string, d fs.DirEntry, err error) error {
			if err != nil {
				color.Red("[!] fatal error while crawling directory: " + err.Error())
				color.Red("[!] exiting due error")
				os.Exit(1)
			}

			// Ignore symbolic links
			if d.Type()&os.ModeSymlink != 0 {
				return nil
			}

			if !d.IsDir() {
				if _, previouslyScanned := alreadyPreviousScannedFiles[s]; !previouslyScanned {
					fileQueue <- newFile(d.Name(), s)
				}
			}
			return nil
		})
		if err != nil {
			color.Red("[!] fatal error while crawling directory: " + err.Error())
			color.Red("[!] exiting due error")
			os.Exit(1)
		}
		close(fileQueue)
		logMessage(options.silent, "loading files which should be scanned finished")
	}()

	scanner := pkg.NewScanner(fileQueue, outputFormat, llmConnector)
	for _, detector := range detectors.LoadBasicDetectors() {
		scanner.AddDetector(detector)
	}

	if options.optionLogScannedFiles != "" {
		logMessage(options.silent, "logging scanned files to: "+options.optionLogScannedFiles)
		scanner.SetOptionLogScannedFiles()
		go logScannedFiles(scanner, options.optionLogScannedFiles)
	}
	if options.optionLogFilesWhichCouldNotBeScanned != "" {
		logMessage(options.silent, "logging scanned files which could not be scanned to: "+options.optionLogFilesWhichCouldNotBeScanned)
		scanner.SetOptionLogFilesWhichCouldNotBeScanned()
		go logFilesWithError(scanner, options.optionLogFilesWhichCouldNotBeScanned)
	}

	numberOfScannerProcesses := int(math.Max(float64(runtime.NumCPU()-1), 1))
	logMessage(options.silent, "start "+fmt.Sprint(numberOfScannerProcesses)+" processes to scan for credentials")
	outputFormat.Start()
	scanner.StartScanning(numberOfScannerProcesses)

	startTime := time.Now()

	for scanner.IsRunning() {
		time.Sleep(5 * time.Second)

		timeRunning := time.Now().Sub(startTime)
		scannedFiles := scanner.GetNumberOfFilesScanned()

		secondsPerFile := timeRunning.Seconds() / float64(scannedFiles)
		remainingTimeInMilliseconds := float64(len(fileQueue)) * secondsPerFile
		remainingTime := time.Second * time.Duration(remainingTimeInMilliseconds)

		totalNumberOfFiles := len(fileQueue) + scannedFiles

		message := fmt.Sprintf("running for %v; processed files: %v / %v; remaining time: %v", timeRunning, scannedFiles, totalNumberOfFiles, remainingTime)
		logMessage(options.silent, message)
	}

	logMessage(options.silent, "finished scanning")
	outputFormat.Finished()
	logMessage(options.silent, "terminating")
}
