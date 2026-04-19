import { useState, useEffect } from 'react'
import './App.css'
import ProductPage from './pages/ProductPage'
import PaymentPage from './pages/PaymentPage'

function App() {
  const [count, setCount] = useState(0)

  return (
    <>
    <h1>Welcome to Zadanie05</h1>
    <ProductPage />
    <PaymentPage />
    </>
  )
}

export default App
