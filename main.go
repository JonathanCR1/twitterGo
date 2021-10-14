package main
import (
	"github.com/JonathanCR1/twitterGo/bd"
	"github.com/JonathanCR1/twitterGo/handlers"
	"log"
)
func main()  {
	if bd.ConnectionCheck() == 0 {
		log.Fatal("No connection to BBDD.")
		return
	}
	handlers.Controllers()
}

