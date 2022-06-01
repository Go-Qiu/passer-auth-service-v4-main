package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"
)

func ParseBody(r *http.Request, x interface{}) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// exception handling code here
		return
	}

	// ok. parse json to struct.
	err = json.Unmarshal([]byte(body), x)
	if err != nil {
		// exception handling code here
		return
	}
}

// GetRandomNumber returns a random number.
func GetRandomNumber() int {

	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(9999999) + 1
	return n
}

// GenerateID returns a randomly generated ID string.
// The randomly generated ID string is in the following format:
// YYYY.MM.DD.RRRRRRRR
// where
// - YYYY represents the Year of the system date
// - MM represents the Month of the system date
// - DD represents the Day of the system date
// - RRRRRRRR is the zero-padded random number
func GenerateID() string {

	year, month, day := time.Now().Local().Date()
	randomNumber := GetRandomNumber()

	ID := fmt.Sprintf("%d.%02d.%02d.%08d", year, int(month), day, randomNumber)
	return ID
}
