package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"helloworld/internal/conf"
	"time"

	pb "helloworld/api/helloworld/v1"
)

type DemoService struct {
	pb.UnimplementedDemoServer
	flag *conf.Flag
}

func NewDemoService(f *conf.Flag) *DemoService {
	return &DemoService{flag: f}
}

func (s *DemoService) CreateDemo(ctx context.Context, req *pb.CreateDemoRequest) (*pb.CreateDemoReply, error) {
	return &pb.CreateDemoReply{}, nil
}
func (s *DemoService) UpdateDemo(ctx context.Context, req *pb.UpdateDemoRequest) (*pb.UpdateDemoReply, error) {
	return &pb.UpdateDemoReply{}, nil
}
func (s *DemoService) DeleteDemo(ctx context.Context, req *pb.DeleteDemoRequest) (*pb.DeleteDemoReply, error) {
	return &pb.DeleteDemoReply{}, nil
}
func (s *DemoService) GetDemo(ctx context.Context, req *pb.GetDemoRequest) (*pb.GetDemoReply, error) {
	return &pb.GetDemoReply{}, nil
}
func (s *DemoService) ListDemo(ctx context.Context, req *pb.ListDemoRequest) (*pb.ListDemoReply, error) {
	log.Log(log.LevelError, "ListDemo", time.Now().Format("20060102"))
	log.Log(log.LevelWarn, s.flag.GetMe())
	return &pb.ListDemoReply{}, nil
}
