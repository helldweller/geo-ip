# geo-ip (in development)
Http/grpc API to convert IP addresses from Maxmind database to countries or cities

## Setup

### You need for development

* A kubernetes cluster (or minikube or docker desktop k8s)
* kubectl with kubeconfig to that cluster
* helm3
* skaffold

Commands for dev env

    # start
    skaffold dev

    # cleanup
    skaffold delete

## Maxmind database using

This https://hub.docker.com/r/maxmindinc/geoipupdate image is used here.

### Secrets

You must register a maxmind account and get account id and key for using in variables:

    GEOIPUPDATE_ACCOUNT_ID
    GEOIPUPDATE_LICENSE_KEY

And put to the secrets ./skaffold/secrets.yaml, see example ./skaffold/secrets-example.yaml

### Get last-modified database (for debug)

    curl -I 'https://download.maxmind.com/app/geoip_download?edition_id=GeoLite2-City&license_key=<KEY>&suffix=tar.gz'
    # or
    docker run --rm -it -e GEOIPUPDATE_LICENSE_KEY=<KEY> -e GEOIPUPDATE_ACCOUNT_ID=<ID> -e GEOIPUPDATE_EDITION_IDS=GeoLite2-City -v "C:\maxmind:/usr/share/GeoIP" maxmindinc/geoipupdate:latest

## REST

3 endpoints are available:

    /api/cities/{ip_addr}    - provides information about city
    /api/countries/{ip_addr} - provides information about country
    /api/summary/{ip_addr}   - provides both city, country and subdivisions

Also system endpoints:

    /metrics - prometheus metrics
    /healthcheck/startup   - startup probe
    /healthcheck/liveness  - liveness probe
    /healthcheck/readiness - readiness probe

## Metrics (if u want to visualisate)

    helm repo add prometheus https://prometheus-community.github.io/helm-charts
    helm repo update
    helm upgrade -n monitoring -i prometheus prometheus/kube-prometheus-stack \
        --version 30.2.0 \
        --create-namespace \
        --set grafana.enabled=true

    # helm show values prometheus/kube-prometheus-stack --version 30.2.0

    # To reset grafana admin password (exec to container)
    grafana-cli admin reset-admin-password <newpass>
    # default user: admin, password: prom-operator

    # forward ports
    kubectl port-forward -n monitoring  svc/prometheus-grafana 3000:3000
    kubectl port-forward -n monitoring  svc/prometheus-kube-prometheus-prometheus 9090:9090

    # GUI
    # grafana: http://127.0.0.1:3000
    # prometheus is also available on: http://127.0.0.1:9090

    # Grafana Go-lang dashboards examples
    https://grafana.com/grafana/dashboards/10826
    https://grafana.com/grafana/dashboards/14061

    # How to make custom metrics
    https://gabrieltanner.org/blog/collecting-prometheus-metrics-in-golang/
    https://prometheus.io/docs/guides/go-application/
