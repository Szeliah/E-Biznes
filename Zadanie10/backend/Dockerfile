FROM node:22
WORKDIR /backend
COPY package*.json ./
RUN npm ci
COPY . .

RUN npm run build
EXPOSE 3000
CMD ["npm", "start"]