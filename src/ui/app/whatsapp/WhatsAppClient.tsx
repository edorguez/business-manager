'use client';

import { Tab, TabList, TabPanel, TabPanels, Tabs } from "@chakra-ui/react";
import SimpleCard from "../components/cards/SimpleCard";
import { Icon } from "@iconify/react";
import ChatList from "../components/whatsapp/ChatList";

const WhatsAppClient = () => {
  return (
    <SimpleCard>
      <div className="h-[85vh] flex my-3 rounded border-2 border-slate-200">
        <div className="w-1/3 border-r-2 border-r-slate-200 overflow-scroll">
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

        </div>
      </div>
    </SimpleCard>
  )
}

export default WhatsAppClient;
