package services

import (
	"context"
	"io"
	"time"
)

type UserService struct {
}

func (*UserService) GetUserScore(ctx context.Context, in *UserScoreRequest) (*UserScoreResponse, error) {
	var score int32 = 101
	users := make([]*UserInfo, 0)
	for _, user := range in.Users {
		user.UserScore = score
		score++
		users = append(users, user)
	}
	return &UserScoreResponse{Users: users}, nil
}

//服务端流
func (*UserService) GetUserScoreServerStream(in *UserScoreRequest, stream UserService_GetUserScoreServerStreamServer) error {
	var score int32 = 101
	users := make([]*UserInfo, 0)
	for index, user := range in.Users {
		user.UserScore = score
		score++
		users = append(users, user)

		//每次发送两条数据
		if (index+1)%2 == 0 && index > 0 {
			err := stream.Send(&UserScoreResponse{Users: users})
			if err != nil {
				return err
			}
			//清空切片
			users = (users)[0:0]
			//模拟耗时
			time.Sleep(time.Second)
		}
	}

	if len(users) > 0 {
		err := stream.Send(&UserScoreResponse{Users: users})
		if err != nil {
			return err
		}
	}
	return nil
}

//客户端流
func (*UserService) GetUserScoreClientStream(stream UserService_GetUserScoreClientStreamServer) error {
	users := make([]*UserInfo, 0)
	var score int32 = 100
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			err = stream.SendAndClose(&UserScoreResponse{Users: users})
			return err
		}

		if err != nil {
			return err
		}

		for _, user := range req.Users {
			user.UserScore = score
			users = append(users, user)
			score++
		}
	}

}

//双向流
func (*UserService) GetUserScoreTwsStream(stream UserService_GetUserScoreTwsStreamServer) error {
	users := make([]*UserInfo, 0)
	var score int32 = 100
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			return err
		}

		for _, user := range req.Users {
			user.UserScore = score
			users = append(users, user)
			score++
		}
		
		stream.Send(&UserScoreResponse{Users: users})
		users = users[0:0]
	}
}
