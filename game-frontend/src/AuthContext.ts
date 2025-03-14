import { createContext, useContext } from "react";
import { Session } from "@supabase/supabase-js";

export const AuthContext = createContext<Session | undefined >(undefined); 

export function useAuthContext() : Session {
    const session = useContext(AuthContext); 
    if (session === undefined) {
        throw new Error("User is not logged in"); 
    } 
    return session; 
} 