export type CreateOrder = {
  companyId: number;
  customer: CreateOrderCustomer;
  products: CreateOrderProduct[];
};

export type CreateOrderCustomer = {
  firstName: string;
  lastName: string;
  identificationNumber: string;
  identificationType: string;
  phone: string;
};

export type CreateOrderProduct = {
  productId: number;
  quantity: number;
  price: number;
};
