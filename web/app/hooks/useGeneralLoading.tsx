import { create } from 'zustand';

interface LoadingStore {
  isLoading: boolean;
  onStartLoading: () => void;
  onEndLoading: () => void;
}

const useGeneralLoading = create<LoadingStore>((set) => ({
  isLoading: true,
  onStartLoading: () => set({ isLoading: true }),
  onEndLoading: () => set({ isLoading: false })
}));

export default useGeneralLoading;