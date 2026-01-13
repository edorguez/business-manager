"use client";

import { Button } from "@chakra-ui/react";
import { useCallback, useEffect, useState } from "react";
import { Company } from "../types/company";
import { Product } from "../types/product";
import ProductCard from "../components/cards/ProductCard";
import { GetCompanyByNameUrlRequest } from "../services/companies";
import { notFound, useParams, useRouter } from "next/navigation";
import useGeneralLoading from "../hooks/useGeneralLoading";
import { GetProductsRequest } from "../services/products";
import ProductsCartDrawer from "../components/drawers/ProductsCartDrawer";
import useProductsCart from "@/app/hooks/useProductsCart";
import useCompanyInfo from "../hooks/useCompanyInfo";
import NotFound from "../components/NotFound";

const SitePage = () => {
  const cart = useProductsCart();
  const companyInfo = useCompanyInfo();
  const router = useRouter();
  const params = useParams();
  const isLoading = useGeneralLoading();
  const [company, setCompany] = useState<Company | null>(null);
  const [isNotFound, setIsNotFound] = useState<boolean>(false);
  const [products, setProducts] = useState<Product[]>([]);

  const getCompany = useCallback(async () => {
    let getCompany: Company = await GetCompanyByNameUrlRequest(
      params.site_id.toString()
    );
    if (!getCompany?.id || getCompany?.lastPaymentDate < new Date()) {
      setIsNotFound(true);
    } else {
      setCompany(getCompany);
      companyInfo.setCompany(getCompany);
      getProductsByCompanyId(getCompany);
    }
    isLoading.onEndLoading();
  }, [params.site_id, router]);

  const getProductsByCompanyId = useCallback(async (comp: Company) => {
    let getProducts: Product[] = await GetProductsRequest({
      companyId: comp.id,
      name: "",
      sku: "",
      productStatus: 1,
      limit: 10,
      offset: 0,
    });

    setProducts(getProducts);
  }, []);

  useEffect(() => {
    getCompany();
  }, [getCompany]);

  const getTotalItems = (): number => {
    return cart.items.reduce((total, item) => total + item.quantity, 0);
  };

   if (isNotFound) {
    return <NotFound />;
  }

  return (
    <div>
      <div className="min-h-screen bg-gray-200">
        {/* Header */}
        <header className="bg-white shadow-md">
          <div className="container mx-auto px-4 py-4 flex items-center justify-between">
            <h1 className="text-2xl font-bold text-defaultcolor">
              {company?.name}
            </h1>
            <div className="flex items-center space-x-4">
              <Button variant="default" onClick={cart.onOpen}>
                Carrito ({getTotalItems()})
              </Button>
            </div>
          </div>
        </header>

        {/* Main Content */}
        <main className="container mx-auto px-4 py-8">
          {/* <h2 className="text-2xl font-semibold mb-6">Our Menu</h2> */}
          <div className="grid grid-cols-2 lg:grid-cols-5 gap-4">
            {products.map((product) => (
              <ProductCard
                key={product.id}
                product={product}
                onAddToCard={() => {
                  cart.onAddToCart(product);
                }}
              />
            ))}
            {
              products?.length <= 0 && (
                <div className="col-span-full text-center w-full">
                  <h2 className="text-2xl font-semibold mb-6">No hay productos disponibles</h2>
                </div>
              )
            }
          </div>
        </main>

        {/* Shopping Cart Drawer */}
        <ProductsCartDrawer />
      </div>
    </div>
  );
};

export default SitePage;
