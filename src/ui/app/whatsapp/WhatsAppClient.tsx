'use client';

import { Tab, TabList, TabPanel, TabPanels, Tabs } from "@chakra-ui/react";
import SimpleCard from "../components/cards/SimpleCard";
import { Icon } from "@iconify/react";
import ChatList from "../components/whatsapp/ChatList";
import MessageList from "../components/whatsapp/MessageList";
import MessageBar from "../components/whatsapp/MessageBar";
import UserBar from "../components/whatsapp/UserBar";

const WhatsAppClient = () => {
  return (
    <SimpleCard>
      <div className="h-[85vh] flex my-3 rounded border-2 border-slate-200">
        <div className="w-3/6 border-r-2 border-r-slate-200 overflow-scroll">
          <Tabs colorScheme='green'>
            <TabList>
              <Tab><Icon icon="ic:baseline-message" /><span className="text-sm"> Mensajes</span></Tab>
              <Tab><Icon icon="fluent:person-support-16-filled" /><span className="text-sm flex items-center"> Atendiendo</span></Tab>
              <Tab><Icon icon="lets-icons:done-duotone" /><span className="text-sm flex items-center"> Finalizado</span></Tab>
            </TabList>

            <TabPanels>
              <TabPanel>
                <ChatList />
              </TabPanel>
              <TabPanel>
                <ChatList />
              </TabPanel>
              <TabPanel>
                <ChatList />
              </TabPanel>
            </TabPanels>
          </Tabs>
        </div>
        <div className="w-full" style={{ backgroundImage: 'url(/images/whatsapp/ws_bg.png)' }}>
          <div className="h-full flex flex-col">
            <UserBar />
            <MessageList />
            <div className="mt-auto">
              <MessageBar />
            </div>
          </div>

        </div>
      </div>
    </SimpleCard>
  )
}

export default WhatsAppClient;
