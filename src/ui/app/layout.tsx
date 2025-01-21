import type { Metadata } from "next";
import { Inter } from "next/font/google";
import "./globals.css";
import Providers from "./components/Providers";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "Edezco",
  icons: {
    icon: "/images/favicon.ico",
  },
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <head>
        <title>Edezco</title>
        <meta
          name="description"
          content="Crea tu tienda online y lleva control de todos tus productos y ventas en cualquier lugar"
        />
        <meta property="og:title" content="Edezco" />
        <meta
          property="og:description"
          content="Crea tu tienda online y lleva control de todos tus productos y ventas en cualquier lugar"
        />
        <meta property="og:image" content="https://example.com/thumbnail.jpg" />
        <meta name="twitter:card" content="summary_large_image" />
      </head>
      <body className={inter.className}>
        <Providers>{children}</Providers>
      </body>
    </html>
  );
}
