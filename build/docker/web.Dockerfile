FROM node:12.13.0-alpine

VOLUME [ "/react-app" ]

WORKDIR /react-app

COPY ./web ./

# Path env
ENV NODE_PATH=/node_modules
ENV PATH /node_modules/.bin:$PATH

# Install Yarn
RUN npm install -g yarn
RUN npm install -g react-scripts

# Install packages
RUN apk add python2
RUN npm install -g serve
RUN yarn;

# Define Env Vars
ARG PORT=3000
ENV PORT $PORT

# Run app
EXPOSE ${PORT}
# RUN yarn build
# CMD [ "serve", "-s", "build", "-l", "3000" ]
CMD [ "yarn", "start" ]
