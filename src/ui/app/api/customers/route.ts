import { CreateCustomer } from "@/app/types/customer";

const baseUrl: string = 'http://localhost:3001/api/customers';

export async function CreateCustomerRequest(
  request: CreateCustomer
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
