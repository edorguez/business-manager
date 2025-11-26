import { Login, SignUp } from "@/app/types/auth"
import Cookies from "js-cookie";

const baseUrl: string =
  process.env.NEXT_PUBLIC_ENVIRONMENT === "production"
    ? "http://edezco.com/api/auth"
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

export async function signUp(
  request: SignUp,
) {
  try {
    request.company.name = request.company.name.trim();

    const formData = new FormData();

    // Add JSON data as a string
    formData.append("json", JSON.stringify(request));

    // Add images to the FormData
    request.company.images?.forEach((image) => {
      formData.append(`files`, image);
    });

    const res = await fetch(`${baseUrl}/signup`, {
      method: "POST",
      body: formData,
    });

    let response = await res.json();

    if(!response.error) {
      Cookies.set('token', `Bearer ${response.token}`);
    }

    return response;
  } catch (error: any) {
    console.log(error.toString());
  }
}
