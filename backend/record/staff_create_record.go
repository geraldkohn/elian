package record

import (
	"context"
	"errors"
	"time"

	query "github.com/geraldkohn/elian/data/biz"
	fabricQuery "github.com/geraldkohn/elian/data/fabric"
	"github.com/geraldkohn/elian/data/model"
	jwt "github.com/geraldkohn/elian/middleware/jwt"
	uuid "github.com/geraldkohn/elian/middleware/uid"
	pb "github.com/geraldkohn/elian/proto/record"
)

func (s *server) StaffCreateRecord(ctx context.Context, in *pb.StaffCreateRecordRequest) (out *pb.StaffCreateRecordResponse, err error) {

	staffUid, err := jwt.ParseToken(in.StaffToken)

	// TODO 令牌认证失败, 需要重新登录.
	if err != nil {
		return &pb.StaffCreateRecordResponse{
			ErrorCodeAndInfo: &pb.ErrorCodeAndInfo{
				ErrorCode: 1,
				Msg:       "请重新登录",
			},
		}, errors.New("staffTokenAuthenticationFailed")
	}

	// TODO 令牌认证成功, 开始创建患者病历记录

	// TODO 根据患者身份证号查询患者uid, 格式化可读可写的医生的uid
	patientRepo := query.NewPatientRepo()
	patient, err := patientRepo.SelectByIdCardNumber(context.Background(),  in.GetRecordRequest().GetPatientIdCardNumber())
	if err != nil {
		return &pb.StaffCreateRecordResponse{
			ErrorCodeAndInfo: &pb.ErrorCodeAndInfo{
				ErrorCode: 1,
				Msg:       "无患者记录",
			},
		}, errors.New("ErrorNoPatientRecord")
	}
	recordUid := uuid.NewUid()
	rwStaffUid, _ := marshal("", staffUid)
	rOnlyStaffUid, _ := marshal("", "")

	// TODO 添加patientUid到recordUid的映射
	uidmapRepo := query.NewUidmapRepo()
	uidmapUid := uuid.NewUid()
	err = uidmapRepo.Create(context.Background(), &model.PatientuidToRecorduid{
		Uid: uidmapUid,
		PatientUid: patient.Uid,
		RecordUid: recordUid,
	})
	if err != nil {
		return &pb.StaffCreateRecordResponse{
			ErrorCodeAndInfo: &pb.ErrorCodeAndInfo{
				ErrorCode: 1,
				Msg:       "创建失败, 请重试",
			},
		}, err
	}

	// TODO 添加病历记录
	recordRepo := fabricQuery.NewRecordRepo()
	err = recordRepo.Create(context.Background(), &model.Record{
		Uid:              recordUid,
		PatientUid:       patient.Uid,
		WriteStaffUid:    rwStaffUid,
		ReadOnlyStaffUid: rOnlyStaffUid,
		PhotoHash:        in.GetRecordRequest().GetPhotoHash(),
		DocumentHash:     in.GetRecordRequest().GetDocumentHash(),
		Description:      in.GetRecordRequest().GetDescription(),
		LastChangeTime:   jwt.FormatTime(time.Now()),
	})
	if err != nil {
		return &pb.StaffCreateRecordResponse{
			ErrorCodeAndInfo: &pb.ErrorCodeAndInfo{
				ErrorCode: 1,
				Msg:       "创建失败, 请重试",
			},
		}, err
	}

	return &pb.StaffCreateRecordResponse{
		ErrorCodeAndInfo: &pb.ErrorCodeAndInfo{
			ErrorCode: 0,
			Msg:       "创建病历成功",
		},
	}, nil
}
