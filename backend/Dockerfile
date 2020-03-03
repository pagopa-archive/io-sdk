FROM actionloop/standalone:latest
ADD public /app/public/
ADD actions /app/actions/
RUN sed -i -e 's!"/!"/app/!g' /app/public/index.html
