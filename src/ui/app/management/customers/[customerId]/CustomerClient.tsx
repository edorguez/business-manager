"use client";

import { Button, Input, Select, useToast } from "@chakra-ui/react";
import { BreadcrumItem } from "@/app/types";
import { Icon } from "@iconify/react";
import Link from "next/link";
import SimpleCard from "@/app/components/cards/SimpleCard";
import BreadcrumbNavigation from "@/app/components/BreadcrumbNavigation";
import { useCallback, useEffect, useState } from "react";
import { CreateCustomer } from "@/app/types/customer";
import { useRouter, useParams, useSearchParams } from "next/navigation";
import {
  CreateCustomerRequest,
  EditCustomerRequest,
  GetCustomerRequest,
} from "@/app/api/customers/route";
import { CurrentUser } from "@/app/types/auth";
import getCurrentUser from "@/app/actions/getCurrentUser";

const CustomerClient = () => {
  const router = useRouter();
  const params = useParams();
  const searchParams = useSearchParams();

  const bcItems: BreadcrumItem[] = [
    {
      label: "Clientes",
      href: "/management/customers",
    },
    {
      label: "Cliente",
      href: `/management/customers/${params.customerId}`,
    },
  ];

  const toast = useToast();
  const [isEdit, setIsEdit] = useState(false);
  const [formData, setFormData] = useState<CreateCustomer>({
    companyId: 0,
    firstName: "",
    lastName: "",
    email: "",
    identificationNumber: "",
    identificationType: "",
    phone: "",
  });

  const getCustomer = useCallback(async () => {
    let customer: any = await GetCustomerRequest({ id: +params.customerId });
    if (customer) {
      setFormData({
        companyId: customer.companyId,
        firstName: customer.firstName,
        lastName: customer.lastName ?? "",
        email: customer.email ?? "",
        identificationNumber: customer.identificationNumber ?? "",
        identificationType: customer.identificationType ?? "",
        phone: customer.phone ?? "",
      });
    }
  }, [params.customerId]);

  useEffect(() => {
    let paramIsEdit = searchParams.get('isEdit');
    if(paramIsEdit) {
      setIsEdit(true);
    }
    getCustomer();
  }, [getCustomer]);

  const handleChange = (event: any) => {
    const { name, value } = event.target;
    setFormData((prevFormData) => ({ ...prevFormData, [name]: value }));
  };

  const onSubmit = async () => {
    if (isFormValid()) {
      let editCustomer: any = await EditCustomerRequest({id: +params.customerId, ...formData});
      if (editCustomer?.error) {
        showErrorMessage(editCustomer.error);
      } else {
        showSuccessCreationMessage("Cliente editado exitosamente");
        router.push("/management/customers");
      }
    } else {
      showErrorMessage("Algunos campos son requeridos o inválidos");
    }
  };

  const isFormValid = (): boolean => {
    if (!formData.firstName) return false;

    if (!formData.identificationNumber || !formData.identificationType)
      return false;

    return true;
  };

  const showSuccessCreationMessage = (msg: string) => {
    toast({
      title: "Cliente",
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
            <Link href="/management/customers">
              <div className="rounded p-2 hover:bg-thirdcolor hover:text-white duration-150">
                <Icon icon="fa-solid:arrow-left" />
              </div>
            </Link>
          </div>
          <h1 className="ml-2 font-bold">{`${isEdit ? 'Editar' : ''} Cliente`}</h1>
        </div>
      </SimpleCard>

      <div className="mt-3">
        <SimpleCard>
          <div className="mt-2">
            <label className="text-sm">
              Nombre <span className="text-thirdcolor">*</span>
            </label>
            <Input
              size="sm"
              name="firstName"
              value={formData.firstName}
              onChange={handleChange}
              maxLength={20}
              disabled={!isEdit}
            />
          </div>
          <div className="mt-2">
            <label className="text-sm">Apellido</label>
            <Input
              size="sm"
              name="lastName"
              value={formData.lastName}
              onChange={handleChange}
              maxLength={20}
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
            <label className="text-sm">
              Cédula <span className="text-thirdcolor">*</span>
            </label>
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

export default CustomerClient;
