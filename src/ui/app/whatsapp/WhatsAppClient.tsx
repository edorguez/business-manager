'use client';

import { Tab, TabList, TabPanel, TabPanels, Tabs } from "@chakra-ui/react";
import SimpleCard from "../components/cards/SimpleCard";
import { Icon } from "@iconify/react";
import ChatList from "../components/whatsapp/ChatList";
import MessageList from "../components/whatsapp/MessageList";
import MessageBar from "../components/whatsapp/MessageBar";
import UserBar from "../components/whatsapp/UserBar";
import useWebSocket from 'react-use-websocket';
import { useState } from "react";
import ConnectQr from "../components/whatsapp/ConnectQr";
import { WebsocketMessage, WebsocketMessageTypes } from "../types/websocket";
import { WhatsappConversation } from "../types/whatsapp";
import SelectChatMessage from "../components/whatsapp/SelectChatMessage";

const WS_URL = 'ws://localhost:50055/ws';

const WhatsAppClient = () => {
  const [showQR, setShowQR] = useState<boolean>(true);
  const [qrString, setQrString] = useState<string>('');
  const [messageData, setMessageData] = useState<string>('');
  const [conversations, setConversations] = useState<WhatsappConversation[]>();
  const [selectedConversation, setSelectedConversation] = useState<WhatsappConversation>();
  const { sendMessage, lastMessage, readyState } = useWebSocket(WS_URL, {
    share: true,
    shouldReconnect: () => false,
    onOpen: () => {
      console.log('WebSocket connection established.');


      // const result: WhatsappConversation[] = JSON.parse('[{"id":"584123921057@s.whatsapp.net","name":"584123921057","profilePictureUrl":"","unreadCount":2,"messages":[{"ID":"658183A1277BD7E4CE46A44528FAA3DE","message":"Hola como estas avísame cuando tienes tiempo para pasar por mi pc e instalar los mods si tienes tiempo y ganas que esta ves son muchos mods😥","date":"2024-03-31T10:25:51-04:00","wasReceipt":false,"wasRead":false,"fromMe":false},{"ID":"13751FA61B6B97F4B1EA0E5245B5B871","message":"Hola como as estado no me dejes botado avísame cuando tengas tiempo para instalarme los mods por fa","date":"2024-06-25T12:14:21-04:00","wasReceipt":false,"wasRead":false,"fromMe":false}]},{"id":"584141280555@s.whatsapp.net","name":"584141280555","profilePictureUrl":"https://pps.whatsapp.net/v/t61.24694-24/426814528_7640383659333996_3273897047108463021_n.jpg?ccb=11-4&oh=01_Q5AaIDsf0HiwhhJMfIbPTh7Qh3nPBdgilRlRfshgsfv2iCHH&oe=668F383E&_nc_sid=e6ed6c&_nc_cat=106","unreadCount":0,"messages":[{"ID":"0161B7F992B71240071433195FC82952","message":"","date":"2024-06-19T21:40:07-04:00","wasReceipt":false,"wasRead":false,"fromMe":true}]},{"id":"584241397041@s.whatsapp.net","name":"584241397041","profilePictureUrl":"https://pps.whatsapp.net/v/t61.24694-24/408210515_6882485975201907_7238111534288855850_n.jpg?ccb=11-4&oh=01_Q5AaIHZepO-pgU0sFFRBJN8tQlfyo4P_uuAMknR8la8bApJz&oe=668F21D0&_nc_sid=e6ed6c&_nc_cat=102","unreadCount":0,"messages":[{"ID":"C82185A417F6AD2CEED7106ACE231EDA","message":"","date":"2024-06-10T21:48:02-04:00","wasReceipt":false,"wasRead":false,"fromMe":true},{"ID":"79E5431EDEF34B67DECD05F0D6C8E3C8","message":"","date":"2024-06-10T21:49:05-04:00","wasReceipt":false,"wasRead":false,"fromMe":true},{"ID":"7C282A4C7FCC55AF072CB192B3E8B1C2","message":"Hola buenas noches, quisiera por favor presupuesto aproximado para una manga completa","date":"2024-06-10T21:49:05-04:00","wasReceipt":false,"wasRead":false,"fromMe":true},{"ID":"E362FCDDE0A2FDD925FF5AB268DC74A7","message":"Artista de preferencia: Emersson, Daniel Lopez o OtinZona: brazo completo Color: negro y gris","date":"2024-06-10T21:50:46-04:00","wasReceipt":false,"wasRead":false,"fromMe":true},{"ID":"0C0E3333105B3ECFFF083D9BFC955FBA","message":"","date":"2024-06-10T21:51:11-04:00","wasReceipt":false,"wasRead":false,"fromMe":true},{"ID":"219FBAB5B5E8171E73B2933393217FBA","message":"Cobran por sesion o por proyecto?","date":"2024-06-10T21:51:34-04:00","wasReceipt":false,"wasRead":false,"fromMe":true},{"ID":"096EB16C86DA09A1948FC300E4F05743","message":"Hola buenos dias","date":"2024-06-11T09:25:20-04:00","wasReceipt":false,"wasRead":false,"fromMe":false},{"ID":"CF58F16FBEBD42E333182F535AF22795","message":"La sesión con Oti tiene un valor de 400$ y Daniel 600$\\nDaniel tiene cupos disponibles para Julio y oti tiene cupos disponibles","date":"2024-06-11T09:25:58-04:00","wasReceipt":false,"wasRead":false,"fromMe":false},{"ID":"1EA63B76D69886BCDD560253C9C0BBD2","message":"El valor es por sesion","date":"2024-06-11T09:26:09-04:00","wasReceipt":false,"wasRead":false,"fromMe":false},{"ID":"7BF1C31BE40BD7DE5FDB44AB9067558F","message":"Aproximadamente cuantas sesiones tardarian para realizar esa pieza ?","date":"2024-06-11T10:00:48-04:00","wasReceipt":false,"wasRead":false,"fromMe":true}]}]');
      // setConversations(result);
      // setShowQR(false);
    },
    onMessage: (event: WebSocketEventMap['message']) => {
      console.log('-------------------')
      const data: WebsocketMessage = JSON.parse(event?.data);
      console.log(data)
      if(messageData !== data?.message)
        setMessageData(data.message);

      if(data.messageType === WebsocketMessageTypes.QR_CODE) {
        setQrString(data.message);
        setShowQR(true);
      }

      if(data.messageType === WebsocketMessageTypes.CONVERSATIONS_CODE) {
        const result: WhatsappConversation[] = JSON.parse(data.message);
        setConversations(result);
        setShowQR(false);
      }
    }
  });

  const selectConversation = (id: string): void => {
    const findConversation: WhatsappConversation | undefined = conversations?.find(x => x.id === id);
    if (findConversation)
      setSelectedConversation(findConversation);
  }

  const sendWhatsappMessage = (message: string): void => {
    console.log('Enviar primero el mensaje');
    console.log(lastMessage);

    const result: any = {
      type: 'send_message',
      payload: {
        message: message,
        from: 'Eduardo'
      }
    };
    sendMessage(JSON.stringify(result));
  }

  // const handleClickSendMessage = useCallback(() => sendMessage('Hello'), []);

  return (
    <SimpleCard>
      {showQR && (
        <ConnectQr qrString={qrString} />
      )}

      {!showQR &&
        conversations && (
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
        )}
    </SimpleCard>
  )
}

export default WhatsAppClient;
