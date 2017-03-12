FROM gcr.io/google_containers/ubuntu-slim:0.7
EXPOSE 8080
copy blob /usr/local/bin/blob
ENV BLOB_PORT 9080
ENTRYPOINT BLOB_PORT=9080 /usr/local/bin/blob
