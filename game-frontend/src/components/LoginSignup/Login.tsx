import { JSX } from "react"; 
import { useState } from "react";
import { useNavigate } from "react-router-dom";
import { supabase } from "../../SupabaseClient";

import "./loginsignup.css"; 

export default function Login(): JSX.Element {
    const [email, setEmail] = useState<string>(''); 
    const [password, setPassword] = useState<string>(''); 
    const [showPassword, setShowPassword] = useState<boolean>(false);
    const navigate = useNavigate(); 

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault(); 

        const { data, error } = await supabase.auth.signInWithPassword({
            email,
            password
        });

        if (error) {
            alert("Invalid email or password. Please try again.");
            console.error("Login error:", error.message);
        } else {
            if(data.session) {
                console.log("user session created"); 
            } 

            alert("Login successful!");
            navigate("/"); 
        }
    };

    return (
        <>
            <div className="game-title">
                <h1>Tatak.ai</h1>
            </div>
            <div className="form-container">
                <form onSubmit={handleSubmit} className="login-register-form">
                    <div className="login-container">
                        <h1>Login</h1>

                        {/* EMAIL FIELD */}
                        <div className="input-box">
                            <label htmlFor="emailid" />
                            <input 
                                type="email" 
                                name="emailid" 
                                id="email" 
                                placeholder="username" 
                                onChange={(e) => setEmail(e.target.value)} 
                                required
                            />
                            <i className='bx bx-user-circle'></i>
                        </div>

                        <div className="input-box password-box">
                            <label htmlFor="pswd" />
                            <input 
                                type={showPassword ? "text" : "password"} 
                                name="pswd" 
                                id="password" 
                                placeholder="password"
                                onChange={(e) => setPassword(e.target.value)} 
                                required
                            />
                            <i 
                                className={showPassword ? 'bx bxs-lock-alt' : 'bx bx-low-vision'}
                                onClick={() => setShowPassword(!showPassword)}
                                style={{ cursor: 'pointer' }} 
                                aria-label="Toggle password visibility"
                            ></i>
                        </div>

                        <div className="input-labels">
                            <label htmlFor="remember-me"> 
                                <input type="radio" name="remember-me" id="remember-me" /> Remember Me 
                            </label>    
                            <label htmlFor="forgot-password">
                                <a href="/">Forgot Password?</a>
                            </label>
                        </div>

                        <div className="submit-btn">
                            <button type="submit" className="submit-btn">Login</button>
                        </div>

                        <div className="register">
                            <label htmlFor="register">Don't have an account? <a href="/signup">Signup</a></label>
                        </div>
                    </div>
                </form>
            </div> 
        </>
    );
}
