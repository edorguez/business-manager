"use client";

import getCurrentUser from "@/app/actions/getCurrentUser";
import { GetCustomersByMonthsRequest } from "@/app/api/customers/route";
import { GetLatestProductsRequest } from "@/app/api/products/route";
import WelcomeBanner from "@/app/components/banners/WelcomeBanner";
import ListCard from "@/app/components/cards/ListCard";
import SimpleLineChartCard from "@/app/components/cards/SimpleLineChartCard";
import { ColumnType, SimpleTableColumn } from "@/app/components/tables/SimpleTable.types";
import useLoading from "@/app/hooks/useLoading";
import { CurrentUser } from "@/app/types/auth";
import { CustomerByMonth, CustomerMonths } from "@/app/types/customer";
import { Product } from "@/app/types/product";
import { convertToTimezone, formatTitleValue } from "@/app/utils/Utils";
import { useCallback, useEffect, useState } from "react";

const formatCustomerMonths = (dates: Date[]): CustomerMonths => {
  let now: Date = new Date();

  const array30Days: Date[] = Array.from({ length: 30 }, (_, days) => {
    let day = new Date(now); // clone "now"
    day.setDate(now.getDate() - days); // change the date
    return day;
  });

  const array60Days: Date[] = Array.from({ length: 60 }, (_, days) => {
    let day = new Date(now); // clone "now"
    day.setDate(now.getDate() - days); // change the date
    return day;
  });

  const array90Days: Date[] = Array.from({ length: 90 }, (_, days) => {
    let day = new Date(now); // clone "now"
    day.setDate(now.getDate() - days); // change the date
    return day;
  });

  let result: CustomerMonths = {
    oneMonth: {
      labels: array30Days,
      data: Array(30).fill(0),
      total: 0,
    },
    twoMonths: {
      labels: array60Days,
      data: Array(60).fill(0),
      total: 0,
    },
    threeMonths: {
      labels: array90Days,
      data: Array(90).fill(0),
      total: 0,
    },
  };

  for (let key in dates) {
    let baseDate: Date = new Date(dates[key]);
    convertToTimezone(baseDate, new Date().getTimezoneOffset());

    const currentUserMonth: number = new Date().getMonth();

    if (baseDate.getMonth() === currentUserMonth) {
      const idx: number = array30Days.findIndex(
        (x) => x.toLocaleDateString() === baseDate.toLocaleDateString()
      );
      result.oneMonth.data[idx]++;
      result.oneMonth.total++;
    } else if (baseDate.getMonth() === currentUserMonth - 1) {
      const idx: number = array60Days.findIndex(
        (x) => x.toLocaleDateString() === baseDate.toLocaleDateString()
      );
      result.twoMonths.data[idx] += 1;
      result.twoMonths.total++;
    } else if (baseDate.getMonth() === currentUserMonth - 2) {
      const idx: number = array90Days.findIndex(
        (x) => x.toLocaleDateString() === baseDate.toLocaleDateString()
      );
      result.threeMonths.data[idx] += 1;
      result.threeMonths.total++;
    }
  }

  return result;
};

const HomeClient = () => {
  const isLoading = useLoading();
  const [customerMonths, setCustomerMonths] = useState<CustomerMonths>();
  const [products, setProducts] = useState<Product[]>([]);

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
      key: "quantity",
      name: "Cantidad",
      type: ColumnType.Number
    },
    {
      key: "price",
      name: "precio",
      type: ColumnType.Money
    },
  ];

  const getCustomersByMonths = useCallback(async () => {
    isLoading.onStartLoading();
    const currentUser: CurrentUser | null = getCurrentUser();

    if (currentUser) {
      let data: CustomerByMonth = await GetCustomersByMonthsRequest({
        companyId: currentUser.companyId,
        months: 3,
      });
      setCustomerMonths(formatCustomerMonths(data.dates));
    }
    isLoading.onEndLoading();
  }, []);

  const getLatestProducts = useCallback(async () => {
    isLoading.onStartLoading();
    const currentUser: CurrentUser | null = getCurrentUser();

    if (currentUser) {
      let data: Product[] = await GetLatestProductsRequest({
        companyId: currentUser.companyId,
        limit: 5,
      });
      setProducts(data);
    }
    isLoading.onEndLoading();
  }, []);

  useEffect(() => {
    getCustomersByMonths();
    getLatestProducts();
  }, [getCustomersByMonths]);

  return (
    <>
      <WelcomeBanner />

      {customerMonths && (
        <div className="mt-4 grid grid-cols-1 lg:grid-cols-3 gap-4">
          <SimpleLineChartCard
            data={customerMonths.threeMonths.data}
            labels={customerMonths.threeMonths.labels}
            title="3 Meses"
            subtitle="Nuevos Clientes"
            total={customerMonths.threeMonths.total.toString()}
          />
          <SimpleLineChartCard
            data={customerMonths.twoMonths.data}
            labels={customerMonths.twoMonths.labels}
            title="2 Meses"
            subtitle="Nuevos Clientes"
            total={customerMonths.twoMonths.total.toString()}
          />
          <SimpleLineChartCard
            data={customerMonths.oneMonth.data}
            labels={customerMonths.oneMonth.labels}
            title="1 Mes"
            subtitle="Nuevos Clientes"
            total={customerMonths.oneMonth.total.toString()}
          />
        </div>
      )}

      {/* <div className="mt-4">
        <BarChartCard />
      </div> */}

      <div className="mt-4 grid grid-cols-1 lg:grid-cols-3 gap-4">
        {/* <DoughnutChartCard /> */}
        <ListCard title="Últimos Productos" columns={productCols} data={products} />
      </div>
    </>
  );
};

export default HomeClient;
