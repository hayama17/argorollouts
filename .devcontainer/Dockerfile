FROM mcr.microsoft.com/devcontainers/base:bullseye

# ArgoCD CLI
RUN curl -sSfL https://github.com/argoproj/argo-cd/releases/latest/download/argocd-linux-amd64 -o /usr/local/bin/argocd \
    && chmod +x /usr/local/bin/argocd

# Argo Rollouts CLI
RUN curl -sSfL https://github.com/argoproj/argo-rollouts/releases/latest/download/kubectl-argo-rollouts-linux-amd64 -o /usr/local/bin/kubectl-argo-rollouts \
    && chmod +x /usr/local/bin/kubectl-argo-rollouts

# BuildPacks CLI
RUN curl -sSfL https://github.com/buildpacks/pack/releases/download/v0.28.0/pack-v0.28.0-linux.tgz \
    | sudo tar -C /usr/local/bin/ --no-same-owner -xzv pack