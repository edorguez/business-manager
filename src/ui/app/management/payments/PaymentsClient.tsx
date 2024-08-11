"use client";

import getCurrentUser from "@/app/actions/getCurrentUser";
import { DeletePaymentRequest, GetPaymentsRequest } from "@/app/api/payment/route";
import { GetPaymentTypesRequest } from "@/app/api/paymentType/route";
import BreadcrumbNavigation from "@/app/components/BreadcrumbNavigation";
import PaymentCard from "@/app/components/cards/PaymentCard";
import PaymentFilterCard from "@/app/components/cards/PaymentFilterCard";
import SimpleCard from "@/app/components/cards/SimpleCard";
import DeleteModal from "@/app/components/modals/DeleteModal";
import useDeleteModal from "@/app/hooks/useDeleteModal";
import useLoading from "@/app/hooks/useLoading";
import { BreadcrumItem } from "@/app/types";
import { CurrentUser } from "@/app/types/auth";
import { Payment } from "@/app/types/payment";
import { PaymentType } from "@/app/types/paymentType";
import { Button, useToast } from "@chakra-ui/react";
import Link from "next/link";
import { useCallback, useEffect, useState } from "react";

const PaymentsClient = () => {
  const bcItems: BreadcrumItem[] = [
    {
      label: "Métodos de Pago",
      href: "/management/payments",
    },
  ];

  const isLoading = useLoading();
  const toast = useToast();
  const deletePaymentModal = useDeleteModal();
  const [paymentTypes, setPaymentTypes] = useState<PaymentType[]>([]);
  const [payments, setPayments] = useState<Payment[]>([]);
  const [paymentIdDelete, setPaymentIdDelete] = useState<number>(0);

  const getPaymentTypes = useCallback(async () => {
    isLoading.onStartLoading();
    const pt = await GetPaymentTypesRequest();
    setPaymentTypes(pt);
    isLoading.onEndLoading();
  }, []);

  const getPayments = useCallback(async () => {
    isLoading.onStartLoading();
    const currentUser: CurrentUser | null = getCurrentUser();

    if (currentUser) {
      let data: Payment[] = await GetPaymentsRequest({
        companyId: currentUser.companyId,
        limit: 10,
        offset: 0,
      });
      setPayments(data);
    }
    isLoading.onEndLoading();
  }, []);

  useEffect(() => {
    getPaymentTypes();
    getPayments();
  }, []);
  
  const handleOpenDelete = (id: number) => {
    setPaymentIdDelete(id);
    deletePaymentModal.onOpen();
  }

  const handleSubmitDelete = () => {
    onDelete(paymentIdDelete);
  }
  
  const onDelete = useCallback(async (id: number) => {
    await DeletePaymentRequest({ id });
    getPayments();
    deletePaymentModal.onClose();
    toast({
      title: 'Método de Pago',
      description: 'Método de pago eliminado exitosamente',
      variant: 'customsuccess',
      position: 'top-right',
      duration: 3000,
      isClosable: true,
    });
  }, [])

  return (
    <div>
      <SimpleCard>
        <DeleteModal onSubmit={handleSubmitDelete} title="Eliminar Método de Pago" description="¿Estás seguro que quieres eliminar este método de pago?" />
        <BreadcrumbNavigation items={bcItems} />

        <hr className="my-3" />

        <div className="grid grid-cols-1 lg:grid-cols-5 gap-4">
          <div>
            <h1 className="ml-2 font-bold">Métodos de Pago</h1>
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
                  <PaymentFilterCard
                    isSelected={true}
                    onlyAll={true}
                    onSelectPayment={() => {}}
                  />
                }
                {paymentTypes.map((val: any, index: number) => (
                  <PaymentFilterCard
                    key={index}
                    paymentType={val}
                    isSelected={index == 0}
                    onlyAll={false}
                    onSelectPayment={() => {}}
                  />
                ))}
              </div>
            </div>
          </SimpleCard>
          <SimpleCard>
            <div className="p-2">
              <h3>Métodos de Pago</h3>
              {payments.map((item) => (
                <div key={item.id} className="mt-1">
                  <PaymentCard payment={item} onDelete={handleOpenDelete} />
                </div>
              ))}
            </div>
          </SimpleCard>
        </div>
      </div>
    </div>
  );
};

export default PaymentsClient;
