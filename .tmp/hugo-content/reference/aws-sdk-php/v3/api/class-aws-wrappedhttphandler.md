Menu

- [Aws](namespace-aws.md)

## WrappedHttpHandler        in package    - [Aws](package-aws.md)

Converts an HTTP handler into a Command HTTP handler.

HTTP handlers have the following signature:
function(RequestInterface $request, array $options) : PromiseInterface

The promise returned form an HTTP handler must resolve to a PSR-7 response
object when fulfilled or an error array when rejected. The error array
can contain the following data:

- exception: (required, Exception) Exception that was encountered.
- response: (ResponseInterface) PSR-7 response that was received (if a
response) was received.
- connection\_error: (bool) True if the error is the result of failing to
connect.

### Table of Contents  [header link](class-aws-wrappedhttphandler-toc.md)

#### Methods  [header link](class-aws-wrappedhttphandler-toc-methods.md)

[\_\_construct()](class-aws-wrappedhttphandler-method-construct.md)
: mixed [\_\_invoke()](class-aws-wrappedhttphandler-method-invoke.md)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Calls the simpler HTTP specific handler and wraps the returned promise
with AWS specific values (e.g., a result object or AWS exception).

### Methods  [header link](class-aws-wrappedhttphandler-methods.md)

#### \_\_construct()  [header link](class-aws-wrappedhttphandler-method-construct.md)

`
    public
                    __construct(callable $httpHandler, callable $parser, callable $errorParser[, string $exceptionClass = AwsException::class ][, bool $collectStats = false ]) : mixed`

##### Parameters

$httpHandler
: callable

Function that accepts a request and array
of request options and returns a promise
that fulfills with a response or rejects
with an error array.

$parser
: callable

Function that accepts a response object
and returns an AWS result object.

$errorParser
: callable

Function that parses a response object
into AWS error data.

$exceptionClass
: string
= AwsException::class

Exception class to throw.

$collectStats
: bool
= false

Whether to collect HTTP transfer
information.

#### \_\_invoke()  [header link](class-aws-wrappedhttphandler-method-invoke.md)

Calls the simpler HTTP specific handler and wraps the returned promise
with AWS specific values (e.g., a result object or AWS exception).

`
    public
                    __invoke(CommandInterface $command, RequestInterface $request) : PromiseInterface`

##### Parameters

$command
: [CommandInterface](class-aws-commandinterface.md)

Command being executed.

$request
: [RequestInterface](class-psr-http-message-requestinterface.md)

Request to send.

##### Return values

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](class-aws-wrappedhttphandler-toc-methods.md)
- Methods
  - [\_\_construct()](class-aws-wrappedhttphandler-method-construct.md)
  - [\_\_invoke()](class-aws-wrappedhttphandler-method-invoke.md)

[Back To Top](class-aws-wrappedhttphandler-top.md)
