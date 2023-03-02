package sidecar

import (
	"context"
	"net"
	"time"

	"github.com/baribari2/pulp-calculator/common/types"
	"github.com/baribari2/pulp-calculator/proto"
	"github.com/baribari2/pulp-calculator/tree"
	"github.com/go-echarts/go-echarts/v2/charts"
	"google.golang.org/grpc"
)

// Sidecar is a gRPC server that serves the tree
type Sidecar struct {
	proto.CalcTreeServer
	Server    *grpc.Server
	Tree      *tree.Tree
	Config    *types.Config
	LineChart *charts.Line
	Tick      time.Duration
	EndTime   time.Duration
	Frequency int64
}

func NewSidecar(tree *tree.Tree) *Sidecar {
	return &Sidecar{
		Tree: tree,
	}
}

func (s *Sidecar) Start() error {
	sv := grpc.NewServer()
	proto.RegisterCalcTreeServer(sv, s)

	s.Server = sv

	conn, err := net.Listen("tcp", ":8081")
	if err != nil {
		return err
	}

	return sv.Serve(conn)
}

func (s *Sidecar) GetCalcTree(ctx context.Context, req *proto.CalcTreeRequest) (*proto.CalcTreeResponse, error) {
	children := make([]*proto.Node, len(s.Tree.Root.Children))

	for i, c := range s.Tree.Root.Children {
		children[i] = &proto.Node{
			// Id:            int64(c.Id),
			// ParentId:      int64(c.ParentId),
			Confidence:    float32(c.Confidence),
			Score:         c.Score,
			LastScore:     c.LastScore,
			InactiveCount: c.InactiveCount,
			Children:      nil,
		}
	}

	return &proto.CalcTreeResponse{
		Root: &proto.Node{
			// Id:            int64(s.Tree.Root.Id),
			// ParentId:      int64(s.Tree.Root.ParentId),
			Confidence:    float32(s.Tree.Root.Confidence),
			Score:         s.Tree.Root.Score,
			LastScore:     s.Tree.Root.LastScore,
			InactiveCount: s.Tree.Root.InactiveCount,
			Children:      children,
		},
		Timestamps:    s.Tree.Timestamps,
		LastScore:     s.Tree.LastScore,
		InactiveCount: s.Tree.InactiveCount,
		Nodes:         nil,
	}, nil
}
