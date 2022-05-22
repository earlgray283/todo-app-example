FROM --platform=linux/amd64 node:16

ARG _REACT_APP_BACKEND_URL
ENV REACT_APP_BACKEND_URL ${_REACT_APP_BACKEND_URL}

WORKDIR /work

RUN yarn global add serve

COPY package.json yarn.lock tsconfig.json /work/

RUN yarn install

COPY . .

RUN yarn build
EXPOSE 8080
ENTRYPOINT [ "serve", "-s", "build", "-p", "8080" ]
