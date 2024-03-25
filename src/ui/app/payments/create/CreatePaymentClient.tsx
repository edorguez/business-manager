'use client';

import SimpleCard from '../../components/cards/SimpleCard';
import BreadcrumbNavigation from "../../components/BreadcrumbNavigation";
import { BreadcrumItem } from "../../types";

const CreatePaymentClient = () => {
  const bcItems: BreadcrumItem[] = [
    {
      label: "Métodos de Pago",
      href: "/payments"
    },
    {
      label: "Crear Métodos",
      href: "/payments/create"
    }
  ];

  return (
    <SimpleCard>
      <BreadcrumbNavigation items={bcItems} />
    </SimpleCard>
  )
}

export default CreatePaymentClient;
