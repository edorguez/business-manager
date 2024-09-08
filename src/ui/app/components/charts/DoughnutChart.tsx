"use client";

import { useEffect, useRef } from "react";
import { chartColors } from "./ChartjsConfig";
import {
  ArcElement,
  Chart,
  DoughnutController,
  TimeScale,
  Tooltip,
} from "chart.js";

interface DoughnutChartProps {
  data: any;
  width: number;
  height: number;
}

Chart.register(DoughnutController, ArcElement, TimeScale, Tooltip);

const DoughnutChart: React.FC<DoughnutChartProps> = ({
  data,
  width,
  height,
}) => {
  const canvas = useRef(null);

  useEffect(() => {
    const ctx: any = canvas.current;
    const newChart: any = new Chart(ctx, {
      type: 'doughnut',
      data: data,
      options: {
        cutout: '80%',
        layout: {
          padding: 44,
        },
        plugins: {
          legend: {
            display: false,
          },
        },
        interaction: {
          mode: 'nearest',
        },
        animation: {
          duration: 500,
        },
        // maintainAspectRatio: false,
        resizeDelay: 200,
      },
    });
    return () => newChart.destroy();
  }, []);

  return (
    <div className="grow flex flex-col justify-center">
      <div>
        <canvas ref={canvas} width={width} height={height}></canvas>
      </div>
    </div>
  );
};

export default DoughnutChart;
