// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"goframe-websocket/internal/dao/internal"
)

// internalUserDao is internal type for wrapping internal DAO implements.
type internalUserDao = *internal.UserDao

// userDao is the data access object for table user.
// You can define custom methods on it to extend its functionality as you wish.
type userDao struct {
	internalUserDao
}

var (
	// User is globally public accessible object for table user operations.
	User = userDao{
		internal.NewUserDao(),
	}
)

// Fill with you ideas below.
