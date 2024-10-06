"use client";

import getCurrentUser from "@/app/actions/getCurrentUser";
import {
  ChangeStatusRequest,
  DeleteProductRequest,
  GetProductsRequest,
} from "@/app/api/products/route";
import BreadcrumbNavigation from "@/app/components/BreadcrumbNavigation";
import SimpleCard from "@/app/components/cards/SimpleCard";
import WarningModal from "@/app/components/modals/WarningModal";
import SimpleTable from "@/app/components/tables/SimpleTable";
import {
  ColumnType,
  SimpleTableColumn,
} from "@/app/components/tables/SimpleTable.types";
import { PLAN_ID, PRODUCT } from "@/app/constants";
import useLoading from "@/app/hooks/useLoading";
import useWarningModal from "@/app/hooks/useWarningModal";
import { BreadcrumItem } from "@/app/types";
import { CurrentUser } from "@/app/types/auth";
import { Product, SearchProduct } from "@/app/types/product";
import { Button, Input, useToast } from "@chakra-ui/react";
import { Icon } from "@iconify/react";
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
      type: ColumnType.ArrayImageFirst,
    },
    {
      key: "name",
      name: "Producto",
      type: ColumnType.String,
    },
    {
      key: "sku",
      name: "SKU",
      type: ColumnType.String,
    },
    {
      key: "quantity",
      name: "Cantidad",
      type: ColumnType.Number,
    },
    {
      key: "price",
      name: "precio",
      type: ColumnType.Money,
    },
  ];

  const { push } = useRouter();
  const isLoading = useLoading();
  const toast = useToast();
  const deleteProductModal = useWarningModal();
  const [searchProduct, setSearchProduct] = useState<SearchProduct>({
    name: "",
    sku: "",
  });
  const [offset, setOffset] = useState<number>(0);
  const [products, setProducts] = useState<Product[]>([]);
  const [productIdDelete, setProductIdDelete] = useState<string>("");
  const [showCreateBtn, setShowCreateBtn] = useState<boolean>(false);

  const getProducts = useCallback(
    async (searchParams: SearchProduct) => {
      isLoading.onStartLoading();
      const currentUser: CurrentUser | null = getCurrentUser();

      if (currentUser) {
        let data: Product[] = await GetProductsRequest({
          companyId: currentUser.companyId,
          name: searchParams.name ?? "",
          sku: searchParams.sku ?? "",
          limit: 10,
          offset: offset,
        });
        data = data.map((x) => {
          return {
            ...x,
            isActive: x.productStatus === 1 ? true : false,
          };
        });

        setShowCreateBtn(
          currentUser.planId === PLAN_ID.PRO ||
            data.length < PRODUCT.MAX_BASIC_PLAN_ITEMS
        );
        setProducts(data);
      }
      isLoading.onEndLoading();
    },
    [offset]
  );

  useEffect(() => {
    getProducts(searchProduct);
  }, [getProducts]);

  const handleChangePage = (val: string) => {
    setOffset((prevValue) =>
      val === "NEXT" ? (prevValue += 10) : (prevValue -= 10)
    );
  };

  const handleChange = (event: any) => {
    const { name, value } = event.target;
    setSearchProduct((prevData) => ({ ...prevData, [name]: value }));
  };

  const onSearchProduct = () => {
    getProducts(searchProduct);
  };

  const handleOpenDelete = (product: any) => {
    setProductIdDelete(product.id);
    deleteProductModal.onOpen();
  };

  const handleSubmitDelete = () => {
    onDelete(productIdDelete);
  };

  const onDelete = useCallback(async (id: string) => {
    isLoading.onStartLoading();
    await DeleteProductRequest({ id });
    getProducts(searchProduct);
    deleteProductModal.onClose();
    isLoading.onEndLoading();
    toast({
      title: "Producto",
      description: "Producto eliminado exitosamente",
      variant: "customsuccess",
      position: "top-right",
      duration: 3000,
      isClosable: true,
    });
  }, []);

  const onChangeStatus = useCallback(async (id: string, status: boolean) => {
    isLoading.onStartLoading();
    await ChangeStatusRequest({ id: id, productStatus: status ? 1 : 0 });
    getProducts(searchProduct);
    isLoading.onEndLoading();
  }, []);

  const handleOpenEdit = (val: any) => {
    push(`products/${val.id}?isEdit=true`);
  };

  const handleOpenDetail = (val: any) => {
    push(`products/${val.id}`);
  };

  return (
    <div>
      <SimpleCard>
        <WarningModal
          onSubmit={handleSubmitDelete}
          title="Eliminar PRoducto"
          description="¿Estás seguro que quieres eliminar este producto?"
          confirmText="Eliminar"
        />
        <BreadcrumbNavigation items={bcItems} />

        <hr className="my-3" />

        <div className="grid grid-cols-1 md:grid-cols-3 lg:grid-cols-5 xl:grid-cols-6 gap-4">
          <div>
            <label className="text-sm">Producto</label>
            <Input size="sm" name="name" onChange={handleChange} />
          </div>
          <div>
            <label className="text-sm">SKU</label>
            <Input size="sm" name="sku" onChange={handleChange} />
          </div>
          <div className="flex flex-col">
            <span className="opacity-0">.</span>
            <Button size="sm" variant="main" onClick={onSearchProduct}>
              <Icon icon="tabler:search" />
            </Button>
          </div>
          <div className="xl:col-start-6">
            <div className="flex flex-col">
              <span className="opacity-0">.</span>
              {showCreateBtn && (
                <Link href="/management/products/create">
                  <Button size="sm" variant="main" className="w-full">
                    Crear Producto
                  </Button>
                </Link>
              )}
            </div>
          </div>
        </div>
      </SimpleCard>

      <div className="mt-3">
        <SimpleCard>
          <SimpleTable
            columns={productCols}
            data={products}
            showToggleActive
            showDetails
            showEdit
            showDelete
            onEdit={handleOpenEdit}
            onDelete={handleOpenDelete}
            onDetail={handleOpenDetail}
            onChangeStatus={onChangeStatus}
            onChangePage={handleChangePage}
            offset={offset}
          />
        </SimpleCard>
      </div>
    </div>
  );
};

export default ProductsClient;
