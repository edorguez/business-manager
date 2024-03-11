'use client';

import SimpleCard from "../components/cards/SimpleCard";
import SimpleTable from "../components/tables/SimpleTable";
import { SimpleTableColumn } from "../components/tables/SimpleTable.types";
import BreadcrumbNavigation from "../components/BreadcrumbNavigation";
import { Button, Input } from "@chakra-ui/react";
import { Icon } from '@iconify/react';
import { Customer } from '../types/customer';

const CustomersClient = () => {
  const bcItems: string[] = ["Clientes", "Crear"];

  const customerCols: SimpleTableColumn[] = [
    {
      key: "firstName",
      name: "Nombre"
    },
    {
      key: "lastName",
      name: "Apellido"
    },
    {
      key: "identificationNumber",
      name: "Cédula"
    },
    {
      key: "email",
      name: "Correo"
    },
    {
      key: "phone",
      name: "Teléfono"
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
  ]

  return (
    <div>
      <SimpleCard>
        <BreadcrumbNavigation items={bcItems} />

        <hr className="my-3" />

        <div className="grid grid-cols-6 gap-4">
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
          <div className="col-start-6">
            <div className="flex flex-col">
              <span className="opacity-0">.</span>
              <Button size="sm" variant="main">
                Crear Cliente
              </Button>
            </div>
          </div>
        </div>
      </SimpleCard>

      <div className="mt-3">
        <SimpleCard>
          <SimpleTable columns={customerCols} data={customerData} showEdit showDelete />
        </SimpleCard>
      </div>
    </div>
  )
}

export default CustomersClient;
