"use client";

import BarChart from "../charts/BarChart";
import { tailwindConfig, convertToTimezone, generateMonthList, convertTimestampToDate } from "@/app/utils/Utils";
import { GetOrdersByMonthRequest } from "@/app/services/orders";
import getCurrentUser from "@/app/actions/getCurrentUser";
import { CurrentUser } from "@/app/types/auth";
import { Select } from "@chakra-ui/react";
import { useEffect, useState } from "react";

const OrdersBarChartCard = () => {
  const monthOptions = generateMonthList(12);
  const [selectedMonthValue, setSelectedMonthValue] = useState<string>(() => {
    // Default to current month in YYYY-MM format
    const now = new Date();
    return `${now.getFullYear()}-${String(now.getMonth() + 1).padStart(2, '0')}`;
  });
  const [ordersData, setOrdersData] = useState<(number | null)[]>([]);
  const [labels, setLabels] = useState<Date[]>([]);
  const [totalOrders, setTotalOrders] = useState<number>(0);

  // Compute current month Date from selected value
  const currentMonthDate = (() => {
    const [year, month] = selectedMonthValue.split('-').map(Number);
    return new Date(year, month - 1, 1);
  })();

  const getMonthName = (date: Date): string => {
    const monthYear = date.toLocaleString('es-ES', { month: 'long', year: 'numeric' });
    return monthYear[0].toUpperCase() + monthYear.substring(1);
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
    if (value === null || value === undefined) return '';
    return Math.round(Number(value)).toString();
  };

  const handleMonthChange = (event: React.ChangeEvent<HTMLSelectElement>) => {
    setSelectedMonthValue(event.target.value);
  };

  useEffect(() => {
    const fetchOrdersByMonth = async () => {
      const currentUser: CurrentUser | null = getCurrentUser();
      if (!currentUser) {
        return;
      }

      const [year, month] = selectedMonthValue.split('-').map(Number);
      const daysInMonth = getDaysInMonth(year, month - 1);
      // Initialize with null for gaps (no bar for zero orders)
      const dayCounts: (number | null)[] = new Array(daysInMonth.length).fill(null);

      try {
        const response = await GetOrdersByMonthRequest({
          companyId: currentUser.companyId,
          year,
          month,
        });

        if (response && response.status !== 200) {
          console.error('Error fetching orders:', response.error);
        }

        if (response && response.createdAt && Array.isArray(response.createdAt)) {
          // Count orders per day
          const userTimezoneOffset = new Date().getTimezoneOffset();
          response.createdAt.forEach((timestampStr: any) => {
            let timestamp: Date = convertTimestampToDate(timestampStr);
            timestamp = convertToTimezone(timestamp, userTimezoneOffset);
            const dayOfMonth = timestamp.getDate() - 1; // zero-index
            if (dayOfMonth >= 0 && dayOfMonth < dayCounts.length) {
              // If currently null, set to 1, else increment
              dayCounts[dayOfMonth] = dayCounts[dayOfMonth] === null ? 1 : (dayCounts[dayOfMonth]! + 1);
            }
          });
        }
      } catch (error) {
        console.error("Failed to fetch orders by month:", error);
      }

      setLabels(daysInMonth);
      setOrdersData(dayCounts);
      // Total orders: sum only non-null values
      const total = dayCounts.reduce<number>((sum, count) => sum + (count || 0), 0);
      setTotalOrders(total);
    };

    fetchOrdersByMonth();
  }, [selectedMonthValue]);

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
        label: "Órdenes",
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
              Órdenes del Mes
            </h2>
            <p className="text-sm text-gray-600">{getMonthName(currentMonthDate)}</p>
          </div>
          <Select
            size="sm"
            value={selectedMonthValue}
            onChange={handleMonthChange}
            variant="default"
            width="180px"
          >
            {monthOptions.map((option) => (
              <option key={option.value} value={option.value}>
                {option.label}
              </option>
            ))}
          </Select>
        </div>
        <div className="mt-2 text-lg font-bold text-black">
          Número de Órdenes: {totalOrders}
        </div>
      </header>
      <div className="grow">
         <BarChart data={chartData} width={389} height={260} unit="day" valueFormatter={integerFormatter} showAllTicks={true} />
      </div>
    </div>
  );
};

export default OrdersBarChartCard;
