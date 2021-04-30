# Rudeus Telegram Bot Project
# Copyright (C) 2021 wotoTeam, ALiwoto
# This file is subject to the terms and conditions defined in
# file 'LICENSE', which is part of the source code.

FROM heroku/heroku:18-build as build

COPY . /app
WORKDIR /app

# Setup the buildpacks of the project Rudeus bot 2021
RUN mkdir -p /tmp/buildpack/heroku/go /tmp/build_cache /tmp/env
RUN curl https://codon-buildpacks.s3.amazonaws.com/buildpacks/heroku/go.tgz | tar xz -C /tmp/buildpack/heroku/go

#Execute the Buildpack of the prject Rudeus bot 2021
RUN STACK=heroku-18 /tmp/buildpack/heroku/go/bin/compile /app /tmp/build_cache /tmp/env

# Prepare the final, minimal image of the project Rudeus bot 2021
FROM heroku/heroku:18

# Copy app directory from build directory
COPY --from=build /app /app
ENV HOME /app
WORKDIR /app
RUN useradd -m aliwoto
USER heroku

# run the binary file
CMD /app/bin/rudeus01