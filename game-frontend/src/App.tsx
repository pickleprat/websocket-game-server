// import { useState } from 'react'
// import { createClient, Session } from '@supabase/supabase-js'

import { JSX } from "react";
import "./App.css"
import { useAuthContext } from "./AuthContext";

export default function App() : JSX.Element {
  // const [session, setSession] = useState<Session | null> (null); 
  const authSesh = useAuthContext(); 
  console.log("The auth session data : = ", authSesh) 
  return (
    <>
      <div className="game-title">
          <h1>Tatak.ai</h1>
      </div>
      <a href="/login">Login</a>
      <a href="signup"> Signup</a>
    </>
  )
}
