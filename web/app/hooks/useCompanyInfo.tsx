import { create } from "zustand";
import { Company } from "../types/company";

interface CompanyInfoStore {
  company: Company | null;
  setCompany: (company: Company) => void;
}

const useCompanyInfo = create<CompanyInfoStore>((set) => ({
  company: null,
  setCompany: (company: Company) => {
    set((state) => {
      return { ...state, company };
    });
  }
}));

export default useCompanyInfo;