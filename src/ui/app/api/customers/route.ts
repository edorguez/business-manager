import { CreateCustomer, GetCustomer } from "@/app/types/customer";

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

export async function GetCustomersRequest(
  request: GetCustomer
) {
  try {
    const headers = new Headers();
    headers.append("Authorization", <string>localStorage.getItem('token'));

    const res = await fetch(`${baseUrl}?` + new URLSearchParams({
      companyId: request.companyId.toString(),
      name: request.name,
      lastName: request.lastName,
      identificationNumber: request.identificationNumber,
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
