"use client";

import config from "@/chakraui.config";
import {
  ChakraProvider,
  extendTheme,
} from "@chakra-ui/react";

const theme = extendTheme(config);

const Providers = ({ children }: { children: React.ReactNode }) => {
  return (
    <>
      <ChakraProvider theme={theme}>{children}</ChakraProvider>
    </>
  );
}

export default Providers;
