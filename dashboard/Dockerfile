FROM node:20-alpine
WORKDIR /app

# Init minimal React app
COPY package.json yarn.lock* ./
RUN yarn install

COPY . .
RUN yarn build || echo "Dashboard skeleton build skipped"

EXPOSE 5174
CMD ["yarn", "dev"]