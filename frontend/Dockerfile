FROM node:18-alpine

WORKDIR /careerscrape/frontend



COPY package*.json ./

RUN npm install

COPY . .



EXPOSE 3000
# CMD
# RUN npm run build

ENTRYPOINT [ "npm","run","dev"]
# RUN npm run dev
# RUN ls
# ENTRYPOINT ["npm", "run", "start"]
