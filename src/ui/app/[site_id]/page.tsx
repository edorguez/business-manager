"use client";

import {
  Button,
  Drawer,
  DrawerBody,
  DrawerCloseButton,
  DrawerContent,
  DrawerHeader,
  DrawerOverlay,
  Menu,
  MenuButton,
  MenuItem,
  MenuList,
  useDisclosure,
} from "@chakra-ui/react";
import { useCallback, useEffect, useState } from "react";
import { Company } from "../types/company";
import { Product } from "../types/product";
import ProductCard from "../components/cards/ProductCard";
import { GetCompanyByNameRequest } from "../services/companies";
import { useParams, useRouter } from "next/navigation";
import useGeneralLoading from "../hooks/useGeneralLoading";

interface CartItem extends Product {
  quantity: number;
}

const products: Product[] = [
  {
    id: "1",
    name: "Classic Burger",
    description: "Juicy beef patty with fresh toppings",
    price: 9.99,
    images: ["/placeholder.svg?height=200&width=200"],
    sku: '',
    quantity: 1,
    productStatus: 1
  },
  {
    id: "2",
    name: "Classic Burger",
    description: "Juicy beef patty with fresh toppings",
    price: 9.99,
    images: ["/placeholder.svg?height=200&width=200"],
    sku: '',
    quantity: 1,
    productStatus: 1
  },
  {
    id: "3",
    name: "Classic Burger",
    description: "Juicy beef patty with fresh toppings",
    price: 9.99,
    images: ["/placeholder.svg?height=200&width=200"],
    sku: '',
    quantity: 1,
    productStatus: 1
  },
  {
    id: "4",
    name: "Classic Burger",
    description: "Juicy beef patty with fresh toppings",
    price: 9.99,
    images: ["/placeholder.svg?height=200&width=200"],
    sku: '',
    quantity: 1,
    productStatus: 1
  },
  {
    id: "5",
    name: "Classic Burger",
    description: "Juicy beef patty with fresh toppings",
    price: 9.99,
    images: ["/placeholder.svg?height=200&width=200"],
    sku: '',
    quantity: 1,
    productStatus: 1
  },
];

const SitePage = () => {
  const [cart, setCart] = useState<CartItem[]>([]);
  const { isOpen, onOpen, onClose } = useDisclosure();
  // const data: Company = useData();


  const router = useRouter();
  const params = useParams();
  // const [isLoading, setIsLoading] = useState<boolean>(true);
  const isLoading = useGeneralLoading();
  const [company, setCompany] = useState<Company | null>(null);
  
  const getCompany = useCallback(async () => {
    let getCompany: Company = await GetCompanyByNameRequest(params.site_id.toString());
    console.log(getCompany);
     if(!getCompany?.id || getCompany?.lastPaymentDate < new Date()) {
      console.log('EPA FUERA');
      // I need to create my not found route
      router.push('/404')
     } else {
      setCompany(getCompany)
      // setIsLoading(false);

      setTimeout(() => {
        isLoading.onEndLoading();
      }, 2000);
     }
  }, [params.site_id, router]);

  useEffect(() => {
    getCompany();
  }, [getCompany])

  const addToCart = (product: Product): void => {
    setCart((prevCart) => {
      const existingItem = prevCart.find((item) => item.id === product.id);
      if (existingItem) {
        return prevCart.map((item) =>
          item.id === product.id
            ? { ...item, quantity: item.quantity + 1 }
            : item
        );
      } else {
        return [...prevCart, { ...product, quantity: 1 }];
      }
    });
  };

  // const removeFromCart = (productId: number): void => {
  //   setCart((prevCart) => {
  //     const existingItem = prevCart.find((item) => item.id === productId);
  //     if (existingItem && existingItem.quantity === 1) {
  //       return prevCart.filter((item) => item.id !== productId);
  //     } else {
  //       return prevCart.map((item) =>
  //         item.id === productId
  //           ? { ...item, quantity: item.quantity - 1 }
  //           : item
  //       );
  //     }
  //   });
  // };

  const getTotalPrice = (): string => {
    return cart
      .reduce((total, item) => total + item.price * item.quantity, 0)
      .toFixed(2);
  };

  const getTotalItems = (): number => {
    return cart.reduce((total, item) => total + item.quantity, 0);
  };

  return (
    <div>
      <div className="min-h-screen bg-gray-100">
        {/* Header */}
        <header className="bg-white shadow-md">
          <div className="container mx-auto px-4 py-4 flex items-center justify-between">
            <h1 className="text-2xl font-bold text-green-500">
              Burger Haven
            </h1>
            <div className="flex items-center space-x-4">
              <Button colorScheme="green" onClick={onOpen}>
                Cart ({getTotalItems()})
              </Button>
              <Menu>
                <MenuButton as={Button}>Account</MenuButton>
                <MenuList>
                  <MenuItem>Profile</MenuItem>
                  <MenuItem>Orders</MenuItem>
                  <MenuItem>Settings</MenuItem>
                  <MenuItem>Logout</MenuItem>
                </MenuList>
              </Menu>
            </div>
          </div>
        </header>

        {/* Main Content */}
        <main className="container mx-auto px-4 py-8">
          <h2 className="text-2xl font-semibold mb-6">Our Menu</h2>
          <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-8">
            {products.map((product) => (
              <ProductCard key={product.id} product={product} onAddToCard={() => { alert("hola") }} />
            ))}
          </div>
        </main>

        {/* Shopping Cart Drawer */}
        <Drawer isOpen={isOpen} placement="right" onClose={onClose}>
          <DrawerOverlay />
          <DrawerContent>
            <DrawerCloseButton />
            <DrawerHeader>Your Cart</DrawerHeader>
            <DrawerBody>
              {cart.length === 0 ? (
                <p>Your cart is empty.</p>
              ) : (
                <>
                  {cart.map((item) => (
                    <div
                      key={item.id}
                      className="flex justify-between items-center mb-4"
                    >
                      <div>
                        <h3 className="font-semibold">{item.name}</h3>
                        <p className="text-sm text-gray-600">
                          ${item.price.toFixed(2)} x {item.quantity}
                        </p>
                      </div>
                      <div className="flex items-center">
                        <Button
                          size="sm"
                          // onClick={() => removeFromCart(item.id)}
                        >
                          -
                        </Button>
                        <span className="mx-2">{item.quantity}</span>
                        <Button size="sm" onClick={() => addToCart(item)}>
                          +
                        </Button>
                      </div>
                    </div>
                  ))}
                  <div className="mt-4 pt-4 border-t">
                    <div className="flex justify-between items-center mb-4">
                      <span className="font-semibold">Total:</span>
                      <span className="font-bold">${getTotalPrice()}</span>
                    </div>
                    <Button colorScheme="green" width="100%">
                      Proceed to Checkout
                    </Button>
                  </div>
                </>
              )}
            </DrawerBody>
          </DrawerContent>
        </Drawer>
      </div>
    </div>
  );
}

export default SitePage;