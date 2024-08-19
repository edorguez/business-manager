/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: false,
  images: {
    remotePatterns: [
      {
        protocol: 'https', //DELETE THIS OBJECT
        hostname: 'marketplace.canva.com', // DELETE THIS OBJECT
      }
    ],
  },
};

export default nextConfig;
