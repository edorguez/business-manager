'use server'

import { jwtDecode } from "jwt-decode";
import { CurrentUser } from "../types/auth";
import dayjs from "dayjs";
import { cookies } from "next/headers";

export default function isValidLogin(): boolean {
  try {
    const cookieStore = cookies();
    const token: string | undefined = cookieStore.get('token')?.value;
    if (token) {
      const currentUser: CurrentUser = jwtDecode(token);
      return currentUser.exp > dayjs(new Date()).unix();
    }
  } catch (error: any) {
    return false;
  }

  return false;
}
