package main

import (
	"fmt"

	"github.com/gorilla/mux"
)

type Movie struct{
 ID string `json:"id"`
 Isbn string `json:"isbn"`
 Title string `json:"title"`
 Director *Director `json:"director"`

}

type Director struct{
 FirstName string `json:"firstname"`
 LastName string `json:"lastname"`
 Nationality string `json:"nationality"`
 BirthYear int `json:"birthyear"`

}

var movies []Movie 

func main(){
  r := mux.NewRouter()
  r.
}