FROM node:18 as builder
WORKDIR /app
COPY . .
# Copy production environment file to the right location
RUN mkdir -p src/environments
COPY src/environments-docker/environment.prod.ts src/environments/
# Set OpenSSL configuration for compatibility
ENV NODE_OPTIONS="--openssl-legacy-provider"
RUN npm install && npm run build

FROM nginx:stable
WORKDIR /app
COPY ./docker_files/default.conf /etc/nginx/conf.d/default.conf
COPY --from=builder /app/dist/ /app/
# Fix permissions for nginx user
RUN chmod -R 755 /app && \
    chown -R nginx:nginx /app
EXPOSE 80 443