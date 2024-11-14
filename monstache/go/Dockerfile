# Ref - https://github.com/rwynn/monstache/wiki/Go-plugin-guide

####################################################################################################
# Step 1: Build the app and plugin
####################################################################################################

FROM rwynn/monstache-builder-cache-rel6:1.0.8 AS build-app

RUN mkdir /app

WORKDIR /app

RUN git clone --depth 1 --branch rel6 https://github.com/rwynn/monstache.git .

COPY ./plugin.go .

RUN make release

RUN go build -buildmode=plugin -o build/plugin.so plugin.go

####################################################################################################
# Step 2: Copy output build files to an alpine image
####################################################################################################

FROM rwynn/monstache-alpine:3.15.0 AS final

ENTRYPOINT ["/bin/monstache"]

COPY --from=build-app /app/build/linux-amd64/monstache /bin/monstache
COPY --from=build-app /app/build/plugin.so /bin/plugin.so
