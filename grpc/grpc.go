package grpc

import (
	"context"
	"net"
	"time"

	"github.com/baribari2/pulp-calculator/common/types"
	"github.com/baribari2/pulp-calculator/proto"
	"github.com/baribari2/pulp-calculator/simulator"
	"github.com/go-echarts/go-echarts/v2/charts"
	"google.golang.org/grpc"
)

type grpcServer struct {
	proto.CalcTreeServer
	Server    *grpc.Server
	Debate    *simulator.Debate
	Config    *types.Config
	LineChart *charts.Line
	Tick      time.Duration
	EndTime   time.Duration
	Frequency int64
}

func NewGrpcCalcServer(debate *simulator.Debate) *grpcServer {
	return &grpcServer{
		Debate: debate,
	}
}

func (s *grpcServer) Start() error {
	sv := grpc.NewServer()
	proto.RegisterCalcTreeServer(sv, s)

	s.Server = sv

	conn, err := net.Listen("tcp", ":8081")
	if err != nil {
		return err
	}

	return sv.Serve(conn)
}

func (s *grpcServer) GetCalcTree(ctx context.Context, req *proto.CalcTreeRequest) (*proto.CalcTreeResponse, error) {
	children := make([]*proto.Node, len(s.Debate.Root.Children))

	for i, c := range s.Debate.Root.Children {
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
			// Id:            int64(s.Debate.Root.Id),
			// ParentId:      int64(s.Debate.Root.ParentId),
			Confidence:    float32(s.Debate.Root.Confidence),
			Score:         s.Debate.Root.Score,
			LastScore:     s.Debate.Root.LastScore,
			InactiveCount: s.Debate.Root.InactiveCount,
			Children:      children,
		},
		Timestamps:    s.Debate.Timestamps,
		LastScore:     s.Debate.LastScore,
		InactiveCount: s.Debate.InactiveCount,
		Nodes:         nil,
	}, nil
}
