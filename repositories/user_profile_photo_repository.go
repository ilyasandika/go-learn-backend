package repositories

import (
	"context"
	"database/sql"
	"errors"
	"uaspw2/helper"
	"uaspw2/models/entity"
)

type UserProfilePhotoRepository interface {
	Update(ctx context.Context, tx *sql.Tx, profiles entity.UserProfilePhoto) entity.UserProfilePhoto
	FindByUserID(ctx context.Context, tx *sql.Tx, userId int) (entity.UserProfilePhoto, error)
}

type UserProfilePhotoRepositoryImpl struct {
}

func NewUserProfilePhotoRepository() UserProfilePhotoRepository {
	return &UserProfilePhotoRepositoryImpl{}
}

func (repository *UserProfilePhotoRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, profiles entity.UserProfilePhoto) entity.UserProfilePhoto {
	SQL := `UPDATE user_profile_photos SET path = ? WHERE user_id = ?`
	_, err := tx.ExecContext(ctx, SQL, profiles.Path, profiles.UserId)
	helper.PanicIfErr(err)

	return profiles
}

func (repository *UserProfilePhotoRepositoryImpl) FindByUserID(ctx context.Context, tx *sql.Tx, userId int) (entity.UserProfilePhoto, error) {
	SQL := `SELECT user_id, path, created_at, updated_at FROM user_profile_photos WHERE user_id = ?`
	row, err := tx.QueryContext(ctx, SQL, userId)
	helper.PanicIfErr(err)
	defer row.Close()

	var profiles entity.UserProfilePhoto
	if row.Next() {
		err := row.Scan(&profiles.UserId, &profiles.Path, &profiles.CreatedAt, &profiles.UpdatedAt)
		helper.PanicIfErr(err)
		return profiles, nil
	} else {
		return profiles, errors.New("user profile photo not found")
	}
}
