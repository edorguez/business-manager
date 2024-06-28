export type WhatsappConversation = {
  id: string;
  name: string;
  profilePictureUrl: string;
  unreadCount: number;
  messages: WhatsappMessage[];
}

export type WhatsappMessage = {
  id: string;
  message: string;
  date: Date;
  wasReceipt: boolean;
  wasRead: boolean;
  fromMe: boolean;
}
