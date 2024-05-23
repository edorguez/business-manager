'use client';

import useWhatsappMessage from "@/app/hooks/useWhatsappMessage";
import { WhatsappMessage, WhatsappMessageType } from "@/app/types/whatsapp";
import { useEffect, useState } from "react";

const MessageList = () => {
  const wsm: WhatsappMessage[] = [
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
      message: 'Bien y tu, oi que tuviste un gran dia el cual quieres hablar de el pero no he tenido tiempo debido a que estoy pogramando jajaja es que es muy fino wn sbias?',
      type: WhatsappMessageType.Received,
      date: new Date()
    },
    {
      id: 3,
      message: 'Chevere gracias bro es que este es un mensaje largo que quiero probar para ver como se ve con el sistema porque aja si es muy largo no tiene sentido que pruebe esto!',
      type: WhatsappMessageType.Sent,
      date: new Date()
    },
  ];

  const [messages, setMessages] = useState<WhatsappMessage[]>(wsm);
  const whatsappMessage = useWhatsappMessage();

  const getDateMessageFormat = (date: Date): string => {
    return date.toLocaleString('en-US', { hour: 'numeric', minute: 'numeric', hour12: true });
  }


  useEffect(() => {
    const newMessage: WhatsappMessage = {
      id: messages.length,
      message: whatsappMessage.message,
      type: WhatsappMessageType.Sent,
      date: new Date()
    }
    
    setMessages(prevVal => [...prevVal, newMessage]);
  }, [whatsappMessage.message])

  return (
    <div className="flex flex-col p-1">
      {
        messages.map((message: WhatsappMessage, idx: number) => (
          <div key={idx} 
                className={`
                  max-w-[70%]
                  rounded
                  mt-1
                  p-1
                  ${message.type === WhatsappMessageType.Sent ? 'self-end bg-whatsappmessage' : 'self-start bg-white'}
                `}
          >

            {message.type === WhatsappMessageType.Sent && (
              <label className="text-sm">
                {message.message}
                <small className="ml-1.5 text-[10px] text-slate-400 select-none">{getDateMessageFormat(message.date)}</small>
              </label>
            )}

            {message.type === WhatsappMessageType.Received && (
              <label className="text-sm">
                {message.message}
                <small className="ml-1.5 text-[10px] text-slate-400 select-none">{getDateMessageFormat(message.date)}</small>
              </label>
            )}

          </div>
        ))
      }
    </div>
  )
}

export default MessageList;
