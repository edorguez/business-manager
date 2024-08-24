import { ChangeStatusProduct, CreateProduct, DeleteProduct, EditProduct, GetProduct, GetProducts } from "@/app/types/product";

const baseUrl: string = 'http://localhost:3001/api/products';

export async function CreateProductRequest(
  request: CreateProduct
) {
  try {
    const headers = new Headers();
    headers.append("Authorization", <string>localStorage.getItem('token'));

    request.price = +request.price;
    request.quantity = +request.quantity;

    const res = await fetch(baseUrl, {
      method: 'POST',
      headers: headers,
      body: JSON.stringify(request),
    });

    let response = await res.json();

    return response;
  } catch (error: any) {
    console.log(error.toString())
  }
}

export async function EditProductRequest(
  request: EditProduct
) {
  try {
    const headers = new Headers();
    headers.append("Authorization", <string>localStorage.getItem('token'));

    request.price = +request.price;
    request.quantity = +request.quantity;

    const res = await fetch(`${baseUrl}/${request.id}`, {
      method: 'PUT',
      headers: headers,
      body: JSON.stringify(request),
    });
  } catch (error: any) {
    console.log(error.toString())
  }
}

export async function ChangeStatusRequest(
  request: ChangeStatusProduct
) {
  try {
    const headers = new Headers();
    headers.append("Authorization", <string>localStorage.getItem('token'));

    await fetch(`${baseUrl}/${request.id}/status`, {
      method: 'PUT',
      headers: headers,
      body: JSON.stringify(request),
    });
  } catch (error: any) {
    console.log(error.toString());
  }
}

export async function GetProductRequest(
  request: GetProduct
) {
  try {
    const headers = new Headers();
    headers.append("Authorization", <string>localStorage.getItem('token'));

    const res = await fetch(`${baseUrl}/${request.id}`, {
      method: 'GET',
      headers: headers,
    });

    let response = await res.json();

    return response;
  } catch (error: any) {
    console.log(error.toString())
  }
}

export async function GetProductsRequest(
  request: GetProducts
) {
  try {
    const headers = new Headers();
    headers.append("Authorization", <string>localStorage.getItem('token'));

    const res = await fetch(`${baseUrl}?` + new URLSearchParams({
      companyId: request.companyId.toString(),
      limit: request.limit.toString(),
      offset: request.offset.toString()
    }).toString(), {
      method: 'GET',
      headers: headers,
    });

    let response = await res.json();

    return response;
  } catch (error: any) {
    console.log(error.toString())
  }
}

export async function DeleteProductRequest(
  request: DeleteProduct
) {
  try {
    const headers = new Headers();
    headers.append("Authorization", <string>localStorage.getItem('token'));

    await fetch(`${baseUrl}/${request.id}`, {
      method: 'DELETE',
      headers: headers,
    });

  } catch (error: any) {
    console.log(error.toString())
  }
}
