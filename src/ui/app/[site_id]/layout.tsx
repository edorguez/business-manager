import { Inter } from "next/font/google";
import { notFound } from "next/navigation";
import { ReactNode } from "react";
import Providers from "../components/Providers";

const inter = Inter({ subsets: ["latin"] });

export default async function SiteLayout({
  params,
  children,
}: {
  params: { site_id: string };
  children: ReactNode;
}) {
  const result: any = {}; //await readSiteById(params?.site_id);

  if (!result) {
    notFound();
  }

  const siteName = result?.[0]?.site_name;
  const siteDomain = result?.[0]?.site_custom_domain;
  const siteDescription = result?.[0]?.site_description;
  const siteLogo = result?.[0]?.site_logo;
  const siteCover = result?.[0]?.site_cover_image;

  return (
    <>
      <head>
        <title>{siteName}</title>
        <meta name="site_name" content={siteName} />
        <meta name="description" content={siteDescription} />
        {siteCover && <meta name="image" content={siteCover} />}
        <meta
          name="url"
          content={siteDomain + "." + process.env.NEXT_PUBLIC_FRONTEND_URL}
        ></meta>
        {siteLogo && <link rel="icon" href={siteLogo} />}
      </head>
      <body className={inter.className}>
        <Providers>{children}</Providers>
      </body>
    </>
  );
}
