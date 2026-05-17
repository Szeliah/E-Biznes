const API_URL = "http://localhost:8080/products"


export const productService = {

    async getProducts() {
        const resp = await fetch(API_URL);

        if (!resp.ok) {
            const text = await resp.text();
            throw new Error(text || "Failed to fetch products") 
        }

        return await resp.json();
    }

}
