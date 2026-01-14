"use client";

import { ReactNode, useEffect, useState } from "react";
import { PuffLoader } from "react-spinners";
import React from "react";
import useGeneralLoading from "../hooks/useGeneralLoading";
import Image from "next/image";

export default function SiteLayout({
  // params,
  children,
}: {
  params: { site_id: string };
  children: ReactNode;
}) {
  const generalLoading = useGeneralLoading();
  const [isLoading, setIsLoading] = useState<boolean>(true);

  useEffect(() => {
    setIsLoading(generalLoading.isLoading);
  }, [generalLoading]);

  return (
    <>
      {isLoading && (
        <div className="w-screen h-screen bg-white absolute z-50 flex justify-center items-center">
          <div className="relative">
            <PuffLoader
              color={"#000000"}
              loading={isLoading}
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
      <div className={isLoading ? "hidden" : ""}>{children}</div>
    </>
  );
}
