# REVERSE PROXY
# ============

### This is a reverse proxy for the web application. It is used to
### redirect all requests to the web application to the web application
### server. It is also used to redirect all requests to the web application
### server to the web application.

# Here is the `docker-compose.yml` that powers the whole setup.

```yaml
version: '3.7'

services:
  reverse-proxy:
    image: us.gcr.io/learn-cloud-0809/reverse-proxy:latest
    container_name: reverse-proxy
    restart: always
    ports:
      - 8080:8080
    volumes:
      - ./proxy.yaml:/etc/reverse-proxy/proxy.yaml
    command: --config=/etc/reverse-proxy/proxy.yaml

```


# create yaml file with the following content
```yaml
port: 8080
proxy:
- path: /middleware
  proxy_pass: https://docs.gofiber.io/api/middleware/proxy
- path: /proxy
  proxy_pass: https://zetcode.com/golang/fiber/
- path: /cfg/style.css
  proxy_pass: https://zetcode.com/cfg/style.css
- path: /favicon.ico
  proxy_pass: https://zetcode.com/favicon.ico
- path: /terraform
  proxy_pass: https://ifconfig.me
- path: /styles/style.css
  proxy_pass: https://ifconfig.me/styles/style.css
```
