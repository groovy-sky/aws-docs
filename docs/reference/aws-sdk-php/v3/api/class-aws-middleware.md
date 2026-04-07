Menu

- [Aws](namespace-aws.md)

## Middleware        in package    - [Aws](package-aws.md)

FinalYes

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Middleware.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Middleware.html\#toc-methods)

[contentType()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Middleware.html#method_contentType)
: callable Middleware wrapper function that adds a Content-Type header to requests.[history()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Middleware.html#method_history)
: callable Tracks command and request history using a history container.[invocationId()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Middleware.html#method_invocationId)
: callable Middleware wrapper function that adds an invocation id header to
requests, which is only applied after the build step.[mapCommand()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Middleware.html#method_mapCommand)
: callable Creates a middleware that applies a map function to commands as they
pass through the middleware.[mapRequest()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Middleware.html#method_mapRequest)
: callable Creates a middleware that applies a map function to requests as they
pass through the middleware.[mapResult()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Middleware.html#method_mapResult)
: callable Creates a middleware that applies a map function to results.[recursionDetection()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Middleware.html#method_recursionDetection)
: callable Middleware wrapper function that adds a trace id header to requests
from clients instantiated in supported Lambda runtime environments.[requestBuilder()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Middleware.html#method_requestBuilder)
: callable Builds an HTTP request for a command.[retry()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Middleware.html#method_retry)
: callable Middleware wrapper function that retries requests based on the boolean
result of invoking the provided "decider" function.[signer()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Middleware.html#method_signer)
: callable Creates a middleware that signs requests for a command.[sourceFile()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Middleware.html#method_sourceFile)
: callable Middleware used to allow a command parameter (e.g., "SourceFile") to
be used to specify the source of data for an upload operation.[tap()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Middleware.html#method_tap)
: callable Creates a middleware that invokes a callback at a given step.[timer()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Middleware.html#method_timer)
: mixed [validation()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Middleware.html#method_validation)
: callable Adds a middleware that uses client-side validation.

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Middleware.html\#methods)

#### contentType()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Middleware.html\#method_contentType)

Middleware wrapper function that adds a Content-Type header to requests.

`
    public
            static        contentType(array<string|int, mixed> $operations) : callable`

This is only done when the Content-Type has not already been set, and the
request body's URI is available. It then checks the file extension of the
URI to determine the mime-type.

##### Parameters

$operations
: array<string\|int, mixed>

Operations that Content-Type should be added to.

##### Return values

callable

#### history()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Middleware.html\#method_history)

Tracks command and request history using a history container.

`
    public
            static        history(History $history) : callable`

This is useful for testing.

##### Parameters

$history
: [History](class-aws-history.md)

History container to store entries.

##### Return values

callable

#### invocationId()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Middleware.html\#method_invocationId)

Middleware wrapper function that adds an invocation id header to
requests, which is only applied after the build step.

`
    public
            static        invocationId() : callable`

This is a uniquely generated UUID to identify initial and subsequent
retries as part of a complete request lifecycle.

##### Return values

callable

#### mapCommand()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Middleware.html\#method_mapCommand)

Creates a middleware that applies a map function to commands as they
pass through the middleware.

`
    public
            static        mapCommand(callable $f) : callable`

##### Parameters

$f
: callable

Map function that accepts a command and returns a
command.

##### Return values

callable

#### mapRequest()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Middleware.html\#method_mapRequest)

Creates a middleware that applies a map function to requests as they
pass through the middleware.

`
    public
            static        mapRequest(callable $f) : callable`

##### Parameters

$f
: callable

Map function that accepts a RequestInterface and
returns a RequestInterface.

##### Return values

callable

#### mapResult()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Middleware.html\#method_mapResult)

Creates a middleware that applies a map function to results.

`
    public
            static        mapResult(callable $f) : callable`

##### Parameters

$f
: callable

Map function that accepts an Aws\\ResultInterface and
returns an Aws\\ResultInterface.

##### Return values

callable

#### recursionDetection()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Middleware.html\#method_recursionDetection)

Middleware wrapper function that adds a trace id header to requests
from clients instantiated in supported Lambda runtime environments.

`
    public
            static        recursionDetection() : callable`

The purpose for this header is to track and stop Lambda functions
from being recursively invoked due to misconfigured resources.

##### Return values

callable

#### requestBuilder()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Middleware.html\#method_requestBuilder)

Builds an HTTP request for a command.

`
    public
            static        requestBuilder(callable $serializer) : callable`

##### Parameters

$serializer
: callable

Function used to serialize a request for a
command.

##### Return values

callable

#### retry()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Middleware.html\#method_retry)

Middleware wrapper function that retries requests based on the boolean
result of invoking the provided "decider" function.

`
    public
            static        retry([callable $decider = null ][, callable $delay = null ][, bool $stats = false ]) : callable`

If no delay function is provided, a simple implementation of exponential
backoff will be utilized.

##### Parameters

$decider
: callable
= null

Function that accepts the number of retries,
a request, \[result\], and \[exception\] and
returns true if the command is to be retried.

$delay
: callable
= null

Function that accepts the number of retries and
returns the number of milliseconds to delay.

$stats
: bool
= false

Whether to collect statistics on retries and the
associated delay.

##### Return values

callable

#### signer()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Middleware.html\#method_signer)

Creates a middleware that signs requests for a command.

`
    public
            static        signer(callable $credProvider, callable $signatureFunction[, mixed $tokenProvider = null ][, mixed $config = [] ]) : callable`

##### Parameters

$credProvider
: callable

Credentials provider function that
returns a promise that is resolved
with a CredentialsInterface object.

$signatureFunction
: callable

Function that accepts a Command
object and returns a
SignatureInterface.

$tokenProvider
: mixed
= null$config
: mixed
= \[\]

##### Return values

callable

#### sourceFile()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Middleware.html\#method_sourceFile)

Middleware used to allow a command parameter (e.g., "SourceFile") to
be used to specify the source of data for an upload operation.

`
    public
            static        sourceFile(Service $api[, string $bodyParameter = 'Body' ][, string $sourceParameter = 'SourceFile' ]) : callable`

##### Parameters

$api
: [Service](class-aws-api-service.md)$bodyParameter
: string
= 'Body'$sourceParameter
: string
= 'SourceFile'

##### Return values

callable

#### tap()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Middleware.html\#method_tap)

Creates a middleware that invokes a callback at a given step.

`
    public
            static        tap(callable $fn) : callable`

The tap callback accepts a CommandInterface and RequestInterface as
arguments but is not expected to return a new value or proxy to
downstream middleware. It's simply a way to "tap" into the handler chain
to debug or get an intermediate value.

##### Parameters

$fn
: callable

Tap function

##### Return values

callable

#### timer()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Middleware.html\#method_timer)

`
    public
            static        timer() : mixed`

#### validation()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Middleware.html\#method_validation)

Adds a middleware that uses client-side validation.

`
    public
            static        validation(Service $api[, Validator|null $validator = null ]) : callable`

##### Parameters

$api
: [Service](class-aws-api-service.md)

API being accessed.

$validator
: [Validator](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Validator.html) \|null
= null

##### Return values

callable
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Middleware.html#toc-methods)
- Methods
  - [contentType()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Middleware.html#method_contentType)
  - [history()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Middleware.html#method_history)
  - [invocationId()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Middleware.html#method_invocationId)
  - [mapCommand()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Middleware.html#method_mapCommand)
  - [mapRequest()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Middleware.html#method_mapRequest)
  - [mapResult()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Middleware.html#method_mapResult)
  - [recursionDetection()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Middleware.html#method_recursionDetection)
  - [requestBuilder()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Middleware.html#method_requestBuilder)
  - [retry()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Middleware.html#method_retry)
  - [signer()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Middleware.html#method_signer)
  - [sourceFile()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Middleware.html#method_sourceFile)
  - [tap()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Middleware.html#method_tap)
  - [timer()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Middleware.html#method_timer)
  - [validation()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Middleware.html#method_validation)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Middleware.html#top)
