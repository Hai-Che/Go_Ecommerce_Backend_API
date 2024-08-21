package main

import (
	"github.com/Hai-Che/Go_Ecommerce_Backend_API/internal/routers"
)

func main() {
	r := routers.NewRouter()
	r.Run("localhost:8000")
}
