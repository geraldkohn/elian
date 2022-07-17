package record

import (
	"context"
	"errors"
	"log"

	query "github.com/geraldkohn/elian/data/biz"
	jwt "github.com/geraldkohn/elian/middleware/jwt"
	pb "github.com/geraldkohn/elian/proto/record"
)

func (s *server) StaffSetReadAndWritePermission(ctx context.Context, in *pb.StaffSetReadAndWritePermissionRequest) (out *pb.StaffSetReadAndWritePermissionResponse, err error) {
	// TODO 验证令牌身份
	staffUid, err := jwt.ParseToken(in.GetStaffToken())
	if err != nil {
		return &pb.StaffSetReadAndWritePermissionResponse{
			ErrorCodeAndInfo: &pb.ErrorCodeAndInfo{
				ErrorCode: 1,
				Msg:       "请重新登录",
			},
		}, errors.New("tokenInvalidError")
	}
	log.Println("医生uid: " + staffUid + " 设置读写权限.")

	// TODO 查找对应的医生信息
	staffRepo := query.NewStaffRepo()
	var staffUids []string
	for _, req := range in.GetPermissionRequests() {
		staff, err := staffRepo.SelectByHospitalAndDepartmentAndName(context.Background(), req.Hospital, req.Department, req.Name)
		if err != nil {
			_ = staff
			continue
		}
		staffUids = append(staffUids, staff.Uid)
	}

	// TODO 在病历记录中添加医生权限(读写)
	err = addRWPermission(staffUids, in.GetRecordUid())
	if err != nil {
		return &pb.StaffSetReadAndWritePermissionResponse{
			ErrorCodeAndInfo: &pb.ErrorCodeAndInfo{
				ErrorCode: 1,
				Msg:       "请重试",
			},
		}, err
	}

	return &pb.StaffSetReadAndWritePermissionResponse{
		ErrorCodeAndInfo: &pb.ErrorCodeAndInfo{
			ErrorCode: 0,
			Msg:       "设置权限成功",
		},
	}, nil
} 