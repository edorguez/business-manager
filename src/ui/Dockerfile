# Use Node.js 18 as the base image
FROM node:18-alpine

# Set the working directory
WORKDIR /app

# Copy package.json and package-lock.json
COPY package*.json ./

# Install dependencies
RUN npm install --production

# Copy the rest of the application code
COPY . .

# Copy the .env file to the container
COPY .env .env

# Build the Next.js application
RUN npm run build

# Set the environment variable to production
ENV ENVIRONMENT=production
ENV NEXT_PUBLIC_ENVIRONMENT=production

# Define the command to start the Next.js app
CMD ["npm", "start"]