package agency

import (
	"context"
	"encoding/json"
	"errors"
	"time"

	"github.com/geraldkohn/elian/config"
	query "github.com/geraldkohn/elian/data/biz"
	jwt "github.com/geraldkohn/elian/middleware/jwt"
	pb "github.com/geraldkohn/elian/proto/agency"
)

func (s *server) AgencyLogin(ctx context.Context, in *pb.AgencyLoginRequest) (out *pb.AgencyLoginResponse, err error) {
	uid, err := jwt.ParseToken(in.GetToken())

	//token失效, 验证license
	if err != nil {
		licenses := []config.License{}
		err = json.Unmarshal(config.LicenseData, &licenses)
		if err != nil {
			return &pb.AgencyLoginResponse{ErrorCode: 1, Token: "", Msg: "请重试"}, errors.New("后台解析json出错, error: " + err.Error())
		}

		for _, license := range licenses {
			if in.GetLicense() == license.LicenseCode {
				agencyRepo := query.NewAgencyRepo()
				a, err := agencyRepo.SelectByLicense(context.Background(), in.GetLicense())
				if err != nil {
					return &pb.AgencyLoginResponse{ErrorCode: 1, Token: "", Msg: "机构不存在"}, errors.New("未注册的机构试图登录, license: " + in.GetLicense())
				}

				newToken, _ := jwt.SignedToken(a.Uid)
				agencyRepo.UpdateTokenAndLoginTime(context.Background(), a.Uid, newToken, jwt.FormatTime(time.Now()))
				return &pb.AgencyLoginResponse{ErrorCode: 0, Token: newToken, Msg: "登录成功"}, nil
			}
		}

		return &pb.AgencyLoginResponse{ErrorCode: 1, Token: "", Msg: "许可证不正确"}, errors.New("无许可证机构试图登录, license: " + in.GetLicense())
	}

	agencyRepo := query.NewAgencyRepo()
	newToken, _ := jwt.SignedToken(uid)
	agencyRepo.UpdateTokenAndLoginTime(context.Background(), uid, newToken, jwt.FormatTime(time.Now()))
	return &pb.AgencyLoginResponse{ErrorCode: 0, Token: newToken, Msg: "登录成功"}, nil
}
