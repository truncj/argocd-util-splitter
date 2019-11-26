# argocd-util-splitter
A utility to split up the superyaml that ArgoCD creates as a backup into individual manifests

### Build
`go build`

### Flags

```
Usage of ./argocd-util-splitter:
  -dst string
        output directory (default "/tmp/")
  -src string
        path of argocd backup yaml (default "/home/backup.yaml")
```

### Run
`./argocd-util-splitter -src=/home/backup.yaml -dst=/tmp`

### Output Format:

{manifest_name}-{manifest_kind}.yaml

<i>myargoapplication-Application.yaml</i>
