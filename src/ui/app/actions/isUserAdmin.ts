import { USER_ROLE_ID } from "../utils/constants";

export default function isUserAdmin(userRoleId: number): boolean {
  return (
    userRoleId === USER_ROLE_ID.SUPER_ADMIN || userRoleId === USER_ROLE_ID.ADMIN
  );
}
