import { Session } from "@supabase/supabase-js";
import { User } from "@supabase/supabase-js"


export interface UserSession {
    session: Session 
    user: User
} 

export interface CreateRoomResponse {
    roomName: string 
    createdAt: string 
    roomId: string 
    roomStatus: boolean
} 
 


