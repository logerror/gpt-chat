package configs

type Configuration struct {
	OpenAISecretKey  string  `mapstructure:"open_ai_secret_key"`
	BotName          string  `mapstructure:"bot_name"`
	Model            string  `mapstructure:"model"`
	MaxTokens        int     `mapstructure:"max_tokens"`
	TopP             float32 `mapstructure:"top_p"`
	FrequencyPenalty float32 `mapstructure:"frequency_penalty"`
	PresencePenalty  float32 `mapstructure:"presence_penalty"`
}
