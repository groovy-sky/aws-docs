Menu

- [Aws](namespace-aws.md)
- [Endpoint](namespace-aws-endpoint.md)

## EndpointProvider        in package    - [Aws](package-aws.md)

Endpoint providers.

An endpoint provider is a function that accepts a hash of endpoint options,
including but not limited to "service" and "region" key value pairs. The
endpoint provider function returns a hash of endpoint data, which MUST
include an "endpoint" key value pair that represents the resolved endpoint
or NULL if an endpoint cannot be determined.

You can wrap your calls to an endpoint provider with the
EndpointProvider::resolve function to ensure that an endpoint hash is
created. If an endpoint hash is not created, then the resolve() function
will throw an UnresolvedEndpointException.

```prettyprint
use Aws\Endpoint\EndpointProvider;
$provider = EndpointProvider::defaultProvider();
// Returns an array or NULL.
$endpoint = $provider(['service' => 'ec2', 'region' => 'us-west-2']);
// Returns an endpoint array or throws.
$endpoint = EndpointProvider::resolve($provider, [
    'service' => 'ec2',
    'region'  => 'us-west-2'
]);

```

You can compose multiple providers into a single provider using
or\_chain. This function accepts providers as arguments and
returns a new function that will invoke each provider until a non-null value
is returned.

```prettyprint
$a = function (array $args) {
    if ($args['region'] === 'my-test-region') {
        return ['endpoint' => 'http://localhost:123/api'];
    }
};
$b = EndpointProvider::defaultProvider();
$c = \Aws\or_chain($a, $b);
$config = ['service' => 'ec2', 'region' => 'my-test-region'];
$res = $c($config);  // $a handles this.
$config['region'] = 'us-west-2';
$res = $c($config); // $b handles this.

```

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.EndpointProvider.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.EndpointProvider.html\#toc-methods)

[defaultProvider()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.EndpointProvider.html#method_defaultProvider)
: callable Creates and returns the default SDK endpoint provider.[patterns()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.EndpointProvider.html#method_patterns)
: callable Creates and returns an endpoint provider that uses patterns from an
array.[resolve()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.EndpointProvider.html#method_resolve)
: array<string\|int, mixed> Resolves and endpoint provider and ensures a non-null return value.

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.EndpointProvider.html\#methods)

#### defaultProvider()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.EndpointProvider.html\#method_defaultProvider)

Creates and returns the default SDK endpoint provider.

`
    public
            static        defaultProvider() : callable`

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.EndpointProvider.html\#method_defaultProvider\#tags)

deprecated

Use an instance of \\Aws\\Endpoint\\Partition instead.

##### Return values

callable

#### patterns()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.EndpointProvider.html\#method_patterns)

Creates and returns an endpoint provider that uses patterns from an
array.

`
    public
            static        patterns(array<string|int, mixed> $patterns) : callable`

##### Parameters

$patterns
: array<string\|int, mixed>

Endpoint patterns

##### Return values

callable

#### resolve()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.EndpointProvider.html\#method_resolve)

Resolves and endpoint provider and ensures a non-null return value.

`
    public
            static        resolve(callable $provider[, array<string|int, mixed> $args = [] ]) : array<string|int, mixed>`

##### Parameters

$provider
: callable

Provider function to invoke.

$args
: array<string\|int, mixed>
= \[\]

Endpoint arguments to pass to the provider.

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.EndpointProvider.html\#method_resolve\#tags)

throws[UnresolvedEndpointException](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Exception.UnresolvedEndpointException.html)

##### Return values

array<string\|int, mixed>
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.EndpointProvider.html#toc-methods)
- Methods
  - [defaultProvider()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.EndpointProvider.html#method_defaultProvider)
  - [patterns()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.EndpointProvider.html#method_patterns)
  - [resolve()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.EndpointProvider.html#method_resolve)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Endpoint.EndpointProvider.html#top)
