package main

import (
	"context"
	"errors"
	"time"

	query "github.com/geraldkohn/elian/data/biz"
	"github.com/geraldkohn/elian/data/model"
	check "github.com/geraldkohn/elian/middleware/format"
	jwt "github.com/geraldkohn/elian/middleware/jwt"
	"github.com/geraldkohn/elian/middleware/uid"
	pb "github.com/geraldkohn/elian/proto/patient"
)

func (s *server) CreatePatient(ctx context.Context, in *pb.CreatePatientRequest) (*pb.CreatePatientResponse, error) {
	if !check.VerifyIdCardNumberAndNameFormat(in.IdCardNumber, in.Name) {
		return &pb.CreatePatientResponse{ErrorCode: 1, Token: "", Msg: "身份证与名称不匹配"}, errors.New("创建患者账号时, 身份者格式和名称不正确或者不匹配. " + "IdCardNumber: " + in.IdCardNumber + " name: " + in.Name)
	}
	if !check.VerifyPasswordFormat(in.Password) {
		return &pb.CreatePatientResponse{ErrorCode: 1, Token: "", Msg: "密码格式不正确"}, errors.New("创建患者账号时, 密码格式不正确." + "IdCardNumber: " + in.IdCardNumber + " name: " + in.Name + " password: " + in.Password)
	}

	now := time.Now()
	uid := uid.NewUid()
	token, err := jwt.SignedToken(uid)
	if err != nil {
		return &pb.CreatePatientResponse{ErrorCode: 1, Token: "", Msg: "生成私钥出错"}, err
	}

	p := &model.Patient{
		Uid:           uid,
		Name:          in.Name,
		Password:      in.Password,
		IdCardNumber:  in.IdCardNumber,
		RegisterTime:  jwt.FormatTime(now),
		LastLoginTime: jwt.FormatTime(now),
		Token:         token,
	}

	patientRepo := query.NewPatientRepo()
	err = patientRepo.Create(context.Background(), p)
	if err != nil {
		return &pb.CreatePatientResponse{ErrorCode: 1, Token: "", Msg: "注册失败, 未知错误"}, err
	}

	return &pb.CreatePatientResponse{ErrorCode: 0, Token: token, Msg: "注册成功"}, nil
}
