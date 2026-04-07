Menu

- [Aws](namespace-aws.md)

## MultiRegionClient        in package    - [Aws](package-aws.md)       implements  [AwsClientInterface](class-aws-awsclientinterface.md)  Uses  [AwsClientTrait](class-aws-awsclienttrait.md)

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MultiRegionClient.html\#toc)

#### Interfaces  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MultiRegionClient.html\#toc-interfaces)

[AwsClientInterface](class-aws-awsclientinterface.md)Represents an AWS client.

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MultiRegionClient.html\#toc-methods)

[\_\_call()](class-aws-awsclienttrait.md#method___call)
: mixed [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MultiRegionClient.html#method___construct)
: mixed The multi-region client constructor accepts the following options:[execute()](class-aws-awsclienttrait.md#method_execute)
: mixed [executeAsync()](class-aws-awsclienttrait.md#method_executeAsync)
: mixed [getApi()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MultiRegionClient.html#method_getApi)
: [Service](class-aws-api-service.md)Get the service description associated with the client.[getArguments()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MultiRegionClient.html#method_getArguments)
: mixed [getCommand()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MultiRegionClient.html#method_getCommand)
: [CommandInterface](class-aws-commandinterface.md)Create a command for an operation name.[getConfig()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MultiRegionClient.html#method_getConfig)
: mixed\|null Get a client configuration value.[getCredentials()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MultiRegionClient.html#method_getCredentials)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Returns a promise that is fulfilled with an
{@see \\Aws\\Credentials\\CredentialsInterface} object.[getEndpoint()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MultiRegionClient.html#method_getEndpoint)
: [UriInterface](class-psr-http-message-uriinterface.md)Gets the default endpoint, or base URL, used by the client.[getHandlerList()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MultiRegionClient.html#method_getHandlerList)
: [HandlerList](class-aws-handlerlist.md)Get the handler list used to transfer commands.[getIterator()](class-aws-awsclienttrait.md#method_getIterator)
: mixed [getPaginator()](class-aws-awsclienttrait.md#method_getPaginator)
: mixed [getRegion()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MultiRegionClient.html#method_getRegion)
: string Get the region to which the client is configured to send requests by
default.[getWaiter()](class-aws-awsclienttrait.md#method_getWaiter)
: mixed [useCustomHandler()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MultiRegionClient.html#method_useCustomHandler)
: mixed [waitUntil()](class-aws-awsclienttrait.md#method_waitUntil)
: mixed

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MultiRegionClient.html\#methods)

#### \_\_call()  [header link](class-aws-awsclienttrait.md\#method___call)

`
    public
                    __call(mixed $name, array<string|int, mixed> $args) : mixed`

##### Parameters

$name
: mixed$args
: array<string\|int, mixed>

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MultiRegionClient.html\#method___construct)

The multi-region client constructor accepts the following options:

`
    public
                    __construct([array<string|int, mixed> $args = [] ]) : mixed`

- client\_factory: (callable) An optional callable that takes an array of
client configuration arguments and returns a regionalized client.
- partition: (Aws\\Endpoint\\Partition\|string) AWS partition to connect to.
Valid partitions include "aws," "aws-cn," and "aws-us-gov." Used to
restrict the scope of the mapRegions method.
- region: (string) Region to connect to when no override is provided.
Used to create the default client factory and determine the appropriate
AWS partition when present.

##### Parameters

$args
: array<string\|int, mixed>
= \[\]

Client configuration arguments.

#### execute()  [header link](class-aws-awsclienttrait.md\#method_execute)

`
    public
                    execute(CommandInterface $command) : mixed`

##### Parameters

$command
: [CommandInterface](class-aws-commandinterface.md)

#### executeAsync()  [header link](class-aws-awsclienttrait.md\#method_executeAsync)

`
    public
                    executeAsync(CommandInterface $command) : mixed`

##### Parameters

$command
: [CommandInterface](class-aws-commandinterface.md)

#### getApi()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MultiRegionClient.html\#method_getApi)

Get the service description associated with the client.

`
    public
                    getApi() : Service`

##### Return values

[Service](class-aws-api-service.md)

#### getArguments()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MultiRegionClient.html\#method_getArguments)

`
    public
            static        getArguments() : mixed`

#### getCommand()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MultiRegionClient.html\#method_getCommand)

Create a command for an operation name.

`
    public
                    getCommand(string $name[, array<string|int, mixed> $args = [] ]) : CommandInterface`

Special keys may be set on the command to control how it behaves,
including:

- @http: Associative array of transfer specific options to apply to the
request that is serialized for this command. Available keys include
"proxy", "verify", "timeout", "connect\_timeout", "debug", "delay", and
"headers".
- @region: The region to which the command should be sent.

##### Parameters

$name
: string

Name of the operation to use in the command

$args
: array<string\|int, mixed>
= \[\]

Arguments to pass to the command

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MultiRegionClient.html\#method_getCommand\#tags)

throwsInvalidArgumentException

if no command can be found by name

##### Return values

[CommandInterface](class-aws-commandinterface.md)

#### getConfig()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MultiRegionClient.html\#method_getConfig)

Get a client configuration value.

`
    public
                    getConfig([mixed $option = null ]) : mixed|null`

##### Parameters

$option
: mixed
= null

The option to retrieve. Pass null to retrieve
all options.

##### Return values

mixed\|null

#### getCredentials()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MultiRegionClient.html\#method_getCredentials)

Returns a promise that is fulfilled with an
{@see \\Aws\\Credentials\\CredentialsInterface} object.

`
    public
                    getCredentials() : PromiseInterface`

If you need the credentials synchronously, then call the wait() method
on the returned promise.

##### Return values

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)

#### getEndpoint()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MultiRegionClient.html\#method_getEndpoint)

Gets the default endpoint, or base URL, used by the client.

`
    public
                    getEndpoint() : UriInterface`

##### Return values

[UriInterface](class-psr-http-message-uriinterface.md)

#### getHandlerList()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MultiRegionClient.html\#method_getHandlerList)

Get the handler list used to transfer commands.

`
    public
                    getHandlerList() : HandlerList`

This list can be modified to add middleware or to change the underlying
handler used to send HTTP requests.

##### Return values

[HandlerList](class-aws-handlerlist.md)

#### getIterator()  [header link](class-aws-awsclienttrait.md\#method_getIterator)

`
    public
                    getIterator(mixed $name[, array<string|int, mixed> $args = [] ]) : mixed`

##### Parameters

$name
: mixed$args
: array<string\|int, mixed>
= \[\]

#### getPaginator()  [header link](class-aws-awsclienttrait.md\#method_getPaginator)

`
    public
                    getPaginator(mixed $name[, array<string|int, mixed> $args = [] ]) : mixed`

##### Parameters

$name
: mixed$args
: array<string\|int, mixed>
= \[\]

#### getRegion()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MultiRegionClient.html\#method_getRegion)

Get the region to which the client is configured to send requests by
default.

`
    public
                    getRegion() : string`

##### Return values

string

#### getWaiter()  [header link](class-aws-awsclienttrait.md\#method_getWaiter)

`
    public
                    getWaiter(mixed $name[, array<string|int, mixed> $args = [] ]) : mixed`

##### Parameters

$name
: mixed$args
: array<string\|int, mixed>
= \[\]

#### useCustomHandler()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MultiRegionClient.html\#method_useCustomHandler)

`
    public
                    useCustomHandler(callable $handler) : mixed`

##### Parameters

$handler
: callable

#### waitUntil()  [header link](class-aws-awsclienttrait.md\#method_waitUntil)

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
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MultiRegionClient.html#toc-methods)
- Methods
  - [\_\_call()](class-aws-awsclienttrait.md#method___call)
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MultiRegionClient.html#method___construct)
  - [execute()](class-aws-awsclienttrait.md#method_execute)
  - [executeAsync()](class-aws-awsclienttrait.md#method_executeAsync)
  - [getApi()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MultiRegionClient.html#method_getApi)
  - [getArguments()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MultiRegionClient.html#method_getArguments)
  - [getCommand()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MultiRegionClient.html#method_getCommand)
  - [getConfig()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MultiRegionClient.html#method_getConfig)
  - [getCredentials()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MultiRegionClient.html#method_getCredentials)
  - [getEndpoint()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MultiRegionClient.html#method_getEndpoint)
  - [getHandlerList()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MultiRegionClient.html#method_getHandlerList)
  - [getIterator()](class-aws-awsclienttrait.md#method_getIterator)
  - [getPaginator()](class-aws-awsclienttrait.md#method_getPaginator)
  - [getRegion()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MultiRegionClient.html#method_getRegion)
  - [getWaiter()](class-aws-awsclienttrait.md#method_getWaiter)
  - [useCustomHandler()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MultiRegionClient.html#method_useCustomHandler)
  - [waitUntil()](class-aws-awsclienttrait.md#method_waitUntil)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.MultiRegionClient.html#top)
