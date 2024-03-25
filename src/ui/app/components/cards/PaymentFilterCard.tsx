'use client';

import React from "react";
import Image from 'next/image';
import { PaymentTypeEnum } from "@/app/types/payment";

interface PaymentFilterCardProps {
  description: string;
  paymentTypeEnum: PaymentTypeEnum;
  isSelected: boolean;
}

const PaymentFilterCard: React.FC<PaymentFilterCardProps> = ({
  description,
  paymentTypeEnum,
  isSelected
}) => {
  const paymentType: string = PaymentTypeEnum[paymentTypeEnum];

  return (
    <div className={`
      shadow-md 
      cursor-pointer 
      bg-white 
      hover:bg-thirdcolorhov 
      hover:scale-105 
      transition 
      duration-150
      rounded-md 
      border 
      border-slate-200 
      p-2
      ${isSelected && 'border-2 border-maincolor bg-maincolorhov hover:bg-maincolorhov'}
    `}>
      <div className="flex flex-col items-center select-none">
        <div className="mt-2 rounded-md border border-slate-200 p-1 bg-white inline-flex">
          <Image src={`/images/payments/${paymentType}.png`} alt="Logo" width={64} height={64} />
        </div>
        <span className="font-bold mt-4">{paymentType.charAt(0).toUpperCase() + paymentType.slice(1)}</span>
        <span className="text-xs">{description}</span>
      </div>
    </div>
  )
}

export default PaymentFilterCard;
