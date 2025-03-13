import { JSX } from "react"; 
import "./loginsignup.css"; 

export default function Login() : JSX.Element {
    return (
        <>
        <div className="form-container">
            <form action="" className="login-register-form">
                <div className="login-container">
                    <h1>Login</h1>
                    <div className="input-box">
                        <label htmlFor="emailid" />
                        <input type="email" name="emailid" id="email" placeholder="username" required/>
                        <i className='bx bx-user-circle'></i>
                    </div>
                    <div className="input-box">
                        <label htmlFor="pswd" />
                        <input type="password" name="pswd" id="password" placeholder="password" required/>
                        <i className='bx bxs-lock-alt'></i>
                    </div>
                    <div className="input-labels">
                        <label htmlFor="remember-me"> 
                            <input type="radio" name="remember-me" id="remember-me" /> Remember Me </label>    
                        <label htmlFor="forgot-password">
                            <a href="/">Forgot Password?</a>
                        </label>
                    </div>
                    <div className="submit-btn">
                        <button className="submit-btn">Login</button>
                    </div>
                </div>
                    <div className="register">
                        <label htmlFor="register">Don't have an account? <a href="/signup">Signup</a></label>
                    </div>
                </form>
            </div> 
        </>
    )
}  