# Lino Proxy

## single-site-proxy

```bash
docker run -it -d --name openai-proxy -p 10080:80 ghcr.io/linolabx/lino-proxy/single-site-proxy \
  --target https://api.openai.com \
  --proxy socks5://myproxy-host:1080
```
