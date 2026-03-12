export type Company = {
  id: number;
  name: string;
  nameFormatUrl: string;
  imageUrl: string;
  isFreeTrial: boolean;
  planId: boolean
  lastPaymentDate: Date;
}

export type GetCompany = {
  id: number;
}

export type EditCompany = {
  id: number;
  name: string;
  nameFormatUrl: string;
}
