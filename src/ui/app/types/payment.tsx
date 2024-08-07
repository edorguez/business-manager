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