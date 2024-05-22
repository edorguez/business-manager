import type { Config } from "tailwindcss";

const config: Config = {
  content: [
    "./pages/**/*.{js,ts,jsx,tsx,mdx}",
    "./components/**/*.{js,ts,jsx,tsx,mdx}",
    "./app/**/*.{js,ts,jsx,tsx,mdx}",
  ],
  theme: {
    screens: {
      'xs': '350px',
      'sm': '640px',
      'md': '768px',
      'lg': '1024px',
      'xl': '1280px',
      '2xl': '1536px',
    },
    extend: {
      backgroundImage: {
        "gradient-radial": "radial-gradient(var(--tw-gradient-stops))",
        "gradient-conic":
          "conic-gradient(from 180deg at 50% 50%, var(--tw-gradient-stops))",
      },
      colors: {
        whitebackground: '#FFFFFF',
        graybackground: '#F2F2F2',
        iconcolor: '#6B7280',
        maincolor: '#14A098',
        maincolorhov: '#14A0981A',
        secondcolor: '#0F292F',
        secondcolorhov: '#0F292F1A',
        thirdcolor: '#CB2D6F',
        thirdcolorhov: '#CB2D6F1A',
        fourthcolor: '#501F3A',
        fourthcolorhov: '#501F3A1A',
        fifthcolor: '#00A9FF',
        fifthcolorhov: '#00A9FF1A',
        whatsappmessage: '#DCF8C6'
      }
    },
  },
  plugins: [],
};
export default config;
