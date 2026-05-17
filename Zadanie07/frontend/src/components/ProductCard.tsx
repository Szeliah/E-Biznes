type Product = {
    id: number,
    name: string,
    price: number,
    description: string
};


export default function ProductCard({ product }: { product: Product }){
    return (
        <div className="product-item">
            <h3>{product.name}</h3>
            <p>{product.description}</p>
            <p>Cena: {product.price} zł</p>
        </div>

    );
}