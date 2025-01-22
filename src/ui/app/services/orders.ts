import { CreateOrder } from "../types/order";

const baseUrl: string =
  process.env.ENVIRONMENT === "production"
    ? "http://gateway:3001/api/orders"
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