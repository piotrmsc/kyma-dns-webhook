module github.com/kyma-incubator/kyma-dns-webhook/dns-proxy

go 1.13

require (
	github.com/go-acme/lego/v3 v3.2.0
	github.com/pkg/errors v0.8.1
	github.com/sirupsen/logrus v1.4.2
)

replace github.com/kyma-incubator/kyma-dns-webhook/dns-challenger => ./dns-challenger
