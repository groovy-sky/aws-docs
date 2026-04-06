Menu

- [Aws](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.html)

## AwsClientInterface     in    - [Aws](https://docs.aws.amazon.com/aws-sdk-php/v3/api/package-Aws.html)

Represents an AWS client.

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientInterface.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientInterface.html\#toc-methods)

[\_\_call()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientInterface.html#method___call)
: [ResultInterface](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ResultInterface.html)Creates and executes a command for an operation by name.[execute()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientInterface.html#method_execute)
: [ResultInterface](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ResultInterface.html)Execute a single command.[executeAsync()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientInterface.html#method_executeAsync)
: [PromiseInterface](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.PromiseInterface.html)Execute a command asynchronously.[getApi()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientInterface.html#method_getApi)
: [Service](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html)Get the service description associated with the client.[getCommand()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientInterface.html#method_getCommand)
: [CommandInterface](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CommandInterface.html)Create a command for an operation name.[getConfig()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientInterface.html#method_getConfig)
: mixed\|null Get a client configuration value.[getCredentials()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientInterface.html#method_getCredentials)
: [PromiseInterface](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.PromiseInterface.html)Returns a promise that is fulfilled with an
{@see \\Aws\\Credentials\\CredentialsInterface} object.[getEndpoint()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientInterface.html#method_getEndpoint)
: [UriInterface](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.UriInterface.html)Gets the default endpoint, or base URL, used by the client.[getHandlerList()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientInterface.html#method_getHandlerList)
: [HandlerList](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html)Get the handler list used to transfer commands.[getIterator()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientInterface.html#method_getIterator)
: IteratorGet a resource iterator for the specified operation.[getPaginator()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientInterface.html#method_getPaginator)
: [ResultPaginator](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ResultPaginator.html)Get a result paginator for the specified operation.[getRegion()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientInterface.html#method_getRegion)
: string Get the region to which the client is configured to send requests.[getWaiter()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientInterface.html#method_getWaiter)
: [Waiter](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Waiter.html)Get a waiter that waits until a resource is in a particular state.[waitUntil()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientInterface.html#method_waitUntil)
: void Wait until a resource is in a particular state.

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientInterface.html\#methods)

#### \_\_call()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientInterface.html\#method___call)

Creates and executes a command for an operation by name.

`
    public
                    __call(string $name, array<string|int, mixed> $arguments) : ResultInterface`

Suffixing an operation name with "Async" will return a
promise that can be used to execute commands asynchronously.

##### Parameters

$name
: string

Name of the command to execute.

$arguments
: array<string\|int, mixed>

Arguments to pass to the getCommand method.

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientInterface.html\#method___call\#tags)

throwsException

##### Return values

[ResultInterface](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ResultInterface.html)

#### execute()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientInterface.html\#method_execute)

Execute a single command.

`
    public
                    execute(CommandInterface $command) : ResultInterface`

##### Parameters

$command
: [CommandInterface](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CommandInterface.html)

Command to execute

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientInterface.html\#method_execute\#tags)

throwsException

##### Return values

[ResultInterface](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ResultInterface.html)

#### executeAsync()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientInterface.html\#method_executeAsync)

Execute a command asynchronously.

`
    public
                    executeAsync(CommandInterface $command) : PromiseInterface`

##### Parameters

$command
: [CommandInterface](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CommandInterface.html)

Command to execute

##### Return values

[PromiseInterface](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.PromiseInterface.html)

#### getApi()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientInterface.html\#method_getApi)

Get the service description associated with the client.

`
    public
                    getApi() : Service`

##### Return values

[Service](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html)

#### getCommand()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientInterface.html\#method_getCommand)

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

##### Parameters

$name
: string

Name of the operation to use in the command

$args
: array<string\|int, mixed>
= \[\]

Arguments to pass to the command

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientInterface.html\#method_getCommand\#tags)

throwsInvalidArgumentException

if no command can be found by name

##### Return values

[CommandInterface](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.CommandInterface.html)

#### getConfig()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientInterface.html\#method_getConfig)

Get a client configuration value.

`
    public
                    getConfig([string|null $option = null ]) : mixed|null`

##### Parameters

$option
: string\|null
= null

The option to retrieve. Pass null to retrieve
all options.

##### Return values

mixed\|null

#### getCredentials()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientInterface.html\#method_getCredentials)

Returns a promise that is fulfilled with an
{@see \\Aws\\Credentials\\CredentialsInterface} object.

`
    public
                    getCredentials() : PromiseInterface`

If you need the credentials synchronously, then call the wait() method
on the returned promise.

##### Return values

[PromiseInterface](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.PromiseInterface.html)

#### getEndpoint()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientInterface.html\#method_getEndpoint)

Gets the default endpoint, or base URL, used by the client.

`
    public
                    getEndpoint() : UriInterface`

##### Return values

[UriInterface](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.UriInterface.html)

#### getHandlerList()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientInterface.html\#method_getHandlerList)

Get the handler list used to transfer commands.

`
    public
                    getHandlerList() : HandlerList`

This list can be modified to add middleware or to change the underlying
handler used to send HTTP requests.

##### Return values

[HandlerList](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html)

#### getIterator()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientInterface.html\#method_getIterator)

Get a resource iterator for the specified operation.

`
    public
                    getIterator(string $name[, array<string|int, mixed> $args = [] ]) : Iterator`

##### Parameters

$name
: string

Name of the iterator to retrieve.

$args
: array<string\|int, mixed>
= \[\]

Command arguments to use with each command.

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientInterface.html\#method_getIterator\#tags)

throwsUnexpectedValueException

if the iterator config is invalid.

##### Return values

Iterator

#### getPaginator()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientInterface.html\#method_getPaginator)

Get a result paginator for the specified operation.

`
    public
                    getPaginator(string $name[, array<string|int, mixed> $args = [] ]) : ResultPaginator`

##### Parameters

$name
: string

Name of the operation used for iterator

$args
: array<string\|int, mixed>
= \[\]

Command args to be used with each command

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientInterface.html\#method_getPaginator\#tags)

throwsUnexpectedValueException

if the iterator config is invalid.

##### Return values

[ResultPaginator](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.ResultPaginator.html)

#### getRegion()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientInterface.html\#method_getRegion)

Get the region to which the client is configured to send requests.

`
    public
                    getRegion() : string`

##### Return values

string

#### getWaiter()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientInterface.html\#method_getWaiter)

Get a waiter that waits until a resource is in a particular state.

`
    public
                    getWaiter(string|callable $name[, array<string|int, mixed> $args = [] ]) : Waiter`

Retrieving a waiter can be useful when you wish to wait asynchronously:

$waiter = $client->getWaiter('foo', \['bar' => 'baz'\]);
$waiter->promise()->then(function () { echo 'Done!'; });

##### Parameters

$name
: string\|callable

Name of the waiter that defines the wait
configuration and conditions.

$args
: array<string\|int, mixed>
= \[\]

Args to be used with each command executed
by the waiter. Waiter configuration options
can be provided in an associative array in
the @waiter key.

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientInterface.html\#method_getWaiter\#tags)

throwsUnexpectedValueException

if the waiter is invalid.

##### Return values

[Waiter](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Waiter.html)

#### waitUntil()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientInterface.html\#method_waitUntil)

Wait until a resource is in a particular state.

`
    public
                    waitUntil(string|callable $name[, array<string|int, mixed> $args = [] ]) : void`

##### Parameters

$name
: string\|callable

Name of the waiter that defines the wait
configuration and conditions.

$args
: array<string\|int, mixed>
= \[\]

Args to be used with each command executed
by the waiter. Waiter configuration options
can be provided in an associative array in
the @waiter key.

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientInterface.html\#method_waitUntil\#tags)

throwsUnexpectedValueException

if the waiter is invalid.

<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientInterface.html#toc-constants)
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientInterface.html#toc-methods)
- Methods
  - [\_\_call()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientInterface.html#method___call)
  - [execute()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientInterface.html#method_execute)
  - [executeAsync()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientInterface.html#method_executeAsync)
  - [getApi()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientInterface.html#method_getApi)
  - [getCommand()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientInterface.html#method_getCommand)
  - [getConfig()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientInterface.html#method_getConfig)
  - [getCredentials()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientInterface.html#method_getCredentials)
  - [getEndpoint()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientInterface.html#method_getEndpoint)
  - [getHandlerList()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientInterface.html#method_getHandlerList)
  - [getIterator()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientInterface.html#method_getIterator)
  - [getPaginator()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientInterface.html#method_getPaginator)
  - [getRegion()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientInterface.html#method_getRegion)
  - [getWaiter()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientInterface.html#method_getWaiter)
  - [waitUntil()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientInterface.html#method_waitUntil)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientInterface.html#top)
