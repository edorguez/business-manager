'use client';

import { Button, Input, Select, useToast } from '@chakra-ui/react'
import { BreadcrumItem } from '@/app/types';
import { Icon } from '@iconify/react';
import Link from 'next/link';
import SimpleCard from '@/app/components/cards/SimpleCard';
import BreadcrumbNavigation from '@/app/components/BreadcrumbNavigation';
import PaymentFilterCard from '@/app/components/cards/PaymentFilterCard';
import useLoading from '@/app/hooks/useLoading';
import { useCallback, useEffect, useState } from 'react';
import { PaymentType } from '@/app/types/paymentType';
import { GetPaymentTypesRequest } from '@/app/services/paymentType';
import { useRouter } from 'next/navigation';
import { CreatePayment } from '@/app/types/payment';
import { CreatePaymentRequest } from '@/app/services/payment';
import { CurrentUser } from '@/app/types/auth';
import getCurrentUser from '@/app/actions/getCurrentUser';

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

  const isLoading = useLoading();
  const [paymentTypes, setPaymentTypes] = useState<PaymentType[]>([]);
  const toast = useToast();
  const { push } = useRouter();
  const [formData, setFormData] = useState<CreatePayment>({
    companyId: 0,
    name: '',
    bank: '',
    accountNumber: '',
    accountType: '',
    identificationNumber: '',
    identificationType: '',
    phone: '',
    email: '',
    paymentTypeId: 0
  });

  const getPaymentTypes = useCallback(async () => {
    isLoading.onStartLoading();
    const pt = await GetPaymentTypesRequest();
    setPaymentTypes(pt);
    isLoading.onEndLoading();
  }, []);

  useEffect(() => {
    getPaymentTypes();
    
    const currentUser: CurrentUser | null = getCurrentUser();
    if (currentUser) {
      formData.companyId = currentUser.companyId;
    }
  }, []);
  
  const handleChange = (event: any) => {
    const { name, value } = event.target;
    setFormData((prevFormData) => ({ ...prevFormData, [name]: value }));
  }

  const handlePaymentTypeSelected = (val: number) => {
    setFormData((prevFormData) => ({ ...prevFormData, paymentTypeId: val }));
  }
  
  const onSubmit = async () => {
    if (isFormValid()) {
      isLoading.onStartLoading();
      let createCustomer: any = await CreatePaymentRequest(formData);
      if (!createCustomer.error) {
        isLoading.onEndLoading();
        showSuccessCreationMessage('Método de Pago creado exitosamente');
        push('/management/payments');
      } else {
        showErrorMessage(createCustomer.error);
        isLoading.onEndLoading();
      }
    } else {
      showErrorMessage('Algunos campos son requeridos o inválidos');
    }
  }

  const isFormValid = (): boolean => {
    if (!formData.name)
      return false;

    if (!formData.paymentTypeId)
      return false;

    return true;
  }

  const showSuccessCreationMessage = (msg: string) => {
    toast({
      title: 'Método de Pago',
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
                paymentTypes.map((val: any, index: number) => (
                  <PaymentFilterCard key={index} paymentType={val} isSelected={formData.paymentTypeId == val.id} onlyAll={false} onSelectPayment={handlePaymentTypeSelected} />
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
            <Input size="sm" name='name' value={formData.name} onChange={handleChange} maxLength={50} />
          </div>
          <div className='mt-2'>
            <label className='text-sm'>Banco</label>
            <Input size="sm" name='bank' value={formData.bank} onChange={handleChange} maxLength={50} />
          </div>
          <div className='mt-2'>
            <label className='text-sm'>Número de Cuenta</label>
            <Input size="sm" name='accountNumber' value={formData.accountNumber} onChange={handleChange} maxLength={20} />
          </div>
          <div className='mt-2'>
            <label className='text-sm'>Tipo de Cuenta</label>
            <Input size="sm" name='accountType' value={formData.accountType} onChange={handleChange} maxLength={20} />
          </div>
          <div className='mt-2'>
            <label className='text-sm'>Cédula</label>
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
          <div className='mt-2'>
            <label className='text-sm'>Teléfono</label>
            <Input size="sm" name='phone' value={formData.phone} onChange={handleChange} maxLength={11} />
          </div>
          <div className='mt-2'>
            <label className='text-sm'>Correo</label>
            <Input size="sm" name='email' value={formData.email} onChange={handleChange} maxLength={100} />
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

export default CreatePaymentClient;
