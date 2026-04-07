Menu

- [Aws](namespace-aws.md)

## TraceMiddleware        in package    - [Aws](package-aws.md)

Traces state changes between middlewares.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.TraceMiddleware.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.TraceMiddleware.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.TraceMiddleware.html#method___construct)
: mixed Configuration array can contain the following key value pairs.[\_\_invoke()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.TraceMiddleware.html#method___invoke)
: mixed

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.TraceMiddleware.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.TraceMiddleware.html\#method___construct)

Configuration array can contain the following key value pairs.

`
    public
                    __construct([array<string|int, mixed> $config = [] ][, Service|null $service = null ]) : mixed`

- logfn: (callable) Function that is invoked with log messages. By
default, PHP's "echo" function will be utilized.
- stream\_size: (int) When the size of a stream is greater than this
number, the stream data will not be logged. Set to "0" to not log any
stream data.
- scrub\_auth: (bool) Set to false to disable the scrubbing of auth data
from the logged messages.
- http: (bool) Set to false to disable the "debug" feature of lower
level HTTP adapters (e.g., verbose curl output).
- auth\_strings: (array) A mapping of authentication string regular
expressions to scrubbed strings. These mappings are passed directly to
preg\_replace (e.g., preg\_replace($key, $value, $debugOutput) if
"scrub\_auth" is set to true.
- auth\_headers: (array) A mapping of header names known to contain
sensitive data to what the scrubbed value should be. The value of any
headers contained in this array will be replaced with the if
"scrub\_auth" is set to true.

##### Parameters

$config
: array<string\|int, mixed>
= \[\]$service
: [Service](class-aws-api-service.md) \|null
= null

#### \_\_invoke()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.TraceMiddleware.html\#method___invoke)

`
    public
                    __invoke(mixed $step, mixed $name) : mixed`

##### Parameters

$step
: mixed$name
: mixed
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.TraceMiddleware.html#toc-methods)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.TraceMiddleware.html#method___construct)
  - [\_\_invoke()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.TraceMiddleware.html#method___invoke)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.TraceMiddleware.html#top)
