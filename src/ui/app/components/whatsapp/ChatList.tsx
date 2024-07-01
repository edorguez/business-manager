'use client';

import { WhatsappConversation } from "@/app/types/whatsapp";
import { Button, Skeleton, Stack } from "@chakra-ui/react";
import Image from "next/image";

interface ChatListProps {
  conversations: WhatsappConversation[] | undefined
}

const ChatList: React.FC<ChatListProps> = ({
  conversations
}) => {

  const src = 'https://canto-wp-media.s3.amazonaws.com/app/uploads/2019/08/19194138/image-url-3.jpg';

  const chatsLoading: any[] = [...Array(13)].map(item => {
    return {...item, id: crypto.randomUUID()}
  });

  return (
    <>
      {
        !conversations &&
        <Stack>
          {
            chatsLoading.map((item) => (
              <Skeleton key={item.id} height='60px' />
            ))
          }
        </Stack>
      }

      {
        conversations &&
        conversations.map((conversation: WhatsappConversation) => (
           <div key={conversation.id}>
              <div className="flex items-center px-1 py-2 hover:bg-thirdcolorhov hover:cursor-pointer select-none">
                <div className="w-[50px] h-[50px]">
                  <Image className="rounded-full" alt={conversation.name} loader={() => src} src={src}  width={50} height={50}/>
                </div>
                <div className="ml-2 flex justify-between items-center w-full">
                  <div>
                    <span className="text-sm font-medium">{conversation.name}</span>
                    <br />
                    <span className="text-xs">{conversation.messages[0].message.substring(0, 20)}...</span>
                  </div>
                  <Button variant="main" size="xs">Atender</Button>
                </div>
              </div>
          </div> 
        ))
      }
    </>
  )
}

export default ChatList;
