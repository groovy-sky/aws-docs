Menu

- [Aws](namespace-aws.md)
- [EndpointDiscovery](namespace-aws-endpointdiscovery.md)

## EndpointList        in package    - [Aws](package-aws.md)

### Table of Contents  [header link](class-aws-endpointdiscovery-endpointlist-toc.md)

#### Methods  [header link](class-aws-endpointdiscovery-endpointlist-toc-methods.md)

[\_\_construct()](class-aws-endpointdiscovery-endpointlist-method-construct.md)
: mixed [getActive()](class-aws-endpointdiscovery-endpointlist-method-getactive.md)
: null\|string Gets an active (unexpired) endpoint. Returns null if none found.[getEndpoint()](class-aws-endpointdiscovery-endpointlist-method-getendpoint.md)
: null\|string Gets an active endpoint if possible, then an expired endpoint if possible.[remove()](class-aws-endpointdiscovery-endpointlist-method-remove.md)
: mixed Removes an endpoint from both lists.

### Methods  [header link](class-aws-endpointdiscovery-endpointlist-methods.md)

#### \_\_construct()  [header link](class-aws-endpointdiscovery-endpointlist-method-construct.md)

`
    public
                    __construct(array<string|int, mixed> $endpoints) : mixed`

##### Parameters

$endpoints
: array<string\|int, mixed>

#### getActive()  [header link](class-aws-endpointdiscovery-endpointlist-method-getactive.md)

Gets an active (unexpired) endpoint. Returns null if none found.

`
    public
                    getActive() : null|string`

##### Return values

null\|string

#### getEndpoint()  [header link](class-aws-endpointdiscovery-endpointlist-method-getendpoint.md)

Gets an active endpoint if possible, then an expired endpoint if possible.

`
    public
                    getEndpoint() : null|string`

Returns null if no endpoints found.

##### Return values

null\|string

#### remove()  [header link](class-aws-endpointdiscovery-endpointlist-method-remove.md)

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
  - [Methods](class-aws-endpointdiscovery-endpointlist-toc-methods.md)
- Methods
  - [\_\_construct()](class-aws-endpointdiscovery-endpointlist-method-construct.md)
  - [getActive()](class-aws-endpointdiscovery-endpointlist-method-getactive.md)
  - [getEndpoint()](class-aws-endpointdiscovery-endpointlist-method-getendpoint.md)
  - [remove()](class-aws-endpointdiscovery-endpointlist-method-remove.md)

[Back To Top](class-aws-endpointdiscovery-endpointlist-top.md)
