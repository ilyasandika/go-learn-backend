package repositories

import (
	"context"
	"database/sql"
	"errors"
	"uaspw2/helper"
	"uaspw2/models/entity"
)

type UserProfileRepository interface {
	Update(ctx context.Context, tx *sql.Tx, userProfile entity.UserProfile) entity.UserProfile
	Delete(ctx context.Context, tx *sql.Tx, userId int)
	FindByUserID(ctx context.Context, tx *sql.Tx, userId int) (entity.UserProfile, error)
	FindAll(ctx context.Context, tx *sql.Tx) []entity.UserProfile
}

type UserProfileRepositoryImpl struct {
}

func NewUserProfileRepository() UserProfileRepository {
	return &UserProfileRepositoryImpl{}
}

func (repository *UserProfileRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, userProfile entity.UserProfile) entity.UserProfile {
	SQL := `UPDATE user_profiles SET full_name = ?, gender = ?, birthdate = ?, phone_number = ?, address = ?  WHERE user_id = ?`
	_, err := tx.ExecContext(ctx, SQL, userProfile.FullName, userProfile.Gender, userProfile.BirthDate, userProfile.PhoneNumber, userProfile.Address, userProfile.UserId)
	helper.PanicIfErr(err)

	return userProfile
}

func (repository *UserProfileRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, userId int) {
	SQL := `DELETE FROM user_profiles WHERE user_id = ?`
	_, err := tx.ExecContext(ctx, SQL, userId)
	helper.PanicIfErr(err)
}

func (repository *UserProfileRepositoryImpl) FindByUserID(ctx context.Context, tx *sql.Tx, id int) (entity.UserProfile, error) {
	SQL := `SELECT user_id, full_name, gender, birthdate, phone_number, address, created_at, updated_at FROM user_profiles WHERE user_id = ?`
	row, err := tx.QueryContext(ctx, SQL, id)
	helper.PanicIfErr(err)

	var user entity.UserProfile
	if row.Next() {
		var fullName sql.NullString
		var gender sql.NullString
		var birthDate sql.NullString
		var phoneNumber sql.NullString
		var address sql.NullString

		err = row.Scan(&user.UserId, &fullName, &gender, &birthDate, &phoneNumber, &address, &user.CreatedAt, &user.UpdatedAt)
		helper.PanicIfErr(err)

		user.FullName = helper.NullStringToString(fullName)
		user.Gender = helper.NullStringToString(gender)
		user.BirthDate = helper.NullStringToString(birthDate)
		user.PhoneNumber = helper.NullStringToString(phoneNumber)
		user.Address = helper.NullStringToString(address)
	} else {
		return user, errors.New("user not found")
	}

	return user, nil
}

func (repository *UserProfileRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []entity.UserProfile {
	SQL := `SELECT user_id, full_name, gender, birthdate, phone_number, address, created_at, updated_at FROM user_profiles`
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfErr(err)
	defer rows.Close()

	var userProfiles []entity.UserProfile

	for rows.Next() {
		var user entity.UserProfile

		var fullName sql.NullString
		var gender sql.NullString
		var birthDate sql.NullString
		var phoneNumber sql.NullString
		var address sql.NullString

		err = rows.Scan(&user.UserId, &fullName, &gender, &birthDate, &phoneNumber, &address, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			helper.PanicIfErr(err)
		}

		user.FullName = helper.NullStringToString(fullName)
		user.Gender = helper.NullStringToString(gender)
		user.BirthDate = helper.NullStringToString(birthDate)
		user.PhoneNumber = helper.NullStringToString(phoneNumber)
		user.Address = helper.NullStringToString(address)

		userProfiles = append(userProfiles, user)
	}

	return userProfiles
}
