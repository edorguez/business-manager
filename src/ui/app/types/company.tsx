export type Company = {
  id: number;
  name: string;
  imageUrl: string;
  lastPaymentDate: Date;
}

export type GetCompany = {
  id: number;
}

export type EditCompany = {
  id: number;
  name: string;
  imageUrl: string;
}