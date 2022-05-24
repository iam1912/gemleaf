package internal

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/iam1912/gemseries/gemleaf/leaf"
	pb "github.com/iam1912/gemseries/gemleaf/proto"
	"gorm.io/gorm"
)

type GemLeafServer struct {
	Client *leaf.GemLeaf
}

func NewLeafServer(db *gorm.DB) *GemLeafServer {
	return &GemLeafServer{
		Client: leaf.NewGemLeaf(db),
	}
}

func (ser *GemLeafServer) GetLeafID(ctx context.Context, req *pb.GetLeafReq) (*pb.GetLeafResp, error) {
	if req.Name == "" {
		return nil, errors.New("invalid params")
	}
	id, err := ser.Client.GetID(req.Name)
	if err != nil {
		return nil, fmt.Errorf("failed get leaf id from %s\n", req.Name)
	}
	return &pb.GetLeafResp{Message: strconv.Itoa(int(id))}, nil
}

func (ser *GemLeafServer) CreateLeaf(ctx context.Context, req *pb.CreateLeafReq) (*pb.CreateLeafResq, error) {
	if req.Name == "" {
		return nil, errors.New("invalid params")
	}
	err := ser.Client.Create(req.Name, req.MaxId, int(req.Step), req.Desc)
	if err != nil {
		return nil, err
	}
	return &pb.CreateLeafResq{Message: "success create"}, nil
}
