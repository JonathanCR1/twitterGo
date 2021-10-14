package bd
import "golang.org/x/crypto/bcrypt"

//EncriptPassw function to encript the password of the user
func EncriptPassw(pass string) (string,error){
	costo:= 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass),costo)
	return string(bytes), err
}
