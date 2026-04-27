package perms

import (
	"encoding/json"
	"errors"
	"fmt"
)

var ErrUnknownPermission = errors.New("unknown permission")

type Permissions []string

var permissionsMap = map[string]Permissions{
	"op":      Permissions{"op", "present", "message", "caption", "token"},
	"present": Permissions{"present", "message"},
	"message": Permissions{"message"},
	"observe": Permissions{},
	"caption": Permissions{"caption"},
	"admin":   Permissions{"admin"},
}

func ExpandPermissions(name string) (Permissions, error) {
	perms, ok := permissionsMap[name]
	if !ok {
		return Permissions{}, ErrUnknownPermission
	}
	return perms, nil
}

func (p Permissions) String() string {
	v, err := json.Marshal([]string(p))
	if err != nil {
		return fmt.Sprintf("(ERROR=%v)", err)
	}
	return string(v)
}

func (p *Permissions) UnmarshalJSON(b []byte) error {
	var a []string
	if err := json.Unmarshal(b, &a); err == nil {
		*p = Permissions(a)
		return nil
	}
	var s string
	if err := json.Unmarshal(b, &s); err == nil {
		a, err := ExpandPermissions(s)
		if err != nil {
			return err
		}
		*p = a
		return nil
	}
	return errors.New("permissions must be a string or array of strings")
}

func (p Permissions) MarshalJSON() ([]byte, error) {
	return json.Marshal([]string(p))
}
