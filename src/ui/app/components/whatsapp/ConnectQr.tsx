'use client';

import { Icon } from "@iconify/react";

const ConnectQr = () => {
  return (
    <>
      <div className='flex justify-center mt-5'>
        <div className='rounded-full bg-thirdcolorhov text-thirdcolor text-5xl p-2'>
          <Icon icon="gravity-ui:plug-connection" />
        </div>
      </div>
      <div className='text-center my-3'>
        <h1 className='font-bold text-lg'>No Conectado</h1>
        <span className='text-sm mt-2'>Parece que no estás conectado a Whatsapp</span><br />
        <span className='text-sm mt-2'>Escanea el código QR de abajo para conectarte</span>
      </div>
    </>
  );
}

export default ConnectQr;
