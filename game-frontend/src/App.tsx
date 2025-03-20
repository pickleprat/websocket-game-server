import { JSX } from "react";
import "./App.css"

export default function App() : JSX.Element {
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
