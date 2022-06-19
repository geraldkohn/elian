package main

import (
	"context"
	"errors"
	"log"

	fabricQuery "github.com/geraldkohn/elian/data/fabric"
	jwt "github.com/geraldkohn/elian/middleware/jwt"
	pb "github.com/geraldkohn/elian/proto/record"
)

func (s *server) StaffUpdateDescription(ctx context.Context, in *pb.StaffUpdateDescriptionRequest) (out *pb.StaffUpdateDescriptionResponse, err error) {

	// TODO 验证令牌身份
	staffUid, err := jwt.ParseToken(in.StaffToken)
	if err != nil {
		return &pb.StaffUpdateDescriptionResponse{
			ErrorCodeAndInfo: &pb.ErrorCodeAndInfo{
				ErrorCode: 1,
				Msg:       "请重新登录",
			},
		}, errors.New("tokenInvalidError")
	}
	log.Print("医生uid" + staffUid + "更新病历简介")

	// TODO 判断是否有写权限
	ok := judgeStaffRWPermission(staffUid, in.RecordUid)
	if !ok {
		return &pb.StaffUpdateDescriptionResponse{
			ErrorCodeAndInfo: &pb.ErrorCodeAndInfo{
				ErrorCode: 1,
				Msg:       "无修改权限",
			},
		}, errors.New("noPermissionError")
	}

	recordRepo := fabricQuery.NewRecordRepo()
	err = recordRepo.UpdateDescriptionByUid(context.Background(), in.RecordUid, in.Description)
	if err != nil {
		return &pb.StaffUpdateDescriptionResponse{
			ErrorCodeAndInfo: &pb.ErrorCodeAndInfo{
				ErrorCode: 1,
				Msg:       "请重试",
			},
		}, err
	}

	return &pb.StaffUpdateDescriptionResponse{
		ErrorCodeAndInfo: &pb.ErrorCodeAndInfo{
			ErrorCode: 0,
			Msg:       "修改成功",
		},
	}, nil
}
