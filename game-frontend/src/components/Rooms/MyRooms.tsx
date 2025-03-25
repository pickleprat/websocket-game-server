import { JSX, useEffect  } from "react";
import { useState } from "react"; 
import { useAuthContext } from "../../AuthContext";
import { useNavigate } from "react-router-dom";
import { Room } from "../../types";

export default function ShowMyRooms() : JSX.Element {
    const [myRooms, setMyRooms] = useState<Room[]>([]); 
    const navigate = useNavigate(); 
    const goServerUrl = import.meta.env.VITE_GO_SERVER_URL; 
    const authSesh = useAuthContext();
    console.log("Showing yoour rooms")

    useEffect(() => {
        console.log("use effect ran")
        if (authSesh?.userSession && authSesh.userSession?.session) {
            console.log("user is auth'ed")
            const userId = authSesh.userSession.user.id; 
            const jwtToken = authSesh.userSession.session.access_token; 
            const fetchMyRooms = async() => {
                try {
                    console.log("fetching rooms...")
                    const response = await fetch(goServerUrl + "/getMyRooms", {
                        method: "POST", 
                        body: JSON.stringify({
                            id: userId, 
                        }), 
                        headers: {
                            "Content-Type": "application/json",
                            "Authorization": `Bearer ${jwtToken}`
                        }, 
                    })

                    const roomsFetched = await response.json() as Room[]; 
                    setMyRooms(roomsFetched)
                    console.log(roomsFetched)
                }catch (error) {
                    console.log("err fetching rooms")
                } 
            } 
            fetchMyRooms(); 
        } else {
            navigate("/") 
        } 

    }, [])

    return (
        <>
        <div className="room-container">
            {myRooms.map((room) => (
                <div key={room.id} className="room-card">
                <h2>{room.name}</h2>
                <p>{room.description}</p>
                <p><strong>Genre:</strong> {room.genre}</p>
                </div>
            ))}
        </div>
        </>
    )
} 