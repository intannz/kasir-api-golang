package util

import "github.com/spf13/viper"

// Config
type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

// LoadConfig
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	// baca environment variables
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		// 2. INI KUNCINYA:
		// Cek jenis errornya. Kalau errornya cuma "File Gak Ketemu", ABAIKAN.
		// Jangan langsung return error.
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// File gak ada? Gpp, lanjut aja. Kita kan punya environment variable.
		} else {
			// Kalau errornya lain (misal syntax salah), baru lapor error.
			return
		}
	}

	err = viper.Unmarshal(&config)
	return
}
