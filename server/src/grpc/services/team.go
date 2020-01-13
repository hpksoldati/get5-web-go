package services

import (
	"fmt"
	// "github.com/FlowingSPDG/get5-web-go/server/src/api"
	"github.com/FlowingSPDG/get5-web-go/server/src/db"
	pb "github.com/FlowingSPDG/get5-web-go/server/src/grpc/proto"
	// "github.com/FlowingSPDG/get5-web-go/server/src/util"
	"context"
	// "google.golang.org/grpc"
	// "log"
)

func (s Server) RegisterTeam(ctx context.Context, req *pb.RegisterTeamRequest) (*pb.RegisterTeamReply, error) {
	team := &db.TeamData{}
	reqteam := req.GetTeam()
	team, err := team.Create(int(reqteam.GetUserid()), reqteam.GetName(), reqteam.GetTag(), reqteam.GetFlag(), reqteam.GetLogo(), reqteam.GetAuths(), reqteam.GetPublicteam())
	if err != nil {
		return &pb.RegisterTeamReply{
			Error:        true,
			Errormessage: "Failed to register team",
		}, err
	}
	resp := &pb.RegisterTeamReply{
		Error:        false,
		Errormessage: "",
	}
	return resp, nil
}

func (s Server) GetTeam(ctx context.Context, req *pb.GetTeamRequest) (*pb.GetTeamReply, error) {
	team := &db.TeamData{}
	rec := db.SQLAccess.Gorm.First(&team, req.GetId())
	if rec.RecordNotFound() {
		return &pb.GetTeamReply{}, fmt.Errorf("Team not found")
	}
	return &pb.GetTeamReply{
		Team: &pb.TeamData{
			Id:     int32(team.ID),
			Userid: int32(team.UserID),
			Name:   team.Name,
			Tag:    team.Tag,
			Flag:   team.Flag,
			Logo:   team.Logo,
			Auths:  team.Auths,
		},
	}, nil
}

func (s Server) EditTeam(ctx context.Context, req *pb.EditTeamRequest) (*pb.EditTeamReply, error) {
	return nil, nil // TODO
}

func (s Server) DeleteTeam(ctx context.Context, req *pb.DeleteTeamRequest) (*pb.DeleteTeamReply, error) {
	return nil, nil // TODO
}
