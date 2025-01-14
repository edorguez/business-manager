"use client";

import { ReactNode, useEffect, useState } from "react";
import { MoonLoader } from "react-spinners";
import React from "react";
import useGeneralLoading from "../hooks/useGeneralLoading";


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
  }, [generalLoading])

  return (
    <>
      {isLoading && (
        <div className="w-screen h-screen bg-white absolute z-50 flex justify-center items-center">
          <MoonLoader
            color={"#14A098"}
            loading={isLoading}
            size={150}
            aria-label="Loading Spinner"
            data-testid="loader"
          />
        </div>
      )}
      
          <div className={`${isLoading ? 'hidden' : ''}`}>
            {children}
          </div>
    </>
  );
}
