/** @type {import('next').NextConfig} */
const nextConfig = {
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
