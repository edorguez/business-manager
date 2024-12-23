import { create } from "zustand";

interface WindowSizeStore {
  width: number | undefined;
  height: number | undefined;
  setWindowSize: (width: number, height: number) => void;
}

const useWindowSize = create<WindowSizeStore>((set) => ({
  width: undefined,
  height: undefined,
  setWindowSize: (width, height) => set({ width, height }),
}));

export default useWindowSize;
