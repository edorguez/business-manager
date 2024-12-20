import { create } from "zustand";
import { Product } from "../types/product";

interface ProductsCartStore {
  isOpen: boolean;
  items: CartItem[];
  onOpen: () => void;
  onClose: () => void;
  onAddToCart: (product: Product) => void;
  onRemoveFromCart: (productId: string) => void;
}

export interface CartItem extends Product {
  quantity: number;
}

const useProductsCart = create<ProductsCartStore>((set) => ({
  isOpen: false,
  items: [],
  onOpen: () => set({ isOpen: true }),
  onClose: () => set({ isOpen: false }),
  onAddToCart: (product: Product) => {
    set((state) => {
      const existingItem = state.items.find((item) => item.id === product.id);

      const updatedCarts = existingItem
        ? state.items.map((item) =>
            item.id === product.id
              ? { ...item, quantity: item.quantity + 1 }
              : item
          )
        : [...state.items, { ...product, quantity: 1 }];

      return { ...state, items: updatedCarts };
    });
  },
  onRemoveFromCart: (productId: string) => {
    set((state) => {
      const updatedCarts = state.items
        .map((item) =>
          item.id === productId
            ? { ...item, quantity: item.quantity - 1 }
            : item
        )
        .filter((item) => item.quantity > 0);

      return { ...state, items: updatedCarts };
    });
  },
}));

export default useProductsCart;
