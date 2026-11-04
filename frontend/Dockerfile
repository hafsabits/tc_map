FROM node:20-alpine

WORKDIR /app

# Copy package files first to cache dependencies
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the rest of your React code
COPY . .

# Expose the port React runs on
EXPOSE 3000

# Start the development server
CMD ["npm", "start"]