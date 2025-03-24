import { JSX, useEffect } from "react";

export default function ShowMyRooms() : JSX.Element {
    useEffect(() => {
        console.log("Use effect ran")
    }, [])

    return (
        <>
        <div className="room-container">

        </div>
        </>
    )
} 