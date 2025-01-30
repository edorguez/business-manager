'use server'

import { Login } from "@/app/types/auth"
import { cookies } from "next/headers";

const baseUrl: string =
  process.env.ENVIRONMENT === "production"
    ? "http://gateway:3001/api/auth"
    : "http://localhost:3001/api/auth";

export async function login(
  request: Login,
) {
  try {
    console.log('check borma')
    console.log(baseUrl);
    console.log(process.env.ENVIRONMENT);

    const res = await fetch(`${baseUrl}/login`, {
      method: 'POST',
      body: JSON.stringify(request)
    });
    let response = await res.json();

    if(!response.error) {
      cookies().set('token', `Bearer ${response.token}`);
    }

    return response;
  } catch (error: any) {
    console.log(error.toString())
  }
}
