export type Product = {
  id: number;
  name: string;
  description: string;
  images: string[];
  sku: string;
  quantity: number;
  price: number;
}

export type CreateProduct = {
  companyId: number;
  name: string;
  description: string;
  images: string[];
  sku: string;
  quantity: number;
  price: number;
  productStatus: number;
}

export type EditProduct = {
  id: number;
  companyId: number;
  name: string;
  description: string;
  images: string[];
  sku: string;
  quantity: number;
  price: number;
  productStatus: number;
}

export type GetProduct = {
  id: number;
}

export type GetProducts = {
  companyId: number;
  limit: number;
  offset: number;
}

export type SearchProduct = {
  name: string;
  sku: string;
}

export type DeleteProduct = {
  id: number;
}
