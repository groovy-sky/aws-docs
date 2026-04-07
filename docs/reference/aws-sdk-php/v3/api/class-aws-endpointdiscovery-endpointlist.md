Menu

- [Aws](namespace-aws.md)
- [EndpointDiscovery](namespace-aws-endpointdiscovery.md)

## EndpointList        in package    - [Aws](package-aws.md)

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointDiscovery.EndpointList.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointDiscovery.EndpointList.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointDiscovery.EndpointList.html#method___construct)
: mixed [getActive()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointDiscovery.EndpointList.html#method_getActive)
: null\|string Gets an active (unexpired) endpoint. Returns null if none found.[getEndpoint()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointDiscovery.EndpointList.html#method_getEndpoint)
: null\|string Gets an active endpoint if possible, then an expired endpoint if possible.[remove()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointDiscovery.EndpointList.html#method_remove)
: mixed Removes an endpoint from both lists.

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointDiscovery.EndpointList.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointDiscovery.EndpointList.html\#method___construct)

`
    public
                    __construct(array<string|int, mixed> $endpoints) : mixed`

##### Parameters

$endpoints
: array<string\|int, mixed>

#### getActive()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointDiscovery.EndpointList.html\#method_getActive)

Gets an active (unexpired) endpoint. Returns null if none found.

`
    public
                    getActive() : null|string`

##### Return values

null\|string

#### getEndpoint()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointDiscovery.EndpointList.html\#method_getEndpoint)

Gets an active endpoint if possible, then an expired endpoint if possible.

`
    public
                    getEndpoint() : null|string`

Returns null if no endpoints found.

##### Return values

null\|string

#### remove()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointDiscovery.EndpointList.html\#method_remove)

Removes an endpoint from both lists.

`
    public
                    remove(string $key) : mixed`

##### Parameters

$key
: string
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointDiscovery.EndpointList.html#toc-methods)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointDiscovery.EndpointList.html#method___construct)
  - [getActive()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointDiscovery.EndpointList.html#method_getActive)
  - [getEndpoint()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointDiscovery.EndpointList.html#method_getEndpoint)
  - [remove()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointDiscovery.EndpointList.html#method_remove)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointDiscovery.EndpointList.html#top)
