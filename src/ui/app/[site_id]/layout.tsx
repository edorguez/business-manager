"use client";

import { Inter } from "next/font/google";
import { ReactNode, useCallback, useEffect, useState } from "react";
import { MoonLoader } from "react-spinners";
import { Company } from "../types/company";
import { GetCompanyByNameRequest } from "../services/companies";
import { useRouter } from "next/navigation";

const inter = Inter({ subsets: ["latin"] });

export default function SiteLayout({
  params,
  children,
}: {
  params: { site_id: string };
  children: ReactNode;
}) {
  const router = useRouter();
  const [isLoading, setIsLoading] = useState<boolean>(true);
  
  const getCompany = useCallback(async () => {
    let company: Company = await GetCompanyByNameRequest(params.site_id);
    console.log(company);
     if(!company?.id) {
      // I need to create my not found route
      router.push('/404')
     }
  }, [params.site_id, router]);

  useEffect(() => {
    getCompany();
  }, [getCompany])

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
      
      {!isLoading && (
        <div className={inter.className}>
          {children}
        </div>
      )}
    </>
  );
}
