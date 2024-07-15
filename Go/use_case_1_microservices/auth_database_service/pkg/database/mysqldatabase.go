package database

import (
    "database/sql"
    "log"

)

// MySQLUserRepository
type MySQLUserRepository struct {
    db *sql.DB
}

// creates new MySQLUserRepository instance
func NewMySQLUserRepository(db *sql.DB) UserRepository {
    return &MySQLUserRepository{
        db: db,
    }
}

// Add user details
func (r *MySQLUserRepository) SaveUser(user *User) error {
    query := "INSERT INTO users (username, email, version) VALUES (?, ?, 1)"
    _, err := r.db.ExecContext(ctx, query, user.Username, user.Email)
    if err != nil {
        log.Printf("[ERROR] Failed to add user to MySQL: %v", err)
        return fmt.Errorf("failed to save user: %w", err)
    }
    return nil
}

// Get user details
func (r *MySQLUserRepository) GetUserByID(id int) (*User, error) {
    var user User
    query := "SELECT id, username, email, version FROM users WHERE id = ?"
    err := r.db.QueryRowContext(ctx, query, id).Scan(&user.ID, &user.Username, &user.Email, &user.Version)
    if err != nil {
        if err == sql.ErrNoRows {
            return nil, fmt.Errorf("user with ID %d not found: %w", id, err)
        }
        log.Printf("[ERROR] Failed to get user from MySQL: %v", err)
        return nil, fmt.Errorf("failed to fetch user: %w", err)
    }
    return &user, nil
}

// Update user details
func (r *MySQLUserRepository) UpdateUser(user *User) error {
    query := "UPDATE users SET username = ?, email = ?, version = version + 1 WHERE id = ? AND version = ?"
    result, err := r.db.ExecContext(ctx, query, user.Username, user.Email, user.ID, user.Version)
    if err != nil {
        log.Printf("[ERROR] failed to update user in MySQL: %v", err)
        return fmt.Errorf("failed to update user: %w", err)
    }

    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return fmt.Errorf("[ERROR] failed to get rows affected: %w", err)
    }

    if rowsAffected == 0 {
        return fmt.Errorf("no rows updated, possible version mismatch")
    }

    return nil
}

// Delete user
func (r *MySQLUserRepository) DeleteUser(id int) error {
    query := "DELETE FROM users WHERE id = ?"
    _, err := r.db.ExecContext(ctx, query, id)
    if err != nil {
        log.Printf("[ERROR] Failed to delete user from MySQL: %v", err)
        return fmt.Errorf("failed to delete user: %w", err)
    }
    return nil
}
