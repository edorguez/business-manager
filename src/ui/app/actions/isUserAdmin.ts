export default function isUserAdmin(userRoleId: number): boolean {
  return userRoleId === 1 || userRoleId === 2;
}
