FROM mcr.microsoft.com/devcontainers/base:bullseye

# ArgoCD CLI
ARG ARGOCD_VERSION
RUN curl -sSfL https://github.com/argoproj/argo-cd/releases/download/${ARGOCD_VERSION}/argocd-linux-amd64 -o /usr/local/bin/argocd \
    && chmod +x /usr/local/bin/argocd

# Argo Rollouts CLI
ARG ARGOROLLOUT_VERSION
RUN curl -sSfL https://github.com/argoproj/argo-rollouts/releases/download/${ARGOROLLOUT_VERSION}/kubectl-argo-rollouts-linux-amd64 -o /usr/local/bin/kubectl-argo-rollouts \
    && chmod +x /usr/local/bin/kubectl-argo-rollouts

# BuildPacks CLI
ARG BUILDPACKS_VERSION
RUN curl -sSfL https://github.com/buildpacks/pack/releases/download/${BUILDPACKS_VERSION}/pack-${BUILDPACKS_VERSION}-linux.tgz \
    | sudo tar -C /usr/local/bin/ --no-same-owner -xzv pack
