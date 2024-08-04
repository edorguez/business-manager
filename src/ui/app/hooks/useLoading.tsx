import { create } from 'zustand';

interface LoadingStore {
  isLoading: boolean;
  onStartLoading: () => void;
  onEndLoading: () => void;
}

const useLoading = create<LoadingStore>((set) => ({
  isLoading: false,
  onStartLoading: () => set({ isLoading: true }),
  onEndLoading: () => set({ isLoading: false })
}));

export default useLoading;