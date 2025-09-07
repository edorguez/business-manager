/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: false,
  images: {
    remotePatterns: [
      {
        protocol: 'https',
        hostname: 'business-manager-bucket-s3.s3.us-east-1.amazonaws.com',
      }
    ],
  },
};

export default nextConfig;
