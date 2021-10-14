package middlew
import (
	"github.com/JonathanCR1/twitterGo/bd"
	"net/http"
)

//CheckBD function to verify if BBDD is ready to start the operations
func CheckBD(next http.HandlerFunc) http.HandlerFunc  {
	return func(w http.ResponseWriter , r *http.Request){
		if bd.ConnectionCheck() == 0{
			http.Error(w,"Lost connection",500)
			return
		}
		next.ServeHTTP(w,r)
	}
}
