FROM node:13.12.0-alpine
WORKDIR /app

COPY src/lider-app/package.json ./

RUN npm install react-scripts@5.0.0 -g --silent
COPY src/lider-app/ ./
RUN npm run build

EXPOSE 80
# COPY ${source:-build} /app/build
RUN npm install -g serve

# FROM node:13.12.0-alpine as build
# WORKDIR /app
# ENV PATH /app/node_modules/.bin:$PATH
# COPY src/lider-app/package.json ./
# COPY src/lider-app/package-lock.json ./
# RUN npm ci --silent
# RUN npm install react-scripts@5.0.0 -g --silent
# COPY src/lider-app/ ./
# RUN npm run build

# # production environment
# FROM nginx:stable-alpine
# COPY --from=build /app/build /usr/share/nginx/html
# EXPOSE 80
# CMD ["nginx", "-g", "daemon off;"]