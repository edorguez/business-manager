import { PaymentType } from "./paymentType";

export type Payment = {
  id: number;
  companyId: number;
  name: string;
  bank: string;
  accountNumber: string;
  accountType: string;
  identificationNumber: string;
  identificationType: string;
  phone: string;
  email: string;
  paymentTypeId: number;
  isActive: boolean;
  paymentType: PaymentType;
}

export type CreatePayment = {
  companyId: number;
  name: string;
  bank: string;
  accountNumber: string;
  accountType: string;
  identificationNumber: string;
  identificationType: string;
  phone: string;
  email: string;
  paymentTypeId: number;
}

export type GetPayments = {
  companyId: number;
  limit: number;
  offset: number;
}

export type DeletePayment = {
  id: number;
}

export type ChangeStatusPayment = {
  id: number;
  status: boolean;
}