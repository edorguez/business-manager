
const baseUrl: string = 'http://localhost:3001/api/paymentTypes';

export async function GetPaymentTypesRequest() {
  try {
    const headers = new Headers();
    headers.append("Authorization", <string>localStorage.getItem('token'));

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