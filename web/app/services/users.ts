import {
  CreateUser,
  DeleteUser,
  EditEmail,
  EditPassword,
  EditUser,
  GetUser,
  GetUsers,
} from "@/app/types/user";
import Cookies from "js-cookie";

const baseUrl: string =
  process.env.NEXT_PUBLIC_ENVIRONMENT === "production"
    ? "https://edezco.com/api/users"
    : "http://localhost:3001/api/users";

export async function CreateUserRequest(request: CreateUser) {
  try {
    const headers = new Headers();
    const token = Cookies.get("token");
    headers.append("Authorization", <string>token);

    request.email = request.email.trim();
    request.password = request.password.trim();
    request.roleId = +request.roleId;

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

export async function GetUserRequest(request: GetUser) {
  try {
    const headers = new Headers();
    const token = Cookies.get("token");
    headers.append("Authorization", <string>token);

    const res = await fetch(`${baseUrl}/${request.id}`, {
      method: "GET",
      headers: headers,
    });

    let response = await res.json();

    return response;
  } catch (error: any) {
    console.log(error.toString());
  }
}

export async function GetUsersRequest(request: GetUsers) {
  try {
    const headers = new Headers();
    const token = Cookies.get("token");
    headers.append("Authorization", <string>token);

    const res = await fetch(
      `${baseUrl}?` +
        new URLSearchParams({
          companyId: request.companyId.toString(),
          limit: request.limit.toString(),
          offset: request.offset.toString(),
        }),
      {
        method: "GET",
        headers: headers,
      }
    );

    let response = await res.json();

    return response;
  } catch (error: any) {
    console.log(error.toString());
  }
}

export async function EditUserRequest(request: EditUser) {
  try {
    const headers = new Headers();
    const token = Cookies.get("token");
    headers.append("Authorization", <string>token);

    request.email = request.email.trim();
    request.roleId = +request.roleId;
    if(request.password)
      request.password = request.password.trim();

    const res = await fetch(`${baseUrl}/${request.id}`, {
      method: "PUT",
      headers: headers,
      body: JSON.stringify(request),
    });

    if (res.status !== 204) return await res.json();
  } catch (error: any) {
    console.log(error.toString());
  }
}

export async function EditEmailRequest(request: EditEmail) {
  try {
    const headers = new Headers();
    const token = Cookies.get("token");
    headers.append("Authorization", <string>token);

    request.email = request.email.trim();

    const res = await fetch(`${baseUrl}/${request.id}/email`, {
      method: "PUT",
      headers: headers,
      body: JSON.stringify(request),
    });

    if (res.status !== 204) return await res.json();
  } catch (error: any) {
    console.log(error.toString());
  }
}

export async function EditPasswordRequest(request: EditPassword) {
  try {
    const headers = new Headers();
    const token = Cookies.get("token");
    headers.append("Authorization", <string>token);

    request.password = request.password.trim();

    const res = await fetch(`${baseUrl}/${request.id}/password`, {
      method: "PUT",
      headers: headers,
      body: JSON.stringify(request),
    });

    if (res.status !== 204) return await res.json();
  } catch (error: any) {
    console.log(error.toString());
  }
}

export async function DeleteUserRequest(request: DeleteUser) {
  try {
    const headers = new Headers();
    const token = Cookies.get("token");
    headers.append("Authorization", <string>token);

    await fetch(`${baseUrl}/${request.id}`, {
      method: "DELETE",
      headers: headers,
    });
  } catch (error: any) {
    console.log(error.toString());
  }
}