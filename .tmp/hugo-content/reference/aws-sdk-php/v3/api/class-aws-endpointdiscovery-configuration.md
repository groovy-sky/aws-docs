Menu

- [Aws](namespace-aws.md)
- [EndpointDiscovery](namespace-aws-endpointdiscovery.md)

## Configuration        in package    - [Aws](package-aws.md)       implements  [ConfigurationInterface](class-aws-endpointdiscovery-configurationinterface.md)

### Table of Contents  [header link](class-aws-endpointdiscovery-configuration-toc.md)

#### Interfaces  [header link](class-aws-endpointdiscovery-configuration-toc-interfaces.md)

[ConfigurationInterface](class-aws-endpointdiscovery-configurationinterface.md)Provides access to endpoint discovery configuration options:
'enabled', 'cache\_limit'

#### Methods  [header link](class-aws-endpointdiscovery-configuration-toc-methods.md)

[\_\_construct()](class-aws-endpointdiscovery-configuration-method-construct.md)
: mixed [getCacheLimit()](class-aws-endpointdiscovery-configuration-method-getcachelimit.md)
: string\|null Returns the cache limit, if available.[isEnabled()](class-aws-endpointdiscovery-configuration-method-isenabled.md)
: bool Checks whether or not endpoint discovery is enabled.[toArray()](class-aws-endpointdiscovery-configuration-method-toarray.md)
: array<string\|int, mixed> Returns the configuration as an associative array

### Methods  [header link](class-aws-endpointdiscovery-configuration-methods.md)

#### \_\_construct()  [header link](class-aws-endpointdiscovery-configuration-method-construct.md)

`
    public
                    __construct(mixed $enabled[, mixed $cacheLimit = 1000 ]) : mixed`

##### Parameters

$enabled
: mixed$cacheLimit
: mixed
= 1000

#### getCacheLimit()  [header link](class-aws-endpointdiscovery-configuration-method-getcachelimit.md)

Returns the cache limit, if available.

`
    public
                    getCacheLimit() : string|null`

##### Return values

string\|null

#### isEnabled()  [header link](class-aws-endpointdiscovery-configuration-method-isenabled.md)

Checks whether or not endpoint discovery is enabled.

`
    public
                    isEnabled() : bool`

##### Return values

bool

#### toArray()  [header link](class-aws-endpointdiscovery-configuration-method-toarray.md)

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
  - [Methods](class-aws-endpointdiscovery-configuration-toc-methods.md)
- Methods
  - [\_\_construct()](class-aws-endpointdiscovery-configuration-method-construct.md)
  - [getCacheLimit()](class-aws-endpointdiscovery-configuration-method-getcachelimit.md)
  - [isEnabled()](class-aws-endpointdiscovery-configuration-method-isenabled.md)
  - [toArray()](class-aws-endpointdiscovery-configuration-method-toarray.md)

[Back To Top](class-aws-endpointdiscovery-configuration-top.md)
