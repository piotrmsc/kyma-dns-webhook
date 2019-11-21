module github.com/kyma-incubator/kyma-dns-webhook/dns-webhook

go 1.13

require (
	github.com/avast/retry-go v2.4.1+incompatible
	github.com/imdario/mergo v0.3.7 // indirect
	github.com/jetstack/cert-manager v0.8.0-alpha.0

	github.com/sirupsen/logrus v1.4.2
	k8s.io/client-go v11.0.0+incompatible
)

replace github.com/kyma-incubator/kyma-dns-webhook/dns-webhook => ./dns-webhook

replace k8s.io/client-go => k8s.io/client-go v0.0.0-20190413052642-108c485f896e

replace github.com/evanphx/json-patch => github.com/evanphx/json-patch v0.0.0-20190203023257-5858425f7550

replace github.com/miekg/dns => github.com/miekg/dns v0.0.0-20170721150254-0f3adef2e220
