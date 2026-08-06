FROM node:24-alpine AS build
WORKDIR /app
RUN corepack enable
COPY package.json ./
RUN pnpm install --no-frozen-lockfile
COPY . .
ARG VITE_API_URL
ARG VITE_LIFF_ID
ENV VITE_API_URL=$VITE_API_URL
ENV VITE_LIFF_ID=$VITE_LIFF_ID
RUN pnpm build

FROM nginx:1.29-alpine
COPY --from=build /app/dist /usr/share/nginx/html
COPY nginx.conf /etc/nginx/templates/default.conf.template
EXPOSE 8080
