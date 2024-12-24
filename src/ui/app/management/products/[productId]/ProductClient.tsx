"use client";

import { EditProductRequest, GetProductRequest } from "@/app/services/products";
import BreadcrumbNavigation from "@/app/components/BreadcrumbNavigation";
import SimpleCard from "@/app/components/cards/SimpleCard";
import ImagesUpload from "@/app/components/uploads/ImagesUpload";
import useLoading from "@/app/hooks/useLoading";
import { BreadcrumItem } from "@/app/types";
import { CreateProduct } from "@/app/types/product";
import {
  Button,
  Input,
  NumberInput,
  NumberInputField,
  useToast,
} from "@chakra-ui/react";
import { Icon } from "@iconify/react";
import Link from "next/link";
import { useParams, useRouter, useSearchParams } from "next/navigation";
import { useCallback, useEffect, useState } from "react";
import { validLettersAndNumbers, validNumbers } from "@/app/utils/InputUtils";

const ProductClient = () => {
  const router = useRouter();
  const params = useParams();
  const searchParams = useSearchParams();
  const isLoading = useLoading();

  const bcItems: BreadcrumItem[] = [
    {
      label: "Productos",
      href: "/management/products",
    },
    {
      label: "Cliente",
      href: `/management/products/${params.productId}`,
    },
  ];

  const toast = useToast();
  const [isEdit, setIsEdit] = useState(false);
  const [formData, setFormData] = useState<CreateProduct>({
    companyId: 0,
    name: "",
    description: "",
    images: [],
    sku: "",
    quantity: undefined,
    price: undefined,
    productStatus: 1,
  });

  const getProduct = useCallback(async () => {
    let product: any = await GetProductRequest({
      id: String(params.productId),
    });
    if (product) {
      setFormData({
        companyId: product.companyId,
        name: product.name ?? "",
        description: product.description ?? "",
        images: product.images ?? [],
        sku: product.sku ?? "",
        quantity: product.quantity,
        price: product.price,
        productStatus: product.status,
      });
    }
  }, [params.productId]);

  useEffect(() => {
    let paramIsEdit = searchParams.get("isEdit");
    if (paramIsEdit) {
      setIsEdit(true);
    }
    getProduct();
  }, [getProduct]);

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

  const onSubmit = async () => {
    if (isFormValid()) {
      isLoading.onStartLoading();
      let editProduct: any = await EditProductRequest({
        id: String(params.productId),
        ...formData,
      });
      if (editProduct?.error) {
        showErrorMessage(editProduct.error);
        isLoading.onEndLoading();
      } else {
        showSuccessEditMessage("Producto editado exitosamente");
        isLoading.onEndLoading();
        router.push("/management/products");
      }
    } else {
      showErrorMessage("Algunos campos son requeridos o inválidos");
    }
  };

  const isFormValid = (): boolean => {
    if (!validLettersAndNumbers(formData.name, true)) return false;

    if (
      formData.description &&
      !validLettersAndNumbers(formData.description, true)
    )
      return false;

    if (formData.sku && !validLettersAndNumbers(formData.sku)) return false;

    if (!formData.quantity || !validLettersAndNumbers(formData.sku))
      return false;

    if (!formData.price || !validLettersAndNumbers(formData.price?.toString()))
      return false;

    return true;
  };

  const showSuccessEditMessage = (msg: string) => {
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
          <h1 className="ml-2 font-bold">{`${
            isEdit ? "Editar" : ""
          } Cliente`}</h1>
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
              disabled={!isEdit}
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
              disabled={!isEdit}
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
              disabled={!isEdit}
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
                disabled={!isEdit}
                maxLength={15}
              />
            </NumberInput>
          </div>
          <div className="mt-2">
            <label className="text-sm">
              Precio <span className="text-thirdcolor">*</span>
            </label>
            <NumberInput size="sm" precision={2} value={formData.price}>
              <NumberInputField
                name="price"
                onChange={handleNumberChange}
                disabled={!isEdit}
                maxLength={15}
              />
            </NumberInput>
          </div>
        </SimpleCard>
      </div>

      <div className="mt-3">
        <SimpleCard>
          <div className="p-1">
            <label className="text-sm">Imágenes</label>
            <div className="border rounded py-5 px-3">
              <ImagesUpload showAddImage={isEdit} />
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

export default ProductClient;
