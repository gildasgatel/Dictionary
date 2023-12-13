package rows

import (
	"fmt"
)

type Rows struct {
	Key  []byte
	Desc []byte
	Date []byte
}

func (r Rows) String() string {
	return fmt.Sprintf("%s\t\t%s\t%s", r.Key, r.Desc, r.Date)
}
