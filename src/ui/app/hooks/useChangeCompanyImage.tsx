import { create } from "zustand";

interface ChangeCompanyImageStore {
  emitSignal: boolean;
  triggerSignal: () => void;
  resetSignal: () => void;
}

const useChangeCompanyImage = create<ChangeCompanyImageStore>((set) => ({
  emitSignal: false,
  triggerSignal: () => set({ emitSignal: true }),
  resetSignal: () => set({ emitSignal: false }),
}));

export default useChangeCompanyImage;
