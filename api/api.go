package api

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"strconv"
	"time"

	"github.com/donh/shortner/pkg/models"
	"github.com/donh/shortner/pkg/storage"
	"github.com/donh/shortner/pkg/util"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// SetAPIRoutes set API routes
func SetAPIRoutes() {
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/add", add).Methods("POST")
	router.HandleFunc("/{key}", parse).Methods("GET")

	c := cors.New(cors.Options{
		AllowedMethods: []string{"GET", "POST"},
	})

	handler := cors.Default().Handler(router)
	handler = c.Handler(handler)
	port := util.Config().Port
	s := ":" + strconv.Itoa(port)
	log.Println("API server started. Listening on port:", port)
	log.Fatal(http.ListenAndServe(s, handler))
}

// characters used for conversion
const alphabet = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const base = "http://localhost:8000/"
const multiplier = 100000

// converts number to base62
func encode(number int) string {
	if number == 0 {
		return string(alphabet[0])
	}

	chars := make([]byte, 0)

	length := len(alphabet)

	for number > 0 {
		result := number / length
		remainder := number % length
		chars = append(chars, alphabet[remainder])
		number = result
	}

	for i, j := 0, len(chars)-1; i < j; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}

	return string(chars)
}

// converts base62 token to int
func decode(token string) int {
	number := 0
	idx := 0.0
	chars := []byte(alphabet)

	charsLength := float64(len(chars))
	tokenLength := float64(len(token))

	for _, c := range []byte(token) {
		power := tokenLength - (idx + 1)
		index := bytes.IndexByte(chars, c)
		number += index * int(math.Pow(charsLength, power))
		idx++
	}

	return number
}

func add(w http.ResponseWriter, r *http.Request) {
	var err error
	req := map[string]interface{}{}
	body, errParseBody := ioutil.ReadAll(r.Body)
	if errParseBody != nil {
		http.Error(w, errParseBody.Error(), http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(body, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	url := ""
	if val, ok := req["url"]; ok {
		url = val.(string)
	}

	id, _, short, err := storage.Query("url", url)
	if id > 0 {
		if short == "" {
			short = encode(id*multiplier + len(url))
			err = storage.Update(id, short)
			if err == nil {
				setResponse(w, base+short, http.StatusOK, "")
				return
			}
		}
		setResponse(w, base+short, http.StatusOK, "")
		return
	}

	id = storage.Insert(url)
	if id == 0 {
		setResponse(w, "Failed", http.StatusBadRequest, err.Error())
		return
	}

	short = encode(id*multiplier + len(url))
	err = storage.Update(id, short)
	if err != nil {
		setResponse(w, "Failed", http.StatusBadRequest, err.Error())
		return
	}

	setResponse(w, base+short, http.StatusOK, "")
}

func parse(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	short := vars["key"]
	id, url, _, _ := storage.Query("short", short)
	if id > 0 {
		if short == "" {
			short = encode(id*multiplier + len(url))
			_ = storage.Update(id, short)
		}
	}
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func setResponse(w http.ResponseWriter, result interface{}, status int, errorMessage string) {
	response := models.ResponseWrapper{
		Result: result,
		Status: status,
		Error:  errorMessage,
		Time:   getNow(),
	}
	js, err := json.Marshal(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(js)
}

func getNow() string {
	t := time.Now()
	now := t.Format("2006-01-02 15:04:05")
	return now
}
