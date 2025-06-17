FROM node:22 as builder
WORKDIR /app
COPY . .
# Copy production environment file to the right location
RUN mkdir -p src/environments
COPY src/environments-docker/environment.prod.ts src/environments/
# Configure Node options for Angular build
# No need for --openssl-legacy-provider with Node 22 as it's not required
ENV NODE_OPTIONS="--max-old-space-size=4096"
# Clean any existing node_modules to prevent conflicts
RUN rm -rf node_modules package-lock.json
# Use legacy-peer-deps to bypass peer dependency issues
# Copy build script and use it to build with type checking disabled
COPY build.sh /app/build.sh
RUN npm install --legacy-peer-deps && \
    chmod +x /app/build.sh && \
    /app/build.sh

FROM nginx:stable
WORKDIR /app
COPY ./docker_files/default.conf /etc/nginx/conf.d/default.conf
COPY --from=builder /app/dist/ /app/
# Fix permissions for nginx user
RUN chmod -R 755 /app && \
    chown -R nginx:nginx /app
EXPOSE 80 443