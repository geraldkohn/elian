package fabric

import (
	"errors"
)

/**
钱包需要的功能接口
*/

type recordWallet interface {
	GetRecord(k string) (v string, err error)
	SetRecord(k string, v string) (err error)
	DeleteRecord(k string) (err error)
	UpdateRecord(k string, v string) (err error)
}

/**
钱包功能接口的简单实现, 目前没有用fabric链码. 在下一版会加入链码.
*/

type recordWalletManager struct {
	tempdata tempDataRepo
}

func newRecordWallet() recordWallet {
	return &recordWalletManager{
		tempdata: newTempData(),
	}
}

// GetRecord 查询记录, 如果没有记录需要返回noRecordError
func (r *recordWalletManager) GetRecord(k string) (v string, err error) {
	r.tempdata.mu.RLock()
	defer r.tempdata.mu.RUnlock()
	if v, ok := r.tempdata.data[k]; ok {
		return v, nil
	}
	return "", errors.New("noRecordError")
}

func (r *recordWalletManager) SetRecord(k string, v string) (err error) {
	r.tempdata.mu.Lock()
	defer r.tempdata.mu.Unlock()
	r.tempdata.data[k] = v
	return nil
}

// DeleteRecord删除记录, 如果没有记录返回noRecordError
func (r *recordWalletManager) DeleteRecord(k string) (err error) {
	r.tempdata.mu.Lock()
	defer r.tempdata.mu.Unlock()
	if _, ok := r.tempdata.data[k]; ok {
		delete(r.tempdata.data, k)
		return nil
	}
	return errors.New("noRecordError")
}

// UpdateRecord更新记录, 如果没有记录返回noRecordError
func (r *recordWalletManager) UpdateRecord(k string, v string) (err error) {
	r.tempdata.mu.Lock()
	defer r.tempdata.mu.Unlock()
	if _, ok := r.tempdata.data[k]; ok {
		r.tempdata.data[k] = v
		return nil
	}
	return errors.New("noRecordError")
}
