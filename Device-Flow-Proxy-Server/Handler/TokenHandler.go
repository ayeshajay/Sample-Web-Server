package Handler

import (
	"encoding/json"
	"github.com/ayeshajay/Device-Flow-Proxy-Server/Parameters"
	error2 "github.com/ayeshajay/Device-Flow-Proxy-Server/error"
	"github.com/ayeshajay/Device-Flow-Proxy-Server/validator"
	"github.com/myesui/uuid"
	"log"
	"net/http"
)


func TokenHandler(w http.ResponseWriter, r *http.Request) {

	key:= r.URL.Query()
	ClientId:= key.Get("client_id")
	GrantType:= key.Get("grant_type")
	DeviceCode:=key.Get("device_code")
	//keys2 := key.Get("key2")
	//log.Println(ClientId)

	if ClientId== "" || GrantType=="" || DeviceCode==""|| len(key) != 3 {
		log.Println("Parameter is missing")
		response := error2.Error{
			Error:"Invalid request",
			ErrorDescription: "Missing required parameter",
		}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
	}

	// Query()["key"] will return an array of items,
	// we only want the single item.


	if validator.ValidateClientId(ClientId) && validator.ValidateGrantType(GrantType){
		log.Println("Url Param 'client_id' is: " + string(ClientId))
		response := Parameters.ParamTokenResponse{
			AccessToken:      uuid.NewV4().String(),
			TokenType:        "bearer",
			ExpiresInToken: "1000",
		}
		json.NewEncoder(w).Encode(response)
	} else {
		log.Println("Url Param 'client_id' is not valid")
	}

}
