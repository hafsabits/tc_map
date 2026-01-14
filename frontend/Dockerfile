FROM node:22-alpine
WORKDIR /app
# Copy package files first to speed up builds (cache)
COPY package*.json ./
# install dependencies
RUN npm install
COPY . .
# Expose the port React runs on
EXPOSE 3000
# Start the React development server
CMD ["npm", "start"]