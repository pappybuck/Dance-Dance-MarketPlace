# Creates a connection to the Cockroachdb cluster and populates the database with the given schema.

import psycopg2
from psycopg2.extensions import ISOLATION_LEVEL_AUTOCOMMIT

conn = psycopg2.connect(
    host="cockroachdb-public.cockroachdb.svc.cluster.local",
    database="marketplace",
    user="Patrick",
    password="password",
    port="26257",
    sslmode="require",
    options="--cluster=cockroachdb-public"
)

conn.set_isolation_level(ISOLATION_LEVEL_AUTOCOMMIT)
with conn.cursor() as cur:
    cur.execute(open("migrationsV2.sql", "r").read())

cur.close()
conn.close()