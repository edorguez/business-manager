import { CreateOrder, GetOrders } from "../types/order";
import Cookies from 'js-cookie';

const baseUrl: string =
  process.env.NEXT_PUBLIC_ENVIRONMENT === "production"
    ? "https://edezco.com/api/orders"
    : "http://localhost:3001/api/orders";

export async function CreateOrderRequest(request: CreateOrder) {
  try {
    const headers = new Headers();

    const res = await fetch(baseUrl, {
      method: "POST",
      headers: headers,
      body: JSON.stringify(request),
    });

    let response = await res.json();

    return response;
  } catch (error: any) {
    console.log(error.toString());
  }
}

export async function GetOrdersRequest(
  request: GetOrders
) {
  try {
    const headers = new Headers();
    const token = Cookies.get('token');
    headers.append("Authorization", <string>token);

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
