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

    // Store user details in DB (assuming this is done in your client)
    user := &pb.User{
        Username:     req.Username,
        PasswordHash: hashedPassword,
        Email:        req.Email,
    }

    // Call AuthDataService to save user details
    _, err = s.client.SaveUser(ctx, &pb.SaveUserRequest{User: user})
    if err != nil {
        log.Printf("[ERROR] Failed to save user: %v", err)
        return nil, status.Errorf(codes.Internal, "Failed to register user")
    }

    /// user registered successfully
    return &RegisterResponse{Message: "user registered successfully"}, nil
}

// Implement Authenticate method for AuthService
func (s *AuthService) Authenticate(ctx context.Context, req *pb.AuthenticateRequest) (*pb.AuthenticateResponse, error) {
    token := req.Token

    // Validate JWT token
    claims, err := s.validateJWTToken(token)
    if err != nil {
        return nil, status.Errorf(codes.Unauthenticated, "Invalid token")
    }

    // Token is valid
    return &pb.AuthenticateResponse{Valid: true, Message: "Token is valid", Claims: claims}, nil
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

// Validate JWT token
func (s *AuthService) validateJWTToken(tokenString string) (jwt.MapClaims, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        // Validate the token signing method
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
        }
        return []byte(s.config.JWTSecretKey()), nil
    })
    if err != nil {
        return nil, err
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
        return claims, nil
    }
    return nil, fmt.Errorf("Invalid token")
}

// hashPassword hashes the given password using bcrypt
func hashPassword(password string) (string, error) {
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        return "", err
    }
    return string(hashedPassword), nil
}