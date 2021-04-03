FROM node:alpine AS dev
WORKDIR /goTemp/web/sapper
COPY ./sapper/package*.json ./
RUN npm install
COPY ./sapper .
EXPOSE 3000
EXPOSE 10000
ENV HOST=0.0.0.0
CMD [ "npm", "run", "dev" ]
#CMD [ "npm", "build" ]
#CMD [ "npm", "start" ]

FROM dev
RUN npm run build
CMD ["npm", "run", "start"]