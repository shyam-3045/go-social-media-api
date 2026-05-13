	package server

	import( 
		"github.com/gofiber/fiber/v3"
		"github.com/gofiber/fiber/v3/middleware/recover"
	)

	var app *fiber.App

	func New() *fiber.App{
		return app
	}
	func Setup() {
		app = fiber.New(fiber.Config{
			ErrorHandler: errorHandler,
			BodyLimit:    16 * 1024 * 1024,
		})

		defer app.Use(notFoundHandler)
		defer app.Use(recover.New())
		
		middlewares(app)
		addRoutes(app)
		
		
		

	}
