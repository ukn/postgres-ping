# postgres-ping

Simple postgres ping tool.

Could be useful to track db reachability during maintenance work.

# Configuration

Use the following environment variables can be set to override default connection params:

user - username (postgres)
pass - password (empty string)
host - db hostname or ip (127.0.0.1)
port - db port (5432)
dbname - database name (postgres)
sslmode - ssl mode (disable)
timeout - connection timeout (2s)
verbose - verbose console logging (false)
interval - ping interval (2s)
