Menu

- [Aws](namespace-aws.md)

## AwsClientTrait

A trait providing generic functionality for interacting with Amazon Web
Services. This is meant to be used in classes implementing
\\Aws\\AwsClientInterface

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientTrait.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientTrait.html\#toc-methods)

[\_\_call()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientTrait.html#method___call)
: mixed [execute()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientTrait.html#method_execute)
: mixed [executeAsync()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientTrait.html#method_executeAsync)
: mixed [getApi()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientTrait.html#method_getApi)
: [Service](class-aws-api-service.md)[getCommand()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientTrait.html#method_getCommand)
: [CommandInterface](class-aws-commandinterface.md)[getIterator()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientTrait.html#method_getIterator)
: mixed [getPaginator()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientTrait.html#method_getPaginator)
: mixed [getWaiter()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientTrait.html#method_getWaiter)
: mixed [waitUntil()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientTrait.html#method_waitUntil)
: mixed

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientTrait.html\#methods)

#### \_\_call()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientTrait.html\#method___call)

`
    public
                    __call(mixed $name, array<string|int, mixed> $args) : mixed`

##### Parameters

$name
: mixed$args
: array<string\|int, mixed>

#### execute()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientTrait.html\#method_execute)

`
    public
                    execute(CommandInterface $command) : mixed`

##### Parameters

$command
: [CommandInterface](class-aws-commandinterface.md)

#### executeAsync()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientTrait.html\#method_executeAsync)

`
    public
                    executeAsync(CommandInterface $command) : mixed`

##### Parameters

$command
: [CommandInterface](class-aws-commandinterface.md)

#### getApi()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientTrait.html\#method_getApi)

`
    public
    abstract                getApi() : Service`

##### Return values

[Service](class-aws-api-service.md)

#### getCommand()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientTrait.html\#method_getCommand)

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

#### getIterator()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientTrait.html\#method_getIterator)

`
    public
                    getIterator(mixed $name[, array<string|int, mixed> $args = [] ]) : mixed`

##### Parameters

$name
: mixed$args
: array<string\|int, mixed>
= \[\]

#### getPaginator()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientTrait.html\#method_getPaginator)

`
    public
                    getPaginator(mixed $name[, array<string|int, mixed> $args = [] ]) : mixed`

##### Parameters

$name
: mixed$args
: array<string\|int, mixed>
= \[\]

#### getWaiter()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientTrait.html\#method_getWaiter)

`
    public
                    getWaiter(mixed $name[, array<string|int, mixed> $args = [] ]) : mixed`

##### Parameters

$name
: mixed$args
: array<string\|int, mixed>
= \[\]

#### waitUntil()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientTrait.html\#method_waitUntil)

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
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientTrait.html#toc-methods)
- Methods
  - [\_\_call()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientTrait.html#method___call)
  - [execute()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientTrait.html#method_execute)
  - [executeAsync()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientTrait.html#method_executeAsync)
  - [getApi()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientTrait.html#method_getApi)
  - [getCommand()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientTrait.html#method_getCommand)
  - [getIterator()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientTrait.html#method_getIterator)
  - [getPaginator()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientTrait.html#method_getPaginator)
  - [getWaiter()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientTrait.html#method_getWaiter)
  - [waitUntil()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientTrait.html#method_waitUntil)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientTrait.html#top)
