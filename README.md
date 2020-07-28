
Build and push docker image to dockerhub:

```bash
./build.sh
docker tag gosse:latest mhewedy/gosse:latest
docker push mhewedy/gosse:latest
```

Generate nginx ssl certificate:

```bash
openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout nginx-selfsigned.key -out nginx-selfsigned.crt
```


--Resources:
https://stackoverflow.com/questions/13672743/eventsource-server-sent-events-through-nginx
