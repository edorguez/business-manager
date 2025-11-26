# Use Node.js 18 as the base image
FROM node:18-alpine

# Set environment variables for the build
ENV NEXT_PUBLIC_ENVIRONMENT=production

# Set the working directory
WORKDIR /app

# Copy package.json and package-lock.json from the web directory
COPY web/package*.json ./

# Install dependencies
RUN npm install --production

# Copy the rest of the web application code
COPY web/ ./

# Build the Next.js application
RUN npm run build

# Define the command to start the Next.js app
CMD ["npm", "start"]
