# geo-ip
Http/grpc API to convert IP addresses from Maxmind database to countries or cities

## Setup for development

    minikube start
    eval $(minikube docker-env)
    skaffold dev

## metrics

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
