import { createContext, JSX, useContext } from "react";
import { useNavigate } from "react-router-dom";
import { UserSession } from "./types";
import { ReactNode } from "react";
import { useState } from "react";

interface AuthSession {
    userSession: UserSession | null;  
    setUserSession: (session: UserSession | null) => void; 
} 

export const AuthContext = createContext<AuthSession| undefined >(undefined); 

export function AuthContextProvider({ children }: {children : ReactNode }) : JSX.Element {
    const [userSession, setUserSession] = useState<UserSession | null>(null); 

    return (
        <AuthContext.Provider value={{ userSession, setUserSession }}>
          {children}
        </AuthContext.Provider>
      );
} 

export function useAuthContext() : AuthSession | undefined  {
    const navigate = useNavigate(); 
    const userSessionContext = useContext(AuthContext); 
    if (userSessionContext === undefined) {
        alert("Session doesn't exist")
        navigate("/"); 
    } 
    return userSessionContext; 
} 