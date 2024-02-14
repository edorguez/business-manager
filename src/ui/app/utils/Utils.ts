import resolveConfig from "tailwindcss/resolveConfig";
import config from "../../tailwind.config";

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

export const formatValue = (value: any) => Intl.NumberFormat('en-US', {
  style: 'currency',
  currency: 'USD',
  maximumSignificantDigits: 3,
  notation: 'compact',
}).format(value);
