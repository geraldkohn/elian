package biz

import (
	"context"

	"github.com/geraldkohn/elian/config"
	"github.com/geraldkohn/elian/data/model"
	query "github.com/geraldkohn/elian/data/query"
	"gorm.io/gorm"
)

type patientRepo interface {
	Create(ctx context.Context, pateint *model.Patient) (err error)
	SelectByUid(ctx context.Context, uid string) (patient *model.Patient, err error)
	SelectByIdCardNumber(ctx context.Context, idCardNumber string) (patient *model.Patient, err error)
	UpdatePassword(ctx context.Context, uid string, password string) (err error)
	UpdateTokenAndLoginTime(ctx context.Context, uid string, newToken string, newLoginTime string) (err error)
}

type patientManager struct {
	DB *gorm.DB
}

func NewPatientRepo() patientRepo {
	return &patientManager{DB: config.DB}
}

func (r *patientManager) Create(ctx context.Context, pateint *model.Patient) (err error) {
	err = query.Use(r.DB).Patient.WithContext(ctx).Create(pateint)
	return
}

func (r *patientManager) SelectByUid(ctx context.Context, uid string) (patient *model.Patient, err error) {
	p := query.Use(r.DB).Patient
	patient, err = p.WithContext(ctx).Where(p.Uid.Eq(uid)).First()
	return
}

func (r *patientManager) SelectByIdCardNumber(ctx context.Context, idCardNumber string) (patient *model.Patient, err error) {
	p := query.Use(r.DB).Patient
	patient, err = p.WithContext(ctx).Where(p.IdCardNumber.Eq(idCardNumber)).First()
	return
}

func (r *patientManager) UpdatePassword(ctx context.Context, uid string, password string) (err error) {
	p := query.Use(r.DB).Patient
	_, err = p.WithContext(ctx).Where(p.Uid.Eq(uid)).Updates(map[string]interface{}{"password": password})
	return err
}

func (r *patientManager) UpdateTokenAndLoginTime(ctx context.Context, uid string, newToken string, newLoginTime string) (err error) {
	p := query.Use(r.DB).Patient
	_, err = p.WithContext(ctx).Where(p.Uid.Eq(uid)).Updates(map[string]interface{}{"token": newToken, "last_login_time": newLoginTime})
	return err
}
