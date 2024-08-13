package main

import (
	"context"
	"database/sql" // add this
	"fmt"
	"log"
	"net/url"
	"os"
	DAL "teamplayer/dal"
	"teamplayer/ent"

	CHAT "teamplayer/models"

	"github.com/gofiber/contrib/websocket"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html/v2"
	_ "github.com/lib/pq" // add this
)

// This is a test for a PLUGIN FOR GIT
// render the login page
func indexHandler(c *fiber.Ctx) error {
	// var res string
	// var todos []string
	// rows, err := db.Query("SELECT * FROM todos")
	// defer rows.Close()
	// if err != nil {
	// 	log.Fatalln(err)
	// 	c.JSON("An error occured")
	// }
	// for rows.Next() {
	// 	rows.Scan(&res)
	//
	// 	todos = append(todos, res)
	// }

	return c.Render("logindashboard", fiber.Map{
		//"Todos": todos,
	})
}

// add the login handler here

// need to validate, then authenticate
// either go home screen to re-login or go to the dashboard

// Structs for the application
type todo struct {
	Item string
}

func postHandler(c *fiber.Ctx, db *sql.DB) error {
	newTodo := todo{}
	if err := c.BodyParser(&newTodo); err != nil {
		log.Printf("An error occured: %v", err)
		return c.SendString(err.Error())
	}
	fmt.Printf("%v", newTodo)
	fmt.Printf("New item is added to the todos")
	if newTodo.Item != "" {
		_, err := db.Exec("INSERT into todos VALUES ($1)", newTodo.Item)
		if err != nil {
			log.Fatalf("An error occured while executing query: %v", err)
		}
	}

	return c.Redirect("/")
}

func loginUserHandler(c *fiber.Ctx, client *ent.Client, ctx context.Context) error {
	// this is to deal with returning users

	// validate request,
	return nil
}

// new user handler
func newUserHandler(c *fiber.Ctx, client *ent.Client, ctx context.Context) error {
	var checkError error
	if checkError = DAL.CheckUser(ctx, c, client); checkError != nil {
		return checkError
	}
	if checkError = DAL.CreateUser(ctx, c, client); checkError != nil {
		// return the message and take the user back to create an account
		return c.SendString(checkError.Error())
		// return c.Redirect("/")
	}
	// take the user to the dashboard
	// return c.SendString("Usser created successfully")

	return c.Render("logindashboard", fiber.Map{
		//"Todos": todos,
	})
}

// update user handler
func updateUserHandler(c *fiber.Ctx, client *ent.Client, ctx context.Context) error {
	// check to see if the desired user exists
	var checkError error

	if checkError = DAL.CheckUser(ctx, c, client); checkError != nil {
		return checkError
	}

	// if so then update
	if checkError = DAL.UpdateUser(ctx, c, client); checkError != nil {
		return checkError
	}

	// good status message
	// may need to build out custom messages tbh
	// or return nil
	return checkError
}

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
	connStr := "postgresql://postgres:gopher@localhost/todos?sslmode=disable"
	// Connect to database
	client, err := ent.Open(
		"postgres",
		"host=192.168.0.53 port=5432 user=postgres dbname=todos password=postgres sslmode=disable",
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	defer client.Close()

	ctx := context.Background()

	// running automigration
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	// static route and directory
	app.Static("/static/", "./static")

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Allow all origins, or specify your frontend domain
		AllowMethods: "GET,POST,PUT,DELETE",
	}))

	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)

			// can also handle taking apart token can't do that in websocket conn vs ctx
			log.Print("Upgraded web socket")

			c.Next()
		}

		// do not need to upgrade unless going to use web socket functionality
		return nil
	})

	chatHub := CHAT.NewChatroomServer()
	chatHubHandler := CHAT.NewChatRoomHandler(chatHub)

	go chatHub.StartServer()
	app.Post("/ws/createRoom", func(c *fiber.Ctx) error {
		return chatHubHandler.CreateNewRoom(c)
	})

	app.Get("/ws/joinRoom/:roomId", websocket.New(func(c *websocket.Conn) {
		log.Println("Made it into a chatroom")
		chatHubHandler.JoinRoom(c)
		// join a room
	}))

	// this will return the default login page
	app.Get("/", func(c *fiber.Ctx) error {
		// render the login and join form
		return indexHandler(c)
	})

	app.Get("/chatroom", func(c *fiber.Ctx) error {
		return c.Render("chatroom", fiber.Map{})
	})

	// This will deal with the post methods adding new todos, new users, new chatrooms, etc
	app.Post("/", func(cfib *fiber.Ctx) error {
		// adding user to the system
		return newUserHandler(cfib, client, ctx)
		// return postHandler(c, db)
	})

	// this is for a single parameter at the moment
	app.Put("/update/:olditem/:newitem", func(c *fiber.Ctx) error {
		return putHandler(c, db)
	})

	app.Delete("/delete", func(c *fiber.Ctx) error {
		return deleteHandler(c, db)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	app.Static("/", "./public")
	log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
}
