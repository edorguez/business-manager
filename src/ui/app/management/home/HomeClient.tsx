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
import { useCallback, useEffect, useState } from "react";

const formatCustomerMonths = (dates: Date[]): CustomerMonths => {
  let result: CustomerMonths = {
    oneMonth: {
      labels: [],
      data: []
    },
    twoMonths: {
      labels: [],
      data: []
    },
    threeMonths: {
      labels: [],
      data: []
    },
  };

  for (let key in dates) {
    // Transform date depending on user's timezone
    let baseDate: Date = new Date(dates[key]);
    let timeOffsetInMS: number = new Date().getTimezoneOffset() * 60000;
    baseDate.setTime(baseDate.getTime() + timeOffsetInMS);

    const currentUserMonth: number = new Date().getMonth();

    if (baseDate.getMonth() === currentUserMonth) {
      result.oneMonth.data.push(baseDate);
    } else if (baseDate.getMonth() === currentUserMonth - 1) {
      result.twoMonths.data.push(baseDate);
    } else if (baseDate.getMonth() === currentUserMonth - 2) {
      result.threeMonths.data.push(baseDate);
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
