"use client";

import React from "react";
import Image from "next/image";
import { PaymentType } from "@/app/types/paymentType";

interface PaymentFilterCardProps {
  paymentType?: PaymentType;
  isSelected: boolean;
  onlyAll: boolean;
  onSelectPayment?: (val: number) => void;
}

const PaymentFilterCard: React.FC<PaymentFilterCardProps> = ({
  paymentType,
  isSelected,
  onlyAll,
  onSelectPayment,
}) => {
  const handleSelectPayment = () => {
    if (paymentType && onSelectPayment) onSelectPayment(paymentType.id);
    else if (onlyAll && onSelectPayment) onSelectPayment(0);
  };

  return (
    <div
      className={`
      shadow-md
      cursor-pointer
      hover:scale-105
      transition
      duration-150
      rounded-md
      p-2
      ${
        isSelected
          ? "border-2 border-maincolor bg-maincolorhov hover:bg-maincolorhov"
          : "border border-slate-200 bg-white hover:bg-thirdcolorhov"
      }
    `}
      onClick={handleSelectPayment}
    >
      {onlyAll && (
        <div className="flex flex-col items-center select-none">
          <div className="mt-2 rounded-md border border-slate-200 p-1 bg-white inline-flex">
            <Image
              src={`/images/payments/all.png`}
              alt="Logo"
              width={64}
              height={64}
            />
          </div>
          <span className="font-bold mt-4">Todos</span>
        </div>
      )}

      {!onlyAll && paymentType && (
        <div className="flex flex-col items-center select-none">
          <div className="mt-2 rounded-md border border-slate-200 p-1 bg-white inline-flex">
            <Image
              src={`/images/payments/${paymentType.imagePath}`}
              alt="Logo"
              width={64}
              height={64}
            />
          </div>
          <span className="font-bold mt-4">
            {paymentType.name.charAt(0).toUpperCase() +
              paymentType.name.slice(1)}
          </span>
        </div>
      )}
    </div>
  );
};

export default PaymentFilterCard;
