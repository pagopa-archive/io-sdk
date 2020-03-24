FROM actionloop/standalone:latest
ADD public /app/public/
ADD actions/manifest.yaml /app/actions/manifest.yaml
ADD actions/deploy  /app/actions/deploy
ADD waitready /bin/waitready
RUN sed -i -e 's!"/!"/app/!g' /app/public/index.html
