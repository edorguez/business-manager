export type Customer = {
  id: number;
  firstName: string;
  lastName: string;
  identificationNumber: string;
  identificationType: string;
  phone: string;
  email: string;
}

export type CustomerByMonth = {
  dates: Date[];
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

export type EditCustomer = {
  id: number;
  companyId: number;
  firstName: string;
  lastName: string;
  identificationNumber: string;
  identificationType: string;
  phone: string;
  email: string;
}

export type GetCustomer = {
  id: number;
}

export type GetCustomers = {
  companyId: number;
  name: string;
  lastName: string;
  identificationNumber: string;
  limit: number;
  offset: number;
}

export type GetCustomersByMonths = {
  companyId: number;
  months: number;
}

export type SearchCustomer = {
  name: string;
  lastName: string;
  identificationNumber: string;
}

export type DeleteCustomer = {
  id: number;
}

export type CustomerMonths = {
  oneMonth: { labels: Date[], data: number[] };
  twoMonths: { labels: Date[], data: number[] };
  threeMonths: { labels: Date[], data: number[] };
}