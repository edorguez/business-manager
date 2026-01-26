"use client";

import BreadcrumbNavigation from "@/app/components/BreadcrumbNavigation";
import SimpleCard from "@/app/components/cards/SimpleCard";
import { BreadcrumItem } from "@/app/types";
import { Input, Select } from "@chakra-ui/react";
import { Icon } from "@iconify/react";
import Link from "next/link";
import { useParams } from "next/navigation";
import { useCallback, useEffect, useState } from "react";
import { OrderDetails, OrderProduct, OrderProductsTable } from "@/app/types/order";
import { GetOrderRequest } from "@/app/services/orders";
import dayjs from 'dayjs';
import { convertTimestampToDate, convertToTimezone, numberMoveDecimal } from "@/app/utils/Utils";
import SimpleTable from "@/app/components/tables/SimpleTable";
import { ColumnType, SimpleTableColumn } from "@/app/components/tables/SimpleTable.types";

const OrderClient = () => {
  const params = useParams();

  const bcItems: BreadcrumItem[] = [
    {
      label: "Órdenes",
      href: "/management/orders",
    },
    {
      label: "Órden",
      href: `/management/orders/${params.orderId}`,
    },
  ];

  const [orderDateString, setOrderDateString] = useState<string>("");
  const [formData, setFormData] = useState<OrderDetails>({
    order: {
      id: 0,
      companyId: 0,
      customerId: 0,
      orderNumber: 0,
      createdAt: new Date()
    },
    customer: {
      id: 0,
      firstName: "",
      lastName: "",
      identificationNumber: "",
      identificationType: "",
      phone: "",
      email: ""
    },
    products: []
  });

  const [orderProductsTableData, setOrderProductsTableData] = useState<OrderProductsTable[]>([]);
  const productCols: SimpleTableColumn[] = [
    {
      key: "imageUrl",
      name: "Foto",
      type: ColumnType.Image
    },
    {
      key: "name",
      name: "Producto",
      type: ColumnType.String
    },
    {
      key: "quantity",
      name: "Cantidad",
      type: ColumnType.Number
    },
    {
      key: "price",
      name: "Precio",
      type: ColumnType.Money
    },
    {
      key: "total",
      name: "Total",
      type: ColumnType.Money
    },
  ]

  const getPayment = useCallback(async () => {
    let data: OrderDetails = await GetOrderRequest(+params.orderId);
    if (data) {
      const userTimeZone: number = new Date().getTimezoneOffset();
      const orderDate: Date = convertTimestampToDate(data.order.createdAt);
      const orderDateStringFormat: string = dayjs(convertToTimezone(orderDate, userTimeZone)).format('DD-MM-YYYY')

      setOrderDateString(orderDateStringFormat);
      setFormData({
        order: {
          id: data.order.id,
          companyId: data.order.companyId,
          customerId: data.order.customerId,
          orderNumber: data.order.orderNumber,
          createdAt: data.order.createdAt
        },
        customer: {
          id: data.customer.id,
          firstName: data.customer.firstName,
          lastName: data.customer.lastName,
          identificationNumber: data.customer.identificationNumber,
          identificationType: data.customer.identificationType,
          phone: data.customer.phone,
          email: ""
        },
        products: data.products
      });

      if(data.products) {
        data.products.forEach((x: OrderProduct) => {
          setOrderProductsTableData(prev =>  [...prev, {
            imageUrl: x.imageUrl,
            name: x.name,
            quantity: x.quantity,
            price: numberMoveDecimal(x.price, 2),
            total: x.quantity * numberMoveDecimal(x.price, 2)
          }])
        });
      }
    }
  }, [params.orderId]);

  useEffect(() => {
    getPayment();
  }, [getPayment]);

  return (
    <div>
      <SimpleCard>
        <BreadcrumbNavigation items={bcItems} />
        <hr className="my-3" />
        <div className="flex items-center">
          <div>
            <Link href="/management/orders">
              <div className="rounded p-2 hover:bg-thirdcolor hover:text-white duration-150">
                <Icon icon="fa-solid:arrow-left" />
              </div>
            </Link>
          </div>
          <h1 className="ml-2 font-bold">Órden</h1>
        </div>
      </SimpleCard>      

      <div className="mt-3">
        <SimpleCard>

          <h1 className="font-bold">Detalles de Órden</h1>
          
          <div className="mt-2">
            <label className="text-sm">Número de Órden</label>
            <Input
              size="sm"
              value={formData.order.orderNumber}
              disabled={true}
            />
          </div>

          <div className="mt-2">
            <label className="text-sm">Fecha</label>
            <Input
              size="sm"
              value={orderDateString}
              disabled={true}
            />
          </div>

        </SimpleCard>
      </div>

      <div className="mt-3">
        <SimpleCard>

          <h1 className="font-bold">Cliente</h1>
          
          <div className="mt-2">
            <label className="text-sm">Nombre</label>
            <Input
              size="sm"
              value={formData.customer.firstName}
              disabled={true}
            />
          </div>

          <div className="mt-2">
            <label className="text-sm">Apellido</label>
            <Input
              size="sm"
              value={formData.customer.lastName}
              disabled={true}
            />
          </div>

          <div className="mt-2">
            <label className="text-sm">Cédula</label>
            <div className="flex">
              <div className="w-24 mr-1">
                <Select
                  size="sm"
                  value={formData.customer.identificationType}
                  disabled={true}
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
                value={formData.customer.identificationNumber}
                disabled={true}
              />
            </div>
          </div>
          <div className="mt-2">
            <label className="text-sm">Teléfono</label>
            <Input
              size="sm"
              value={formData.customer.phone}
              disabled={true}
            />
          </div>
        </SimpleCard>
      </div>

      <div className="mt-3">
        <SimpleCard>

          <h1 className="font-bold">Productos</h1>

          <SimpleTable columns={productCols} data={orderProductsTableData} offset={undefined} />
        </SimpleCard>
      </div>
    </div>
  );
};

export default OrderClient;
