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

- [**2015-03-31**](api-lambda-2015-03-31.md)

  - [AddLayerVersionPermission](api-lambda-2015-03-31-addlayerversionpermission.md)
  - [AddPermission](api-lambda-2015-03-31-addpermission.md)
  - [CheckpointDurableExecution](api-lambda-2015-03-31-checkpointdurableexecution.md)
  - [CreateAlias](api-lambda-2015-03-31-createalias.md)
  - [CreateCapacityProvider](api-lambda-2015-03-31-createcapacityprovider.md)
  - [CreateCodeSigningConfig](api-lambda-2015-03-31-createcodesigningconfig.md)
  - [CreateEventSourceMapping](api-lambda-2015-03-31-createeventsourcemapping.md)
  - [CreateFunction](api-lambda-2015-03-31-createfunction.md)
  - [CreateFunctionUrlConfig](api-lambda-2015-03-31-createfunctionurlconfig.md)
  - [DeleteAlias](api-lambda-2015-03-31-deletealias.md)
  - [DeleteCapacityProvider](api-lambda-2015-03-31-deletecapacityprovider.md)
  - [DeleteCodeSigningConfig](api-lambda-2015-03-31-deletecodesigningconfig.md)
  - [DeleteEventSourceMapping](api-lambda-2015-03-31-deleteeventsourcemapping.md)
  - [DeleteFunction](api-lambda-2015-03-31-deletefunction.md)
  - [DeleteFunctionCodeSigningConfig](api-lambda-2015-03-31-deletefunctioncodesigningconfig.md)
  - [DeleteFunctionConcurrency](api-lambda-2015-03-31-deletefunctionconcurrency.md)
  - [DeleteFunctionEventInvokeConfig](api-lambda-2015-03-31-deletefunctioneventinvokeconfig.md)
  - [DeleteFunctionUrlConfig](api-lambda-2015-03-31-deletefunctionurlconfig.md)
  - [DeleteLayerVersion](api-lambda-2015-03-31-deletelayerversion.md)
  - [DeleteProvisionedConcurrencyConfig](api-lambda-2015-03-31-deleteprovisionedconcurrencyconfig.md)
  - [GetAccountSettings](api-lambda-2015-03-31-getaccountsettings.md)
  - [GetAlias](api-lambda-2015-03-31-getalias.md)
  - [GetCapacityProvider](api-lambda-2015-03-31-getcapacityprovider.md)
  - [GetCodeSigningConfig](api-lambda-2015-03-31-getcodesigningconfig.md)
  - [GetDurableExecution](api-lambda-2015-03-31-getdurableexecution.md)
  - [GetDurableExecutionHistory](api-lambda-2015-03-31-getdurableexecutionhistory.md)
  - [GetDurableExecutionState](api-lambda-2015-03-31-getdurableexecutionstate.md)
  - [GetEventSourceMapping](api-lambda-2015-03-31-geteventsourcemapping.md)
  - [GetFunction](api-lambda-2015-03-31-getfunction.md)
  - [GetFunctionCodeSigningConfig](api-lambda-2015-03-31-getfunctioncodesigningconfig.md)
  - [GetFunctionConcurrency](api-lambda-2015-03-31-getfunctionconcurrency.md)
  - [GetFunctionConfiguration](api-lambda-2015-03-31-getfunctionconfiguration.md)
  - [GetFunctionEventInvokeConfig](api-lambda-2015-03-31-getfunctioneventinvokeconfig.md)
  - [GetFunctionRecursionConfig](api-lambda-2015-03-31-getfunctionrecursionconfig.md)
  - [GetFunctionScalingConfig](api-lambda-2015-03-31-getfunctionscalingconfig.md)
  - [GetFunctionUrlConfig](api-lambda-2015-03-31-getfunctionurlconfig.md)
  - [GetLayerVersion](api-lambda-2015-03-31-getlayerversion.md)
  - [GetLayerVersionByArn](api-lambda-2015-03-31-getlayerversionbyarn.md)
  - [GetLayerVersionPolicy](api-lambda-2015-03-31-getlayerversionpolicy.md)
  - [GetPolicy](api-lambda-2015-03-31-getpolicy.md)
  - [GetProvisionedConcurrencyConfig](api-lambda-2015-03-31-getprovisionedconcurrencyconfig.md)
  - [GetRuntimeManagementConfig](api-lambda-2015-03-31-getruntimemanagementconfig.md)
  - [Invoke](api-lambda-2015-03-31-invoke.md)
  - [InvokeAsynchronous](api-lambda-2015-03-31-invokeasynchronous.md)
  - [InvokeWithResponseStream](api-lambda-2015-03-31-invokewithresponsestream.md)
  - [ListAliases](api-lambda-2015-03-31-listaliases.md)
  - [ListCapacityProviders](api-lambda-2015-03-31-listcapacityproviders.md)
  - [ListCodeSigningConfigs](api-lambda-2015-03-31-listcodesigningconfigs.md)
  - [ListDurableExecutionsByFunction](api-lambda-2015-03-31-listdurableexecutionsbyfunction.md)
  - [ListEventSourceMappings](api-lambda-2015-03-31-listeventsourcemappings.md)
  - [ListFunctionEventInvokeConfigs](api-lambda-2015-03-31-listfunctioneventinvokeconfigs.md)
  - [ListFunctionUrlConfigs](api-lambda-2015-03-31-listfunctionurlconfigs.md)
  - [ListFunctionVersionsByCapacityProvider](api-lambda-2015-03-31-listfunctionversionsbycapacityprovider.md)
  - [ListFunctions](api-lambda-2015-03-31-listfunctions.md)
  - [ListFunctionsByCodeSigningConfig](api-lambda-2015-03-31-listfunctionsbycodesigningconfig.md)
  - [ListLayerVersions](api-lambda-2015-03-31-listlayerversions.md)
  - [ListLayers](api-lambda-2015-03-31-listlayers.md)
  - [ListProvisionedConcurrencyConfigs](api-lambda-2015-03-31-listprovisionedconcurrencyconfigs.md)
  - [ListTags](api-lambda-2015-03-31-listtags.md)
  - [ListVersionsByFunction](api-lambda-2015-03-31-listversionsbyfunction.md)
  - [PublishLayerVersion](api-lambda-2015-03-31-publishlayerversion.md)
  - [PublishVersion](api-lambda-2015-03-31-publishversion.md)
  - [PutFunctionCodeSigningConfig](api-lambda-2015-03-31-putfunctioncodesigningconfig.md)
  - [PutFunctionConcurrency](api-lambda-2015-03-31-putfunctionconcurrency.md)
  - [PutFunctionEventInvokeConfig](api-lambda-2015-03-31-putfunctioneventinvokeconfig.md)
  - [PutFunctionRecursionConfig](api-lambda-2015-03-31-putfunctionrecursionconfig.md)
  - [PutFunctionScalingConfig](api-lambda-2015-03-31-putfunctionscalingconfig.md)
  - [PutProvisionedConcurrencyConfig](api-lambda-2015-03-31-putprovisionedconcurrencyconfig.md)
  - [PutRuntimeManagementConfig](api-lambda-2015-03-31-putruntimemanagementconfig.md)
  - [RemoveLayerVersionPermission](api-lambda-2015-03-31-removelayerversionpermission.md)
  - [RemovePermission](api-lambda-2015-03-31-removepermission.md)
  - [SendDurableExecutionCallbackFailure](api-lambda-2015-03-31-senddurableexecutioncallbackfailure.md)
  - [SendDurableExecutionCallbackHeartbeat](api-lambda-2015-03-31-senddurableexecutioncallbackheartbeat.md)
  - [SendDurableExecutionCallbackSuccess](api-lambda-2015-03-31-senddurableexecutioncallbacksuccess.md)
  - [StopDurableExecution](api-lambda-2015-03-31-stopdurableexecution.md)
  - [TagResource](api-lambda-2015-03-31-tagresource.md)
  - [UntagResource](api-lambda-2015-03-31-untagresource.md)
  - [UpdateAlias](api-lambda-2015-03-31-updatealias.md)
  - [UpdateCapacityProvider](api-lambda-2015-03-31-updatecapacityprovider.md)
  - [UpdateCodeSigningConfig](api-lambda-2015-03-31-updatecodesigningconfig.md)
  - [UpdateEventSourceMapping](api-lambda-2015-03-31-updateeventsourcemapping.md)
  - [UpdateFunctionCode](api-lambda-2015-03-31-updatefunctioncode.md)
  - [UpdateFunctionConfiguration](api-lambda-2015-03-31-updatefunctionconfiguration.md)
  - [UpdateFunctionEventInvokeConfig](api-lambda-2015-03-31-updatefunctioneventinvokeconfig.md)
  - [UpdateFunctionUrlConfig](api-lambda-2015-03-31-updatefunctionurlconfig.md)

## Examples

### Basics, Actions and Scenarios

The following code examples show you how to perform actions and implement common scenarios by using the AWS SDK for PHP with AWS Lambda.

- [See examples on AWS Docs](../../../sdk-for-php/v3/developer-guide/php-lambda-code-examples.md)

### Table of Contents  [header link](class-aws-lambda-lambdaclient-toc.md)

#### Methods  [header link](class-aws-lambda-lambdaclient-toc-methods.md)

[\_\_call()](class-aws-awsclienttrait.md#method___call)
: mixed [\_\_construct()](class-aws-lambda-lambdaclient-method-construct.md)
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
{@see \\Aws\\Credentials\\CredentialsInterface} object.[getDefaultCurlOptionsMiddleware()](class-aws-lambda-lambdaclient-method-getdefaultcurloptionsmiddleware.md)
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

### Methods  [header link](class-aws-lambda-lambdaclient-methods.md)

#### \_\_call()  [header link](class-aws-awsclienttrait.md\#method___call)

`
    public
                    __call(mixed $name, array<string|int, mixed> $args) : mixed`

##### Parameters

$name
: mixed$args
: array<string\|int, mixed>

#### \_\_construct()  [header link](class-aws-lambda-lambdaclient-method-construct.md)

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

#### getDefaultCurlOptionsMiddleware()  [header link](class-aws-lambda-lambdaclient-method-getdefaultcurloptionsmiddleware.md)

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
  - [Methods](class-aws-lambda-lambdaclient-toc-methods.md)
- Methods
  - [\_\_call()](class-aws-awsclienttrait.md#method___call)
  - [\_\_construct()](class-aws-lambda-lambdaclient-method-construct.md)
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
  - [getDefaultCurlOptionsMiddleware()](class-aws-lambda-lambdaclient-method-getdefaultcurloptionsmiddleware.md)
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

[Back To Top](class-aws-lambda-lambdaclient-top.md)
