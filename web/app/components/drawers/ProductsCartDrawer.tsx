"use client";

import useCheckoutModal from "@/app/hooks/useCheckoutModal";
import useProductsCart from "@/app/hooks/useProductsCart";
import { numberMoveDecimal } from "@/app/utils/Utils";
import {
  Button,
  Drawer,
  DrawerBody,
  DrawerCloseButton,
  DrawerContent,
  DrawerHeader,
  DrawerOverlay,
} from "@chakra-ui/react";
import CheckoutModal from "../modals/CheckoutModal";

const ProductsCartDrawer = () => {
  const cart = useProductsCart();
  const checkoutModal = useCheckoutModal();

  const getTotalPrice = (): string => {
    let res: number = cart.items.reduce((total, item) => total + numberMoveDecimal(item.price, 2) * item.quantity, 0);
    return res.toFixed(2);
  };

  const onCheckout = (): void => {
    cart.onClose();
    checkoutModal.onOpen();
  }

  return (
    <>
      <CheckoutModal />
      <Drawer isOpen={cart.isOpen} placement="right" onClose={cart.onClose}>
        <DrawerOverlay />
        <DrawerContent>
          <DrawerCloseButton />
          <DrawerHeader>Carrito</DrawerHeader>
          <DrawerBody>
            {cart.items.length === 0 ? (
              <p>Carrito vacío.</p>
            ) : (
              <>
                {cart.items.map((item) => {
                    return (
                        <div
                            key={item.id}
                            className="flex justify-between items-center mb-4"
                        >
                            <div>
                                <h3 className="font-semibold">{item.name}</h3>
                                <p className="text-sm text-gray-600">
                                    ${numberMoveDecimal(item.price, 2)} x {item.quantity}
                                </p>
                            </div>
                            <div className="flex items-center">
                                <Button
                                    size="sm"
                                    onClick={() => cart.onRemoveFromCart(item.id)}
                                >
                                    -
                                </Button>
                                <span className="mx-2">{item.quantity}</span>
                                <Button size="sm" onClick={() => cart.onAddToCart(item)}>
                                    +
                                </Button>
                            </div>
                        </div>
                    );
                })}
                <div className="mt-4 pt-4 border-t">
                  <div className="flex justify-between items-center mb-4">
                    <span className="font-semibold">Total:</span>
                    <span className="font-bold">${getTotalPrice()}</span>
                  </div>
                  <Button variant="default" width="100%" onClick={onCheckout}>
                    Realizar Pedido
                  </Button>
                </div>
              </>
            )}
          </DrawerBody>
        </DrawerContent>
      </Drawer>
    </>
  );
};

export default ProductsCartDrawer;
