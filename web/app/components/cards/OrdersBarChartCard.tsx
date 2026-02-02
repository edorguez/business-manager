"use client";

import BarChart from "../charts/BarChart";
import { tailwindConfig, convertToTimezone } from "@/app/utils/Utils";
import { GetOrdersByMonthRequest } from "@/app/services/orders";
import getCurrentUser from "@/app/actions/getCurrentUser";
import { CurrentUser } from "@/app/types/auth";
import { useEffect, useState } from "react";

const OrdersBarChartCard = () => {
  const [currentMonth, setCurrentMonth] = useState<Date>(new Date());
  const [ordersData, setOrdersData] = useState<number[]>([]);
  const [labels, setLabels] = useState<Date[]>([]);
  const [totalOrders, setTotalOrders] = useState<number>(0);

  const getMonthName = (date: Date): string => {
    return date.toLocaleString('default', { month: 'long', year: 'numeric' });
  };

  const getDaysInMonth = (year: number, month: number): Date[] => {
    const date = new Date(year, month, 1);
    const days: Date[] = [];
    while (date.getMonth() === month) {
      days.push(new Date(date));
      date.setDate(date.getDate() + 1);
    }
    return days;
  };

  const integerFormatter = (value: any): string => {
    return Math.round(Number(value)).toString();
  };

  const prevMonth = () => {
    setCurrentMonth(prev => {
      const newDate = new Date(prev);
      newDate.setMonth(prev.getMonth() - 1);
      return newDate;
    });
  };

  const nextMonth = () => {
    setCurrentMonth(prev => {
      const newDate = new Date(prev);
      newDate.setMonth(prev.getMonth() + 1);
      return newDate;
    });
  };

  const goToCurrentMonth = () => {
    setCurrentMonth(new Date());
  };

  useEffect(() => {
    const fetchOrdersByMonth = async () => {
      const currentUser: CurrentUser | null = getCurrentUser();
      if (!currentUser) {
        return;
      }

      const year = currentMonth.getFullYear();
      const month = currentMonth.getMonth() + 1; // JavaScript months are 0-indexed, backend expects 1-indexed

      const daysInMonth = getDaysInMonth(year, month - 1);
      const dayCounts = new Array(daysInMonth.length).fill(0);

      try {
        const response = await GetOrdersByMonthRequest({
          companyId: currentUser.companyId,
          year,
          month,
        });

        if (response && response.createdAt) {
          // Count orders per day
          const userTimezoneOffset = new Date().getTimezoneOffset();
          response.createdAt.forEach((timestampStr) => {
            let timestamp = new Date(timestampStr);
            timestamp = convertToTimezone(timestamp, userTimezoneOffset);
            const dayOfMonth = timestamp.getDate() - 1; // zero-index
            if (dayOfMonth >= 0 && dayOfMonth < dayCounts.length) {
              dayCounts[dayOfMonth]++;
            }
          });
        }
      } catch (error) {
        console.error("Failed to fetch orders by month:", error);
      }

      setLabels(daysInMonth);
      setOrdersData(dayCounts);
      setTotalOrders(dayCounts.reduce((sum, count) => sum + count, 0));
    };

    fetchOrdersByMonth();
  }, [currentMonth]);

  const chartData = {
    labels: labels.map(date => {
      // Format as MM-DD-YYYY for Chart.js time parser
      const mm = String(date.getMonth() + 1).padStart(2, '0');
      const dd = String(date.getDate()).padStart(2, '0');
      const yyyy = date.getFullYear();
      return `${mm}-${dd}-${yyyy}`;
    }),
    datasets: [
      {
        label: "Orders",
        data: ordersData,
        backgroundColor: tailwindConfig().theme.colors.blue[400],
        hoverBackgroundColor: tailwindConfig().theme.colors.blue[500],
        barPercentage: 0.66,
        categoryPercentage: 0.66,
        minBarLength: 2,
      },
    ],
  };

  return (
    <div className="bg-white shadow-lg rounded-md">
      <header className="px-5 py-4">
        <div className="flex justify-between items-center">
          <div>
            <h2 className="font-semibold text-maincolor text-md">
              Orders per Day
            </h2>
            <p className="text-sm text-gray-600">{getMonthName(currentMonth)}</p>
          </div>
          <div className="flex items-center space-x-2">
            <button
              onClick={prevMonth}
              className="p-2 rounded-md bg-gray-100 hover:bg-gray-200 text-gray-700"
              aria-label="Previous month"
            >
              &larr;
            </button>
            <button
              onClick={goToCurrentMonth}
              className="px-3 py-1 text-sm rounded-md bg-maincolor text-white hover:bg-thirdcolor"
            >
              Current Month
            </button>
            <button
              onClick={nextMonth}
              className="p-2 rounded-md bg-gray-100 hover:bg-gray-200 text-gray-700"
              aria-label="Next month"
            >
              &rarr;
            </button>
          </div>
        </div>
        <div className="mt-2 text-lg font-bold text-black">
          Total Orders: {totalOrders}
        </div>
      </header>
      {/* Chart built with Chart.js 3 */}
      <div className="grow">
         <BarChart data={chartData} width={595} height={248} unit="day" valueFormatter={integerFormatter} showAllTicks={true} />
      </div>
    </div>
  );
};

export default OrdersBarChartCard;