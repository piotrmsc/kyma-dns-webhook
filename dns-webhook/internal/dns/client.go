package dns

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/avast/retry-go"
	"log"
	"net/http"
)

type DNSChallengeClient interface {
	Present(domain string, keyAuth string) error
	CleanUp(domain string, keyAuth string) error
}

type dnsChallengeClient struct {
	httpClient          *http.Client
	dnsChallengeAddress string
	retryOptions        []retry.Option
}

func NewDNSChallengeClient(dnsChallengeAddress string, skipVerify bool, retryOptions []retry.Option) DNSChallengeClient {
	return dnsChallengeClient{
		dnsChallengeAddress: dnsChallengeAddress,
		httpClient:          newHttpClient(skipVerify),
		retryOptions:        retryOptions,
	}
}

func newHttpClient(skipVerify bool) *http.Client {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: skipVerify},
	}
	client := &http.Client{Transport: tr}
	return client
}

func (c dnsChallengeClient) Present(domain string, keyAuth string) error {
	err := c.dnsChallengeReq("present", domain, keyAuth)

	if err != nil {
		log.Println("Error while communicating with dns-challenger : "+ err.Error())
	}
	log.Println("Finish present")
	return err
}

func (c dnsChallengeClient) CleanUp(domain string, keyAuth string) error {
	return c.dnsChallengeReq("cleanup", domain, keyAuth)
}

func (c dnsChallengeClient) dnsChallengeReq(action string, domain string, keyAuth string) error {
	reqBody, err := getReqBody(domain, keyAuth)
	if err != nil {
		return err
	}

	endpoint := fmt.Sprintf("%s/%s", c.dnsChallengeAddress, action)

	return c.postWithRetry(endpoint, reqBody)
}

func getReqBody(domain string, keyAuth string) ([]byte, error) {
	return json.Marshal(map[string]string{
		"domain":  domain,
		"keyAuth": keyAuth,
	})
}

func (c dnsChallengeClient) postWithRetry(endpoint string, body []byte) error {
	return retry.Do(func() error {

		resp, postErr := c.httpClient.Post(endpoint, "application/json", bytes.NewBuffer(body))
		log.Printf("status : %d ",resp.StatusCode)
		if postErr != nil {
			return postErr
		}

		return nil
	},
		c.retryOptions...,
	)
}
