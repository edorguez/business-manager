export type WebsocketMessage = {
  from: string;
  message: string;
  messageType: WebsocketMessageTypes;
  sent: Date;
}

export enum WebsocketMessageTypes {
    QR_CODE = 'QR_CODE',
    CONVERSATIONS_CODE = 'CONVERSATIONS_CODE',
    MESSAGE_CODE = 'MESSAGE_CODE',
}
