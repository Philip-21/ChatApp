//web socket implementation

var socket = new WebSocket("ws://localhost:8080/");

/*
connect() triggers a callback whenever it receives
 a new message from our WebSocket connection:
*/


//listen for webSocket connections
let connect = cb =>{
      console.log("Attempting Connection .....");

      socket.onopen = ()=>{
            console.log("Opened Successfully");
      };

      socket.onmessage = msg=>{
            console.log(msg);
            cb(msg);
      };

      socket.onclose = event => {
            console.log("Socket Closed Connection: ", event);
      };

      socket.onerror = error =>{
            console.log("Socket Error: ",error);
      };
};

//send messages from our frontend to our 
//backend via our WebSocket connection using socket.send()
let sendMsg = msg => {
      console.log("Sending Msg:", msg);
      socket.send(msg)
};

export {connect, sendMsg};
