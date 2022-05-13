package postgres

import (
	"edison-takehome/pkg/crypto"
	"edison-takehome/pkg/model"
	"edison-takehome/pkg/repository"
	"errors"
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

type PostgresDB struct {
	db *pg.DB
}

func NewPostgresDB(hostname string, port int, username string, password string) (repository.Repository, error) {
	addr := fmt.Sprintf("%s:%d", hostname, port)
	db := pg.Connect(&pg.Options{
		Addr:     addr,
		User:     username,
		Password: password,
	})

	pgdb := &PostgresDB{
		db: db,
	}

	err := pgdb.createUserSchema()
	if err != nil {
		return nil, err
	}

	return pgdb, nil
}

func (p *PostgresDB) InsertUserAccount(acct model.UserAccount) error {
	hashedPassword := crypto.Hash(acct.Password)
	acct.Password = hashedPassword
	_, err := p.db.Model(&acct).Insert()
	return err
}

func (p *PostgresDB) UpdateUserAccount(userEmail model.UserEmail) (*model.UserAccount, error) {
	userAcct := model.UserAccount{
		Email:     userEmail.Email,
		FirstName: userEmail.FirstName,
		LastName:  userEmail.LastName,
	}

	res, err := p.db.Model(&userAcct).Where("email = ?", userEmail.Email).Column("first_name", "last_name").Update()
	if err != nil {
		return nil, err
	}
	if res.RowsAffected() == 0 {
		return nil, errors.New("record not found")
	}

	return &userAcct, nil
}

func (p *PostgresDB) GetUserAccount(email string) (*model.UserAccount, error) {
	var userAcct model.UserAccount

	err := p.db.Model(&userAcct).Where("email = ?", email).Select()
	if err != nil {
		return nil, err
	}

	return &userAcct, nil
}

func (p *PostgresDB) GetUserAccounts() ([]model.UserAccount, error) {
	var userAccts []model.UserAccount

	err := p.db.Model(&userAccts).Select()
	if err != nil {
		return nil, err
	}

	return userAccts, nil
}

func (p *PostgresDB) Close() error {
	return p.db.Close()
}

func (p *PostgresDB) createUserSchema() error {
	err := p.db.Model(&model.UserAccount{}).DropTable(&orm.DropTableOptions{
		IfExists: true,
	})
	if err != nil {
		return err
	}
	err = p.db.Model(&model.UserAccount{}).CreateTable(&orm.CreateTableOptions{
		IfNotExists: true,
	})
	if err != nil {
		return err
	}
	return nil
}
