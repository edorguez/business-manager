export type WhatsappMessage = {
  id: number;
  message: string;
  type: WhatsappMessageType;
  date: Date;
}

export enum WhatsappMessageType {
  Sent,
  Received
}
