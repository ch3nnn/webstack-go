/**
 * @Author: chentong
 * @Date: 2025/02/08 14:31
 */

package tools

import (
	"net/http"
)

var _ http.RoundTripper = (*httpTransport)(nil)

type httpTransport struct {
	transport http.RoundTripper

	userAgent string
}

func (h *httpTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Accept", "*/*")
	req.Header.Set("User-Agent", h.userAgent)
	return h.transport.RoundTrip(req)
}

func NewHTTPTransport(transport *http.Transport) http.RoundTripper {
	return &httpTransport{
		transport: transport,
		userAgent: RandomUserAgent(),
	}
}
