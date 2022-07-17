package record

import (
	"context"
	"errors"
	"log"

	query "github.com/geraldkohn/elian/data/biz"
	jwt "github.com/geraldkohn/elian/middleware/jwt"
	pb "github.com/geraldkohn/elian/proto/record"
)

func (s *server) AgencySetReadOnlyPermission(ctx context.Context, in *pb.AgencySetReadOnlyPermissionRequest) (out *pb.AgencySetReadOnlyPermissionResponse, err error) {

	// TODO 验证令牌身份
	agencyUid, err := jwt.ParseToken(in.GetAgencyToken())
	if err != nil {
		return &pb.AgencySetReadOnlyPermissionResponse{
			ErrorCodeAndInfo: &pb.ErrorCodeAndInfo{
				ErrorCode: 1,
				Msg:       "请重新登录",
			},
		}, errors.New("tokenInvalidError")
	}
	log.Println("机构uid: " + agencyUid + " 设置只读权限.")

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

	// TODO 在病历记录中添加医生权限(只读)
	err = addROnlyPermission(staffUids, in.GetRecordUid())
	if err != nil {
		return &pb.AgencySetReadOnlyPermissionResponse{
			ErrorCodeAndInfo: &pb.ErrorCodeAndInfo{
				ErrorCode: 1,
				Msg:       "请重试",
			},
		}, err
	}

	return &pb.AgencySetReadOnlyPermissionResponse{
		ErrorCodeAndInfo: &pb.ErrorCodeAndInfo{
			ErrorCode: 0,
			Msg:       "设置权限成功",
		},
	}, nil
}