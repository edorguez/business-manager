import { mode } from '@chakra-ui/theme-tools';

const config: any = {
  colors: {
    main: {
      50: '#bbf7f4',
      100: '#8ef1ec',
      200: '#60ebe4',
      300: '#33e6dd',
      400: '#19ccc3',
      500: '#14A098',
      600: '#118882',
      700: '#0e716d',
      800: '#0b5b57',
      900: '#084441',
    },
    third: {
      50: '#f6d5e3',
      100: '#f1c0d5',
      200: '#edabc7',
      300: '#e382aa',
      400: '#da588e',
      500: '#CB2D6F',
      600: '#a7255b',
      700: '#7d1c44',
      800: '#54122e',
      900: '#2a0917',
    },
  },
  components: {
    Button: {
      variants: {
        main: (props: any) => ({
          color: mode('white', 'gray.800')(props),
          backgroundColor: mode('main.500', 'main.200')(props),
          _hover: {
            backgroundColor: mode('main.600', 'main.300')(props),
          },
          _active: {
            backgroundColor: mode('main.700', 'main.400')(props),
          },
        }),
        third: (props: any) => ({
          color: mode('white', 'gray.800')(props),
          backgroundColor: mode('third.500', 'third.200')(props),
          _hover: {
            backgroundColor: mode('third.600', 'third.300')(props),
          },
          _active: {
            backgroundColor: mode('third.700', 'third.400')(props),
          },
        }),
      },
    },
  },
}
export default config;
