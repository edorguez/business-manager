'use client';

import useWhatsappMessage from "@/app/hooks/useWhatsappMessage";
import { WhatsappMessage } from "@/app/types/whatsapp";
import { useEffect, useRef, useState } from "react";
import dayjs from 'dayjs';

interface MessageListProps {
  messages: WhatsappMessage[]
}

const MessageList: React.FC<MessageListProps> = ({
  messages
}) => {
  const whatsappMessage = useWhatsappMessage();
  const scrollRef = useRef<any>(null);

  const getDateMessageFormat = (stringDate: string): string => {
    return dayjs(stringDate).format('hh:mm A')
  }

  useEffect(() => {
    if(whatsappMessage.message) {
      const newMessage: WhatsappMessage = {
        id: "A",
        message: whatsappMessage.message,
        fromMe: true,
        date: new Date(),
        wasReceipt: false,
        wasRead: false
      }

      messages.push(newMessage);
    
      // setMessages(prevVal => [...prevVal, newMessage]);

      // Scroll chat to bottom
      setTimeout(() => { scrollChatToBottom(); }, 1);

    }
  }, [whatsappMessage.message])

  const scrollChatToBottom = () => {
    scrollRef.current?.scrollIntoView({ behavior: "smooth" })
  }

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
                  ${message.fromMe ? 'self-end bg-whatsappmessage' : 'self-start bg-white'}
                `}
          >

            {message.fromMe && (
              <label className="text-sm">
                {message.message}
                <small className="ml-1.5 text-[10px] text-slate-400 select-none">{getDateMessageFormat(message.date.toDateString())}</small>
              </label>
            )}

            {!message.fromMe && (
              <label className="text-sm">
                {message.message}
                <small className="ml-1.5 text-[10px] text-slate-400 select-none">{getDateMessageFormat(message.date.toDateString())}</small>
              </label>
            )}

          </div>
        ))
      }
      <div ref={scrollRef}/>
    </div>
  )
}

export default MessageList;
