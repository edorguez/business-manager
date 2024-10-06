"use client";

import getCurrentUser from "@/app/actions/getCurrentUser";
import {
  ChangeStatusRequest,
  DeletePaymentRequest,
  GetPaymentsRequest,
} from "@/app/api/payment/route";
import { GetPaymentTypesRequest } from "@/app/api/paymentType/route";
import BreadcrumbNavigation from "@/app/components/BreadcrumbNavigation";
import PaymentCard from "@/app/components/cards/PaymentCard";
import PaymentFilterCard from "@/app/components/cards/PaymentFilterCard";
import SimpleCard from "@/app/components/cards/SimpleCard";
import WarningModal from "@/app/components/modals/WarningModal";
import { PAYMENT, PLAN_ID } from "@/app/constants";
import useLoading from "@/app/hooks/useLoading";
import useWarningModal from "@/app/hooks/useWarningModal";
import { BreadcrumItem } from "@/app/types";
import { CurrentUser } from "@/app/types/auth";
import { Payment } from "@/app/types/payment";
import { PaymentType } from "@/app/types/paymentType";
import { Button, useToast } from "@chakra-ui/react";
import Link from "next/link";
import { useRouter } from "next/navigation";
import { useCallback, useEffect, useState } from "react";

const PaymentsClient = () => {
  const bcItems: BreadcrumItem[] = [
    {
      label: "Métodos de Pago",
      href: "/management/payments",
    },
  ];

  const { push } = useRouter();
  const isLoading = useLoading();
  const toast = useToast();
  const deletePaymentModal = useWarningModal();
  const [paymentTypes, setPaymentTypes] = useState<PaymentType[]>([]);
  const [payments, setPayments] = useState<Payment[]>([]);
  const [paymentIdDelete, setPaymentIdDelete] = useState<number>(0);
  const [filterPaymentIdSelected, setFilterPaymentIdSelected] =
    useState<number>(0);
  const [showCreateBtn, setShowCreateBtn] = useState<boolean>(false);

  const getPaymentTypes = useCallback(async () => {
    isLoading.onStartLoading();
    const pt = await GetPaymentTypesRequest();
    setPaymentTypes(pt);
    isLoading.onEndLoading();
  }, []);

  const getPayments = useCallback(async (paymentTypeId: number) => {
    isLoading.onStartLoading();
    const currentUser: CurrentUser | null = getCurrentUser();

    if (currentUser) {
      let data: Payment[] = await GetPaymentsRequest({
        companyId: currentUser.companyId,
        paymentTypeId: paymentTypeId,
        limit: 30,
        offset: 0,
      });
      setShowCreateBtn(
        currentUser.planId === PLAN_ID.PRO ||
          data.length < PAYMENT.MAX_BASIC_PLAN_ITEMS
      );
      setPayments(data);
    }
    isLoading.onEndLoading();
  }, []);

  useEffect(() => {
    getPaymentTypes();
    getPayments(filterPaymentIdSelected);
  }, []);

  const handleOpenDelete = (id: number) => {
    setPaymentIdDelete(id);
    deletePaymentModal.onOpen();
  };

  const handleSubmitDelete = () => {
    onDelete(paymentIdDelete);
  };

  const onDelete = useCallback(async (id: number) => {
    isLoading.onStartLoading();
    await DeletePaymentRequest({ id });
    getPayments(filterPaymentIdSelected);
    deletePaymentModal.onClose();
    isLoading.onEndLoading();
    toast({
      title: "Método de Pago",
      description: "Método de pago eliminado exitosamente",
      variant: "customsuccess",
      position: "top-right",
      duration: 3000,
      isClosable: true,
    });
  }, []);

  const onChangeStatus = useCallback(async (id: number, status: boolean) => {
    isLoading.onStartLoading();
    await ChangeStatusRequest({ id: id, status: status });
    getPayments(filterPaymentIdSelected);
    isLoading.onEndLoading();
  }, []);

  const handleOpenEdit = (id: number) => {
    push(`payments/${id}?isEdit=true`);
  };

  const handleOpenDetail = (id: number) => {
    push(`payments/${id}`);
  };

  const handleFilterPaymentType = (id: number) => {
    setFilterPaymentIdSelected(id);
    getPayments(id);
  };

  return (
    <div>
      <SimpleCard>
        <WarningModal
          onSubmit={handleSubmitDelete}
          title="Eliminar Método de Pago"
          description="¿Estás seguro que quieres eliminar este método de pago?"
          confirmText="Eliminar"
        />
        <BreadcrumbNavigation items={bcItems} />

        <hr className="my-3" />

        <div className="grid grid-cols-1 lg:grid-cols-5 gap-4">
          <div>
            <h1 className="ml-2 font-bold">Métodos de Pago</h1>
          </div>
          <div className="lg:col-start-5">
            {showCreateBtn && (
              <Link href="/management/payments/create">
                <Button size="sm" variant="main" className="w-full">
                  Crear Método de Pago
                </Button>
              </Link>
            )}
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
                    isSelected={filterPaymentIdSelected === 0}
                    onlyAll={true}
                    onSelectPayment={handleFilterPaymentType}
                  />
                }
                {paymentTypes.map((val: any, index: number) => (
                  <PaymentFilterCard
                    key={index}
                    paymentType={val}
                    isSelected={filterPaymentIdSelected === val.id}
                    onlyAll={false}
                    onSelectPayment={handleFilterPaymentType}
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
                  <PaymentCard
                    payment={item}
                    onDelete={handleOpenDelete}
                    onDetails={handleOpenDetail}
                    onEdit={handleOpenEdit}
                    onChangeStatus={onChangeStatus}
                  />
                </div>
              ))}
            </div>
            {payments.length === 0 && (
              <div className="h-full flex justify-center">
                <span className="pt-4">No hay métodos de pago creados</span>
              </div>
            )}
          </SimpleCard>
        </div>
      </div>
    </div>
  );
};

export default PaymentsClient;
