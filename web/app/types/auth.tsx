export type Login = {
  email: string;
  password: string;
}

export type CurrentUser = {
  id: number;
  email: string;
  roleId: number;
  planId: number;
  companyId: number;
  exp: number;
}
