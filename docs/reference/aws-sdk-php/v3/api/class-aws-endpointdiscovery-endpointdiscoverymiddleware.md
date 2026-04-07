Menu

- [Aws](namespace-aws.md)
- [EndpointDiscovery](namespace-aws-endpointdiscovery.md)

## EndpointDiscoveryMiddleware        in package    - [Aws](package-aws.md)

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointDiscovery.EndpointDiscoveryMiddleware.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointDiscovery.EndpointDiscoveryMiddleware.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointDiscovery.EndpointDiscoveryMiddleware.html#method___construct)
: mixed [\_\_invoke()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointDiscovery.EndpointDiscoveryMiddleware.html#method___invoke)
: mixed [wrap()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointDiscovery.EndpointDiscoveryMiddleware.html#method_wrap)
: mixed

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointDiscovery.EndpointDiscoveryMiddleware.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointDiscovery.EndpointDiscoveryMiddleware.html\#method___construct)

`
    public
                    __construct(callable $handler, AwsClient $client, array<string|int, mixed> $args, mixed $config) : mixed`

##### Parameters

$handler
: callable$client
: [AwsClient](class-aws-awsclient.md)$args
: array<string\|int, mixed>$config
: mixed

#### \_\_invoke()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointDiscovery.EndpointDiscoveryMiddleware.html\#method___invoke)

`
    public
                    __invoke(CommandInterface $cmd, RequestInterface $request) : mixed`

##### Parameters

$cmd
: [CommandInterface](class-aws-commandinterface.md)$request
: [RequestInterface](class-psr-http-message-requestinterface.md)

#### wrap()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointDiscovery.EndpointDiscoveryMiddleware.html\#method_wrap)

`
    public
            static        wrap(mixed $client, mixed $args, mixed $config) : mixed`

##### Parameters

$client
: mixed$args
: mixed$config
: mixed
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointDiscovery.EndpointDiscoveryMiddleware.html#toc-methods)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointDiscovery.EndpointDiscoveryMiddleware.html#method___construct)
  - [\_\_invoke()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointDiscovery.EndpointDiscoveryMiddleware.html#method___invoke)
  - [wrap()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointDiscovery.EndpointDiscoveryMiddleware.html#method_wrap)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.EndpointDiscovery.EndpointDiscoveryMiddleware.html#top)
