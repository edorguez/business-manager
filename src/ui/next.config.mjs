/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: false,
  images: {
    remotePatterns: [
      {
        protocol: 'https',
        hostname: 'marketplace.canva.com',
      },
    ],
  },
};

export default nextConfig;
