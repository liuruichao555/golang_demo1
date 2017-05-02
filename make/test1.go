package main

import (
	"fmt"
	google_protobuf1 "github.com/golang/protobuf/ptypes/timestamp"
)

type ChaincodeMessage struct {
	Timestamp *google_protobuf1.Timestamp `protobuf:"bytes,2,opt,name=timestamp" json:"timestamp,omitempty"`
	Payload   []byte                      `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	Txid      string                      `protobuf:"bytes,4,opt,name=txid" json:"txid,omitempty"`
	//Proposal  *SignedProposal             `protobuf:"bytes,5,opt,name=proposal" json:"proposal,omitempty"`
}

func main() {
	var str chan string
	fmt.Println("no make: ", str)

	str = make(chan string, 10)
	fmt.Println("make no value: ", str)

	str <- "one"
	fmt.Println("value one: ", str)

	<- str
	fmt.Println("out value: ", str)
}
