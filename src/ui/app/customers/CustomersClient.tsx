'use client';

import SimpleCard from "../components/cards/SimpleCard";
import SimpleTable from "../components/tables/SimpleTable";
import { ColumnType, SimpleTableColumn } from "../components/tables/SimpleTable.types";
import BreadcrumbNavigation from "../components/BreadcrumbNavigation";
import { Button, Input } from "@chakra-ui/react";
import { Icon } from '@iconify/react';
import { Customer } from '../types/customer';
import Link from "next/link";
import { BreadcrumItem } from "../types";
import DeleteModal from "../components/modals/DeleteModal";
import useDeleteModal from "../hooks/useDeleteModal";

const CustomersClient = () => {
  const bcItems: BreadcrumItem[] = [
    {
      label: "Clientes",
      href: "/customers"
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

  const customerData: Customer[] = [
    {
      id: 1,
      firstName: "Eduardo",
      lastName: "Rodriguez",
      identificationNumber: "V-12123123",
      phone: "04161234567",
      email: "test@test.com"
    },
    {
      id: 2,
      firstName: "Eduardo",
      lastName: "Rodriguez",
      identificationNumber: "V-12123123",
      phone: "04161234567",
      email: "test@test.com"
    },
    {
      id: 3,
      firstName: "Eduardo",
      lastName: "Rodriguez",
      identificationNumber: "V-12123123",
      phone: "04161234567",
      email: "test@test.com"
    },
    {
      id: 4,
      firstName: "Eduardo",
      lastName: "Rodriguez",
      identificationNumber: "V-12123123",
      phone: "04161234567",
      email: "test@test.com"
    },
    {
      id: 5,
      firstName: "Eduardo",
      lastName: "Rodriguez",
      identificationNumber: "V-12123123",
      phone: "04161234567",
      email: "test@test.com"
    },
    {
      id: 6,
      firstName: "Eduardo",
      lastName: "Rodriguez",
      identificationNumber: "V-12123123",
      phone: "04161234567",
      email: "test@test.com"
    },
    {
      id: 7,
      firstName: "Eduardo",
      lastName: "Rodriguez",
      identificationNumber: "V-12123123",
      phone: "04161234567",
      email: "test@test.com"
    },
    {
      id: 8,
      firstName: "Eduardo",
      lastName: "Rodriguez",
      identificationNumber: "V-12123123",
      phone: "04161234567",
      email: "test@test.com"
    },
    {
      id: 9,
      firstName: "Eduardo",
      lastName: "Rodriguez",
      identificationNumber: "V-12123123",
      phone: "04161234567",
      email: "test@test.com"
    },
    {
      id: 10,
      firstName: "Eduardo",
      lastName: "Rodriguez",
      identificationNumber: "V-12123123",
      phone: "04161234567",
      email: "test@test.com"
    },
  ]

  const deleteCustomerModal = useDeleteModal();

  return (
    <div>
      <SimpleCard>
        <DeleteModal onSubmit={()=>{}} title="Eliminar Cliente" description="¿Estás seguro que quieres eliminar este cliente?"/>
        <BreadcrumbNavigation items={bcItems} />

        <hr className="my-3" />

        <div className="grid grid-cols-1 md:grid-cols-3 lg:grid-cols-5 xl:grid-cols-6 gap-4">
          <div>
            <label className="text-sm">Nombre</label>
            <Input size="sm" />
          </div>
          <div>
            <label className="text-sm">Apellido</label>
            <Input size="sm" />
          </div>
          <div>
            <label className="text-sm">Cédula</label>
            <Input size="sm" />
          </div>
          <div className="flex flex-col">
            <span className="opacity-0">.</span>
            <Button size="sm" variant="main">
              <Icon icon="tabler:search" />
            </Button>
          </div>
          <div className="xl:col-start-6">
            <div className="flex flex-col">
              <span className="opacity-0">.</span>
              <Link href="/customers/create">
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
          <SimpleTable columns={customerCols} data={customerData} showDetails showEdit showDelete onDelete={deleteCustomerModal.onOpen} />
        </SimpleCard>
      </div>
    </div>
  )
}

export default CustomersClient;
