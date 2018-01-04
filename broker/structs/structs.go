package structs

import (
	"bytes"

	"github.com/ugorji/go/codec"
)

type MessageType uint8

const (
	RegisterRequestType   MessageType = 0
	DeregisterRequestType             = 1
)

type RegisterRequest struct {
	Node string
}

type DeregisterRequest struct {
	Node string
}

// msgpackHandle is a shared handle for encoding/decoding of structs
var msgpackHandle = &codec.MsgpackHandle{}

// Decode is used to encode a MsgPack object with type prefix.
func Decode(buf []byte, out interface{}) error {
	return codec.NewDecoder(bytes.NewReader(buf), msgpackHandle).Decode(out)
}

// Encode is used to encode a MsgPack object with type prefix
func Encode(t MessageType, msg interface{}) ([]byte, error) {
	var buf bytes.Buffer
	buf.WriteByte(uint8(t))
	err := codec.NewEncoder(&buf, msgpackHandle).Encode(msg)
	return buf.Bytes(), err
}

type RaftIndex struct {
	CreateIndex uint64
	ModifyIndex uint64
}

// Node is used to return info about a node
type Node struct {
	ID      string
	Node    string
	Address string
	Meta    map[string]string
	RaftIndex
}

// NodeService is a service provided by a node
type NodeService struct {
	ID      string
	Service string
	Tags    []string
	Address string
	Port    int
	RaftIndex
}

// HealthCheck represents a single check on a given node
type HealthCheck struct {
	Node        string
	CheckID     string
	Status      string
	Nodes       string
	ServiceID   string
	ServiceName string
	ServiceTags []string
	RaftIndex
}
