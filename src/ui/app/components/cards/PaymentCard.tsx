'use client';

import React from "react";
import { Icon } from "@iconify/react";
import Image from 'next/image';
import { Button, Switch } from "@chakra-ui/react";
import { Payment } from "@/app/types/payment";

interface PaymentCardProps {
  payment: Payment;
  onDelete?: (id: number) => void;
  onChangeStatus?: (id: number, status: boolean) => void;
}

const PaymentCard: React.FC<PaymentCardProps> = ({
  payment,
  onDelete,
  onChangeStatus
}) => {

  const handleDelete = (id: number) => {
    if(onDelete)
      onDelete(id);
  }
  
  const handleChangeStatus = (id: number, status: boolean) => {
    if(onChangeStatus)
      onChangeStatus(id, status);
  }

  return (
    <div className="bg-white hover:bg-thirdcolorhov transition duration-150 rounded-md border border-slate-200 p-2">
      <div className="grid grid-cols-1 lg:grid-cols-2 items-center">
        <div className="flex items-center">
          <div className="rounded-md border border-slate-200 p-1 bg-white">
            <Image src={`/images/payments/${payment.paymentType.imagePath}`} alt="Logo" width={28} height={28} />
          </div>
          <span className="text-sm font-bold ml-2">{payment.name}</span>
        </div>
        <div className="flex items-center justify-end">
          <span className="text-xs font-bold text-maincolor mr-2">Activo</span>
          <Switch size='md' colorScheme='main' isChecked={payment.isActive} onChange={() => handleChangeStatus(payment.id, !payment.isActive)}/>
          <div className="flex ml-4">
            <Button size="sm" variant="fifth">
              <Icon icon="lucide:info" />
            </Button>

            <Button size="sm" variant="main" className="mx-1">
              <Icon icon="lucide:edit" />
            </Button>

            <Button size="sm" variant="third" onClick={() => handleDelete(payment.id)}>
              <Icon icon="wpf:delete" />
            </Button>
          </div>
        </div>
      </div>
    </div>
  )
}

export default PaymentCard;
