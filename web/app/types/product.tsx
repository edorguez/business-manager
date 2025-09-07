export type Product = {
  id: string;
  name: string;
  description: string;
  images: string[];
  sku: string;
  quantity: number;
  price: number;
  productStatus: number;
}

export type CreateProduct = {
  companyId: number;
  name: string;
  description: string;
  sku: string;
  quantity: number | undefined;
  price: number | undefined;
  productStatus: number;
}

export type EditProduct = {
  id: string;
  companyId: number;
  name: string;
  description: string;
  sku: string;
  quantity: number | undefined;
  price: number | undefined;
  productStatus: number;
}

export type GetProduct = {
  id: string;
}

export type GetProducts = {
  companyId: number;
  name: string;
  sku: string;
  productStatus: number | undefined;
  limit: number;
  offset: number;
}

export type GetLatestProducts = {
  companyId: number;
  limit: number;
}

export type SearchProduct = {
  name: string;
  sku: string;
}

export type DeleteProduct = {
  id: string;
}

export type ChangeStatusProduct = {
  id: string;
  productStatus: number;
}
