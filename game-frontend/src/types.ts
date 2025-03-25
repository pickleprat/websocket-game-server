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

interface Member {
  id: string;
  full_name: string;
  created_at: string;
}

export interface Room {
    id: string;
    full_name: string;
    created_at: string;
    "owner-name": string;
    name: string;
    genre: string;
    description: string;
    members: Member[];
    owner: string;
} 
 


