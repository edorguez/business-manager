export type WebsocketMessage = {
  from: string;
  message: string;
  messageType: WebsocketMessageTypes;
  sent: Date;
}

export enum WebsocketMessageTypes {
    QR = 'QR',
    Conversations = 'CONVERSATIONS',
    Message = 'MESSAGE',
}
