package dao

import (
	"context"
	"fmt"
	"os"
	"testing"

	"notion/shared/mongo/objid"
	mongotesting "notion/shared/mongo/testing"

	accountpb "notion/account/api/gen/v1"
)

func TestAccountLifecycle(t *testing.T) {
	c := context.Background()
	mc, err := mongotesting.NewClient(c)
	if err != nil {
		t.Fatal(err)
	}
	db := mc.Database("notion")
	err = mongotesting.SetupIndexes(c, db)
	if err != nil {
		t.Fatalf("cannot setup indexes: %v", err)
	}
	m := NewMongo(db)
	account := AccountRecord{
		Account: &accountpb.Account{
			Email:    "test@gmail.com",
			Username: "test",
			Password: "123456",
		},
	}
	cases := []struct {
		name    string
		op      func() error
		wantErr bool
	}{
		{
			name: "add_account_should_success",
			op: func() error {
				a, err := m.CreateAccount(context.Background(), &account)
				if err != nil {
					return err
				}
				if a.Account.Email != account.Account.Email {
					return fmt.Errorf("create account email [%v] but got [%v]", account.Account.Email, a.Account.Email)
				}
				if a.Account.Username != account.Account.Username {
					return fmt.Errorf("create account username [%v] but got [%v]", account.Account.Username, a.Account.Username)
				}
				if a.Account.Password != account.Account.Password {
					return fmt.Errorf("create account passowrd [%v] but got [%v]", account.Account.Password, a.Account.Password)
				}
				account = *a
				return nil
			},
			wantErr: false,
		},
		{
			name: "add_account_again_should_fail",
			op: func() error {
				again := account
				_, err := m.CreateAccount(context.Background(), &again)
				if err != nil {
					return err
				}
				return nil
			},
			wantErr: true,
		},
		{
			name: "get_account_should_success",
			op: func() error {
				a, err := m.GetAccount(context.Background(), objid.ToAccountID(account.ID))
				if err != nil {
					return err
				}
				if a.Account.Email != account.Account.Email {
					return fmt.Errorf("account email [%v] but got [%v]", account.Account.Email, a.Account.Email)
				}
				if a.Account.Username != account.Account.Username {
					return fmt.Errorf("create account username [%v] but got [%v]", account.Account.Username, a.Account.Username)
				}
				if a.Account.Password != account.Account.Password {
					return fmt.Errorf("create account passowrd [%v] but got [%v]", account.Account.Password, a.Account.Password)
				}
				return nil
			},
			wantErr: false,
		},
		{
			name: "update_account_should_success",
			op: func() error {
				update := &AccountUpdate{
					Email:    "update@email.com",
					Username: "update",
				}
				a, err := m.UpdateAccount(context.Background(), objid.ToAccountID(account.ID), update)
				if err != nil {
					return err
				}
				if a.Account.Email != update.Email {
					return fmt.Errorf("update email [%v] but got [%v]", update.Email, a.Account.Email)
				}
				if a.Account.Username != update.Username {
					return fmt.Errorf("update account username [%v] but got [%v]", update.Username, a.Account.Username)
				}
				return nil
			},
			wantErr: false,
		},
		{
			name: "change_password_should_success",
			op: func() error {
				newPassowrd := "newPassword"
				account.Account.Password = newPassowrd
				return m.ChangePassword(context.Background(), objid.ToAccountID(account.ID), newPassowrd)
			},
			wantErr: false,
		},
		{
			name: "check_change_password_should_success",
			op: func() error {
				a, err := m.GetAccount(context.Background(), objid.ToAccountID(account.ID))
				if err != nil {
					return err
				}
				if a.Account.Password != account.Account.Password {
					return fmt.Errorf("chnage account passowrd [%v] but got [%v]", account.Account.Password, a.Account.Password)
				}
				return nil
			},
			wantErr: false,
		},
	}

	for _, cc := range cases {
		err := cc.op()
		if cc.wantErr && err == nil {
			t.Errorf("%s: want error; got none", cc.name)
		}
		if !cc.wantErr && err != nil {
			t.Errorf("%s: operation failed: %v", cc.name, err)
		}
	}
}

func TestMain(m *testing.M) {
	os.Exit(mongotesting.RunWithMongoInDocker(m))
}
