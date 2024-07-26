package repository

import (
	"food-reserve/db/repository/interface"
	"gorm.io/gorm"
)

type IUnitOfWork interface {
	Begin() error
	Commit() error
	Rollback() error
	UserRepository() repository.IUserRepository
	RoleRepository() repository.IRoleRepository
	AuthRepository() repository.IAuthRepository
	reset()
}

type UnitOfWork struct {
	db       *gorm.DB
	tx       *gorm.DB
	userRepo repository.IUserRepository
	roleRepo repository.IRoleRepository
	authRepo repository.IAuthRepository
}

func NewUnitOfWork(db *gorm.DB) IUnitOfWork {
	u := &UnitOfWork{
		db: db,
	}
	u.userRepo = NewUserRepository(db)
	u.roleRepo = NewRoleRepository(db)
	u.authRepo = NewAuthRepository(db)
	return u
}

func (u *UnitOfWork) Begin() error {
	tx := u.db.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	u.tx = tx
	u.userRepo = NewUserRepository(u.tx)
	u.roleRepo = NewRoleRepository(u.tx)
	u.authRepo = NewAuthRepository(u.tx)
	return nil
}

func (u *UnitOfWork) Commit() error {
	if u.tx == nil {
		return nil
	}
	err := u.tx.Commit().Error
	u.reset()
	return err
}

func (u *UnitOfWork) Rollback() error {
	if u.tx == nil {
		return nil
	}
	err := u.tx.Rollback().Error
	u.reset()
	return err
}

func (u *UnitOfWork) UserRepository() repository.IUserRepository {
	return u.userRepo
}

func (u *UnitOfWork) RoleRepository() repository.IRoleRepository {
	return u.roleRepo
}

func (u *UnitOfWork) AuthRepository() repository.IAuthRepository {
	return u.authRepo
}

func (u *UnitOfWork) reset() {
	u.tx = nil
	u.userRepo = NewUserRepository(u.db)
	u.roleRepo = NewRoleRepository(u.db)
	u.authRepo = NewAuthRepository(u.db)
}
