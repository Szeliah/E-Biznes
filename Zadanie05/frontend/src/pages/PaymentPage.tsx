import { useEffect, useState } from "react";
import { paymentService } from "../services/paymentService";

type Payment = {
    cart_id: number,
    amount: number,
    status: string,
    method: string
};

export default function PaymentPage() {
  const [payment, setPayment] = useState<Payment>({
      cart_id: 1,
      amount: 221.97,
      status: "pending",
      method: "card"
  });
  const [status, setStatus] = useState<"idle" | "success" | "error">("idle");

  const handleSubmit = async () => {
    try {
      const resp = await paymentService.sendPaymentData(payment)
     

      if (!resp.ok) {
        throw new Error("Payment failed");
      }

      setStatus("success");
    } catch (err) {
      console.error(err);
      setStatus("error");
    }
  };


  return (
    <div>
      <h1>Płatność</h1>


      <p>Do zapłaty: {payment.amount}</p>

      <input
        placeholder="Metoda"
        value={payment.method}
        onChange={(e) =>
          setPayment({ ...payment, method: e.target.value })
        }
      />

      <button onClick={handleSubmit}>Zapłać</button>

      {status === "success" && (
        <p style={{ color: "green" }}>Płatność udała się</p>
      )}

      {status === "error" && (
        <p style={{ color: "red" }}> Błąd płatności</p>
      )}
    </div>
  );
}