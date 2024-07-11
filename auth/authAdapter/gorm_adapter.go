package authadapter

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	authcore "github.com/tanasinp/go-inventory-management/auth/authCore"
	"github.com/tanasinp/go-inventory-management/database"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type gormUserRepository struct {
	db *gorm.DB
}

func NewGormUserRepository(db *gorm.DB) authcore.UserRepository {
	return &gormUserRepository{db: db}
}

func (r *gormUserRepository) CreateUser(user *database.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	if result := r.db.Create(user); result != nil {
		return result.Error
	}
	return nil
}

func (r *gormUserRepository) LoginUser(user *database.User) (string, error) {
	selectedUser := new(database.User)
	if result := r.db.Where("email = ?", user.Email).First(selectedUser); result.Error != nil {
		return "", result.Error
	}
	err := bcrypt.CompareHashAndPassword([]byte(selectedUser.Password), []byte(user.Password))
	if err != nil {
		return "", err
	}
	jwtSecretKey := os.Getenv("JWT_SECRETKEY")
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = selectedUser.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return "", err
	}
	return t, nil
}
