import { CreateOrder } from "../types/order";

const baseUrl: string = "http://localhost:3001/api/orders";

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