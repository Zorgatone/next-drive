package api

import (
	"next-drive/constants"
	httpclient "next-drive/http-client"
)

type ApiClient struct {
	httpClient *httpclient.HttpClient
}

func NewApiClient() *ApiClient {
	var httpClient = httpclient.NewHttpClient()
	httpClient.SetUserAgent(constants.USER_AGENT)

	return &ApiClient{
		httpClient: httpClient,
	}
}
