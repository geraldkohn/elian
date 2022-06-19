package uid_test

import (
	"testing"

	"github.com/geraldkohn/elian/middleware/uid"
)

func TestNewUid(t *testing.T) {
	uid := uid.NewUid()
	t.Logf("uid: %s", uid)
}
