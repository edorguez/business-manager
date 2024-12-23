import resolveConfig from "tailwindcss/resolveConfig";
import config from "../../tailwind.config";
import dayjs from "dayjs";

export const tailwindConfig = () => {
  // Tailwind config
  return resolveConfig(config);
};

export const hexToRGB = (h: any) => {
  let r: string = "0";
  let g: string = "0";
  let b: string = "0";
  if (h.length === 4) {
    r = `0x${h[1]}${h[1]}`;
    g = `0x${h[2]}${h[2]}`;
    b = `0x${h[3]}${h[3]}`;
  } else if (h.length === 7) {
    r = `0x${h[1]}${h[2]}`;
    g = `0x${h[3]}${h[4]}`;
    b = `0x${h[5]}${h[6]}`;
  }
  return `${+r},${+g},${+b}`;
};

export const formatValue = (value: any) =>
  Intl.NumberFormat("en-US", {
    style: "currency",
    currency: "USD",
    maximumSignificantDigits: 3,
    notation: "compact",
  }).format(value);

export const formatTitleValue = (value: any): string => {
  if (dayjs(value).isValid()) {
    return dayjs(value).format("DD/MM/YYYY");
  }

  return value;
};

export const convertToTimezone = (date: Date, targetTimezoneOffset: number) => {
  // Get the timezone offset for the original date (in minutes)
  const originalTimezoneOffset = date.getTimezoneOffset();
  // Calculate the difference in timezone offsets
  const offsetDifference = targetTimezoneOffset - originalTimezoneOffset;
  // Create a new date adjusted by the difference
  const targetDate = new Date(date.getTime() + offsetDifference * 60000);

  return targetDate;
};

export const validateUserInRoles = (
  roleIds: number[],
  userRoleId: number
): boolean => {
  return roleIds.some((x) => x === userRoleId);
};

export const numberMoveDecimal = (num: number, places: number): number => {
  return num / Math.pow(10, places);
};
