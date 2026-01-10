package main

import (
	"context"
	"database/sql" // add this
	"fmt"
	"log"
	"net/url"
	"os"
	"time"
	middleware "teamplayer/middleware"
	BLL "teamplayer/bll/auth"
	DAL "teamplayer/dal"
	models "teamplayer/models"
	viewmodels "teamplayer/models/viewmodels"
	"teamplayer/ent"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/template/html/v2"
	_ "github.com/lib/pq" // add this
)

// This is a test for a PLUGIN FOR GIT
// render the login page
func indexHandler(c *fiber.Ctx) error {
	

	return c.Render("index", fiber.Map{
		//"Todos": todos,
	})
}

// gets the new user form so a user get created.
func newUserGetHandler(c *fiber.Ctx) error {

	// new view model
	vm := viewmodels.SignUpViewModel{
		Form:   models.UserSignUp{},
        Errors: models.SignUpErrors{},
	}
	log.Println("NEWWWW UUUUUSSSESARERERERER method to create a new user")
	return c.Render("createuser", vm)
}


func newUserPostHandler(c *fiber.Ctx, client *ent.Client, ctx context.Context) error {

	log.Println("Post method to create a new user")

	// new view model
	vm := viewmodels.SignUpViewModel{}

	// Get the request for the form data
	var NewUserObj models.UserSignUp

	//parse the object to get user information
	if err := c.BodyParser(&NewUserObj); err != nil {
		// handle bad request data maybe a code and 
		// information to tell the user to retry
	}

	// assign into View Model
	vm.Form = NewUserObj


	//either format object to go to bll or bll will handle it
	if SignUpErr := BLL.UserSignUp(ctx, client , NewUserObj); SignUpErr != nil {

		// get the fields that have errors and assign error messages
		vm = BLL.CreateAccError(SignUpErr, vm)

		// clear password
		vm.Form.Password = ""

		// return with errors
		return c.Render("create_user_form", vm)

	}
		// // return error minimal information, db related
		// log.Printf("signup failed: %v", SignUpErr)

		// return c.Status(fiber.StatusInternalServerError).Render("createuser", fiber.Map{
		// 	"values": fiber.Map{
		// 		"first_name": NewUserObj.FirstName,
		// 		"last_name":  NewUserObj.LastName,
		// 		"email":      NewUserObj.Email,
		// 		"user_name":  NewUserObj.UserName,
		// 	},
		// 	"globalError": "Something went wrong while creating your account. Please try again.",
		// })

	// redirect 
	if c.Get("HX-Request") == "true" {
		c.Set("HX-Redirect", "logindashboard")
		return nil
	}
	
	return c.Redirect("logindashboard")


}
// add the login handler here

// need to validate, then authenticate
// either go home screen to re-login or go to the dashboard



// func postHandler(c *fiber.Ctx, db *sql.DB) error {
// 	newTodo := todo{}
// 	if err := c.BodyParser(&newTodo); err != nil {
// 		log.Printf("An error occured: %v", err)
// 		return c.SendString(err.Error())
// 	}
// 	fmt.Printf("%v", newTodo)
// 	fmt.Printf("New item is added to the todos")
// 	if newTodo.Item != "" {
// 		_, err := db.Exec("INSERT into todos VALUES ($1)", newTodo.Item)
// 		if err != nil {
// 			log.Fatalf("An error occured while executing query: %v", err)
// 		}
// 	}

// 	return c.Redirect("/")
// }

func loginUserHandler(c *fiber.Ctx, client *ent.Client, ctx context.Context) error {
	// this is to deal with returning users

	// validate request,
	return nil
}

// new user handler
// func newUserHandler(c *fiber.Ctx, client *ent.Client, ctx context.Context) error {
// 	// var checkError error
// 	// if checkError = DAL.CheckUser(ctx, c, client); checkError != nil {
// 	// 	return checkError
// 	// }
	
// 	var newUser models.UserSignUp

// 	if err := c.BodyParser(&newUser); err != nil {
// 		log.Printf("An error occured while parsing new user data: %v", err)
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 			"error": "invalid new user sign up body",
// 		})
// 	}
	
// 	err := BLL.UserSignUp(ctx, client, newUser)
// 	if err != nil {
// 		switch e := err.(type) {
	
// 		case models.NewUserValidationError:
// 			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 				"field":   e.SignUpField,
// 				"message": e.ValidationMessage,
// 			})
	
// 		case *ent.ConstraintError:
// 			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
// 				"error": "username or email already exists",
// 			})
	
// 		default:
// 			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
// 				"error": "internal server error",
// 			})
// 		}
// 	}
// 	return c.SendStatus(fiber.StatusCreated)
// }

// update user handler
// func updateUserHandler(c *fiber.Ctx, client *ent.Client, ctx context.Context) error {
// 	// check to see if the desired user exists
// 	var checkError error

// 	if checkError = DAL.CheckUser(ctx, c, client); checkError != nil {
// 		return checkError
// 	}

// 	// if so then update
// 	if checkError = DAL.UpdateUser(ctx, c, client); checkError != nil {
// 		return checkError
// 	}

// 	// good status message
// 	// may need to build out custom messages tbh
// 	// or return nil
// 	return checkError
// }

func putHandler(c *fiber.Ctx, db *sql.DB) error {
	olditem, err := url.PathUnescape(c.Params("olditem"))
	if err != nil {
		return err // handle error
	}

	newitem, err := url.PathUnescape(c.Params("newitem"))
	if err != nil {
		return err // handle error
	}

	db.Exec("UPDATE todos SET item=$1 WHERE item=$2", newitem, olditem)

	return c.SendString("updated")

	// producing 405 error
	// should update on the rocket guideline about this
	// return c.Redirect("/")
}

func deleteHandler(c *fiber.Ctx, db *sql.DB) error {
	todoToDelete := c.Query("item")
	db.Exec("DELETE from todos WHERE item=$1", todoToDelete)
	return c.SendString("deleted")
}

// end of sql stuff

// start of websocket stuff
func main() {
	ctx := context.Background()

	dbConn, err := DAL.NewDbConnection(ctx)

	if err != nil {
		log.Fatalf("failed to init DB: %v", err)
	}

	defer dbConn.SQL.Close()
	defer dbConn.Ent.Close()


	engine := html.New("./views", ".html")
	
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	chatHub := models.NewChatRoomServer()

	// Create the default "lobby" room can convert to a function later
	chatHub.Rooms["lobby"] = &models.ChatRoom{
		ID:      "lobby",
		Name:    "Lobby",
		Project: "default",
		Clients: make(map[string]*models.Client),
	}
	
	chatHubHandler := models.NewChatRoomHandler(chatHub)

	go chatHub.StartServer()

	// END of startup dependencies

	// Middleware section
	// static route and directory
	app.Static("/static/", "./static")

	// browser level, do I let this talk to the server?
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Allow all origins, or specify your frontend domain. When I decide to get a front end I need to specify
		AllowMethods: "GET,POST,PUT,DELETE",
	}))

	// session level 
	var store = session.New(session.Config{
		CookieSecure:    false,
		CookieHTTPOnly: true,
		CookieSameSite: "Lax",
		Expiration:    30 * time.Minute,
	})
	// app.Use(session.New(session.Config{
	// 	// Storage:           storage,		  // This will be changes when I get redis installed and working
	// 	CookieSecure:      false,              // HTTPS only, need to change when in testing or production
	// 	CookieHTTPOnly:    true,              // Prevent XSS
	// 	CookieSameSite:    "Lax",             // CSRF protection
	// 	IdleTimeout:       30 * time.Minute,  // Session timeout
	// 	AbsoluteTimeout:   24 * time.Hour,    // Maximum session life
	// 	// Extractor:         extractors.FromCookie("__Host-session_id"), // will also change in production as well
	// }))

	// user level middleware

	app.Use(middleware.LoadUserSession(dbConn.Ent, store))


	// END of middleware section

	/// START public routes: login, signup, home page

	// this will return the default login page
	app.Get("/", func(c *fiber.Ctx) error {
		// render the login and join form
		log.Println("Home landing page.")
		return indexHandler(c)
	})
	
	app.Get("/signup", func(c *fiber.Ctx) error {
		log.Println("Creating a new user here")
		return newUserGetHandler(c)
	})

	app.Post("/signup", func(c *fiber.Ctx) error {
		// adding user to the system
		return newUserPostHandler(c, dbConn.Ent, ctx)
		
	})
	// END public routes

	// START protected routes
	// This will deal with the post methods adding new todos, new users, new chatrooms, etc
	auth := app.Group("/", middleware.RequireAuth)
	

	// database health check
	auth.Get("/health/db", func(c *fiber.Ctx) error {
		if err := dbConn.DbHealthCheck(c.Context()); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
		}
		return c.SendString("ok")
	})

	// GET requested chatroom
	auth.Get("/chatroom", func(c *fiber.Ctx) error {
		log.Println("Here we goooo")
		return c.Render("chatroom", fiber.Map{})
	})


	// GET user logindashobard 
	auth.Get("/logindashboard", func(c *fiber.Ctx) error {
		log.Println("Successfully created user. On user home page")
		return c.Render("logindashboard", fiber.Map{})
	})

	// GET Websocket request
	auth.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			log.Println("Upgraded the websocket")
			return c.Next()
			// can also handle taking apart token can't do that in websocket conn vs ctx
		} 

		return fiber.ErrUpgradeRequired
	})

	// GET chatroom request
	auth.Get("/ws/joinRoom/:roomId", func(c *fiber.Ctx) error {
		u := c.Locals("currentUser").(*models.PublicUserProfile)
		c.Locals("user_uuid", u.UUID_Id.String())
		c.Locals("username", u.UserName)
		u, ok := c.Locals("currentUser").(*models.PublicUserProfile)
		if !ok || u == nil {
			return fiber.ErrUnauthorized
		}

		return websocket.New(chatHubHandler.JoinRoom)(c)
	})
	
	// POST Create chatroom request
	auth.Post("/ws/createRoom", func(c *fiber.Ctx) error {
		return chatHubHandler.CreateNewRoom(c)
	})

	// END protected routes
	
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	app.Static("/", "./public")
	log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
}
