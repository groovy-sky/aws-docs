# (Optional) Using environment variables to configure AWS AppConfig Agent for Amazon ECS and Amazon EKS

You can configure AWS AppConfig Agent by changing the following environment variables for
your agent container.

###### Note

The following table includes a **Sample values** column. Depending
on your monitor resolution, you might need to scroll to the bottom of the table and then
scroll to the right to view the column.

Environment variableDetailsDefault valueSample value(s)

`ACCESS_TOKEN`

This environment variable defines a token that must be provided when
requesting configuration data from the agent HTTP server. The value of the token
must be set in the HTTP request authorization header with an authorization type
of `Bearer`. Here is an example.

```

GET /applications/my_app/...
                  Host: localhost:2772
                  Authorization: Bearer <token value>
```

NoneMyAccessToken

`BACKUP_DIRECTORY`

This environment variable enables AWS AppConfig Agent to save a backup of each
configuration it retrieves to the specified directory.

###### Important

Configurations backed up to disk are not encrypted. If your configuration
contains sensitive data, AWS AppConfig recommends that you practice the principle of
least privilege with your filesystem permissions. For more information, see
[Security in AWS AppConfig](appconfig-security.md).

None/path/to/backups

`HTTP_PORT`

This environment variable specifies the port on which the HTTP server for
the agent runs.

27722772

`HTTP_HOST`

The HTTP\_HOST variable controls how the AWS AppConfig Agent binds to network interfaces. The binding behavior differs based on the runtime environment to ensure optimal security and accessibility.

ECS, EKS

- Default binding: All network interfaces (0.0.0.0)

EC2 and on-prem

- Default binding: localhost only

- IPv4 address: 127.0.0.1:2772

- IPv6 address: \[::1\]:2772

Custom Configuration Options. You can override the default behavior using these values:

- ` all` (binds to all interfaces)

- `localhost`(explicitly binds to localhost interfaces)

- Specific IP address (e.g `192.168.1.1`)

- Custom hostname

`LOG_LEVEL`

This environment variable specifies the level of detail that the agent logs.
Each level includes the current level and all higher levels. The value is case
insensitive. From most to least detailed, the log levels are: `trace`, `debug`,
`info`, `warn`, `error`, `fatal`, and `none`. The `trace` log includes detailed information,
including timing information, about the agent.

info

trace

debug

info

warn

error

fatal

none

`LOG_PATH`

The disk location where logs are written. If not specified, logs are written
to stderr.

None

/path/to/logs/agent.log

`MANIFEST`

This environment variable configures AWS AppConfig Agent to take advantage of
additional per-configuration features like multi-account retrievals and save
configuration to disk. For more information about these features, see [Using a manifest to enable additional retrieval features](appconfig-agent-how-to-use-additional-features.md).

None

When using AWS AppConfig configuration as manifest:
`MyApp:MyEnv:MyManifestConfig`.

When loading manifest
from disk: `file:/path/to/manifest.json`

`MAX_CONNECTIONS`

This environment variable configures the maximum number of connections that
the agent uses to retrieve configurations from AWS AppConfig.

33

`POLL_INTERVAL`

This environment variable controls how often the agent polls AWS AppConfig for
updated configuration data. You can specify a number of seconds for the
interval. You can also specify a number with a time unit: s for seconds, m for
minutes, and h for hours. If a unit isn't specified, the agent defaults to
seconds. For example, 60, 60s, and 1m result in the same poll interval.

45 seconds

45

45s

5m

1h

`PREFETCH_LIST`

This environment variable specifies the configuration data the agent requests from AWS AppConfig as soon as it starts.
Multiple configuration identifiers may be provided in a comma-separated list.

None

MyApp:MyEnv:MyConfig

abcd123:efgh456:ijkl789

MyApp:MyEnv:Config1,MyApp:MyEnv:Config2

`PRELOAD_BACKUPS`

If set to `true`, AWS AppConfig Agent loads configuration backups found
in the `BACKUP_DIRECTORY` into memory and immediately checks to see
if a newer version exists from the service. If set to `false`, AWS AppConfig
Agent only loads the contents from a configuration backup if it cannot retrieve
configuration data from the service, for example if there is a problem with your
network.

true

true

false

`PROXY_HEADERS`This environment variable specifies headers that are required by the proxy
referenced in the `PROXY_URL` environment variable. The value is a
comma-separated list of headers. None

header: value

h1: v1, h2: v2

`PROXY_URL`This environment variable specifies the proxy URL to use for connections from
the agent to AWS services, including AWS AppConfig. `HTTPS` and
`HTTP` URLs are supported.None

http://localhost:7474

https://my-proxy.example.com

`REQUEST_TIMEOUT`

This environment variable controls the amount of time the agent waits for a
response from AWS AppConfig. If the service does not respond, the request fails.

If the request is for the initial data retrieval, the agent returns an error
to your application.

If the timeout occurs during a background check for updated data, the agent
logs the error and tries again after a short delay.

You can specify the number of milliseconds for the timeout. You can also
specify a number with a time unit: ms for milliseconds and s for seconds. If a
unit isn't specified, the agent defaults to milliseconds. As an example, 5000,
5000ms and 5s result in the same request timeout value.

3000ms

3000

3000ms

5s

`ROLE_ARN`This environment variable specifies the Amazon Resource Name (ARN) of an
IAM role. AWS AppConfig Agent assumes this role to retrieve configuration data.Nonearn:aws:iam::123456789012:role/MyRole`ROLE_EXTERNAL_ID`This environment variable specifies the external ID to use with the assumed
role ARN.NoneMyExternalId`ROLE_SESSION_NAME`This environment variable specifies the session name to be associated with
the credentials for the assumed IAM role.NoneAWSAppConfigAgentSession`SERVICE_REGION`This environment variable specifies an alternative AWS Region that AWS AppConfig
Agent uses to call the AWS AppConfig service. If left undefined, the agent attempts to
determine the current Region. If it can't, the agent fails to start.None

us-east-1

eu-west-1

`WAIT_ON_MANIFEST`

This environment variable configures AWS AppConfig Agent to wait until the manifest
is processed before completing startup.

true

true

false

[Document Conventions](../../../../general/latest/gr/docconventions.md)

(Optional) Running AWS AppConfig as a DaemonSet in Amazon EKS

Retrieving configuration data for applications running in Amazon ECS and Amazon EKS

All content copied from https://docs.aws.amazon.com/.
