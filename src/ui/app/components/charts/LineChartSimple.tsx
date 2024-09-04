"use client";

import { useEffect, useRef, useState } from "react";
import { chartColors } from "./ChartjsConfig";
import { Chart, ChartData, Filler, LineController, LineElement, LinearScale, PointElement, TimeScale, Tooltip, TooltipItem } from "chart.js";
import { formatTitleValue, formatValue } from "@/app/utils/Utils";
import 'chartjs-adapter-moment';

interface LineChartSimpleProps {
  data: any,
  width: number,
  height: number
}

Chart.register(LineController, LineElement, Filler, PointElement, LinearScale, TimeScale, Tooltip);

const LineChartSimple: React.FC<LineChartSimpleProps> = ({
  data,
  width,
  height
}) => {
  const [chart, setChart] = useState<any>(null);
  const canvas = useRef(null);
  const { tooltipBodyColor, tooltipBgColor, tooltipBorderColor, chartAreaBg } = chartColors;

  useEffect(() => {
    const ctx: any = canvas.current;
    // eslint-disable-next-line no-unused-vars
    const newChart: any = new Chart(ctx, {
      type: "line",
      data: data,
      options: {
        backgroundColor: chartAreaBg.light,
        // chartArea: {
        //   backgroundColor: chartAreaBg.dark,
        // },
        layout: {
          padding: 20,
        },
        scales: {
          y: {
            display: false,
            beginAtZero: false,
          },
          x: {
            type: "time",
            time: {
              parser: "MM-DD-YYYY",
              unit: "month",
            },
            display: false,
          },
        },
        plugins: {
          tooltip: {
            callbacks: {
              title: (tooltipItems: TooltipItem<'line'>[]): string => {
                // Access the first tooltip item
                const tooltipItem: any = tooltipItems[0];
                // Use the label from the data
                const label = tooltipItem.label;
                // Return your custom title
                return formatTitleValue(label);
            },
              //title: () => false, // Disable tooltip title
              // label: (context) => formatValue(context.parsed.y),
            },
            bodyColor: tooltipBodyColor.dark,
            backgroundColor: tooltipBgColor.dark,
            borderColor: tooltipBorderColor.dark,
          },
          legend: {
            display: false,
          },
        },
        interaction: {
          intersect: false,
          mode: "nearest",
        },
        maintainAspectRatio: false,
        resizeDelay: 200,
      },
    });
    setChart(newChart);
    return () => newChart.destroy();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);

  return <canvas ref={canvas} width={width} height={height}></canvas>;
};

export default LineChartSimple;
