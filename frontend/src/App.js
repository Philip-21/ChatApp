/*
src/App. js : This is the file for App Component. 
App Component is the main component in React which acts as a container for all other components
App. js is built to serve makers of static single-page apps. 
This means that it keeps all page navigation within the session of the webpage, 
defining "pages" as DOM nodes that can be instantiated
*/
import React, {Component} from "react";
import "./App.css";
import {connect, sendMsg} from "./api";
import Header from "./components/Header";





class App extends Component {
  
     //a connection that triggers the sendMsg()
     //when the button is clicked
    constructor(props){
        super(props);// invoke a super class constructor
        connect();
    }
    send() {
        console.log("hello")
        sendMsg("hello");
    }

    //defining chatChannel state using a constructor 
    constructor(props){
        super(props);
        this.state={
            ChatChannel : []
        }
    }

    /*
    move our connect() call from the constructor into a componentDidMount() 
    function which will be called automatically as part of our Components life-cycle.
    */

    componentDidMount(){
        connect((msg)=>{
            console.log("New Message")
            this.setState(prevState=>({
                ChatChannel: [...this.state.ChatChannel, msg]
            }))
            console.log(this.state)
        });
    }

    render() {
        return (
            <div className="App">
                <buttton onClick={this.send}>Hit</buttton>
            </div>
        );
    }
}

export default App; 