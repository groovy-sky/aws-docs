Menu

- [Aws](namespace-aws.md)
- [ClientSideMonitoring](namespace-aws-clientsidemonitoring.md)

## Configuration        in package    - [Aws](package-aws.md)       implements  [ConfigurationInterface](class-aws-clientsidemonitoring-configurationinterface.md)

### Table of Contents  [header link](class-aws-clientsidemonitoring-configuration-toc.md)

#### Interfaces  [header link](class-aws-clientsidemonitoring-configuration-toc-interfaces.md)

[ConfigurationInterface](class-aws-clientsidemonitoring-configurationinterface.md)Provides access to client-side monitoring configuration options:
'client\_id', 'enabled', 'host', 'port'

#### Methods  [header link](class-aws-clientsidemonitoring-configuration-toc-methods.md)

[\_\_construct()](class-aws-clientsidemonitoring-configuration-method-construct.md)
: mixed Constructs a new Configuration object with the specified CSM options set.[getClientId()](class-aws-clientsidemonitoring-configuration-method-getclientid.md)
: string\|null Returns the Client ID, if available.[getHost()](class-aws-clientsidemonitoring-configuration-method-gethost.md)
: string\|null /{@inheritdoc}[getPort()](class-aws-clientsidemonitoring-configuration-method-getport.md)
: int\|null Returns the configured port.[isEnabled()](class-aws-clientsidemonitoring-configuration-method-isenabled.md)
: bool Checks whether or not client-side monitoring is enabled.[toArray()](class-aws-clientsidemonitoring-configuration-method-toarray.md)
: array<string\|int, mixed> Returns the configuration as an associative array.

### Methods  [header link](class-aws-clientsidemonitoring-configuration-methods.md)

#### \_\_construct()  [header link](class-aws-clientsidemonitoring-configuration-method-construct.md)

Constructs a new Configuration object with the specified CSM options set.

`
    public
                    __construct(mixed $enabled, string $host, string|int $port[, string $clientId = '' ]) : mixed`

##### Parameters

$enabled
: mixed$host
: string$port
: string\|int$clientId
: string
= ''

#### getClientId()  [header link](class-aws-clientsidemonitoring-configuration-method-getclientid.md)

Returns the Client ID, if available.

`
    public
                    getClientId() : string|null`

##### Return values

string\|null

#### getHost()  [header link](class-aws-clientsidemonitoring-configuration-method-gethost.md)

/{@inheritdoc}

`
    public
                    getHost() : string|null`

##### Return values

string\|null

#### getPort()  [header link](class-aws-clientsidemonitoring-configuration-method-getport.md)

Returns the configured port.

`
    public
                    getPort() : int|null`

##### Return values

int\|null

#### isEnabled()  [header link](class-aws-clientsidemonitoring-configuration-method-isenabled.md)

Checks whether or not client-side monitoring is enabled.

`
    public
                    isEnabled() : bool`

##### Return values

bool

#### toArray()  [header link](class-aws-clientsidemonitoring-configuration-method-toarray.md)

Returns the configuration as an associative array.

`
    public
                    toArray() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](class-aws-clientsidemonitoring-configuration-toc-methods.md)
- Methods
  - [\_\_construct()](class-aws-clientsidemonitoring-configuration-method-construct.md)
  - [getClientId()](class-aws-clientsidemonitoring-configuration-method-getclientid.md)
  - [getHost()](class-aws-clientsidemonitoring-configuration-method-gethost.md)
  - [getPort()](class-aws-clientsidemonitoring-configuration-method-getport.md)
  - [isEnabled()](class-aws-clientsidemonitoring-configuration-method-isenabled.md)
  - [toArray()](class-aws-clientsidemonitoring-configuration-method-toarray.md)

[Back To Top](class-aws-clientsidemonitoring-configuration-top.md)
