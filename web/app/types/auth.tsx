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

export type SignUp = {
  company: SignUpCompany;
  user: SignUpUser;
}

export type SignUpCompany = {
  name: string;
  nameFormatUrl: string;
  phone: string;
  images?: File[];
}

export type SignUpUser = {
  email: string;
  password: string;
  passwordRepeat: string;
}
