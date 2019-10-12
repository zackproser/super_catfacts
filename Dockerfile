FROM ubuntu

RUN apt-get update \
    && apt-get dist-upgrade -y \
    && apt-get install -y --no-install-recommends \
          language-pack-en \
          ca-certificates \
          curl \
          lsb-release \
      && apt-get purge -y \
          krb5-locales \
      && apt-get clean -y \
      && apt-get autoremove -y \
      && rm -rf /tmp/* /var/tmp/* \
      && rm -rf /var/lib/apt/lists/*

COPY super_catfacts /super_catfacts

COPY static static

COPY data/catfacts.json data/catfacts.json

COPY config.yml config.yml

CMD ["/super_catfacts", "serve"]
