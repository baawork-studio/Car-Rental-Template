FROM node:24-alpine AS build
WORKDIR /app
RUN corepack enable
COPY package.json pnpm-lock.yaml pnpm-workspace.yaml ./
COPY apps/admin/package.json apps/admin/package.json
RUN pnpm install --frozen-lockfile
COPY apps/admin apps/admin
RUN pnpm --filter @car-rental/admin build

FROM nginx:1.29-alpine
COPY --from=build /app/apps/admin/dist /usr/share/nginx/html
EXPOSE 80
