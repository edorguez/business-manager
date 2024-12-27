"use client";

import DoughnutChart from "../charts/DoughnutChart";

export interface DoughnutChartCardProps {
  title: string;
  labels: string[];
  data: number[];
}

const DoughnutChartCard: React.FC<DoughnutChartCardProps> = ({
  title,
  labels,
  data,
}) => {
  const chartData = {
    labels: labels,
    datasets: [
      {
        label: title,
        data: data,
        backgroundColor: ["#14A098", "#CB2D6F", "#501F3A", "#0F292F"],
        hoverBackgroundColor: ["#0e716d", "#922050", "#250e1b", "#061113"],
      },
    ],
  };

  return (
    <div className="bg-white shadow-lg rounded-md">
      <header className="px-5 py-4">
        <h2 className="font-semibold text-md text-maincolor">{title}</h2>
      </header>
      {data?.length === 0 && (
        <div className="flex justify-center">
          <span>No hay registros para mostrar</span>
        </div>
      )}
      {data && <DoughnutChart data={chartData} width={389} height={260} />}
    </div>
  );
};

export default DoughnutChartCard;
