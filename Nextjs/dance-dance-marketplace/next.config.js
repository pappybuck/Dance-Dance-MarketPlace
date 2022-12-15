/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
  swcMinify: true,
  output: 'standalone',
  experimental: {
    isrMemoryCacheSize: 0,
  }
}

module.exports = nextConfig
