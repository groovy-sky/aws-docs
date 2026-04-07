Menu

- [Aws](namespace-aws.md)

## UserAgentMiddleware        in package    - [Aws](package-aws.md)

Builds and injects the user agent header values.

This middleware must be appended into step where all the
metrics to be gathered are already resolved. As of now it should be
after the signing step.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.UserAgentMiddleware.html\#toc)

#### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.UserAgentMiddleware.html\#toc-constants)

[AGENT\_VERSION](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.UserAgentMiddleware.html#constant_AGENT_VERSION)
= 2.1

#### Properties  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.UserAgentMiddleware.html\#toc-properties)

[$userAgentFnList](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.UserAgentMiddleware.html#property_userAgentFnList)
: mixed

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.UserAgentMiddleware.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.UserAgentMiddleware.html#method___construct)
: mixed [\_\_invoke()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.UserAgentMiddleware.html#method___invoke)
: mixed When invoked, its injects the user agent header into the
request headers.[wrap()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.UserAgentMiddleware.html#method_wrap)
: ClosureReturns a middleware wrapper function.

### Constants  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.UserAgentMiddleware.html\#constants)

#### AGENT\_VERSION  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.UserAgentMiddleware.html\#constant_AGENT_VERSION)

`
    public
        mixed
    AGENT_VERSION
    = 2.1
`

### Properties  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.UserAgentMiddleware.html\#properties)

#### $userAgentFnList  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.UserAgentMiddleware.html\#property_userAgentFnList)

`
    public
    static    mixed
    $userAgentFnList
     = ['getSdkVersion', 'getUserAgentVersion', 'getHhvmVersion', 'getOsName', 'getLangVersion', 'getExecEnv', 'getEndpointDiscovery', 'getAppId', 'getMetrics']`

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.UserAgentMiddleware.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.UserAgentMiddleware.html\#method___construct)

`
    public
                    __construct(callable $nextHandler[, array<string|int, mixed> $args = [] ]) : mixed`

##### Parameters

$nextHandler
: callable$args
: array<string\|int, mixed>
= \[\]

#### \_\_invoke()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.UserAgentMiddleware.html\#method___invoke)

When invoked, its injects the user agent header into the
request headers.

`
    public
                    __invoke(CommandInterface $command, RequestInterface $request) : mixed`

##### Parameters

$command
: [CommandInterface](class-aws-commandinterface.md)$request
: [RequestInterface](class-psr-http-message-requestinterface.md)

#### wrap()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.UserAgentMiddleware.html\#method_wrap)

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
  - [Constants](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.UserAgentMiddleware.html#toc-constants)
  - [Properties](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.UserAgentMiddleware.html#toc-properties)
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.UserAgentMiddleware.html#toc-methods)
- Constants
  - [AGENT\_VERSION](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.UserAgentMiddleware.html#constant_AGENT_VERSION)
- Properties
  - [$userAgentFnList](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.UserAgentMiddleware.html#property_userAgentFnList)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.UserAgentMiddleware.html#method___construct)
  - [\_\_invoke()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.UserAgentMiddleware.html#method___invoke)
  - [wrap()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.UserAgentMiddleware.html#method_wrap)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.UserAgentMiddleware.html#top)
