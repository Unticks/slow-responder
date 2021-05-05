FROM debian:stable-slim

COPY slow-responder /app

CMD /app
