package hearths

import (
	"fmt"
)

func newErr(err string) error {
	return fmt.Errorf(err)
}
