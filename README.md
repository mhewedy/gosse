
1. Build and push docker image to docker hub:

```bash
./build.sh
docker tag gosse:latest mhewedy/gosse:latest
docker push mhewedy/gosse:latest
```

2. Generate nginx ssl certificate to be used to enable ssl & http2:

```bash
openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout nginx-selfsigned.key -out nginx-selfsigned.crt
```

3. create Kubernetes objects:

```bash
kubectl apply -f k8s.yaml
```


Resources:
https://stackoverflow.com/questions/13672743/eventsource-server-sent-events-through-nginx
