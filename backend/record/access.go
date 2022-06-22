package record

import (
	"context"

	fabricQuery "github.com/geraldkohn/elian/data/fabric"
)

// rwPermission 根据医生uid和病历记录uid来判断是否有读写权限
func judgeStaffRWPermission(staffUid string, recordUid string) bool {
	recordRepo := fabricQuery.NewRecordRepo()
	record, err := recordRepo.FindByUid(context.Background(), recordUid)
	if err != nil {
		return false
	}

	rwStaffUid := record.WriteStaffUid
	uids, err := unmarshal(rwStaffUid)
	if err != nil || uids == nil {
		return false
	}

	for _, uid := range uids {
		if uid == staffUid {
			return true
		}
	}
	return false
}

func judgeStaffROnlyPermission(staffUid string, recordUid string) bool {
	recordRepo := fabricQuery.NewRecordRepo()
	record, err := recordRepo.FindByUid(context.Background(), recordUid)
	if err != nil {
		return false
	}

	rOnlyStaffUid := record.ReadOnlyStaffUid
	uids, err := unmarshal(rOnlyStaffUid)
	if err != nil || uids == nil {
		return false
	}

	for _, uid := range uids {
		if uid == staffUid {
			return true
		}
	}
	return false
}

// func judgeAgencyRWPermission(agencyUid string, recordUid string) bool {
// 	return true
// }

// func judgeAgencyROnlyPermission(agencyUid string, recordUid string) bool {
// 	return true
// }

func addRWPermission(staffUids []string, recordUid string) (err error) {
	recordRepo := fabricQuery.NewRecordRepo()
	record, err := recordRepo.FindByUid(context.Background(), recordUid)
	if err != nil {
		return err
	}

	oldUidStr := record.WriteStaffUid
	for _, staffUid := range staffUids {
		oldUidStr, err = marshal(oldUidStr, staffUid)
		if err != nil {
			return err
		}
	}
	newUidStr := oldUidStr

	return recordRepo.UpdateWriteStaffUidByUid(context.Background(), recordUid, newUidStr)
}

func addROnlyPermission(staffUids []string, recordUid string) (err error) {
	recordRepo := fabricQuery.NewRecordRepo()
	record, err := recordRepo.FindByUid(context.Background(), recordUid)
	if err != nil {
		return err
	}

	oldUidStr := record.ReadOnlyStaffUid
	for _, staffUid := range staffUids {
		oldUidStr, err = marshal(oldUidStr, staffUid)
		if err != nil {
			return err
		}
	}
	newUidStr := oldUidStr

	return recordRepo.UpdateReadOnlyStaffUidByUid(context.Background(), recordUid, newUidStr)
}

