/*
  
*/
import React, {Component} from "react";
import "./ChatChannel.scss";


class ChatChannel extends Component {
      //the render() returns the jsx we wish to render in
      //our app for this particular component
      render () {
            const messages= this.props.ChatChannel.map((msg, index)=>(
                  <p key={index}>{msg.data}</p>
            ));
            return (
                  <div className="ChatChannel">
                        <h2>Chat Channel</h2>
                        {messages}
                  </div>
            );
      }
}
export default ChatChannel;