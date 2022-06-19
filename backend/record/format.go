package main

import (
	"encoding/json"
)

type uidStruct struct {
	Uids []string `json:"uids"`
}

// @title marshal
// @description
func marshal(originalStr string, newUid string) (newStr string, err error) {
	uids, err := unmarshal(originalStr)
	if err != nil {
		return "", err
	}
	if newUid == "" {
		return "", nil
	}

	uids = append(uids, newUid)
	u := &uidStruct{
		Uids: uids,
	}
	b, err := json.Marshal(u)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func unmarshal(str string) (uids []string, err error) {
	if str == "" {
		return nil, nil
	}
	u := uidStruct{}
	err = json.Unmarshal([]byte(str), &u)
	if err != nil {
		return nil, err
	}
	return u.Uids, nil
}
