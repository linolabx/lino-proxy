package main

import (
	"flag"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var (
	listen    string
	target    string
	targetURL *url.URL
	proxy     string
	proxyURL  *url.URL
)

func init() {
	flag.StringVar(&listen, "listen", ":80", "port to listen on")
	flag.StringVar(&target, "target", "https://api.openai.com", "target website")
	flag.StringVar(&proxy, "proxy", "", "proxy server")
}

func ReverseProxyHandler(w http.ResponseWriter, r *http.Request) {
	rproxy := &httputil.ReverseProxy{}

	rproxy.Director = func(req *http.Request) {
		req.URL.Scheme = targetURL.Scheme
		req.URL.Host = targetURL.Host
		req.Host = targetURL.Host
	}

	if proxy != "" {
		rproxy.Transport = &http.Transport{Proxy: http.ProxyURL(proxyURL)}
	}

	rproxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		log.Printf("[!] %v: %v\n", r.URL, err)
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
	}

	rproxy.ServeHTTP(w, r)
}

func main() {
	flag.Parse()

	err := error(nil)

	if targetURL, err = url.Parse(target); err != nil {
		log.Fatalf("[!] Invalid target website: %s\n", target)
	}
	log.Printf("[*] Target website: %s\n", target)

	if proxy != "" {
		if proxyURL, err = url.Parse(proxy); err != nil {
			log.Fatalf("[!] Invalid proxy server: %s\n", proxy)
		}
		log.Printf("[*] Proxy server: %s\n", proxy)
	}

	log.Printf("[*] Starting server at port %v\n", listen)

	if err := http.ListenAndServe(listen, http.HandlerFunc(ReverseProxyHandler)); err != nil {
		log.Fatalf("[!] Failed to start server: %v\n", err)
	}
}
