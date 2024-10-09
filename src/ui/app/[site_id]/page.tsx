'use client';

import { Button, Drawer, DrawerBody, DrawerCloseButton, DrawerContent, DrawerHeader, DrawerOverlay, Menu, MenuButton, MenuItem, MenuList, useDisclosure } from "@chakra-ui/react";
import { useState } from "react";

interface Product {
  id: number;
  name: string;
  description: string;
  price: number;
  image: string;
}

interface CartItem extends Product {
  quantity: number;
}

const products: Product[] = [
  { id: 1, name: "Classic Burger", description: "Juicy beef patty with fresh toppings", price: 9.99, image: "/placeholder.svg?height=200&width=200" },
  { id: 2, name: "Veggie Delight", description: "Plant-based patty with avocado", price: 8.99, image: "/placeholder.svg?height=200&width=200" },
  { id: 3, name: "Chicken Sandwich", description: "Grilled chicken breast with special sauce", price: 10.99, image: "/placeholder.svg?height=200&width=200" },
  { id: 4, name: "Loaded Fries", description: "Crispy fries topped with cheese and bacon", price: 6.99, image: "/placeholder.svg?height=200&width=200" },
  { id: 5, name: "Milkshake", description: "Creamy shake in various flavors", price: 4.99, image: "/placeholder.svg?height=200&width=200" },
  { id: 6, name: "Caesar Salad", description: "Fresh romaine lettuce with Caesar dressing", price: 7.99, image: "/placeholder.svg?height=200&width=200" },
]

export default function Hola() {
    const [cart, setCart] = useState<CartItem[]>([])
  const { isOpen, onOpen, onClose } = useDisclosure()

  const addToCart = (product: Product): void => {
    setCart(prevCart => {
      const existingItem = prevCart.find(item => item.id === product.id)
      if (existingItem) {
        return prevCart.map(item => 
          item.id === product.id ? { ...item, quantity: item.quantity + 1 } : item
        )
      } else {
        return [...prevCart, { ...product, quantity: 1 }]
      }
    })
  }

  const removeFromCart = (productId: number): void => {
    setCart(prevCart => {
      const existingItem = prevCart.find(item => item.id === productId)
      if (existingItem && existingItem.quantity === 1) {
        return prevCart.filter(item => item.id !== productId)
      } else {
        return prevCart.map(item => 
          item.id === productId ? { ...item, quantity: item.quantity - 1 } : item
        )
      }
    })
  }

  const getTotalPrice = (): string => {
    return cart.reduce((total, item) => total + item.price * item.quantity, 0).toFixed(2)
  }

  const getTotalItems = (): number => {
    return cart.reduce((total, item) => total + item.quantity, 0)
  }

  return (
    <div>
         <h1 className="bg-rose-800">HOLAAAAAAAAAAAAA</h1>
      <div className="min-h-screen bg-gray-100">
        {/* Header */}
        <header className="bg-white shadow-md">
          <div className="container mx-auto px-4 py-4 flex items-center justify-between">
            <h1 className="text-2xl font-bold text-green-500">Burger Haven</h1>
            <div className="flex items-center space-x-4">
              <Button colorScheme="green" onClick={onOpen}>
                Cart ({getTotalItems()})
              </Button>
              <Menu>
                <MenuButton as={Button}>
                  Account
                </MenuButton>
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
              <div key={product.id} className="bg-white rounded-lg shadow-md overflow-hidden hover:shadow-lg transition-shadow duration-300">
                <img src={product.image} alt={product.name} className="w-full h-48 object-cover" />
                <div className="p-4">
                  <h3 className="text-lg font-semibold mb-2">{product.name}</h3>
                  <p className="text-gray-600 mb-2">{product.description}</p>
                  <div className="flex items-center justify-between">
                    <span className="text-lg font-bold">${product.price.toFixed(2)}</span>
                    <Button colorScheme="green" size="sm" onClick={() => addToCart(product)}>
                      Add to Cart
                    </Button>
                  </div>
                </div>
              </div>
            ))}
          </div>
        </main>

        {/* Footer */}
        <footer className="bg-gray-800 text-white py-8">
          <div className="container mx-auto px-4">
            <div className="grid grid-cols-1 md:grid-cols-3 gap-8">
              <div>
                <h3 className="text-lg font-semibold mb-4">About Burger Haven</h3>
                <ul className="space-y-2">
                  <li><a href="#" className="hover:text-green-400">Our Story</a></li>
                  <li><a href="#" className="hover:text-green-400">Locations</a></li>
                  <li><a href="#" className="hover:text-green-400">Careers</a></li>
                </ul>
              </div>
              <div>
                <h3 className="text-lg font-semibold mb-4">Customer Service</h3>
                <ul className="space-y-2">
                  <li><a href="#" className="hover:text-green-400">Contact Us</a></li>
                  <li><a href="#" className="hover:text-green-400">FAQs</a></li>
                  <li><a href="#" className="hover:text-green-400">Allergen Information</a></li>
                </ul>
              </div>
              <div>
                <h3 className="text-lg font-semibold mb-4">Connect With Us</h3>
                <ul className="space-y-2">
                  <li><a href="#" className="hover:text-green-400">Facebook</a></li>
                  <li><a href="#" className="hover:text-green-400">Instagram</a></li>
                  <li><a href="#" className="hover:text-green-400">Twitter</a></li>
                </ul>
              </div>
            </div>
            <div className="mt-8 text-center text-gray-400">
              © 2023 Burger Haven. All rights reserved.
            </div>
          </div>
        </footer>

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
                    <div key={item.id} className="flex justify-between items-center mb-4">
                      <div>
                        <h3 className="font-semibold">{item.name}</h3>
                        <p className="text-sm text-gray-600">${item.price.toFixed(2)} x {item.quantity}</p>
                      </div>
                      <div className="flex items-center">
                        <Button size="sm" onClick={() => removeFromCart(item.id)}>-</Button>
                        <span className="mx-2">{item.quantity}</span>
                        <Button size="sm" onClick={() => addToCart(item)}>+</Button>
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
  )
}