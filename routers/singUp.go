package routers

import (
	"encoding/json"
	"github.com/JonathanCR1/twitterGo/bd"
	"github.com/JonathanCR1/twitterGo/models"
	"net/http"
)

//SingUp function to register new users
func SingUp(w http.ResponseWriter , r *http.Request){

	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil{
		http.Error(w, "ERROR: The data is not correctly " + err.Error(), 400)
		return
	}

	if len(user.Email) == 0 {
		http.Error(w,"ERROR: The email is mandatory.",400)
		return
	}
	if len(user.Password) == 6 {
		http.Error(w,"ERROR: The password must be greater than 6.",400)
		return
	}

	_,encontrado,_ := bd.UserExist(user.Email)
	if encontrado{
		http.Error(w,"ERROR: This email is registered on the BBDD.",400)
		return
	}
	_,status,err:=bd.RegisterUser(user)
	if err!=nil{
		http.Error(w,"ERROR: Cant register the user."+err.Error(),400)
		return
	}
	if status == false{
		http.Error(w,"ERROR: Cant register the user correctly.",400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}