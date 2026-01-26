'use client';

import getCurrentUser from '@/app/actions/getCurrentUser';
import BreadcrumbNavigation from '@/app/components/BreadcrumbNavigation';
import SimpleCard from '@/app/components/cards/SimpleCard';
import SimpleTable from '@/app/components/tables/SimpleTable';
import { ColumnType, SimpleTableColumn } from '@/app/components/tables/SimpleTable.types';
import { BreadcrumItem } from '@/app/types';
import { CurrentUser } from '@/app/types/auth';
import { useCallback, useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';
import useLoading from '@/app/hooks/useLoading';
import { GetOrdersRequest } from '@/app/services/orders';
import { GetOrdersResponse, OrdersTable } from '@/app/types/order';
import { convertTimestampToDate, convertToTimezone } from '@/app/utils/Utils';
import dayjs from 'dayjs';

const OrdersClient = () => {
  const bcItems: BreadcrumItem[] = [
    {
      label: "Órdenes",
      href: "/management/orders"
    }
  ];

  const orderCols: SimpleTableColumn[] = [
    {
      key: "orderNumber",
      name: "N. Órden",
      type: ColumnType.String
    },
    {
      key: "fullName",
      name: "Cliente",
      type: ColumnType.String
    },
    {
      key: "identificationNumber",
      name: "Cédula",
      type: ColumnType.String
    },
    {
      key: "date",
      name: "Fecha",
      type: ColumnType.String
    },
    {
      key: "products",
      name: "Productos",
      type: ColumnType.String
    },
  ]

  const isLoading = useLoading();
  const { push } = useRouter();
  const [ordersTableData, setOrdersTableData] = useState<OrdersTable[]>([]);
  const [offset, setOffset] = useState<number>(0);

  const getOrders = useCallback(async () => {
    isLoading.onStartLoading();
    const currentUser: CurrentUser | null = getCurrentUser();

    if (currentUser) {
      let data: GetOrdersResponse = await GetOrdersRequest({ companyId: currentUser.companyId, limit: 10, offset: offset });
      if(data?.orders) {
        const formatData: OrdersTable[] = data.orders.map((x): OrdersTable => {
          const userTimeZone: number = new Date().getTimezoneOffset();
          const orderDate: Date = convertTimestampToDate(x.order.createdAt);
          const productsString: string = x.products.map(x => `- (${x.quantity}) ${x.name.substring(0, 20)}`)
                                                    .slice(0, 5)
                                                    .join('\n');

          return {
            id: x.order.id,
            orderNumber: x.order.orderNumber,
            fullName: `${x.customer.firstName} ${x.customer.lastName}`,
            identificationNumber: `${x.customer.identificationType}-${x.customer.identificationNumber}`,
            date: dayjs(convertToTimezone(orderDate, userTimeZone)).format('DD-MM-YYYY'),
            products: productsString
          }
        })

        setOrdersTableData(formatData);
      }
    }

    isLoading.onEndLoading();
  }, [offset])

  useEffect(() => {
    getOrders();
  }, [getOrders]);

  const handleChangePage = (val: string) => {
    setOffset((prevValue) => val === 'NEXT' ? prevValue += 10 : prevValue -= 10);
  }

  const handleOpenDetail = (val: any) => {
    push(`orders/${val.id}`);
  }

  return (
    <div>
      <SimpleCard>
        <BreadcrumbNavigation items={bcItems} />
      </SimpleCard>

      <div className="mt-3">
        <SimpleCard>
          <SimpleTable columns={orderCols} data={ordersTableData} showDetails onDetail={handleOpenDetail} onChangePage={handleChangePage} offset={offset} />
        </SimpleCard>
      </div>
    </div>
  )
}

export default OrdersClient;
