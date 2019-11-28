package dns

import (
	"github.com/jetstack/cert-manager/pkg/acme/webhook"
	"github.com/jetstack/cert-manager/pkg/acme/webhook/apis/acme/v1alpha1"
	log "github.com/sirupsen/logrus"
	"k8s.io/client-go/rest"
)

// Solver implements `github.com/jetstack/cert-manager/pkg/acme/webhook.Solver`
// In order to delegate DNS challenge to DNS client deployed in different namespace/k8s cluster.
type Solver struct {
	dnsClient DNSChallengeClient
}

// NewSolver returned initialied Solver
func NewSolver(dnsClient DNSChallengeClient) webhook.Solver {
	return &Solver{dnsClient: dnsClient}
}

//Name is used as the name for this DNS solver when referencing it on the ACME
// Issuer resource.
func (s *Solver) Name() string {
	return "kyma-dns"
}

func (s *Solver) Present(cr *v1alpha1.ChallengeRequest) error {
	log.Info("DNS Challenge Present ")
	err := s.dnsClient.Present(cr.DNSName, cr.Key)
	if err != nil {
		log.Error(err)
	}
	return err
}

func (s *Solver) CleanUp(cr *v1alpha1.ChallengeRequest) error {
	log.Info("DNS Challenge CleanUp ")
	err := s.dnsClient.CleanUp(cr.DNSName, cr.Key)
	if err != nil {
		log.Error(err)
	}
	return err
}

func (s *Solver) Initialize(kubeClientConfig *rest.Config, stopCh <-chan struct{}) error {

	return nil
}
