'use client';

import { Alert, AlertIcon, Button } from "@chakra-ui/react";
import { Icon } from "@iconify/react";

interface ConnectedProps {
  phone: string
}

const Connected: React.FC<ConnectedProps> = ({
  phone
}) => {

  const formatPhone = () => {
    const rgx = /^\d+/g;
    return '+' + rgx.exec(phone);
  }

  return (
    <>
      {/*
      <div className='flex justify-center mt-5'>
        <div className='rounded-full bg-maincolorhov text-maincolor text-9xl p-2'>
          <Icon icon="ic:twotone-check-circle" />
        </div>
      </div>
      <div className='text-center my-3'>
        <h1 className='font-bold text-lg'>Conectado en WhatsApp</h1>
        <h5 className="font-medium text-maincolor">{formatPhone()}</h5>
      </div>
      <div className="text-center my-3">
        <span className='text-sm mt-2'>Ya puedes recibir órdenes a través de Whatsapp, en caso que quieras cambiar de dispositivo solo haz clic en <b>&quot;Desconectar&quot;</b></span>
      </div>
      <div className="mt-9 flex justify-center">
        <Button variant="third" size="sm">Desconectar</Button>
      </div>
      */}
      <Alert status='success' size='sm'>
        <AlertIcon />
        <div className="w-full flex justify-between items-center">
          <span className="text-sm">
            Conectado en WhatsApp {formatPhone()}
          </span>
          <Button variant="third" size="sm">Desconectar</Button>
        </div>
      </Alert>
    </>
  );
}

export default Connected;
