'use client'

import { Product } from "@/app/types/product";
import { numberMoveDecimal } from "@/app/utils/Utils";
import { Button } from "@chakra-ui/react";
import Image from "next/image";

interface ProductCardProps {
    product: Product;
    onAddToCard: () => void;
}

const ProductCard: React.FC<ProductCardProps> = ({
    product,
    onAddToCard
}) => {
    return (
        <div
        key={product.id}
        className="bg-white rounded-lg shadow-md overflow-hidden hover:shadow-lg transition-shadow duration-300"
        >
        <Image src={'/images/products/no_product.png'} alt={product.name} width={38} height={38} className="w-full h-48 object-cover" />
        <div className="p-4">
            <h3 className="text-lg font-semibold mb-2">{product.name}</h3>
            <p className="text-gray-600 mb-2">{product.description}</p>
            <div className="flex items-center justify-between">
            <span className="text-lg font-bold">
                ${numberMoveDecimal(product.price, 2)}
            </span>
            <Button
                variant="main"
                size="sm"
                onClick={() => onAddToCard()}
            >
                + Agregar
            </Button>
            </div>
        </div>
        </div>
    );
}

export default ProductCard;