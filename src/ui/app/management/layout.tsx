"use client";

import Header from "../components/Header";
import HeaderMobile from "../components/HeaderMobile";
import MarginWidthWrapper from "../components/MarginWidthWrapper";
import PageWrapper from "../components/PageWrapper";
import SideNav from "../components/SideNav";
import { ScaleLoader } from "react-spinners";
import useLoading from "../hooks/useLoading";

export default function ManagementLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  const isLoading = useLoading();
  return (
    <section>
      {isLoading.isLoading && (
        <div className="w-screen h-screen bg-black/50 absolute z-50 flex justify-center items-center">
          <ScaleLoader
            color={"#14A098"}
            loading={isLoading.isLoading}
            height={70}
            width={8}
            aria-label="Loading Spinner"
            data-testid="loader"
          />
        </div>
      )}
      <div className="flex">
        <SideNav />
        <main className="flex-1">
          <MarginWidthWrapper>
            <Header />
            <HeaderMobile />
            <PageWrapper>{children}</PageWrapper>
          </MarginWidthWrapper>
        </main>
      </div>
    </section>
  );
}
