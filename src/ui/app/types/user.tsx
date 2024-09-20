export type User = {
  id: number;
  companyId: number;
  roleId: number;
  email: string;
}

export type GetUser = {
  id: number;
}

export type EditUser = {
  id: number;
  roleId: number;
  email: string;
  password: string;
}

export type EditEmail = {
  id: number;
  email: string;
}

export type EditPassword = {
  id: number;
  password: string;
  passwordRepeat: string;
}