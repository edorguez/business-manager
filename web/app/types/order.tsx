import { Customer } from "./customer";

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
  productId: string;
  name: string;
  quantity: number;
  price: number;
};

export type GetOrders = {
  companyId: number;
  limit: number;
  offset: number;
}

export type GetOrdersResponse = {
  orders: OrderDetails[];
  total: number;
}

export type Order = {
  id: number;
  companyId: number;
  customerId: number;
  createdAt: any;
}

export type OrderProduct = {
  id: number;
  orderId: number;
  productId: string;
  imageUrl: string;
  name: string;
  quantity: number;
  price: number;
}

export type OrderDetails = {
  order: Order;
  customer: Customer;
  products: OrderProduct[];
}

export type OrdersTable = {
  id: number;
  fullName: string;
  identificationNumber: string;
  date: string;
  products: string;
}

export type OrderProductsTable = {
  imageUrl: string;
  name: string;
  quantity: number;
  price: number;
  total: number;
}
