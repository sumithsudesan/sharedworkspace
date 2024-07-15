package database

// Interface UserRepository
// to different different Db types
type UserRepository interface {
    SaveUser(username, password string) error
    GetUserByUsername(username string) (*User, error)
    UpdateUserPassword(username, newPassword string) error
    DeleteUserByUsername(username string) error
}