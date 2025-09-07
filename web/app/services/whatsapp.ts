import Cookies from 'js-cookie';
import { CreateBusinessPhone, EditBusinessPhone } from '../types/whatsapp';

const baseUrl: string =
  process.env.NEXT_PUBLIC_ENVIRONMENT === "production"
    ? "https://edezco.com/api/whatsapp"
    : "http://localhost:3001/api/whatsapp";

export async function CreateBusinessPhoneRequest(
  request: CreateBusinessPhone
) {
  try {
    const headers = new Headers();
    const token = Cookies.get('token');
    headers.append("Authorization", <string>token);

    const res = await fetch(`${baseUrl}/businessPhone`, {
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

export async function EditBusinessPhoneRequest(
  request: EditBusinessPhone
) {
  try {
    const headers = new Headers();
    const token = Cookies.get('token');
    headers.append("Authorization", <string>token);

    const res = await fetch(`${baseUrl}/businessPhone`, {
      method: 'PUT',
      headers: headers,
      body: JSON.stringify(request),
    });
  } catch (error: any) {
    console.log(error.toString())
  }
}

export async function GetBusinessPhoneByCompanyIdRequest(
    companyId: number
) {
  try {
    const headers = new Headers();
    const token = Cookies.get('token');
    headers.append("Authorization", <string>token);

    const res = await fetch(`${baseUrl}/businessPhone/${companyId}`, {
      method: 'GET',
      headers: headers,
    });

    let response = await res.json();

    return response;
  } catch (error: any) {
    console.log(error.toString())
  }
}