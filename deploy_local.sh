#!/bin/sh

. ./config.sh
. ./minikube_start.sh

set -e

if [ "$DEBUG" = "1" ]; then
  set -x
fi

cmd_installed() {
    if command -v $1 >/dev/null 2>&1; then
        echo 1
    else
        echo 0
    fi
}

echo_warning() {
  echo "\033[31m$1\033[m"
}

# Setup environment to build images inside minikube
eval `minikube docker-env -p $MINIKUBE_PROFILE`

export ADMIN_API_IMAGE_TAG="latest"
export ADMIN_UI_IMAGE_TAG="latest"
export K8S_MANAGER_IMAGE_TAG="latest"
export SDK_RELEASE_VERSION="v0.13.0"
export OPERATOR_SDK_INSTALLED=$(cmd_installed operator-sdk)

./scripts/replace_env_path.sh

if [ "$SKIP_BUILD" != "1" ]; then
    docker build -t konstellation/kre-admin-api:latest admin-api
    docker build -t konstellation/kre-admin-ui:latest admin-ui
    docker build -t konstellation/kre-k8s-manager:latest k8s-manager
    docker build -t konstellation/kre-runtime-api:latest runtime-api

    if [ "$OPERATOR_SDK_INSTALLED" = "1" ]; then
      cd operator && operator-sdk build konstellation/kre-operator:latest && cd ..
    fi
fi

echo "Create Namespace if not exist...\n"
kubectl create ns kre --dry-run -o yaml | kubectl apply -f -

echo "Init helm tiller...\n"
helm init --upgrade --wait

helm dep update helm/kre
helm upgrade \
  --wait --recreate-pods \
  --install ${DEPLOY_NAME} --namespace ${NAMESPACE} \
  helm/kre

./scripts/show_minikube_etc_hosts.sh $MINIKUBE_PROFILE

if [ "$OPERATOR_SDK_INSTALLED" != "1" ]; then
      echo_warning "\n\n\n¡¡¡¡¡WARNING: Operator SDK not installed. Operator image was not built!!!\n\n\n"
fi

echo "Done."
