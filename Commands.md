# Useful Commands

## Create a local cluster using K3D

    k3d cluster create dev -p "80:80@loadbalancer" --registry-create dev-registry --k3s-arg="--disable=traefik@server:0" --agents 2

## Add User to CockRoachDB

    kubectl exec -it cockroachdb-client-secure -- ./cockroach sql --certs-dir=./cockroach-certs --host=cockroachdb-public
    CREATE USER ____ WITH PASSWORD '______';
    GRANT admin TO _____;
    \q
    kubectl port-forward service/cockroachdb-public 8080

    http://localhost:8080/

## Grafana Dashboard

    kubectl get secret --namespace default grafana-1663824153 -o jsonpath="{.data.admin-password}" | base64 --decode ; echo

    kubectl get pods --namespace default -l "app=prometheus,component=alertmanager"

    kubectl --namespace default port-forward $POD_NAME 3000

    http://localhost:3000/

    username = admin
    password = base64 decoded password

## Prometheus Dashboard

    kubectl get pods --namespace default -l "app=prometheus,component=server"

    kubectl --namespace default port-forward $POD_NAME 9090

## Pushing images to local k3D registry registry

    docker build -t localhost:$PORT/$NAME .

    docker push localhost:$PORT/$NAME

## Access CockRoachDB through local database visualizer

    kubectl port-forward service/cockroachdb-public 26257

## Install Istio

    kubectl create namespace istio-system

    helm install istio-base istio/base -n istio-system

    helm install istiod istio/istiod -n istio-system --wait

## Istio enable injection

    kubectl label namespace ____ istio-injection=enabled

## Install metrics server

    kubectl apply -f https://github.com/kubernetes-sigs/metrics-server/releases/latest/download/components.yaml
