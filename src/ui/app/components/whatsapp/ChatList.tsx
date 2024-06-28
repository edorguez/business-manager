'use client';

import { WhatsappConversation } from "@/app/types/whatsapp";
import { Skeleton, Stack } from "@chakra-ui/react";

interface ChatListProps {
  conversations: WhatsappConversation[] | undefined
}

const ChatList: React.FC<ChatListProps> = ({
  conversations
}) => {
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
              <span>{conversation.name}</span>
          </div> 
        ))
      }
    </>
  )
}

export default ChatList;
