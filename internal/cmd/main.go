package main

import (
	impl "github.com/luyaops/cmdb/internal/impl/service"
	pb "github.com/luyaops/cmdb/proto"
	"github.com/luyaops/fw/common/log"
	"github.com/luyaops/fw/core"
)

func main() {
	server := core.NewRpcServer()
	pb.RegisterAddServer(server.Server, &impl.CmdbAddTmpl{})
	if err := server.Run(); err != nil {
		log.Fatalf("Failed to run rpc server: %v", err)
	}
}
