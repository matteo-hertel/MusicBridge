FROM node:carbon

LABEL maintainer="Matteo Hertel <info@matteohertel.com"

ARG env 
ENV ENV $env

RUN groupadd -r nodejs && useradd -m -r -g  nodejs nodejs

USER nodejs

RUN mkdir /home/nodejs/app 
WORKDIR /home/nodejs/app

COPY package.json .
COPY package-lock.json .
RUN npm install --$ENV
COPY src/ ./src 

CMD npm run start

EXPOSE 3450
