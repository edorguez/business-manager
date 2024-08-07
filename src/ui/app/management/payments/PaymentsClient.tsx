'use client';

import { GetPaymentTypesRequest } from "@/app/api/paymentType/route";
import BreadcrumbNavigation from "@/app/components/BreadcrumbNavigation";
import PaymentCard from "@/app/components/cards/PaymentCard";
import PaymentFilterCard from "@/app/components/cards/PaymentFilterCard";
import SimpleCard from "@/app/components/cards/SimpleCard";
import useLoading from "@/app/hooks/useLoading";
import { BreadcrumItem } from "@/app/types";
import { PaymentType } from "@/app/types/paymentType";
import { Button } from "@chakra-ui/react";
import Link from "next/link";
import { useCallback, useEffect, useState } from "react";

const PaymentsClient = () => {

  const bcItems: BreadcrumItem[] = [
    {
      label: "Métodos de Pago",
      href: "/management/payments"
    }
  ];

  const isLoading = useLoading();
  const [paymentTypes, setPaymentTypes] = useState<PaymentType[]>([]);

  const getPaymentTypes = useCallback(async () => {
    isLoading.onStartLoading();
    const pt = await GetPaymentTypesRequest();
    setPaymentTypes(pt);
    isLoading.onEndLoading();
  }, []);

  useEffect(() => {
    getPaymentTypes();
  }, []);


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
            <Link href="/management/payments/create">
              <Button size="sm" variant="main" className="w-full">
                Crear Método de Pago
              </Button>
            </Link>
          </div>
        </div>
      </SimpleCard>

      <div className="mt-3">
        <div className="grid grid-cols-1 xl:grid-cols-2 gap-3">
          <SimpleCard>
            <div className="p-2">
              <h3>Filtrar</h3>
              <div className="grid grid-cols-2 sm:grid-cols-3 lg:grid-cols-4 gap-3 mt-1">
                {
                  <PaymentFilterCard isSelected={true} onlyAll={true} onSelectPayment={() => {}} />
                }
                {
                  paymentTypes.map((val: any, index: number) => (
                    <PaymentFilterCard key={index} paymentType={val} isSelected={index == 0} onlyAll={false} onSelectPayment={() => {}} />
                  ))
                }
              </div>
            </div>
          </SimpleCard>
          <SimpleCard>
            <div className="p-2">
              <h3>Métodos de Pago</h3>
              {/* <div className="mt-1">
                <PaymentCard name="Transferencia" paymentTypeEnum={0} />
              </div> */}
              
            </div>
          </SimpleCard>
        </div>
      </div>
    </div>
  )

}

export default PaymentsClient;
