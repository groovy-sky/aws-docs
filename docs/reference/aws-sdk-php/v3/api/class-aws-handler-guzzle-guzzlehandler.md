Menu

- [Aws](namespace-aws.md)
- [Handler](namespace-aws-handler.md)
- [Guzzle](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.handler.guzzle.html)

## GuzzleHandler        in package    - [Aws](package-aws.md)

A request handler that sends PSR-7-compatible requests with Guzzle.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Handler.Guzzle.GuzzleHandler.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Handler.Guzzle.GuzzleHandler.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Handler.Guzzle.GuzzleHandler.html#method___construct)
: mixed [\_\_invoke()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Handler.Guzzle.GuzzleHandler.html#method___invoke)
: [Promise](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.Promise.html)

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Handler.Guzzle.GuzzleHandler.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Handler.Guzzle.GuzzleHandler.html\#method___construct)

`
    public
                    __construct([ClientInterface $client = null ]) : mixed`

##### Parameters

$client
: ClientInterface
= null

#### \_\_invoke()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Handler.Guzzle.GuzzleHandler.html\#method___invoke)

`
    public
                    __invoke(RequestInterface $request[, array<string|int, mixed> $options = [] ]) : Promise`

##### Parameters

$request
: [RequestInterface](class-psr-http-message-requestinterface.md)$options
: array<string\|int, mixed>
= \[\]

##### Return values

[Promise](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.Promise.html)
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Handler.Guzzle.GuzzleHandler.html#toc-methods)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Handler.Guzzle.GuzzleHandler.html#method___construct)
  - [\_\_invoke()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Handler.Guzzle.GuzzleHandler.html#method___invoke)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Handler.Guzzle.GuzzleHandler.html#top)
