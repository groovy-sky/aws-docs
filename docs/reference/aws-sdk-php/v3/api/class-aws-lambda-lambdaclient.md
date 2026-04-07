Menu

- [Aws](namespace-aws.md)
- [Lambda](namespace-aws-lambda.md)

## LambdaClient     extends [AwsClient](class-aws-awsclient.md)   in package    - [Aws](package-aws.md)

This client is used to interact with AWS Lambda

## Supported API Versions

This class uses a _service description model_ that is associated at
runtime based on the `version` option given when constructing the
client. The `version` option will determine which API operations,
waiters, and paginators are available for a client. Creating a command or a
specific API operation can be done using magic methods (e.g.,
`$client->commandName(/** parameters */)`, or using the
`$client->getCommand` method of the client.

- [**2015-03-31**](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html)

  - [AddLayerVersionPermission](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#addlayerversionpermission)
  - [AddPermission](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#addpermission)
  - [CheckpointDurableExecution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#checkpointdurableexecution)
  - [CreateAlias](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#createalias)
  - [CreateCapacityProvider](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#createcapacityprovider)
  - [CreateCodeSigningConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#createcodesigningconfig)
  - [CreateEventSourceMapping](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#createeventsourcemapping)
  - [CreateFunction](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#createfunction)
  - [CreateFunctionUrlConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#createfunctionurlconfig)
  - [DeleteAlias](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#deletealias)
  - [DeleteCapacityProvider](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#deletecapacityprovider)
  - [DeleteCodeSigningConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#deletecodesigningconfig)
  - [DeleteEventSourceMapping](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#deleteeventsourcemapping)
  - [DeleteFunction](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#deletefunction)
  - [DeleteFunctionCodeSigningConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#deletefunctioncodesigningconfig)
  - [DeleteFunctionConcurrency](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#deletefunctionconcurrency)
  - [DeleteFunctionEventInvokeConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#deletefunctioneventinvokeconfig)
  - [DeleteFunctionUrlConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#deletefunctionurlconfig)
  - [DeleteLayerVersion](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#deletelayerversion)
  - [DeleteProvisionedConcurrencyConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#deleteprovisionedconcurrencyconfig)
  - [GetAccountSettings](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#getaccountsettings)
  - [GetAlias](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#getalias)
  - [GetCapacityProvider](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#getcapacityprovider)
  - [GetCodeSigningConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#getcodesigningconfig)
  - [GetDurableExecution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#getdurableexecution)
  - [GetDurableExecutionHistory](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#getdurableexecutionhistory)
  - [GetDurableExecutionState](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#getdurableexecutionstate)
  - [GetEventSourceMapping](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#geteventsourcemapping)
  - [GetFunction](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#getfunction)
  - [GetFunctionCodeSigningConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#getfunctioncodesigningconfig)
  - [GetFunctionConcurrency](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#getfunctionconcurrency)
  - [GetFunctionConfiguration](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#getfunctionconfiguration)
  - [GetFunctionEventInvokeConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#getfunctioneventinvokeconfig)
  - [GetFunctionRecursionConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#getfunctionrecursionconfig)
  - [GetFunctionScalingConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#getfunctionscalingconfig)
  - [GetFunctionUrlConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#getfunctionurlconfig)
  - [GetLayerVersion](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#getlayerversion)
  - [GetLayerVersionByArn](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#getlayerversionbyarn)
  - [GetLayerVersionPolicy](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#getlayerversionpolicy)
  - [GetPolicy](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#getpolicy)
  - [GetProvisionedConcurrencyConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#getprovisionedconcurrencyconfig)
  - [GetRuntimeManagementConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#getruntimemanagementconfig)
  - [Invoke](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#invoke)
  - [InvokeAsynchronous](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#invokeasynchronous)
  - [InvokeWithResponseStream](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#invokewithresponsestream)
  - [ListAliases](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#listaliases)
  - [ListCapacityProviders](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#listcapacityproviders)
  - [ListCodeSigningConfigs](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#listcodesigningconfigs)
  - [ListDurableExecutionsByFunction](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#listdurableexecutionsbyfunction)
  - [ListEventSourceMappings](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#listeventsourcemappings)
  - [ListFunctionEventInvokeConfigs](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#listfunctioneventinvokeconfigs)
  - [ListFunctionUrlConfigs](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#listfunctionurlconfigs)
  - [ListFunctionVersionsByCapacityProvider](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#listfunctionversionsbycapacityprovider)
  - [ListFunctions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#listfunctions)
  - [ListFunctionsByCodeSigningConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#listfunctionsbycodesigningconfig)
  - [ListLayerVersions](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#listlayerversions)
  - [ListLayers](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#listlayers)
  - [ListProvisionedConcurrencyConfigs](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#listprovisionedconcurrencyconfigs)
  - [ListTags](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#listtags)
  - [ListVersionsByFunction](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#listversionsbyfunction)
  - [PublishLayerVersion](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#publishlayerversion)
  - [PublishVersion](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#publishversion)
  - [PutFunctionCodeSigningConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#putfunctioncodesigningconfig)
  - [PutFunctionConcurrency](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#putfunctionconcurrency)
  - [PutFunctionEventInvokeConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#putfunctioneventinvokeconfig)
  - [PutFunctionRecursionConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#putfunctionrecursionconfig)
  - [PutFunctionScalingConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#putfunctionscalingconfig)
  - [PutProvisionedConcurrencyConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#putprovisionedconcurrencyconfig)
  - [PutRuntimeManagementConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#putruntimemanagementconfig)
  - [RemoveLayerVersionPermission](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#removelayerversionpermission)
  - [RemovePermission](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#removepermission)
  - [SendDurableExecutionCallbackFailure](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#senddurableexecutioncallbackfailure)
  - [SendDurableExecutionCallbackHeartbeat](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#senddurableexecutioncallbackheartbeat)
  - [SendDurableExecutionCallbackSuccess](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#senddurableexecutioncallbacksuccess)
  - [StopDurableExecution](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#stopdurableexecution)
  - [TagResource](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#tagresource)
  - [UntagResource](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#untagresource)
  - [UpdateAlias](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#updatealias)
  - [UpdateCapacityProvider](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#updatecapacityprovider)
  - [UpdateCodeSigningConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#updatecodesigningconfig)
  - [UpdateEventSourceMapping](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#updateeventsourcemapping)
  - [UpdateFunctionCode](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#updatefunctioncode)
  - [UpdateFunctionConfiguration](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#updatefunctionconfiguration)
  - [UpdateFunctionEventInvokeConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#updatefunctioneventinvokeconfig)
  - [UpdateFunctionUrlConfig](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-lambda-2015-03-31.html#updatefunctionurlconfig)

## Examples

### Basics, Actions and Scenarios

The following code examples show you how to perform actions and implement common scenarios by using the AWS SDK for PHP with AWS Lambda.

- [See examples on AWS Docs](https://docs.aws.amazon.com/sdk-for-php/v3/developer-guide/php_lambda_code_examples.html)

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Lambda.LambdaClient.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Lambda.LambdaClient.html\#toc-methods)

[\_\_call()](class-aws-awsclienttrait.md#method___call)
: mixed [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Lambda.LambdaClient.html#method___construct)
: mixed The client constructor accepts the following options:[\_\_sleep()](class-aws-awsclient.md#method___sleep)
: mixed [execute()](class-aws-awsclienttrait.md#method_execute)
: mixed [executeAsync()](class-aws-awsclienttrait.md#method_executeAsync)
: mixed [factory()](class-aws-awsclient.md#method_factory)
: static [getApi()](class-aws-awsclienttrait.md#method_getApi)
: [Service](class-aws-api-service.md)[getArguments()](class-aws-awsclient.md#method_getArguments)
: array<string\|int, mixed> Get an array of client constructor arguments used by the client.[getClientBuiltIns()](class-aws-awsclient.md#method_getClientBuiltIns)
: array<string\|int, mixed> Provides the set of built-in keys and values
used for endpoint resolution[getClientContextParams()](class-aws-awsclient.md#method_getClientContextParams)
: array<string\|int, mixed> Provides the set of service context parameter
key-value pairs used for endpoint resolution.[getCommand()](class-aws-awsclienttrait.md#method_getCommand)
: [CommandInterface](class-aws-commandinterface.md)[getConfig()](class-aws-awsclient.md#method_getConfig)
: mixed\|null Get a client configuration value.[getCredentials()](class-aws-awsclient.md#method_getCredentials)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Returns a promise that is fulfilled with an
{@see \\Aws\\Credentials\\CredentialsInterface} object.[getDefaultCurlOptionsMiddleware()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Lambda.LambdaClient.html#method_getDefaultCurlOptionsMiddleware)
: callable Provides a middleware that sets default Curl options for the command[getEndpoint()](class-aws-awsclient.md#method_getEndpoint)
: [UriInterface](class-psr-http-message-uriinterface.md)Gets the default endpoint, or base URL, used by the client.[getEndpointProvider()](class-aws-awsclient.md#method_getEndpointProvider)
: mixed [getEndpointProviderArgs()](class-aws-awsclient.md#method_getEndpointProviderArgs)
: array<string\|int, mixed> Retrieves arguments to be used in endpoint resolution.[getHandlerList()](class-aws-awsclient.md#method_getHandlerList)
: [HandlerList](class-aws-handlerlist.md)Get the handler list used to transfer commands.[getIterator()](class-aws-awsclienttrait.md#method_getIterator)
: mixed [getPaginator()](class-aws-awsclienttrait.md#method_getPaginator)
: mixed [getRegion()](class-aws-awsclient.md#method_getRegion)
: string Get the region to which the client is configured to send requests.[getSignatureProvider()](class-aws-awsclient.md#method_getSignatureProvider)
: callable Get the signature\_provider function of the client.[getToken()](class-aws-awsclient.md#method_getToken)
: mixed [getWaiter()](class-aws-awsclienttrait.md#method_getWaiter)
: mixed [waitUntil()](class-aws-awsclienttrait.md#method_waitUntil)
: mixed

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Lambda.LambdaClient.html\#methods)

#### \_\_call()  [header link](class-aws-awsclienttrait.md\#method___call)

`
    public
                    __call(mixed $name, array<string|int, mixed> $args) : mixed`

##### Parameters

$name
: mixed$args
: array<string\|int, mixed>

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Lambda.LambdaClient.html\#method___construct)

The client constructor accepts the following options:

`
    public
                    __construct(array<string|int, mixed> $args) : mixed`

##### Parameters

$args
: array<string\|int, mixed>

Client configuration arguments.

#### \_\_sleep()  [header link](class-aws-awsclient.md\#method___sleep)

`
    public
                    __sleep() : mixed`

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

#### factory()  [header link](class-aws-awsclient.md\#method_factory)

`
    public
            static        factory([array<string|int, mixed> $config = [] ]) : static`

##### Parameters

$config
: array<string\|int, mixed>
= \[\]

##### Tags  [header link](class-aws-awsclient.md\#method_factory\#tags)

deprecated

##### Return values

static

#### getApi()  [header link](class-aws-awsclienttrait.md\#method_getApi)

`
    public
    abstract                getApi() : Service`

##### Return values

[Service](class-aws-api-service.md)

#### getArguments()  [header link](class-aws-awsclient.md\#method_getArguments)

Get an array of client constructor arguments used by the client.

`
    public
            static        getArguments() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getClientBuiltIns()  [header link](class-aws-awsclient.md\#method_getClientBuiltIns)

Provides the set of built-in keys and values
used for endpoint resolution

`
    public
                    getClientBuiltIns() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getClientContextParams()  [header link](class-aws-awsclient.md\#method_getClientContextParams)

Provides the set of service context parameter
key-value pairs used for endpoint resolution.

`
    public
                    getClientContextParams() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getCommand()  [header link](class-aws-awsclienttrait.md\#method_getCommand)

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

#### getConfig()  [header link](class-aws-awsclient.md\#method_getConfig)

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

#### getCredentials()  [header link](class-aws-awsclient.md\#method_getCredentials)

Returns a promise that is fulfilled with an
{@see \\Aws\\Credentials\\CredentialsInterface} object.

`
    public
                    getCredentials() : PromiseInterface`

If you need the credentials synchronously, then call the wait() method
on the returned promise.

##### Return values

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)

#### getDefaultCurlOptionsMiddleware()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Lambda.LambdaClient.html\#method_getDefaultCurlOptionsMiddleware)

Provides a middleware that sets default Curl options for the command

`
    public
                    getDefaultCurlOptionsMiddleware() : callable`

##### Return values

callable

#### getEndpoint()  [header link](class-aws-awsclient.md\#method_getEndpoint)

Gets the default endpoint, or base URL, used by the client.

`
    public
                    getEndpoint() : UriInterface`

##### Return values

[UriInterface](class-psr-http-message-uriinterface.md)

#### getEndpointProvider()  [header link](class-aws-awsclient.md\#method_getEndpointProvider)

`
    public
                    getEndpointProvider() : mixed`

#### getEndpointProviderArgs()  [header link](class-aws-awsclient.md\#method_getEndpointProviderArgs)

Retrieves arguments to be used in endpoint resolution.

`
    public
                    getEndpointProviderArgs() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getHandlerList()  [header link](class-aws-awsclient.md\#method_getHandlerList)

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

#### getRegion()  [header link](class-aws-awsclient.md\#method_getRegion)

Get the region to which the client is configured to send requests.

`
    public
                    getRegion() : string`

##### Return values

string

#### getSignatureProvider()  [header link](class-aws-awsclient.md\#method_getSignatureProvider)

Get the signature\_provider function of the client.

`
    public
        final            getSignatureProvider() : callable`

##### Return values

callable

#### getToken()  [header link](class-aws-awsclient.md\#method_getToken)

`
    public
                    getToken() : mixed`

#### getWaiter()  [header link](class-aws-awsclienttrait.md\#method_getWaiter)

`
    public
                    getWaiter(mixed $name[, array<string|int, mixed> $args = [] ]) : mixed`

##### Parameters

$name
: mixed$args
: array<string\|int, mixed>
= \[\]

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
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Lambda.LambdaClient.html#toc-methods)
- Methods
  - [\_\_call()](class-aws-awsclienttrait.md#method___call)
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Lambda.LambdaClient.html#method___construct)
  - [\_\_sleep()](class-aws-awsclient.md#method___sleep)
  - [execute()](class-aws-awsclienttrait.md#method_execute)
  - [executeAsync()](class-aws-awsclienttrait.md#method_executeAsync)
  - [factory()](class-aws-awsclient.md#method_factory)
  - [getApi()](class-aws-awsclienttrait.md#method_getApi)
  - [getArguments()](class-aws-awsclient.md#method_getArguments)
  - [getClientBuiltIns()](class-aws-awsclient.md#method_getClientBuiltIns)
  - [getClientContextParams()](class-aws-awsclient.md#method_getClientContextParams)
  - [getCommand()](class-aws-awsclienttrait.md#method_getCommand)
  - [getConfig()](class-aws-awsclient.md#method_getConfig)
  - [getCredentials()](class-aws-awsclient.md#method_getCredentials)
  - [getDefaultCurlOptionsMiddleware()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Lambda.LambdaClient.html#method_getDefaultCurlOptionsMiddleware)
  - [getEndpoint()](class-aws-awsclient.md#method_getEndpoint)
  - [getEndpointProvider()](class-aws-awsclient.md#method_getEndpointProvider)
  - [getEndpointProviderArgs()](class-aws-awsclient.md#method_getEndpointProviderArgs)
  - [getHandlerList()](class-aws-awsclient.md#method_getHandlerList)
  - [getIterator()](class-aws-awsclienttrait.md#method_getIterator)
  - [getPaginator()](class-aws-awsclienttrait.md#method_getPaginator)
  - [getRegion()](class-aws-awsclient.md#method_getRegion)
  - [getSignatureProvider()](class-aws-awsclient.md#method_getSignatureProvider)
  - [getToken()](class-aws-awsclient.md#method_getToken)
  - [getWaiter()](class-aws-awsclienttrait.md#method_getWaiter)
  - [waitUntil()](class-aws-awsclienttrait.md#method_waitUntil)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Lambda.LambdaClient.html#top)
