import { JSX } from "react"; 
import "./loginsignup.css"; 

export default function Signup(): JSX.Element {
  return (
      <>
        <div className="game-title">
            <h1>Tatak.ai</h1>
        </div>
        <div className="form-container">
            <form action="" className="login-register-form">
            <div className="register-container">
                <h1>Sign Up</h1>

                <div className="input-box">
                    <label htmlFor="fullname"></label>
                    <input type="text" name="fullname" id="fullname" placeholder="Enter your full name" required />
                    <i className='bx bxs-user-account'></i>
                </div>

                <div className="input-box">
                    <label htmlFor="emailid"></label>
                    <input type="email" name="emailid" id="email" placeholder="Enter your email ID" required />
                    <i className='bx bxs-user-check' ></i>
                </div>

                <div className="input-box">
                    <label htmlFor="pswd"></label>
                    <input type="password" name="pswd" id="password" placeholder="Enter your password" required />
                    <i className='bx bx-lock'></i>
                </div>

                <div className="input-box">
                    <label htmlFor="confirm-pswd"></label>
                    <input type="password" name="confirm-pswd" id="confirm-password" placeholder="Re-enter your password" required />
                    <i className='bx bx-target-lock' ></i>
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
