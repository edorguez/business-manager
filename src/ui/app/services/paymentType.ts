import Cookies from 'js-cookie';

const baseUrl: string =
  process.env.NEXT_PUBLIC_ENVIRONMENT === "production"
    ? "http://0.0.0.0:3001/api/paymentTypes"
    : "http://localhost:3001/api/paymentTypes";

export async function GetPaymentTypesRequest() {
  try {
    const headers = new Headers();
    const token = Cookies.get('token');
    headers.append("Authorization", <string>token);

    const res = await fetch(baseUrl, {
      method: 'GET',
      headers: headers,
    });

    let response = await res.json();

    return response;
  } catch (error: any) {
    console.log(error.toString())
  }
}