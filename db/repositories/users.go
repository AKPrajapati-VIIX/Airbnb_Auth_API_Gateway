package db

import (
	"AuthInGo/models"
	"database/sql"
	"fmt"
)

type UserRepository interface {
	GetByID(id string) (*models.User, error)

	// Create() error
	Create(username string, email string, hashedPassword string) (*models.User, error) //error
	GetByEmail(email string) (*models.User, error)
	GetAll() ([]*models.User, error)
	// DeleteByID(id int64) error
	DeleteUserById(userID int64) error
}

type UserRepositoryImpl struct {
	db *sql.DB
}


func NewUserRepository(_db *sql.DB) UserRepository {
	return &UserRepositoryImpl{
		db: _db,
	}
}


// GetAll retrieves all users from the database
func (u *UserRepositoryImpl) GetAll() ([]*models.User, error) {
	query := "SELECT id, username, email, created_at, updated_at FROM users"
	rows, err := u.db.Query(query)
	if err != nil {
		fmt.Println("Error querying all users:", err)
		return nil, err
	}
	defer rows.Close()

	users := make([]*models.User, 0)

	for rows.Next() {
		user := &models.User{}
		err := rows.Scan(&user.Id, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt)
		if err != nil {
			fmt.Println("Error scanning user:", err)
			return nil, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		fmt.Println("Error iterating over rows:", err)
		return nil, err
	}

	return users, nil
}



// DeleteUserById deletes a user by ID from the database
//    RESTful API: DELETE /users/{id}
func (u *UserRepositoryImpl) DeleteUserById(id int64) error {
	query := "DELETE FROM users WHERE id = ?"
	result, err := u.db.Exec(query, id)
	if err != nil {
		fmt.Println("Error deleting user:", err)
		return err
	}
	rowsAffected, rowErr := result.RowsAffected()
	if rowErr != nil {
		fmt.Println("Error getting rows affected:", rowErr)
		return rowErr
	}
	if rowsAffected == 0 {
		fmt.Println("No rows were affected, user not deleted")
		return nil
	}
	fmt.Printf("User with ID %d deleted successfully. Rows affected: %d\n", id, rowsAffected)
	return nil
}



// Create inserts a new user into the database
// It takes username, email, and hashed password as parameters
func (u *UserRepositoryImpl) Create(username string, email string, hashedPassword string) (*models.User, error) {
	query := "INSERT INTO users (username, email, password) VALUES (?, ?, ?)"

	result, err := u.db.Exec(query, username, email, hashedPassword)

	if err != nil {
		fmt.Println("Error creating user:", err)
		return nil, err
	}

	// rowsAffected, rowErr := result.RowsAffected()
	lastInsertID, rowErr := result.LastInsertId()

	if rowErr != nil {
		fmt.Println("Error getting rows affected:", rowErr)
		return nil, rowErr
	}

	user := &models.User{
		Id:       lastInsertID,
		Username: username,
		Email:    email,
		
	}

	fmt.Println("User created successfully:", user)

	return user, nil
}

func (u *UserRepositoryImpl) GetByID( id string) (*models.User, error) {
	fmt.Println("Fetching user in UserRepository")

	// Step 1: Prepare the query
	query := "SELECT id, username, email, created_at, updated_at FROM users WHERE id = ?"

	// Step 2: Execute the query
	row := u.db.QueryRow(query, id)

	// Step 3: Process the result
	user := &models.User{}

	err := row.Scan(&user.Id, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No user found with the given ID")
			return nil, err
		} else {
			fmt.Println("Error scanning user:", err)
			return nil, err
		}
	}

	// Step 4: Print the user details
	fmt.Println("User fetched successfully:", user)

	return user, nil
}

func (r *UserRepositoryImpl) Save(user *models.User) error {
	// Add DB logic to insert user
	return nil
}

func (u *UserRepositoryImpl) GetByEmail(email string) (*models.User, error) {
	query := "SELECT id, email, password FROM users WHERE email = ?"

	row := u.db.QueryRow(query, email)

	user := &models.User{}

	err := row.Scan(&user.Id, &user.Email, &user.Password) // hashed password

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No user found with the given email")
			return nil, err
		} else {
			fmt.Println("Error scanning user:", err)
			return nil, err
		}
	}

	return user, nil
}