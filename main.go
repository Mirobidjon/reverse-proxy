package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
)

var configFile = flag.String("config", "config.yaml", "path to the config file")

type cfg struct {
	Path      string `yaml:"path"`
	ProxyPass string `yaml:"proxy_pass"`
}

type config struct {
	Port  string `yaml:"port"`
	Proxy []cfg  `yaml:"proxy"`
}

func init() {
	flag.Usage = func() {
		flag.PrintDefaults()
	}

	flag.Parse()

	if *configFile == "" {
		flag.Usage()
		os.Exit(1)
	}
}

func getProxy(path, remoteUrl string) func(c *gin.Context) {
	return func(c *gin.Context) {
		remote, err := url.Parse(remoteUrl)
		if err != nil {
			panic(err)
		}

		proxy := httputil.NewSingleHostReverseProxy(remote)
		proxy.Director = func(req *http.Request) {
			req.Header = c.Request.Header
			req.Host = remote.Host
			req.URL.Scheme = remote.Scheme
			req.URL.Host = remote.Host
			req.URL.Path = strings.TrimSuffix(remote.Path, "/") + strings.TrimPrefix(c.FullPath(), path)

			fmt.Printf("\nreq.Url: %s\n", req.URL.Path)
		}

		proxy.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	r := gin.Default()

	var conf config
	yfile, err := ioutil.ReadFile(*configFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "could not load config: %v\n", err)
		os.Exit(1)
	}

	err = yaml.Unmarshal(yfile, &conf)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	fmt.Printf("%+v", conf)

	for _, val := range conf.Proxy {
		r.Any(val.Path, getProxy(val.Path, val.ProxyPass))
	}

	if err := r.Run(":" + conf.Port); err != nil {
		log.Printf("Server exited with error: %v", err)
		os.Exit(1)
	}
}
