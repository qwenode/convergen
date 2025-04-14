package as

import "github.com/qwenode/convergen/as/data"

// ================================================================= //
// convergen
//
//go:generate convergen
type Gen interface {
	// :map $2 Age
	// :map $3 User
	// :skip Slice
	// :typecast
	UserA_UserB(data.UserA, int, data.UserA) data.UserB
	// :map $2 Age
	// :map $3 Slice
	// :map $4 User
	// :typecast
	UserA_UserB2(data.UserA, int, []data.UserA, data.UserA) data.UserB
}
