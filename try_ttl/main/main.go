package main

import (
	"fmt"
	"time"
)

func main() {
	//sonos := ttl_service.TTSService{}
	//token := sonos.CreateToken()
	//token.Token
	n := name{}
	n.hh = time.Second
	fmt.Println(n.hh)
}

type name struct {
	hh time.Duration
}
