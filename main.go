package main

import (
	"html/template"
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	// "github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"

	// "github.com/gofiber/helmet/v2"
	fiberTemplating "github.com/gofiber/template/html/v2"

	// "github.com/golang-jwt/jwt/v5"

	// _ "auta/docs"
	// swag "github.com/gofiber/swagger"
	// "github.com/gofiber/contrib/swagger"

	"github.com/joho/godotenv"
)

var htmlContainerConfig = fiber.Map{
	"Title": "Go fiber template",
	// "USE_CDN":     true,
	"USE_HTMX":        true,
	"USE_FLEXBOXGRID": true,
	"USE_BULMA":       true,
	"EMBED_VIEWS":     "{{embed}}",
	// "USE_HYPERSCRIPT": true,
	// "USE_PICOCSS":     true,
	// "USE_UIKIT":   true,
	// "USE_MATERIALIZE": true,
}

func PlainPageRender(name string) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		return c.Render(name, nil)
	}
}

func CreateCar(c *fiber.Ctx) error {
	mForm, err := c.MultipartForm()
	if err != nil {
		panic(err)
	}
	files := mForm.File["carPhotos"]
	for i, file := range files {
		log.Println(i, file.Filename)
	}

	return c.SendStatus(fiber.StatusAccepted)
}

// custom middlewares
type FiberMiddleware func(c *fiber.Ctx) error

// This verifies that the incoming http Connect-Type header matches contentType
func BuildMiddlewareEnsureMIME(contentType string) FiberMiddleware {
	return func(c *fiber.Ctx) error {
		// Compare the Content-Type header with the expected value
		contentType := c.Get(fiber.HeaderContentType)
		if strings.HasPrefix(contentType, contentType) {
			// The request has the correct Content-Type for form submission
			// You can proceed with form processing here
			return c.Next()
		} else {
			// The request does not have the expected Content-Type
			// You can return an error response or handle it as needed
			return c.Status(fiber.StatusBadRequest).SendString("Invalid Content-Type")
		}
	}
}

func main() {
	godotenv.Load()
	log.SetOutput(os.Stdout)
	// connectToSupportServices()
	// setupModels(dbConn)

	// generate HTML boilerplate template from the index template
	HTMLBoilerplateTemplate, err := template.ParseFiles("./views/index.html")
	if err != nil {
		panic(err)
	}
	generatedHTMLBoilerplate, err := os.Create("./views/generated/index.html")
	if err != nil {
		panic(err)
	}
	err = HTMLBoilerplateTemplate.Execute(generatedHTMLBoilerplate, htmlContainerConfig)
	if err != nil {
		panic(err)
	}
	err = generatedHTMLBoilerplate.Close()
	if err != nil {
		panic(err)
	}

	engine := fiberTemplating.New("./", ".html")
	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "views/generated/index",
	})

	// fiber middlewares
	{
		// set up logging
		app.Use(logger.New())
		// set up swagger
		// cfg := swagger.Config{
		// 	BasePath: "/",
		// 	FilePath: "./docs/swagger.json",
		// 	Path:     "swagger",
		// 	Title:    "Swagger API Docs",
		// }
		// app.Use(swagger.New(cfg))
		// app.Get("/docs/*", swag.HandlerDefault)

		// security stuffies
		// app.Use(helmet.New())
		config := limiter.ConfigDefault
		config.LimitReached = func(c *fiber.Ctx) error {
			log.Println("[fiber rate limit exceeded]", c.IP())
			return c.SendStatus(fiber.StatusTooManyRequests)
		}
		// config.Storage = ...
		config.Max = 100
		app.Use(limiter.New(config))

		// TODO: CORS, CSRF
	}
	// app.Use(cors.New(cors.Config{
	// 	AllowOriginsFunc: func(origin string) bool {
	// 		return os.Getenv("DEPLOYMENT") == "dev"
	// 	},
	// 	AllowOrigins:  "*",
	// 	ExposeHeaders: "*",
	// }))

	// set up routes
	app.Get("/", PlainPageRender("views/pages/home"))
	app.Get("/robots.txt", func(c *fiber.Ctx) error { return c.SendString("") }) // TODO:
	app.Post("/createCar", CreateCar)
	app.Get("/404", PlainPageRender("views/pages/404"))
	app.Get("/401", PlainPageRender("views/pages/401"))
	app.Get("/favicon.ico", func(c *fiber.Ctx) error { return c.Redirect("https://bulma.io/images/bulma-logo.png") })
	app.Static("/static", "./static")

	log.Fatal(app.Listen(":3300"))
}
