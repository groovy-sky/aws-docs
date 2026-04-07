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

- [**2018-08-20**](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html)

  - [AssociateAccessGrantsIdentityCenter](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#associateaccessgrantsidentitycenter)
  - [CreateAccessGrant](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#createaccessgrant)
  - [CreateAccessGrantsInstance](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#createaccessgrantsinstance)
  - [CreateAccessGrantsLocation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#createaccessgrantslocation)
  - [CreateAccessPoint](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#createaccesspoint)
  - [CreateAccessPointForObjectLambda](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#createaccesspointforobjectlambda)
  - [CreateBucket](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#createbucket)
  - [CreateJob](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#createjob)
  - [CreateMultiRegionAccessPoint](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#createmultiregionaccesspoint)
  - [CreateStorageLensGroup](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#createstoragelensgroup)
  - [DeleteAccessGrant](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#deleteaccessgrant)
  - [DeleteAccessGrantsInstance](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#deleteaccessgrantsinstance)
  - [DeleteAccessGrantsInstanceResourcePolicy](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#deleteaccessgrantsinstanceresourcepolicy)
  - [DeleteAccessGrantsLocation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#deleteaccessgrantslocation)
  - [DeleteAccessPoint](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#deleteaccesspoint)
  - [DeleteAccessPointForObjectLambda](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#deleteaccesspointforobjectlambda)
  - [DeleteAccessPointPolicy](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#deleteaccesspointpolicy)
  - [DeleteAccessPointPolicyForObjectLambda](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#deleteaccesspointpolicyforobjectlambda)
  - [DeleteAccessPointScope](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#deleteaccesspointscope)
  - [DeleteBucket](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#deletebucket)
  - [DeleteBucketLifecycleConfiguration](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#deletebucketlifecycleconfiguration)
  - [DeleteBucketPolicy](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#deletebucketpolicy)
  - [DeleteBucketReplication](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#deletebucketreplication)
  - [DeleteBucketTagging](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#deletebuckettagging)
  - [DeleteJobTagging](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#deletejobtagging)
  - [DeleteMultiRegionAccessPoint](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#deletemultiregionaccesspoint)
  - [DeletePublicAccessBlock](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#deletepublicaccessblock)
  - [DeleteStorageLensConfiguration](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#deletestoragelensconfiguration)
  - [DeleteStorageLensConfigurationTagging](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#deletestoragelensconfigurationtagging)
  - [DeleteStorageLensGroup](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#deletestoragelensgroup)
  - [DescribeJob](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#describejob)
  - [DescribeMultiRegionAccessPointOperation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#describemultiregionaccesspointoperation)
  - [DissociateAccessGrantsIdentityCenter](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#dissociateaccessgrantsidentitycenter)
  - [GetAccessGrant](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#getaccessgrant)
  - [GetAccessGrantsInstance](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#getaccessgrantsinstance)
  - [GetAccessGrantsInstanceForPrefix](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#getaccessgrantsinstanceforprefix)
  - [GetAccessGrantsInstanceResourcePolicy](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#getaccessgrantsinstanceresourcepolicy)
  - [GetAccessGrantsLocation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#getaccessgrantslocation)
  - [GetAccessPoint](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#getaccesspoint)
  - [GetAccessPointConfigurationForObjectLambda](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#getaccesspointconfigurationforobjectlambda)
  - [GetAccessPointForObjectLambda](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#getaccesspointforobjectlambda)
  - [GetAccessPointPolicy](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#getaccesspointpolicy)
  - [GetAccessPointPolicyForObjectLambda](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#getaccesspointpolicyforobjectlambda)
  - [GetAccessPointPolicyStatus](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#getaccesspointpolicystatus)
  - [GetAccessPointPolicyStatusForObjectLambda](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#getaccesspointpolicystatusforobjectlambda)
  - [GetAccessPointScope](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#getaccesspointscope)
  - [GetBucket](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#getbucket)
  - [GetBucketLifecycleConfiguration](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#getbucketlifecycleconfiguration)
  - [GetBucketPolicy](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#getbucketpolicy)
  - [GetBucketReplication](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#getbucketreplication)
  - [GetBucketTagging](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#getbuckettagging)
  - [GetBucketVersioning](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#getbucketversioning)
  - [GetDataAccess](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#getdataaccess)
  - [GetJobTagging](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#getjobtagging)
  - [GetMultiRegionAccessPoint](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#getmultiregionaccesspoint)
  - [GetMultiRegionAccessPointPolicy](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#getmultiregionaccesspointpolicy)
  - [GetMultiRegionAccessPointPolicyStatus](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#getmultiregionaccesspointpolicystatus)
  - [GetMultiRegionAccessPointRoutes](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#getmultiregionaccesspointroutes)
  - [GetPublicAccessBlock](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#getpublicaccessblock)
  - [GetStorageLensConfiguration](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#getstoragelensconfiguration)
  - [GetStorageLensConfigurationTagging](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#getstoragelensconfigurationtagging)
  - [GetStorageLensGroup](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#getstoragelensgroup)
  - [ListAccessGrants](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#listaccessgrants)
  - [ListAccessGrantsInstances](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#listaccessgrantsinstances)
  - [ListAccessGrantsLocations](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#listaccessgrantslocations)
  - [ListAccessPoints](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#listaccesspoints)
  - [ListAccessPointsForDirectoryBuckets](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#listaccesspointsfordirectorybuckets)
  - [ListAccessPointsForObjectLambda](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#listaccesspointsforobjectlambda)
  - [ListCallerAccessGrants](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#listcalleraccessgrants)
  - [ListJobs](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#listjobs)
  - [ListMultiRegionAccessPoints](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#listmultiregionaccesspoints)
  - [ListRegionalBuckets](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#listregionalbuckets)
  - [ListStorageLensConfigurations](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#liststoragelensconfigurations)
  - [ListStorageLensGroups](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#liststoragelensgroups)
  - [ListTagsForResource](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#listtagsforresource)
  - [PutAccessGrantsInstanceResourcePolicy](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#putaccessgrantsinstanceresourcepolicy)
  - [PutAccessPointConfigurationForObjectLambda](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#putaccesspointconfigurationforobjectlambda)
  - [PutAccessPointPolicy](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#putaccesspointpolicy)
  - [PutAccessPointPolicyForObjectLambda](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#putaccesspointpolicyforobjectlambda)
  - [PutAccessPointScope](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#putaccesspointscope)
  - [PutBucketLifecycleConfiguration](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#putbucketlifecycleconfiguration)
  - [PutBucketPolicy](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#putbucketpolicy)
  - [PutBucketReplication](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#putbucketreplication)
  - [PutBucketTagging](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#putbuckettagging)
  - [PutBucketVersioning](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#putbucketversioning)
  - [PutJobTagging](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#putjobtagging)
  - [PutMultiRegionAccessPointPolicy](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#putmultiregionaccesspointpolicy)
  - [PutPublicAccessBlock](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#putpublicaccessblock)
  - [PutStorageLensConfiguration](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#putstoragelensconfiguration)
  - [PutStorageLensConfigurationTagging](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#putstoragelensconfigurationtagging)
  - [SubmitMultiRegionAccessPointRoutes](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#submitmultiregionaccesspointroutes)
  - [TagResource](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#tagresource)
  - [UntagResource](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#untagresource)
  - [UpdateAccessGrantsLocation](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#updateaccessgrantslocation)
  - [UpdateJobPriority](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#updatejobpriority)
  - [UpdateJobStatus](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#updatejobstatus)
  - [UpdateStorageLensGroup](https://docs.aws.amazon.com/aws-sdk-php/v3/api/api-s3control-2018-08-20.html#updatestoragelensgroup)

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3Control.S3ControlClient.html\#toc)

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3Control.S3ControlClient.html\#toc-methods)

[\_\_call()](class-aws-awsclienttrait.md#method___call)
: mixed [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3Control.S3ControlClient.html#method___construct)
: mixed The client constructor accepts the following options:[\_\_sleep()](class-aws-awsclient.md#method___sleep)
: mixed [\_apply\_use\_arn\_region()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3Control.S3ControlClient.html#method__apply_use_arn_region)
: mixed [execute()](class-aws-awsclienttrait.md#method_execute)
: mixed [executeAsync()](class-aws-awsclienttrait.md#method_executeAsync)
: mixed [factory()](class-aws-awsclient.md#method_factory)
: static [getApi()](class-aws-awsclienttrait.md#method_getApi)
: [Service](class-aws-api-service.md)[getArguments()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3Control.S3ControlClient.html#method_getArguments)
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

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3Control.S3ControlClient.html\#methods)

#### \_\_call()  [header link](class-aws-awsclienttrait.md\#method___call)

`
    public
                    __call(mixed $name, array<string|int, mixed> $args) : mixed`

##### Parameters

$name
: mixed$args
: array<string\|int, mixed>

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3Control.S3ControlClient.html\#method___construct)

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

#### \_apply\_use\_arn\_region()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3Control.S3ControlClient.html\#method__apply_use_arn_region)

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

#### getArguments()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3Control.S3ControlClient.html\#method_getArguments)

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
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3Control.S3ControlClient.html#toc-methods)
- Methods
  - [\_\_call()](class-aws-awsclienttrait.md#method___call)
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3Control.S3ControlClient.html#method___construct)
  - [\_\_sleep()](class-aws-awsclient.md#method___sleep)
  - [\_apply\_use\_arn\_region()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3Control.S3ControlClient.html#method__apply_use_arn_region)
  - [execute()](class-aws-awsclienttrait.md#method_execute)
  - [executeAsync()](class-aws-awsclienttrait.md#method_executeAsync)
  - [factory()](class-aws-awsclient.md#method_factory)
  - [getApi()](class-aws-awsclienttrait.md#method_getApi)
  - [getArguments()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3Control.S3ControlClient.html#method_getArguments)
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

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3Control.S3ControlClient.html#top)
