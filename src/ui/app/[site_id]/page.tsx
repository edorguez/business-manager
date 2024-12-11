"use client";

import {
  Button,
} from "@chakra-ui/react";
import { useCallback, useEffect, useState } from "react";
import { Company } from "../types/company";
import { Product } from "../types/product";
import ProductCard from "../components/cards/ProductCard";
import { GetCompanyByNameRequest } from "../services/companies";
import { useParams, useRouter } from "next/navigation";
import useGeneralLoading from "../hooks/useGeneralLoading";
import { GetProductsRequest } from "../services/products";
import ProductsCartDrawer from "../components/drawers/ProductsCartDrawer";
import useProductsCart from "@/app/hooks/useProductsCart";

const SitePage = () => {
  const cart = useProductsCart();
  const router = useRouter();
  const params = useParams();
  const isLoading = useGeneralLoading();
  const [company, setCompany] = useState<Company | null>(null);
  const [products, setProducts] = useState<Product[]>([]);
  
  const getCompany = useCallback(async () => {
    let getCompany: Company = await GetCompanyByNameRequest(params.site_id.toString());
     if(!getCompany?.id || getCompany?.lastPaymentDate < new Date()) {
      console.log('EPA FUERA');
      // I need to create my not found route
      router.push('/404')
     } else {
      setCompany(getCompany);
      getProductsByCompanyId(getCompany);
      
      isLoading.onEndLoading();
     }
  }, [params.site_id, router]);

  const getProductsByCompanyId = useCallback(async (comp: Company) => {
    let getProducts: Product[] = await GetProductsRequest({
      companyId: comp.id,
      name: '',
      sku: '',
      limit: 10,
      offset: 0
    });

    setProducts(getProducts);
  }, []);

  useEffect(() => {
    getCompany();
  }, [getCompany])

  const getTotalItems = (): number => {
    return cart.items.reduce((total, item) => total + item.quantity, 0);
  };

  return (
    <div>
      <div className="min-h-screen bg-gray-100">
        {/* Header */}
        <header className="bg-white shadow-md">
          <div className="container mx-auto px-4 py-4 flex items-center justify-between">
            <h1 className="text-2xl font-bold text-green-500">
              { company?.name }
            </h1>
            <div className="flex items-center space-x-4">
              <Button variant="main" onClick={cart.onOpen}>
                Carrito ({getTotalItems()})
              </Button>
            </div>
          </div>
        </header>

        {/* Main Content */}
        <main className="container mx-auto px-4 py-8">
          {/* <h2 className="text-2xl font-semibold mb-6">Our Menu</h2> */}
          <div className="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-8">
            {products.map((product) => (
              <ProductCard key={product.id} product={product} onAddToCard={() => { cart.onAddToCart(product) }} />
            ))}
          </div>
        </main>

        {/* Shopping Cart Drawer */}
        <ProductsCartDrawer />
      </div>
    </div>
  );
}

export default SitePage;