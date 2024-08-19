'use client';

import getCurrentUser from "@/app/actions/getCurrentUser";
import { DeleteProductRequest, GetProductsRequest } from "@/app/api/products/route";
import BreadcrumbNavigation from "@/app/components/BreadcrumbNavigation";
import SimpleCard from "@/app/components/cards/SimpleCard";
import DeleteModal from "@/app/components/modals/DeleteModal";
import SimpleTable from "@/app/components/tables/SimpleTable";
import { ColumnType, SimpleTableColumn } from "@/app/components/tables/SimpleTable.types";
import useDeleteModal from "@/app/hooks/useDeleteModal";
import useLoading from "@/app/hooks/useLoading";
import { BreadcrumItem } from "@/app/types";
import { CurrentUser } from "@/app/types/auth";
import { Product } from "@/app/types/product";
import { Button, Input, useToast } from "@chakra-ui/react";
import { Icon } from '@iconify/react';
import Link from "next/link";
import { useRouter } from "next/navigation";
import { useCallback, useEffect, useState } from "react";


const ProductsClient = () => {
  
  const bcItems: BreadcrumItem[] = [
    {
      label: "Productos",
      href: "/management/products",
    },
  ];
  
  const productCols: SimpleTableColumn[] = [
    {
      key: "images",
      name: "",
      type: ColumnType.ArrayImageFirst
    },
    {
      key: "name",
      name: "Producto",
      type: ColumnType.String
    },
    {
      key: "sku",
      name: "SKU",
      type: ColumnType.String
    },
    {
      key: "quantity",
      name: "Cantidad",
      type: ColumnType.Number
    },
    {
      key: "price",
      name: "precio",
      type: ColumnType.Money
    },
  ]

  const { push } = useRouter();
  const isLoading = useLoading();
  const toast = useToast();
  const deleteProductModal = useDeleteModal();
  const [offset, setOffset] = useState<number>(0);
  const [products, setProducts] = useState<Product[]>([]);
  const [productIdDelete, setProductIdDelete] = useState<number>(0);

  const getProducts = useCallback(async () => {
    isLoading.onStartLoading();
    const currentUser: CurrentUser | null = getCurrentUser();

    if (currentUser) {
      let data: Product[] = await GetProductsRequest({
        companyId: currentUser.companyId,
        limit: 10,
        offset: offset,
      });
      setProducts(data);
    }
    isLoading.onEndLoading();
  }, [offset]);

  useEffect(() => {
    getProducts();
  }, []);
  
  const handleChangePage = (val: string) => {
    setOffset((prevValue) => val === 'NEXT' ? prevValue += 10 : prevValue -= 10);
  }

  const handleOpenDelete = (id: number) => {
    setProductIdDelete(id);
    deleteProductModal.onOpen();
  }

  const handleSubmitDelete = () => {
    onDelete(productIdDelete);
  }
  
  const onDelete = useCallback(async (id: number) => {
    isLoading.onStartLoading();
    await DeleteProductRequest({ id });
    getProducts();
    deleteProductModal.onClose();
    isLoading.onEndLoading()
    toast({
      title: 'Producto',
      description: 'Producto eliminado exitosamente',
      variant: 'customsuccess',
      position: 'top-right',
      duration: 3000,
      isClosable: true,
    });
  }, [])

  const onChangeStatus = useCallback(async (id: number, status: boolean) => {
    isLoading.onStartLoading();
    // await ChangeStatusRequest({ id: id, status:  status});
    getProducts();
    isLoading.onEndLoading()
  }, [])
  
  const handleOpenEdit = (id: number) => {
    push(`products/${id}?isEdit=true`);
  }

  const handleOpenDetail = (id: number) => {
    push(`products/${id}`);
  }
  
  return (
    <div>
      <SimpleCard>
        <DeleteModal onSubmit={handleSubmitDelete} title="Eliminar PRoducto" description="¿Estás seguro que quieres eliminar este producto?" />
        <BreadcrumbNavigation items={bcItems} />

        <hr className="my-3" />

        <div className="grid grid-cols-1 md:grid-cols-3 lg:grid-cols-5 xl:grid-cols-6 gap-4">
          <div>
            <label className="text-sm">Producto</label>
            <Input size="sm" />
          </div>
          <div>
            <label className="text-sm">SKU</label>
            <Input size="sm" />
          </div>
          <div className="flex flex-col">
            <span className="opacity-0">.</span>
            <Button size="sm" variant="main">
              <Icon icon="tabler:search" />
            </Button>
          </div>
          <div className="xl:col-start-6">
            <div className="flex flex-col">
              <span className="opacity-0">.</span>
              <Link href="/management/products/create">
                <Button size="sm" variant="main" className="w-full">
                  Crear Producto
                </Button>
              </Link>
            </div>
          </div>
        </div>
      </SimpleCard>

      <div className="mt-3">
        <SimpleCard>
          <SimpleTable columns={productCols} data={products} showToggleActive showDetails showEdit showDelete onDelete={deleteProductModal.onOpen} onChangePage={handleChangePage} offset={offset} />
        </SimpleCard>
      </div>
    </div>
  )

}

export default ProductsClient;
