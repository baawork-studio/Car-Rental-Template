package repositories

import "testing"

func TestNewRedisRepositoryRejectsInvalidURL(t *testing.T) {
	repository, err := NewRedisRepository("://invalid")
	if err == nil || repository != nil {
		t.Fatalf("expected invalid Redis URL to fail, got repository=%#v error=%v", repository, err)
	}
}

func TestNewRedisRepositoryAcceptsValidURL(t *testing.T) {
	repository, err := NewRedisRepository("redis://localhost:6379/0")
	if err != nil {
		t.Fatalf("expected valid Redis URL to succeed: %v", err)
	}
	if repository.Client == nil {
		t.Fatal("expected Redis client")
	}
	if err := repository.Close(); err != nil {
		t.Fatalf("close Redis client: %v", err)
	}
}
