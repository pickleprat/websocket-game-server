import { JSX  } from "react";
import { useParams } from "react-router-dom";

export default function LiveRoom(): JSX.Element {
    const { roomId } = useParams(); 
    console.log(roomId); 
    return (
        <>
        </>
    )
} 