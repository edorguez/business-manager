"use client";

import { formatValue, tailwindConfig } from "@/app/utils/Utils";
import {
  BarController,
  BarElement,
  Chart,
  Legend,
  LinearScale,
  TimeScale,
  Tooltip,
  TimeUnit,
} from "chart.js";
import React, { useEffect, useRef, useState } from "react";
import { chartColors } from "./ChartjsConfig";

interface BarChartProps {
  data: any;
  width: number;
  height: number;
  unit?: TimeUnit;
  valueFormatter?: (value: any) => string;
  showAllTicks?: boolean;
}

Chart.register(
  BarController,
  BarElement,
  LinearScale,
  TimeScale,
  Tooltip,
  Legend
);

const BarChart: React.FC<BarChartProps> = ({ 
  data, 
  width, 
  height, 
  unit = 'month', 
  valueFormatter = formatValue, 
  showAllTicks = false 
}) => {
  const [chart, setChart] = useState<Chart | null>(null);
  const canvas = useRef<HTMLCanvasElement>(null);
  const legend = useRef<HTMLUListElement>(null);
  
  const {
    textColor,
    gridColor,
    tooltipBodyColor,
    tooltipBgColor,
    tooltipBorderColor,
  } = chartColors;

  useEffect(() => {
    if (!canvas.current) return;

    const ctx = canvas.current;
    
    const newChart = new Chart(ctx, {
      type: "bar",
      data: data,
      options: {
        layout: {
          padding: { top: 12, bottom: 16, left: 20, right: 20 },
        },
        scales: {
          y: {
            border: { display: false },
            ticks: {
              maxTicksLimit: 5,
              callback: (value) => valueFormatter(value),
              color: textColor.dark,
            },
            grid: { color: gridColor.dark },
          },
          x: {
            type: "time",
            time: {
              parser: "MM-DD-YYYY",
              unit: unit,
              displayFormats: { month: "MMM YY", day: "MMM DD" },
            },
            border: { display: false },
            grid: { display: false },
            ticks: {
              color: textColor.dark,
              ...(unit === 'day' && showAllTicks ? { autoSkip: false, maxTicksLimit: 31 } : {}),
            },
          },
        },
        plugins: {
          legend: { display: false },
          tooltip: {
            callbacks: {
              label: (context) => valueFormatter(context.parsed.y),
            },
            bodyColor: tooltipBodyColor.dark,
            backgroundColor: tooltipBgColor.dark,
            borderColor: tooltipBorderColor.dark,
          },
        },
        interaction: { intersect: false, mode: "nearest" },
        animation: { duration: 500 },
        maintainAspectRatio: false,
        resizeDelay: 200,
      },
      plugins: [
        {
          id: "htmlLegend",
          afterUpdate(c: any) {
            const ul = legend.current;
            if (!ul) return;
            while (ul.firstChild) { ul.firstChild.remove(); }
            const items = c.options.plugins.legend.labels.generateLabels(c);
            items.forEach((item: any) => {
              const li = document.createElement("li");
              li.style.marginRight = tailwindConfig().theme.margin[4];
              const button = document.createElement("button");
              button.style.display = "inline-flex";
              button.style.alignItems = "center";
              button.style.opacity = item.hidden ? ".3" : "";
              button.onclick = () => {
                c.setDatasetVisibility(item.datasetIndex, !c.isDatasetVisible(item.datasetIndex));
                c.update();
              };
              const box = document.createElement("span");
              box.style.display = "block";
              box.style.width = tailwindConfig().theme.width[3];
              box.style.height = tailwindConfig().theme.height[3];
              box.style.borderRadius = "9999px";
              box.style.marginRight = tailwindConfig().theme.margin[2];
              box.style.borderWidth = "3px";
              box.style.borderColor = item.fillStyle;
              const labelContainer = document.createElement("span");
              labelContainer.style.display = "flex";
              labelContainer.style.alignItems = "center";
              const value = document.createElement("span");
              value.classList.add("text-black");
              value.style.fontWeight = "bold";
              value.style.marginRight = tailwindConfig().theme.margin[2];
              const totalValue = c.data.datasets[item.datasetIndex].data.reduce((a: number, b: number) => a + b, 0);
              value.innerText = valueFormatter(totalValue);
              const label = document.createElement("span");
              label.innerText = item.text;
              labelContainer.appendChild(value);
              labelContainer.appendChild(label);
              button.appendChild(box);
              button.appendChild(labelContainer);
              li.appendChild(button);
              ul.appendChild(li);
            });
          },
        },
      ],
    });

    setChart(newChart);
    return () => {
      newChart.destroy();
    };
  }, [data, unit, showAllTicks, gridColor.dark, textColor.dark, tooltipBgColor.dark, tooltipBodyColor.dark, tooltipBorderColor.dark, valueFormatter]);

  return (
    <React.Fragment>
      <div className="px-5 py-3">
        <ul ref={legend} className="flex flex-wrap"></ul>
      </div>
      <div className="grow">
        <canvas ref={canvas} width={width} height={height}></canvas>
      </div>
    </React.Fragment>
  );
};

export default BarChart;
