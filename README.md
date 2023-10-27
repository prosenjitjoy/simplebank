# simplebank
Backend master class [Golang + Postgres + Kubernetes + gRPC] 

## Run
### kubernetes cluster
```bash
kubectl apply -f simplebank-pod.yaml
kubectl port-forward simplebank 3000:3000
```
   
To view api specific log, run the following in a new terminal session.
```bash
kubectl logs -c simplebank-api -f simplebank
```
### dev environment
```bash
make dev_deploy
```

## Docs
https://dbdocs.io/prosenjitjoy/SimpleBank
http://localhost:3000/doc/swagger