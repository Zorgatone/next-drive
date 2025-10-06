package constants

import (
	"net/http"

	. "github.com/Zorgatone/next-drive-api/constants"
)

func GetDefaultHeaders() http.Header {
	headers := http.Header{}

	headers.Set(UserAgentHeaderName, UserAgentHeaderValue)
	headers.Set(AcceptHeaderName, "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	headers.Set(AcceptEncodingHeaderName, "gzip, deflate, br, zstd")
	headers.Set(AcceptLanguageHeaderName, "en-US,en")

	return headers
}
