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
