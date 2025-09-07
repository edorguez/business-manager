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

declare global {
  interface Window {
    dataLayer: any[];
  }
}

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
        <meta
          property="og:image"
          content="https://edezco.com/_next/image?url=%2Fimages%2Flogo.png&w=128&q=75"
        />
        <meta name="twitter:card" content="summary_large_image" />
        <script
          async
          src="https://www.googletagmanager.com/gtag/js?id=G-QDMNTWDHZQ"
        ></script>
        <script
          dangerouslySetInnerHTML={{
            __html: `
              window.dataLayer = window.dataLayer || [];
              function gtag(){dataLayer.push(arguments);}
              gtag('js', new Date());
              gtag('config', 'G-QDMNTWDHZQ');
            `,
          }}
        />
      </head>
      <body className={inter.className}>
        <Providers>{children}</Providers>
      </body>
    </html>
  );
}
