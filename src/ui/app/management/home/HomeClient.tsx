"use client";

import getCurrentUser from "@/app/actions/getCurrentUser";
import { GetCustomersByMonthsRequest } from "@/app/api/customers/route";
import WelcomeBanner from "@/app/components/banners/WelcomeBanner";
import BarChartCard from "@/app/components/cards/BarChartCard";
import DoughnutChartCard from "@/app/components/cards/DoughnutChartCard";
import ListCard from "@/app/components/cards/ListCard";
import SimpleLineChartCard from "@/app/components/cards/SimpleLineChartCard";
import useLoading from "@/app/hooks/useLoading";
import { CurrentUser } from "@/app/types/auth";
import { CustomerByMonth, CustomerMonths } from "@/app/types/customer";
import { convertToTimezone, formatTitleValue } from "@/app/utils/Utils";
import dayjs from "dayjs";
import { useCallback, useEffect, useState } from "react";

const formatCustomerMonths = (dates: Date[]): CustomerMonths => {
  let now: Date = new Date();

  const array30Days: Date[] = Array.from({ length: 30 }, (_, days) => {
    let day = new Date(now) // clone "now"
    day.setDate(now.getDate() - days) // change the date
    return day;
  });

  const array60Days: Date[] = Array.from({ length: 60 }, (_, days) => {
    let day = new Date(now) // clone "now"
    day.setDate(now.getDate() - days) // change the date
    return day;
  });

  const array90Days: Date[] = Array.from({ length: 90 }, (_, days) => {
    let day = new Date(now) // clone "now"
    day.setDate(now.getDate() - days) // change the date
    return day;
  });

  let result: CustomerMonths = {
    oneMonth: {
      labels: array30Days,
      data: Array(30).fill(0)
    },
    twoMonths: {
      labels: array60Days ,
      data: Array(60).fill(0)
    },
    threeMonths: {
      labels: array90Days,
      data: Array(90).fill(0)
    },
  };

  for (let key in dates) {
    let baseDate: Date = new Date(dates[key]);
    convertToTimezone(baseDate, new Date().getTimezoneOffset());

    const currentUserMonth: number = new Date().getMonth();

    if (baseDate.getMonth() === currentUserMonth) {
      const idx: number = array30Days.findIndex(x => x.toLocaleDateString() === baseDate.toLocaleDateString());
      result.oneMonth.data[idx] += 1;
    } else if (baseDate.getMonth() === currentUserMonth - 1) {
      const idx: number = array60Days.findIndex(x => x.toLocaleDateString() === baseDate.toLocaleDateString());
      result.twoMonths.data[idx] += 1;
    } else if (baseDate.getMonth() === currentUserMonth - 2) {
      const idx: number = array90Days.findIndex(x => x.toLocaleDateString() === baseDate.toLocaleDateString());
      result.threeMonths.data[idx] += 1;
    }
  }

  return result;
};

const HomeClient = () => {
  const isLoading = useLoading();
  const [customerMonths, setCustomerMonths] = useState<CustomerMonths>();
  const [customerMonthsLabels, setCustomerMonthsLabels] = useState<string[]>([]);

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

  useEffect(() => {
    getCustomersByMonths();
  }, [getCustomersByMonths]);

  return (
    <>
      <WelcomeBanner />

      {customerMonths && customerMonthsLabels && (
        <div className="mt-4 grid grid-cols-1 lg:grid-cols-3 gap-4">
          <SimpleLineChartCard data={customerMonths.threeMonths.data} labels={customerMonths.threeMonths.labels}/>
          <SimpleLineChartCard data={customerMonths.twoMonths.data} labels={customerMonths.twoMonths.labels} />
          <SimpleLineChartCard data={customerMonths.oneMonth.data} labels={customerMonths.oneMonth.labels} />
        </div>
      )}

      {/* <div className="mt-4">
        <BarChartCard />
      </div> */}

      <div className="mt-4 grid grid-cols-1 lg:grid-cols-3 gap-4">
        {/* <DoughnutChartCard /> */}
        <ListCard />
      </div>
    </>
  );
};

export default HomeClient;
