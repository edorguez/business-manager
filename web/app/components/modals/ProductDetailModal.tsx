"use client";

import useProductDetailModal from "@/app/hooks/useProductDetailModal";
import { numberMoveDecimal } from "@/app/utils/Utils";
import {
  Box,
  Modal,
  ModalBody,
  ModalContent,
  ModalOverlay,
  ModalHeader,
  ModalCloseButton,
} from "@chakra-ui/react";
import { Icon } from "@iconify/react";
import Image from "next/image";
import { useState, useEffect } from "react";
import { PhotoProvider, PhotoView } from "react-photo-view";
import "react-photo-view/dist/react-photo-view.css";

const ProductDetailModal = () => {
  const productDetailModal = useProductDetailModal();
  const [currentImageIndex, setCurrentImageIndex] = useState<number>(0);

  const product = productDetailModal.product;
  const images = product?.images || [];

  useEffect(() => {
    setCurrentImageIndex(0);
  }, [product]);

  useEffect(() => {
    if (currentImageIndex >= images.length && images.length > 0) {
      setCurrentImageIndex(0);
    }
  }, [images, currentImageIndex]);

  const nextImage = () => {
    setCurrentImageIndex((prevIndex) => (prevIndex + 1) % (images.length || 1));
  };

  const prevImage = () => {
    setCurrentImageIndex((prevIndex) =>
      prevIndex === 0 ? (images.length || 1) - 1 : prevIndex - 1
    );
  };

  if (!product) return null;

  return (
    <Modal
      isOpen={productDetailModal.isOpen}
      onClose={productDetailModal.onClose}
      size={{ base: "full", md: "lg" }}
    >
      <ModalOverlay />
      <ModalContent className="max-h-[90vh] overflow-hidden">
        <ModalHeader className="border-b">
          <div className="flex items-center justify-between">
            <h2 className="text-xl font-bold">{product.name}</h2>
            <ModalCloseButton />
          </div>
        </ModalHeader>
        <ModalBody className="p-0 overflow-y-auto">
          <div className="flex flex-col">
            {images.length > 0 ? (
              <PhotoProvider>
                {/* Hidden gallery images - non-current images for gallery navigation */}
                {images.length > 1 && (
                  <div style={{ visibility: 'hidden', position: 'absolute', width: 0, height: 0, overflow: 'hidden' }}>
                    {images
                      .filter((_, idx) => idx !== currentImageIndex)
                      .map((image, index) => (
                        <PhotoView key={index} src={image}>
                          <div />
                        </PhotoView>
                      ))}
                  </div>
                )}

                {/* Main Image Area */}
                <div className="relative w-full h-64 md:h-80 bg-gray-100">
                  <PhotoView src={images[currentImageIndex]}>
                    <div className="relative w-full h-full cursor-zoom-in">
                      <Image
                        src={images[currentImageIndex]}
                        alt={product.name}
                        layout="fill"
                        objectFit="contain"
                        className="select-none"
                        draggable="false"
                      />
                    </div>
                  </PhotoView>

                  {images.length > 1 && (
                    <>
                      <button
                        onClick={prevImage}
                        className="absolute left-2 top-1/2 transform -translate-y-1/2 bg-white bg-opacity-75 rounded-full p-2 hover:bg-opacity-100 transition-all duration-200 shadow-md"
                        aria-label="Previous image"
                      >
                        <Icon icon="material-symbols:chevron-left" />
                      </button>
                      <button
                        onClick={nextImage}
                        className="absolute right-2 top-1/2 transform -translate-y-1/2 bg-white bg-opacity-75 rounded-full p-2 hover:bg-opacity-100 transition-all duration-200 shadow-md"
                        aria-label="Next image"
                      >
                        <Icon icon="material-symbols:chevron-right" />
                      </button>
                    </>
                  )}
                </div>

                {/* Image Thumbnails */}
                {images.length > 1 && (
                  <div className="flex space-x-2 p-4 overflow-x-auto bg-gray-50">
                    {images.map((image, index) => (
                      <button
                        key={index}
                        onClick={() => setCurrentImageIndex(index)}
                        className={`flex-shrink-0 w-16 h-16 rounded-md overflow-hidden border-2 ${
                          currentImageIndex === index
                            ? "border-blue-500"
                            : "border-transparent"
                        }`}
                      >
                        <div className="relative w-full h-full">
                          <Image
                            src={image}
                            alt={`${product.name} thumbnail ${index + 1}`}
                            layout="fill"
                            objectFit="cover"
                            className="select-none"
                            draggable="false"
                          />
                        </div>
                      </button>
                    ))}
                  </div>
                )}
              </PhotoProvider>
            ) : (
              <div className="relative w-full h-64 md:h-80 bg-gray-100">
                <div className="relative w-full h-full">
                  <Image
                    src="/images/products/no_product.png"
                    alt={product.name}
                    layout="fill"
                    objectFit="contain"
                    className="select-none"
                    draggable="false"
                  />
                </div>
              </div>
            )}

            {/* Product Details */}
            <div className="p-4 md:p-6 space-y-4">
              <div>
                <h3 className="text-lg font-semibold mb-2">Descripción</h3>
                <p className="text-gray-700">{product.description}</p>
              </div>

              <div className="grid grid-cols-2 gap-4">
                <Box className="bg-gray-50 p-3 rounded-lg">
                  <p className="text-sm text-gray-500">SKU</p>
                  <p className="font-medium">{product.sku}</p>
                </Box>
                <Box className="bg-gray-50 p-3 rounded-lg">
                  <p className="text-sm text-gray-500">Cantidad disponible</p>
                  <p className="font-medium">{product.quantity}</p>
                </Box>
                <Box className="bg-gray-50 p-3 rounded-lg">
                  <p className="text-sm text-gray-500">Precio</p>
                  <p className="font-medium text-green-600">
                    ${numberMoveDecimal(product.price, 2)}
                  </p>
                </Box>
                <Box className="bg-gray-50 p-3 rounded-lg">
                  <p className="text-sm text-gray-500">Estado</p>
                  <p className="font-medium">
                    {product.productStatus === 1 ? "Activo" : "Inactivo"}
                  </p>
                </Box>
              </div>
            </div>
          </div>
        </ModalBody>
      </ModalContent>
    </Modal>
  );
};

export default ProductDetailModal;