FROM --platform=linux/amd64 node:16

WORKDIR /work

RUN yarn global add serve

COPY package.json yarn.lock tsconfig.json /work/

RUN yarn install

COPY . .

RUN yarn build
EXPOSE 8080
ENTRYPOINT [ "serve", "-s", "build", "-p", "8080" ]
