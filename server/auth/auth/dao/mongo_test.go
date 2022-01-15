package dao

import (
	"context"

	"notion/shared/id"
	mgutil "notion/shared/mongo"
	"notion/shared/mongo/objid"
	mongotesting "notion/shared/mongo/testing"
	"os"
	"testing"
)

func TestAuthLifecycle(t *testing.T) {
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
	type accountRecord struct {
		mgutil.IDField `bson:"inline"`
		Account        struct {
			Username string
			Email    string
			Password string
		} `bson:"account"`
	}

	aid := id.AccountID("61e177a7e9a4d4c171abadfe")
	objID, err := objid.FromID(id.AccountID(aid))
	if err != nil {
		t.Fatalf("cannot converts an id to objected id: %+v", err)
	}
	var a accountRecord
	a.ID = objID
	a.Account.Username = "test_account"
	a.Account.Email = "test_email@gmail.com"
	a.Account.Password = "fjdsalgjlksadweuriwqjrklwfsafd"
	m.col.InsertOne(context.Background(), &a)
	cases := []struct {
		name    string
		op      func() error
		wantErr bool
	}{
		{
			name: "get_exists_account_should_success",
			op: func() error {
				_, err := m.Exists(context.Background(), aid)
				return err
			},
			wantErr: false,
		},
		{
			name: "get_no_exists_account_should_fail",
			op: func() error {
				_, err := m.Exists(context.Background(), id.AccountID(""))
				return err
			},
			wantErr: true,
		},
		{
			name: "resolve_exists_account_should_success",
			op: func() error {
				_, err := m.ResolveAccount(context.Background(), a.Account.Email)
				return err
			},
			wantErr: false,
		},
		{
			name: "resolve_no_exists_account_should_fail",
			op: func() error {
				_, err := m.ResolveAccount(context.Background(), "fake_email")
				return err
			},
			wantErr: true,
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
