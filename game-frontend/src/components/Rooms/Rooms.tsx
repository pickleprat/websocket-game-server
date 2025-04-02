import { JSX } from "react";
import CreateRoomForm from "./CreateRoomForm";
import ShowLiveRooms from "./LiveRooms";
// import ShowMyRooms from "./MyRooms";

export default function Rooms() : JSX.Element {
    return (
        <>
            <CreateRoomForm /> 
            {/* <ShowMyRooms /> */}
            <ShowLiveRooms /> 
        </>
    ) 
} 