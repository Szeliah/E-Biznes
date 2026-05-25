import { Routes, Route } from "react-router-dom";
import { useState, useEffect } from "react";

import Header from "./components/Header.tsx";
import Signin from "./components/Signin.tsx";
import Register from "./components/Register.tsx";

import "./App.css";

export default function App() {
  const [isLoggedIn, setIsLoggedIn] =  useState(false);
  const [loading, setLoading] = useState(true);
  const [userEmail, setUserEmail] =useState("");

  useEffect(() => {

    async function checkAuth() {

        try {

          const res = await fetch(
              "http://localhost:3000/me",
              {
                  credentials: "include",
              }
          );

          if (res.ok) {
            const data = await res.json();
            setIsLoggedIn(true);
            setUserEmail(data.email);

          } else {
            setIsLoggedIn(false);
            setUserEmail("");
          }

        } catch {
          setIsLoggedIn(false);
        }

        setLoading(false);
    }

    checkAuth();

  }, []);

  if (loading) {
      return <p>Loading...</p>;
  }

  return (
      <Routes>
          <Route path="/" element={ <Header isLoggedIn={ isLoggedIn } setIsLoggedIn={ setIsLoggedIn } userEmail={ userEmail }/>}/>
          <Route path="/signin" element={ <Signin setIsLoggedIn={ setIsLoggedIn } setUserEmail={setUserEmail}/> }/>
          <Route path="/register" element={ <Register setIsLoggedIn={ setIsLoggedIn } setUserEmail={setUserEmail}/> }/>
      </Routes>
  );
}