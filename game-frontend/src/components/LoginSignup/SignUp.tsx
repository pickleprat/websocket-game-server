import React, { JSX, useState } from "react"; 
import { supabase } from "../../SupabaseClient";

import "./loginsignup.css"; 

export default function Signup(): JSX.Element {
    const [username, setUserName] = useState<string>(''); 
    const [emailId, setEmailId] = useState<string>(''); 
    const [password, setPassword] = useState<string>(''); 

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault(); 
        const {data, error} = await supabase.auth.signUp({
            email: emailId, 
            password:  password, 
            options: {
                data: { username }, 
            }
        })

        if (error) {
            throw new Error("Error signing up user"); 
        } else {
            console.log(data)
        } 
    } 

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
                        <input type="text" onChange={(e) => {setUserName(e.target.value)}} 
                            name="fullname" id="fullname" placeholder="full name"  required />
                        <i className='bx bxs-user-account'></i>
                    </div>

                    <div className="input-box">
                        <label htmlFor="emailid"></label>
                        <input type="email" name="emailid" 
                            onChange={(e) => {setEmailId(e.target.value)}} id="email" 
                            placeholder="email id" required />
                        <i className='bx bxs-user-check' ></i>
                    </div>

                    <div className="input-box">
                        <label htmlFor="pswd"></label>
                        <input type="password" name="pswd" 
                            onChange={(e) => {setPassword(e.target.value)}} id="password" 
                            placeholder="password" required />
                        <i className='bx bx-lock'></i>
                    </div>

                    <div className="input-labels">
                        <label htmlFor="terms">
                            <input type="checkbox" name="terms" id="terms" required /> I agree to the <a href="/">Terms & Conditions</a>
                        </label>
                    </div>

                    <div className="submit-btn">
                        <button className="submit-btn">Sign Up</button>
                    </div>

                    <div className="register">
                        <label htmlFor="login">Already have an account? <a href="/login">Login</a></label>
                    </div>

                </div>
                </form>
            </div>
        </>
    );
}
