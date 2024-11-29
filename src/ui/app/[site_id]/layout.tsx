"use client";

import { Inter } from "next/font/google";
import { createContext, ReactNode, useCallback, useContext, useEffect, useState } from "react";
import { MoonLoader } from "react-spinners";
import { Company } from "../types/company";
import { GetCompanyByNameRequest } from "../services/companies";
import { useRouter } from "next/navigation";
import React from "react";
import useGeneralLoading from "../hooks/useGeneralLoading";

const inter = Inter({ subsets: ["latin"] });

// const DataContext = createContext<Company | undefined>(undefined);

// export const useData = (): Company => {
//   const context = useContext(DataContext);
//   if (!context) {
//     throw new Error("useData must be used within a DataContextProvider");
//   }
//   return context;
// };

export default function SiteLayout({
  // params,
  children,
}: {
  params: { site_id: string };
  children: ReactNode;
}) {
  const generalLoading = useGeneralLoading();
  const [isLoading, setIsLoading] = useState<boolean>(true);
  // const router = useRouter();
  // const [isLoading, setIsLoading] = useState<boolean>(true);
  // const [company, setCompany] = useState<Company | null>(null);
  
  // const getCompany = useCallback(async () => {
  //   let getCompany: Company = await GetCompanyByNameRequest(params.site_id);
  //   console.log(getCompany);
  //    if(!getCompany?.id || getCompany?.lastPaymentDate < new Date()) {
  //     console.log('EPA FUERA');
  //     // I need to create my not found route
  //     router.push('/404')
  //    } else {
  //     setCompany(getCompany)
  //     setIsLoading(false);
  //    }
  // }, [params.site_id, router]);

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
      
        {/* <DataContext.Provider value={company!}> */}
          <div className={`${isLoading ? 'hidden' : ''}`}>
            {children}
          </div>
        {/* </DataContext.Provider> */}
    </>
  );
}
