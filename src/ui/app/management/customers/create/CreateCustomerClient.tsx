'use client';

import { Button, Input, Select, useToast } from '@chakra-ui/react'
import { BreadcrumItem } from '@/app/types';
import { Icon } from '@iconify/react';
import Link from 'next/link';
import SimpleCard from '@/app/components/cards/SimpleCard';
import BreadcrumbNavigation from '@/app/components/BreadcrumbNavigation';
import { useEffect, useState } from 'react';
import { CreateCustomer } from '@/app/types/customer';
import { useRouter } from 'next/navigation';
import { CreateCustomerRequest } from '@/app/services/customers';
import { CurrentUser } from '@/app/types/auth';
import getCurrentUser from '@/app/actions/getCurrentUser';
import useLoading from '@/app/hooks/useLoading';

const CreateCustomerClient = () => {
  const bcItems: BreadcrumItem[] = [
    {
      label: "Clientes",
      href: "/management/customers"
    },
    {
      label: "Crear Cliente",
      href: "/management/customers/create"
    }
  ];

  const isLoading = useLoading();
  const toast = useToast();
  const { push } = useRouter();
  const [formData, setFormData] = useState<CreateCustomer>({
    companyId: 0,
    firstName: '',
    lastName: '',
    email: '',
    identificationNumber: '',
    identificationType: '',
    phone: ''
  });

  useEffect(() => {
    const currentUser: CurrentUser | null = getCurrentUser();
    if (currentUser) {
      formData.companyId = currentUser.companyId;
    }
  }, []);

  const handleChange = (event: any) => {
    const { name, value } = event.target;
    setFormData((prevFormData) => ({ ...prevFormData, [name]: value }));
  }

  const onSubmit = async () => {
    if (isFormValid()) {
      isLoading.onStartLoading();
      let createCustomer: any = await CreateCustomerRequest(formData);
      if (!createCustomer.error) {
        isLoading.onEndLoading();
        showSuccessCreationMessage('Cliente creado exitosamente');
        push('/management/customers');
      } else {
        showErrorMessage(createCustomer.error);
        isLoading.onEndLoading();
      }
    } else {
      showErrorMessage('Algunos campos son requeridos o inválidos');
    }
  }

  const isFormValid = (): boolean => {
    if (!formData.firstName)
      return false;

    if (!formData.identificationNumber || !formData.identificationType)
      return false;

    return true;
  }

  const showSuccessCreationMessage = (msg: string) => {
    toast({
      title: 'Cliente',
      description: msg,
      variant: 'customsuccess',
      position: 'top-right',
      duration: 3000,
      isClosable: true,
    });
  }

  const showErrorMessage = (msg: string) => {
    toast({
      title: 'Error',
      description: msg,
      variant: 'customerror',
      position: 'top-right',
      duration: 3000,
      isClosable: true,
    });
  }

  return (
    <div>
      <SimpleCard>
        <BreadcrumbNavigation items={bcItems} />
        <hr className='my-3' />
        <div className='flex items-center'>
          <div>
            <Link href="/management/customers">
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
            <Input size="sm" name='firstName' value={formData.firstName} onChange={handleChange} maxLength={20} />
          </div>
          <div className='mt-2'>
            <label className='text-sm'>Apellido</label>
            <Input size="sm" name='lastName' value={formData.lastName} onChange={handleChange} maxLength={20} />
          </div>
          <div className='mt-2'>
            <label className='text-sm'>Correo</label>
            <Input size="sm" name='email' value={formData.email} onChange={handleChange} maxLength={100} />
          </div>
          <div className='mt-2'>
            <label className='text-sm'>Teléfono</label>
            <Input size="sm" name='phone' value={formData.phone} onChange={handleChange} maxLength={11} />
          </div>
          <div className='mt-2'>
            <label className='text-sm'>Cédula <span className='text-thirdcolor'>*</span></label>
            <div className='flex'>
              <div className='w-24 mr-1'>
                <Select size='sm' name='identificationType' value={formData.identificationType} onChange={handleChange}>
                  <option value=''>-</option>
                  <option value='V'>V</option>
                  <option value='E'>E</option>
                  <option value='P'>P</option>
                  <option value='J'>J</option>
                  <option value='G'>G</option>
                </Select>
              </div>
              <Input size="sm" name='identificationNumber' value={formData.identificationNumber} onChange={handleChange} maxLength={20} />
            </div>
          </div>
        </SimpleCard>
      </div>

      <div className="mt-3">
        <Button variant="main" className='w-full' onClick={onSubmit}>
          Crear
        </Button>
      </div>
    </div>
  )
}

export default CreateCustomerClient;
