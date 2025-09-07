"use client";

import LineChartSimple from "../charts/LineChartSimple";
import { hexToRGB, tailwindConfig } from "@/app/utils/Utils";

interface SimpleLineChartCardProps {
  labels: any[];
  data: any[];
  title: string;
  subtitle: string;
  total: string;
}

const SimpleLineChartCard: React.FC<SimpleLineChartCardProps> = ({
  labels,
  data,
  title,
  subtitle,
  total,
}) => {
  const chartData = {
    labels: labels,
    datasets: [
      // Indigo line
      {
        data: data,
        fill: true,
        backgroundColor: 'rgba(203, 45, 111, 0.08)',
        borderColor: 'rgb(20, 160, 155)',
        borderWidth: 2,
        tension: 0,
        pointRadius: 0,
        pointHoverRadius: 3,
        pointBackgroundColor: 'rgba(203, 45, 111, 1)',
        pointHoverBackgroundColor: 'rgba(203, 45, 111, 1)',
        pointBorderWidth: 0,
        pointHoverBorderWidth: 0,
        clip: 20,
      },
      // Gray line
      // {
      //   data: [
      //     532, 532, 532, 404, 404, 314, 314, 314, 314, 314, 234, 314, 234, 234,
      //     314, 314, 314, 388, 314, 202, 202, 202, 202, 314, 720, 642,
      //   ],
      //   borderColor: `rgba(${hexToRGB(
      //     tailwindConfig().theme.colors.slate[500]
      //   )}, 0.25)`,
      //   borderWidth: 2,
      //   tension: 0,
      //   pointRadius: 0,
      //   pointHoverRadius: 3,
      //   pointBackgroundColor: `rgba(${hexToRGB(
      //     tailwindConfig().theme.colors.slate[500]
      //   )}, 0.25)`,
      //   pointHoverBackgroundColor: `rgba(${hexToRGB(
      //     tailwindConfig().theme.colors.slate[500]
      //   )}, 0.25)`,
      //   pointBorderWidth: 0,
      //   pointHoverBorderWidth: 0,
      //   clip: 20,
      // },
    ],
  };

  return (
    <div className="bg-white shadow-lg rounded-md">
      <div className="px-5 pt-5">
        <h2 className="text-md font-semibold text-maincolor mb-2">{title}</h2>
        <div className="text-xs font-semibold text-black uppercase mb-1">
          {subtitle}
        </div>
        <div className="flex items-start">
          <div className="text-2xl font-bold text-black mr-2">{total}</div>
          {/* <div className="text-sm font-semibold text-white px-1.5 bg-emerald-500 rounded-full">
            +49%
          </div> */}
        </div>
      </div>
      {/* Chart built with Chart.js 3 */}
      <div className="grow max-sm:max-h-[128px] xl:max-h-[128px]">
        {/* Change the height attribute to adjust the chart height */}
        <LineChartSimple data={chartData} width={389} height={128} />
      </div>
    </div>
  );
};

export default SimpleLineChartCard;
