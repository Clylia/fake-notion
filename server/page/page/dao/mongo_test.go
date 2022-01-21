package dao

import (
	"context"
	"fmt"
	"os"
	"testing"

	"notion/shared/errs"
	"notion/shared/id"
	mgutil "notion/shared/mongo"
	"notion/shared/mongo/objid"
	mongotesting "notion/shared/mongo/testing"
)

func TestPageLifecycle(t *testing.T) {
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
	now := mgutil.UpdatedAt()
	accountID := mgutil.NewObjID()
	page := PageRecord{
		CreatedAtField: mgutil.CreatedAtField{
			CreatedAt: now,
		},
		UpdatedAtField: mgutil.UpdatedAtField{
			UpdatedAt: now,
		},
		CreatorID: accountID,
		Blocks: []*BlockEmtity{
			{
				IDField: mgutil.IDField{
					ID: mgutil.NewObjID(),
				},
				HTML:     "<div>hello notion<div>",
				Tag:      "Html",
				ImageURL: "",
			},
		},
	}
	cases := []struct {
		name    string
		op      func() error
		wantErr bool
	}{
		{
			name: "add_page_should_success",
			op: func() error {
				p, err := m.CreatePage(context.Background(), objid.ToAccountID(accountID), page.Blocks)
				if err != nil {
					return err
				}

				page.ID = objid.MustFromID(id.PageID(p.Id))
				return nil
			},
			wantErr: false,
		},
		{
			name: "update_page_should_success",
			op: func() error {
				blocks := []*BlockEmtity{
					{
						IDField: mgutil.IDField{
							ID: mgutil.NewObjID(),
						},
						HTML:     "<div>hello notion2<div>",
						Tag:      "Html",
						ImageURL: "",
					},
				}
				p, err := m.UpdatePage(context.Background(), objid.ToPageID(page.ID), objid.ToAccountID(accountID), blocks)
				if err != nil {
					return err
				}
				if len(p.Blocks) != 1 {
					return fmt.Errorf("expected update page has %v block but got %v", len(blocks), len(p.Blocks))
				}
				for i := 0; i < len(blocks); i++ {
					if blocks[i].ID.Hex() != p.Blocks[i].Id {
						return fmt.Errorf("expected update page block id has %v but got %v", blocks[i].ID.Hex(), p.Blocks[i].Id)
					}

					if blocks[i].HTML != p.Blocks[i].Html {
						return fmt.Errorf("expected update page block HTML has %v but got %v", blocks[i].HTML, p.Blocks[i].Html)
					}
				}
				return nil
			},
			wantErr: false,
		},
		{
			name: "gets_pages_should_success",
			op: func() error {
				_, err := m.CreatePage(context.Background(), objid.ToAccountID(accountID), page.Blocks)
				if err != nil {
					return err
				}
				pages, err := m.GetPages(c, objid.ToAccountID(accountID))
				if err != nil {
					return err
				}
				if len(pages) != 2 {
					return fmt.Errorf("expected page has 2 but got %d", len(pages))
				}
				return nil
			},
			wantErr: false,
		},
		{
			name: "delete_page_should_success",
			op: func() error {
				return m.DeletePage(context.Background(), objid.ToPageID(page.ID), objid.ToAccountID(accountID))
			},
			wantErr: false,
		},
		{
			name: "check_delete_page_should_success",
			op: func() error {
				_, err := m.GetPage(context.Background(), objid.ToPageID(page.ID), objid.ToAccountID(accountID))
				if errs.IsNoDocumentsErr(err) {
					return nil
				}
				return err
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
