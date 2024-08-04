package repositories

import (
	"context"
	"database/sql"
	"errors"
	"uaspw2/helper"
	"uaspw2/models/entity"
)

type AuthRepository interface {
	GetUserByUsername(ctx context.Context, tx *sql.Tx, username string) (entity.UserWithProfile, error)
	RegisterUser(ctx context.Context, tx *sql.Tx, request entity.User) entity.User
	CreateUserProfileOnRegisterUser(ctx context.Context, tx *sql.Tx, userId int, fullName string)
	CreateUserPhotoProfileOnRegisterUser(ctx context.Context, tx *sql.Tx, profile entity.UserProfilePhoto)
}

type AuthRepositoryImpl struct {
}

func NewAuthenticationRepository() AuthRepository {
	return &AuthRepositoryImpl{}
}

func (repository *AuthRepositoryImpl) GetUserByUsername(ctx context.Context, tx *sql.Tx, username string) (entity.UserWithProfile, error) {
	SQL := `SELECT 
				u.id,
				u.username,
				u.password,
				u.role,
				u.created_at AS user_created_at,
				u.updated_at AS user_updated_at,
				p.user_id,
				p.full_name,
				p.gender,
				p.birthdate,
				p.phone_number,
				p.address,
				p.created_at AS profile_created_at,
				p.updated_at AS profile_updated_at
			FROM 
				users AS u
			JOIN 
				user_profiles AS p
			ON 
				u.id = p.user_id
			WHERE 
				u.username = ?;`
	row, err := tx.QueryContext(ctx, SQL, username)
	helper.PanicIfErr(err)
	defer row.Close()

	var user entity.UserWithProfile
	if row.Next() {
		var fullName sql.NullString
		var gender sql.NullString
		var birthDate sql.NullString
		var phoneNumber sql.NullString
		var address sql.NullString

		err := row.Scan(&user.Id, &user.Username, &user.Password, &user.Role, &user.CreatedAt, &user.UpdatedAt,
			&user.Profile.UserId, &fullName, &gender, &birthDate, &phoneNumber, &address, &user.Profile.CreatedAt, &user.Profile.UpdatedAt)
		helper.PanicIfErr(err)

		user.Profile.FullName = helper.NullStringToString(fullName)
		user.Profile.Gender = helper.NullStringToString(gender)
		user.Profile.BirthDate = helper.NullStringToString(birthDate)
		user.Profile.PhoneNumber = helper.NullStringToString(phoneNumber)
		user.Profile.Address = helper.NullStringToString(address)

		return user, nil
	} else {
		return user, errors.New("user does not exist")
	}
}

func (repository *AuthRepositoryImpl) RegisterUser(ctx context.Context, tx *sql.Tx, user entity.User) entity.User {
	SQL := `INSERT INTO users (username, password, role) VALUES (?, ?, ?)`

	result, err := tx.ExecContext(ctx, SQL, user.Username, user.Password, user.Role)
	helper.PanicIfErr(err)

	lastInsertId, err := result.LastInsertId()
	helper.PanicIfErr(err)

	user.Id = int(lastInsertId)

	return user
}
func (repository *AuthRepositoryImpl) CreateUserProfileOnRegisterUser(ctx context.Context, tx *sql.Tx, userId int, fullName string) {
	SQL := `INSERT INTO user_profiles (user_id, full_name) VALUES (?, ?)`
	_, err := tx.ExecContext(ctx, SQL, userId, fullName)
	helper.PanicIfErr(err)
}

func (repository *AuthRepositoryImpl) CreateUserPhotoProfileOnRegisterUser(ctx context.Context, tx *sql.Tx, profile entity.UserProfilePhoto) {
	SQL := `INSERT INTO user_profile_photos (user_id, path) VALUES (?,?)`
	_, err := tx.ExecContext(ctx, SQL, profile.UserId, profile.Path)
	helper.PanicIfErr(err)
}
