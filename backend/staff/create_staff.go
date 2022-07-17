package staff

import (
	"context"
	"errors"
	"time"

	query "github.com/geraldkohn/elian/data/biz"
	"github.com/geraldkohn/elian/data/model"
	check "github.com/geraldkohn/elian/middleware/format"
	jwt "github.com/geraldkohn/elian/middleware/jwt"
	"github.com/geraldkohn/elian/middleware/uid"
	pb "github.com/geraldkohn/elian/proto/staff"
)

func (s *server) CreateStaff(ctx context.Context, in *pb.CreateStaffRequset) (out *pb.CreateStaffResponse, err error) {
	if !check.VerifyIdCardNumberAndNameFormat(in.IdCardNumber, in.Name) {
		return &pb.CreateStaffResponse{ErrorCode: 1, Token: "", Msg: "身份证与名称不匹配"}, errors.New("创建医生账号时, 身份者格式和名称不正确或者不匹配. " + "IdCardNumber: " + in.IdCardNumber + " name: " + in.Name)
	}
	if !check.VerifyPasswordFormat(in.Password) {
		return &pb.CreateStaffResponse{ErrorCode: 1, Token: "", Msg: "密码格式不正确"}, errors.New("创建医生账号时, 密码格式不正确." + "IdCardNumber: " + in.IdCardNumber + " name: " + in.Name + " password: " + in.Password)
	}

	now := time.Now()
	uid := uid.NewUid()
	token, err := jwt.SignedToken(uid)
	if err != nil {
		return &pb.CreateStaffResponse{ErrorCode: 1, Token: "", Msg: "生成私钥出错"}, err
	}

	staff := &model.Staff{
		Uid:           uid,
		Name:          in.GetName(),
		Hospital:      in.GetHospital(),
		Department:    in.GetDepartment(),
		JobNumber:     in.GetJobNumber(),
		IdCardNumber:  in.GetIdCardNumber(),
		RegisterTime:  jwt.FormatTime(now),
		LastLoginTime: jwt.FormatTime(now),
		Token:         token,
	}

	staffRepo := query.NewStaffRepo()
	err = staffRepo.Create(context.Background(), staff)
	if err != nil {
		return &pb.CreateStaffResponse{ErrorCode: 1, Token: "", Msg: "注册失败, 未知错误"}, err
	}

	return &pb.CreateStaffResponse{ErrorCode: 0, Token: token, Msg: "注册成功"}, nil
}
