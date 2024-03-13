'use client';

import SimpleCard from '../../components/cards/SimpleCard';
import { Button, Input, Select } from '@chakra-ui/react'
import BreadcrumbNavigation from "../../components/BreadcrumbNavigation";
import { BreadcrumItem } from '@/app/types';
import { Icon } from '@iconify/react';
import Link from 'next/link';

const CreateCustomerClient = () => {
  const bcItems: BreadcrumItem[] = [
    {
      label: "Clientes",
      href: "/customers"
    },
    {
      label: "Crear Cliente",
      href: "/customers/create"
    }
  ];

  return (
    <div>
      <SimpleCard>
        <BreadcrumbNavigation items={bcItems} />
        <hr className='my-3' />
        <div className='flex items-center'>
          <div>
            <Link href="/customers">
              <div className='rounded p-2 hover:bg-thirdcolor hover:text-white duration-150'>
                <Icon icon="fa-solid:arrow-left" />
              </div>
            </Link>
          </div>
          <h1 className='ml-2 font-bold'>Crear Cliente</h1>
        </div>
      </SimpleCard>

      <div className='mt-3'>
        <SimpleCard>
          <div className='mt-2'>
            <label className='text-sm'>Nombre <span className='text-thirdcolor'>*</span></label>
            <Input size="sm" />
          </div>
          <div className='mt-2'>
            <label className='text-sm'>Apellido</label>
            <Input size="sm" />
          </div>
          <div className='mt-2'>
            <label className='text-sm'>Correo</label>
            <Input size="sm" />
          </div>
          <div className='mt-2'>
            <label className='text-sm'>Teléfono</label>
            <Input size="sm" />
          </div>
          <div className='mt-2'>
            <label className='text-sm'>Cédula <span className='text-thirdcolor'>*</span></label>
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
        </SimpleCard>

        <div className="mt-3">
          <Button variant="main" className='w-full'>
            Crear
          </Button>
        </div>
      </div>
    </div>
  )
}

export default CreateCustomerClient;
