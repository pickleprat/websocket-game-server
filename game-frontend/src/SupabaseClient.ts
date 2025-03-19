import { createClient } from "@supabase/supabase-js";

// constants
const SUPABASE_URL : string = import.meta.env.VITE_SUPABASE_API_URL as string; 
const SUPABASE_KEY : string = import.meta.env.VITE_SUPABASE_ANON_KEY as string; 

export const supabase = createClient(SUPABASE_URL, SUPABASE_KEY); 