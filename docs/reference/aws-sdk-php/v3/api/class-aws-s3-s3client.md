Menu

- [Aws](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.html)
- [S3](https://docs.aws.amazon.com/aws-sdk-php/v3/api/namespace-Aws.s3.html)

## S3Client     extends [AwsClient](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClient.html)   in package    - [Aws](https://docs.aws.amazon.com/aws-sdk-php/v3/api/package-Aws.html)       implements  [S3ClientInterface](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientInterface.html)  Uses  [S3ClientTrait](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html)

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

- [See examples on AWS Docs](https://docs.aws.amazon.com/sdk-for-php/v3/developer-guide/php_s3_code_examples.html)

### Legacy Code Examples With Guidance

The following examples demonstrate how to use this service with the AWS SDK for PHP. These code examples are available in the [AWS SDK for PHP Developer Guide](https://docs.aws.amazon.com/sdk-for-php/v3/developer-guide/s3-examples.html).

- [Creating and using Amazon S3 buckets](https://docs.aws.amazon.com/sdk-for-php/v3/developer-guide/s3-examples-creating-buckets.html)
- [Managing Amazon S3 bucket access permissions](https://docs.aws.amazon.com/sdk-for-php/v3/developer-guide/s3-examples-access-permissions.html)
- [Configuring Amazon S3 buckets](https://docs.aws.amazon.com/sdk-for-php/v3/developer-guide/s3-examples-configuring-a-bucket.html)
- [Amazon S3 multipart uploads](https://docs.aws.amazon.com/sdk-for-php/v3/developer-guide/s3-multipart-upload.html)
- [Amazon S3 pre-signed URL](https://docs.aws.amazon.com/sdk-for-php/v3/developer-guide/s3-presigned-url.html)
- [Creating S3 pre-signed POSTs](https://docs.aws.amazon.com/sdk-for-php/v3/developer-guide/s3-presigned-post.html)
- [Using an Amazon S3 bucket as a static web host](https://docs.aws.amazon.com/sdk-for-php/v3/developer-guide/s3-examples-static-web-host.html)
- [Working with Amazon S3 bucket policies](https://docs.aws.amazon.com/sdk-for-php/v3/developer-guide/s3-examples-bucket-policies.html)
- [Using S3 access point ARNs](https://docs.aws.amazon.com/sdk-for-php/v3/developer-guide/s3-examples-access-point-arn.html)
- [Use Multi-Region Access Points](https://docs.aws.amazon.com/sdk-for-php/v3/developer-guide/s3-multi-region-access-points.html)

### Table of Contents  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Client.html\#toc)

#### Interfaces  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Client.html\#toc-interfaces)

[S3ClientInterface](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientInterface.html)\*\*Amazon Simple Storage Service\*\* client.

#### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Client.html\#toc-methods)

[\_\_call()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientTrait.html#method___call)
: mixed [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Client.html#method___construct)
: mixed The client constructor accepts the following options:[\_\_sleep()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClient.html#method___sleep)
: mixed [\_apply\_request\_checksum\_calculation()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Client.html#method__apply_request_checksum_calculation)
: void [\_apply\_response\_checksum\_validation()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Client.html#method__apply_response_checksum_validation)
: void [\_apply\_use\_arn\_region()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Client.html#method__apply_use_arn_region)
: mixed [\_default\_disable\_express\_session\_auth()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Client.html#method__default_disable_express_session_auth)
: mixed [\_default\_request\_checksum\_calculation()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Client.html#method__default_request_checksum_calculation)
: string [\_default\_response\_checksum\_validation()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Client.html#method__default_response_checksum_validation)
: string [\_default\_s3\_express\_identity\_provider()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Client.html#method__default_s3_express_identity_provider)
: mixed [copy()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_copy)
: mixed [copyAsync()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_copyAsync)
: mixed [createPresignedRequest()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Client.html#method_createPresignedRequest)
: [RequestInterface](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.RequestInterface.html)Create a pre-signed URL for the given S3 command object.[deleteMatchingObjects()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_deleteMatchingObjects)
: mixed [deleteMatchingObjectsAsync()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_deleteMatchingObjectsAsync)
: mixed [determineBucketRegion()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_determineBucketRegion)
: mixed [determineBucketRegionAsync()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_determineBucketRegionAsync)
: [PromiseInterface](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.PromiseInterface.html)[doesBucketExist()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_doesBucketExist)
: mixed [doesBucketExistV2()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_doesBucketExistV2)
: mixed [doesObjectExist()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_doesObjectExist)
: mixed [doesObjectExistV2()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_doesObjectExistV2)
: mixed [downloadBucket()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_downloadBucket)
: mixed [downloadBucketAsync()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_downloadBucketAsync)
: mixed [encodeKey()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Client.html#method_encodeKey)
: string Raw URL encode a key and allow for '/' characters[execute()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientTrait.html#method_execute)
: mixed [executeAsync()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientTrait.html#method_executeAsync)
: mixed [factory()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClient.html#method_factory)
: static [getApi()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientTrait.html#method_getApi)
: [Service](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html)[getArguments()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Client.html#method_getArguments)
: array<string\|int, mixed> Get an array of client constructor arguments used by the client.[getClientBuiltIns()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClient.html#method_getClientBuiltIns)
: array<string\|int, mixed> Provides the set of built-in keys and values
used for endpoint resolution[getClientContextParams()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClient.html#method_getClientContextParams)
: array<string\|int, mixed> Provides the set of service context parameter
key-value pairs used for endpoint resolution.[getCommand()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientTrait.html#method_getCommand)
: [CommandInterface](class-aws-commandinterface.md)[getConfig()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClient.html#method_getConfig)
: mixed\|null Get a client configuration value.[getCredentials()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClient.html#method_getCredentials)
: [PromiseInterface](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.PromiseInterface.html)Returns a promise that is fulfilled with an
{@see \\Aws\\Credentials\\CredentialsInterface} object.[getEndpoint()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClient.html#method_getEndpoint)
: [UriInterface](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.UriInterface.html)Gets the default endpoint, or base URL, used by the client.[getEndpointProvider()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClient.html#method_getEndpointProvider)
: mixed [getEndpointProviderArgs()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClient.html#method_getEndpointProviderArgs)
: array<string\|int, mixed> Retrieves arguments to be used in endpoint resolution.[getHandlerList()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClient.html#method_getHandlerList)
: [HandlerList](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html)Get the handler list used to transfer commands.[getIterator()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientTrait.html#method_getIterator)
: mixed [getObjectUrl()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Client.html#method_getObjectUrl)
: string Returns the URL to an object identified by its bucket and key.[getPaginator()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientTrait.html#method_getPaginator)
: mixed [getRegion()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClient.html#method_getRegion)
: string Get the region to which the client is configured to send requests.[getSignatureProvider()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClient.html#method_getSignatureProvider)
: callable Get the signature\_provider function of the client.[getToken()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClient.html#method_getToken)
: mixed [getWaiter()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientTrait.html#method_getWaiter)
: mixed [isBucketDnsCompatible()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Client.html#method_isBucketDnsCompatible)
: bool Determine if a string is a valid name for a DNS compatible Amazon S3
bucket.[isDirectoryBucket()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Client.html#method_isDirectoryBucket)
: bool Determines whether a bucket is a directory bucket.[registerStreamWrapper()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_registerStreamWrapper)
: mixed [registerStreamWrapperV2()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_registerStreamWrapperV2)
: mixed [upload()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_upload)
: mixed [uploadAsync()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_uploadAsync)
: mixed [uploadDirectory()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_uploadDirectory)
: mixed [uploadDirectoryAsync()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_uploadDirectoryAsync)
: mixed [waitUntil()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientTrait.html#method_waitUntil)
: mixed

### Methods  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Client.html\#methods)

#### \_\_call()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientTrait.html\#method___call)

`
    public
                    __call(mixed $name, array<string|int, mixed> $args) : mixed`

##### Parameters

$name
: mixed$args
: array<string\|int, mixed>

#### \_\_construct()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Client.html\#method___construct)

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

#### \_\_sleep()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClient.html\#method___sleep)

`
    public
                    __sleep() : mixed`

#### \_apply\_request\_checksum\_calculation()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Client.html\#method__apply_request_checksum_calculation)

`
    public
            static        _apply_request_checksum_calculation(string $value, array<string|int, mixed> &$args) : void`

##### Parameters

$value
: string$args
: array<string\|int, mixed>

#### \_apply\_response\_checksum\_validation()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Client.html\#method__apply_response_checksum_validation)

`
    public
            static        _apply_response_checksum_validation(mixed $value, array<string|int, mixed> &$args) : void`

##### Parameters

$value
: mixed$args
: array<string\|int, mixed>

#### \_apply\_use\_arn\_region()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Client.html\#method__apply_use_arn_region)

`
    public
            static        _apply_use_arn_region(mixed $value, array<string|int, mixed> &$args, HandlerList $list) : mixed`

##### Parameters

$value
: mixed$args
: array<string\|int, mixed>$list
: [HandlerList](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html)

#### \_default\_disable\_express\_session\_auth()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Client.html\#method__default_disable_express_session_auth)

`
    public
            static        _default_disable_express_session_auth(array<string|int, mixed> &$args) : mixed`

##### Parameters

$args
: array<string\|int, mixed>

#### \_default\_request\_checksum\_calculation()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Client.html\#method__default_request_checksum_calculation)

`
    public
            static        _default_request_checksum_calculation(array<string|int, mixed> $args) : string`

##### Parameters

$args
: array<string\|int, mixed>

##### Return values

string

#### \_default\_response\_checksum\_validation()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Client.html\#method__default_response_checksum_validation)

`
    public
            static        _default_response_checksum_validation(array<string|int, mixed> $args) : string`

##### Parameters

$args
: array<string\|int, mixed>

##### Return values

string

#### \_default\_s3\_express\_identity\_provider()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Client.html\#method__default_s3_express_identity_provider)

`
    public
            static        _default_s3_express_identity_provider(array<string|int, mixed> $args) : mixed`

##### Parameters

$args
: array<string\|int, mixed>

#### copy()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_copy)

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

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_copy\#tags)

see[S3ClientInterface::copy()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientInterface.html#method_copy)

#### copyAsync()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_copyAsync)

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

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_copyAsync\#tags)

see[S3ClientInterface::copyAsync()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientInterface.html#method_copyAsync)

#### createPresignedRequest()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Client.html\#method_createPresignedRequest)

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

[RequestInterface](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.RequestInterface.html)

#### deleteMatchingObjects()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_deleteMatchingObjects)

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

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_deleteMatchingObjects\#tags)

see[S3ClientInterface::deleteMatchingObjects()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientInterface.html#method_deleteMatchingObjects)

#### deleteMatchingObjectsAsync()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_deleteMatchingObjectsAsync)

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

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_deleteMatchingObjectsAsync\#tags)

see[S3ClientInterface::deleteMatchingObjectsAsync()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientInterface.html#method_deleteMatchingObjectsAsync)

#### determineBucketRegion()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_determineBucketRegion)

`
    public
                    determineBucketRegion(mixed $bucketName) : mixed`

##### Parameters

$bucketName
: mixed

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_determineBucketRegion\#tags)

see[S3ClientInterface::determineBucketRegion()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientInterface.html#method_determineBucketRegion)

#### determineBucketRegionAsync()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_determineBucketRegionAsync)

`
    public
                    determineBucketRegionAsync(string $bucketName) : PromiseInterface`

##### Parameters

$bucketName
: string

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_determineBucketRegionAsync\#tags)

see[S3ClientInterface::determineBucketRegionAsync()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientInterface.html#method_determineBucketRegionAsync)

##### Return values

[PromiseInterface](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.PromiseInterface.html)

#### doesBucketExist()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_doesBucketExist)

`
    public
                    doesBucketExist(mixed $bucket) : mixed`

##### Parameters

$bucket
: mixed

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_doesBucketExist\#tags)

see[S3ClientInterface::doesBucketExist()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientInterface.html#method_doesBucketExist)

#### doesBucketExistV2()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_doesBucketExistV2)

`
    public
                    doesBucketExistV2(mixed $bucket[, mixed $accept403 = false ]) : mixed`

##### Parameters

$bucket
: mixed$accept403
: mixed
= false

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_doesBucketExistV2\#tags)

see[S3ClientInterface::doesBucketExistV2()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientInterface.html#method_doesBucketExistV2)

#### doesObjectExist()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_doesObjectExist)

`
    public
                    doesObjectExist(mixed $bucket, mixed $key[, array<string|int, mixed> $options = [] ]) : mixed`

##### Parameters

$bucket
: mixed$key
: mixed$options
: array<string\|int, mixed>
= \[\]

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_doesObjectExist\#tags)

see[S3ClientInterface::doesObjectExist()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientInterface.html#method_doesObjectExist)

#### doesObjectExistV2()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_doesObjectExistV2)

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

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_doesObjectExistV2\#tags)

see[S3ClientInterface::doesObjectExistV2()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientInterface.html#method_doesObjectExistV2)

#### downloadBucket()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_downloadBucket)

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

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_downloadBucket\#tags)

see[S3ClientInterface::downloadBucket()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientInterface.html#method_downloadBucket)

#### downloadBucketAsync()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_downloadBucketAsync)

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

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_downloadBucketAsync\#tags)

see[S3ClientInterface::downloadBucketAsync()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientInterface.html#method_downloadBucketAsync)

#### encodeKey()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Client.html\#method_encodeKey)

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

#### factory()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClient.html\#method_factory)

`
    public
            static        factory([array<string|int, mixed> $config = [] ]) : static`

##### Parameters

$config
: array<string\|int, mixed>
= \[\]

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClient.html\#method_factory\#tags)

deprecated

##### Return values

static

#### getApi()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientTrait.html\#method_getApi)

`
    public
    abstract                getApi() : Service`

##### Return values

[Service](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.Api.Service.html)

#### getArguments()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Client.html\#method_getArguments)

Get an array of client constructor arguments used by the client.

`
    public
            static        getArguments() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getClientBuiltIns()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClient.html\#method_getClientBuiltIns)

Provides the set of built-in keys and values
used for endpoint resolution

`
    public
                    getClientBuiltIns() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getClientContextParams()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClient.html\#method_getClientContextParams)

Provides the set of service context parameter
key-value pairs used for endpoint resolution.

`
    public
                    getClientContextParams() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

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

#### getConfig()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClient.html\#method_getConfig)

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

#### getCredentials()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClient.html\#method_getCredentials)

Returns a promise that is fulfilled with an
{@see \\Aws\\Credentials\\CredentialsInterface} object.

`
    public
                    getCredentials() : PromiseInterface`

If you need the credentials synchronously, then call the wait() method
on the returned promise.

##### Return values

[PromiseInterface](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-GuzzleHttp.Promise.PromiseInterface.html)

#### getEndpoint()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClient.html\#method_getEndpoint)

Gets the default endpoint, or base URL, used by the client.

`
    public
                    getEndpoint() : UriInterface`

##### Return values

[UriInterface](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Psr.Http.Message.UriInterface.html)

#### getEndpointProvider()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClient.html\#method_getEndpointProvider)

`
    public
                    getEndpointProvider() : mixed`

#### getEndpointProviderArgs()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClient.html\#method_getEndpointProviderArgs)

Retrieves arguments to be used in endpoint resolution.

`
    public
                    getEndpointProviderArgs() : array<string|int, mixed>`

##### Return values

array<string\|int, mixed>

#### getHandlerList()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClient.html\#method_getHandlerList)

Get the handler list used to transfer commands.

`
    public
                    getHandlerList() : HandlerList`

This list can be modified to add middleware or to change the underlying
handler used to send HTTP requests.

##### Return values

[HandlerList](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.HandlerList.html)

#### getIterator()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientTrait.html\#method_getIterator)

`
    public
                    getIterator(mixed $name[, array<string|int, mixed> $args = [] ]) : mixed`

##### Parameters

$name
: mixed$args
: array<string\|int, mixed>
= \[\]

#### getObjectUrl()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Client.html\#method_getObjectUrl)

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

#### getPaginator()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientTrait.html\#method_getPaginator)

`
    public
                    getPaginator(mixed $name[, array<string|int, mixed> $args = [] ]) : mixed`

##### Parameters

$name
: mixed$args
: array<string\|int, mixed>
= \[\]

#### getRegion()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClient.html\#method_getRegion)

Get the region to which the client is configured to send requests.

`
    public
                    getRegion() : string`

##### Return values

string

#### getSignatureProvider()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClient.html\#method_getSignatureProvider)

Get the signature\_provider function of the client.

`
    public
        final            getSignatureProvider() : callable`

##### Return values

callable

#### getToken()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClient.html\#method_getToken)

`
    public
                    getToken() : mixed`

#### getWaiter()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientTrait.html\#method_getWaiter)

`
    public
                    getWaiter(mixed $name[, array<string|int, mixed> $args = [] ]) : mixed`

##### Parameters

$name
: mixed$args
: array<string\|int, mixed>
= \[\]

#### isBucketDnsCompatible()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Client.html\#method_isBucketDnsCompatible)

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

#### isDirectoryBucket()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Client.html\#method_isDirectoryBucket)

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

#### registerStreamWrapper()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_registerStreamWrapper)

`
    public
                    registerStreamWrapper() : mixed`

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_registerStreamWrapper\#tags)

see[S3ClientInterface::registerStreamWrapper()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientInterface.html#method_registerStreamWrapper)

#### registerStreamWrapperV2()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_registerStreamWrapperV2)

`
    public
                    registerStreamWrapperV2() : mixed`

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_registerStreamWrapperV2\#tags)

see[S3ClientInterface::registerStreamWrapperV2()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientInterface.html#method_registerStreamWrapperV2)

#### upload()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_upload)

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

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_upload\#tags)

see[S3ClientInterface::upload()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientInterface.html#method_upload)

#### uploadAsync()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_uploadAsync)

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

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_uploadAsync\#tags)

see[S3ClientInterface::uploadAsync()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientInterface.html#method_uploadAsync)

#### uploadDirectory()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_uploadDirectory)

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

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_uploadDirectory\#tags)

see[S3ClientInterface::uploadDirectory()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientInterface.html#method_uploadDirectory)

#### uploadDirectoryAsync()  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_uploadDirectoryAsync)

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

##### Tags  [header link](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html\#method_uploadDirectoryAsync\#tags)

see[S3ClientInterface::uploadDirectoryAsync()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientInterface.html#method_uploadDirectoryAsync)

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
  - [Methods](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Client.html#toc-methods)
- Methods
  - [\_\_call()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientTrait.html#method___call)
  - [\_\_construct()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Client.html#method___construct)
  - [\_\_sleep()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClient.html#method___sleep)
  - [\_apply\_request\_checksum\_calculation()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Client.html#method__apply_request_checksum_calculation)
  - [\_apply\_response\_checksum\_validation()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Client.html#method__apply_response_checksum_validation)
  - [\_apply\_use\_arn\_region()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Client.html#method__apply_use_arn_region)
  - [\_default\_disable\_express\_session\_auth()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Client.html#method__default_disable_express_session_auth)
  - [\_default\_request\_checksum\_calculation()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Client.html#method__default_request_checksum_calculation)
  - [\_default\_response\_checksum\_validation()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Client.html#method__default_response_checksum_validation)
  - [\_default\_s3\_express\_identity\_provider()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Client.html#method__default_s3_express_identity_provider)
  - [copy()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_copy)
  - [copyAsync()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_copyAsync)
  - [createPresignedRequest()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Client.html#method_createPresignedRequest)
  - [deleteMatchingObjects()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_deleteMatchingObjects)
  - [deleteMatchingObjectsAsync()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_deleteMatchingObjectsAsync)
  - [determineBucketRegion()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_determineBucketRegion)
  - [determineBucketRegionAsync()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_determineBucketRegionAsync)
  - [doesBucketExist()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_doesBucketExist)
  - [doesBucketExistV2()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_doesBucketExistV2)
  - [doesObjectExist()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_doesObjectExist)
  - [doesObjectExistV2()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_doesObjectExistV2)
  - [downloadBucket()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_downloadBucket)
  - [downloadBucketAsync()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_downloadBucketAsync)
  - [encodeKey()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Client.html#method_encodeKey)
  - [execute()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientTrait.html#method_execute)
  - [executeAsync()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientTrait.html#method_executeAsync)
  - [factory()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClient.html#method_factory)
  - [getApi()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientTrait.html#method_getApi)
  - [getArguments()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Client.html#method_getArguments)
  - [getClientBuiltIns()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClient.html#method_getClientBuiltIns)
  - [getClientContextParams()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClient.html#method_getClientContextParams)
  - [getCommand()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientTrait.html#method_getCommand)
  - [getConfig()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClient.html#method_getConfig)
  - [getCredentials()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClient.html#method_getCredentials)
  - [getEndpoint()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClient.html#method_getEndpoint)
  - [getEndpointProvider()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClient.html#method_getEndpointProvider)
  - [getEndpointProviderArgs()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClient.html#method_getEndpointProviderArgs)
  - [getHandlerList()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClient.html#method_getHandlerList)
  - [getIterator()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientTrait.html#method_getIterator)
  - [getObjectUrl()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Client.html#method_getObjectUrl)
  - [getPaginator()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientTrait.html#method_getPaginator)
  - [getRegion()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClient.html#method_getRegion)
  - [getSignatureProvider()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClient.html#method_getSignatureProvider)
  - [getToken()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClient.html#method_getToken)
  - [getWaiter()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientTrait.html#method_getWaiter)
  - [isBucketDnsCompatible()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Client.html#method_isBucketDnsCompatible)
  - [isDirectoryBucket()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Client.html#method_isDirectoryBucket)
  - [registerStreamWrapper()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_registerStreamWrapper)
  - [registerStreamWrapperV2()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_registerStreamWrapperV2)
  - [upload()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_upload)
  - [uploadAsync()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_uploadAsync)
  - [uploadDirectory()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_uploadDirectory)
  - [uploadDirectoryAsync()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3ClientTrait.html#method_uploadDirectoryAsync)
  - [waitUntil()](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.AwsClientTrait.html#method_waitUntil)

[Back To Top](https://docs.aws.amazon.com/aws-sdk-php/v3/api/class-Aws.S3.S3Client.html#top)
