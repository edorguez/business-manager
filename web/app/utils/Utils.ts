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

export const convertTimestampToDate = (timestamp: {
  seconds: number;
  nanos: number;
}): Date => {
  // Convert seconds to milliseconds and add nanoseconds as milliseconds
  return new Date(timestamp.seconds * 1000 + timestamp.nanos / 1000000);
};

export const formatPriceStringToNumberBackend = (price: string): number => {
  if (price.indexOf(".") !== -1) {
    let parts = price.split(".");

    // Handle cases where there is no integer part
    let integerPart = parts[0].length > 0 ? parts[0] : "0";
    let decimalPart = parts[1] ? parts[1] : "0";

    // Ensure decimal part has exactly two digits
    if (decimalPart.length === 1) {
      decimalPart += "0";
    } else if (decimalPart.length > 2) {
      decimalPart = decimalPart.substring(0, 2);
    }

    return parseInt(integerPart + decimalPart);
  }

  // Handle cases where there is no comma
  return parseInt(price) * 100;
};

export const formatPriceNumberBackendToString = (price: number): string => {
  let priceString = price.toString();
  if (priceString.length === 1) {
    return `0.0${priceString}`;
  } else if (priceString.length === 2) {
    return `0.${priceString}`;
  } else {
    return `${priceString.slice(0, -2)}.${priceString.slice(-2)}`;
  }
};

export const formatCompanyNameToUrlName = (name: string): string => {
  return name
    .toLowerCase()
    // Remove all characters except letters, numbers, and spaces
    .replace(/[^a-zA-Z0-9\s]/g, '')
    // Trim leading and trailing spaces
    .trim()
    // Replace one or more consecutive spaces with a single dash
    .replace(/\s+/g, '-');
}
