FROM node:24-alpine AS build
WORKDIR /app
RUN corepack enable
COPY package.json pnpm-lock.yaml pnpm-workspace.yaml ./
COPY apps/customer/package.json apps/customer/package.json
RUN pnpm install --frozen-lockfile
COPY apps/customer apps/customer
RUN pnpm --filter @car-rental/customer build

FROM nginx:1.29-alpine
COPY --from=build /app/apps/customer/dist /usr/share/nginx/html
EXPOSE 80
