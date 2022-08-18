package main

import (
	"context"
	"fmt"
	"log"

	"github.com/batchcorp/natty"
)

// Connect to NATS
func main() {

	var natsUrl = []string{"tls://nats-url-goes-here"}
	var certFile = `
	-----BEGIN CERTIFICATE-----
	[yeet]
	-----END CERTIFICATE-----`

	n, err := natty.New(&natty.Config{
		NatsURL:       natsUrl,
		UseTLS:        true,
		TLSCACertFile: certFile,
	})

	if err != nil {
		log.Fatal(err)
	}
	keys, err := n.Keys(context.Background(), "bucketname")

	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(keys)

}
