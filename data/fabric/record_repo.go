package fabric

import (
	"context"

	"github.com/geraldkohn/elian/data/model"
)


// record数据操作接口
// uid是指record的uid(即record的唯一识别码)
type recordRepo interface {
	Create(ctx context.Context, record *model.Record) (err error)
	UpdateWriteStaffUidByUid(ctx context.Context, uid string, writeStaffUid string) (err error)
	UpdateReadOnlyStaffUidByUid(ctx context.Context, uid string, readOnlyStaffUid string) (err error)
	UpdatePhotoHashByUid(ctx context.Context, uid string, photoHash string) (err error)
	UpdateDocumentHashByUid(ctx context.Context, uid string, documentHash string) (err error)
	UpdateLastChangeTimeByUid(ctx context.Context, uid string, lastChangeTime string) (err error)
	UpdateDescriptionByUid(ctx context.Context, uid string, description string) (err error)
	FindByUid(ctx context.Context, uid string) (record *model.Record, err error)
	FindWriteStaffUidByUid(ctx context.Context, uid string) (str string, err error)
	FindReadOnlyStaffUidByUid(ctx context.Context, uid string) (str string, err error)
	FindPhotoHashByUid(ctx context.Context, uid string) (photoHash string, err error)
	FindDocumentHashByUid(ctx context.Context, uid string) (documentHash string, err error)
	FindLastChangeTimeByUid(ctx context.Context, uid string) (lastChangeTime string, err error)
	FindDescriptionByUid(ctx context.Context, uid string) (description string, err error)
}

type recordRepoManager struct {
	wallet recordWallet
}

func NewRecordRepo() recordRepo {
	return &recordRepoManager{
		wallet: newRecordWallet(),
	}
}

func (r *recordRepoManager) Create(ctx context.Context, record *model.Record) (err error) {
	k := record.Uid
	v := marshalRecord(record)
	r.wallet.SetRecord(k, v)
	return
}

func (r *recordRepoManager) UpdateWriteStaffUidByUid(ctx context.Context, uid string, writeStaffUid string) (err error) {
	k := uid
	v, err := r.wallet.GetRecord(k)
	if err != nil {
		return err
	}
	record := unmarshalRecord(v)
	newRecord := &model.Record{
		Uid:              uid,
		PatientUid:       record.PatientUid,
		WriteStaffUid:    writeStaffUid,
		ReadOnlyStaffUid: record.ReadOnlyStaffUid,
		PhotoHash:        record.PhotoHash,
		DocumentHash:     record.DocumentHash,
		LastChangeTime:   record.LastChangeTime,
		Description:      record.Description,
	}
	v = marshalRecord(newRecord)
	err = r.wallet.UpdateRecord(k, v)
	return err
}

func (r *recordRepoManager) UpdateReadOnlyStaffUidByUid(ctx context.Context, uid string, readOnlyStaffUid string) (err error) {
	k := uid
	v, err := r.wallet.GetRecord(k)
	if err != nil {
		return err
	}
	record := unmarshalRecord(v)
	newRecord := &model.Record{
		Uid:              uid,
		PatientUid:       record.PatientUid,
		WriteStaffUid:    record.WriteStaffUid,
		ReadOnlyStaffUid: readOnlyStaffUid,
		PhotoHash:        record.PhotoHash,
		DocumentHash:     record.DocumentHash,
		LastChangeTime:   record.LastChangeTime,
		Description:      record.Description,
	}
	v = marshalRecord(newRecord)
	err = r.wallet.UpdateRecord(k, v)
	return err
}

func (r *recordRepoManager) UpdatePhotoHashByUid(ctx context.Context, uid string, photoHash string) (err error) {
	k := uid
	v, err := r.wallet.GetRecord(k)
	if err != nil {
		return err
	}
	record := unmarshalRecord(v)
	newRecord := &model.Record{
		Uid:              uid,
		PatientUid:       record.PatientUid,
		WriteStaffUid:    record.WriteStaffUid,
		ReadOnlyStaffUid: record.ReadOnlyStaffUid,
		PhotoHash:        photoHash,
		DocumentHash:     record.DocumentHash,
		LastChangeTime:   record.LastChangeTime,
		Description:      record.Description,
	}
	v = marshalRecord(newRecord)
	err = r.wallet.UpdateRecord(k, v)
	return err
}

func (r *recordRepoManager) UpdateDocumentHashByUid(ctx context.Context, uid string, documentHash string) (err error) {
	k := uid
	v, err := r.wallet.GetRecord(k)
	if err != nil {
		return err
	}
	record := unmarshalRecord(v)
	newRecord := &model.Record{
		Uid:              uid,
		PatientUid:       record.PatientUid,
		WriteStaffUid:    record.WriteStaffUid,
		ReadOnlyStaffUid: record.ReadOnlyStaffUid,
		PhotoHash:        record.PhotoHash,
		DocumentHash:     documentHash,
		LastChangeTime:   record.LastChangeTime,
		Description:      record.Description,
	}
	v = marshalRecord(newRecord)
	err = r.wallet.UpdateRecord(k, v)
	return err
}

func (r *recordRepoManager) UpdateLastChangeTimeByUid(ctx context.Context, uid string, lastChangeTime string) (err error) {
	k := uid
	v, err := r.wallet.GetRecord(k)
	if err != nil {
		return err
	}
	record := unmarshalRecord(v)
	newRecord := &model.Record{
		Uid:              uid,
		PatientUid:       record.PatientUid,
		WriteStaffUid:    record.WriteStaffUid,
		ReadOnlyStaffUid: record.ReadOnlyStaffUid,
		PhotoHash:        record.PhotoHash,
		DocumentHash:     record.DocumentHash,
		LastChangeTime:   lastChangeTime,
		Description:      record.Description,
	}
	v = marshalRecord(newRecord)
	err = r.wallet.UpdateRecord(k, v)
	return err
}

func (r *recordRepoManager) UpdateDescriptionByUid(ctx context.Context, uid string, description string) (err error) {
	k := uid
	v, err := r.wallet.GetRecord(k)
	if err != nil {
		return err
	}
	record := unmarshalRecord(v)
	newRecord := &model.Record{
		Uid:              uid,
		PatientUid:       record.PatientUid,
		WriteStaffUid:    record.WriteStaffUid,
		ReadOnlyStaffUid: record.ReadOnlyStaffUid,
		PhotoHash:        record.PhotoHash,
		DocumentHash:     record.DocumentHash,
		LastChangeTime:   record.LastChangeTime,
		Description:      description,
	}
	v = marshalRecord(newRecord)
	err = r.wallet.UpdateRecord(k, v)
	return err
}

func (r *recordRepoManager) FindByUid(ctx context.Context, uid string) (record *model.Record, err error) {
	k := uid
	v, err := r.wallet.GetRecord(k)
	if err != nil {
		return nil, err
	}
	record = unmarshalRecord(v)
	return record, nil
}

func (r *recordRepoManager) FindWriteStaffUidByUid(ctx context.Context, uid string) (str string, err error) {
	k := uid
	v, err := r.wallet.GetRecord(k)
	if err != nil {
		return "", err
	}
	record := unmarshalRecord(v)
	return record.WriteStaffUid, nil
}

func (r *recordRepoManager) FindReadOnlyStaffUidByUid(ctx context.Context, uid string) (str string, err error) {
	k := uid
	v, err := r.wallet.GetRecord(k)
	if err != nil {
		return "", err
	}
	record := unmarshalRecord(v)
	return record.ReadOnlyStaffUid, nil
}

func (r *recordRepoManager) FindPhotoHashByUid(ctx context.Context, uid string) (photoHash string, err error) {
	k := uid
	v, err := r.wallet.GetRecord(k)
	if err != nil {
		return "", err
	}
	record := unmarshalRecord(v)
	return record.PhotoHash, nil
}

func (r *recordRepoManager) FindDocumentHashByUid(ctx context.Context, uid string) (documentHash string, err error) {
	k := uid
	v, err := r.wallet.GetRecord(k)
	if err != nil {
		return "", err
	}
	record := unmarshalRecord(v)
	return record.DocumentHash, nil
}

func (r *recordRepoManager) FindLastChangeTimeByUid(ctx context.Context, uid string) (lastChangeTime string, err error) {
	k := uid
	v, err := r.wallet.GetRecord(k)
	if err != nil {
		return "", err
	}
	record := unmarshalRecord(v)
	return record.LastChangeTime, nil
}

func (r *recordRepoManager) FindDescriptionByUid(ctx context.Context, uid string) (description string, err error) {
	k := uid
	v, err := r.wallet.GetRecord(k)
	if err != nil {
		return "", err
	}
	record := unmarshalRecord(v)
	return record.Description, nil
}
