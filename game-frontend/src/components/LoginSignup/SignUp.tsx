import { JSX } from "react"; 
import "./loginsignup.css"; 

export default function Signup(): JSX.Element {
  return (
      <>
          <div className="form-container">
              <form action="" className="login-register-form">
                  <h1>Sign Up</h1>

                  <div className="input-label">
                      <label htmlFor="fullname">Full Name</label>
                      <input type="text" name="fullname" id="fullname" placeholder="Enter your full name" required />
                  </div>

                  <div className="input-label">
                      <label htmlFor="emailid">Email ID</label>
                      <input type="email" name="emailid" id="email" placeholder="Enter your email ID" required />
                  </div>

                  <div className="input-label">
                      <label htmlFor="pswd">Password</label>
                      <input type="password" name="pswd" id="password" placeholder="Enter your password" required />
                  </div>

                  <div className="input-label">
                      <label htmlFor="confirm-pswd">Confirm Password</label>
                      <input type="password" name="confirm-pswd" id="confirm-password" placeholder="Re-enter your password" required />
                  </div>

                  <div className="input-label">
                      <label htmlFor="terms">
                          <input type="checkbox" name="terms" id="terms" required /> I agree to the <a href="/">Terms & Conditions</a>
                      </label>
                  </div>

                  <div className="register">
                      <label htmlFor="login">Already have an account? <a href="/login">Login</a></label>
                  </div>

                  <div className="submit-btn">
                      <button className="submit-btn">Sign Up</button>
                  </div>

              </form>
          </div>
      </>
  );
}
