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
* CLI導入
  ```sh
  curl -sSL -o argocd-linux-amd64 https://github.com/argoproj/argo-cd/releases/latest/download/argocd-linux-amd64
  sudo install -m 555 argocd-linux-amd64 /usr/local/bin/argocd
  rm argocd-linux-amd64
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
* ダッシュボード表示
  ```sh
  kubectl argo rollouts dashboard
  ```

## Buildpacks / Docker
* インストール
  ```sh
  (curl -sSL "https://github.com/buildpacks/pack/releases/download/v0.28.0/pack-v0.28.0-linux.tgz" | sudo tar -C /usr/local/bin/ --no-same-owner -xzv pack)
  ```
* ビルド
  ```sh
  pack build hello-world --builder paketobuildpacks/builder:tiny --env BP_GO_TARGETS="./apps/hello-world" --env BP_GO_BUILD_LDFLAGS="-X main.version=<タグ名>"
  ```
* タグ付け
  ```sh
  docker tag <イメージID> registry.gitlab.com/h-meru/argorollouts-hands-on/hello-world:<タグ名>
  ```
* プッシュ
  ```sh
  docker push registry.gitlab.com/h-meru/argorollouts-hands-on/hello-world:<タグ名>
  ```

## Test
* curl実行
  ```sh
  watch -de -n 0.5 curl -sf <ドメイン名>.dev.sample.io
  ```
