export type SignUp = {
  company: SignUpCompany;
  user: SignUpUser;
}

export type SignUpCompany = {
  name: string;
  phone: string;
}

export type SignUpUser = {
  email: string;
  password: string;
}
