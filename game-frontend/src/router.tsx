import { createBrowserRouter } from "react-router-dom";
import { RouteObject } from "react-router-dom"; 

import App from "./App"; 
import Login from "./components/LoginSignup/Login"; 
import SignUp from "./components/LoginSignup/SignUp"; 
import Rooms from "./components/Rooms/Rooms"; 
import LiveRoom from "./components/Rooms/LiveRooms"

export const router = createBrowserRouter([
    { path: "/", element: <App/> } as RouteObject, 
    { path: "/login", element: <Login/> } as RouteObject, 
    { path: "/signup", element: <SignUp/> } as RouteObject, 
    { path: "/rooms", element: <Rooms /> } as RouteObject, 
    { path: "/live-room/:roomId", element: <LiveRoom/> } as RouteObject, 
]);  
