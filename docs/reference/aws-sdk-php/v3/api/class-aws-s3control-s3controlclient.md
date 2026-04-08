Menu

- [Aws](namespace-aws.md)
- [S3Control](namespace-aws-s3control.md)

## S3ControlClient     extends [AwsClient](class-aws-awsclient.md)   in package    - [Aws](package-aws.md)

This client is used to interact with the **AWS S3 Control** service.

## Supported API Versions

This class uses a _service description model_ that is associated at
runtime based on the `version` option given when constructing the
client. The `version` option will determine which API operations,
waiters, and paginators are available for a client. Creating a command or a
specific API operation can be done using magic methods (e.g.,
`$client->commandName(/** parameters */)`, or using the
`$client->getCommand` method of the client.

- [**2018-08-20**](api-s3control-2018-08-20.md)

  - [AssociateAccessGrantsIdentityCenter](api-s3control-2018-08-20-associateaccessgrantsidentitycenter.md)
  - [CreateAccessGrant](api-s3control-2018-08-20-createaccessgrant.md)
  - [CreateAccessGrantsInstance](api-s3control-2018-08-20-createaccessgrantsinstance.md)
  - [CreateAccessGrantsLocation](api-s3control-2018-08-20-createaccessgrantslocation.md)
  - [CreateAccessPoint](api-s3control-2018-08-20-createaccesspoint.md)
  - [CreateAccessPointForObjectLambda](api-s3control-2018-08-20-createaccesspointforobjectlambda.md)
  - [CreateBucket](api-s3control-2018-08-20-createbucket.md)
  - [CreateJob](api-s3control-2018-08-20-createjob.md)
  - [CreateMultiRegionAccessPoint](api-s3control-2018-08-20-createmultiregionaccesspoint.md)
  - [CreateStorageLensGroup](api-s3control-2018-08-20-createstoragelensgroup.md)
  - [DeleteAccessGrant](api-s3control-2018-08-20-deleteaccessgrant.md)
  - [DeleteAccessGrantsInstance](api-s3control-2018-08-20-deleteaccessgrantsinstance.md)
  - [DeleteAccessGrantsInstanceResourcePolicy](api-s3control-2018-08-20-deleteaccessgrantsinstanceresourcepolicy.md)
  - [DeleteAccessGrantsLocation](api-s3control-2018-08-20-deleteaccessgrantslocation.md)
  - [DeleteAccessPoint](api-s3control-2018-08-20-deleteaccesspoint.md)
  - [DeleteAccessPointForObjectLambda](api-s3control-2018-08-20-deleteaccesspointforobjectlambda.md)
  - [DeleteAccessPointPolicy](api-s3control-2018-08-20-deleteaccesspointpolicy.md)
  - [DeleteAccessPointPolicyForObjectLambda](api-s3control-2018-08-20-deleteaccesspointpolicyforobjectlambda.md)
  - [DeleteAccessPointScope](api-s3control-2018-08-20-deleteaccesspointscope.md)
  - [DeleteBucket](api-s3control-2018-08-20-deletebucket.md)
  - [DeleteBucketLifecycleConfiguration](api-s3control-2018-08-20-deletebucketlifecycleconfiguration.md)
  - [DeleteBucketPolicy](api-s3control-2018-08-20-deletebucketpolicy.md)
  - [DeleteBucketReplication](api-s3control-2018-08-20-deletebucketreplication.md)
  - [DeleteBucketTagging](api-s3control-2018-08-20-deletebuckettagging.md)
  - [DeleteJobTagging](api-s3control-2018-08-20-deletejobtagging.md)
  - [DeleteMultiRegionAccessPoint](api-s3control-2018-08-20-deletemultiregionaccesspoint.md)
  - [DeletePublicAccessBlock](api-s3control-2018-08-20-deletepublicaccessblock.md)
  - [DeleteStorageLensConfiguration](api-s3control-2018-08-20-deletestoragelensconfiguration.md)
  - [DeleteStorageLensConfigurationTagging](api-s3control-2018-08-20-deletestoragelensconfigurationtagging.md)
  - [DeleteStorageLensGroup](api-s3control-2018-08-20-deletestoragelensgroup.md)
  - [DescribeJob](api-s3control-2018-08-20-describejob.md)
  - [DescribeMultiRegionAccessPointOperation](api-s3control-2018-08-20-describemultiregionaccesspointoperation.md)
  - [DissociateAccessGrantsIdentityCenter](api-s3control-2018-08-20-dissociateaccessgrantsidentitycenter.md)
  - [GetAccessGrant](api-s3control-2018-08-20-getaccessgrant.md)
  - [GetAccessGrantsInstance](api-s3control-2018-08-20-getaccessgrantsinstance.md)
  - [GetAccessGrantsInstanceForPrefix](api-s3control-2018-08-20-getaccessgrantsinstanceforprefix.md)
  - [GetAccessGrantsInstanceResourcePolicy](api-s3control-2018-08-20-getaccessgrantsinstanceresourcepolicy.md)
  - [GetAccessGrantsLocation](api-s3control-2018-08-20-getaccessgrantslocation.md)
  - [GetAccessPoint](api-s3control-2018-08-20-getaccesspoint.md)
  - [GetAccessPointConfigurationForObjectLambda](api-s3control-2018-08-20-getaccesspointconfigurationforobjectlambda.md)
  - [GetAccessPointForObjectLambda](api-s3control-2018-08-20-getaccesspointforobjectlambda.md)
  - [GetAccessPointPolicy](api-s3control-2018-08-20-getaccesspointpolicy.md)
  - [GetAccessPointPolicyForObjectLambda](api-s3control-2018-08-20-getaccesspointpolicyforobjectlambda.md)
  - [GetAccessPointPolicyStatus](api-s3control-2018-08-20-getaccesspointpolicystatus.md)
  - [GetAccessPointPolicyStatusForObjectLambda](api-s3control-2018-08-20-getaccesspointpolicystatusforobjectlambda.md)
  - [GetAccessPointScope](api-s3control-2018-08-20-getaccesspointscope.md)
  - [GetBucket](api-s3control-2018-08-20-getbucket.md)
  - [GetBucketLifecycleConfiguration](api-s3control-2018-08-20-getbucketlifecycleconfiguration.md)
  - [GetBucketPolicy](api-s3control-2018-08-20-getbucketpolicy.md)
  - [GetBucketReplication](api-s3control-2018-08-20-getbucketreplication.md)
  - [GetBucketTagging](api-s3control-2018-08-20-getbuckettagging.md)
  - [GetBucketVersioning](api-s3control-2018-08-20-getbucketversioning.md)
  - [GetDataAccess](api-s3control-2018-08-20-getdataaccess.md)
  - [GetJobTagging](api-s3control-2018-08-20-getjobtagging.md)
  - [GetMultiRegionAccessPoint](api-s3control-2018-08-20-getmultiregionaccesspoint.md)
  - [GetMultiRegionAccessPointPolicy](api-s3control-2018-08-20-getmultiregionaccesspointpolicy.md)
  - [GetMultiRegionAccessPointPolicyStatus](api-s3control-2018-08-20-getmultiregionaccesspointpolicystatus.md)
  - [GetMultiRegionAccessPointRoutes](api-s3control-2018-08-20-getmultiregionaccesspointroutes.md)
  - [GetPublicAccessBlock](api-s3control-2018-08-20-getpublicaccessblock.md)
  - [GetStorageLensConfiguration](api-s3control-2018-08-20-getstoragelensconfiguration.md)
  - [GetStorageLensConfigurationTagging](api-s3control-2018-08-20-getstoragelensconfigurationtagging.md)
  - [GetStorageLensGroup](api-s3control-2018-08-20-getstoragelensgroup.md)
  - [ListAccessGrants](api-s3control-2018-08-20-listaccessgrants.md)
  - [ListAccessGrantsInstances](api-s3control-2018-08-20-listaccessgrantsinstances.md)
  - [ListAccessGrantsLocations](api-s3control-2018-08-20-listaccessgrantslocations.md)
  - [ListAccessPoints](api-s3control-2018-08-20-listaccesspoints.md)
  - [ListAccessPointsForDirectoryBuckets](api-s3control-2018-08-20-listaccesspointsfordirectorybuckets.md)
  - [ListAccessPointsForObjectLambda](api-s3control-2018-08-20-listaccesspointsforobjectlambda.md)
  - [ListCallerAccessGrants](api-s3control-2018-08-20-listcalleraccessgrants.md)
  - [ListJobs](api-s3control-2018-08-20-listjobs.md)
  - [ListMultiRegionAccessPoints](api-s3control-2018-08-20-listmultiregionaccesspoints.md)
  - [ListRegionalBuckets](api-s3control-2018-08-20-listregionalbuckets.md)
  - [ListStorageLensConfigurations](api-s3control-2018-08-20-liststoragelensconfigurations.md)
  - [ListStorageLensGroups](api-s3control-2018-08-20-liststoragelensgroups.md)
  - [ListTagsForResource](api-s3control-2018-08-20-listtagsforresource.md)
  - [PutAccessGrantsInstanceResourcePolicy](api-s3control-2018-08-20-putaccessgrantsinstanceresourcepolicy.md)
  - [PutAccessPointConfigurationForObjectLambda](api-s3control-2018-08-20-putaccesspointconfigurationforobjectlambda.md)
  - [PutAccessPointPolicy](api-s3control-2018-08-20-putaccesspointpolicy.md)
  - [PutAccessPointPolicyForObjectLambda](api-s3control-2018-08-20-putaccesspointpolicyforobjectlambda.md)
  - [PutAccessPointScope](api-s3control-2018-08-20-putaccesspointscope.md)
  - [PutBucketLifecycleConfiguration](api-s3control-2018-08-20-putbucketlifecycleconfiguration.md)
  - [PutBucketPolicy](api-s3control-2018-08-20-putbucketpolicy.md)
  - [PutBucketReplication](api-s3control-2018-08-20-putbucketreplication.md)
  - [PutBucketTagging](api-s3control-2018-08-20-putbuckettagging.md)
  - [PutBucketVersioning](api-s3control-2018-08-20-putbucketversioning.md)
  - [PutJobTagging](api-s3control-2018-08-20-putjobtagging.md)
  - [PutMultiRegionAccessPointPolicy](api-s3control-2018-08-20-putmultiregionaccesspointpolicy.md)
  - [PutPublicAccessBlock](api-s3control-2018-08-20-putpublicaccessblock.md)
  - [PutStorageLensConfiguration](api-s3control-2018-08-20-putstoragelensconfiguration.md)
  - [PutStorageLensConfigurationTagging](api-s3control-2018-08-20-putstoragelensconfigurationtagging.md)
  - [SubmitMultiRegionAccessPointRoutes](api-s3control-2018-08-20-submitmultiregionaccesspointroutes.md)
  - [TagResource](api-s3control-2018-08-20-tagresource.md)
  - [UntagResource](api-s3control-2018-08-20-untagresource.md)
  - [UpdateAccessGrantsLocation](api-s3control-2018-08-20-updateaccessgrantslocation.md)
  - [UpdateJobPriority](api-s3control-2018-08-20-updatejobpriority.md)
  - [UpdateJobStatus](api-s3control-2018-08-20-updatejobstatus.md)
  - [UpdateStorageLensGroup](api-s3control-2018-08-20-updatestoragelensgroup.md)

### Table of Contents  [header link](class-aws-s3control-s3controlclient-toc.md)

#### Methods  [header link](class-aws-s3control-s3controlclient-toc-methods.md)

[\_\_call()](class-aws-awsclienttrait.md#method___call)
: mixed [\_\_construct()](class-aws-s3control-s3controlclient-method-construct.md)
: mixed The client constructor accepts the following options:[\_\_sleep()](class-aws-awsclient.md#method___sleep)
: mixed [\_apply\_use\_arn\_region()](class-aws-s3control-s3controlclient-method-apply-use-arn-region.md)
: mixed [execute()](class-aws-awsclienttrait.md#method_execute)
: mixed [executeAsync()](class-aws-awsclienttrait.md#method_executeAsync)
: mixed [factory()](class-aws-awsclient.md#method_factory)
: static [getApi()](class-aws-awsclienttrait.md#method_getApi)
: [Service](class-aws-api-service.md)[getArguments()](class-aws-s3control-s3controlclient-method-getarguments.md)
: array<string\|int, mixed> Get an array of client constructor arguments used by the client.[getClientBuiltIns()](class-aws-awsclient.md#method_getClientBuiltIns)
: array<string\|int, mixed> Provides the set of built-in keys and values
used for endpoint resolution[getClientContextParams()](class-aws-awsclient.md#method_getClientContextParams)
: array<string\|int, mixed> Provides the set of service context parameter
key-value pairs used for endpoint resolution.[getCommand()](class-aws-awsclienttrait.md#method_getCommand)
: [CommandInterface](class-aws-commandinterface.md)[getConfig()](class-aws-awsclient.md#method_getConfig)
: mixed\|null Get a client configuration value.[getCredentials()](class-aws-awsclient.md#method_getCredentials)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Returns a promise that is fulfilled with an
{@see \\Aws\\Credentials\\CredentialsInterface} object.[getEndpoint()](class-aws-awsclient.md#method_getEndpoint)
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

### Methods  [header link](class-aws-s3control-s3controlclient-methods.md)

#### \_\_call()  [header link](class-aws-awsclienttrait.md\#method___call)

`
    public
                    __call(mixed $name, array<string|int, mixed> $args) : mixed`

##### Parameters

$name
: mixed$args
: array<string\|int, mixed>

#### \_\_construct()  [header link](class-aws-s3control-s3controlclient-method-construct.md)

The client constructor accepts the following options:

`
    public
                    __construct(array<string|int, mixed> $args) : mixed`

In addition to the options available to
AwsClient::\_\_construct, S3ControlClient accepts the following
option:

- use\_dual\_stack\_endpoint: (bool) Set to true to send requests to an S3
Control Dual Stack endpoint by default, which enables IPv6 Protocol.
Can be enabled or disabled on individual operations by setting
'@use\_dual\_stack\_endpoint' to true or false. Note:
you cannot use it together with an accelerate endpoint.

##### Parameters

$args
: array<string\|int, mixed>

#### \_\_sleep()  [header link](class-aws-awsclient.md\#method___sleep)

`
    public
                    __sleep() : mixed`

#### \_apply\_use\_arn\_region()  [header link](class-aws-s3control-s3controlclient-method-apply-use-arn-region.md)

`
    public
            static        _apply_use_arn_region(mixed $value, array<string|int, mixed> &$args, HandlerList $list) : mixed`

##### Parameters

$value
: mixed$args
: array<string\|int, mixed>$list
: [HandlerList](class-aws-handlerlist.md)

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

#### getArguments()  [header link](class-aws-s3control-s3controlclient-method-getarguments.md)

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
  - [Methods](class-aws-s3control-s3controlclient-toc-methods.md)
- Methods
  - [\_\_call()](class-aws-awsclienttrait.md#method___call)
  - [\_\_construct()](class-aws-s3control-s3controlclient-method-construct.md)
  - [\_\_sleep()](class-aws-awsclient.md#method___sleep)
  - [\_apply\_use\_arn\_region()](class-aws-s3control-s3controlclient-method-apply-use-arn-region.md)
  - [execute()](class-aws-awsclienttrait.md#method_execute)
  - [executeAsync()](class-aws-awsclienttrait.md#method_executeAsync)
  - [factory()](class-aws-awsclient.md#method_factory)
  - [getApi()](class-aws-awsclienttrait.md#method_getApi)
  - [getArguments()](class-aws-s3control-s3controlclient-method-getarguments.md)
  - [getClientBuiltIns()](class-aws-awsclient.md#method_getClientBuiltIns)
  - [getClientContextParams()](class-aws-awsclient.md#method_getClientContextParams)
  - [getCommand()](class-aws-awsclienttrait.md#method_getCommand)
  - [getConfig()](class-aws-awsclient.md#method_getConfig)
  - [getCredentials()](class-aws-awsclient.md#method_getCredentials)
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

[Back To Top](class-aws-s3control-s3controlclient-top.md)
