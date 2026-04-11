Menu

- [Aws](namespace-aws.md)
- [S3](namespace-aws-s3.md)

## S3Client     extends [AwsClient](class-aws-awsclient.md)   in package    - [Aws](package-aws.md)       implements  [S3ClientInterface](class-aws-s3-s3clientinterface.md)  Uses  [S3ClientTrait](class-aws-s3-s3clienttrait.md)

Client used to interact with **Amazon Simple Storage Service (Amazon S3)**.

## Supported API Versions

This class uses a _service description model_ that is associated at
runtime based on the `version` option given when constructing the
client. The `version` option will determine which API operations,
waiters, and paginators are available for a client. Creating a command or a
specific API operation can be done using magic methods (e.g.,
`$client->commandName(/** parameters */)`, or using the
`$client->getCommand` method of the client.

- [**2006-03-01**](api-s3-2006-03-01.md)

  - [AbortMultipartUpload](api-s3-2006-03-01.md#abortmultipartupload)
  - [CompleteMultipartUpload](api-s3-2006-03-01.md#completemultipartupload)
  - [CopyObject](api-s3-2006-03-01.md#copyobject)
  - [CreateBucket](api-s3-2006-03-01.md#createbucket)
  - [CreateBucketMetadataConfiguration](api-s3-2006-03-01.md#createbucketmetadataconfiguration)
  - [CreateBucketMetadataTableConfiguration](api-s3-2006-03-01.md#createbucketmetadatatableconfiguration)
  - [CreateMultipartUpload](api-s3-2006-03-01.md#createmultipartupload)
  - [CreateSession](api-s3-2006-03-01.md#createsession)
  - [DeleteBucket](api-s3-2006-03-01.md#deletebucket)
  - [DeleteBucketAnalyticsConfiguration](api-s3-2006-03-01.md#deletebucketanalyticsconfiguration)
  - [DeleteBucketCors](api-s3-2006-03-01.md#deletebucketcors)
  - [DeleteBucketEncryption](api-s3-2006-03-01.md#deletebucketencryption)
  - [DeleteBucketIntelligentTieringConfiguration](api-s3-2006-03-01.md#deletebucketintelligenttieringconfiguration)
  - [DeleteBucketInventoryConfiguration](api-s3-2006-03-01.md#deletebucketinventoryconfiguration)
  - [DeleteBucketLifecycle](api-s3-2006-03-01.md#deletebucketlifecycle)
  - [DeleteBucketMetadataConfiguration](api-s3-2006-03-01.md#deletebucketmetadataconfiguration)
  - [DeleteBucketMetadataTableConfiguration](api-s3-2006-03-01.md#deletebucketmetadatatableconfiguration)
  - [DeleteBucketMetricsConfiguration](api-s3-2006-03-01.md#deletebucketmetricsconfiguration)
  - [DeleteBucketOwnershipControls](api-s3-2006-03-01.md#deletebucketownershipcontrols)
  - [DeleteBucketPolicy](api-s3-2006-03-01.md#deletebucketpolicy)
  - [DeleteBucketReplication](api-s3-2006-03-01.md#deletebucketreplication)
  - [DeleteBucketTagging](api-s3-2006-03-01.md#deletebuckettagging)
  - [DeleteBucketWebsite](api-s3-2006-03-01.md#deletebucketwebsite)
  - [DeleteObject](api-s3-2006-03-01.md#deleteobject)
  - [DeleteObjectTagging](api-s3-2006-03-01.md#deleteobjecttagging)
  - [DeleteObjects](api-s3-2006-03-01.md#deleteobjects)
  - [DeletePublicAccessBlock](api-s3-2006-03-01.md#deletepublicaccessblock)
  - [GetBucketAbac](api-s3-2006-03-01.md#getbucketabac)
  - [GetBucketAccelerateConfiguration](api-s3-2006-03-01.md#getbucketaccelerateconfiguration)
  - [GetBucketAcl](api-s3-2006-03-01.md#getbucketacl)
  - [GetBucketAnalyticsConfiguration](api-s3-2006-03-01.md#getbucketanalyticsconfiguration)
  - [GetBucketCors](api-s3-2006-03-01.md#getbucketcors)
  - [GetBucketEncryption](api-s3-2006-03-01.md#getbucketencryption)
  - [GetBucketIntelligentTieringConfiguration](api-s3-2006-03-01.md#getbucketintelligenttieringconfiguration)
  - [GetBucketInventoryConfiguration](api-s3-2006-03-01.md#getbucketinventoryconfiguration)
  - [GetBucketLifecycle](api-s3-2006-03-01.md#getbucketlifecycle)
  - [GetBucketLifecycleConfiguration](api-s3-2006-03-01.md#getbucketlifecycleconfiguration)
  - [GetBucketLocation](api-s3-2006-03-01.md#getbucketlocation)
  - [GetBucketLogging](api-s3-2006-03-01.md#getbucketlogging)
  - [GetBucketMetadataConfiguration](api-s3-2006-03-01.md#getbucketmetadataconfiguration)
  - [GetBucketMetadataTableConfiguration](api-s3-2006-03-01.md#getbucketmetadatatableconfiguration)
  - [GetBucketMetricsConfiguration](api-s3-2006-03-01.md#getbucketmetricsconfiguration)
  - [GetBucketNotification](api-s3-2006-03-01.md#getbucketnotification)
  - [GetBucketNotificationConfiguration](api-s3-2006-03-01.md#getbucketnotificationconfiguration)
  - [GetBucketOwnershipControls](api-s3-2006-03-01.md#getbucketownershipcontrols)
  - [GetBucketPolicy](api-s3-2006-03-01.md#getbucketpolicy)
  - [GetBucketPolicyStatus](api-s3-2006-03-01.md#getbucketpolicystatus)
  - [GetBucketReplication](api-s3-2006-03-01.md#getbucketreplication)
  - [GetBucketRequestPayment](api-s3-2006-03-01.md#getbucketrequestpayment)
  - [GetBucketTagging](api-s3-2006-03-01.md#getbuckettagging)
  - [GetBucketVersioning](api-s3-2006-03-01.md#getbucketversioning)
  - [GetBucketWebsite](api-s3-2006-03-01.md#getbucketwebsite)
  - [GetObject](api-s3-2006-03-01.md#getobject)
  - [GetObjectAcl](api-s3-2006-03-01.md#getobjectacl)
  - [GetObjectAttributes](api-s3-2006-03-01.md#getobjectattributes)
  - [GetObjectLegalHold](api-s3-2006-03-01.md#getobjectlegalhold)
  - [GetObjectLockConfiguration](api-s3-2006-03-01.md#getobjectlockconfiguration)
  - [GetObjectRetention](api-s3-2006-03-01.md#getobjectretention)
  - [GetObjectTagging](api-s3-2006-03-01.md#getobjecttagging)
  - [GetObjectTorrent](api-s3-2006-03-01.md#getobjecttorrent)
  - [GetPublicAccessBlock](api-s3-2006-03-01.md#getpublicaccessblock)
  - [HeadBucket](api-s3-2006-03-01.md#headbucket)
  - [HeadObject](api-s3-2006-03-01.md#headobject)
  - [ListBucketAnalyticsConfigurations](api-s3-2006-03-01.md#listbucketanalyticsconfigurations)
  - [ListBucketIntelligentTieringConfigurations](api-s3-2006-03-01.md#listbucketintelligenttieringconfigurations)
  - [ListBucketInventoryConfigurations](api-s3-2006-03-01.md#listbucketinventoryconfigurations)
  - [ListBucketMetricsConfigurations](api-s3-2006-03-01.md#listbucketmetricsconfigurations)
  - [ListBuckets](api-s3-2006-03-01.md#listbuckets)
  - [ListDirectoryBuckets](api-s3-2006-03-01.md#listdirectorybuckets)
  - [ListMultipartUploads](api-s3-2006-03-01.md#listmultipartuploads)
  - [ListObjectVersions](api-s3-2006-03-01.md#listobjectversions)
  - [ListObjects](api-s3-2006-03-01.md#listobjects)
  - [ListObjectsV2](api-s3-2006-03-01.md#listobjectsv2)
  - [ListParts](api-s3-2006-03-01.md#listparts)
  - [PutBucketAbac](api-s3-2006-03-01.md#putbucketabac)
  - [PutBucketAccelerateConfiguration](api-s3-2006-03-01.md#putbucketaccelerateconfiguration)
  - [PutBucketAcl](api-s3-2006-03-01.md#putbucketacl)
  - [PutBucketAnalyticsConfiguration](api-s3-2006-03-01.md#putbucketanalyticsconfiguration)
  - [PutBucketCors](api-s3-2006-03-01.md#putbucketcors)
  - [PutBucketEncryption](api-s3-2006-03-01.md#putbucketencryption)
  - [PutBucketIntelligentTieringConfiguration](api-s3-2006-03-01.md#putbucketintelligenttieringconfiguration)
  - [PutBucketInventoryConfiguration](api-s3-2006-03-01.md#putbucketinventoryconfiguration)
  - [PutBucketLifecycle](api-s3-2006-03-01.md#putbucketlifecycle)
  - [PutBucketLifecycleConfiguration](api-s3-2006-03-01.md#putbucketlifecycleconfiguration)
  - [PutBucketLogging](api-s3-2006-03-01.md#putbucketlogging)
  - [PutBucketMetricsConfiguration](api-s3-2006-03-01.md#putbucketmetricsconfiguration)
  - [PutBucketNotification](api-s3-2006-03-01.md#putbucketnotification)
  - [PutBucketNotificationConfiguration](api-s3-2006-03-01.md#putbucketnotificationconfiguration)
  - [PutBucketOwnershipControls](api-s3-2006-03-01.md#putbucketownershipcontrols)
  - [PutBucketPolicy](api-s3-2006-03-01.md#putbucketpolicy)
  - [PutBucketReplication](api-s3-2006-03-01.md#putbucketreplication)
  - [PutBucketRequestPayment](api-s3-2006-03-01.md#putbucketrequestpayment)
  - [PutBucketTagging](api-s3-2006-03-01.md#putbuckettagging)
  - [PutBucketVersioning](api-s3-2006-03-01.md#putbucketversioning)
  - [PutBucketWebsite](api-s3-2006-03-01.md#putbucketwebsite)
  - [PutObject](api-s3-2006-03-01.md#putobject)
  - [PutObjectAcl](api-s3-2006-03-01.md#putobjectacl)
  - [PutObjectLegalHold](api-s3-2006-03-01.md#putobjectlegalhold)
  - [PutObjectLockConfiguration](api-s3-2006-03-01.md#putobjectlockconfiguration)
  - [PutObjectRetention](api-s3-2006-03-01.md#putobjectretention)
  - [PutObjectTagging](api-s3-2006-03-01.md#putobjecttagging)
  - [PutPublicAccessBlock](api-s3-2006-03-01.md#putpublicaccessblock)
  - [RenameObject](api-s3-2006-03-01.md#renameobject)
  - [RestoreObject](api-s3-2006-03-01.md#restoreobject)
  - [SelectObjectContent](api-s3-2006-03-01.md#selectobjectcontent)
  - [UpdateBucketMetadataInventoryTableConfiguration](api-s3-2006-03-01.md#updatebucketmetadatainventorytableconfiguration)
  - [UpdateBucketMetadataJournalTableConfiguration](api-s3-2006-03-01.md#updatebucketmetadatajournaltableconfiguration)
  - [UpdateObjectEncryption](api-s3-2006-03-01.md#updateobjectencryption)
  - [UploadPart](api-s3-2006-03-01.md#uploadpart)
  - [UploadPartCopy](api-s3-2006-03-01.md#uploadpartcopy)
  - [WriteGetObjectResponse](api-s3-2006-03-01.md#writegetobjectresponse)

## Examples

### Basics, Actions and Scenarios

The following code examples show you how to perform actions and implement common scenarios by using the AWS SDK for PHP with Amazon Simple Storage Service.

- [See examples on AWS Docs](../../../sdk-for-php/v3/developer-guide/php-s3-code-examples.md)

### Legacy Code Examples With Guidance

The following examples demonstrate how to use this service with the AWS SDK for PHP. These code examples are available in the [AWS SDK for PHP Developer Guide](../../../sdk-for-php/v3/developer-guide/s3-examples.md).

- [Creating and using Amazon S3 buckets](../../../sdk-for-php/v3/developer-guide/s3-examples-creating-buckets.md)
- [Managing Amazon S3 bucket access permissions](../../../sdk-for-php/v3/developer-guide/s3-examples-access-permissions.md)
- [Configuring Amazon S3 buckets](../../../sdk-for-php/v3/developer-guide/s3-examples-configuring-a-bucket.md)
- [Amazon S3 multipart uploads](../../../sdk-for-php/v3/developer-guide/s3-multipart-upload.md)
- [Amazon S3 pre-signed URL](../../../sdk-for-php/v3/developer-guide/s3-presigned-url.md)
- [Creating S3 pre-signed POSTs](../../../sdk-for-php/v3/developer-guide/s3-presigned-post.md)
- [Using an Amazon S3 bucket as a static web host](../../../sdk-for-php/v3/developer-guide/s3-examples-static-web-host.md)
- [Working with Amazon S3 bucket policies](../../../sdk-for-php/v3/developer-guide/s3-examples-bucket-policies.md)
- [Using S3 access point ARNs](../../../sdk-for-php/v3/developer-guide/s3-examples-access-point-arn.md)
- [Use Multi-Region Access Points](../../../sdk-for-php/v3/developer-guide/s3-multi-region-access-points.md)

### Table of Contents  [header link](class-aws-s3-s3client-toc.md)

#### Interfaces  [header link](class-aws-s3-s3client-toc-interfaces.md)

[S3ClientInterface](class-aws-s3-s3clientinterface.md)\*\*Amazon Simple Storage Service\*\* client.

#### Methods  [header link](class-aws-s3-s3client-toc-methods.md)

[\_\_call()](class-aws-awsclienttrait-method-call.md)
: mixed [\_\_construct()](class-aws-s3-s3client-method-construct.md)
: mixed The client constructor accepts the following options:[\_\_sleep()](class-aws-awsclient-method-sleep.md)
: mixed [\_apply\_request\_checksum\_calculation()](class-aws-s3-s3client-method-apply-request-checksum-calculation.md)
: void [\_apply\_response\_checksum\_validation()](class-aws-s3-s3client-method-apply-response-checksum-validation.md)
: void [\_apply\_use\_arn\_region()](class-aws-s3-s3client-method-apply-use-arn-region.md)
: mixed [\_default\_disable\_express\_session\_auth()](class-aws-s3-s3client-method-default-disable-express-session-auth.md)
: mixed [\_default\_request\_checksum\_calculation()](class-aws-s3-s3client-method-default-request-checksum-calculation.md)
: string [\_default\_response\_checksum\_validation()](class-aws-s3-s3client-method-default-response-checksum-validation.md)
: string [\_default\_s3\_express\_identity\_provider()](class-aws-s3-s3client-method-default-s3-express-identity-provider.md)
: mixed [copy()](class-aws-s3-s3clienttrait-method-copy.md)
: mixed [copyAsync()](class-aws-s3-s3clienttrait-method-copyasync.md)
: mixed [createPresignedRequest()](class-aws-s3-s3client-method-createpresignedrequest.md)
: [RequestInterface](class-psr-http-message-requestinterface.md)Create a pre-signed URL for the given S3 command object.[deleteMatchingObjects()](class-aws-s3-s3clienttrait-method-deletematchingobjects.md)
: mixed [deleteMatchingObjectsAsync()](class-aws-s3-s3clienttrait-method-deletematchingobjectsasync.md)
: mixed [determineBucketRegion()](class-aws-s3-s3clienttrait-method-determinebucketregion.md)
: mixed [determineBucketRegionAsync()](class-aws-s3-s3clienttrait-method-determinebucketregionasync.md)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)[doesBucketExist()](class-aws-s3-s3clienttrait-method-doesbucketexist.md)
: mixed [doesBucketExistV2()](class-aws-s3-s3clienttrait-method-doesbucketexistv2.md)
: mixed [doesObjectExist()](class-aws-s3-s3clienttrait-method-doesobjectexist.md)
: mixed [doesObjectExistV2()](class-aws-s3-s3clienttrait-method-doesobjectexistv2.md)
: mixed [downloadBucket()](class-aws-s3-s3clienttrait-method-downloadbucket.md)
: mixed [downloadBucketAsync()](class-aws-s3-s3clienttrait-method-downloadbucketasync.md)
: mixed [encodeKey()](class-aws-s3-s3client-method-encodekey.md)
: string Raw URL encode a key and allow for '/' characters[execute()](class-aws-awsclienttrait-method-execute.md)
: mixed [executeAsync()](class-aws-awsclienttrait-method-executeasync.md)
: mixed [factory()](class-aws-awsclient-method-factory.md)
: static [getApi()](class-aws-awsclienttrait-method-getapi.md)
: [Service](class-aws-api-service.md)[getArguments()](class-aws-s3-s3client-method-getarguments.md)
: array<string\|int, mixed> Get an array of client constructor arguments used by the client.[getClientBuiltIns()](class-aws-awsclient-method-getclientbuiltins.md)
: array<string\|int, mixed> Provides the set of built-in keys and values
used for endpoint resolution[getClientContextParams()](class-aws-awsclient-method-getclientcontextparams.md)
: array<string\|int, mixed> Provides the set of service context parameter
key-value pairs used for endpoint resolution.[getCommand()](class-aws-awsclienttrait-method-getcommand.md)
: [CommandInterface](class-aws-commandinterface.md)[getConfig()](class-aws-awsclient-method-getconfig.md)
: mixed\|null Get a client configuration value.[getCredentials()](class-aws-awsclient-method-getcredentials.md)
: [PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)Returns a promise that is fulfilled with an
{@see \\Aws\\Credentials\\CredentialsInterface} object.[getEndpoint()](class-aws-awsclient-method-getendpoint.md)
: [UriInterface](class-psr-http-message-uriinterface.md)Gets the default endpoint, or base URL, used by the client.[getEndpointProvider()](class-aws-awsclient-method-getendpointprovider.md)
: mixed [getEndpointProviderArgs()](class-aws-awsclient-method-getendpointproviderargs.md)
: array<string\|int, mixed> Retrieves arguments to be used in endpoint resolution.[getHandlerList()](class-aws-awsclient-method-gethandlerlist.md)
: [HandlerList](class-aws-handlerlist.md)Get the handler list used to transfer commands.[getIterator()](class-aws-awsclienttrait-method-getiterator.md)
: mixed [getObjectUrl()](class-aws-s3-s3client-method-getobjecturl.md)
: string Returns the URL to an object identified by its bucket and key.[getPaginator()](class-aws-awsclienttrait-method-getpaginator.md)
: mixed [getRegion()](class-aws-awsclient-method-getregion.md)
: string Get the region to which the client is configured to send requests.[getSignatureProvider()](class-aws-awsclient-method-getsignatureprovider.md)
: callable Get the signature\_provider function of the client.[getToken()](class-aws-awsclient-method-gettoken.md)
: mixed [getWaiter()](class-aws-awsclienttrait-method-getwaiter.md)
: mixed [isBucketDnsCompatible()](class-aws-s3-s3client-method-isbucketdnscompatible.md)
: bool Determine if a string is a valid name for a DNS compatible Amazon S3
bucket.[isDirectoryBucket()](class-aws-s3-s3client-method-isdirectorybucket.md)
: bool Determines whether a bucket is a directory bucket.[registerStreamWrapper()](class-aws-s3-s3clienttrait-method-registerstreamwrapper.md)
: mixed [registerStreamWrapperV2()](class-aws-s3-s3clienttrait-method-registerstreamwrapperv2.md)
: mixed [upload()](class-aws-s3-s3clienttrait-method-upload.md)
: mixed [uploadAsync()](class-aws-s3-s3clienttrait-method-uploadasync.md)
: mixed [uploadDirectory()](class-aws-s3-s3clienttrait-method-uploaddirectory.md)
: mixed [uploadDirectoryAsync()](class-aws-s3-s3clienttrait-method-uploaddirectoryasync.md)
: mixed [waitUntil()](class-aws-awsclienttrait-method-waituntil.md)
: mixed

### Methods  [header link](class-aws-s3-s3client-methods.md)

#### \_\_call()  [header link](class-aws-awsclienttrait-method-call.md)

`
    public
                    __call(mixed $name, array<string|int, mixed> $args) : mixed`

##### Parameters

$name
: mixed$args
: array<string\|int, mixed>

#### \_\_construct()  [header link](class-aws-s3-s3client-method-construct.md)

The client constructor accepts the following options:

`
    public
                    __construct(array<string|int, mixed> $args) : mixed`

In addition to the options available to
AwsClient::\_\_construct, S3Client accepts the following
options:

- bucket\_endpoint: (bool) Set to true to send requests to a
hardcoded bucket endpoint rather than create an endpoint as a result
of injecting the bucket into the URL. This option is useful for
interacting with CNAME endpoints. Note: if you are using version 2.243.0
and above and do not expect the bucket name to appear in the host, you will
also need to set `use_path_style_endpoint` to `true`.
- calculate\_md5: (bool) Set to false to disable calculating an MD5
for all Amazon S3 signed uploads.
- s3\_us\_east\_1\_regional\_endpoint:
(Aws\\S3\\RegionalEndpoint\\ConfigurationInterface\|Aws\\CacheInterface\|callable\|string\|array)
Specifies whether to use regional or legacy endpoints for the us-east-1
region. Provide an Aws\\S3\\RegionalEndpoint\\ConfigurationInterface object, an
instance of Aws\\CacheInterface, a callable configuration provider used
to create endpoint configuration, a string value of `legacy` or
`regional`, or an associative array with the following keys:
endpoint\_types: (string) Set to `legacy` or `regional`, defaults to
`legacy`
- use\_accelerate\_endpoint: (bool) Set to true to send requests to an S3
Accelerate endpoint by default. Can be enabled or disabled on
individual operations by setting '@use\_accelerate\_endpoint' to true or
false. Note: you must enable S3 Accelerate on a bucket before it can be
accessed via an Accelerate endpoint.
- use\_arn\_region: (Aws\\S3\\UseArnRegion\\ConfigurationInterface,
Aws\\CacheInterface, bool, callable) Set to true to enable the client
to use the region from a supplied ARN argument instead of the client's
region. Provide an instance of Aws\\S3\\UseArnRegion\\ConfigurationInterface,
an instance of Aws\\CacheInterface, a callable that provides a promise for
a Configuration object, or a boolean value. Defaults to false (i.e.
the SDK will not follow the ARN region if it conflicts with the client
region and instead throw an error).
- use\_dual\_stack\_endpoint: (bool) Set to true to send requests to an S3
Dual Stack endpoint by default, which enables IPv6 Protocol.
Can be enabled or disabled on individual operations by setting
'@use\_dual\_stack\_endpoint' to true or false. Note:
you cannot use it together with an accelerate endpoint.
- use\_path\_style\_endpoint: (bool) Set to true to send requests to an S3
path style endpoint by default.
Can be enabled or disabled on individual operations by setting
'@use\_path\_style\_endpoint' to true or false. Note:
you cannot use it together with an accelerate endpoint.
- disable\_multiregion\_access\_points: (bool) Set to true to disable
sending multi region requests. They are enabled by default.
Can be enabled or disabled on individual operations by setting
'@disable\_multiregion\_access\_points' to true or false. Note:
you cannot use it together with an accelerate or dualstack endpoint.

##### Parameters

$args
: array<string\|int, mixed>

#### \_\_sleep()  [header link](class-aws-awsclient-method-sleep.md)

`
    public
                    __sleep() : mixed`

#### \_apply\_request\_checksum\_calculation()  [header link](class-aws-s3-s3client-method-apply-request-checksum-calculation.md)

`
    public
            static        _apply_request_checksum_calculation(string $value, array<string|int, mixed> &$args) : void`

##### Parameters

$value
: string$args
: array<string\|int, mixed>

#### \_apply\_response\_checksum\_validation()  [header link](class-aws-s3-s3client-method-apply-response-checksum-validation.md)

`
    public
            static        _apply_response_checksum_validation(mixed $value, array<string|int, mixed> &$args) : void`

##### Parameters

$value
: mixed$args
: array<string\|int, mixed>

#### \_apply\_use\_arn\_region()  [header link](class-aws-s3-s3client-method-apply-use-arn-region.md)

`
    public
            static        _apply_use_arn_region(mixed $value, array<string|int, mixed> &$args, HandlerList $list) : mixed`

##### Parameters

$value
: mixed$args
: array<string\|int, mixed>$list
: [HandlerList](class-aws-handlerlist.md)

#### \_default\_disable\_express\_session\_auth()  [header link](class-aws-s3-s3client-method-default-disable-express-session-auth.md)

`
    public
            static        _default_disable_express_session_auth(array<string|int, mixed> &$args) : mixed`

##### Parameters

$args
: array<string\|int, mixed>

#### \_default\_request\_checksum\_calculation()  [header link](class-aws-s3-s3client-method-default-request-checksum-calculation.md)

`
    public
            static        _default_request_checksum_calculation(array<string|int, mixed> $args) : string`

##### Parameters

$args
: array<string\|int, mixed>

##### Return values

string

#### \_default\_response\_checksum\_validation()  [header link](class-aws-s3-s3client-method-default-response-checksum-validation.md)

`
    public
            static        _default_response_checksum_validation(array<string|int, mixed> $args) : string`

##### Parameters

$args
: array<string\|int, mixed>

##### Return values

string

#### \_default\_s3\_express\_identity\_provider()  [header link](class-aws-s3-s3client-method-default-s3-express-identity-provider.md)

`
    public
            static        _default_s3_express_identity_provider(array<string|int, mixed> $args) : mixed`

##### Parameters

$args
: array<string\|int, mixed>

#### copy()  [header link](class-aws-s3-s3clienttrait-method-copy.md)

`
    public
                    copy(mixed $fromB, mixed $fromK, mixed $destB, mixed $destK[, mixed $acl = 'private' ][, array<string|int, mixed> $opts = [] ]) : mixed`

##### Parameters

$fromB
: mixed$fromK
: mixed$destB
: mixed$destK
: mixed$acl
: mixed
= 'private'$opts
: array<string\|int, mixed>
= \[\]

##### Tags  [header link](class-aws-s3-s3clienttrait-method-copy-tags.md)

see[S3ClientInterface::copy()](class-aws-s3-s3clientinterface-method-copy.md)

#### copyAsync()  [header link](class-aws-s3-s3clienttrait-method-copyasync.md)

`
    public
                    copyAsync(mixed $fromB, mixed $fromK, mixed $destB, mixed $destK[, mixed $acl = 'private' ][, array<string|int, mixed> $opts = [] ]) : mixed`

##### Parameters

$fromB
: mixed$fromK
: mixed$destB
: mixed$destK
: mixed$acl
: mixed
= 'private'$opts
: array<string\|int, mixed>
= \[\]

##### Tags  [header link](class-aws-s3-s3clienttrait-method-copyasync-tags.md)

see[S3ClientInterface::copyAsync()](class-aws-s3-s3clientinterface-method-copyasync.md)

#### createPresignedRequest()  [header link](class-aws-s3-s3client-method-createpresignedrequest.md)

Create a pre-signed URL for the given S3 command object.

`
    public
                    createPresignedRequest(CommandInterface $command, mixed $expires[, array<string|int, mixed> $options = [] ]) : RequestInterface`

##### Parameters

$command
: [CommandInterface](class-aws-commandinterface.md)

Command to create a pre-signed
URL for.

$expires
: mixed

The time at which the URL should
expire. This can be a Unix
timestamp, a PHP DateTime object,
or a string that can be evaluated
by strtotime().

$options
: array<string\|int, mixed>
= \[\]

##### Return values

[RequestInterface](class-psr-http-message-requestinterface.md)

#### deleteMatchingObjects()  [header link](class-aws-s3-s3clienttrait-method-deletematchingobjects.md)

`
    public
                    deleteMatchingObjects(mixed $bucket[, mixed $prefix = '' ][, mixed $regex = '' ][, array<string|int, mixed> $options = [] ]) : mixed`

##### Parameters

$bucket
: mixed$prefix
: mixed
= ''$regex
: mixed
= ''$options
: array<string\|int, mixed>
= \[\]

##### Tags  [header link](class-aws-s3-s3clienttrait-method-deletematchingobjects-tags.md)

see[S3ClientInterface::deleteMatchingObjects()](class-aws-s3-s3clientinterface-method-deletematchingobjects.md)

#### deleteMatchingObjectsAsync()  [header link](class-aws-s3-s3clienttrait-method-deletematchingobjectsasync.md)

`
    public
                    deleteMatchingObjectsAsync(mixed $bucket[, mixed $prefix = '' ][, mixed $regex = '' ][, array<string|int, mixed> $options = [] ]) : mixed`

##### Parameters

$bucket
: mixed$prefix
: mixed
= ''$regex
: mixed
= ''$options
: array<string\|int, mixed>
= \[\]

##### Tags  [header link](class-aws-s3-s3clienttrait-method-deletematchingobjectsasync-tags.md)

see[S3ClientInterface::deleteMatchingObjectsAsync()](class-aws-s3-s3clientinterface-method-deletematchingobjectsasync.md)

#### determineBucketRegion()  [header link](class-aws-s3-s3clienttrait-method-determinebucketregion.md)

`
    public
                    determineBucketRegion(mixed $bucketName) : mixed`

##### Parameters

$bucketName
: mixed

##### Tags  [header link](class-aws-s3-s3clienttrait-method-determinebucketregion-tags.md)

see[S3ClientInterface::determineBucketRegion()](class-aws-s3-s3clientinterface-method-determinebucketregion.md)

#### determineBucketRegionAsync()  [header link](class-aws-s3-s3clienttrait-method-determinebucketregionasync.md)

`
    public
                    determineBucketRegionAsync(string $bucketName) : PromiseInterface`

##### Parameters

$bucketName
: string

##### Tags  [header link](class-aws-s3-s3clienttrait-method-determinebucketregionasync-tags.md)

see[S3ClientInterface::determineBucketRegionAsync()](class-aws-s3-s3clientinterface-method-determinebucketregionasync.md)

##### Return values

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)

#### doesBucketExist()  [header link](class-aws-s3-s3clienttrait-method-doesbucketexist.md)

`
    public
                    doesBucketExist(mixed $bucket) : mixed`

##### Parameters

$bucket
: mixed

##### Tags  [header link](class-aws-s3-s3clienttrait-method-doesbucketexist-tags.md)

see[S3ClientInterface::doesBucketExist()](class-aws-s3-s3clientinterface-method-doesbucketexist.md)

#### doesBucketExistV2()  [header link](class-aws-s3-s3clienttrait-method-doesbucketexistv2.md)

`
    public
                    doesBucketExistV2(mixed $bucket[, mixed $accept403 = false ]) : mixed`

##### Parameters

$bucket
: mixed$accept403
: mixed
= false

##### Tags  [header link](class-aws-s3-s3clienttrait-method-doesbucketexistv2-tags.md)

see[S3ClientInterface::doesBucketExistV2()](class-aws-s3-s3clientinterface-method-doesbucketexistv2.md)

#### doesObjectExist()  [header link](class-aws-s3-s3clienttrait-method-doesobjectexist.md)

`
    public
                    doesObjectExist(mixed $bucket, mixed $key[, array<string|int, mixed> $options = [] ]) : mixed`

##### Parameters

$bucket
: mixed$key
: mixed$options
: array<string\|int, mixed>
= \[\]

##### Tags  [header link](class-aws-s3-s3clienttrait-method-doesobjectexist-tags.md)

see[S3ClientInterface::doesObjectExist()](class-aws-s3-s3clientinterface-method-doesobjectexist.md)

#### doesObjectExistV2()  [header link](class-aws-s3-s3clienttrait-method-doesobjectexistv2.md)

`
    public
                    doesObjectExistV2(mixed $bucket, mixed $key[, mixed $includeDeleteMarkers = false ][, array<string|int, mixed> $options = [] ]) : mixed`

##### Parameters

$bucket
: mixed$key
: mixed$includeDeleteMarkers
: mixed
= false$options
: array<string\|int, mixed>
= \[\]

##### Tags  [header link](class-aws-s3-s3clienttrait-method-doesobjectexistv2-tags.md)

see[S3ClientInterface::doesObjectExistV2()](class-aws-s3-s3clientinterface-method-doesobjectexistv2.md)

#### downloadBucket()  [header link](class-aws-s3-s3clienttrait-method-downloadbucket.md)

`
    public
                    downloadBucket(mixed $directory, mixed $bucket[, mixed $keyPrefix = '' ][, array<string|int, mixed> $options = [] ]) : mixed`

##### Parameters

$directory
: mixed$bucket
: mixed$keyPrefix
: mixed
= ''$options
: array<string\|int, mixed>
= \[\]

##### Tags  [header link](class-aws-s3-s3clienttrait-method-downloadbucket-tags.md)

see[S3ClientInterface::downloadBucket()](class-aws-s3-s3clientinterface-method-downloadbucket.md)

#### downloadBucketAsync()  [header link](class-aws-s3-s3clienttrait-method-downloadbucketasync.md)

`
    public
                    downloadBucketAsync(mixed $directory, mixed $bucket[, mixed $keyPrefix = '' ][, array<string|int, mixed> $options = [] ]) : mixed`

##### Parameters

$directory
: mixed$bucket
: mixed$keyPrefix
: mixed
= ''$options
: array<string\|int, mixed>
= \[\]

##### Tags  [header link](class-aws-s3-s3clienttrait-method-downloadbucketasync-tags.md)

see[S3ClientInterface::downloadBucketAsync()](class-aws-s3-s3clientinterface-method-downloadbucketasync.md)

#### encodeKey()  [header link](class-aws-s3-s3client-method-encodekey.md)

Raw URL encode a key and allow for '/' characters

`
    public
            static        encodeKey(string $key) : string`

##### Parameters

$key
: string

Key to encode

##### Return values

string
—

Returns the encoded key

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

#### factory()  [header link](class-aws-awsclient-method-factory.md)

`
    public
            static        factory([array<string|int, mixed> $config = [] ]) : static`

##### Parameters

$config
: array<string\|int, mixed>
= \[\]

##### Tags  [header link](class-aws-awsclient-method-factory-tags.md)

deprecated

##### Return values

static

#### getApi()  [header link](class-aws-awsclienttrait-method-getapi.md)

`
    public
    abstract                getApi() : Service`

##### Return values

[Service](class-aws-api-service.md)

#### getArguments()  [header link](class-aws-s3-s3client-method-getarguments.md)

Get an array of client constructor arguments used by the client.

`
    public
            static        getArguments() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getClientBuiltIns()  [header link](class-aws-awsclient-method-getclientbuiltins.md)

Provides the set of built-in keys and values
used for endpoint resolution

`
    public
                    getClientBuiltIns() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getClientContextParams()  [header link](class-aws-awsclient-method-getclientcontextparams.md)

Provides the set of service context parameter
key-value pairs used for endpoint resolution.

`
    public
                    getClientContextParams() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

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

#### getConfig()  [header link](class-aws-awsclient-method-getconfig.md)

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

#### getCredentials()  [header link](class-aws-awsclient-method-getcredentials.md)

Returns a promise that is fulfilled with an
{@see \\Aws\\Credentials\\CredentialsInterface} object.

`
    public
                    getCredentials() : PromiseInterface`

If you need the credentials synchronously, then call the wait() method
on the returned promise.

##### Return values

[PromiseInterface](class-guzzlehttp-promise-promiseinterface.md)

#### getEndpoint()  [header link](class-aws-awsclient-method-getendpoint.md)

Gets the default endpoint, or base URL, used by the client.

`
    public
                    getEndpoint() : UriInterface`

##### Return values

[UriInterface](class-psr-http-message-uriinterface.md)

#### getEndpointProvider()  [header link](class-aws-awsclient-method-getendpointprovider.md)

`
    public
                    getEndpointProvider() : mixed`

#### getEndpointProviderArgs()  [header link](class-aws-awsclient-method-getendpointproviderargs.md)

Retrieves arguments to be used in endpoint resolution.

`
    public
                    getEndpointProviderArgs() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getHandlerList()  [header link](class-aws-awsclient-method-gethandlerlist.md)

Get the handler list used to transfer commands.

`
    public
                    getHandlerList() : HandlerList`

This list can be modified to add middleware or to change the underlying
handler used to send HTTP requests.

##### Return values

[HandlerList](class-aws-handlerlist.md)

#### getIterator()  [header link](class-aws-awsclienttrait-method-getiterator.md)

`
    public
                    getIterator(mixed $name[, array<string|int, mixed> $args = [] ]) : mixed`

##### Parameters

$name
: mixed$args
: array<string\|int, mixed>
= \[\]

#### getObjectUrl()  [header link](class-aws-s3-s3client-method-getobjecturl.md)

Returns the URL to an object identified by its bucket and key.

`
    public
                    getObjectUrl(string $bucket, string $key) : string`

The URL returned by this method is not signed nor does it ensure that the
bucket and key given to the method exist. If you need a signed URL, then
use the S3Client::createPresignedRequest method and get
the URI of the signed request.

##### Parameters

$bucket
: string

The name of the bucket where the object is located

$key
: string

The key of the object

##### Return values

string
—

The URL to the object

#### getPaginator()  [header link](class-aws-awsclienttrait-method-getpaginator.md)

`
    public
                    getPaginator(mixed $name[, array<string|int, mixed> $args = [] ]) : mixed`

##### Parameters

$name
: mixed$args
: array<string\|int, mixed>
= \[\]

#### getRegion()  [header link](class-aws-awsclient-method-getregion.md)

Get the region to which the client is configured to send requests.

`
    public
                    getRegion() : string`

##### Return values

string

#### getSignatureProvider()  [header link](class-aws-awsclient-method-getsignatureprovider.md)

Get the signature\_provider function of the client.

`
    public
        final            getSignatureProvider() : callable`

##### Return values

callable

#### getToken()  [header link](class-aws-awsclient-method-gettoken.md)

`
    public
                    getToken() : mixed`

#### getWaiter()  [header link](class-aws-awsclienttrait-method-getwaiter.md)

`
    public
                    getWaiter(mixed $name[, array<string|int, mixed> $args = [] ]) : mixed`

##### Parameters

$name
: mixed$args
: array<string\|int, mixed>
= \[\]

#### isBucketDnsCompatible()  [header link](class-aws-s3-s3client-method-isbucketdnscompatible.md)

Determine if a string is a valid name for a DNS compatible Amazon S3
bucket.

`
    public
            static        isBucketDnsCompatible(string $bucket) : bool`

DNS compatible bucket names can be used as a subdomain in a URL (e.g.,
".s3.amazonaws.com").

##### Parameters

$bucket
: string

Bucket name to check.

##### Return values

bool

#### isDirectoryBucket()  [header link](class-aws-s3-s3client-method-isdirectorybucket.md)

Determines whether a bucket is a directory bucket.

`
    public
            static        isDirectoryBucket(string $bucket) : bool`

Only considers the availability zone/suffix format

##### Parameters

$bucket
: string

##### Return values

bool

#### registerStreamWrapper()  [header link](class-aws-s3-s3clienttrait-method-registerstreamwrapper.md)

`
    public
                    registerStreamWrapper() : mixed`

##### Tags  [header link](class-aws-s3-s3clienttrait-method-registerstreamwrapper-tags.md)

see[S3ClientInterface::registerStreamWrapper()](class-aws-s3-s3clientinterface-method-registerstreamwrapper.md)

#### registerStreamWrapperV2()  [header link](class-aws-s3-s3clienttrait-method-registerstreamwrapperv2.md)

`
    public
                    registerStreamWrapperV2() : mixed`

##### Tags  [header link](class-aws-s3-s3clienttrait-method-registerstreamwrapperv2-tags.md)

see[S3ClientInterface::registerStreamWrapperV2()](class-aws-s3-s3clientinterface-method-registerstreamwrapperv2.md)

#### upload()  [header link](class-aws-s3-s3clienttrait-method-upload.md)

`
    public
                    upload(mixed $bucket, mixed $key, mixed $body[, mixed $acl = 'private' ][, array<string|int, mixed> $options = [] ]) : mixed`

##### Parameters

$bucket
: mixed$key
: mixed$body
: mixed$acl
: mixed
= 'private'$options
: array<string\|int, mixed>
= \[\]

##### Tags  [header link](class-aws-s3-s3clienttrait-method-upload-tags.md)

see[S3ClientInterface::upload()](class-aws-s3-s3clientinterface-method-upload.md)

#### uploadAsync()  [header link](class-aws-s3-s3clienttrait-method-uploadasync.md)

`
    public
                    uploadAsync(mixed $bucket, mixed $key, mixed $body[, mixed $acl = 'private' ][, array<string|int, mixed> $options = [] ]) : mixed`

##### Parameters

$bucket
: mixed$key
: mixed$body
: mixed$acl
: mixed
= 'private'$options
: array<string\|int, mixed>
= \[\]

##### Tags  [header link](class-aws-s3-s3clienttrait-method-uploadasync-tags.md)

see[S3ClientInterface::uploadAsync()](class-aws-s3-s3clientinterface-method-uploadasync.md)

#### uploadDirectory()  [header link](class-aws-s3-s3clienttrait-method-uploaddirectory.md)

`
    public
                    uploadDirectory(mixed $directory, mixed $bucket[, mixed $keyPrefix = null ][, array<string|int, mixed> $options = [] ]) : mixed`

##### Parameters

$directory
: mixed$bucket
: mixed$keyPrefix
: mixed
= null$options
: array<string\|int, mixed>
= \[\]

##### Tags  [header link](class-aws-s3-s3clienttrait-method-uploaddirectory-tags.md)

see[S3ClientInterface::uploadDirectory()](class-aws-s3-s3clientinterface-method-uploaddirectory.md)

#### uploadDirectoryAsync()  [header link](class-aws-s3-s3clienttrait-method-uploaddirectoryasync.md)

`
    public
                    uploadDirectoryAsync(mixed $directory, mixed $bucket[, mixed $keyPrefix = null ][, array<string|int, mixed> $options = [] ]) : mixed`

##### Parameters

$directory
: mixed$bucket
: mixed$keyPrefix
: mixed
= null$options
: array<string\|int, mixed>
= \[\]

##### Tags  [header link](class-aws-s3-s3clienttrait-method-uploaddirectoryasync-tags.md)

see[S3ClientInterface::uploadDirectoryAsync()](class-aws-s3-s3clientinterface-method-uploaddirectoryasync.md)

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
  - [Methods](class-aws-s3-s3client-toc-methods.md)
- Methods
  - [\_\_call()](class-aws-awsclienttrait-method-call.md)
  - [\_\_construct()](class-aws-s3-s3client-method-construct.md)
  - [\_\_sleep()](class-aws-awsclient-method-sleep.md)
  - [\_apply\_request\_checksum\_calculation()](class-aws-s3-s3client-method-apply-request-checksum-calculation.md)
  - [\_apply\_response\_checksum\_validation()](class-aws-s3-s3client-method-apply-response-checksum-validation.md)
  - [\_apply\_use\_arn\_region()](class-aws-s3-s3client-method-apply-use-arn-region.md)
  - [\_default\_disable\_express\_session\_auth()](class-aws-s3-s3client-method-default-disable-express-session-auth.md)
  - [\_default\_request\_checksum\_calculation()](class-aws-s3-s3client-method-default-request-checksum-calculation.md)
  - [\_default\_response\_checksum\_validation()](class-aws-s3-s3client-method-default-response-checksum-validation.md)
  - [\_default\_s3\_express\_identity\_provider()](class-aws-s3-s3client-method-default-s3-express-identity-provider.md)
  - [copy()](class-aws-s3-s3clienttrait-method-copy.md)
  - [copyAsync()](class-aws-s3-s3clienttrait-method-copyasync.md)
  - [createPresignedRequest()](class-aws-s3-s3client-method-createpresignedrequest.md)
  - [deleteMatchingObjects()](class-aws-s3-s3clienttrait-method-deletematchingobjects.md)
  - [deleteMatchingObjectsAsync()](class-aws-s3-s3clienttrait-method-deletematchingobjectsasync.md)
  - [determineBucketRegion()](class-aws-s3-s3clienttrait-method-determinebucketregion.md)
  - [determineBucketRegionAsync()](class-aws-s3-s3clienttrait-method-determinebucketregionasync.md)
  - [doesBucketExist()](class-aws-s3-s3clienttrait-method-doesbucketexist.md)
  - [doesBucketExistV2()](class-aws-s3-s3clienttrait-method-doesbucketexistv2.md)
  - [doesObjectExist()](class-aws-s3-s3clienttrait-method-doesobjectexist.md)
  - [doesObjectExistV2()](class-aws-s3-s3clienttrait-method-doesobjectexistv2.md)
  - [downloadBucket()](class-aws-s3-s3clienttrait-method-downloadbucket.md)
  - [downloadBucketAsync()](class-aws-s3-s3clienttrait-method-downloadbucketasync.md)
  - [encodeKey()](class-aws-s3-s3client-method-encodekey.md)
  - [execute()](class-aws-awsclienttrait-method-execute.md)
  - [executeAsync()](class-aws-awsclienttrait-method-executeasync.md)
  - [factory()](class-aws-awsclient-method-factory.md)
  - [getApi()](class-aws-awsclienttrait-method-getapi.md)
  - [getArguments()](class-aws-s3-s3client-method-getarguments.md)
  - [getClientBuiltIns()](class-aws-awsclient-method-getclientbuiltins.md)
  - [getClientContextParams()](class-aws-awsclient-method-getclientcontextparams.md)
  - [getCommand()](class-aws-awsclienttrait-method-getcommand.md)
  - [getConfig()](class-aws-awsclient-method-getconfig.md)
  - [getCredentials()](class-aws-awsclient-method-getcredentials.md)
  - [getEndpoint()](class-aws-awsclient-method-getendpoint.md)
  - [getEndpointProvider()](class-aws-awsclient-method-getendpointprovider.md)
  - [getEndpointProviderArgs()](class-aws-awsclient-method-getendpointproviderargs.md)
  - [getHandlerList()](class-aws-awsclient-method-gethandlerlist.md)
  - [getIterator()](class-aws-awsclienttrait-method-getiterator.md)
  - [getObjectUrl()](class-aws-s3-s3client-method-getobjecturl.md)
  - [getPaginator()](class-aws-awsclienttrait-method-getpaginator.md)
  - [getRegion()](class-aws-awsclient-method-getregion.md)
  - [getSignatureProvider()](class-aws-awsclient-method-getsignatureprovider.md)
  - [getToken()](class-aws-awsclient-method-gettoken.md)
  - [getWaiter()](class-aws-awsclienttrait-method-getwaiter.md)
  - [isBucketDnsCompatible()](class-aws-s3-s3client-method-isbucketdnscompatible.md)
  - [isDirectoryBucket()](class-aws-s3-s3client-method-isdirectorybucket.md)
  - [registerStreamWrapper()](class-aws-s3-s3clienttrait-method-registerstreamwrapper.md)
  - [registerStreamWrapperV2()](class-aws-s3-s3clienttrait-method-registerstreamwrapperv2.md)
  - [upload()](class-aws-s3-s3clienttrait-method-upload.md)
  - [uploadAsync()](class-aws-s3-s3clienttrait-method-uploadasync.md)
  - [uploadDirectory()](class-aws-s3-s3clienttrait-method-uploaddirectory.md)
  - [uploadDirectoryAsync()](class-aws-s3-s3clienttrait-method-uploaddirectoryasync.md)
  - [waitUntil()](class-aws-awsclienttrait-method-waituntil.md)

[Back To Top](class-aws-s3-s3client-top.md)
