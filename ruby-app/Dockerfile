FROM ruby:2.3

RUN mkdir /app
ADD . /app

WORKDIR /app
RUN bundle install -j 4

EXPOSE 9292

CMD ["rackup", "config.ru", "-p", "9292", "-o", "0.0.0.0"]
