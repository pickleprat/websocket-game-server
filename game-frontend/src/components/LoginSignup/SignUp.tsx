import React, { JSX, useState } from "react"; 
import { supabase } from "../../SupabaseClient";
import { useNavigate } from "react-router-dom";
import { useAuthContext } from "../../AuthContext";
import { UserSession } from "../../types";

import "./loginsignup.css"; 

export default function Signup(): JSX.Element {
    const [username, setUserName] = useState<string>(''); 
    const [emailId, setEmailId] = useState<string>(''); 
    const [password, setPassword] = useState<string>(''); 
    const [showPassword, setShowPassword] = useState<boolean>(false);
    const authSesh = useAuthContext(); 
    const navigate = useNavigate(); 

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault(); 
        const { data, error } = await supabase.auth.signUp({
            email: emailId, 
            password: password, 
            options: {
                data: { username }, 
            }
        });

        if (error) {
            if (error.message.includes("User already registered")) {
                alert("User already exists. Please go to the login page.");
            } else {
                alert("Some error signing up the user. It isn't that your email exists though");  
            }
        } else {
            const user = data.user; 
            if (user) {
                authSesh?.setUserSession({
                    session: data.session, 
                    user: user, 
                } as UserSession)

                const {data: insertData, error: insertError } = await supabase
                .from("profiles") 
                .insert([{
                    id: user.id, 
                    full_name: username, 
                }]); 

                if (insertError) {
                    console.error("Error inserting user profile:", insertError);
                    alert("Profile creation failed.");
                } else if(insertData) {
                    navigate("/login");
                } 
            } 
            navigate("/login") 
        } 
    };

    return (
        <>
            <div className="game-title">
                <h1>Tatak.ai</h1>
            </div>
            <div className="form-container">
                <form onSubmit={handleSubmit} className="login-register-form">
                    <div className="register-container">
                        <h1>Sign Up</h1>

                        <div className="input-box">
                            <label htmlFor="fullname"></label>
                            <input 
                                type="text" 
                                onChange={(e) => setUserName(e.target.value)} 
                                name="fullname" 
                                id="fullname" 
                                placeholder="full name" 
                                required 
                            />
                            <i className='bx bxs-user-account'></i>
                        </div>

                        <div className="input-box">
                            <label htmlFor="emailid"></label>
                            <input 
                                type="email" 
                                name="emailid" 
                                onChange={(e) => setEmailId(e.target.value)} 
                                id="email" 
                                placeholder="email id" 
                                required 
                            />
                            <i className='bx bxs-user-check'></i>
                        </div>

                        <div className="input-box">
                            <label htmlFor="pswd"></label>
                            <input 
                                type={showPassword ? "text" : "password"} 
                                name="pswd" 
                                onChange={(e) => setPassword(e.target.value)} 
                                id="password" 
                                placeholder="password" 
                                required 
                            />
                            <i 
                                className={showPassword ? 'bx bxs-lock-alt' : 'bx bx-low-vision'} 
                                onClick={() => setShowPassword(!showPassword)}
                                style={{ cursor: 'pointer' }}
                                aria-label="Toggle password visibility"
                            />
                        </div>

                        <div className="input-labels">
                            <label htmlFor="terms">
                                <input type="checkbox" name="terms" id="terms" required /> 
                                I agree to the <a href="/">Terms & Conditions</a>
                            </label>
                        </div>

                        <div className="submit-btn">
                            <button className="submit-btn">Sign Up</button>
                        </div>

                        <div className="register">
                            <label htmlFor="login">
                                Already have an account? <a href="/login">Login</a>
                            </label>
                        </div>
                    </div>
                </form>
            </div>
        </>
    );
}
