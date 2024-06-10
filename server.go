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
   "bytes"
   "encoding/json"
   "html/template"
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

// end of sql stuff

// start of websocket stuff
type Message struct {
	Text       string `json:"text"`
}
// websocket structs 
type WebSocketServer struct {
	clients map[*websocket.Conn]bool
	broadcast chan *Message
}

func NewWebSocket() *WebSocketServer {
	return &WebSocketServer{
		clients: make(map[*websocket.Conn]bool),
		broadcast: make(chan *Message),
	}
}

func (s *WebSocketServer) HandleWebSocket(ctx *websocket.Conn){

	// REgister a new Client
	s.clients[ctx] = true
	defer func() {
		delete(s.clients, ctx)
		ctx.Close()
	}()


	for {
		_, msg, err := ctx.ReadMessage()
		if err != nil {
			log.Println("Read Error", err)
			break
		}


		// send the message to the broadcast channel
		var message Message
		if err := json.Unmarshal(msg, &message); err != nil{
			log.Fatalf("Error Umarshalling")
		}
		s.broadcast <- &message
	}
}

func (s *WebSocketServer) HandleMessages(){
	for { 
		msg := <- s.broadcast


		// Send the message to all Clients


		for client := range s.clients {
			err := client.WriteMessage(websocket.TextMessage, getMessageTemplate(msg))
			if err != nil {
				log.Printf("Write Error: %v", err)
				client.Close()
				delete(s.clients, client)
			}
		}
	}
}


func getMessageTemplate( msg *Message) []byte {
	tmpl, err := template.ParseFiles("views/message.html")
	if err != nil {
		log.Fatal("template parsing: %s" ,err)
	}

	// Render the template with the message as data
	var renderedMessage bytes.Buffer
	err = tmpl.Execute(&renderedMessage, msg)
	if err != nil {
		log.Fatalf("template execution: %s", err)
	}

	return renderedMessage.Bytes()
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

   // static route and directory
  app.Static("/static/", "./static") 


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



	// create new websocket
	server := NewWebSocket()


	app.Get("/ws", websocket.New(func(c *websocket.Conn) {
	    // Handle WebSocket connection here

	    server.HandleWebSocket(c)
	}))

	go server.HandleMessages()


   app.Get("/", func(c *fiber.Ctx) error {
       return indexHandler(c, db)
   })

   app.Get("/chatroom", func(c *fiber.Ctx) error {
	   return c.Render("chatroom", fiber.Map{})
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
