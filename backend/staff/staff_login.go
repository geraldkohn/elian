package staff

import (
	"context"
	"errors"
	"time"

	query "github.com/geraldkohn/elian/data/biz"
	jwt "github.com/geraldkohn/elian/middleware/jwt"
	pb "github.com/geraldkohn/elian/proto/staff"
)

func (s *server) StaffLogin(ctx context.Context, in *pb.StaffLoginRequest) (out *pb.StaffLoginResponse, err error) {
	uid, err := jwt.ParseToken(in.Token)

	//token失效
	if err != nil {
		staffRepo := query.NewStaffRepo()
		p, err := staffRepo.SelectByIdCardNumber(context.Background(), in.IdCardNumber)
		if err != nil {
			return &pb.StaffLoginResponse{ErrorCode: 1, Token: "", Msg: "用户不存在"}, errors.New("不存在的用户试图登录. " + "IdCardNumber: " + in.IdCardNumber)
		}
		if p.Password == in.Password {
			newToken, err := jwt.SignedToken(p.Uid)
			if err != nil {
				return &pb.StaffLoginResponse{ErrorCode: 1, Token: "", Msg: "请重新登录"}, errors.New("患者登录时生成token错误, 错误信息为: " + err.Error())
			}
			newLoginTime := jwt.FormatTime(time.Now())
			err = staffRepo.UpdateTokenAndLoginTime(context.Background(), p.Uid, newToken, newLoginTime)
			if err != nil {
				return &pb.StaffLoginResponse{ErrorCode: 1, Token: "", Msg: "请重新登录"}, errors.New("患者登录时无法更新token和最新登录时间, 错误信息为: " + err.Error())
			}
			return &pb.StaffLoginResponse{ErrorCode: 0, Token: newToken, Msg: "登录成功"}, nil
		}
	}

	staffRepo := query.NewStaffRepo()
	newToken, err := jwt.SignedToken(uid)
	if err != nil {
		return &pb.StaffLoginResponse{ErrorCode: 1, Token: "", Msg: "请重新登录"}, errors.New("患者登录时生成token错误, 错误信息为: " + err.Error())
	}
	newLoginTime := jwt.FormatTime(time.Now())
	err = staffRepo.UpdateTokenAndLoginTime(context.Background(), uid, newToken, newLoginTime)
	if err != nil {
		return &pb.StaffLoginResponse{ErrorCode: 1, Token: "", Msg: "请重新登录"}, errors.New("患者登录时无法更新token和最新登录时间, 错误信息为: " + err.Error())
	}
	return &pb.StaffLoginResponse{ErrorCode: 0, Token: newToken, Msg: "登录成功"}, nil
}
