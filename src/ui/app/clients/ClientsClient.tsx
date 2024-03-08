'use client';

import SimpleCard from "../components/cards/SimpleCard";
import BreadcrumbNavigation from "../components/BreadcrumbNavigation";
import { Button, Input } from "@chakra-ui/react";
import { Icon } from '@iconify/react';

const ClientsClient = () => {
  const bcItems: string[] = ["Clientes", "Crear"];

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

        <hr className="my-3"/>
      </SimpleCard>
    </div>
  )
}

export default ClientsClient;
