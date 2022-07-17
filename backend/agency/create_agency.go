package agency

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"time"

	"github.com/geraldkohn/elian/config"
	query "github.com/geraldkohn/elian/data/biz"
	"github.com/geraldkohn/elian/data/model"
	jwt "github.com/geraldkohn/elian/middleware/jwt"
	"github.com/geraldkohn/elian/middleware/uid"
	pb "github.com/geraldkohn/elian/proto/agency"
)

func (s *server) CreateAgency(ctx context.Context, in *pb.CreateAgencyRequest) (out *pb.CreateAgencyResponse, err error) {
	licenses := []config.License{}
	err = json.Unmarshal(config.LicenseData, &licenses)
	if err != nil {
		return &pb.CreateAgencyResponse{ErrorCode: 1, Token: "", Msg: "请重试"}, errors.New("后台解析json出错, error: " + err.Error())
	}

	for _, license := range licenses {
		if in.GetLicense() == license.LicenseCode {
			agencyRepo := query.NewAgencyRepo()
			uid := uid.NewUid()
			token, err := jwt.SignedToken(uid)
			if err != nil {
				log.Println("机构注册时生成token失败, " + err.Error())
				return &pb.CreateAgencyResponse{ErrorCode: 1, Token: "", Msg: "注册失败, 请重试"}, err
			}
			now := time.Now()

			agency := &model.Agency{
				Uid:           uid,
				License:       in.GetLicense(),
				RegisterTime:  jwt.FormatTime(now),
				LastLoginTime: jwt.FormatTime(now),
				Token:         token,
			}

			err = agencyRepo.Create(context.Background(), agency)
			if err != nil {
				log.Println("机构注册时出现错误, " + err.Error())
				return &pb.CreateAgencyResponse{ErrorCode: 1, Token: "", Msg: "注册失败, 请重试"}, err
			}

			return &pb.CreateAgencyResponse{ErrorCode: 0, Token: "", Msg: "注册成功"}, nil
		}
	}
	return &pb.CreateAgencyResponse{ErrorCode: 1, Token: "", Msg: "许可证不正确"}, errors.New("无许可证机构试图注册, license: " + in.GetLicense())
}
