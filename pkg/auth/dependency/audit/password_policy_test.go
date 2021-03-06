package audit

import (
	"encoding/json"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	. "github.com/skygeario/skygear-server/pkg/core/skytest"
)

func TestPasswordPolicyJSON(t *testing.T) {
	Convey("PasswordPolicy JSON serialization", t, func() {
		v := PasswordPolicy{
			Name: PasswordTooShort,
			Info: map[string]interface{}{
				"min_length": 8,
				"pw_length":  6,
			},
		}
		b, err := json.Marshal(v)
		So(err, ShouldBeNil)
		So(b, ShouldEqualJSON, `{"kind":"PasswordTooShort","min_length":8,"pw_length":6}`)
	})
}
