'use client';

import { Icon } from "@iconify/react";
import QRCode from "react-qr-code";

interface ConnectQrProps {
  qrString: string;
}

const ConnectQr: React.FC<ConnectQrProps> = ({
  qrString
}) => {
  return (
    <>
      <div className='flex justify-center mt-5'>
        <div className='rounded-full bg-thirdcolorhov text-thirdcolor text-5xl p-2'>
          <Icon icon="ic:baseline-whatsapp" />
        </div>
      </div>
      <div className='text-center my-3'>
        <h1 className='font-bold text-lg'>No Conectado</h1>
        <span className='text-sm mt-2'>Parece que no estás conectado a Whatsapp</span><br />
        <span className='text-sm mt-2'>Escanea el código QR de abajo para conectarte</span>
      </div>

      <div className="flex justify-center my-5">
        <QRCode value={qrString} />
      </div>
    </>
  );
}

export default ConnectQr;
