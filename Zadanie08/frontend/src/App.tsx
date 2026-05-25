import { Routes, Route } from "react-router-dom";
import { useState } from "react";

import Header from "./components/Header.tsx"
import Signin from "./components/Signin.tsx"
import Register from "./components/Register.tsx"

import "./App.css"

export default function App() {
  const [isLoggedIn, setIsLoggedIn] =
  useState(false);

  return (
      <Routes>
        <Route path="/" element={<Header isLoggedIn={isLoggedIn}/>} />
        <Route path="/signin" element={<Signin setIsLoggedIn={setIsLoggedIn}/>} />
        <Route path="/register" element={<Register setIsLoggedIn={setIsLoggedIn}/>}/>
      </Routes>
  );
}


