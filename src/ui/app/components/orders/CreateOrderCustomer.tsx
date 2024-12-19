"use client";

import { Button, Input, Select, useToast } from "@chakra-ui/react";
import { useState } from "react";
import { CreateOrderCustomer } from "@/app/types/order";

interface CreateCustomerProps {
  onStartCreateOrderCustomer: (customer: CreateOrderCustomer) => void
}

const CreateOrderCustomerComponent: React.FC<CreateCustomerProps> = ({
  onStartCreateOrderCustomer,
}) => {
  const [isCreateLoading, setIsCreateLoading] = useState<boolean>(false);
  const toast = useToast();

  const [formData, setFormData] = useState<CreateOrderCustomer>({
    firstName: "",
    lastName: "",
    identificationNumber: "",
    identificationType: "",
    phone: "",
  });

  const handleChange = (event: any) => {
    const { name, value } = event.target;
    setFormData((prevFormData) => ({ ...prevFormData, [name]: value }));
  };

  const onSubmit = async () => {
    if (isFormValid()) {
      setIsCreateLoading(true);
      onStartCreateOrderCustomer(formData);
    } else {
      showErrorMessage("Algunos campos son requeridos o inválidos");
    }
  };

  const isFormValid = (): boolean => {
    if (!formData.firstName) return false;

    if (!formData.identificationNumber || !formData.identificationType)
      return false;

    if (!formData.phone)
      return false;

    return true;
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
    <>
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
          disabled={isCreateLoading}
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
          disabled={isCreateLoading}
        />
      </div>
      <div className="mt-2">
        <label className="text-sm">Teléfono <span className="text-thirdcolor">*</span></label>
        <Input
          size="sm"
          name="phone"
          value={formData.phone}
          onChange={handleChange}
          maxLength={11}
          disabled={isCreateLoading}
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
              disabled={isCreateLoading}
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
            disabled={isCreateLoading}
          />
        </div>
      </div>
      <div className="mt-3">
        <Button variant="main" size="sm" className="w-full" onClick={onSubmit} isLoading={isCreateLoading} loadingText="Creando Order">
          Enviar Datos
        </Button>
      </div>
    </>
  );
};

export default CreateOrderCustomerComponent;
