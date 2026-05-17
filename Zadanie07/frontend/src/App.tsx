import { BrowserRouter, Routes, Route } from "react-router-dom";
import './App.css'
import ProductPage from './pages/ProductPage'
import PaymentPage from './pages/PaymentPage'

function App() {
  
  return (
    <>
      <h1>Welcome to Zadanie05</h1>
      <BrowserRouter>
        <Routes>
          <Route path='/' element={<ProductPage/>}/>
          <Route path="/payment" element={<PaymentPage/>} />
        </Routes>
      </BrowserRouter>
    </>
  )
}

export default App
