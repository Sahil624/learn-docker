FROM node:17.0.1-bullseye-slim as build
WORKDIR /usr/src/app

COPY package.json /usr/src/app
COPY package-lock.json /usr/src/app

RUN npm ci

COPY . /usr/src/app

ENV NODE_OPTIONS=--openssl-legacy-provider

RUN npm run build

FROM nginx:1.21.5-alpine
WORKDIR /usr/share/nginx/html

COPY --from=build /usr/src/app/dist /usr/share/nginx/html
COPY my.conf /etc/nginx/conf.d/default.conf
