import React, { useState, type SyntheticEvent } from "react";
import { useNavigate } from "react-router-dom";

import "./Register.css"
const API_URL = import.meta.env.VITE_API_URL;

type RegisterProps = {
    setIsLoggedIn: React.Dispatch<
    React.SetStateAction<boolean>
    >;
    setUserEmail: React.Dispatch<
    React.SetStateAction<string>
    >;
};


export default function Signup({ setIsLoggedIn, setUserEmail }: RegisterProps) {
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
                `${API_URL}/register`,
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
                const errorData = await res.json();
                setError(errorData.message);

                return;
            }

            console.log(res.status);

            const data = await res.json();
            console.log(data);

            setIsLoggedIn(true);
            setSuccess("Account created!");
            setUserEmail(email);
            setError("");

            setTimeout(() => {
                navigate("/");
            }, 1500);
        } catch (err) {
            console.log(err);
            setError("Something went wrong");
        }

    } 

    return (
        <div className="signup-card">
            <h2>Register</h2>

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
                    register
                </button>
            </form>
            <p className="auth-switch">
                Already have an account?{" "}

                <span onClick={() => navigate("/signin")}>
                    Sign in
                </span>
            </p>
        </div>
    );
}