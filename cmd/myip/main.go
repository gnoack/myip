package main

import (
	"context"
	"fmt"
	"log"

	"github.com/gnoack/myip"
)

func main() {
	ctx := context.Background()
	ip, err := myip.LookupV4(ctx)
	if err != nil {
		log.Fatalf("Looking up IP: %v", err)
	}
	fmt.Println(ip)
}
