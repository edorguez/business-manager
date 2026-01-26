import { create } from 'zustand';
import { Product } from '../types/product';

interface ProductDetailModalStore {
  isOpen: boolean;
  product: Product | null;
  onOpen: (product: Product) => void;
  onClose: () => void;
}

const useProductDetailModal = create<ProductDetailModalStore>((set) => ({
  isOpen: false,
  product: null,
  onOpen: (product: Product) => set({ isOpen: true, product }),
  onClose: () => set({ isOpen: false, product: null })
}));

export default useProductDetailModal;