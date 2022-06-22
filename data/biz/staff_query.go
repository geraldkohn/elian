package biz

import (
	"context"

	"github.com/geraldkohn/elian/config"
	"github.com/geraldkohn/elian/data/model"
	"github.com/geraldkohn/elian/data/query"
	"gorm.io/gorm"
)

type staffRepo interface {
	Create(ctx context.Context, staff *model.Staff) (err error)

	SelectByUid(ctx context.Context, uid string) (staff *model.Staff, err error)
	SelectByIdCardNumber(ctx context.Context, idCardNumber string) (staff *model.Staff, err error)
	SelectByHospital(ctx context.Context, hospital string) (staffs []*model.Staff, err error)
	SelectByHospitalAndDepartment(ctx context.Context, hospital string, department string) (staffs []*model.Staff, err error)
	SelectByHospitalAndJobNumber(ctx context.Context, hospital string, jobNumber string) (staff *model.Staff, err error)
	SelectByHospitalAndDepartmentAndJobNumber(ctx context.Context, hospital string, department string, jobNumber string) (staff *model.Staff, err error)
	SelectByHospitalAndDepartmentAndName(ctx context.Context, hospital string, department string, name string) (staff *model.Staff, err error)

	UpdatePassword(ctx context.Context, uid string, password string) (err error)
	UpdateTokenAndLoginTime(ctx context.Context, uid string, newToken string, newLoginTime string) (err error)
}

type staffManager struct {
	DB *gorm.DB
}

func NewStaffRepo() staffRepo {
	db, _ := config.GenerateDB(config.Dsn)
	//建表
	if !db.Migrator().HasTable("staff") {
		err := db.Migrator().CreateTable(&model.Staff{})
		if err != nil {
			panic("创建staff表错误")
		}
	}
	return &staffManager{DB: db}
}

func (r *staffManager) Create(ctx context.Context, staff *model.Staff) (err error) {
	err = query.Use(r.DB).Staff.WithContext(ctx).Create(staff)
	return
}

func (r *staffManager) SelectByUid(ctx context.Context, uid string) (staff *model.Staff, err error) {
	s := query.Use(r.DB).Staff
	staff, err = s.WithContext(ctx).Where(s.Uid.Eq(uid)).First()
	return
}

func (r *staffManager) SelectByIdCardNumber(ctx context.Context, idCardNumber string) (staff *model.Staff, err error) {
	s := query.Use(r.DB).Staff
	staff, err = s.WithContext(ctx).Where(s.IdCardNumber.Eq(idCardNumber)).First()
	return
}

func (r *staffManager) SelectByHospitalAndJobNumber(ctx context.Context, hospital string, jobNumber string) (staff *model.Staff, err error) {
	s := query.Use(r.DB).Staff
	staff, err = s.WithContext(ctx).Where(s.Hospital.Eq(hospital), s.JobNumber.Eq(jobNumber)).First()
	return
}

func (r *staffManager) SelectByHospital(ctx context.Context, hospital string) (staffs []*model.Staff, err error) {
	s := query.Use(r.DB).Staff
	staffs, err = s.WithContext(ctx).Where(s.Hospital.Eq(hospital)).Find()
	return
}

func (r *staffManager) SelectByHospitalAndDepartment(ctx context.Context, hospital string, department string) (staffs []*model.Staff, err error) {
	s := query.Use(r.DB).Staff
	staffs, err = s.WithContext(ctx).Where(s.Hospital.Eq(hospital), s.Department.Eq(department)).Find()
	return
}

func (r *staffManager) SelectByHospitalAndDepartmentAndJobNumber(ctx context.Context, hospital string, department string, jobNumber string) (staff *model.Staff, err error) {
	s := query.Use(r.DB).Staff
	staff, err = s.WithContext(ctx).Where(s.Hospital.Eq(hospital), s.Department.Eq(department), s.JobNumber.Eq(jobNumber)).First()
	return
}

func (r *staffManager) SelectByHospitalAndDepartmentAndName(ctx context.Context, hospital string, department string, name string) (staff *model.Staff, err error) {
	s := query.Use(r.DB).Staff
	staff, err = s.WithContext(ctx).Where(s.Hospital.Eq(hospital), s.Department.Eq(department), s.Name.Eq(name)).First()
	return
}

func (r *staffManager) UpdatePassword(ctx context.Context, uid string, password string) (err error) {
	s := query.Use(r.DB).Staff
	_, err = s.WithContext(ctx).Where(s.Uid.Eq(uid)).Updates(map[string]interface{}{"password": password})
	return
}

func (r *staffManager) UpdateTokenAndLoginTime(ctx context.Context, uid string, newToken string, newLoginTime string) (err error) {
	s := query.Use(r.DB).Staff
	_, err = s.WithContext(ctx).Where(s.Uid.Eq(uid)).Updates(map[string]interface{}{"token": newToken, "last_login_time": newLoginTime})
	return
}
