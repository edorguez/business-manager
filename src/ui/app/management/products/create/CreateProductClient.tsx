"use client";

import BreadcrumbNavigation from "@/app/components/BreadcrumbNavigation";
import SimpleCard from "@/app/components/cards/SimpleCard";
import { BreadcrumItem } from "@/app/types";
import Link from "next/link";
import { Icon } from "@iconify/react";
import {
  Button,
  Input,
  NumberInput,
  NumberInputField,
  useToast,
} from "@chakra-ui/react";
import ImagesUpload from "@/app/components/uploads/ImagesUpload";
import useLoading from "@/app/hooks/useLoading";
import { useRouter } from "next/navigation";
import { useEffect, useState } from "react";
import { CreateProduct } from "@/app/types/product";
import { CurrentUser } from "@/app/types/auth";
import getCurrentUser from "@/app/actions/getCurrentUser";
import { CreateProductRequest } from "@/app/services/products";

const CreateProductClient = () => {
  const bcItems: BreadcrumItem[] = [
    {
      label: "Productos",
      href: "/management/products",
    },
    {
      label: "Crear Producto",
      href: "/management/products/create",
    },
  ];

  const isLoading = useLoading();
  const toast = useToast();
  const { push } = useRouter();
  const [formData, setFormData] = useState<CreateProduct>({
    companyId: 0,
    name: "",
    description: "",
    images: [],
    sku: "",
    quantity: 0,
    price: 0,
    productStatus: 1
  });

  useEffect(() => {
    const currentUser: CurrentUser | null = getCurrentUser();
    if (currentUser) {
      formData.companyId = currentUser.companyId;
    }
  }, []);

  const handleChange = (event: any) => {
    const { name, value } = event.target;
    setFormData((prevFormData) => ({ ...prevFormData, [name]: value }));
  };

  const onSubmit = async () => {
    if (isFormValid()) {
      isLoading.onStartLoading();
      let createProduct: any = await CreateProductRequest(formData);
      if (!createProduct.error) {
        isLoading.onEndLoading();
        showSuccessCreationMessage("Producto creado exitosamente");
        push("/management/products");
      } else {
        showErrorMessage(createProduct.error);
        isLoading.onEndLoading();
      }
    } else {
      showErrorMessage("Algunos campos son requeridos o inválidos");
    }
  };

  const isFormValid = (): boolean => {
    if (!formData.name) return false;

    if (!formData.quantity) return false;

    if (!formData.price) return false;

    return true;
  };

  const showSuccessCreationMessage = (msg: string) => {
    toast({
      title: "Producto",
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
            <Link href="/management/products">
              <div className="rounded p-2 hover:bg-thirdcolor hover:text-white duration-150">
                <Icon icon="fa-solid:arrow-left" />
              </div>
            </Link>
          </div>
          <h1 className="ml-2 font-bold">Crear Producto</h1>
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
              name="name"
              value={formData.name}
              onChange={handleChange}
              maxLength={50}
            />
          </div>
          <div className="mt-2">
            <label className="text-sm">
              Descripción
            </label>
            <Input
              size="sm"
              name="description"
              value={formData.description}
              onChange={handleChange}
              maxLength={50}
            />
          </div>
          <div className="mt-2">
            <label className="text-sm">SKU</label>
            <Input
              size="sm"
              name="sku"
              value={formData.sku}
              onChange={handleChange}
              maxLength={12}
            />
          </div>
          <div className="mt-2">
            <label className="text-sm">
              Cantidad <span className="text-thirdcolor">*</span>
            </label>
            <NumberInput size="sm" value={formData.quantity}>
              <NumberInputField name="quantity" onChange={handleChange} />
            </NumberInput>
          </div>
          <div className="mt-2">
            <label className="text-sm">
              Precio <span className="text-thirdcolor">*</span>
            </label>
            <NumberInput size="sm" precision={2} value={formData.price}>
              <NumberInputField name="price" onChange={handleChange} />
            </NumberInput>
          </div>
        </SimpleCard>
      </div>

      <div className="mt-3">
        <SimpleCard>
          <div className="p-1">
            <label className="text-sm">Imágenes</label>
            <div className="border rounded py-5 px-3">
              <ImagesUpload showAddImage={true} />
            </div>
          </div>
        </SimpleCard>
      </div>

      <div className="mt-3">
        <Button variant="main" className="w-full" onClick={onSubmit}>
          Crear
        </Button>
      </div>
    </div>
  );
};

export default CreateProductClient;
