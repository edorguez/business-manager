'use client'

import WelcomeBanner from "./components/banners/WelcomeBanner";
import SimpleLineChartCard from "./components/cards/SimpleLineChartCard";
import DoughnutChartCard from "./components/cards/DoughnutChartCard";
import ListCard from "./components/cards/ListCard";
import BarChartCard from "./components/cards/BarChartCard";
import Container from "./components/Container";

export default function Home() {
  return (
    <Container>
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
    </Container>
  );
}
