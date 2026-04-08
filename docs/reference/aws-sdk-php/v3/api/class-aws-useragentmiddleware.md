Menu

- [Aws](namespace-aws.md)

## UserAgentMiddleware        in package    - [Aws](package-aws.md)

Builds and injects the user agent header values.

This middleware must be appended into step where all the
metrics to be gathered are already resolved. As of now it should be
after the signing step.

### Table of Contents  [header link](class-aws-useragentmiddleware-toc.md)

#### Constants  [header link](class-aws-useragentmiddleware-toc-constants.md)

[AGENT\_VERSION](class-aws-useragentmiddleware-constant-agent-version.md)
= 2.1

#### Properties  [header link](class-aws-useragentmiddleware-toc-properties.md)

[$userAgentFnList](class-aws-useragentmiddleware-property-useragentfnlist.md)
: mixed

#### Methods  [header link](class-aws-useragentmiddleware-toc-methods.md)

[\_\_construct()](class-aws-useragentmiddleware-method-construct.md)
: mixed [\_\_invoke()](class-aws-useragentmiddleware-method-invoke.md)
: mixed When invoked, its injects the user agent header into the
request headers.[wrap()](class-aws-useragentmiddleware-method-wrap.md)
: ClosureReturns a middleware wrapper function.

### Constants  [header link](class-aws-useragentmiddleware-constants.md)

#### AGENT\_VERSION  [header link](class-aws-useragentmiddleware-constant-agent-version.md)

`
    public
        mixed
    AGENT_VERSION
    = 2.1
`

### Properties  [header link](class-aws-useragentmiddleware-properties.md)

#### $userAgentFnList  [header link](class-aws-useragentmiddleware-property-useragentfnlist.md)

`
    public
    static    mixed
    $userAgentFnList
     = ['getSdkVersion', 'getUserAgentVersion', 'getHhvmVersion', 'getOsName', 'getLangVersion', 'getExecEnv', 'getEndpointDiscovery', 'getAppId', 'getMetrics']`

### Methods  [header link](class-aws-useragentmiddleware-methods.md)

#### \_\_construct()  [header link](class-aws-useragentmiddleware-method-construct.md)

`
    public
                    __construct(callable $nextHandler[, array<string|int, mixed> $args = [] ]) : mixed`

##### Parameters

$nextHandler
: callable$args
: array<string\|int, mixed>
= \[\]

#### \_\_invoke()  [header link](class-aws-useragentmiddleware-method-invoke.md)

When invoked, its injects the user agent header into the
request headers.

`
    public
                    __invoke(CommandInterface $command, RequestInterface $request) : mixed`

##### Parameters

$command
: [CommandInterface](class-aws-commandinterface.md)$request
: [RequestInterface](class-psr-http-message-requestinterface.md)

#### wrap()  [header link](class-aws-useragentmiddleware-method-wrap.md)

Returns a middleware wrapper function.

`
    public
            static        wrap(array<string|int, mixed> $args) : Closure`

##### Parameters

$args
: array<string\|int, mixed>

##### Return values

Closure
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](class-aws-useragentmiddleware-toc-constants.md)
  - [Properties](class-aws-useragentmiddleware-toc-properties.md)
  - [Methods](class-aws-useragentmiddleware-toc-methods.md)
- Constants
  - [AGENT\_VERSION](class-aws-useragentmiddleware-constant-agent-version.md)
- Properties
  - [$userAgentFnList](class-aws-useragentmiddleware-property-useragentfnlist.md)
- Methods
  - [\_\_construct()](class-aws-useragentmiddleware-method-construct.md)
  - [\_\_invoke()](class-aws-useragentmiddleware-method-invoke.md)
  - [wrap()](class-aws-useragentmiddleware-method-wrap.md)

[Back To Top](class-aws-useragentmiddleware-top.md)
