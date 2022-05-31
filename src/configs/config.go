package configs

type Config struct {
	ServerPort       int    `split_words:"true"`
	TwitterAPIKey    string `split_words:"true"`
	TwitterAPISecret string `split_words:"true"`
}
