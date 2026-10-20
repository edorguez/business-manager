/** @type {import('next').NextConfig} */
const nextConfig = {
  output: "standalone",
  reactStrictMode: false,
  images: {
    // Uploaded images are served by the edge at /media/*. Avoid asking the
    // Next.js optimizer to fetch them from inside the container.
    unoptimized: true,
  },
};

export default nextConfig;
