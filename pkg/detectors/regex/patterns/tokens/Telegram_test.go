package tokens

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_telegram_bot_token_positive(t *testing.T) {
	testContent := `
https://api.telegram.org/bot587523485:AAG7sasdnsd782839213zGo/getUpdates
`
	pattern := TelegramBotToken()
	matches := pattern.GetMatches("", []byte(testContent))
	require.Equal(t, 1, len(matches))
	require.Equal(t, "bot587523485:AAG7sasdnsd782839213zGo", matches[0])
}

func Test_telegram_bot_token_negative(t *testing.T) {
	testContent := `
BoT6:D
BOT7:ug
`
	pattern := TelegramBotToken()
	matches := pattern.GetMatches("", []byte(testContent))
	require.Equal(t, 0, len(matches))
}
