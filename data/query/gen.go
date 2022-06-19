// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package query

import (
	"context"
	"database/sql"

	"gorm.io/gorm"
)

func Use(db *gorm.DB) *Query {
	return &Query{
		db:                    db,
		Agency:                newAgency(db),
		Patient:               newPatient(db),
		PatientuidToRecorduid: newPatientuidToRecorduid(db),
		Record:                newRecord(db),
		Staff:                 newStaff(db),
	}
}

type Query struct {
	db *gorm.DB

	Agency                agency
	Patient               patient
	PatientuidToRecorduid patientuidToRecorduid
	Record                record
	Staff                 staff
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:                    db,
		Agency:                q.Agency.clone(db),
		Patient:               q.Patient.clone(db),
		PatientuidToRecorduid: q.PatientuidToRecorduid.clone(db),
		Record:                q.Record.clone(db),
		Staff:                 q.Staff.clone(db),
	}
}

type queryCtx struct {
	Agency                agencyDo
	Patient               patientDo
	PatientuidToRecorduid patientuidToRecorduidDo
	Record                recordDo
	Staff                 staffDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		Agency:                *q.Agency.WithContext(ctx),
		Patient:               *q.Patient.WithContext(ctx),
		PatientuidToRecorduid: *q.PatientuidToRecorduid.WithContext(ctx),
		Record:                *q.Record.WithContext(ctx),
		Staff:                 *q.Staff.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	return &QueryTx{q.clone(q.db.Begin(opts...))}
}

type QueryTx struct{ *Query }

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}