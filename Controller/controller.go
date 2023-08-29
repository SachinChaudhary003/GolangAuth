package controller

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	helpers "github.com/sachinchaudhary003/golangAuth/Helpers"
	jwt "github.com/sachinchaudhary003/golangAuth/Jwt"
	model "github.com/sachinchaudhary003/golangAuth/Model"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
)

func getHash(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Fatal(err)
	}
	return string(hash)
}
func Userlogin(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	var user model.User
	var dbUser model.User
	json.NewDecoder(request.Body).Decode(&user)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := helpers.Collection.FindOne(ctx, bson.M{"email": user.Email}).Decode(&dbUser)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message":"` + err.Error() + `"}`))
		return
	}

	userPass := []byte(user.Password)
	dbPass := []byte(dbUser.Password)
	passErr := bcrypt.CompareHashAndPassword(dbPass, userPass)
	if passErr != nil {
		log.Println(passErr)
		response.Write([]byte(`{"response":"Wrong Password!"}`))
		return
	}
	jwtToken, err := jwt.GenerateToken()
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"message":"` + err.Error() + `"}`))
		return
	}
	response.Write([]byte(`{"token":"` + jwtToken + `"}`))

}

func UserSignup(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var user model.User
	json.NewDecoder(request.Body).Decode(&user)
	user.Password = getHash([]byte(user.Password))

	helpers.InsertOne(user)
	json.NewEncoder(response).Encode(user)
}
