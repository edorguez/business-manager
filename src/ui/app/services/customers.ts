import { CreateCustomer, DeleteCustomer, EditCustomer, GetCustomer, GetCustomers, GetCustomersByMonths } from "@/app/types/customer";
import Cookies from 'js-cookie';

const baseUrl: string =
  process.env.ENVIRONMENT === "production"
    ? "http://gateway:3001/api/customers"
    : "http://localhost:3001/api/customers";

export async function CreateCustomerRequest(
  request: CreateCustomer
) {
  try {
    const headers = new Headers();
    const token = Cookies.get('token');
    headers.append("Authorization", <string>token);

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

export async function EditCustomerRequest(
  request: EditCustomer
) {
  try {
    const headers = new Headers();
    const token = Cookies.get('token');
    headers.append("Authorization", <string>token);

    const res = await fetch(`${baseUrl}/${request.id}`, {
      method: 'PUT',
      headers: headers,
      body: JSON.stringify(request),
    });
  } catch (error: any) {
    console.log(error.toString())
  }
}

export async function GetCustomerRequest(
  request: GetCustomer
) {
  try {
    const headers = new Headers();
    const token = Cookies.get('token');
    headers.append("Authorization", <string>token);

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

export async function GetCustomersRequest(
  request: GetCustomers
) {
  try {
    const headers = new Headers();
    const token = Cookies.get('token');
    headers.append("Authorization", <string>token);

    const res = await fetch(`${baseUrl}?` + new URLSearchParams({
      companyId: request.companyId.toString(),
      name: request.name.trim(),
      lastName: request.lastName.trim(),
      identificationNumber: request.identificationNumber.trim(),
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

export async function GetCustomersByMonthsRequest(
  request: GetCustomersByMonths
) {
  try {
    const headers = new Headers();
    const token = Cookies.get('token');
    headers.append("Authorization", <string>token);

    const res = await fetch(`${baseUrl}/months?` + new URLSearchParams({
      companyId: request.companyId.toString(),
      months: request.months.toString()
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

export async function DeleteCustomerRequest(
  request: DeleteCustomer
) {
  try {
    const headers = new Headers();
    const token = Cookies.get('token');
    headers.append("Authorization", <string>token);

    await fetch(`${baseUrl}/${request.id}`, {
      method: 'DELETE',
      headers: headers,
    });

  } catch (error: any) {
    console.log(error.toString())
  }
}
