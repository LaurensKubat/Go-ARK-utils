FROM node:10

# normally you would use a seperate docker container for the database.
# since this image will only be used for running the tests in the package, we all have it in the same docker container
RUN apt update
RUN apt install postgresql postgresql-contrib -y

# redis is used for the transaction pool. We build it ourselves, as that is the recommended way to run redis
RUN apt install redis-server -y


# Lerna grabs our dependencies for us. (it seems this one randomly fails sometimes when building the image)
RUN npm install --global lerna --loglevel verbose
RUN git clone -b master https://github.com/ArkEcosystem/core.git
RUN (cd core && lerna bootstrap)

# enabling all APIs by default
ENV ARK_WEBHOOKS_ENABLED=true
ENV ARK_WEBHOOKS_API_ENABLED=true
ENV ARK_GRAPHQL_ENABLED=true
ENV ARK_JSON_RPC_ENABLED=true

# public API, this one is for developers
EXPOSE 4003

# webhook port
EXPOSE 4004

# JSON-RPC
EXPOSE 8080

# public graphql API, including graphiQL explorer
EXPOSE 4005

# internal API, for nodes to communicate
EXPOSE 4000

# postgresql port, if you want to directly query the DB
EXPOSE 5432

COPY entrypoint.sh /

RUN echo "listen_addresses = '*'" >> /etc/postgresql/9.6/main/postgresql.conf
RUN echo "host all all 0.0.0.0/0 trust" >> /etc/postgresql/9.6/main/pg_hba.conf
RUN chmod a+x entrypoint.sh

# this will start an entire test network, including genesis block. To find the secrets, check out:
# https://github.com/ArkEcosystem/core/blob/master/packages/core/lib/config/testnet/delegates.json
ENTRYPOINT ./entrypoint.sh
