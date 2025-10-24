export type SignUp = {
  company: SignUpCompany;
  user: SignUpUser;
}

export type SignUpCompany = {
  name: string;
  phone: string;
  image?: File[];
}

export type SignUpUser = {
  email: string;
  password: string;
  passwordRepeat: string;
}
