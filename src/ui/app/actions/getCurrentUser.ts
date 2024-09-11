import { jwtDecode } from "jwt-decode";
import { CurrentUser } from "../types/auth";
import Cookies from 'js-cookie';

export default function getCurrentUser(): CurrentUser | null {
  try {
    const token = Cookies.get('token');
    if (token) {
      const currentUser: CurrentUser = jwtDecode(token)
      return currentUser;
    }
  } catch (error: any) {
    return null;
  }

  return null;
}
