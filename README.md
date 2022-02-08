# warpd

## Process overview

```mermaid
graph TD;
    commit[New Commit to Branch]-->artifact[Build and publish Artifact];
    artifact-->exist{Environment Exists};
    exist-->|yes|deploy[Deploy to Environment]
    exist-->|no|createEnv[Create Environment]
```

### Env mapping configuration
```yaml
# warpd.yaml
#- branch: regex
#   cluster: name
#   envName: regexWithGroups
envMapping:
  - branch: *
    excludeBranches:
      - main
      - staging
    envName: test-\(1\)  
  - branch: main
    envName: production
```

## Deployer
- helm dep up
- helm install
