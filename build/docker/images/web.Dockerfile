FROM node:12.13.0-alpine as build-deps

WORKDIR /react-app

# Install executables
RUN apk add python2
RUN yarn global add serve
RUN yarn global add react-scripts

# Install modules
COPY ./web .
RUN yarn

# Define Env Vars
ARG PORT=3000
ENV PORT $PORT

# Run app
EXPOSE ${PORT}
# RUN yarn build
# CMD [ "serve", "-s", "build", "-l", "3000" ]
CMD [ "yarn", "start" ]
