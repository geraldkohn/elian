package record

import (
	"context"
	"errors"
	"log"

	query "github.com/geraldkohn/elian/data/biz"
	fabricQuery "github.com/geraldkohn/elian/data/fabric"
	"github.com/geraldkohn/elian/middleware/jwt"
	pb "github.com/geraldkohn/elian/proto/record"
)

func (s *server) StaffQueryRecords(ctx context.Context, in *pb.StaffQueryRecordsRequest) (out *pb.StaffQueryRecordsResponse, err error) {

	// TODO 验证令牌身份
	staffUid, err := jwt.ParseToken(in.GetStaffToken())
	if err != nil {
		return &pb.StaffQueryRecordsResponse{
			RecordResponses: nil,
			ErrorCodeAndInfo: &pb.ErrorCodeAndInfo{
				ErrorCode: 1,
				Msg:       "请重新登录",
			},
		}, errors.New("tokenInvalidError")
	}
	log.Println("医生uid: " + staffUid + " 查找患者记录.")

	var records []*pb.RecordResponse

	// TODO 查找patientuidToRecorduid表, 找出表中的recorduid字段
	uidmapRepo := query.NewUidmapRepo()
	patientRepo := query.NewPatientRepo()
	recordRepo := fabricQuery.NewRecordRepo()

	// TODO 用患者身份证号查找患者记录
	patient, err := patientRepo.SelectByIdCardNumber(context.Background(), in.GetPatientIdCardNumber())
	if err != nil {
		return &pb.StaffQueryRecordsResponse{
			RecordResponses: nil,
			ErrorCodeAndInfo: &pb.ErrorCodeAndInfo{
				ErrorCode: 1,
				Msg:       "无患者记录",
			},
		}, errors.New("noPatientMessageError")
	}
	patientUid := patient.Uid
	patientName := patient.Name

	// TODO 用患者uid来查找患者uid--病历记录uid的映射表
	uidmaps, err := uidmapRepo.SelectAllByPatientUid(context.Background(), patientUid)
	if err != nil {
		return &pb.StaffQueryRecordsResponse{
			RecordResponses: nil,
			ErrorCodeAndInfo: &pb.ErrorCodeAndInfo{
				ErrorCode: 1,
				Msg:       "无患者记录",
			},
		}, errors.New("noPatientMessageError")
	}

	// TODO 遍历这个映射表, 用病历记录uid查找病历记录每一个字段
	for _, uidmap := range uidmaps {
		recordModel, err := recordRepo.FindByUid(context.Background(), uidmap.RecordUid)
		// TODO 判断医生是否有权限读取
		if err != nil {
			log.Println("医生uid: " + staffUid + " 查找患者记录. error: " + err.Error())
			continue
		}
		if !judgeStaffROnlyPermission(staffUid, recordModel.Uid) {
			log.Println("医生uid: " + staffUid + " 查找患者记录, 无权限读取.")
			break
		}
		record := &pb.RecordResponse{
			RecordUid:           recordModel.Uid,
			PatientIdCardNumber: in.GetPatientIdCardNumber(),
			PatientName:         patientName,
			PhotoHash:           recordModel.PhotoHash,
			DocumentHash:        recordModel.DocumentHash,
			LastChangeTime:      recordModel.LastChangeTime,
			Description:         recordModel.Description,
		}
		records = append(records, record)
	}
	return &pb.StaffQueryRecordsResponse{
		RecordResponses: records,
		ErrorCodeAndInfo: &pb.ErrorCodeAndInfo{
			ErrorCode: 0,
			Msg:       "已找到信息",
		},
	}, nil
}
