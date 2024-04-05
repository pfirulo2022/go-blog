FROM nginx:alpine

COPY ./nginx/default.conf /etc/nginx/conf.d/default.conf 

WORKDIR /go-blog

COPY ./client/dist /index.html

COPY ./server/.env.prod .

COPY ./server/build/server.exe .

COPY ./startup.sh .

EXPOSE 80
EXPOSE 8000

ENTRYPOINT [ "/bin/sh", "startup.sh" ]