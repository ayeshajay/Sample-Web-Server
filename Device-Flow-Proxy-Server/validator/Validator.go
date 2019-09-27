package validator

import "github.com/ayeshajay/Device-Flow-Proxy-Server/config"

func ValidateClientId(ClientID string) (out bool) {
	out = ClientID==config.GetClientId()
	return
}

func ValidateGrantType(GrantType string) (out bool){
	out = GrantType==config.GetGrantType()
	return
}