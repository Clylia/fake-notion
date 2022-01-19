package dao

import (
	"context"
	"fmt"
	"os"
	"testing"
	"time"

	"notion/shared/errs"
	"notion/shared/id"
	"notion/shared/mongo/objid"
	mongotesting "notion/shared/mongo/testing"

	pagepb "notion/page/api/gen/v1"
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
	now := time.Now().Unix()
	accountID := id.AccountID("aid")
	page := PageRecord{
		Page: pagepb.Page{
			CreatorId: accountID.String(),
			Blocks: []*pagepb.BlockEmtity{
				{
					Id: "3341dd8d-a9b2-48c1-99ff-ef31a1a7c4f2",
					Block: &pagepb.Block{
						Html:     "<div>hello notion<div>",
						Tag:      "Html",
						ImageUrl: "",
					},
				},
			},
			CreatedAt: int32(now),
			UpdatedAt: int32(now),
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
				p, err := m.CreatePage(context.Background(), &page)
				if err != nil {
					return err
				}
				page.ID = p.ID
				return nil
			},
			wantErr: false,
		},
		{
			name: "update_page_should_success",
			op: func() error {
				update := &PageUpdate{
					Blocks: []*pagepb.BlockEmtity{
						{
							Id: "192b70d5-f0f5-4e43-a0f1-af48da3be186",
							Block: &pagepb.Block{
								Html:     "<div>hello notion2<div>",
								Tag:      "Html",
								ImageUrl: "",
							},
						},
					},
					CreatorID: accountID,
				}
				p, err := m.UpdatePage(context.Background(), objid.ToPageID(page.ID), update)
				if err != nil {
					return err
				}
				if len(p.Blocks) != 1 {
					return fmt.Errorf("expected update page has %v block but got %v", len(update.Blocks), len(p.Blocks))
				}
				for i := 0; i < len(update.Blocks); i++ {
					if update.Blocks[i].Id != p.Blocks[i].Id {
						return fmt.Errorf("expected update page block id has %v but got %v", update.Blocks[i].Id, p.Blocks[i].Id)
					}

					if update.Blocks[i].Block.Html != p.Blocks[i].Block.Html {
						return fmt.Errorf("expected update page block Html has %v but got %v", update.Blocks[i].Block.Html, p.Blocks[i].Block.Html)
					}
				}
				return nil
			},
			wantErr: false,
		},
		{
			name: "delete_page_should_success",
			op: func() error {
				return m.DeletePage(context.Background(), objid.ToPageID(page.ID), accountID)
			},
			wantErr: false,
		},
		{
			name: "check_delete_page_should_success",
			op: func() error {
				_, err := m.GetPage(context.Background(), objid.ToPageID(page.ID))
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
