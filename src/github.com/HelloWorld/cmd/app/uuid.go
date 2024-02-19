package main

import (
	"github.com/ryanfowler/uuid"
)

func GetUuid() string {
	u := uuid.Must(uuid.NewV4())
	return u.String()
}
