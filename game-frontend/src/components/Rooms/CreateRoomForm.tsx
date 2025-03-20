import { JSX, useState } from "react";
import { useAuthContext } from "../../AuthContext";

import "./room.css"

export default function CreateRoomForm(): JSX.Element {
    const [roomName, setRoomName] = useState<string | null>(null); 
    const [roomDescription, setRoomDescription] = useState<string | null>(null); 
    const [roomGenre, setRoomGenre] = useState<string | null>(null); 
    const authSesh = useAuthContext(); 

    const handleCreateRoom = () => {
        const jwtToken = authSesh?.userSession?.session.access_token; 
        const goServerUrl = import.meta.env.VITE_GO_SERVER_URL; 

    } 

    return (
        <div className="room-creation-container">
            <form className="room-creation-form" onSubmit={handleCreateRoom}>
                <div className="room-input">
                    <input 
                        type="text" 
                        className="room-text" 
                        name="room-title" 
                        id="room-title" 
                        placeholder="Give your room a name"
                        onChange={(e) => {setRoomName(e.target.value)}}
                        required
                    />
                </div>
                <div className="room-input">
                    <textarea
                        name="room-description" 
                        className="room-text" 
                        id="room-description" 
                        placeholder="Describe your room for other users..."
                        onChange={(e)=>{setRoomDescription(e.target.value)}}
                        required
                    />
                </div>
                <div className="room-input">
                    <input 
                        type="text"     
                        className="room-text"   
                        name="room-genre"   
                        id="room-genre"
                        placeholder="Which genre would you like your room to fall under?"
                        onChange={(e) => {setRoomGenre(e.target.value)}}
                        required
                    />
                </div>
                <div className="create-room-btn-shell">
                    <button className="create-room-btn">Create</button>
                </div>
            </form>
        </div>
    ) 
} 