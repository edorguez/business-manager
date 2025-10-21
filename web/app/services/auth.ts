import { Login } from "@/app/types/auth"
import Cookies from "js-cookie";

const baseUrl: string =
  process.env.NEXT_PUBLIC_ENVIRONMENT === "production"
    ? "http://gateway:3001/api/auth"
    : "http://localhost:3001/api/auth";

export async function login(
  request: Login,
) {
  try {
    const res = await fetch(`${baseUrl}/login`, {
      method: 'POST',
      body: JSON.stringify(request)
    });
    let response = await res.json();

    if(!response.error) {
      Cookies.set('token', `Bearer ${response.token}`);
    }

    return response;
  } catch (error: any) {
    console.log(error.toString())
  }
}
