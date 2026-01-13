import { mode } from "@chakra-ui/theme-tools";

const config: any = {
  colors: {
    main: {
      50: "#bbf7f4",
      100: "#8ef1ec",
      200: "#60ebe4",
      300: "#33e6dd",
      400: "#19ccc3",
      500: "#14A098",
      600: "#118882",
      700: "#0e716d",
      800: "#0b5b57",
      900: "#084441",
    },
    third: {
      50: "#f6d5e3",
      100: "#f1c0d5",
      200: "#edabc7",
      300: "#e382aa",
      400: "#da588e",
      500: "#CB2D6F",
      600: "#a7255b",
      700: "#7d1c44",
      800: "#54122e",
      900: "#2a0917",
    },
    fifth: {
      50: "#cceeff",
      100: "#99ddff",
      200: "#80d4ff",
      300: "#66ccff",
      400: "#33bbff",
      500: "#00A9FF",
      600: "#0088cc",
      700: "#006699",
      800: "#004466",
      900: "#002233",
    },
    default: {
      50: "#ffffff",
      100: "#d9d9d9",
      200: "#b3b3b3",
      300: "#999999",
      400: "#808080",
      500: "#666666",
      600: "#4d4d4d",
      700: "#333333",
      800: "#1a1a1a",
      900: "#000000",

    }
  },
  components: {
    Button: {
      variants: {
        main: (props: any) => ({
          color: mode("white", "gray.800")(props),
          backgroundColor: mode("main.500", "main.200")(props),
          _hover: {
            backgroundColor: mode("main.600", "main.300")(props),
            _disabled: {
              backgroundColor: mode("main.500", "main.200")(props),
            },
          },
          _active: {
            backgroundColor: mode("main.700", "main.400")(props),
          },
          _loading: {
            backgroundColor: mode("main.500", "main.200")(props),
          },
        }),
        third: (props: any) => ({
          color: mode("white", "gray.800")(props),
          backgroundColor: mode("third.500", "third.200")(props),
          _hover: {
            backgroundColor: mode("third.600", "third.300")(props),
          },
          _active: {
            backgroundColor: mode("third.700", "third.400")(props),
          },
        }),
        fifth: (props: any) => ({
          color: mode("white", "gray.800")(props),
          backgroundColor: mode("fifth.500", "fifth.200")(props),
          _hover: {
            backgroundColor: mode("fifth.600", "fifth.300")(props),
          },
          _active: {
            backgroundColor: mode("fifth.700", "fifth.400")(props),
          },
        }),
        default: (props: any) => ({
          color: mode("white", "gray.800")(props),
          backgroundColor: mode("default.700", "default.400")(props),
          _hover: {
            backgroundColor: mode("default.800", "default.400")(props),
            _disabled: {
              backgroundColor: mode("default.700", "default.300")(props),
            },
          },
          _active: {
            backgroundColor: mode("default.800", "default.500")(props),
          },
          _loading: {
            backgroundColor: mode("default.600", "default.300")(props),
          },
        }),
      },
    },
    Input: {
      variants: {
        default: (props: any) => ({
          field: {
            border: "1px solid",
            borderColor: mode("default.100", "default.600")(props),
            _focus: {
              borderColor: mode("default.500", "default.400")(props),
            },
            _hover: {
              borderColor: mode("default.400", "default.500")(props),
            },
            _invalid: {
              borderColor: mode("third.500", "third.300")(props),
            },
          },
        }),
      },
      defaultProps: {
        focusBorderColor: "main.500",
        errorBorderColor: "third.500",
      },
    },
    NumberInput: {
      variants: {
        default: (props: any) => ({
          field: {
            border: "1px solid",
            borderColor: mode("default.100", "default.600")(props),
            _focus: {
              borderColor: mode("default.500", "default.400")(props),
            },
            _hover: {
              borderColor: mode("default.400", "default.500")(props),
            },
            _invalid: {
              borderColor: mode("third.500", "third.300")(props),
            },
          },
        }),
      },
      defaultProps: {
        focusBorderColor: "main.500",
        errorBorderColor: "third.500",
      },
    },
    Select: {
      variants: {
        default: (props: any) => ({
          field: {
            border: "1px solid",
            borderColor: mode("default.100", "default.600")(props),
            _focus: {
              borderColor: mode("default.500", "default.400")(props),
            },
            _hover: {
              borderColor: mode("default.400", "default.500")(props),
            },
            _invalid: {
              borderColor: mode("third.500", "third.300")(props),
            },
          },
        }),
      },
    },
    Alert: {
      variants: {
        customsuccess: (props: any) => {
          return {
            container: {
              ...props.theme.components.Alert.variants.solid(props).container,
              backgroundColor: mode("main.500", "main.500")(props),
              color: mode("white", "black")(props),
            },
          };
        },
        customerror: (props: any) => {
          return {
            container: {
              ...props.theme.components.Alert.variants.solid(props).container,
              backgroundColor: mode("third.500", "third.500")(props),
              color: mode("white", "black")(props),
            },
          };
        },
      },
    },
  },
};
export default config;
