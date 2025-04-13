package as

import "github.com/qwenode/convergen/as/data"

// ================================================================= //
// convergen
//
//go:generate convergen
type Gen interface {
    // :map $2 Age
    // :map $3 User
    // :typecast
    UserA_UserB(data.UserA, int, data.UserA) data.UserB
}
