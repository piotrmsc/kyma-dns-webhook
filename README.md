# Kyma DNS Webhook

The Kyma DNS Webhook is a component that can be used with [Cert Manager](https://cert-manager.netlify.com/).

//TODO @piotrmsc describe repo, configuration, how to build & test
## Build DNS webhook

 `docker build -f Dockerfile.webhook . `
 
## Deploy webhook on a cluster

Run this command to install webhook on a cluster, note it should be installed in the same namespace as cert-manager. In our case it's `istio-system` (at least for now)

``` helm install --namespace istio-system ./deploy/kyma-dns-webhook  --name kyma-dns-webhook ```
## Sample Issuer and Cert CR for cert-manager
 
 ```
 apiVersion: certmanager.k8s.io/v1alpha1
  kind: Issuer
  metadata:
    name: letsencrypt-staging
  spec:
    acme:
      email: admin@kyma-project.io
      server: https://acme-staging-v02.api.letsencrypt.org/directory
      privateKeySecretRef:
        name: letsencrypt-staging-account-key
      dns01:
        providers:
          - name: kyma-dns
            webhook:
              groupName: acme.kyma-project.io
              solverName: kyma-dns
                           
```                

```
apiVersion: certmanager.k8s.io/v1alpha1
kind: Certificate
metadata:
  name: letsencrypt-staging-cert  
spec:
  secretName: example-tls-cert
  commonName: example.com
  dnsNames:
  - example.com
  - www.example.com
  issuerRef:
    name: letsencrypt-staging
    kind: Issuer
  acme:
    config:
      - dns01:
          provider: kyma-dns
        domains:
          - *.example.com
          
```