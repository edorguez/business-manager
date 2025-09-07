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
import {
  validLettersAndNumbers,
  validNumbers,
  validPrice,
} from "@/app/utils/InputUtils";
import { formatPriceStringToNumberBackend } from "@/app/utils/Utils";

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
  const [inputPrice, setInputPrice] = useState<string>("");
  const [images, setImages] = useState<File[]>([]);
  const [formData, setFormData] = useState<CreateProduct>({
    companyId: 0,
    name: "",
    description: "",
    sku: "",
    quantity: undefined,
    price: undefined,
    productStatus: 1,
  });

  useEffect(() => {
    const currentUser: CurrentUser | null = getCurrentUser();
    if (currentUser) {
      formData.companyId = currentUser.companyId;
    }
  }, []);

  const handleNameChange = (event: any) => {
    const { name, value } = event.target;
    if (value && !validLettersAndNumbers(value, true)) return;
    setFormData((prevFormData) => ({ ...prevFormData, [name]: value }));
  };

  const handleSKUChange = (event: any) => {
    let { name, value } = event.target;
    if (value && !validLettersAndNumbers(value)) return;
    if (value) value = value.toUpperCase();
    setFormData((prevFormData) => ({ ...prevFormData, [name]: value }));
  };

  const handleNumberChange = (event: any) => {
    let { name, value } = event.target;
    if (value && !validNumbers(value)) return;
    setFormData((prevFormData) => ({ ...prevFormData, [name]: value }));
  };

  const handlePriceChange = (event: any) => {
    let { name, value } = event.target;
    if (value && !validPrice(value)) return;
    setInputPrice(value);
    setFormData((prevFormData) => ({
      ...prevFormData,
      [name]: formatPriceStringToNumberBackend(value),
    }));
  };

  const onSubmit = async () => {
    if (isFormValid()) {
      isLoading.onStartLoading();
      let createProduct: any = await CreateProductRequest(formData, images);
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

    if (!validLettersAndNumbers(formData.name, true)) return false;

    if (
      formData.description &&
      !validLettersAndNumbers(formData.description, true)
    )
      return false;

    if (formData.sku && !validLettersAndNumbers(formData.sku)) return false;

    if ((formData.quantity ?? undefined) === undefined) return false;

    if (formData.quantity && !validNumbers(formData.quantity.toString()))
      return false;

    if (!formData.price || !validLettersAndNumbers(formData.price?.toString()))
      return false;

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

  const handleUploadFiles = (files: File[]) => {
    setImages(files);
  }

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
              onChange={handleNameChange}
              maxLength={50}
            />
          </div>
          <div className="mt-2">
            <label className="text-sm">Descripción</label>
            <Input
              size="sm"
              name="description"
              value={formData.description}
              onChange={handleNameChange}
              maxLength={50}
            />
          </div>
          <div className="mt-2">
            <label className="text-sm">SKU</label>
            <Input
              size="sm"
              name="sku"
              value={formData.sku}
              onChange={handleSKUChange}
              maxLength={12}
            />
          </div>
          <div className="mt-2">
            <label className="text-sm">
              Cantidad <span className="text-thirdcolor">*</span>
            </label>
            <NumberInput size="sm" value={formData.quantity}>
              <NumberInputField
                name="quantity"
                onChange={handleNumberChange}
                maxLength={15}
              />
            </NumberInput>
          </div>
          <div className="mt-2">
            <label className="text-sm">
              Precio <span className="text-thirdcolor">*</span>
            </label>
            <Input
              size="sm"
              name="price"
              value={inputPrice}
              onChange={handlePriceChange}
              maxLength={15}
            />
          </div>
        </SimpleCard>
      </div>

      <div className="mt-3">
        <SimpleCard>
          <div className="p-1">
            <label className="text-sm">Imágenes</label>
            <div className="border rounded py-5 px-3">
              <ImagesUpload isViewOnlyImage={false} onUploadFiles={handleUploadFiles}/>
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
