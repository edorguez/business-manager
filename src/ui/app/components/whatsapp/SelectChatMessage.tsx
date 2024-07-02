'use client';

import { Icon } from "@iconify/react";

const SelectChatMessage = () => {
  return (
    <div className="bg-white rounded-lg p-2 h-[250px] w-[300px] flex flex-col justify-center">
      <div className='flex justify-center'>
        <div className='rounded-full bg-thirdcolorhov text-thirdcolor text-5xl p-2'>
          <Icon icon="ic:baseline-whatsapp" />
        </div>
      </div>
      <div className='text-center mt-3'>
        <h1 className='font-bold text-lg'>Seleccionar Chat</h1>
        <span className='text-sm'>Selecciona un chat para comenzar a enviar y recibir mensajes</span>
      </div>
    </div>
  )
}

export default SelectChatMessage;
