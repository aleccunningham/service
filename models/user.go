package models

const (
  host      = "localhost"
  port      = 5432
  user      = "postgresql"
  password  = "postgresql"
  dbname    = "service-db"
)

type User struct {
  gorm.Model
  Name          string
  Email         string `gorm:"not null;unique_index"`
  Color         string
  Orders        []Order
}

type UserDB interface {
  // Methods for querying single users
  ByID(id uint) (*User, error)
  ByEmail(email string) (*User, error)
  ByRemember(token string) (*User, error)

  // Methods for altering users
  Create(user *User) error
  Update(user *User) error
  Delete(id uint) error
}

// Wrapper for the User model to allow
// for manipulation, etc
type UserService interface {
  Authenticate(email, password string) (*User, error)
  InitiateReset(email string) (string, error)
  CompleteReset(token, newPw string) (*User, error)
  UserDB
}

func NewUserService(db *gorm.DB, pepper, hmacKey string) UserService {
  ug := &userGorm{db}
  hmac := hash.NewHMAC(hmacKey)
  uv := newUserValidator(ug, hmac, pepper)
  return &UserService{
    UserDB:     uv,
    pepper:     pepper,
    pwResetDB:  newPwResetValidator(&pwResetGorm{db}, hmac),
  }
}

var _ UserService = &UserService{}
