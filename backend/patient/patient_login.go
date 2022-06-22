package patient

import (
	"context"
	"errors"
	"time"

	query "github.com/geraldkohn/elian/data/biz"
	jwt "github.com/geraldkohn/elian/middleware/jwt"
	pb "github.com/geraldkohn/elian/proto/patient"
)

func (s *server) PatientLogin(ctx context.Context, in *pb.PatientLoginRequest) (*pb.PatientLoginResponse, error) {
	uid, err := jwt.ParseToken(in.Token)

	// TODO 令牌失效, 需要使用用户身份证号和密码确认用户身份, 并且返回新的令牌
	if err != nil {
		patientRepo := query.NewPatientRepo()
		p, err := patientRepo.SelectByIdCardNumber(context.Background(), in.IdCardNumber)
		if err != nil {
			return &pb.PatientLoginResponse{ErrorCode: 1, Token: "", Msg: "用户不存在"}, errors.New("不存在的用户试图登录. " + "IdCardNumber: " + in.IdCardNumber)
		}

		// TODO 用户身份确认, 生成新的令牌
		if p.Password == in.Password {
			newToken, err := jwt.SignedToken(p.Uid)
			if err != nil {
				return &pb.PatientLoginResponse{ErrorCode: 1, Token: "", Msg: "请重新登录"}, errors.New("患者登录时生成token错误, 错误信息为: " + err.Error())
			}
			newLoginTime := jwt.FormatTime(time.Now())
			err = patientRepo.UpdateTokenAndLoginTime(context.Background(), p.Uid, newToken, newLoginTime)
			if err != nil {
				return &pb.PatientLoginResponse{ErrorCode: 1, Token: "", Msg: "请重新登录"}, errors.New("患者登录时无法更新token和最新登录时间, 错误信息为: " + err.Error())
			}
			return &pb.PatientLoginResponse{ErrorCode: 0, Token: newToken, Msg: "登录成功"}, nil
		}
	}

	// TODO 令牌未失效, 更新用户最后一次登录时间和令牌
	patientRepo := query.NewPatientRepo()
	newToken, err := jwt.SignedToken(uid)
	if err != nil {
		return &pb.PatientLoginResponse{ErrorCode: 1, Token: "", Msg: "请重新登录"}, errors.New("患者登录时生成token错误, 错误信息为: " + err.Error())
	}
	newLoginTime := jwt.FormatTime(time.Now())
	err = patientRepo.UpdateTokenAndLoginTime(context.Background(), uid, newToken, newLoginTime)
	if err != nil {
		return &pb.PatientLoginResponse{ErrorCode: 1, Token: "", Msg: "请重新登录"}, errors.New("患者登录时无法更新token和最新登录时间, 错误信息为: " + err.Error())
	}
	return &pb.PatientLoginResponse{ErrorCode: 0, Token: newToken, Msg: "登录成功"}, nil
}
