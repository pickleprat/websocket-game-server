import { createContext, useContext } from "react";
import { Session } from "@supabase/supabase-js";
import { redirect } from "react-router-dom";

export const AuthContext = createContext<Session | undefined >(undefined); 

export function useAuthContext() : Session | undefined  {
    const session = useContext(AuthContext); 
    if (session === undefined) {
        redirect("/"); 
    } 
    return session; 
} 