export type Customer = {
  id: number;
  firstName: string;
  lastName: string;
  identificationNumber: string;
  phone: string;
  email: string;
}

export type CreateCustomer = {
  companyId: number;
  firstName: string;
  lastName: string;
  identificationNumber: string;
  identificationType: string;
  phone: string;
  email: string;
}
