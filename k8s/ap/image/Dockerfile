FROM gcr.io/google_containers/ubuntu-slim:0.7
EXPOSE 8080
copy ap /usr/local/bin/ap
ENV AP_PORT 8080
ENTRYPOINT AP_PORT=8080 /usr/local/bin/ap
