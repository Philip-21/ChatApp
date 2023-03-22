package websockets

import "chatapp/internal/config"

//a Pool struct which will contain all of the channels we need for
//concurrent communication, as well as a map of clients
type Pool struct {
	// registers a new user and will send out New User Joined
	//to all of the clients within this pool when a new client connects
	Register chan *Client
	//unregister a user and notify the pool when a client disconnects.
	Unregister chan *Client
	// a map of clients to a boolean value. We can use the boolean value to dictate active/inactive
	//but not disconnected further down the line based on browser focus.
	Clients map[*Client]bool
	//a channel which when it is passed, a message will loop through all clients
	// in the pool and send the message through the socket connection
	Broadcast chan Message
	App       *config.AppConfig
}

func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			pool.Clients[client] = true
			pool.App.InfoLog.Println("Size of Connection Pool:", len(pool.Clients))

			for client, _ := range pool.Clients {
				pool.App.InfoLog.Println(client)
				client.Conn.WriteJSON(
					Message{Type: 1, Body: "New User Joined..."})
			}
			break
		case client := <-pool.Unregister:
			delete(pool.Clients, client)
			pool.App.InfoLog.Println("Size of Connection Pool:", len(pool.Clients))
			for client, _ := range pool.Clients {
				client.Conn.WriteJSON(
					Message{Type: 1, Body: "User Disconnected"})
			}
			break
		case message := <-pool.Broadcast:
			pool.App.InfoLog.Println("Sending Message to all clients in the pool")
			for client, _ := range pool.Clients {
				err := client.Conn.WriteJSON((message))
				if err != nil {
					pool.App.ErrorLog.Println(err)
				}
			}
		}
	}
}
