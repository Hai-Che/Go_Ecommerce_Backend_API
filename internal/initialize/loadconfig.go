package initialize

import (
	"fmt"

	"github.com/Hai-Che/Go_Ecommerce_Backend_API/global"
	"github.com/spf13/viper"
)

func LoadConfig() {
	viper := viper.New()
	viper.AddConfigPath("./config/")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Failed to read configuration %w \n", err))
	}
	fmt.Println("Server port:", viper.GetInt("server.port"))
	fmt.Println("Server key:", viper.GetString("security.jwt.key"))

	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("Unable to decode configuration %v", err)
	}
}
