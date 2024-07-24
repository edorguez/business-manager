import { jwtDecode } from "jwt-decode";
import { CurrentUser } from "../types/auth";

export default function getCurrentUser(): CurrentUser | null {
  try {
    const token: string | null = localStorage.getItem("token");
    if (token) {
      const currentUser: CurrentUser = jwtDecode(token)
      return currentUser;
    }
  } catch (error: any) {
    return null;
  }

  return null;
}
