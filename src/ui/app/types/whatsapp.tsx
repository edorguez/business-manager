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

export type CreateBusinessPhone = {
  companyId: number;
  phone: string;
}

export type EditBusinessPhone = {
  id: number;
  companyId: number;
  phone: string;
}

export type BusinessPhone = {
  id: number;
  companyId: number;
  phone: string;
}