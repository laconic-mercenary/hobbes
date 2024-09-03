package util

import (
	"bytes"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"regexp"
	"strings"
	"time"

	log "github.com/rs/zerolog/log"
)

const (
	// owasp regex for URL
	URL_REGEX                    = "^(((https?)://)(%[0-9A-Fa-f]{2}|[-()_.!~*';/?:@&=+$,A-Za-z0-9])+)([).!';/?:,])?$"
	HTTP_CLIENT_TIMEOUT_SEC      = 30
	HTTP_IDLE_CONN_TIMEOUT       = 30
	HTTP_MAX_IDLE_CONNS          = 5
	HTTP_MAX_IDLE_CONNS_PER_HOST = 3
	HTTP_MAX_CONNS_PER_HOST      = 5
	HTTP_DISABLE_COMPRESSION     = true
	HTTP_DISABLE_KEEPALIVE       = false
	HTTP_POST_CONTENTTYPE        = "application/json"
	RAND_MULTIPLIER              = 15.0
)

var (
	isValidUrl = regexp.MustCompile(URL_REGEX).MatchString
)

type HttpClient struct {
	client *http.Client
	server string
}

func NewClient(serverUrl string) (httpClient HttpClient, err error) {
	log.Trace().Msg("util.NewClient")
	var revisedUrl string
	if revisedUrl, err = filterServerUrl(serverUrl); err == nil {
		log.Debug().Str("targetServer", revisedUrl).Msg("creating client")
		httpClient = HttpClient{
			server: revisedUrl,
			client: createClient(),
		}
	}
	return
}

func RandomRequestId() float64 {
	return (rand.Float64() * RAND_MULTIPLIER)
}

func RxToErr(statusCode int) (err error) {
	if statusCode > http.StatusIMUsed {
		err = errors.New(fmt.Sprintf("Status Code is %d", statusCode))
	}
	return
}

func RxToErr2(statusCode int, upstream error) (err error) {
	err = upstream
	if err == nil {
		err = RxToErr(statusCode)
	}
	return
}

func (this *HttpClient) Get(target string) (rx *http.Response, err error) {
	if prefixed := strings.HasPrefix(target, "/"); !prefixed {
		target = "/" + target
	}
	url := fmt.Sprintf("%s%s", this.server, target)
	rx, err = this.client.Get(url)
	return
}

func (this *HttpClient) Post(target string, data []byte) (rx *http.Response, err error) {
	if prefixed := strings.HasPrefix(target, "/"); !prefixed {
		target = "/" + target
	}
	url := fmt.Sprintf("%s%s", this.server, target)
	rx, err = this.client.Post(url, HTTP_POST_CONTENTTYPE, bytes.NewBuffer(data))
	return
}

func createClient() *http.Client {
	log.Trace().Msg("util.createClient")
	timeout := HTTP_CLIENT_TIMEOUT_SEC * time.Second
	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.MaxIdleConns = HTTP_MAX_IDLE_CONNS
	transport.MaxIdleConnsPerHost = HTTP_MAX_IDLE_CONNS_PER_HOST
	transport.MaxConnsPerHost = HTTP_MAX_CONNS_PER_HOST
	transport.DisableCompression = HTTP_DISABLE_COMPRESSION
	transport.DisableKeepAlives = HTTP_DISABLE_KEEPALIVE
	return &http.Client{
		Timeout:   timeout,
		Transport: transport,
	}
}

func filterServerUrl(serverUrl string) (result string, err error) {
	log.Trace().Msg("util.filterServerUrl")
	log.Debug().Str("serverUrl", serverUrl).Str("context", "util.filterServerUrl").Msg("validating serverUrl...")
	serverUrl = strings.TrimSpace(serverUrl)
	if valid := isValidUrl(serverUrl); !valid {
		err = errors.New(fmt.Sprintf("invalid base server URL: %s", serverUrl))
		log.Error().Str("context", "util.filterServerUrl").Str("serverUrl", serverUrl).Err(err)
	} else {
		result = serverUrl
		if suffixed := strings.HasSuffix(result, "/"); suffixed {
			result = result[:len(serverUrl)-1]
		}
	}
	return
}
