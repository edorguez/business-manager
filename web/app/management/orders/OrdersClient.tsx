'use client';

import getCurrentUser from '@/app/actions/getCurrentUser';
import { DeleteCustomerRequest, GetCustomersRequest } from '@/app/services/customers';
import BreadcrumbNavigation from '@/app/components/BreadcrumbNavigation';
import SimpleCard from '@/app/components/cards/SimpleCard';
import SimpleTable from '@/app/components/tables/SimpleTable';
import { ColumnType, SimpleTableColumn } from '@/app/components/tables/SimpleTable.types';
import { BreadcrumItem } from '@/app/types';
import { CurrentUser } from '@/app/types/auth';
import { Customer, SearchCustomer } from '@/app/types/customer';
import { Button, Input, useToast } from '@chakra-ui/react';
import { Icon } from '@iconify/react';
import Link from "next/link";
import { useCallback, useEffect, useState } from 'react';
import { useRouter } from 'next/navigation';
import useLoading from '@/app/hooks/useLoading';
import useWarningModal from '@/app/hooks/useWarningModal';
import WarningModal from '@/app/components/modals/WarningModal';
import { validLetters, validNumbers } from '@/app/utils/InputUtils';
import { GetOrdersRequest } from '@/app/services/orders';
import { GetOrdersResponse, OrderDetails, OrdersTable } from '@/app/types/order';
import { convertTimestampToDate, convertToTimezone } from '@/app/utils/Utils';
import dayjs from 'dayjs';

const OrdersClient = () => {
  const bcItems: BreadcrumItem[] = [
    {
      label: "Órdenes",
      href: "/management/orders"
    }
  ];

  const customerCols: SimpleTableColumn[] = [
    {
      key: "id",
      name: "N. Orden",
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
  const toast = useToast();
  const { push } = useRouter();
  const [ordersTableData, setOrdersTableData] = useState<OrdersTable[]>([]);
  const [offset, setOffset] = useState<number>(0);
  const deleteCustomerModal = useWarningModal();
  const [customerIdDelete, setCustomerIdDelete] = useState<number>(0);

  const getOrders = useCallback(async () => {
    isLoading.onStartLoading();
    const currentUser: CurrentUser | null = getCurrentUser();

    if (currentUser) {
      let data: GetOrdersResponse = await GetOrdersRequest({ companyId: currentUser.companyId, limit: 10, offset: offset });
      const formatData: OrdersTable[] = data.orders.map((x): OrdersTable => {
        const userTimeZone: number = new Date().getTimezoneOffset();
        const productDate: Date = convertTimestampToDate(x.order.createdAt);
        const productsString: string = x.products.map(x => `- (${x.quantity}) ${x.name.substring(0, 20)}`)
                                                  .slice(0, 5)
                                                  .join('\n');

        return {
          id: x.order.id,
          fullName: `${x.customer.firstName} ${x.customer.lastName}`,
          identificationNumber: `${x.customer.identificationType}-${x.customer.identificationNumber}`,
          date: dayjs(convertToTimezone(productDate, userTimeZone)).format('DD-MM-YYYY'),
          products: productsString
        }
      })

      setOrdersTableData(formatData);
    }
    isLoading.onEndLoading();
  }, [offset])

  useEffect(() => {
    getOrders();
  }, [getOrders]);

  const handleChangePage = (val: string) => {
    setOffset((prevValue) => val === 'NEXT' ? prevValue += 10 : prevValue -= 10);
  }

  // const handleNameChange = (event: any) => {
  //   const { name, value } = event.target;
  //   if (value && !validLetters(value, true)) return;
  //   setSearchCustomer((prevData) => ({ ...prevData, [name]: value }));
  // }
  //
  // const handleIdentificationNumberChange = (event: any) => {
  //   const { name, value } = event.target;
  //   if (value && !validNumbers(value)) return;
  //   setSearchCustomer((prevData) => ({ ...prevData, [name]: value }));
  // }

  // const onSearchCustomer = () => {
  //   getCustomers(searchCustomer);
  // }

  const handleOpenDetail = (val: any) => {
    push(`orders/${val.id}`);
  }

  return (
    <div>
      <SimpleCard>
        <BreadcrumbNavigation items={bcItems} />

        {/* <hr className="my-3" /> */}
        {/**/}
        {/* <div className="grid grid-cols-1 md:grid-cols-3 lg:grid-cols-5 xl:grid-cols-6 gap-4"> */}
        {/*   <div> */}
        {/*     <label className="text-sm">Nombre</label> */}
        {/*     <Input size="sm" name="name" onChange={handleNameChange} value={searchCustomer.name} maxLength={20} /> */}
        {/*   </div> */}
        {/*   <div> */}
        {/*     <label className="text-sm">Apellido</label> */}
        {/*     <Input size="sm" name="lastName" onChange={handleNameChange} value={searchCustomer.lastName} maxLength={20} /> */}
        {/*   </div> */}
        {/*   <div> */}
        {/*     <label className="text-sm">Cédula</label> */}
        {/*     <Input size="sm" name="identificationNumber" onChange={handleIdentificationNumberChange} value={searchCustomer.identificationNumber} maxLength={9} /> */}
        {/*   </div> */}
        {/*   <div className="flex flex-col"> */}
        {/*     <span className="opacity-0">.</span> */}
        {/*     <Button size="sm" variant="main" onClick={onSearchCustomer}> */}
        {/*       <Icon icon="tabler:search" /> */}
        {/*     </Button> */}
        {/*   </div> */}
        {/* </div> */}
      </SimpleCard>

      <div className="mt-3">
        <SimpleCard>
          <SimpleTable columns={customerCols} data={ordersTableData} showDetails onDetail={handleOpenDetail} onChangePage={handleChangePage} offset={offset} />
        </SimpleCard>
      </div>
    </div>
  )
}

export default OrdersClient;
