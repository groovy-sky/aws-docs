Menu

- [Aws](namespace-aws.md)
- [ClientSideMonitoring](namespace-aws-clientsidemonitoring.md)

## Configuration        in package    - [Aws](package-aws.md)       implements  [ConfigurationInterface](class-aws-clientsidemonitoring-configurationinterface.md)

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ClientSideMonitoring.Configuration.html\#toc)

#### Interfaces  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ClientSideMonitoring.Configuration.html\#toc-interfaces)

[ConfigurationInterface](class-aws-clientsidemonitoring-configurationinterface.md)Provides access to client-side monitoring configuration options:
'client\_id', 'enabled', 'host', 'port'

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ClientSideMonitoring.Configuration.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ClientSideMonitoring.Configuration.html#method___construct)
: mixed Constructs a new Configuration object with the specified CSM options set.[getClientId()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ClientSideMonitoring.Configuration.html#method_getClientId)
: string\|null Returns the Client ID, if available.[getHost()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ClientSideMonitoring.Configuration.html#method_getHost)
: string\|null /{@inheritdoc}[getPort()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ClientSideMonitoring.Configuration.html#method_getPort)
: int\|null Returns the configured port.[isEnabled()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ClientSideMonitoring.Configuration.html#method_isEnabled)
: bool Checks whether or not client-side monitoring is enabled.[toArray()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ClientSideMonitoring.Configuration.html#method_toArray)
: array<string\|int, mixed> Returns the configuration as an associative array.

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ClientSideMonitoring.Configuration.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ClientSideMonitoring.Configuration.html\#method___construct)

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

#### getClientId()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ClientSideMonitoring.Configuration.html\#method_getClientId)

Returns the Client ID, if available.

`
    public
                    getClientId() : string|null`

##### Return values

string\|null

#### getHost()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ClientSideMonitoring.Configuration.html\#method_getHost)

/{@inheritdoc}

`
    public
                    getHost() : string|null`

##### Return values

string\|null

#### getPort()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ClientSideMonitoring.Configuration.html\#method_getPort)

Returns the configured port.

`
    public
                    getPort() : int|null`

##### Return values

int\|null

#### isEnabled()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ClientSideMonitoring.Configuration.html\#method_isEnabled)

Checks whether or not client-side monitoring is enabled.

`
    public
                    isEnabled() : bool`

##### Return values

bool

#### toArray()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ClientSideMonitoring.Configuration.html\#method_toArray)

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
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ClientSideMonitoring.Configuration.html#toc-methods)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ClientSideMonitoring.Configuration.html#method___construct)
  - [getClientId()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ClientSideMonitoring.Configuration.html#method_getClientId)
  - [getHost()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ClientSideMonitoring.Configuration.html#method_getHost)
  - [getPort()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ClientSideMonitoring.Configuration.html#method_getPort)
  - [isEnabled()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ClientSideMonitoring.Configuration.html#method_isEnabled)
  - [toArray()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ClientSideMonitoring.Configuration.html#method_toArray)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ClientSideMonitoring.Configuration.html#top)
