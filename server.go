package main

import (
   "database/sql" // add this
   "fmt"
   "log"
   "os"
   "github.com/gofiber/fiber/v2"
   "github.com/gofiber/contrib/websocket"
   "github.com/gofiber/fiber/v2/middleware/cors"
   _"github.com/lib/pq" // add this
   "github.com/gofiber/template/html/v2"
   "net/url"
)

func indexHandler(c *fiber.Ctx, db *sql.DB) error {
   var res string
   var todos []string
   rows, err := db.Query("SELECT * FROM todos")
   defer rows.Close()
   if err != nil {
       log.Fatalln(err)
       c.JSON("An error occured")
   }
   for rows.Next() {
       rows.Scan(&res)
       todos = append(todos, res)
   }
   return c.Render("index", fiber.Map{
       "Todos": todos,
   })
}


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
   if newTodo.Item != "" {
       _, err := db.Exec("INSERT into todos VALUES ($1)", newTodo.Item)
       if err != nil {
           log.Fatalf("An error occured while executing query: %v", err)
       }
   }

   return c.Redirect("/")
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
   
   //producing 405 error
   // should update on the rocket guideline about this 
   //return c.Redirect("/")
}

func deleteHandler(c *fiber.Ctx, db *sql.DB) error {
   todoToDelete := c.Query("item")
   db.Exec("DELETE from todos WHERE item=$1", todoToDelete)
   return c.SendString("deleted")
}

func main() {
   connStr := "postgresql://postgres:gopher@localhost/todos?sslmode=disable"
   // Connect to database
   db, err := sql.Open("postgres", connStr)
   if err != nil {
       log.Fatal(err)
   }
   engine := html.New("./views", ".html")
   app := fiber.New(fiber.Config{
       Views: engine,
   })


	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Allow all origins, or specify your frontend domain
		AllowMethods: "GET,POST,PUT,DELETE",
	}))


	app.Use("/ws", func(c *fiber.Ctx) error {
	    if websocket.IsWebSocketUpgrade(c) {
		c.Locals("allowed", true)
		return c.Next()
	    }
	    return fiber.ErrUpgradeRequired
	})

	app.Get("/ws/:id", websocket.New(func(c *websocket.Conn) {
	    // Handle WebSocket connection here

	    var (
		    messageType int
		    msg		[]byte
		    err		error
	    )

	    // these access local variables
	    log.Println(c.Locals("allowed"))
            log.Println(c.Params("id"))
            log.Println(c.Query("v"))
	    log.Println(c.Cookies("session"))

	    for {

		    if messageType, msg, err = c.ReadMessage(); err != nil {
			    log.Println("read error :", err)
			    break
		    }

		    response := string(msg)
		    log.Printf("Chatroom: %s", c.Params("id"))
		    log.Printf("Message: %s", msg)



		    if err = c.WriteMessage(messageType, []byte(fmt.Sprintf("Server: %s", response))); err != nil {
			    log.Println("Write error: ", err)
			    break
		    }
	    }
	}))

   app.Get("/", func(c *fiber.Ctx) error {
       return indexHandler(c, db)
   })

   app.Post("/", func(c *fiber.Ctx) error {
       return postHandler(c, db)
   })

   //this is for a single parameter at the moment
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
