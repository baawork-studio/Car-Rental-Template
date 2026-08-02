package repositories

import (
	"context"
	"testing"
)

func TestNewPostgresRepositoryRejectsInvalidURL(t *testing.T) {
	repository, err := NewPostgresRepository(context.Background(), "://invalid")
	if err == nil || repository != nil {
		t.Fatalf("expected invalid PostgreSQL URL to fail, got repository=%#v error=%v", repository, err)
	}
}
