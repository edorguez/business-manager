"use client";

import {
  ChakraProvider,
  extendTheme,
} from "@chakra-ui/react";

const theme = extendTheme();

const Providers = ({ children }: { children: React.ReactNode }) => {
  return (
    <>
      <ChakraProvider theme={theme}>{children}</ChakraProvider>
    </>
  );
}

export default Providers;