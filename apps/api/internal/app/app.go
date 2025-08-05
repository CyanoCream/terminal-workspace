package app

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"

	authHandler "terminal/internal/auth/adapter/handler"
	authRepo "terminal/internal/auth/adapter/repository"
	authPort "terminal/internal/auth/port"
	authService "terminal/internal/auth/service"

	terminalHandler "terminal/internal/terminal/adapter/handler"
	terminalRepo "terminal/internal/terminal/adapter/repository"
	terminalPort "terminal/internal/terminal/port"
	terminalService "terminal/internal/terminal/service"

	userHandler "terminal/internal/user/adapter/handler"
	userRepo "terminal/internal/user/adapter/repository"
	userPort "terminal/internal/user/port"
	userService "terminal/internal/user/service"

	transactionHandler "terminal/internal/transaction/adapter/handler"
	transactionRepo "terminal/internal/transaction/adapter/repository"
	transactionPort "terminal/internal/transaction/port"
	transactionService "terminal/internal/transaction/service"

	"terminal/pkg/config"
	"terminal/pkg/database"
	"terminal/pkg/jwt"
)

type App struct {
	DB     *gorm.DB
	Server *fiber.App
}

func NewApp() (*App, error) {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		return nil, err
	}

	// Initialize database
	db, err := database.NewPostgresConnection(cfg.DBURL)
	if err != nil {
		return nil, err
	}

	// Initialize repositories
	authRepo := authRepo.NewAuthRepository(db)
	terminalRepo := terminalRepo.NewTerminalRepository(db)
	userRepo := userRepo.NewUserRepository(db)
	transactionRepo := transactionRepo.NewTransactionRepository(db)

	// Initialize services
	authSvc := authService.NewAuthService(authRepo, cfg.JWTSecret, cfg.JWTExpiry)
	terminalSvc := terminalService.NewTerminalService(terminalRepo)
	userSvc := userService.NewUserService(userRepo)
	transactionSvc := transactionService.NewTransactionService(
		transactionRepo,
		terminalSvc,
		userRepo,
	)

	// Initialize handlers
	authHandler := authHandler.NewAuthHandler(authSvc)
	terminalHandler := terminalHandler.NewTerminalHandler(terminalSvc)
	userHandler := userHandler.NewUserHandler(userSvc)
	transactionHandler := transactionHandler.NewTransactionHandler(transactionSvc)

	// Create Fiber app
	app := fiber.New()

	// Setup routes
	setupRoutes(app, authHandler, terminalHandler, userHandler, transactionHandler, cfg.JWTSecret)

	return &App{
		DB:     db,
		Server: app,
	}, nil
}

func setupRoutes(
	app *fiber.App,
	authHandler authPort.AuthHandler,
	terminalHandler terminalPort.TerminalHandler,
	userHandler userPort.UserHandler,
	transactionHandler transactionPort.TransactionHandler,
	jwtSecret string,
) {
	api := app.Group("/api")

	// Auth routes
	auth := api.Group("/auth")
	auth.Post("/login", authHandler.Login)
	auth.Post("/register", authHandler.Register)

	// Protected routes
	protected := api.Group("", jwt.JWTProtected(jwtSecret))

	// Terminal routes
	protected.Get("/terminals", terminalHandler.GetAllTerminals)
	protected.Post("/terminals", terminalHandler.CreateTerminal)
	protected.Post("/terminals/:id/gates", terminalHandler.AddGate)
	protected.Post("/terminal-pricing", terminalHandler.SetPricing)

	// User routes
	protected.Get("/users/me", userHandler.GetProfile)
	protected.Post("/cards", userHandler.CreateCard)
	protected.Post("/cards/topup", userHandler.TopUpCard)

	// Transaction routes
	protected.Post("/transactions/checkin", transactionHandler.CheckIn)
	protected.Post("/transactions/checkout", transactionHandler.CheckOut)
	protected.Post("/transactions/sync", transactionHandler.SyncTransactions)
	protected.Get("/transactions", transactionHandler.GetTransactionHistory)
}
