package jwt_test

import (
	"testing"

	"github.com/geraldkohn/elian/middleware/jwt"
	"github.com/geraldkohn/elian/middleware/uid"
)

func TestSignedAndParseToken(t *testing.T) {
	uid := uid.NewUid()
	tokenString, _ := jwt.SignedToken(uid)
	t.Log("tokenString: ", tokenString)
	uidParse, _ := jwt.ParseToken(tokenString)
	t.Log("uidParse: ", uidParse)
	if uid != uidParse {
		t.Fatalf("失败")
	}
}

func TestSignedToken(t *testing.T) {
	uid := uid.NewUid()
	tokenString, _ := jwt.SignedToken(uid)
	t.Log(uid)
	t.Log(tokenString)
}
