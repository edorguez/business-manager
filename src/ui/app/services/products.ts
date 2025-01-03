import {
  ChangeStatusProduct,
  CreateProduct,
  DeleteProduct,
  EditProduct,
  GetLatestProducts,
  GetProduct,
  GetProducts,
} from "@/app/types/product";
import Cookies from "js-cookie";

const baseUrl: string = "http://localhost:3001/api/products";

export async function CreateProductRequest(request: CreateProduct) {
  try {
    const headers = new Headers();
    const token = Cookies.get("token");
    headers.append("Authorization", <string>token);

    if (request.price) request.price = +request.price;
    if (request.quantity) request.quantity = +request.quantity;

    const formData = new FormData();

    // Add JSON data as a string
    formData.append("json", JSON.stringify(request));

    // Add images to the FormData
    request.images.forEach((image, index) => {
      formData.append(`files`, image);
    });

    console.log('check')
    console.log(request.images)
    console.log(formData);

    const res = await fetch(baseUrl, {
      method: "POST",
      headers: headers,
      body: formData,
    });

    let response = await res.json();

    return response;
  } catch (error: any) {
    console.log(error.toString());
  }
}

export async function EditProductRequest(request: EditProduct) {
  try {
    const headers = new Headers();
    const token = Cookies.get("token");
    headers.append("Authorization", <string>token);

    if (request.price) request.price = +request.price;
    if (request.quantity) request.quantity = +request.quantity;

    const res = await fetch(`${baseUrl}/${request.id}`, {
      method: "PUT",
      headers: headers,
      body: JSON.stringify(request),
    });
  } catch (error: any) {
    console.log(error.toString());
  }
}

export async function ChangeStatusRequest(request: ChangeStatusProduct) {
  try {
    const headers = new Headers();
    const token = Cookies.get("token");
    headers.append("Authorization", <string>token);

    await fetch(`${baseUrl}/${request.id}/status`, {
      method: "PUT",
      headers: headers,
      body: JSON.stringify(request),
    });
  } catch (error: any) {
    console.log(error.toString());
  }
}

export async function GetProductRequest(request: GetProduct) {
  try {
    const headers = new Headers();
    const token = Cookies.get("token");
    headers.append("Authorization", <string>token);

    const res = await fetch(`${baseUrl}/${request.id}`, {
      method: "GET",
      headers: headers,
    });

    let response = await res.json();

    return response;
  } catch (error: any) {
    console.log(error.toString());
  }
}

export async function GetProductsRequest(request: GetProducts) {
  try {
    const headers = new Headers();
    const token = Cookies.get("token");
    headers.append("Authorization", <string>token);

    let params: any = {
      companyId: request.companyId.toString(),
      name: request.name.trim(),
      sku: request.sku.trim(),
      limit: request.limit.toString(),
      offset: request.offset.toString(),
    };

    if (request.productStatus) {
      params.productStatus = request.productStatus.toString();
    }

    const res = await fetch(
      `${baseUrl}?` + new URLSearchParams(params).toString(),
      {
        method: "GET",
        headers: headers,
      }
    );

    let response = await res.json();

    return response;
  } catch (error: any) {
    console.log(error.toString());
  }
}

export async function GetLatestProductsRequest(request: GetLatestProducts) {
  try {
    const headers = new Headers();
    const token = Cookies.get("token");
    headers.append("Authorization", <string>token);

    const res = await fetch(
      `${baseUrl}/latest?` +
        new URLSearchParams({
          companyId: request.companyId.toString(),
          limit: request.limit.toString(),
        }).toString(),
      {
        method: "GET",
        headers: headers,
      }
    );

    let response = await res.json();

    return response;
  } catch (error: any) {
    console.log(error.toString());
  }
}

export async function DeleteProductRequest(request: DeleteProduct) {
  try {
    const headers = new Headers();
    const token = Cookies.get("token");
    headers.append("Authorization", <string>token);

    await fetch(`${baseUrl}/${request.id}`, {
      method: "DELETE",
      headers: headers,
    });
  } catch (error: any) {
    console.log(error.toString());
  }
}
