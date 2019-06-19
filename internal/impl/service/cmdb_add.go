package service

import (
	pb "github.com/luyaops/cmdb/proto"
	"golang.org/x/net/context"
	"net/http"
)

type CmdbAddTmpl struct {
}

func (CmdbAddTmpl) AliEcs(context.Context, *pb.AliEcsRequest) (*pb.AliEcsResponse, error) {
	return &pb.AliEcsResponse{
		Data: "AliEcsResponse...",
		Code: http.StatusOK,
	}, nil
}
