Menu

- [Aws](namespace-aws.md)

## Waiter        in package    - [Aws](package-aws.md)       implements  [PromisorInterface](class-guzzlehttp-promise-promisorinterface.md)

"Waiters" are associated with an AWS resource (e.g., EC2 instance), and poll
that resource and until it is in a particular state.

The Waiter object produces a promise that is either a.) resolved once the
waiting conditions are met, or b.) rejected if the waiting conditions cannot
be met or has exceeded the number of allowed attempts at meeting the
conditions. You can use waiters in a blocking or non-blocking way, depending
on whether you call wait() on the promise.

The configuration for the waiter must include information about the operation
and the conditions for wait completion.

### Table of Contents  [header link](class-aws-waiter-toc.md)

#### Interfaces  [header link](class-aws-waiter-toc-interfaces.md)

[PromisorInterface](class-guzzlehttp-promise-promisorinterface.md)Interface used with classes that return a promise.

#### Methods  [header link](class-aws-waiter-toc-methods.md)

[\_\_construct()](class-aws-waiter-method-construct.md)
: mixed The array of configuration options include:[promise()](class-aws-waiter-method-promise.md)
: [Coroutine](class-guzzlehttp-promise-coroutine.md)Returns a promise.

### Methods  [header link](class-aws-waiter-methods.md)

#### \_\_construct()  [header link](class-aws-waiter-method-construct.md)

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

##### Tags  [header link](class-aws-waiter-method-construct-tags.md)

throwsInvalidArgumentException

if the configuration is incomplete.

#### promise()  [header link](class-aws-waiter-method-promise.md)

Returns a promise.

`
    public
                    promise() : Coroutine`

##### Return values

[Coroutine](class-guzzlehttp-promise-coroutine.md)
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](class-aws-waiter-toc-methods.md)
- Methods
  - [\_\_construct()](class-aws-waiter-method-construct.md)
  - [promise()](class-aws-waiter-method-promise.md)

[Back To Top](class-aws-waiter-top.md)
