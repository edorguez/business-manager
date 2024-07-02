'use client';

import { WhatsappConversation } from "@/app/types/whatsapp";
import { Button, Skeleton, Stack } from "@chakra-ui/react";
import { Icon } from "@iconify/react";
import Image from "next/image";

interface ChatListProps {
  conversations: WhatsappConversation[] | undefined,
  onSelectConversation: (id: string) => void
}

const ChatList: React.FC<ChatListProps> = ({
  conversations,
  onSelectConversation
}) => {

  const src = 'https://canto-wp-media.s3.amazonaws.com/app/uploads/2019/08/19194138/image-url-3.jpg';

  const previewMsgFormat = (msg: string): string => {
    const maxLength: number = 40;
    if (msg.length >= maxLength) {
      return msg.substring(0, maxLength) + '...';
    }

    return msg;
  }

  const chatsLoading: any[] = [...Array(13)].map(item => {
    return { ...item, id: crypto.randomUUID() }
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
          <div key={conversation.id} className="flex items-center px-1 py-2 hover:bg-thirdcolorhov hover:cursor-pointer select-none" onClick={() => onSelectConversation(conversation.id)}>
            <div className="w-[50px] h-[50px]">
              <Image className="rounded-full" alt={conversation.name} loader={() => src} src={src} width={50} height={50} />
            </div>
            <div className="ml-2 flex justify-between items-center w-full">
              <div>
                <span className="text-sm font-medium">{conversation.name}</span>
                <br />
                <span className="text-xs flex items-center">
                  <Icon icon="mdi:check-all" />
                  <span className="ml-1">
                    {previewMsgFormat(conversation.messages[0].message)}
                  </span>
                </span>
              </div>
              <Button variant="main" size="xs">Atender</Button>
            </div>
          </div>
        ))
      }
    </>
  )
}

export default ChatList;
