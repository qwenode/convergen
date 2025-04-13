package domain

import (
    "github.com/qwenode/convergen/tests/fixtures/usecase/typecast/enums"
)

type User struct {
    ID     int
    Name   string
    Status enums.Status
}
