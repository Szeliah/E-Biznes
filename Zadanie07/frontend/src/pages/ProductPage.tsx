import { useEffect, useState } from "react";
import { productService } from "../services/productService";
import ProductCard from "../components/ProductCard";
import "./ProductPage.css"

type Product = {
    id: number,
    name: string,
    price: number,
    description: string
};

export default function ProductPage() {
    const [products, setProducts] = useState<Product[]>([]);
    const [loading, setLoading] = useState<boolean>(true); 
    const [error, setError] =  useState<string | null>(null);

    
    useEffect(() => {
        
        const fetchData = async () => {
            try {
                const data = await productService.getProducts();
                setProducts(data);
            } catch(err) {
                console.error(err);
                setError("Failed to fetch products")
            } finally {
                setLoading(false);
            }
        };

        fetchData();

    }, []);

    if (loading) {
        return <p> Loading </p>;
    }

    if (error) {
        return <p style={{ color: "red" }}>{error}</p>
    }

    return (
        <>
            <h2 className="title">Nasze wszystkie produkty</h2>
            <div className="container">
                {products.map((p) => (
                    <ProductCard key={p.id} product={p} />
                ))}
            </div>
        </>
    )
}