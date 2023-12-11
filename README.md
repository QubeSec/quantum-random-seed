# quantum-random-seed

Deploy the quantum-random-seed to your cluster:
```bash
kubectl create deployment quantum-random-seed --image=qubesec/quantum-random-seed:v0.1.4
kubectl expose deployment quantum-random-seed --port=80 --name=quantum-random-seed
```

Deploy netshoot to your cluster:
```bash
kubectl run netshoot --image=nicolaka/netshoot -- sleep infinity
kubectl exec -it netshoot -- sh
```

Test your API server:
```bash
curl quantum-random-seed/?bytes=48 ; echo
```
