package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"
)

type city struct {
	Name string
	Area int
}

// Middleware function to check the content type as json
func filterContentType(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Currently in check content type middleware")
		// Filtering the requests by MIME type
		if r.Header.Get("Content-type") != "application/json" {
			w.WriteHeader(http.StatusUnsupportedMediaType)
			w.Write([]byte("415 - Unsupport Media type, Please send Json\n"))
			return
		}
		handler.ServeHTTP(w, r)
	})
}

// middleware function to add server time stamp for response cookie
func setServerTimeCookie(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.ServeHTTP(w, r)
		// setting cookie to each and every response
		cookie := http.Cookie{
			Name:  "Server-Time(UTC)",
			Value: strconv.FormatInt(time.Now().Unix(), 10),
		}
		http.SetCookie(w, &cookie)
		log.Println("Currently in the set server time middleware")
	})
}

func mainLogic(w http.ResponseWriter, r *http.Request) {
	// check if method is POST
	if r.Method == "POST" {
		var tempCity city
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&tempCity)
		if err != nil {
			panic(err)
		}
		defer r.Body.Close()
		// resource creation logic comes heres, as of now it is plain text
		log.Printf("Got %s city with area of %d sq miles! \n", tempCity.Name, tempCity.Area)
		// Tell everything is fine
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("201 - created\n"))
	} else {
		//say method not allowed
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("405 - Method Not allowed"))
	}
}

func main() {
	mainLogicHandler := http.HandlerFunc(mainLogic)
	http.Handle("/city", filterContentType(setServerTimeCookie(mainLogicHandler)))
	http.ListenAndServe(":8080", nil)
}
