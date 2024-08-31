'use client';

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

const formatCustomerMonths = (customerByMonths: CustomerByMonth[]): CustomerMonths => {
  let result: CustomerMonths = {
    oneMonth: null,
    twoMonths: null,
    threeMonths: null
  };

  for(let key in customerByMonths) {
    // Transform date depending on user's timezone
    let baseDate: Date = new Date(customerByMonths[key].monthInterval);
    let timeOffsetInMS: number = new Date().getTimezoneOffset() * 60000;
    baseDate.setTime(baseDate.getTime() + timeOffsetInMS);

    const currentUserMonth: number = new Date().getMonth();

    if(baseDate.getMonth() === currentUserMonth) {
      result.oneMonth =  {
        monthInterval: baseDate,
        recordCount: customerByMonths[key].recordCount
      }
    } else if(baseDate.getMonth() === currentUserMonth - 1) {
      result.oneMonth =  {
        monthInterval: baseDate,
        recordCount: customerByMonths[key].recordCount
      }
    } else if(baseDate.getMonth() === currentUserMonth - 2) {
      result.oneMonth =  {
        monthInterval: baseDate,
        recordCount: customerByMonths[key].recordCount
      }
    }
  }

  return result;
};

const HomeClient = () => {

  const isLoading = useLoading();
  const [customerMonths, setCustomerMonths] = useState<CustomerMonths>();

  const getCustomersByMonths = useCallback(async () => {
    isLoading.onStartLoading();
    const currentUser: CurrentUser | null = getCurrentUser();

    if (currentUser) {
      let data: CustomerByMonth[] = await GetCustomersByMonthsRequest({ companyId: currentUser.companyId });
      setCustomerMonths(formatCustomerMonths(data))
    }
    isLoading.onEndLoading();
  }, []);

  useEffect(() => {
    getCustomersByMonths();
  }, [getCustomersByMonths]);


  return (
    <>
      <WelcomeBanner />

      <div className="mt-4 grid grid-cols-1 lg:grid-cols-3 gap-4">
        <SimpleLineChartCard />
        <SimpleLineChartCard />
        <SimpleLineChartCard />
      </div>

      {/* <div className="mt-4">
        <BarChartCard />
      </div> */}

      <div className="mt-4 grid grid-cols-1 lg:grid-cols-3 gap-4">
        {/* <DoughnutChartCard /> */}
        <ListCard />
      </div>
    </>
  )
}

export default HomeClient;
