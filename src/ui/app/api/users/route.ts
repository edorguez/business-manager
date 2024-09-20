import { EditEmail, EditPassword, EditUser, GetUser } from '@/app/types/user';
import Cookies from 'js-cookie';

const baseUrl: string = 'http://localhost:3001/api/users';

export async function GetUserRequest(
  request: GetUser
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

export async function EditUserRequest(
  request: EditUser
) {
  try {
    const headers = new Headers();
    const token = Cookies.get('token');
    headers.append("Authorization", <string>token);
    
    request.email = request.email.trim();
    request.password = request.password.trim();

    const res = await fetch(`${baseUrl}/${request.id}`, {
      method: 'PUT',
      headers: headers,
      body: JSON.stringify(request),
    });

    if(res.status !== 204)
      return await res.json();
  } catch (error: any) {
    console.log(error.toString())
  }
}

export async function EditEmailRequest(
  request: EditEmail
) {
  try {
    const headers = new Headers();
    const token = Cookies.get('token');
    headers.append("Authorization", <string>token);

    request.email = request.email.trim();

    const res = await fetch(`${baseUrl}/${request.id}/email`, {
      method: 'PUT',
      headers: headers,
      body: JSON.stringify(request),
    });

    if(res.status !== 204)
      return await res.json();

  } catch (error: any) {
    console.log(error.toString())
  }
}

export async function EditPasswordRequest(
  request: EditPassword
) {
  try {
    const headers = new Headers();
    const token = Cookies.get('token');
    headers.append("Authorization", <string>token);

    request.password = request.password.trim();

    const res = await fetch(`${baseUrl}/${request.id}/password`, {
      method: 'PUT',
      headers: headers,
      body: JSON.stringify(request),
    });

    if(res.status !== 204)
      return await res.json();
    
  } catch (error: any) {
    console.log(error.toString())
  }
}