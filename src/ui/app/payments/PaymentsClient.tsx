'use client';

import SimpleCard from "../components/cards/SimpleCard";
import SimpleCardItem from "../components/cards/SimpleCardItem";
import BreadcrumbNavigation from "../components/BreadcrumbNavigation";
import { BreadcrumItem } from "../types";
import { Button } from "@chakra-ui/react";
import Link from "next/link";

const PaymentsClient = () => {

  const bcItems: BreadcrumItem[] = [
    {
      label: "Métodos de Pago",
      href: "/payments"
    }
  ];

  return (
    <div>
      <SimpleCard>
        <BreadcrumbNavigation items={bcItems} />

        <hr className="my-3" />

        <div className="grid grid-cols-1 lg:grid-cols-5 gap-4">
          <div>
            <h1 className='ml-2 font-bold'>Métodos de Pago</h1>
          </div>
          <div className="lg:col-start-5">
            <Link href="/payments/create">
              <Button size="sm" variant="main" className="w-full">
                Crear Método de Pago
              </Button>
            </Link>
          </div>
        </div>
      </SimpleCard>

      <div className="mt-3">
        <SimpleCard>
          <div className="grid grid-cols-1 xl:grid-cols-2">
            <div>
              <h3>Métodos de Pago</h3>
              <SimpleCardItem>
                <span>hola</span>
              </SimpleCardItem>
            </div>
            <div>
              <h3>Métodos de Pago</h3>
              <SimpleCardItem>
                <span>hola</span>
              </SimpleCardItem>
              <SimpleCardItem>
                <span>hola</span>
              </SimpleCardItem>
              <SimpleCardItem>
                <span>hola</span>
              </SimpleCardItem>
              <SimpleCardItem>
                <span>hola</span>
              </SimpleCardItem>
            </div>
          </div>
        </SimpleCard>
      </div>
    </div>
  )

}

export default PaymentsClient;
