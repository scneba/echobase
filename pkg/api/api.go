package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/pkg/errors"
	"gobase.com/base/pkg/registering"
)

func RegisterUser(service registering.RegisteringInterface) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		user := registering.User{}
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			//write error HTTP request failed to read the body
			writeError(w, errors.New("Failed to read content"), nil)
			return
		}
		if err = json.Unmarshal(bodyBytes, &user); err != nil {
			//write error failed to unmarshal
			writeError(w, errors.New("Failed to unmarshal content"), nil)
			return
		}
		id, regErrors := service.RegisterUser(user)
		if len(regErrors) > 0 {
			writeErrors(w, regErrors, nil)
			return
		}
		writeJSON(w, id)
	}
}
