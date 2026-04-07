Menu

- [Aws](namespace-aws.md)

## Waiter        in package    - [Aws](package-aws.md)       implements  [PromisorInterface](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.PromisorInterface.html)

"Waiters" are associated with an AWS resource (e.g., EC2 instance), and poll
that resource and until it is in a particular state.

The Waiter object produces a promise that is either a.) resolved once the
waiting conditions are met, or b.) rejected if the waiting conditions cannot
be met or has exceeded the number of allowed attempts at meeting the
conditions. You can use waiters in a blocking or non-blocking way, depending
on whether you call wait() on the promise.

The configuration for the waiter must include information about the operation
and the conditions for wait completion.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Waiter.html\#toc)

#### Interfaces  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Waiter.html\#toc-interfaces)

[PromisorInterface](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.PromisorInterface.html)Interface used with classes that return a promise.

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Waiter.html\#toc-methods)

[\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Waiter.html#method___construct)
: mixed The array of configuration options include:[promise()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Waiter.html#method_promise)
: [Coroutine](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.Coroutine.html)Returns a promise.

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Waiter.html\#methods)

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Waiter.html\#method___construct)

The array of configuration options include:

`
    public
                    __construct(AwsClientInterface $client, string $name[, array<string|int, mixed> $args = [] ][, array<string|int, mixed> $config = [] ]) : mixed`

- acceptors: (array) Array of acceptor options
- delay: (int) Number of seconds to delay between attempts
- maxAttempts: (int) Maximum number of attempts before failing
- operation: (string) Name of the API operation to use for polling
- before: (callable) Invoked before attempts. Accepts command and tries.

##### Parameters

$client
: [AwsClientInterface](class-aws-awsclientinterface.md)

Client used to execute commands.

$name
: string

Waiter name.

$args
: array<string\|int, mixed>
= \[\]

Command arguments.

$config
: array<string\|int, mixed>
= \[\]

Waiter config that overrides defaults.

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Waiter.html\#method___construct\#tags)

throwsInvalidArgumentException

if the configuration is incomplete.

#### promise()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Waiter.html\#method_promise)

Returns a promise.

`
    public
                    promise() : Coroutine`

##### Return values

[Coroutine](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.Coroutine.html)
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Waiter.html#toc-methods)
- Methods
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Waiter.html#method___construct)
  - [promise()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Waiter.html#method_promise)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Waiter.html#top)
