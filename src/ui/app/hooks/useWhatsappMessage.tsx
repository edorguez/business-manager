import { create } from 'zustand';

interface WhatsappMessageStore {
  message: string;
  setMessage: (value: string) => void;
}

const useWhatsappMessage = create<WhatsappMessageStore>((set) => ({
  message: '',
  setMessage: (value: string) =>
    set(() => ({
      message: value,
    })),
}));


export default useWhatsappMessage;
