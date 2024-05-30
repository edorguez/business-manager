'use client';

import useWhatsappMessage from "@/app/hooks/useWhatsappMessage";
import { Input } from "@chakra-ui/react";
import { Icon } from "@iconify/react";
import { useState } from "react";

interface MessageBarProps {
  onSendMessage: (message: string) => void;
}

const MessageBar: React.FC<MessageBarProps> = ({
  onSendMessage
}) => {
  
  const whatsappMessage = useWhatsappMessage();
  const [message, setMessage] = useState('');

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setMessage(e.target.value);
  }

  const handleKeys = (e: any) => {
    if(e.key === 'Enter')
      handleSendMessage();
  }

  const handleSendMessage = () => {
    if(message.trim() !== '') {
      whatsappMessage.setMessage(message);
      onSendMessage(message);
      setMessage('');
    }
  }

  return (
    <div className="p-1 bg-slate-200 border-t-2 border-slate-300">
      <div className="flex">
        <div className="mx-2 flex justify-center items-center rounded-full text-slate-500 hover:text-thirdcolor cursor-pointer transition duration-100">
          <Icon icon="fluent:wallet-credit-card-16-filled" width={22} height={22} />
        </div>
        <div className="mx-2 flex justify-center items-center rounded-full text-slate-500 hover:text-thirdcolor cursor-pointer transition duration-100">
          <Icon icon="fluent-mdl2:product" width={22} height={22} />
        </div>
        <Input bg="white" size='sm' placeholder='Mensaje' onChange={handleChange} onKeyDown={handleKeys} value={message} />
        <div onClick={handleSendMessage} className="mx-2 p-2 flex justify-center items-center rounded-full bg-thirdcolor text-white hover:bg-maincolor cursor-pointer transition duration-100">
          <Icon icon="material-symbols:send" />
        </div>
      </div>
    </div>
  )
}

export default MessageBar;
