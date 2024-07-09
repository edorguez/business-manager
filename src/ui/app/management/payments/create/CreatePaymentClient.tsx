'use client';

import { Button, Input, Select } from '@chakra-ui/react'
import { BreadcrumItem } from '@/app/types';
import { Icon } from '@iconify/react';
import Link from 'next/link';
import SimpleCard from '@/app/components/cards/SimpleCard';
import BreadcrumbNavigation from '@/app/components/BreadcrumbNavigation';
import PaymentFilterCard from '@/app/components/cards/PaymentFilterCard';

const CreatePaymentClient = () => {
  const bcItems: BreadcrumItem[] = [
    {
      label: "Métodos de Pago",
      href: "/management/payments"
    },
    {
      label: "Crear Método de Pago",
      href: "/management/payments/create"
    }
  ];

  const payments: any[] = [...Array(11)];

  return (
    <div>
      <SimpleCard>
        <BreadcrumbNavigation items={bcItems} />
        <hr className='my-3' />
        <div className='flex items-center'>
          <div>
            <Link href="/management/payments">
              <div className='rounded p-2 hover:bg-thirdcolor hover:text-white duration-150'>
                <Icon icon="fa-solid:arrow-left" />
              </div>
            </Link>
          </div>
          <h1 className='ml-2 font-bold'>Crear Método de Pago</h1>
        </div>
      </SimpleCard>

      <div className="mt-3">
        <SimpleCard>
          <div className="p-2">
            <label className='text-sm'>Tipo de Cuenta</label>
            <div className="grid grid-cols-2 sm:grid-cols-3 lg:grid-cols-4 gap-3 mt-1">
              {
                payments.map((val: any, index: number) => (
                  <PaymentFilterCard key={index} paymentTypeEnum={index} description="Description" isSelected={index == 0} />
                ))
              }
            </div>
          </div>
        </SimpleCard>
      </div>

      <div className='mt-3'>
        <SimpleCard>
          <div className='mt-2'>
            <label className='text-sm'>Nombre <span className='text-thirdcolor'>*</span></label>
            <Input size="sm" />
          </div>
          <div className='mt-2'>
            <label className='text-sm'>Banco</label>
            <Input size="sm" />
          </div>
          <div className='mt-2'>
            <label className='text-sm'>Número de Cuenta</label>
            <Input size="sm" />
          </div>
          <div className='mt-2'>
            <label className='text-sm'>Tipo de Cuenta</label>
            <Input size="sm" />
          </div>
          <div className='mt-2'>
            <label className='text-sm'>Cédula</label>
            <div className='flex'>
              <div className='w-24 mr-1'>
                <Select size='sm'>
                  <option value=''>-</option>
                  <option value='V'>V</option>
                  <option value='E'>E</option>
                  <option value='P'>P</option>
                  <option value='J'>J</option>
                  <option value='G'>G</option>
                </Select>
              </div>
              <Input size="sm" />
            </div>
          </div>
          <div className='mt-2'>
            <label className='text-sm'>Teléfono</label>
            <Input size="sm" />
          </div>
          <div className='mt-2'>
            <label className='text-sm'>Correo</label>
            <Input size="sm" />
          </div>
        </SimpleCard>
      </div>

      <div className="mt-3">
        <Button variant="main" className='w-full'>
          Crear
        </Button>
      </div>
    </div>
  )
}

export default CreatePaymentClient;
