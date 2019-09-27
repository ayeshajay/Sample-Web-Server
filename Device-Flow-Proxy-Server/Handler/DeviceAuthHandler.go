package Handler

import (
	"encoding/json"
	"github.com/ayeshajay/Device-Flow-Proxy-Server/Parameters"
	"github.com/ayeshajay/Device-Flow-Proxy-Server/config"
	error2 "github.com/ayeshajay/Device-Flow-Proxy-Server/error"
	"github.com/ayeshajay/Device-Flow-Proxy-Server/generator"
	"github.com/ayeshajay/Device-Flow-Proxy-Server/validator"
	"github.com/twinj/uuid"
	"log"
	"net/http"
)

func DeviceAuthHandler(w http.ResponseWriter, r *http.Request) {

	key:= r.URL.Query()
	ClientId:= key.Get("client_id")
	//keys2 := key.Get("key2")
	//log.Println(ClientId)

	if ClientId== "" ||len(key) < 1 {
		log.Println("ClientId is missing")
		response := error2.Error{
			Error:"Invalid request",
			ErrorDescription: "Missing required parameter",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
	}

	// Query()["key"] will return an array of items,
	// we only want the single item.


	if validator.ValidateClientId(ClientId) {
		log.Println("Url Param 'client_id' is: " + string(ClientId))
		response := Parameters.ParamAuthResponse{
			DeviceCode:      uuid.NewV4().String(),
			UserCode:        generator.RandStringRunes(6),
			VerificationURI: config.GetVerificationURI(),
		}
		json.NewEncoder(w).Encode(response)
	} else {
		response := error2.Error{
			Error:"Invalid client",
			ErrorDescription: "Required parameter is wrong",
		}
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(response)
		log.Println("Url Param 'client_id' is not valid")
	}

}