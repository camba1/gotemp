FROM node:alpine

WORKDIR /go/src/goTemp/web/svelte

COPY ./svelte/package*.json ./

RUN npm install

COPY ./svelte .

EXPOSE 5000

ENV HOST=0.0.0.0

CMD [ "npm", "run", "dev" ]
#CMD [ "npm", "start" ]
#CMD [ "npm", "startSPA" ]