'use client';

import BreadcrumbNavigation from "@/app/components/BreadcrumbNavigation";
import SimpleCard from "@/app/components/cards/SimpleCard";
import { BreadcrumItem } from "@/app/types";
import Link from "next/link";
import { Icon } from '@iconify/react';
import { Button, Input, NumberInput, NumberInputField } from '@chakra-ui/react'
import ImagesUpload from "@/app/components/uploads/ImagesUpload";

const CreateProductClient = () => {
  const bcItems: BreadcrumItem[] = [
    {
      label: "Productos",
      href: "/management/products"
    },
    {
      label: "Crear Producto",
      href: "/management/products/create"
    }
  ];

  return (
    <div>
      <SimpleCard>
        <BreadcrumbNavigation items={bcItems} />
        <hr className='my-3' />
        <div className='flex items-center'>
          <div>
            <Link href="/management/products">
              <div className='rounded p-2 hover:bg-thirdcolor hover:text-white duration-150'>
                <Icon icon="fa-solid:arrow-left" />
              </div>
            </Link>
          </div>
          <h1 className='ml-2 font-bold'>Crear Producto</h1>
        </div>
      </SimpleCard>

      <div className='mt-3'>
        <SimpleCard>
          <div className='mt-2'>
            <label className='text-sm'>Nombre <span className='text-thirdcolor'>*</span></label>
            <Input size="sm" />
          </div>
          <div className='mt-2'>
            <label className='text-sm'>SKU</label>
            <Input size="sm" />
          </div>
          <div className='mt-2'>
            <label className='text-sm'>Cantidad <span className='text-thirdcolor'>*</span></label>
            <Input size="sm" />
          </div>
          <div className='mt-2'>
            <label className='text-sm'>Precio <span className='text-thirdcolor'>*</span></label>
            <NumberInput size="sm" precision={2}>
              <NumberInputField />
            </NumberInput>
          </div>
        </SimpleCard>
      </div>
      
      <div className="mt-3">
        <SimpleCard>
          <div className="p-1">
            <label className='text-sm'>Imágenes</label>
            <div className="border rounded py-5 px-3">
              <ImagesUpload />
            </div>
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

export default CreateProductClient;
