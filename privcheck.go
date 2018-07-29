// +build !windows

package privcheck

import (
	"os"
)

// IsAdmin checks to determine if the caller has effective UID of 0 (root)
func IsAdmin() bool {
	return os.Geteuid() == 0
}
