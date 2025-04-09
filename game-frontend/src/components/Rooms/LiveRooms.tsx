import { JSX  } from "react";
import { useParams } from "react-router-dom";
import { useAuthContext } from "../../AuthContext";
import useWebSocket from "react-use-websocket";

export default function LiveRoom(): JSX.Element {
    const { roomId } = useParams(); 
    const authSesh = useAuthContext(); 
    const {sendJsonMessage, lastJsonMessage} = useWebSocket("ws://localhost:8000/api/connectRoom"); 
    console.log(sendJsonMessage, lastJsonMessage); 

    return (
        <>
        </>
    )
} 