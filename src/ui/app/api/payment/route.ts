import { ChangeStatusPayment, CreatePayment, DeletePayment, EditPayment, GetPayment, GetPayments, GetPaymentsTypes } from "@/app/types/payment";

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
    console.log(error.toString());
  }
}

export async function EditPaymentRequest(
  request: EditPayment
) {
  try {
    const headers = new Headers();
    headers.append("Authorization", <string>localStorage.getItem('token'));

    const res = await fetch(`${baseUrl}/${request.id}`, {
      method: 'PUT',
      headers: headers,
      body: JSON.stringify(request),
    });

  } catch (error: any) {
    console.log(error.toString())
  }
}

export async function GetPaymentRequest(
  request: GetPayment
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
    console.log(error.toString());
  }
}

export async function GetPaymentsRequest(
  request: GetPayments
) {
  try {
    const headers = new Headers();
    headers.append("Authorization", <string>localStorage.getItem('token'));

    const res = await fetch(`${baseUrl}?` + new URLSearchParams({
      companyId: request.companyId.toString(),
      paymentTypeId: request.paymentTypeId.toString(),
      limit: request.limit.toString(),
      offset: request.offset.toString()
    }).toString(), {
      method: 'GET',
      headers: headers,
    });

    let response = await res.json();

    return response;
  } catch (error: any) {
    console.log(error.toString());
  }
}

export async function GetPaymentsTypesRequest(
  request: GetPaymentsTypes
) {
  try {
    const headers = new Headers();
    headers.append("Authorization", <string>localStorage.getItem('token'));

    const res = await fetch(`${baseUrl}/types?` + new URLSearchParams({
      companyId: request.companyId.toString()
    }).toString(), {
      method: 'GET',
      headers: headers,
    });

    let response = await res.json();

    return response;
  } catch (error: any) {
    console.log(error.toString());
  }
}

export async function DeletePaymentRequest(
  request: DeletePayment
) {
  try {
    const headers = new Headers();
    headers.append("Authorization", <string>localStorage.getItem('token'));

    await fetch(`${baseUrl}/${request.id}`, {
      method: 'DELETE',
      headers: headers,
    });

  } catch (error: any) {
    console.log(error.toString());
  }
}

export async function ChangeStatusRequest(
  request: ChangeStatusPayment
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