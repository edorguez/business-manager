"use client";

import { EditPaymentRequest, GetPaymentRequest } from "@/app/services/payment";
import { GetPaymentTypesRequest } from "@/app/services/paymentType";
import BreadcrumbNavigation from "@/app/components/BreadcrumbNavigation";
import PaymentFilterCard from "@/app/components/cards/PaymentFilterCard";
import SimpleCard from "@/app/components/cards/SimpleCard";
import useLoading from "@/app/hooks/useLoading";
import { BreadcrumItem } from "@/app/types";
import { CreatePayment } from "@/app/types/payment";
import { PaymentType } from "@/app/types/paymentType";
import { Button, Input, Select, useToast } from "@chakra-ui/react";
import { Icon } from "@iconify/react";
import Link from "next/link";
import { useParams, useSearchParams } from "next/navigation";
import { useRouter } from "next/navigation";
import { useCallback, useEffect, useState } from "react";

const PaymentClient = () => {
  const router = useRouter();
  const params = useParams();

  const bcItems: BreadcrumItem[] = [
    {
      label: "Métodos de Pago",
      href: "/management/payments",
    },
    {
      label: "Método de Pago",
      href: `/management/payments/${params.paymentId}`,
    },
  ];

  const isLoading = useLoading();
  const searchParams = useSearchParams();
  const [isEdit, setIsEdit] = useState(false);
  const [paymentTypes, setPaymentTypes] = useState<PaymentType[]>([]);
  const toast = useToast();
  const [formData, setFormData] = useState<CreatePayment>({
    companyId: 0,
    name: "",
    bank: "",
    accountNumber: "",
    accountType: "",
    identificationNumber: "",
    identificationType: "",
    phone: "",
    email: "",
    paymentTypeId: 0,
  });

  const getPaymentTypes = useCallback(async () => {
    isLoading.onStartLoading();
    const pt = await GetPaymentTypesRequest();
    setPaymentTypes(pt);
    isLoading.onEndLoading();
  }, []);

  const getPayment = useCallback(async () => {
    let payment: any = await GetPaymentRequest({ id: +params.paymentId });
    if (payment) {
      setFormData({
        companyId: payment.companyId,
        name: payment.name,
        bank: payment.bank ?? "",
        accountNumber: payment.accountNumber ?? "",
        accountType: payment.accountType ?? "",
        identificationNumber: payment.identificationNumber ?? "",
        identificationType: payment.identificationType ?? "",
        phone: payment.phone ?? "",
        email: payment.email ?? "",
        paymentTypeId: payment.paymentTypeId,
      });
    }
  }, [params.paymentId]);

  useEffect(() => {
    getPaymentTypes();

    let paramIsEdit = searchParams.get("isEdit");
    if (paramIsEdit) {
      setIsEdit(true);
    }
    getPayment();
  }, []);

  const handleChange = (event: any) => {
    const { name, value } = event.target;
    setFormData((prevFormData) => ({ ...prevFormData, [name]: value }));
  };

  const handlePaymentTypeSelected = (val: number) => {
    setFormData((prevFormData) => ({ ...prevFormData, paymentTypeId: val }));
  };

  const onSubmit = async () => {
    if (isFormValid()) {
      isLoading.onStartLoading();
      let editPayment: any = await EditPaymentRequest({
        id: +params.paymentId,
        ...formData,
      });
      if (!editPayment) {
        isLoading.onEndLoading();
        showSuccessEditMessage("Método de Pago editado exitosamente");
        router.push("/management/payments");
      } else {
        showErrorMessage(editPayment.error);
        isLoading.onEndLoading();
      }
    } else {
      showErrorMessage("Algunos campos son requeridos o inválidos");
    }
  };

  const isFormValid = (): boolean => {
    if (!formData.name) return false;

    if (!formData.paymentTypeId) return false;

    return true;
  };

  const showSuccessEditMessage = (msg: string) => {
    toast({
      title: "Método de Pago",
      description: msg,
      variant: "customsuccess",
      position: "top-right",
      duration: 3000,
      isClosable: true,
    });
  };

  const showErrorMessage = (msg: string) => {
    toast({
      title: "Error",
      description: msg,
      variant: "customerror",
      position: "top-right",
      duration: 3000,
      isClosable: true,
    });
  };

  return (
    <div>
      <SimpleCard>
        <BreadcrumbNavigation items={bcItems} />
        <hr className="my-3" />
        <div className="flex items-center">
          <div>
            <Link href="/management/payments">
              <div className="rounded p-2 hover:bg-thirdcolor hover:text-white duration-150">
                <Icon icon="fa-solid:arrow-left" />
              </div>
            </Link>
          </div>
          <h1 className="ml-2 font-bold">{`${
            isEdit ? "Editar" : ""
          } Método de Pago`}</h1>
        </div>
      </SimpleCard>

      <div className="mt-3">
        <SimpleCard>
          <div className="p-2">
            <label className="text-sm">Tipo de Cuenta</label>
            <div className="grid grid-cols-2 sm:grid-cols-3 lg:grid-cols-4 gap-3 mt-1">
              {paymentTypes.map((val: any, index: number) => {
                return isEdit ? (
                  <PaymentFilterCard
                    key={index}
                    paymentType={val}
                    isSelected={formData.paymentTypeId == val.id}
                    onlyAll={false}
                    onSelectPayment={handlePaymentTypeSelected}
                  />
                ) : (
                  <PaymentFilterCard
                    key={index}
                    paymentType={val}
                    isSelected={formData.paymentTypeId == val.id}
                    onlyAll={false}
                  />
                );
              })}
            </div>
          </div>
        </SimpleCard>
      </div>

      <div className="mt-3">
        <SimpleCard>
          <div className="mt-2">
            <label className="text-sm">
              Nombre <span className="text-thirdcolor">*</span>
            </label>
            <Input
              size="sm"
              name="name"
              value={formData.name}
              onChange={handleChange}
              maxLength={50}
              disabled={!isEdit}
            />
          </div>
          <div className="mt-2">
            <label className="text-sm">Banco</label>
            <Input
              size="sm"
              name="bank"
              value={formData.bank}
              onChange={handleChange}
              maxLength={50}
              disabled={!isEdit}
            />
          </div>
          <div className="mt-2">
            <label className="text-sm">Número de Cuenta</label>
            <Input
              size="sm"
              name="accountNumber"
              value={formData.accountNumber}
              onChange={handleChange}
              maxLength={20}
              disabled={!isEdit}
            />
          </div>
          <div className="mt-2">
            <label className="text-sm">Tipo de Cuenta</label>
            <Input
              size="sm"
              name="accountType"
              value={formData.accountType}
              onChange={handleChange}
              maxLength={20}
              disabled={!isEdit}
            />
          </div>
          <div className="mt-2">
            <label className="text-sm">Cédula</label>
            <div className="flex">
              <div className="w-24 mr-1">
                <Select
                  size="sm"
                  name="identificationType"
                  value={formData.identificationType}
                  onChange={handleChange}
                  disabled={!isEdit}
                >
                  <option value="">-</option>
                  <option value="V">V</option>
                  <option value="E">E</option>
                  <option value="P">P</option>
                  <option value="J">J</option>
                  <option value="G">G</option>
                </Select>
              </div>
              <Input
                size="sm"
                name="identificationNumber"
                value={formData.identificationNumber}
                onChange={handleChange}
                maxLength={20}
                disabled={!isEdit}
              />
            </div>
          </div>
          <div className="mt-2">
            <label className="text-sm">Teléfono</label>
            <Input
              size="sm"
              name="phone"
              value={formData.phone}
              onChange={handleChange}
              maxLength={11}
              disabled={!isEdit}
            />
          </div>
          <div className="mt-2">
            <label className="text-sm">Correo</label>
            <Input
              size="sm"
              name="email"
              value={formData.email}
              onChange={handleChange}
              maxLength={100}
              disabled={!isEdit}
            />
          </div>
        </SimpleCard>
      </div>

      {isEdit && (
        <div className="mt-3">
          <Button variant="main" className="w-full" onClick={onSubmit}>
            Editar
          </Button>
        </div>
      )}
    </div>
  );
};

export default PaymentClient;
