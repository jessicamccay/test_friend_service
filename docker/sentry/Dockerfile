FROM slafs/sentry

ENV SECRET_KEY secret
ENV SENTRY_URL_PREFIX http://sentry

ENV SENTRY_INITIAL_TEAM friends
ENV SENTRY_INITIAL_PROJECT friends
ENV SENTRY_INITIAL_KEY local:secret

RUN /usr/local/bin/sentry_run prepare
