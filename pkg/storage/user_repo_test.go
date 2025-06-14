package storage

import "testing"

func TestUserRepo_GetUserByUserName_NotFound(t *testing.T) {
	memDB := UserMemDB{}
	repo := NewUserRepo(memDB)

	result := repo.GetUserByUserName("unknown")

	if result != nil {
		t.Errorf("Expected nil for unknown user, got %+v", result)
	}
}

func TestUserRepo_SetUserAndGetUser(t *testing.T) {
	memDB := UserMemDB{}
	repo := NewUserRepo(memDB)

	user := &UserModel{Name: "alice"}

	err := repo.SetUser(user)
	if err != nil {
		t.Fatalf("SetUser failed: %v", err)
	}

	got := repo.GetUserByUserName("alice")
	if got == nil {
		t.Fatal("Expected user to be found, got nil")
	}
	if got.Name != "alice" {
		t.Errorf("Expected name 'alice', got '%s'", got.Name)
	}
}

func TestUserRepo_SetUser_OverwriteExisting(t *testing.T) {
	memDB := UserMemDB{
		"bob": &UserModel{Name: "bob"},
	}
	repo := NewUserRepo(memDB)

	updatedUser := &UserModel{Name: "bob"}
	err := repo.SetUser(updatedUser)
	if err != nil {
		t.Fatalf("SetUser failed: %v", err)
	}

	got := repo.GetUserByUserName("bob")
	if got == nil {
		t.Fatal("Expected user to be found, got nil")
	}
	if got != updatedUser {
		t.Errorf("Expected user to be overwritten with %+v, got %+v", updatedUser, got)
	}
}
