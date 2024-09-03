FROM nginx

RUN adduser --system --no-create-home --shell /bin/false --group --disabled-login www

COPY nginx.conf /etc/nginx/nginx.conf
EXPOSE 80