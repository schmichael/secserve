// secserve is a simple https server for serving static content along with their hash
package main

import (
	"flag"
	//"net/http"
	"os"
	"log"
	"net"
	"strings"

	"github.com/schmichael/secserve/util"
)

var (
	path = flag.String("path", ".", "path to serve files from")
	https = flag.String("https", "0.0.0.0:8000", "host:port to bind https server on")
	cert = flag.String("cert", "", "path to certificate pem file")
	key = flag.String("key", "", "path to private key pem file")
	hostnames = flag.String("names", "", "comma seperatd list of valid hostnames")
)

func tmpCert(path string) (string, string, error) {
	hosts := []string{}
	if *https != "" {
		ip, _, err := net.SplitHostPort(*https)
		if err != nil {
			return "", "", err
		}
		hosts = append(hosts, ip)
	}
	if *hostnames != "" {
		hosts = append(hosts, strings.Split(*hostnames, ",")...)
	}

	cert, key, err := util.GenCert(hosts, time.Duration(-1 << 63), false)

	//FIXME Run GenCert and return valid stuffs
	return "", "", nil
}

func main() {
	flag.Parse()

	if *cert == "" || *key == "" {
		// No cert or key specified, create temporary ones
		tmpPath, err := util.SecTempDir()
		if err != nil {
			log.Fatalf("Error creating temporary directory for certificate: %v", err)
		}
		*cert, *key, err = tmpCert(tmpPath)
		if err != nil {
			log.Fatalf("Error generating temporary certificate: %v", err)
		}
		defer func() {
			if err := os.RemoveAll(tmpPath); err != nil {
				log.Printf("Error removing temporary path %s: %v", tmpPath, err)
			}
		}()
	}
}
