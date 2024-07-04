# Use envsubst https://stackoverflow.com/a/56009991
# envsubst < secret.yaml.tpl > secret.yaml
---
apiVersion: v1
kind: Secret
metadata:
  name: external-dns-gcp-service-account
data:
  # Get it from distro-central repo.
  credentials.json: "$EXTERNAL_DNS_GCP_SERVICE_ACCOUNT"
