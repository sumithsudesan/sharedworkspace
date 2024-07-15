package auth

import (
    "context"
    "log"
    "time"
    "github.com/golang-jwt/jwt"
    "golang.org/x/crypto/bcrypt"
    "data-definitions/auth"
    "data-definitions/user"
)

// AuthServer -  server
type AuthService struct {
    config Config
    client *AuthServiceClient
}

// new NewAuthService
func NewAuthService(config Config, pb.AuthDataServiceClient) *AuthService {
    return &AuthService{
        client: NewAuthClient(conn)
        config: config,
    }
}

// Login 
func (s *AuthService) Login(ctx context.Context, req *LoginRequest) (*LoginResponse, error) {
   // get user details from auth_database_service 
   user, err := s.client.GetUserByUsername(ctx, &GetUserByUsernameRequest{Username: req.Username})
   if err != nil {
       log.Printf("[ERROR] Failed to get user by username: %v", err)
       return nil, status.Errorf(codes.Internal, "Failed to get user")
   }
    // check the 
    if req.Username != username || req.Password != password {
        return nil, status.Errorf(codes.Unauthenticated, "invalid credentials")
    }

    token, err := s.generateJWTToken(username)
    if err != nil {
        log.Printf("[ERROR] Failed to generate JWT token: %v", err)
        return nil, status.Errorf(codes.Internal, "Failed to generate JWT token")
    }

    // Return a successful login response with the token
    return &LoginResponse{Token: token}, nil
}

func (s *AuthService) Register(ctx context.Context, req *RegisterRequest) (*RegisterResponse, error) {
    // validates user name and password
    if len(req.Username) < 4 {
        return nil, status.Errorf(codes.InvalidArgument, "Invalid user name, at least 4 characters")
    }
    if len(req.Password) < 6 {
        return nil, status.Errorf(codes.InvalidArgument, "Invalid password, at least 6 characters")
    }

    // Check password
    hashedPassword, err := hashPassword(req.Password)
    if err != nil {
        log.Printf("[ERROR] Failed to hash password: %v", err)
        return nil, status.Errorf(codes.Internal, "Failed to register user")
    }

    // Store the user details in DB 
    log.Printf("[INFO] Registering user: %s", req.Username)

    // In a real implementation, you would store the user details in a database
     // if err := storeUserInDatabase(req.Username, hashedPassword); err != nil {
    //     log.Printf("Failed to store user in database: %v", err)
    //     return nil, status.Errorf(codes.Internal, "could not register user")
    // }

    /// user registered successfully
    return &RegisterResponse{Message: "user registered successfully"}, nil
}

// Generate Authenication token for the user
func (s *AuthService) generateJWTToken(username string) (string, error) {
    // Create a new JWT token
    token := jwt.New(jwt.SigningMethodHS256)

    // Set claims
    claims := token.Claims.(jwt.MapClaims)
    claims["username"] = username
    // Token expiration time (24 hours)
    claims["exp"] = time.Now().Add(time.Hour * 24).Unix() 

    // Sign the token with the secret key from config
    signedToken, err := token.SignedString([]byte(s.config.JWTSecretKey()))
    if err != nil {
        return "", err
    }

    return signedToken, nil
}

// hashPassword hashes the given password using bcrypt
func hashPassword(password string) (string, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }
    return string(hashedPassword), nil
}