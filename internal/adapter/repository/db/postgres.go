package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/ThailanTec/jogo-do-bicho/internal/core/domain"
	"github.com/ThailanTec/jogo-do-bicho/pkg"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type UserDB struct {
	db *gorm.DB
}

func NewUserPostgresRepository() *UserDB {
	host := "localhost"
	port := "5432"
	user := "postgres"
	password := "animals"
	dbname := "postgres"

	conn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host,
		port,
		user,
		dbname,
		password,
	)

	db, err := gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&domain.User{})

	return &UserDB{
		db: db,
	}
}

func (p *UserDB) CreateUser(ctx context.Context, u domain.User) (user *domain.User, err error) {
	usr := u

	pkg.SHA256Enconder(u.Password)
	checkE, err := p.ValideEmail(u.Email)
	if err != nil {
		return nil, err
	}
	if checkE {
		data := p.db.Create(&usr)

		if data.RowsAffected == 0 {
			return nil, errors.New("erro ao criar usuario")
		}
	}

	return &usr, err
}

func (p *UserDB) ListUser(ctx context.Context, id uint) (user *domain.User, err error) {
	user = &domain.User{}
	req := p.db.First(&user, "id = ? ", id)
	if req.RowsAffected == 0 {
		return nil, errors.New("user not found")
	}

	return user, err
}

func (p *UserDB) ListAllUser() (user *domain.User, err error) {
	var use []*domain.User

	req := p.db.Find(&use)
	if req.Error != nil {
		return nil, errors.New(fmt.Sprintf("user not found: %v", req.Error))
	}
	return user, nil
}

func (p *UserDB) DeleteUser(ctx context.Context, id uint) (deleted bool, err error) {
	user := domain.User{}

	req := p.db.Delete(&user, "id= ?", id)
	if req.RowsAffected == 0 {
		return false, errors.New("could not delete")
	}

	return true, err
}

func (p *UserDB) UpdateUser(ctx context.Context, u domain.User) (user *domain.User, err error) {
	usr := domain.User{}

	req := p.db.Save(usr)
	if req.RowsAffected == 0 {
		return nil, errors.New("user not updated")
	}

	return user, err
}

func (p *UserDB) LoginUser(email, password string) (login bool, err error) {
	usr := domain.User{}
	err = p.db.Where("email = ?", email).First(&usr).Error
	if err != nil {
		return false, errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(usr.Password), []byte(password)); err != nil {
		return false, errors.New("incorret password") // Senha incorreta
	}

	return true, nil
}

func (p *UserDB) ValideEmail(email string) (valide bool, err error) {
	var user domain.User

	if err := p.db.Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, errors.New("email exist on db") // E-mail não encontrado no banco de dados
		}
	}

	return true, nil
}
