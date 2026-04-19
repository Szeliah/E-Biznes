const API_URL = "http://localhost:8080/payment"

type Payment = {
    cart_id: number,
    amount: number,
    status: string,
    method: string
};


export const paymentService = {

    async sendPaymentData(payment: Payment) {
        const resp = await fetch(API_URL, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify(payment),
        });

        if (!resp.ok) {
            throw new Error("Payment failed");
        } 

        return resp
    }

}
