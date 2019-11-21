package main

import (
	"github.com/avast/retry-go"
	"github.com/jetstack/cert-manager/pkg/acme/webhook/cmd"
	"github.com/kyma-incubator/kyma-dns-webhook/dns-webhook/internal/dns"
	log "github.com/sirupsen/logrus"
	"net/url"
	"os"
	"time"
)

var GroupName = os.Getenv("GROUP_NAME")
var DNSChallengeAddress = os.Getenv("DNS_CHALLENGE_ADDRESS")

func main() {
	if GroupName == "" {
		panic("GROUP_NAME must be specified")
	}
	if DNSChallengeAddress == "" {
		panic("DNS_CHALLENGE_ADDRESS must be specified")
	}

	_, err := url.Parse(DNSChallengeAddress)
	if err != nil {
		panic("DNS_CHALLENGE_ADDRESS must be valid url")
	}

	//TODO make it configurable
	skipVerify := true
	retryOpts := []retry.Option{
		retry.Delay(2 * time.Second),
		retry.Attempts(3),
		retry.DelayType(retry.FixedDelay),
	}

	dnsClient := dns.NewDNSChallengeClient(DNSChallengeAddress, skipVerify, retryOpts)

	// This will register our custom DNS provider with the webhook serving
	// library, making it available as an API under the provided GroupName.
	// You can register multiple DNS provider implementations with a single
	// webhook, where the Name() method will be used to disambiguate between
	// the different implementations.

	log.Info("Starting WebHook server")
	cmd.RunWebhookServer(GroupName, dns.NewSolver(dnsClient))
}
