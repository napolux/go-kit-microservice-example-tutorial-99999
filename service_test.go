package napodate

import (
	"context"
	"testing"
	"time"
)

func TestStatus(t *testing.T) {
	srv, ctx := setup()

	s, err := srv.Status(ctx)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	// testing status
	ok := s == "ok"
	if !ok {
		t.Errorf("expected service to be ok")
	}
}

func TestGet(t *testing.T) {
	srv, ctx := setup()
	d, err := srv.Get(ctx)
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	time := time.Now()
	today := time.Format("02/01/2006")

	// testing today's date
	ok := today == d
	if !ok {
		t.Errorf("expected dates to be equal")
	}
}
func TestValidate(t *testing.T) {
	srv, ctx := setup()
	b, err := srv.Validate(ctx, "31/12/2019")
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	// testing that the date is valid
	if !b {
		t.Errorf("date should be valid")
	}
	// testing an invalid date
	b, err = srv.Validate(ctx, "31/31/2019")
	if b {
		t.Errorf("date should be invalid")
	}

	// testing a USA date date
	b, err = srv.Validate(ctx, "12/31/2019")
	if b {
		t.Errorf("USA date should be invalid")
	}
}

func setup() (srv Service, ctx context.Context) {
	return NewService(), context.Background()
}
