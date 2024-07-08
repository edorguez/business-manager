'use client';

import { Tab, TabList, TabPanel, TabPanels, Tabs } from "@chakra-ui/react";
import { Icon } from "@iconify/react";
import ChatList from "./ChatList";
import UserBar from "./UserBar";
import MessageList from "./MessageList";
import MessageBar from "./MessageBar";
import SelectChatMessage from "./SelectChatMessage";
import { WhatsappConversation } from "@/app/types/whatsapp";
import { useState } from "react";

interface ChatViewProps {
  conversations: WhatsappConversation[],
}

const ChatView: React.FC<ChatViewProps> = ({
  conversations
}) => {
  const [selectedConversation, setSelectedConversation] = useState<WhatsappConversation>();

  const selectConversation = (id: string): void => {
    const findConversation: WhatsappConversation | undefined = conversations?.find(x => x.id === id);
    if (findConversation)
      setSelectedConversation(findConversation);
  }

  const sendWhatsappMessage = (message: string): void => {
    console.log('Enviar primero el mensaje');
    console.log(message);

    const result: any = {
      type: 'send_message',
      payload: {
        message: message,
        from: 'Eduardo'
      }
    };
    // sendMessage(JSON.stringify(result));
  }

  return (
    <div className="h-[85vh] flex my-3 rounded border-2 border-slate-200">
      <div className="w-3/6 border-r-2 border-r-slate-200 overflow-scroll">
        <Tabs colorScheme='green'>
          <TabList>
            <Tab><Icon icon="ic:baseline-message" /><span className="text-sm"> Mensajes</span></Tab>
            <Tab><Icon icon="fluent:person-support-16-filled" /><span className="text-sm flex items-center"> Atendiendo</span></Tab>
            <Tab><Icon icon="lets-icons:done-duotone" /><span className="text-sm flex items-center"> Finalizado</span></Tab>
          </TabList>

          <TabPanels>
            <TabPanel p={0}>
              <ChatList conversations={conversations} onSelectConversation={selectConversation} />
            </TabPanel>
            <TabPanel p={0}>
              <ChatList conversations={conversations} onSelectConversation={selectConversation} />
            </TabPanel>
            <TabPanel p={0}>
              <ChatList conversations={conversations} onSelectConversation={selectConversation} />
            </TabPanel>
          </TabPanels>
        </Tabs>
      </div>
      <div className="w-full" style={{ backgroundImage: 'url(/images/whatsapp/ws_bg.jpg)', backgroundSize: 'contain', backgroundRepeat: 'repeat' }}>
        {
          selectedConversation &&
          <div className="h-full flex flex-col">
            <UserBar name={selectedConversation.name} imageUrl={selectedConversation.profilePictureUrl} />
            <div className="mt-auto overflow-y-auto">
              <MessageList messages={selectedConversation.messages} />
            </div>
            <MessageBar onSendMessage={sendWhatsappMessage} />
          </div>
        }

        {
          !selectedConversation &&
          <div className="h-full flex justify-center items-center">
            <SelectChatMessage />
          </div>
        }
      </div>
    </div>

  )
}

export default ChatView;
