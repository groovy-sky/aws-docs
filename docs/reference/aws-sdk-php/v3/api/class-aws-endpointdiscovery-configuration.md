Menu

- [Aws](namespace-aws.md)
- [EndpointDiscovery](namespace-aws-endpointdiscovery.md)

## Configuration        in package    - [Aws](package-aws.md)       implements  [ConfigurationInterface](class-aws-endpointdiscovery-configurationinterface.md)

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointDiscovery.Configuration.html\#toc)

#### Interfaces  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointDiscovery.Configuration.html\#toc-interfaces)

[ConfigurationInterface](class-aws-endpointdiscovery-configurationinterface.md)Provides access to endpoint discovery configuration options:
'enabled', 'cache\_limit'

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointDiscovery.Configuration.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointDiscovery.Configuration.html#method___construct)
: mixed [getCacheLimit()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointDiscovery.Configuration.html#method_getCacheLimit)
: string\|null Returns the cache limit, if available.[isEnabled()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointDiscovery.Configuration.html#method_isEnabled)
: bool Checks whether or not endpoint discovery is enabled.[toArray()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointDiscovery.Configuration.html#method_toArray)
: array<string\|int, mixed> Returns the configuration as an associative array

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointDiscovery.Configuration.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointDiscovery.Configuration.html\#method___construct)

`
    public
                    __construct(mixed $enabled[, mixed $cacheLimit = 1000 ]) : mixed`

##### Parameters

$enabled
: mixed$cacheLimit
: mixed
= 1000

#### getCacheLimit()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointDiscovery.Configuration.html\#method_getCacheLimit)

Returns the cache limit, if available.

`
    public
                    getCacheLimit() : string|null`

##### Return values

string\|null

#### isEnabled()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointDiscovery.Configuration.html\#method_isEnabled)

Checks whether or not endpoint discovery is enabled.

`
    public
                    isEnabled() : bool`

##### Return values

bool

#### toArray()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointDiscovery.Configuration.html\#method_toArray)

Returns the configuration as an associative array

`
    public
                    toArray() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointDiscovery.Configuration.html#toc-methods)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointDiscovery.Configuration.html#method___construct)
  - [getCacheLimit()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointDiscovery.Configuration.html#method_getCacheLimit)
  - [isEnabled()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointDiscovery.Configuration.html#method_isEnabled)
  - [toArray()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointDiscovery.Configuration.html#method_toArray)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointDiscovery.Configuration.html#top)
