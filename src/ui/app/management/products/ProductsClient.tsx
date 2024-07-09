'use client';

import BreadcrumbNavigation from "@/app/components/BreadcrumbNavigation";
import SimpleCard from "@/app/components/cards/SimpleCard";
import DeleteModal from "@/app/components/modals/DeleteModal";
import SimpleTable from "@/app/components/tables/SimpleTable";
import { ColumnType, SimpleTableColumn } from "@/app/components/tables/SimpleTable.types";
import useDeleteModal from "@/app/hooks/useDeleteModal";
import { BreadcrumItem } from "@/app/types";
import { Product } from "@/app/types/product";
import { Button, Input } from "@chakra-ui/react";
import { Icon } from '@iconify/react';
import Link from "next/link";


const ProductsClient = () => {
  const bcItems: BreadcrumItem[] = [
    {
      label: "Productos",
      href: "/management/products"
    }
  ];

  const productCols: SimpleTableColumn[] = [
    {
      key: "imageUrl",
      name: "",
      type: ColumnType.Image
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

  const productData: Product[] = [
    {
      id: 1,
      name: "Camisa",
      imageUrl: "https://marketplace.canva.com/print-mockup/bundle/E9Me4jcyzMX/fit:female,pages:double-sided,surface:marketplace/product:171618,surface:marketplace/EAFam5QLuIc/1/0/933w/canva-black-white-typography-motivation-tshirt-WKRZLU21i2c.png?sig=bc03703936ce8090247068bcf3a44f0e&width=800",
      sku: "H19V2312",
      quantity: 100,
      price: 999.99
    },
    {
      id: 2,
      name: "Camisa",
      imageUrl: "https://marketplace.canva.com/print-mockup/bundle/E9Me4jcyzMX/fit:female,pages:double-sided,surface:marketplace/product:171618,surface:marketplace/EAFam5QLuIc/1/0/933w/canva-black-white-typography-motivation-tshirt-WKRZLU21i2c.png?sig=bc03703936ce8090247068bcf3a44f0e&width=800",
      sku: "H19V2312",
      quantity: 100,
      price: 999.99
    },
    {
      id: 3,
      name: "Camisa",
      imageUrl: "https://marketplace.canva.com/print-mockup/bundle/E9Me4jcyzMX/fit:female,pages:double-sided,surface:marketplace/product:171618,surface:marketplace/EAFam5QLuIc/1/0/933w/canva-black-white-typography-motivation-tshirt-WKRZLU21i2c.png?sig=bc03703936ce8090247068bcf3a44f0e&width=800",
      sku: "H19V2312",
      quantity: 100,
      price: 999.99
    },
    {
      id: 4,
      name: "Camisa",
      imageUrl: "https://marketplace.canva.com/print-mockup/bundle/E9Me4jcyzMX/fit:female,pages:double-sided,surface:marketplace/product:171618,surface:marketplace/EAFam5QLuIc/1/0/933w/canva-black-white-typography-motivation-tshirt-WKRZLU21i2c.png?sig=bc03703936ce8090247068bcf3a44f0e&width=800",
      sku: "H19V2312",
      quantity: 100,
      price: 999.99
    },
    {
      id: 5,
      name: "Camisa",
      imageUrl: "https://marketplace.canva.com/print-mockup/bundle/E9Me4jcyzMX/fit:female,pages:double-sided,surface:marketplace/product:171618,surface:marketplace/EAFam5QLuIc/1/0/933w/canva-black-white-typography-motivation-tshirt-WKRZLU21i2c.png?sig=bc03703936ce8090247068bcf3a44f0e&width=800",
      sku: "H19V2312",
      quantity: 100,
      price: 999.99
    },
    {
      id: 6,
      name: "Camisa",
      imageUrl: "https://marketplace.canva.com/print-mockup/bundle/E9Me4jcyzMX/fit:female,pages:double-sided,surface:marketplace/product:171618,surface:marketplace/EAFam5QLuIc/1/0/933w/canva-black-white-typography-motivation-tshirt-WKRZLU21i2c.png?sig=bc03703936ce8090247068bcf3a44f0e&width=800",
      sku: "H19V2312",
      quantity: 100,
      price: 999.99
    },
    {
      id: 7,
      name: "Camisa",
      imageUrl: "https://marketplace.canva.com/print-mockup/bundle/E9Me4jcyzMX/fit:female,pages:double-sided,surface:marketplace/product:171618,surface:marketplace/EAFam5QLuIc/1/0/933w/canva-black-white-typography-motivation-tshirt-WKRZLU21i2c.png?sig=bc03703936ce8090247068bcf3a44f0e&width=800",
      sku: "H19V2312",
      quantity: 100,
      price: 999.99
    },
    {
      id: 8,
      name: "Camisa",
      imageUrl: "https://marketplace.canva.com/print-mockup/bundle/E9Me4jcyzMX/fit:female,pages:double-sided,surface:marketplace/product:171618,surface:marketplace/EAFam5QLuIc/1/0/933w/canva-black-white-typography-motivation-tshirt-WKRZLU21i2c.png?sig=bc03703936ce8090247068bcf3a44f0e&width=800",
      sku: "H19V2312",
      quantity: 100,
      price: 999.99
    },
    {
      id: 9,
      name: "Camisa",
      imageUrl: "https://marketplace.canva.com/print-mockup/bundle/E9Me4jcyzMX/fit:female,pages:double-sided,surface:marketplace/product:171618,surface:marketplace/EAFam5QLuIc/1/0/933w/canva-black-white-typography-motivation-tshirt-WKRZLU21i2c.png?sig=bc03703936ce8090247068bcf3a44f0e&width=800",
      sku: "H19V2312",
      quantity: 100,
      price: 999.99
    },
    {
      id: 10,
      name: "Camisa",
      imageUrl: "https://marketplace.canva.com/print-mockup/bundle/E9Me4jcyzMX/fit:female,pages:double-sided,surface:marketplace/product:171618,surface:marketplace/EAFam5QLuIc/1/0/933w/canva-black-white-typography-motivation-tshirt-WKRZLU21i2c.png?sig=bc03703936ce8090247068bcf3a44f0e&width=800",
      sku: "H19V2312",
      quantity: 100,
      price: 999.99
    },
  ]

  const deleteProductModal = useDeleteModal();

  return (
    <div>
      <SimpleCard>
        <DeleteModal onSubmit={() => { }} title="Eliminar Producto" description="¿Estás seguro que quieres eliminar este producto?" />
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
          <SimpleTable columns={productCols} data={productData} showToggleActive showDetails showEdit showDelete onDelete={deleteProductModal.onOpen} />
        </SimpleCard>
      </div>
    </div>
  )

}

export default ProductsClient;
