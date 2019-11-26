# argocd-util-splitter
A utility to split up the superyaml that ArgoCD creates as a backup into individual manifests

### Build
`go build -o argocd-util-splitter`

### Run
`./argocd-util-splitter -src=/home/backup.yaml -dst=/tmp`

### Output Format:

{manifest_name}-{manifest_kind}.yaml

<i>myargoapplication-Application.yaml</i>
