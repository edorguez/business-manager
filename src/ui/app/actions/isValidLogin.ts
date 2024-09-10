import { jwtDecode } from "jwt-decode";
import { CurrentUser } from "../types/auth";
import dayjs from "dayjs";

export default function isValidLogin(): boolean {
  try {
    const token: string | null = localStorage.getItem("token");
    if (token) {
      const currentUser: CurrentUser = jwtDecode(token);
      return currentUser.exp > dayjs(new Date()).unix();
    }
  } catch (error: any) {
    return false;
  }

  return false;
}
