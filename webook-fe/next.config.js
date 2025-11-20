/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
  async rewrites() {
    return [
      {
        source: '/users/:path*',
        destination: 'http://localhost:8899/users/:path*',
      },
    ]
  },
}

module.exports = nextConfig
