"use server";

import { jwtDecode } from "jwt-decode";
import { CurrentUser } from "../types/auth";
import { cookies } from "next/headers";

export default function getCurrentUserServer(): CurrentUser | null {
  try {
    const cookieStore = cookies();
    const token: string | undefined = cookieStore.get("token")?.value;
    if (token) {
      const currentUser: CurrentUser = jwtDecode(token);
      return currentUser;
    }
  } catch (error: any) {
    return null;
  }

  return null;
}
