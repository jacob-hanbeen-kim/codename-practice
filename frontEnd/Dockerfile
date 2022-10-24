FROM node:14-alpine
WORKDIR /usr/src/app
RUN apk update 
COPY . .
EXPOSE 3000
CMD [ "yarn", "run", "start" ]