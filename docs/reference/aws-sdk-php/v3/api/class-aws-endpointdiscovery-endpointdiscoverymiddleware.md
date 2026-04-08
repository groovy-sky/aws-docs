Menu

- [Aws](namespace-aws.md)
- [EndpointDiscovery](namespace-aws-endpointdiscovery.md)

## EndpointDiscoveryMiddleware        in package    - [Aws](package-aws.md)

### Table of Contents  [header link](class-aws-endpointdiscovery-endpointdiscoverymiddleware-toc.md)

#### Methods  [header link](class-aws-endpointdiscovery-endpointdiscoverymiddleware-toc-methods.md)

[\_\_construct()](class-aws-endpointdiscovery-endpointdiscoverymiddleware-method-construct.md)
: mixed [\_\_invoke()](class-aws-endpointdiscovery-endpointdiscoverymiddleware-method-invoke.md)
: mixed [wrap()](class-aws-endpointdiscovery-endpointdiscoverymiddleware-method-wrap.md)
: mixed

### Methods  [header link](class-aws-endpointdiscovery-endpointdiscoverymiddleware-methods.md)

#### \_\_construct()  [header link](class-aws-endpointdiscovery-endpointdiscoverymiddleware-method-construct.md)

`
    public
                    __construct(callable $handler, AwsClient $client, array<string|int, mixed> $args, mixed $config) : mixed`

##### Parameters

$handler
: callable$client
: [AwsClient](class-aws-awsclient.md)$args
: array<string\|int, mixed>$config
: mixed

#### \_\_invoke()  [header link](class-aws-endpointdiscovery-endpointdiscoverymiddleware-method-invoke.md)

`
    public
                    __invoke(CommandInterface $cmd, RequestInterface $request) : mixed`

##### Parameters

$cmd
: [CommandInterface](class-aws-commandinterface.md)$request
: [RequestInterface](class-psr-http-message-requestinterface.md)

#### wrap()  [header link](class-aws-endpointdiscovery-endpointdiscoverymiddleware-method-wrap.md)

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
  - [Methods](class-aws-endpointdiscovery-endpointdiscoverymiddleware-toc-methods.md)
- Methods
  - [\_\_construct()](class-aws-endpointdiscovery-endpointdiscoverymiddleware-method-construct.md)
  - [\_\_invoke()](class-aws-endpointdiscovery-endpointdiscoverymiddleware-method-invoke.md)
  - [wrap()](class-aws-endpointdiscovery-endpointdiscoverymiddleware-method-wrap.md)

[Back To Top](class-aws-endpointdiscovery-endpointdiscoverymiddleware-top.md)
