package biz

import (
	"context"

	"github.com/geraldkohn/elian/config"
	"github.com/geraldkohn/elian/data/model"
	"github.com/geraldkohn/elian/data/query"
	"gorm.io/gorm"
)

type uidmapRepo interface {
	Create(ctx context.Context, uidmap *model.PatientuidToRecorduid) (err error)
	SelectAllByPatientUid(ctx context.Context, patientUid string) (patientuidToRecorduid []*model.PatientuidToRecorduid, err error)
}

type uidmapManager struct {
	DB *gorm.DB
}

func NewUidmapRepo() uidmapRepo {
	return &uidmapManager{DB: config.DB}
}

func (r *uidmapManager) Create(ctx context.Context, uidmap *model.PatientuidToRecorduid) (err error) {
	err = query.Use(r.DB).PatientuidToRecorduid.WithContext(ctx).Create(uidmap)
	return 
}

func (r *uidmapManager) SelectAllByPatientUid(ctx context.Context, patientUid string) (patientuidToRecorduid []*model.PatientuidToRecorduid, err error) {
	u := query.Use(r.DB).PatientuidToRecorduid
	return u.WithContext(ctx).Where(u.PatientUid.Eq(patientUid)).Find()
}