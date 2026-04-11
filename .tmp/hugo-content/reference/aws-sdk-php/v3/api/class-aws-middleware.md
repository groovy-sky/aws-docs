Menu

- [Aws](namespace-aws.md)

## Middleware        in package    - [Aws](package-aws.md)

FinalYes

### Table of Contents  [header link](class-aws-middleware-toc.md)

#### Methods  [header link](class-aws-middleware-toc-methods.md)

[contentType()](class-aws-middleware-method-contenttype.md)
: callable Middleware wrapper function that adds a Content-Type header to requests.[history()](class-aws-middleware-method-history.md)
: callable Tracks command and request history using a history container.[invocationId()](class-aws-middleware-method-invocationid.md)
: callable Middleware wrapper function that adds an invocation id header to
requests, which is only applied after the build step.[mapCommand()](class-aws-middleware-method-mapcommand.md)
: callable Creates a middleware that applies a map function to commands as they
pass through the middleware.[mapRequest()](class-aws-middleware-method-maprequest.md)
: callable Creates a middleware that applies a map function to requests as they
pass through the middleware.[mapResult()](class-aws-middleware-method-mapresult.md)
: callable Creates a middleware that applies a map function to results.[recursionDetection()](class-aws-middleware-method-recursiondetection.md)
: callable Middleware wrapper function that adds a trace id header to requests
from clients instantiated in supported Lambda runtime environments.[requestBuilder()](class-aws-middleware-method-requestbuilder.md)
: callable Builds an HTTP request for a command.[retry()](class-aws-middleware-method-retry.md)
: callable Middleware wrapper function that retries requests based on the boolean
result of invoking the provided "decider" function.[signer()](class-aws-middleware-method-signer.md)
: callable Creates a middleware that signs requests for a command.[sourceFile()](class-aws-middleware-method-sourcefile.md)
: callable Middleware used to allow a command parameter (e.g., "SourceFile") to
be used to specify the source of data for an upload operation.[tap()](class-aws-middleware-method-tap.md)
: callable Creates a middleware that invokes a callback at a given step.[timer()](class-aws-middleware-method-timer.md)
: mixed [validation()](class-aws-middleware-method-validation.md)
: callable Adds a middleware that uses client-side validation.

### Methods  [header link](class-aws-middleware-methods.md)

#### contentType()  [header link](class-aws-middleware-method-contenttype.md)

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

#### history()  [header link](class-aws-middleware-method-history.md)

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

#### invocationId()  [header link](class-aws-middleware-method-invocationid.md)

Middleware wrapper function that adds an invocation id header to
requests, which is only applied after the build step.

`
    public
            static        invocationId() : callable`

This is a uniquely generated UUID to identify initial and subsequent
retries as part of a complete request lifecycle.

##### Return values

callable

#### mapCommand()  [header link](class-aws-middleware-method-mapcommand.md)

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

#### mapRequest()  [header link](class-aws-middleware-method-maprequest.md)

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

#### mapResult()  [header link](class-aws-middleware-method-mapresult.md)

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

#### recursionDetection()  [header link](class-aws-middleware-method-recursiondetection.md)

Middleware wrapper function that adds a trace id header to requests
from clients instantiated in supported Lambda runtime environments.

`
    public
            static        recursionDetection() : callable`

The purpose for this header is to track and stop Lambda functions
from being recursively invoked due to misconfigured resources.

##### Return values

callable

#### requestBuilder()  [header link](class-aws-middleware-method-requestbuilder.md)

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

#### retry()  [header link](class-aws-middleware-method-retry.md)

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

#### signer()  [header link](class-aws-middleware-method-signer.md)

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

#### sourceFile()  [header link](class-aws-middleware-method-sourcefile.md)

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

#### tap()  [header link](class-aws-middleware-method-tap.md)

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

#### timer()  [header link](class-aws-middleware-method-timer.md)

`
    public
            static        timer() : mixed`

#### validation()  [header link](class-aws-middleware-method-validation.md)

Adds a middleware that uses client-side validation.

`
    public
            static        validation(Service $api[, Validator|null $validator = null ]) : callable`

##### Parameters

$api
: [Service](class-aws-api-service.md)

API being accessed.

$validator
: [Validator](class-aws-api-validator.md) \|null
= null

##### Return values

callable
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](class-aws-middleware-toc-methods.md)
- Methods
  - [contentType()](class-aws-middleware-method-contenttype.md)
  - [history()](class-aws-middleware-method-history.md)
  - [invocationId()](class-aws-middleware-method-invocationid.md)
  - [mapCommand()](class-aws-middleware-method-mapcommand.md)
  - [mapRequest()](class-aws-middleware-method-maprequest.md)
  - [mapResult()](class-aws-middleware-method-mapresult.md)
  - [recursionDetection()](class-aws-middleware-method-recursiondetection.md)
  - [requestBuilder()](class-aws-middleware-method-requestbuilder.md)
  - [retry()](class-aws-middleware-method-retry.md)
  - [signer()](class-aws-middleware-method-signer.md)
  - [sourceFile()](class-aws-middleware-method-sourcefile.md)
  - [tap()](class-aws-middleware-method-tap.md)
  - [timer()](class-aws-middleware-method-timer.md)
  - [validation()](class-aws-middleware-method-validation.md)

[Back To Top](class-aws-middleware-top.md)
