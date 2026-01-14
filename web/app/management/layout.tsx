"use client";

import Header from "../components/Header";
import HeaderMobile from "../components/HeaderMobile";
import MarginWidthWrapper from "../components/MarginWidthWrapper";
import PageWrapper from "../components/PageWrapper";
import SideNav from "../components/SideNav";
import { PuffLoader } from "react-spinners";
import useLoading from "../hooks/useLoading";
import Image from "next/image";

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
          <div className="relative">
            <PuffLoader
              color={"#000000"}
              loading={isLoading.isLoading}
              size={150}
              aria-label="Loading Spinner"
              data-testid="loader"
            />

            <div className="absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2">
              <Image
                src="/images/logo.png"
                width={512}
                height={512}
                alt="Loading"
                className="rounded-full w-16 h-16 object-cover border-2 border-white"
              />
            </div>
          </div>
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
