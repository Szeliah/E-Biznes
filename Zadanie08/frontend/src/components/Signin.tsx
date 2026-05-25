import React, { useState, type SyntheticEvent } from "react";
import { useNavigate } from "react-router-dom";

import "./Signin.css"

type SigninProps = {
    setIsLoggedIn: React.Dispatch<
        React.SetStateAction<boolean>
    >;
};

export default function Signup({ setIsLoggedIn }: SigninProps) {
    const [email, setEmail] = useState<string>("");
    const [password, setPassword] = useState<string>("");

    const [error, setError] = useState<string>("");
    const [success, setSuccess] = useState<string>("");

    const navigate = useNavigate();

    async function handleSubmit(e: SyntheticEvent) {
        e.preventDefault();

        if (!email.includes("@")) {
            setError("Email is invalid");
            return;
        }

        if (password.length < 6) {
            setError("Password must have at least 6 characters");
            return;
        }

        setError("");

        try {
            const res = await fetch(
                "http://localhost:3000/login",
                {
                    method: "POST",
                    credentials: "include",
                    headers: {
                        "Content-Type": "application/json",
                    },
                    body: JSON.stringify({
                        email,
                        password,
                    }),
                }
            );

            if (!res.ok) {
                setError("Invalid credentials");
                return;
            }

            const data = await res.json();
            console.log(data);

            setIsLoggedIn(true);
            setSuccess("Successfully logged in!");
            setError("");



            setTimeout(() => {
                navigate("/");
            }, 1500);
        } catch (err) {
            setError("Something went wrong");
        }

    } 

    return (
        <div className="signup-card">
            <h2>Sign in</h2>

            <form className="signup-form" onSubmit={handleSubmit}>
                <input type="email" placeholder="Email" value={email}
                    onChange={(e) => {
                    setEmail(e.target.value);
                    setError("");
                    setSuccess("");
                }}/>
                <input type="password" placeholder="Password" value={password} onChange={(e) => {
                    setPassword(e.target.value);    
                    setError("");
                    setSuccess("");
                }}/>
                {error && <p className="error">{error}</p>}
                {success && <p className="success">{success}</p>}
                <button type="submit">
                    login
                </button>
            </form>
            <p className="auth-switch">
                Don't have an account?{" "}

                <span onClick={() => navigate("/register")}>
                    Sign up
                </span>
            </p>
        </div>
    );
}