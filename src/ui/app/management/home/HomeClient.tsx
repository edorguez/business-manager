'use client';

import WelcomeBanner from "@/app/components/banners/WelcomeBanner";
import BarChartCard from "@/app/components/cards/BarChartCard";
import DoughnutChartCard from "@/app/components/cards/DoughnutChartCard";
import ListCard from "@/app/components/cards/ListCard";
import SimpleLineChartCard from "@/app/components/cards/SimpleLineChartCard";

const HomeClient = () => {
  return (
    <>
      <WelcomeBanner />

      <div className="mt-4 grid grid-cols-1 lg:grid-cols-3 gap-4">
        <SimpleLineChartCard />
        <SimpleLineChartCard />
        <SimpleLineChartCard />
      </div>

      <div className="mt-4">
        <BarChartCard />
      </div>

      <div className="mt-4 grid grid-cols-1 lg:grid-cols-3 gap-4">
        <DoughnutChartCard />
        <ListCard />
      </div>
    </>
  )
}

export default HomeClient;
