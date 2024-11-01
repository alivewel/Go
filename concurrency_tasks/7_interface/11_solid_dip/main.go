package main

import (
	"fmt"
)

// User - структура пользователя
type User struct {
	ID   int
	Name string
}

// UserRepository - интерфейс для работы с пользователями
type UserRepository interface {
	GetUserByID(id int) (*User, error)
}

// InMemoryUserRepository - реализация UserRepository, работающая с данными в памяти
type InMemoryUserRepository struct {
	users map[int]*User
}

func (repo *InMemoryUserRepository) GetUserByID(id int) (*User, error) {
	user, exists := repo.users[id]
	if !exists {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}

// UserService - сервис для работы с пользователями
type UserService struct {
	repo UserRepository // Зависимость от абстракции
}

func (service *UserService) GetUser(id int) (*User, error) {
	return service.repo.GetUserByID(id)
}

func main() {
	// Создаем репозиторий и добавляем пользователей
	repo := &InMemoryUserRepository{
		users: map[int]*User{
			1: {ID: 1, Name: "Alice"},
			2: {ID: 2, Name: "Bob"},
		},
	}

	// Создаем сервис и получаем пользователя
	service := UserService{repo: repo}
	user, err := service.GetUser(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("User found:", user.Name)
}
