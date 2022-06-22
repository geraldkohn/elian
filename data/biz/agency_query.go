package biz

import (
	"context"

	"github.com/geraldkohn/elian/config"
	"github.com/geraldkohn/elian/data/model"
	"github.com/geraldkohn/elian/data/query"
	"gorm.io/gorm"
)

type agencyRepo interface {
	Create(ctx context.Context, agency *model.Agency) (err error)
	SelectByLicense(ctx context.Context, license string) (agency *model.Agency, err error)
	UpdateTokenAndLoginTime(ctx context.Context, uid string, newToken string, newLoginTime string) (err error)
}

type agencyManager struct {
	DB *gorm.DB
}

func NewAgencyRepo() agencyRepo {
	db, _ := config.GenerateDB(config.Dsn)
	if !db.Migrator().HasTable("agency") {
		err := db.Migrator().CreateTable(&model.Agency{})
		if err != nil {
			panic("创建agency表错误")
		}
	}
	return &agencyManager{DB: db}
}

func (r *agencyManager) Create(ctx context.Context, agency *model.Agency) (err error) {
	err = query.Use(r.DB).Agency.WithContext(ctx).Create(agency)
	return
}

func (r *agencyManager) UpdateTokenAndLoginTime(ctx context.Context, uid string, newToken string, newLoginTime string) (err error) {
	a := query.Use(r.DB).Agency
	_, err = a.WithContext(ctx).Where(a.Uid.Eq(uid)).Updates(map[string]interface{}{"token": newToken, "last_login_time": newLoginTime})
	return
}

func (r *agencyManager) SelectByLicense(ctx context.Context, license string) (agency *model.Agency, err error) {
	a := query.Use(r.DB).Agency
	return a.WithContext(ctx).Where(a.License.Eq(license)).First()
}
