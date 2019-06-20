package service

import (
	pb "github.com/luyaops/cmdb/proto"
	"golang.org/x/net/context"
	"net/http"
)

type CmdbAddTmpl struct {
}

func (c *CmdbAddTmpl) AliEcs(ctx context.Context, req *pb.AliEcsRequest) (*pb.AliEcsResponse, error) {
	return &pb.AliEcsResponse{
		Data: "AliEcsResponse...",
		Code: http.StatusOK,
	}, nil
}
