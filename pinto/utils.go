package pinto

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	// using name "pinto_provider" because "provider" is a reserved word for resources and data sources
	schemaProvider = "pinto_provider"
	// using name "pinto_environment" to keep the same naming schema as schemaProvider
	schemaEnvironment = "pinto_environment"
)

func handleClientError(op string, errorString string, httpResponse *http.Response) string {
	bodyBytes, err := ioutil.ReadAll(httpResponse.Body)
	bodyString := ""
	if err == nil {
		bodyString = string(bodyBytes)
		log.Printf("[ERROR] Unable to perform operation %s. \n Reason: %s \n Details: %s", op, errorString, bodyString)
		return errorString + ": " + bodyString
	} else {
		log.Printf("[ERROR] Unable to perform operation %s. \n Reason: %s", op, errorString)
		return errorString
	}
}

func getProvider(p *PintoProvider, d *schema.ResourceData) (string, error) {
	res := ""
	spec, ok := d.GetOk(schemaProvider)
	if ok {
		res = spec.(string)
	} else {
		if p.provider != "" {
			res = p.provider
		} else {
			return "", fmt.Errorf("invalid configuration. %s has to be set on provider or resource-level", schemaProvider)
		}
	}
	return res, nil
}

func getEnvironment(p *PintoProvider, d *schema.ResourceData) string {
	res := ""
	spec, ok := d.GetOk(schemaEnvironment)
	if ok {
		res = spec.(string)
	} else {
		res = p.environment
	}
	return res
}

// TODO: Clarify missing struct in client
type AccessOptions struct {
	Provider      string `json:"provider"`
	Environment   string `json:"environment"`
	CredentialsId string `json:"credentials_id"`
}
type XApiOptions struct {
	AccessOptions AccessOptions `json:"access_options"`
}
