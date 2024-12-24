export const validLettersAndNumbers = (
  input: string,
  allowSpaces: boolean = false
): boolean => {
  if (!input) return false;
  const regex = allowSpaces ? /^[a-zA-Z0-9\s]*$/ : /^[a-zA-Z0-9]*$/;
  return regex.test(input);
};

export const validLetters = (
  input: string,
  allowSpaces: boolean = false
): boolean => {
  if (!input) return false;
  const regex = allowSpaces ? /^[a-zA-ZÀ-ÖØ-öø-ÿ\s]*$/ : /^[a-zA-ZÀ-ÖØ-öø-ÿ]*$/;
  return regex.test(input);
};

export const validNumbers = (input: string): boolean => {
  if (!input) return false;
  const regex = /^[0-9]*$/;
  return regex.test(input);
};

export const validEmail = (email: string): boolean => {
  if (!email) return false;
  const regex = /^\S+@\S+\.\S+$/;
  return regex.test(email);
};

export const validPhone = (phone: string): boolean => {
  if (!phone) return false;
  const regex = /^(0414|0424|0412|0416|0426|414|424|412|416|426)\d{6,7}$/;
  return regex.test(phone);
};

export const validIdentification = (identification: string): boolean => {
  if (!identification) return false;
  const regex = /^[0-9]{5,9}$/;
  return regex.test(identification);
};

export const validWithNoSpaces = (input: string): boolean => {
  if (!input) return false;
  const regex = /\s/;
  return !regex.test(input);
};
