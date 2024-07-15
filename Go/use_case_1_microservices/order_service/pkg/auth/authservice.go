package auth

import (
    "context"
    "database/sql"
    "log"
    "time"

    "auth_service/src/config"

    _ "github.com/go-sql-driver/mysql"
    "github.com/golang-jwt/jwt"
)

// AuthServer -  server
type AuthService struct {
	UnimplementedAuthServiceServer
    db *sql.DB
}
func (s *AuthService) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
    // Authenticate user and return JWT token
    // Example implementation here
    return &LoginResponse{Token: "dummy-token"}, nil
}

func (s *AuthService) Register(ctx context.Context, req *RegisterRequest) (*RegisterResponse, error) {
    // Register new user
    // Example implementation here
    return &RegisterResponse{Message: "user registered"}, nil
}

func (s *AuthService) initializeDB(dsn string) {
    var err error
    s.db, err = sql.Open("mysql", dsn)
    if err != nil {
        log.Fatalf("[ERROR] Failed to connect to MySQL: %v", err)
    }
}

func (s *AuthService) generateJWT(userID string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user_id": userID,
        "exp":     time.Now().Add(time.Hour * 72).Unix(),
    })
    return token.SignedString([]byte(config.LoadConfig().JWTSecretKey))
}