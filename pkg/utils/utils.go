package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"regexp"
	"strings"
	"time"
)

func ParseBody(r *http.Request, x interface{}) error {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// exception handling code here
		return err
	}
	defer r.Body.Close()

	// ok. parse json to struct.
	err = json.Unmarshal([]byte(body), x)
	if err != nil {
		// exception handling code here
		return err
	}

	// ok.
	return nil
}

// GetRandomNumber returns a random number.
func GetRandomNumber() int {

	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(99999999) + 1
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

	ID := fmt.Sprintf("%d%02d%02d.%09d", year, int(month), day, randomNumber)
	return ID
}

// IsValidEmailFormat checks if the input, v is in a valid email format.
// Return true if valid; false if not valid.
func IsValidEmailFormat(v string) bool {
	pattern := regexp.MustCompile(`^[a-z0-9_]+[.-][a-z0-9]+@\w+\.([a-z0-9]{2,4}|[a-z]{2}.[a-z]{2})$`)
	return pattern.MatchString(v)
}

// IsAlphaAndSpaceOnly checks if the input, v contains only alphabets and white space characters (i.e. space or tab) only.
// Return true if valid; false if not valid.
func IsAlphaAndSpaceOnly(v string) bool {
	pattern := regexp.MustCompile(`^[a-zA-z\s]+$`)
	return pattern.MatchString(v)
}

// AreAllowedCharacters checks if the input, v contains any of the following characters that are not allowed:
// - /
// - \
// - -
// Returns
// - true if v does not contain any not allowed characters;
// - false if v contains any of the not allowed characters.
func AreAllowedCharacters(v string) bool {
	pattern := regexp.MustCompile(`^([a-zA-z0-9]|[^\/-])+$`)
	return pattern.MatchString(v)
}

// SendErrorMsgToClient prepares:
// - a INTERNAL SERVER ERROR response header;
// - a JSON body containing:
//   * "ok" attribute, set to false;
//   * "msg" attribute set to the error message passed in;
//   * "data" attribute set to {}
func SendErrorMsgToClient(w *http.ResponseWriter, err error) {
	(*w).WriteHeader(http.StatusInternalServerError)
	body := fmt.Sprintf(`{
			"ok" : false,
			"msg" : "%s",
			"data" : {}
		}`, err.Error())
	(*w).Write([]byte(body))
	//
}

// SendNotFoundMsgToClient prepares:
// - a NOT FOUND response header;
// - a JSON body containing:
//   * "ok" attribute, set to false;
//   * "msg" attribute set to the error message passed in;
//   * "data" attribute set to {}
func SendNotFoundMsgToClient(w *http.ResponseWriter, err error) {
	(*w).WriteHeader(http.StatusNotFound)
	body := fmt.Sprintf(`{
			"ok" : false,
			"msg" : "%s",
			"data" : {}
		}`, err.Error())
	(*w).Write([]byte(body))
	//
}

// SendBadRequestMsgToClient prepares:
// - a BAD REQUEST response header;
// - a JSON body containing:
//   * "ok" attribute, set to false;
//   * "msg" attribute set to the error message passed in;
//   * "data" attribute set to {}
func SendBadRequestMsgToClient(w *http.ResponseWriter, err error) {
	(*w).WriteHeader(http.StatusBadRequest)
	body := fmt.Sprintf(`{
			"ok" : false,
			"msg" : "%s",
			"data" : {}
		}`, err.Error())
	(*w).Write([]byte(body))
	//
}

// SendDataToClient prepares:
// - a OK header;
// - a JSON body containing:
//   * "ok" attribute, set to true;
//   * "msg" attriubte;
//   * "data" attribute, set to the data passed in
// and send the response to the client.
func SendDataToClient(w *http.ResponseWriter, data []byte, msg string) {

	(*w).WriteHeader(http.StatusOK)

	var _msg string
	if strings.TrimSpace(msg) == "" {
		_msg = "n.a."
	} else {
		_msg = msg
	}

	if data != nil {

		// data is available
		body := fmt.Sprintf(`{
			"ok" : true,
			"msg" : "%s",
			"data" : %s
		}`, _msg, string(data))
		(*w).Write([]byte(body))
		//
	} else {

		// no data
		body := fmt.Sprintf(`{
			"ok" : true,
			"msg" : "%s",
			"data" : {}
		}`, _msg)
		(*w).Write([]byte(body))
		//
	}

	//
}
