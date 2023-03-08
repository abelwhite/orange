// filename: internal/data/schools.go
package data

import (
	"time"

	"honnef.co/go/tools/lintcmd/version"
)

type School struct{
	ID int64
	Name string
	Level string
	Contact string
	Phone string
	Email string
	Website string
	Address string
	Mode []string
	CreatedAt time.Time
	version int32
}