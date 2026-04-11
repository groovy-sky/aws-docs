Menu

- [Aws](namespace-aws.md)

## AwsClientInterface     in    - [Aws](package-aws.md)

Represents an AWS client.

### Table of Contents  [header link](class-aws-awsclientinterface-toc.md)

#### Methods  [header link](class-aws-awsclientinterface-toc-methods.md)

[\_\_call()](class-aws-awsclientinterface-method-call.md)
: [ResultInterface](class-aws-resultinterface.md)Creates and executes a command for an operation by name.[execute()](class-aws-awsclientinterface-method-execute.md)
: [ResultInterface](class-aws-resultinterface.md)Execute a single command.[executeAsync()](class-aws-awsclientinterface-method-executeasync.md)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Execute a command asynchronously.[getApi()](class-aws-awsclientinterface-method-getapi.md)
: [Service](class-aws-api-service.md)Get the service description associated with the client.[getCommand()](class-aws-awsclientinterface-method-getcommand.md)
: [CommandInterface](class-aws-commandinterface.md)Create a command for an operation name.[getConfig()](class-aws-awsclientinterface-method-getconfig.md)
: mixed\|null Get a client configuration value.[getCredentials()](class-aws-awsclientinterface-method-getcredentials.md)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Returns a promise that is fulfilled with an
{@see \\Aws\\Credentials\\CredentialsInterface} object.[getEndpoint()](class-aws-awsclientinterface-method-getendpoint.md)
: [UriInterface](class-psr-http-message-uriinterface.md)Gets the default endpoint, or base URL, used by the client.[getHandlerList()](class-aws-awsclientinterface-method-gethandlerlist.md)
: [HandlerList](class-aws-handlerlist.md)Get the handler list used to transfer commands.[getIterator()](class-aws-awsclientinterface-method-getiterator.md)
: IteratorGet a resource iterator for the specified operation.[getPaginator()](class-aws-awsclientinterface-method-getpaginator.md)
: [ResultPaginator](class-aws-resultpaginator.md)Get a result paginator for the specified operation.[getRegion()](class-aws-awsclientinterface-method-getregion.md)
: string Get the region to which the client is configured to send requests.[getWaiter()](class-aws-awsclientinterface-method-getwaiter.md)
: [Waiter](class-aws-waiter.md)Get a waiter that waits until a resource is in a particular state.[waitUntil()](class-aws-awsclientinterface-method-waituntil.md)
: void Wait until a resource is in a particular state.

### Methods  [header link](class-aws-awsclientinterface-methods.md)

#### \_\_call()  [header link](class-aws-awsclientinterface-method-call.md)

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

##### Tags  [header link](class-aws-awsclientinterface-method-call-tags.md)

throwsException

##### Return values

[ResultInterface](class-aws-resultinterface.md)

#### execute()  [header link](class-aws-awsclientinterface-method-execute.md)

Execute a single command.

`
    public
                    execute(CommandInterface $command) : ResultInterface`

##### Parameters

$command
: [CommandInterface](class-aws-commandinterface.md)

Command to execute

##### Tags  [header link](class-aws-awsclientinterface-method-execute-tags.md)

throwsException

##### Return values

[ResultInterface](class-aws-resultinterface.md)

#### executeAsync()  [header link](class-aws-awsclientinterface-method-executeasync.md)

Execute a command asynchronously.

`
    public
                    executeAsync(CommandInterface $command) : PromiseInterface`

##### Parameters

$command
: [CommandInterface](class-aws-commandinterface.md)

Command to execute

##### Return values

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)

#### getApi()  [header link](class-aws-awsclientinterface-method-getapi.md)

Get the service description associated with the client.

`
    public
                    getApi() : Service`

##### Return values

[Service](class-aws-api-service.md)

#### getCommand()  [header link](class-aws-awsclientinterface-method-getcommand.md)

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

##### Tags  [header link](class-aws-awsclientinterface-method-getcommand-tags.md)

throwsInvalidArgumentException

if no command can be found by name

##### Return values

[CommandInterface](class-aws-commandinterface.md)

#### getConfig()  [header link](class-aws-awsclientinterface-method-getconfig.md)

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

#### getCredentials()  [header link](class-aws-awsclientinterface-method-getcredentials.md)

Returns a promise that is fulfilled with an
{@see \\Aws\\Credentials\\CredentialsInterface} object.

`
    public
                    getCredentials() : PromiseInterface`

If you need the credentials synchronously, then call the wait() method
on the returned promise.

##### Return values

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)

#### getEndpoint()  [header link](class-aws-awsclientinterface-method-getendpoint.md)

Gets the default endpoint, or base URL, used by the client.

`
    public
                    getEndpoint() : UriInterface`

##### Return values

[UriInterface](class-psr-http-message-uriinterface.md)

#### getHandlerList()  [header link](class-aws-awsclientinterface-method-gethandlerlist.md)

Get the handler list used to transfer commands.

`
    public
                    getHandlerList() : HandlerList`

This list can be modified to add middleware or to change the underlying
handler used to send HTTP requests.

##### Return values

[HandlerList](class-aws-handlerlist.md)

#### getIterator()  [header link](class-aws-awsclientinterface-method-getiterator.md)

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

##### Tags  [header link](class-aws-awsclientinterface-method-getiterator-tags.md)

throwsUnexpectedValueException

if the iterator config is invalid.

##### Return values

Iterator

#### getPaginator()  [header link](class-aws-awsclientinterface-method-getpaginator.md)

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

##### Tags  [header link](class-aws-awsclientinterface-method-getpaginator-tags.md)

throwsUnexpectedValueException

if the iterator config is invalid.

##### Return values

[ResultPaginator](class-aws-resultpaginator.md)

#### getRegion()  [header link](class-aws-awsclientinterface-method-getregion.md)

Get the region to which the client is configured to send requests.

`
    public
                    getRegion() : string`

##### Return values

string

#### getWaiter()  [header link](class-aws-awsclientinterface-method-getwaiter.md)

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

##### Tags  [header link](class-aws-awsclientinterface-method-getwaiter-tags.md)

throwsUnexpectedValueException

if the waiter is invalid.

##### Return values

[Waiter](class-aws-waiter.md)

#### waitUntil()  [header link](class-aws-awsclientinterface-method-waituntil.md)

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

##### Tags  [header link](class-aws-awsclientinterface-method-waituntil-tags.md)

throwsUnexpectedValueException

if the waiter is invalid.

<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](class-aws-awsclientinterface-toc-constants.md)
  - [Methods](class-aws-awsclientinterface-toc-methods.md)
- Methods
  - [\_\_call()](class-aws-awsclientinterface-method-call.md)
  - [execute()](class-aws-awsclientinterface-method-execute.md)
  - [executeAsync()](class-aws-awsclientinterface-method-executeasync.md)
  - [getApi()](class-aws-awsclientinterface-method-getapi.md)
  - [getCommand()](class-aws-awsclientinterface-method-getcommand.md)
  - [getConfig()](class-aws-awsclientinterface-method-getconfig.md)
  - [getCredentials()](class-aws-awsclientinterface-method-getcredentials.md)
  - [getEndpoint()](class-aws-awsclientinterface-method-getendpoint.md)
  - [getHandlerList()](class-aws-awsclientinterface-method-gethandlerlist.md)
  - [getIterator()](class-aws-awsclientinterface-method-getiterator.md)
  - [getPaginator()](class-aws-awsclientinterface-method-getpaginator.md)
  - [getRegion()](class-aws-awsclientinterface-method-getregion.md)
  - [getWaiter()](class-aws-awsclientinterface-method-getwaiter.md)
  - [waitUntil()](class-aws-awsclientinterface-method-waituntil.md)

[Back To Top](class-aws-awsclientinterface-top.md)
