FROM node:22.6

WORKDIR /web

COPY package*.json ./

RUN npm install

COPY . .

CMD ["npm", "run", "dev"]

EXPOSE 5173

