# GetBucketLifecycleConfiguration

###### Note

This action gets an Amazon S3 on Outposts bucket's lifecycle configuration. To get an S3
bucket's lifecycle configuration, see [GetBucketLifecycleConfiguration](api-getbucketlifecycleconfiguration.md) in the _Amazon S3 API Reference_.

Returns the lifecycle configuration information set on the Outposts bucket. For more
information, see [Using Amazon S3 on Outposts](../userguide/s3onoutposts.md) and for
information about lifecycle configuration, see [Object Lifecycle\
Management](../dev/object-lifecycle-mgmt.md) in _Amazon S3 User Guide_.

To use this action, you must have permission to perform the
`s3-outposts:GetLifecycleConfiguration` action. The Outposts bucket owner
has this permission, by default. The bucket owner can grant this permission to others. For
more information about permissions, see [Permissions Related to Bucket Subresource Operations](../userguide/using-with-s3-actions.md#using-with-s3-actions-related-to-bucket-subresources) and [Managing\
Access Permissions to Your Amazon S3 Resources](../userguide/s3-access-control.md).

All Amazon S3 on Outposts REST API requests for this action require an additional parameter of `x-amz-outpost-id` to be passed with the request. In addition, you must use an S3 on Outposts endpoint hostname prefix instead of `s3-control`. For an example of the request syntax for Amazon S3 on Outposts that uses the S3 on Outposts endpoint hostname prefix and the `x-amz-outpost-id` derived by using the access point ARN, see the [Examples](api-control-getbucketlifecycleconfiguration.md#API_control_GetBucketLifecycleConfiguration_Examples) section.

`GetBucketLifecycleConfiguration` has the following special error:

- Error code: `NoSuchLifecycleConfiguration`

- Description: The lifecycle configuration does not exist.

- HTTP Status Code: 404 Not Found

- SOAP Fault Code Prefix: Client

The following actions are related to
`GetBucketLifecycleConfiguration`:

- [PutBucketLifecycleConfiguration](api-control-putbucketlifecycleconfiguration.md)

- [DeleteBucketLifecycleConfiguration](api-control-deletebucketlifecycleconfiguration.md)

## Request Syntax

```nohighlight

GET /v20180820/bucket/name/lifecycleconfiguration HTTP/1.1
Host: Bucket.s3-control.amazonaws.com
x-amz-account-id: AccountId

```

## URI Request Parameters

The request uses the following URI parameters.

**[name](#API_control_GetBucketLifecycleConfiguration_RequestSyntax)**

The Amazon Resource Name (ARN) of the bucket.

For using this parameter with Amazon S3 on Outposts with the REST API, you must specify the name and the x-amz-outpost-id as well.

For using this parameter with S3 on Outposts with the AWS SDK and CLI, you must specify the ARN of the bucket accessed in the format `arn:aws:s3-outposts:<Region>:<account-id>:outpost/<outpost-id>/bucket/<my-bucket-name>`. For example, to access the bucket `reports` through Outpost `my-outpost` owned by account `123456789012` in Region `us-west-2`, use the URL encoding of `arn:aws:s3-outposts:us-west-2:123456789012:outpost/my-outpost/bucket/reports`. The value must be URL encoded.

Length Constraints: Minimum length of 3. Maximum length of 255.

Required: Yes

**[x-amz-account-id](#API_control_GetBucketLifecycleConfiguration_RequestSyntax)**

The AWS account ID of the Outposts bucket.

Length Constraints: Maximum length of 64.

Pattern: `^\d{12}$`

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
<?xml version="1.0" encoding="UTF-8"?>
<GetBucketLifecycleConfigurationResult>
   <Rules>
      <Rule>
         <AbortIncompleteMultipartUpload>
            <DaysAfterInitiation>integer</DaysAfterInitiation>
         </AbortIncompleteMultipartUpload>
         <Expiration>
            <Date>timestamp</Date>
            <Days>integer</Days>
            <ExpiredObjectDeleteMarker>boolean</ExpiredObjectDeleteMarker>
         </Expiration>
         <Filter>
            <And>
               <ObjectSizeGreaterThan>long</ObjectSizeGreaterThan>
               <ObjectSizeLessThan>long</ObjectSizeLessThan>
               <Prefix>string</Prefix>
               <Tags>
                  <S3Tag>
                     <Key>string</Key>
                     <Value>string</Value>
                  </S3Tag>
               </Tags>
            </And>
            <ObjectSizeGreaterThan>long</ObjectSizeGreaterThan>
            <ObjectSizeLessThan>long</ObjectSizeLessThan>
            <Prefix>string</Prefix>
            <Tag>
               <Key>string</Key>
               <Value>string</Value>
            </Tag>
         </Filter>
         <ID>string</ID>
         <NoncurrentVersionExpiration>
            <NewerNoncurrentVersions>integer</NewerNoncurrentVersions>
            <NoncurrentDays>integer</NoncurrentDays>
         </NoncurrentVersionExpiration>
         <NoncurrentVersionTransitions>
            <NoncurrentVersionTransition>
               <NoncurrentDays>integer</NoncurrentDays>
               <StorageClass>string</StorageClass>
            </NoncurrentVersionTransition>
         </NoncurrentVersionTransitions>
         <Status>string</Status>
         <Transitions>
            <Transition>
               <Date>timestamp</Date>
               <Days>integer</Days>
               <StorageClass>string</StorageClass>
            </Transition>
         </Transitions>
      </Rule>
   </Rules>
</GetBucketLifecycleConfigurationResult>
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in XML format by the service.

**[GetBucketLifecycleConfigurationResult](#API_control_GetBucketLifecycleConfiguration_ResponseSyntax)**

Root level tag for the GetBucketLifecycleConfigurationResult parameters.

Required: Yes

**[Rules](#API_control_GetBucketLifecycleConfiguration_ResponseSyntax)**

Container for the lifecycle rule of the Outposts bucket.

Type: Array of [LifecycleRule](api-control-lifecyclerule.md) data types

## Examples

### Sample request to get the lifecycle configuration of the Amazon S3 on Outposts bucket

The following example shows how to get the lifecycle configuration of the
Outposts bucket.

```HTTP

            GET /v20180820/bucket/example-outpost-bucket/lifecycleconfiguration HTTP/1.1
            Host: s3-outposts.<Region>.amazonaws.com
            x-amz-account-id: example-account-id
            x-amz-outpost-id: op-01ac5d28a6a232904
            x-amz-date: Thu, 15 Nov 2012 00:17:21 GMT
            Authorization: signatureValue

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/s3control-2018-08-20/getbucketlifecycleconfiguration.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/s3control-2018-08-20/getbucketlifecycleconfiguration.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3control-2018-08-20/getbucketlifecycleconfiguration.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/s3control-2018-08-20/getbucketlifecycleconfiguration.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3control-2018-08-20/getbucketlifecycleconfiguration.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/s3control-2018-08-20/getbucketlifecycleconfiguration.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/s3control-2018-08-20/getbucketlifecycleconfiguration.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/s3control-2018-08-20/getbucketlifecycleconfiguration.md)

- [AWS SDK for Python](../../../goto/boto3/s3control-2018-08-20/getbucketlifecycleconfiguration.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3control-2018-08-20/getbucketlifecycleconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetBucket

GetBucketPolicy

All content copied from https://docs.aws.amazon.com/.
