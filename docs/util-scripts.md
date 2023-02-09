# Util Scripts

## minikube
* 作成/起動:
  ```sh
  minikube start
  ```
* 停止:
  ```sh
  minikube stop
  ```
* 削除:
  ```sh
  minikube delete
  ```
* ingress追加
  ```sh
  minikube addons enable ingress
  minikube addons enable ingress-dns
  ```

## ArgoCD
* インストール:
  ```sh
  kubectl create namespace argocd
  kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
  ```
* ポートフォワード:
  ```sh
  kubectl port-forward svc/argocd-server -n argocd 8080:443
  ```

## ArgoRollouts
* インストール:
  ```sh
  kubectl create namespace argo-rollouts
  kubectl apply -n argo-rollouts -f https://github.com/argoproj/argo-rollouts/releases/latest/download/install.yaml
  ```
* Plugin導入
  ```sh
  curl -LO https://github.com/argoproj/argo-rollouts/releases/latest/download/kubectl-argo-rollouts-linux-amd64
  chmod +x ./kubectl-argo-rollouts-linux-amd64
  sudo mv ./kubectl-argo-rollouts-linux-amd64 /usr/local/bin/kubectl-argo-rollouts
  ```
* 監視
  ```sh
  kubectl argo rollouts get rollout <Rollout> -n <Namespace> --watch
  ```
