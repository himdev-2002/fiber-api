package models

import (
	"html"
	"strings"
	"tde/fiber-api/core/services"
	"time"

	"github.com/lithammer/shortuuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	// gorm.Model
	// ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()"`
	ID        string    `gorm:"primaryKey;size:30" structs:"id"`
	Username  string    `gorm:"unique;size:20;not null" structs:"username"`
	Password  string    `gorm:"not null" structs:"password,omitempty"`
	Email     string    `gorm:"size:100;unique;not null" structs:"email"`
	FirstName string    `gorm:"size:50;not null" structs:"first_name"`
	LastName  *string   `gorm:"size:50" structs:"last_name"`
	Status    uint      `gorm:"default:1;precision:1;size:1;not null" structs:"status"`
	CreatedAt time.Time `gorm:"not null;autoCreateTime" structs:"created_at"`
	UpdatedAt time.Time `gorm:"not null;autoUpdateTime" structs:"updated_at,omitempty"`
}

func (u *User) SaveUser() (*User, error) {
	tx := services.DBCore.Session(&gorm.Session{SkipDefaultTransaction: true})
	err := tx.Table("m_user").Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) BeforeSave(tx *gorm.DB) error {
	//turn password into hash
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)

	uuid := shortuuid.New()
	u.ID = uuid
	//remove spaces in username
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	u.Email = strings.TrimSpace(u.Email)
	u.FirstName = html.EscapeString(strings.TrimSpace(u.FirstName))
	*u.LastName = html.EscapeString(strings.TrimSpace(*u.LastName))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	u.Status = 1

	return nil
}

func (u *User) FindByID(id *string) error {
	tx := services.DBCore.Session(&gorm.Session{PrepareStmt: true})
	err := tx.Table("m_user").Where("id = ?", *id).Find(&u).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *User) FindByUsername(username *string) error {
	tx := services.DBCore.Session(&gorm.Session{PrepareStmt: true})
	err := tx.Table("m_user").Where("username = ?", *username).Find(&u).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *User) FindByUsernameOrEmail(account *string) error {
	tx := services.DBCore.Session(&gorm.Session{PrepareStmt: true})
	err := tx.Table("m_user").Where("username = ?", *account).Or("email = ?", *account).Find(&u).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *User) ChangePassword(pass string) (*User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		return &User{}, err
	}
	// fmt.Println("====")
	// fmt.Println(pass)
	// fmt.Println(hashedPassword)
	// fmt.Println("====")

	// fmt.Println(u)
	// fmt.Println(string(hashedPassword))
	tx := services.DBCore.Session(&gorm.Session{SkipDefaultTransaction: true})
	now := time.Now()
	result := tx.Table("m_user").Where("id = ?", u.ID).Update("password", string(hashedPassword)).Update("updated_at", now)
	// fmt.Println(result.Statement.SQL.String())
	// fmt.Println(result.RowsAffected)
	// fmt.Println(result.Error)
	if result.RowsAffected <= 0 {
		return &User{}, result.Error
	}

	u.UpdatedAt = now
	u.Password = string(hashedPassword)
	return u, nil
}

func SearchActiveUsers(like_cond *[]string, like_value *map[string]interface{}, users *[]map[string]interface{}) error {
	tx := services.DBCore.Session(&gorm.Session{PrepareStmt: true})
	err := tx.Table("m_user").Where(strings.Join(*like_cond, " AND "), *like_value).Find(users).Error

	if err != nil {
		return err
	}
	return nil
}
