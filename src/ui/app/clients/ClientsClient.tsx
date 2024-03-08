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
      </SimpleCard>

      <div className="mt-4">

        <SimpleCard>
          <div className="grid grid-cols-5 gap-4 mb-3">
            <div>
              <span className="text-sm">Nombre</span>
              <Input size="sm" />
            </div>
            <div>
              <span className="text-sm">Apellido</span>
              <Input size="sm" />
            </div>
            <div>
              <span className="text-sm">Cédula</span>
              <Input size="sm" />
            </div>
            <div>
              <Button size="sm" className="bg-maincolor">
                <Icon icon="tabler:search" />
              </Button>
            </div>
          </div>

          <hr />
        </SimpleCard>
      </div>
    </div>
  )
}

export default ClientsClient;
