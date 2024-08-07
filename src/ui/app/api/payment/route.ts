import { CreatePayment } from "@/app/types/payment";

const baseUrl: string = 'http://localhost:3001/api/payments';

export async function CreatePaymentRequest(
  request: CreatePayment
) {
  try {
    const headers = new Headers();
    headers.append("Authorization", <string>localStorage.getItem('token'));

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