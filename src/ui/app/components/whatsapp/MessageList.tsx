'use client';

import { WhatsappMessage, WhatsappMessageType } from "@/app/types/whatsapp";

const MessageList = () => {
  const messages: WhatsappMessage[] = [
    {
      id: 0,
      message: 'Hola',
      type: WhatsappMessageType.Received,
      date: new Date()
    },
    {
      id: 1,
      message: 'Como estas?',
      type: WhatsappMessageType.Sent,
      date: new Date()
    },
    {
      id: 2,
      message: 'Bien y tu',
      type: WhatsappMessageType.Received,
      date: new Date()
    },
    {
      id: 3,
      message: 'Chevere gracias bro!',
      type: WhatsappMessageType.Sent,
      date: new Date()
    },

  ];

  return (
    <div className="flex flex-col">
      {
        messages.map((message: WhatsappMessage) => (
          <div key={message.id} 
                className={`
                  ${message.type === WhatsappMessageType.Sent ? 'self-end' : 'self-auto'}
                `}
          >

            {message.type === WhatsappMessageType.Sent && (
              <label className="bg-slate-200">
                {message.message}
              </label>
            )}

            {message.type === WhatsappMessageType.Received && (
              <label className="bg-blue-200">
                {message.message}
              </label>
            )}

          </div>
        ))
      }
    </div>
  )
}

export default MessageList;
