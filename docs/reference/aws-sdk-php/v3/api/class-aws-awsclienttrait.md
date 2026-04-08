Menu

- [Aws](namespace-aws.md)

## AwsClientTrait

A trait providing generic functionality for interacting with Amazon Web
Services. This is meant to be used in classes implementing
\\Aws\\AwsClientInterface

### Table of Contents  [header link](class-aws-awsclienttrait-toc.md)

#### Methods  [header link](class-aws-awsclienttrait-toc-methods.md)

[\_\_call()](class-aws-awsclienttrait-method-call.md)
: mixed [execute()](class-aws-awsclienttrait-method-execute.md)
: mixed [executeAsync()](class-aws-awsclienttrait-method-executeasync.md)
: mixed [getApi()](class-aws-awsclienttrait-method-getapi.md)
: [Service](class-aws-api-service.md)[getCommand()](class-aws-awsclienttrait-method-getcommand.md)
: [CommandInterface](class-aws-commandinterface.md)[getIterator()](class-aws-awsclienttrait-method-getiterator.md)
: mixed [getPaginator()](class-aws-awsclienttrait-method-getpaginator.md)
: mixed [getWaiter()](class-aws-awsclienttrait-method-getwaiter.md)
: mixed [waitUntil()](class-aws-awsclienttrait-method-waituntil.md)
: mixed

### Methods  [header link](class-aws-awsclienttrait-methods.md)

#### \_\_call()  [header link](class-aws-awsclienttrait-method-call.md)

`
    public
                    __call(mixed $name, array<string|int, mixed> $args) : mixed`

##### Parameters

$name
: mixed$args
: array<string\|int, mixed>

#### execute()  [header link](class-aws-awsclienttrait-method-execute.md)

`
    public
                    execute(CommandInterface $command) : mixed`

##### Parameters

$command
: [CommandInterface](class-aws-commandinterface.md)

#### executeAsync()  [header link](class-aws-awsclienttrait-method-executeasync.md)

`
    public
                    executeAsync(CommandInterface $command) : mixed`

##### Parameters

$command
: [CommandInterface](class-aws-commandinterface.md)

#### getApi()  [header link](class-aws-awsclienttrait-method-getapi.md)

`
    public
    abstract                getApi() : Service`

##### Return values

[Service](class-aws-api-service.md)

#### getCommand()  [header link](class-aws-awsclienttrait-method-getcommand.md)

`
    public
    abstract                getCommand(string $name[, array<string|int, mixed> $args = [] ]) : CommandInterface`

##### Parameters

$name
: string$args
: array<string\|int, mixed>
= \[\]

##### Return values

[CommandInterface](class-aws-commandinterface.md)

#### getIterator()  [header link](class-aws-awsclienttrait-method-getiterator.md)

`
    public
                    getIterator(mixed $name[, array<string|int, mixed> $args = [] ]) : mixed`

##### Parameters

$name
: mixed$args
: array<string\|int, mixed>
= \[\]

#### getPaginator()  [header link](class-aws-awsclienttrait-method-getpaginator.md)

`
    public
                    getPaginator(mixed $name[, array<string|int, mixed> $args = [] ]) : mixed`

##### Parameters

$name
: mixed$args
: array<string\|int, mixed>
= \[\]

#### getWaiter()  [header link](class-aws-awsclienttrait-method-getwaiter.md)

`
    public
                    getWaiter(mixed $name[, array<string|int, mixed> $args = [] ]) : mixed`

##### Parameters

$name
: mixed$args
: array<string\|int, mixed>
= \[\]

#### waitUntil()  [header link](class-aws-awsclienttrait-method-waituntil.md)

`
    public
                    waitUntil(mixed $name[, array<string|int, mixed> $args = [] ]) : mixed`

##### Parameters

$name
: mixed$args
: array<string\|int, mixed>
= \[\]
<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Methods](class-aws-awsclienttrait-toc-methods.md)
- Methods
  - [\_\_call()](class-aws-awsclienttrait-method-call.md)
  - [execute()](class-aws-awsclienttrait-method-execute.md)
  - [executeAsync()](class-aws-awsclienttrait-method-executeasync.md)
  - [getApi()](class-aws-awsclienttrait-method-getapi.md)
  - [getCommand()](class-aws-awsclienttrait-method-getcommand.md)
  - [getIterator()](class-aws-awsclienttrait-method-getiterator.md)
  - [getPaginator()](class-aws-awsclienttrait-method-getpaginator.md)
  - [getWaiter()](class-aws-awsclienttrait-method-getwaiter.md)
  - [waitUntil()](class-aws-awsclienttrait-method-waituntil.md)

[Back To Top](class-aws-awsclienttrait-top.md)
