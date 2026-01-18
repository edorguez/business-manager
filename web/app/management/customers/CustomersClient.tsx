'use client';

import getCurrentUser from '@/app/actions/getCurrentUser';
import { DeleteCustomerRequest, GetCustomersRequest } from '@/app/services/customers';
import BreadcrumbNavigation from '@/app/components/BreadcrumbNavigation';
import SimpleCard from '@/app/components/cards/SimpleCard';
import SimpleTable from '@/app/components/tables/SimpleTable';
import { ColumnType, SimpleTableColumn } from '@/app/components/tables/SimpleTable.types';
import { BreadcrumItem } from '@/app/types';
import { CurrentUser } from '@/app/types/auth';
import { Customer, SearchCustomer } from '@/app/types/customer';
import { Button, Input, useToast } from '@chakra-ui/react';
import { Icon } from '@iconify/react';
import Link from "next/link";
import { useCallback, useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';
import useLoading from '@/app/hooks/useLoading';
import useWarningModal from '@/app/hooks/useWarningModal';
import WarningModal from '@/app/components/modals/WarningModal';
import { validLetters, validNumbers } from '@/app/utils/InputUtils';

const CustomersClient = () => {
  const bcItems: BreadcrumItem[] = [
    {
      label: "Clientes",
      href: "/management/customers"
    }
  ];

  const customerCols: SimpleTableColumn[] = [
    {
      key: "identificationNumber",
      name: "Cédula",
      type: ColumnType.String
    },
    {
      key: "firstName",
      name: "Nombre",
      type: ColumnType.String
    },
    {
      key: "lastName",
      name: "Apellido",
      type: ColumnType.String
    },
    {
      key: "email",
      name: "Correo",
      type: ColumnType.String
    },
    {
      key: "phone",
      name: "Teléfono",
      type: ColumnType.String
    },
  ]

  const isLoading = useLoading();
  const toast = useToast();
  const { push } = useRouter();
  const [searchCustomer, setSearchCustomer] = useState<SearchCustomer>({ name: '', lastName: '', identificationNumber: '' });
  const [customerData, setCustomerData] = useState<Customer[]>([]);
  const [offset, setOffset] = useState<number>(0);
  const deleteCustomerModal = useWarningModal();
  const [customerIdDelete, setCustomerIdDelete] = useState<number>(0);

  const getCustomers = useCallback(async (searchParams: SearchCustomer = searchCustomer) => {
    isLoading.onStartLoading();
    const currentUser: CurrentUser | null = getCurrentUser();

    if (currentUser) {
      let data: Customer[] = await GetCustomersRequest({ companyId: currentUser.companyId, limit: 10, offset: offset, name: searchParams?.name ?? "", lastName: searchParams?.lastName ?? "", identificationNumber: searchParams?.identificationNumber ?? "" });
      const formatData: Customer[] = data.map(x => {
        return {
          ...x,
          identificationNumber: `${x.identificationType}-${x.identificationNumber}`
        }
      })
      setCustomerData(formatData);
    }
    isLoading.onEndLoading();
  }, [offset])

  useEffect(() => {
    getCustomers();
  }, [getCustomers]);

  const handleChangePage = (val: string) => {
    setOffset((prevValue) => val === 'NEXT' ? prevValue += 10 : prevValue -= 10);
  }

  const handleNameChange = (event: any) => {
    const { name, value } = event.target;
    if(value && !validLetters(value, true)) return;
    setSearchCustomer((prevData) => ({ ...prevData, [name]: value }));
  }

  const handleIdentificationNumberChange = (event: any) => {
    const { name, value } = event.target;
    if(value && !validNumbers(value)) return;
    setSearchCustomer((prevData) => ({ ...prevData, [name]: value }));
  }

  const onSearchCustomer = () => {
    getCustomers(searchCustomer);
  }

  const handleOpenDelete = (val: any) => {
    setCustomerIdDelete(val.id);
    deleteCustomerModal.onOpen();
  }

  const handleSubmitDelete = () => {
    onDelete(customerIdDelete);
  }

  const onDelete = useCallback(async (id: number) => {
    await DeleteCustomerRequest({ id });
    getCustomers(searchCustomer);
    deleteCustomerModal.onClose();
    toast({
      title: 'Cliente',
      description: 'Cliente eliminado exitosamente',
      variant: 'customsuccess',
      position: 'top-right',
      duration: 3000,
      isClosable: true,
    });
  }, [])

  const handleOpenEdit = (val: any) => {
    push(`customers/${val.id}?isEdit=true`);
  }

  const handleOpenDetail = (val: any) => {
    push(`customers/${val.id}`);
  }

  return (
    <div>
      <SimpleCard>
        <WarningModal onSubmit={handleSubmitDelete} title="Eliminar Cliente" description="¿Estás seguro que quieres eliminar este cliente?" confirmText="Eliminar" />
        <BreadcrumbNavigation items={bcItems} />

        <hr className="my-3" />

        <div className="grid grid-cols-1 md:grid-cols-3 lg:grid-cols-5 xl:grid-cols-6 gap-4">
          <div>
            <label className="text-sm">Nombre</label>
            <Input size="sm" name="name" onChange={handleNameChange} value={searchCustomer.name} maxLength={20} />
          </div>
          <div>
            <label className="text-sm">Apellido</label>
            <Input size="sm" name="lastName" onChange={handleNameChange} value={searchCustomer.lastName} maxLength={20} />
          </div>
          <div>
            <label className="text-sm">Cédula</label>
            <Input size="sm" name="identificationNumber" onChange={handleIdentificationNumberChange} value={searchCustomer.identificationNumber} maxLength={9} />
          </div>
          <div className="flex flex-col">
            <span className="opacity-0">.</span>
            <Button size="sm" variant="main" onClick={onSearchCustomer}>
              <Icon icon="tabler:search" />
            </Button>
          </div>
          <div className="xl:col-start-6">
            <div className="flex flex-col">
              <span className="opacity-0">.</span>
              <Link href="/management/customers/create">
                <Button size="sm" variant="main" className="w-full">
                  Crear Cliente
                </Button>
              </Link>
            </div>
          </div>
        </div>
      </SimpleCard>

      <div className="mt-3">
        <SimpleCard>
          <SimpleTable columns={customerCols} data={customerData} showDetails showEdit showDelete onEdit={handleOpenEdit} onDelete={handleOpenDelete} onDetail={handleOpenDetail} onChangePage={handleChangePage} offset={offset} />
        </SimpleCard>
      </div>
    </div>
  )
}

export default CustomersClient;
