"use client";

import { Product } from "@/app/types/product";
import { numberMoveDecimal } from "@/app/utils/Utils";
import { Button } from "@chakra-ui/react";
import { Icon } from "@iconify/react";
import Image from "next/image";
import { useState } from "react";

interface ProductCardProps {
  product: Product;
  onAddToCard: () => void;
  onImageClick?: () => void;
}

const ProductCard: React.FC<ProductCardProps> = ({ product, onAddToCard, onImageClick }) => {
  const [currentImageIndex, setCurrentImageIndex] = useState<number>(0);

  const nextImage = () => {
    setCurrentImageIndex(
      (prevIndex) => (prevIndex + 1) % (product.images?.length || 1)
    );
  };

  const prevImage = () => {
    setCurrentImageIndex((prevIndex) =>
      prevIndex === 0 ? (product.images?.length || 1) - 1 : prevIndex - 1
    );
  };

  return (
    <div
      key={product.id}
      className="bg-white rounded-lg shadow-md overflow-hidden hover:shadow-lg transition-shadow duration-300"
    >
       <div 
         className={`relative w-full h-40 lg:h-64 overflow-hidden ${onImageClick ? 'cursor-pointer' : ''}`}
         onClick={onImageClick}
       >
        <Image
          src={
            product.images && product.images.length > 0
              ? product.images[currentImageIndex]
              : "/images/products/no_product.png"
          }
          alt={product.name}
          layout="fill"
          objectFit="contain"
          className="rounded-t-lg select-none object-center scale-100 hover:scale-105 transition-transform duration-300"
          draggable="false"
        />
        {product.images && product.images.length > 1 && (
          <>
            <button
              onClick={(e) => { e.stopPropagation(); prevImage(); }}
              className="absolute left-2 top-1/2 transform -translate-y-1/2 bg-white bg-opacity-50 rounded-full p-2 hover:bg-opacity-75 transition-all duration-200"
              aria-label="Previous image"
            >
              <Icon icon="material-symbols:chevron-left" />
            </button>
            <button
              onClick={(e) => { e.stopPropagation(); nextImage(); }}
              className="absolute right-2 top-1/2 transform -translate-y-1/2 bg-white bg-opacity-50 rounded-full p-2 hover:bg-opacity-75 transition-all duration-200"
              aria-label="Next image"
            >
              <Icon icon="material-symbols:chevron-right" />
            </button>
          </>
        )}
      </div>
      <div className="p-4">
        <h3 className="text-lg font-semibold mb-2">{product.name}</h3>
        <p className="text-gray-600 mb-2">{product.description}</p>
        <div className="flex items-start justify-between flex-col md:flex-row md:items-center">
          <span className="text-lg font-bold">
            ${numberMoveDecimal(product.price, 2)}
          </span>
          <Button variant="default" size="sm" onClick={() => onAddToCard()}>
            + Agregar
          </Button>
        </div>
      </div>
    </div>
  );
};

export default ProductCard;
