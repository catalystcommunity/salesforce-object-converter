package internal

import (
	"encoding/json"
	"fmt"
	"github.com/catalystcommunity/app-utils-go/errorutils"
	"github.com/catalystcommunity/app-utils-go/logging"
	"github.com/joomcode/errorx"
	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"
	"net/http"
	"net/url"
	"strings"
)

// salesforceCredentials represents the response from salesforce's /services/oauth2/token endpoint to get an access token
type salesforceCredentials struct {
	AccessToken string `json:"access_token"`
	InstanceUrl string `json:"instance_url"`
	Id          string `json:"id"`
	TokenType   string `json:"token_type"`
	IssuedAt    int    `json:"issued_at,string"`
	Signature   string `json:"signature"`
}

// GetSalesforceCredentials uses the configured credentials to get an access token
func GetSalesforceCredentials(domain, clientId, clientSecret, username, password, grantType string) (salesforceCredentials, error) {
	var creds salesforceCredentials
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	uri := getAuthUrl(domain, clientId, clientSecret, username, password, grantType)
	req.SetRequestURI(uri)
	req.Header.SetMethod("POST")
	res := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(res)
	err := fasthttp.Do(req, res)
	// return any errors
	if err != nil {
		errorutils.LogOnErr(nil, "error setting access token", err)
		return creds, err
	}
	statusCode := res.StatusCode()
	body := res.Body()
	// ensure we got 200 response
	if statusCode != http.StatusOK {
		entry := logging.Log.WithFields(logrus.Fields{"status_code": statusCode, "body": string(body)})
		errorutils.LogOnErr(entry, "unexpected status code attempting to set access token", errorx.IllegalState.New("unexpected status code"))
	}
	// unmarshal response body to struct
	err = json.Unmarshal(body, &creds)
	if err != nil {
		errorutils.LogOnErr(nil, "error unmarshalling access token response to struct", err)
		return creds, err
	}
	// ensure we actually got an access token
	if creds.AccessToken == "" {
		err = errorx.IllegalState.New("access token should not be empty")
		errorutils.LogOnErr(nil, "access token should not be empty", err)
		return creds, err
	}
	return creds, nil
}

func DescribeObject(domain, apiVersion, object, accessToken string) ([]byte, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.SetRequestURI(getDescribeUrl(domain, apiVersion, object))
	addAuthHeader(req, accessToken)
	res := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(res)
	err := fasthttp.Do(req, res)
	if err != nil {
		return nil, err
	} else if res.StatusCode() != http.StatusOK {
		return nil, errorx.IllegalState.New("unexpected response code: %d with body: %s", res.StatusCode(), res.Body())
	} else {
		return res.Body(), nil
	}
}

// getQueryUrl gets a formatted url to the soql query endpoint
func getQueryUrl(domain, apiVersion string, query string) string {
	formattedQuery := strings.Replace(query, " ", "+", -1)
	return fmt.Sprintf("%s/services/data/v%s/query?q=%s", getBaseUrl(domain), apiVersion, formattedQuery)
}

// getDescribeUrl gets a formatted url to the describe endpoint
func getDescribeUrl(domain, apiVersion, object string) string {
	return fmt.Sprintf("%s/services/data/v%s/sobjects/%s/describe", getBaseUrl(domain), apiVersion, object)
}

// getAuthUrl gets a formatted url to the token endpoint
func getAuthUrl(domain, clientId, clientSecret, username, password, grantType string) string {
	//return fmt.Sprintf("%s/services/oauth2/token?client_id=%s&client_secret=%s&username=%s&grant_type=%s&password=%s", getBaseUrl(domain), clientId, clientSecret, username, grantType, password)
	params := url.Values{}
	params.Add("client_id", clientId)
	params.Add("client_secret", clientSecret)
	params.Add("username", username)
	params.Add("password", password)
	params.Add("grant_type", grantType)
	return fmt.Sprintf("%s/services/oauth2/token?%s", getBaseUrl(domain), params.Encode())
}

// getBaseUrl gets a base url using the configured domain
func getBaseUrl(domain string) string {
	return fmt.Sprintf("https://%s", domain)
}

// addAuthHeader adds the access token from the config to the request
func addAuthHeader(req *fasthttp.Request, accessToken string) {
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))
}
