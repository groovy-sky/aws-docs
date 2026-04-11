Menu

- [Aws](namespace-aws.md)

## MultiRegionClient        in package    - [Aws](package-aws.md)       implements  [AwsClientInterface](class-aws-awsclientinterface.md)  Uses  [AwsClientTrait](class-aws-awsclienttrait.md)

### Table of Contents  [header link](class-aws-multiregionclient-toc.md)

#### Interfaces  [header link](class-aws-multiregionclient-toc-interfaces.md)

[AwsClientInterface](class-aws-awsclientinterface.md)Represents an AWS client.

#### Methods  [header link](class-aws-multiregionclient-toc-methods.md)

[\_\_call()](class-aws-awsclienttrait.md#method___call)
: mixed [\_\_construct()](class-aws-multiregionclient-method-construct.md)
: mixed The multi-region client constructor accepts the following options:[execute()](class-aws-awsclienttrait.md#method_execute)
: mixed [executeAsync()](class-aws-awsclienttrait.md#method_executeAsync)
: mixed [getApi()](class-aws-multiregionclient-method-getapi.md)
: [Service](class-aws-api-service.md)Get the service description associated with the client.[getArguments()](class-aws-multiregionclient-method-getarguments.md)
: mixed [getCommand()](class-aws-multiregionclient-method-getcommand.md)
: [CommandInterface](class-aws-commandinterface.md)Create a command for an operation name.[getConfig()](class-aws-multiregionclient-method-getconfig.md)
: mixed\|null Get a client configuration value.[getCredentials()](class-aws-multiregionclient-method-getcredentials.md)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Returns a promise that is fulfilled with an
{@see \\Aws\\Credentials\\CredentialsInterface} object.[getEndpoint()](class-aws-multiregionclient-method-getendpoint.md)
: [UriInterface](class-psr-http-message-uriinterface.md)Gets the default endpoint, or base URL, used by the client.[getHandlerList()](class-aws-multiregionclient-method-gethandlerlist.md)
: [HandlerList](class-aws-handlerlist.md)Get the handler list used to transfer commands.[getIterator()](class-aws-awsclienttrait.md#method_getIterator)
: mixed [getPaginator()](class-aws-awsclienttrait.md#method_getPaginator)
: mixed [getRegion()](class-aws-multiregionclient-method-getregion.md)
: string Get the region to which the client is configured to send requests by
default.[getWaiter()](class-aws-awsclienttrait.md#method_getWaiter)
: mixed [useCustomHandler()](class-aws-multiregionclient-method-usecustomhandler.md)
: mixed [waitUntil()](class-aws-awsclienttrait.md#method_waitUntil)
: mixed

### Methods  [header link](class-aws-multiregionclient-methods.md)

#### \_\_call()  [header link](class-aws-awsclienttrait.md\#method___call)

`
    public
                    __call(mixed $name, array<string|int, mixed> $args) : mixed`

##### Parameters

$name
: mixed$args
: array<string\|int, mixed>

#### \_\_construct()  [header link](class-aws-multiregionclient-method-construct.md)

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

#### getApi()  [header link](class-aws-multiregionclient-method-getapi.md)

Get the service description associated with the client.

`
    public
                    getApi() : Service`

##### Return values

[Service](class-aws-api-service.md)

#### getArguments()  [header link](class-aws-multiregionclient-method-getarguments.md)

`
    public
            static        getArguments() : mixed`

#### getCommand()  [header link](class-aws-multiregionclient-method-getcommand.md)

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

##### Tags  [header link](class-aws-multiregionclient-method-getcommand-tags.md)

throwsInvalidArgumentException

if no command can be found by name

##### Return values

[CommandInterface](class-aws-commandinterface.md)

#### getConfig()  [header link](class-aws-multiregionclient-method-getconfig.md)

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

#### getCredentials()  [header link](class-aws-multiregionclient-method-getcredentials.md)

Returns a promise that is fulfilled with an
{@see \\Aws\\Credentials\\CredentialsInterface} object.

`
    public
                    getCredentials() : PromiseInterface`

If you need the credentials synchronously, then call the wait() method
on the returned promise.

##### Return values

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)

#### getEndpoint()  [header link](class-aws-multiregionclient-method-getendpoint.md)

Gets the default endpoint, or base URL, used by the client.

`
    public
                    getEndpoint() : UriInterface`

##### Return values

[UriInterface](class-psr-http-message-uriinterface.md)

#### getHandlerList()  [header link](class-aws-multiregionclient-method-gethandlerlist.md)

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

#### getRegion()  [header link](class-aws-multiregionclient-method-getregion.md)

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

#### useCustomHandler()  [header link](class-aws-multiregionclient-method-usecustomhandler.md)

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
  - [Methods](class-aws-multiregionclient-toc-methods.md)
- Methods
  - [\_\_call()](class-aws-awsclienttrait.md#method___call)
  - [\_\_construct()](class-aws-multiregionclient-method-construct.md)
  - [execute()](class-aws-awsclienttrait.md#method_execute)
  - [executeAsync()](class-aws-awsclienttrait.md#method_executeAsync)
  - [getApi()](class-aws-multiregionclient-method-getapi.md)
  - [getArguments()](class-aws-multiregionclient-method-getarguments.md)
  - [getCommand()](class-aws-multiregionclient-method-getcommand.md)
  - [getConfig()](class-aws-multiregionclient-method-getconfig.md)
  - [getCredentials()](class-aws-multiregionclient-method-getcredentials.md)
  - [getEndpoint()](class-aws-multiregionclient-method-getendpoint.md)
  - [getHandlerList()](class-aws-multiregionclient-method-gethandlerlist.md)
  - [getIterator()](class-aws-awsclienttrait.md#method_getIterator)
  - [getPaginator()](class-aws-awsclienttrait.md#method_getPaginator)
  - [getRegion()](class-aws-multiregionclient-method-getregion.md)
  - [getWaiter()](class-aws-awsclienttrait.md#method_getWaiter)
  - [useCustomHandler()](class-aws-multiregionclient-method-usecustomhandler.md)
  - [waitUntil()](class-aws-awsclienttrait.md#method_waitUntil)

[Back To Top](class-aws-multiregionclient-top.md)
