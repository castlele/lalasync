package service

import (
	"testing"

	"github.com/castlele/lalasync/pkg/auth/models"
	"github.com/castlele/lalasync/pkg/storage"
)

func TestAuthService_Register_NewUser(t *testing.T) {
	memDB := storage.UserMemDB{}
	repo := storage.NewUserRepo(memDB)
	user := &models.UserLogin{Name: "alice", Password: "secret"}
	sut := NewAuthService(repo)

	err := sut.Register(user)

	if err != nil {
		t.Fatalf("Expected registration to succeed, got error: %v", err)
	}
	if user.JWT == "" {
		t.Error("Expected JWT to be set on successful registration")
	}
}

func TestAuthService_Register_ExistingUser(t *testing.T) {
	memDB := storage.UserMemDB{}
	repo := storage.NewUserRepo(memDB)
	user := &models.UserLogin{Name: "alice", Password: "secret"}
	sut := NewAuthService(repo)

	_ = sut.Register(user)
	err := sut.Register(user)

	if err == nil || err.Error() != UserRegisteredError {
		t.Errorf("Expected error when registering existing user, got %v", err)
	}
}

func TestAuthService_Register_InvalidInput(t *testing.T) {
	memDB := storage.UserMemDB{}
	repo := storage.NewUserRepo(memDB)
	sut := NewAuthService(repo)

	err := sut.Register(&models.UserLogin{Name: "", Password: ""})

	if err == nil {
		t.Fatal("Expected error for invalid input, got nil")
	}
}

func TestAuthService_Login_Success(t *testing.T) {
	memDB := storage.UserMemDB{}
	repo := storage.NewUserRepo(memDB)
	sut := NewAuthService(repo)
	user := &models.UserLogin{Name: "bob", Password: "hunter2"}
	_ = sut.Register(user)

	loginUser := &models.UserLogin{Name: "bob", Password: "hunter2"}
	err := sut.Login(loginUser)

	if err != nil {
		t.Fatalf("Expected login to succeed, got error: %v", err)
	}

	if loginUser.JWT == "" {
		t.Error("Expected JWT to be set on successful login")
	}
}

func TestAuthService_Login_InvalidUser(t *testing.T) {
	memDB := storage.UserMemDB{}
	repo := storage.NewUserRepo(memDB)
	sut := NewAuthService(repo)

	err := sut.Login(&models.UserLogin{Name: "ghost", Password: "1234"})

	if err == nil || err.Error() != UnknownUserError {
		t.Errorf("Expected error when logging in with non-existent user, got %v", err)
	}
}

func TestAuthService_Login_WrongPassword(t *testing.T) {
	memDB := storage.UserMemDB{}
	repo := storage.NewUserRepo(memDB)
	sut := NewAuthService(repo)

	_ = sut.Register(&models.UserLogin{Name: "eve", Password: "correct-password"})

	err := sut.Login(&models.UserLogin{Name: "eve", Password: "wrong-password"})

	if err == nil || err.Error() != InvalidPasswordError {
		t.Errorf("Expected error for wrong password, got %v", err)
	}
}
