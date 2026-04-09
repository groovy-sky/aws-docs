# GetBucketLifecycleConfiguration

Returns the lifecycle configuration information set on the bucket. For information about lifecycle
configuration, see [Object Lifecycle Management](../dev/object-lifecycle-mgmt.md).

Bucket lifecycle configuration now supports specifying a lifecycle rule using an object key name
prefix, one or more object tags, object size, or any combination of these. Accordingly, this section
describes the latest API, which is compatible with the new functionality. The previous version of the
API supported filtering based only on an object key name prefix, which is supported for general purpose
buckets for backward compatibility. For the related API description, see [GetBucketLifecycle](api-getbucketlifecycle.md).

###### Note

Lifecyle configurations for directory buckets only support expiring objects and cancelling
multipart uploads. Expiring of versioned objects, transitions and tag filters are not
supported.

Permissions

- **General purpose bucket permissions** \- By default, all Amazon S3
resources are private, including buckets, objects, and related subresources (for example,
lifecycle configuration and website configuration). Only the resource owner (that is, the
AWS account that created it) can access the resource. The resource owner can optionally
grant access permissions to others by writing an access policy. For this operation, a user
must have the `s3:GetLifecycleConfiguration` permission.

For more information about permissions, see [Managing Access Permissions to Your\
Amazon S3 Resources](../userguide/s3-access-control.md).

- **Directory bucket permissions** \- You must have the
`s3express:GetLifecycleConfiguration` permission in an IAM identity-based policy
to use this operation. Cross-account access to this API operation isn't supported. The
resource owner can optionally grant access permissions to others by creating a role or user
for them as long as they are within the same account as the owner and resource.

For more information about directory bucket policies and permissions, see [Authorizing Regional endpoint APIs with IAM](../userguide/s3-express-security-iam.md) in the _Amazon S3 User_
_Guide_.

###### Note

**Directory buckets** \- For directory buckets, you must make requests for this API operation to the Regional endpoint. These endpoints support path-style requests in the format `https://s3express-control.region-code.amazonaws.com/bucket-name
                          `. Virtual-hosted-style requests aren't supported.
For more information about endpoints in Availability Zones, see [Regional and Zonal endpoints for directory buckets in Availability Zones](../userguide/endpoint-directory-buckets-az.md) in the
_Amazon S3 User Guide_. For more information about endpoints in Local Zones, see [Concepts for directory buckets in Local Zones](../userguide/s3-lzs-for-directory-buckets.md) in the
_Amazon S3 User Guide_.

HTTP Host header syntax

**Directory buckets** \- The HTTP Host header syntax is
`s3express-control.region.amazonaws.com`.

`GetBucketLifecycleConfiguration` has the following special error:

- Error code: `NoSuchLifecycleConfiguration`

- Description: The lifecycle configuration does not exist.

- HTTP Status Code: 404 Not Found

- SOAP Fault Code Prefix: Client

The following operations are related to `GetBucketLifecycleConfiguration`:

- [GetBucketLifecycle](api-getbucketlifecycle.md)

- [PutBucketLifecycle](api-putbucketlifecycle.md)

- [DeleteBucketLifecycle](api-deletebucketlifecycle.md)

###### Important

You must URL encode any signed header values that contain spaces. For example, if your header value is `my  file.txt`, containing two spaces after `my`, you must URL encode this value to `my%20%20file.txt`.

## Request Syntax

```nohighlight

GET /?lifecycle HTTP/1.1
Host: Bucket.s3.amazonaws.com
x-amz-expected-bucket-owner: ExpectedBucketOwner

```

## URI Request Parameters

The request uses the following URI parameters.

**[Bucket](#API_GetBucketLifecycleConfiguration_RequestSyntax)**

The name of the bucket for which to get the lifecycle information.

Required: Yes

**[x-amz-expected-bucket-owner](#API_GetBucketLifecycleConfiguration_RequestSyntax)**

The account ID of the expected bucket owner. If the account ID that you provide does not match the actual owner of the bucket, the request fails with the HTTP status code `403 Forbidden` (access denied).

###### Note

This parameter applies to general purpose buckets only. It is not supported for directory bucket
lifecycle configurations.

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
x-amz-transition-default-minimum-object-size: TransitionDefaultMinimumObjectSize
<?xml version="1.0" encoding="UTF-8"?>
<LifecycleConfiguration>
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
            <Tag>
               <Key>string</Key>
               <Value>string</Value>
            </Tag>
            ...
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
      <NoncurrentVersionTransition>
         <NewerNoncurrentVersions>integer</NewerNoncurrentVersions>
         <NoncurrentDays>integer</NoncurrentDays>
         <StorageClass>string</StorageClass>
      </NoncurrentVersionTransition>
      ...
      <Prefix>string</Prefix>
      <Status>string</Status>
      <Transition>
         <Date>timestamp</Date>
         <Days>integer</Days>
         <StorageClass>string</StorageClass>
      </Transition>
      ...
   </Rule>
   ...
</LifecycleConfiguration>
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The response returns the following HTTP headers.

**[x-amz-transition-default-minimum-object-size](#API_GetBucketLifecycleConfiguration_ResponseSyntax)**

Indicates which default minimum object size behavior is applied to the lifecycle
configuration.

###### Note

This parameter applies to general purpose buckets only. It isn't supported for directory bucket
lifecycle configurations.

- `all_storage_classes_128K` \- Objects smaller than 128 KB will not transition to
any storage class by default.

- `varies_by_storage_class` \- Objects smaller than 128 KB will transition to Glacier
Flexible Retrieval or Glacier Deep Archive storage classes. By default, all other storage classes
will prevent transitions smaller than 128 KB.

To customize the minimum object size for any transition you can add a filter that specifies a custom
`ObjectSizeGreaterThan` or `ObjectSizeLessThan` in the body of your transition
rule. Custom filters always take precedence over the default transition behavior.

Valid Values: `varies_by_storage_class | all_storage_classes_128K`

The following data is returned in XML format by the service.

**[LifecycleConfiguration](#API_GetBucketLifecycleConfiguration_ResponseSyntax)**

Root level tag for the LifecycleConfiguration parameters.

Required: Yes

**[Rule](#API_GetBucketLifecycleConfiguration_ResponseSyntax)**

Container for a lifecycle rule.

Type: Array of [LifecycleRule](api-lifecyclerule.md) data types

## Examples

### Example 1: Get lifecycle configuration - general purpose bucket

This example illustrates how to use `GetBucketLifecycleConfiguration` to retrieve the
lifecycle configuration for a general purpose bucket:

```

GET /?lifecycle HTTP/1.1
Host: amzn-s3-demo-bucket.s3.<Region>.amazonaws.com
x-amz-date: Thu, 15 Nov 2012 00:17:21 GMT
Authorization: signatureValue

```

### Sample Response

This example shows the response from the preceeding `GetBucketLifecycleConfiguration`
request:

```

HTTP/1.1 200 OK
x-amz-id-2: ITnGT1y4RyTmXa3rPi4hklTXouTf0hccUjo0iCPjz6FnfIutBj3M7fPGlWO2SEWp
x-amz-request-id: 51991C342C575321
Date: Thu, 15 Nov 2012 00:17:23 GMT
Server: AmazonS3
Content-Length: 358

<?xml version="1.0" encoding="UTF-8"?>
<LifecycleConfiguration xmlns="http://s3.amazonaws.com/doc/2006-03-01/">
   <Rule>
      <ID>Archive and then delete rule</ID>
      <Prefix>projectdocs/</Prefix>
      <Status>Enabled</Status>
      <Transition>
         <Days>30</Days>
         <StorageClass>STANDARD_IA</StorageClass>
      </Transition>
      <Transition>
         <Days>365</Days>
         <StorageClass>GLACIER</StorageClass>
      </Transition>
      <Expiration>
         <Days>3650</Days>
      </Expiration>
   </Rule>
</LifecycleConfiguration>

```

### Example 2: Get lifecycle configuration - directory bucket

This example illustrates how to use `GetBucketLifecycleConfiguration` to retrieve the
lifecycle configuration for a directory bucket:

```

GET /?lifecycle HTTP/1.1
Host:s3express-control.us-west-2.amazonaws.com

```

### Sample Response

This example shows the response from the preceeding `GetBucketLifecycleConfiguration`
request:

```

<?xml version="1.0" encoding="UTF-8"?>
<LifecycleConfiguration xmlns="http://s3.amazonaws.com/doc/2006-03-01/">
   <Rule>
      <ID>Lifecycle expiration rule</ID>
      <Filter>
         <And>
            <Prefix>myprefix/</Prefix>
            <ObjectSizeGreaterThan>500</ObjectSizeGreaterThan>
            <ObjectSizeLessThan>64000</ObjectSizeLessThan>
         </And>
      </Filter>
      <Status>Enabled</Status>
      <Expiration>
         <Days>7</Days>
      </Expiration>
   </Rule>
   <Rule>
      <ID>MPU Rule </ID>
      <Filter>
         <Prefix>another_prefix </Prefix>
      </Filter>
      <Status>Enabled</Status>
      <AbortIncompleteMultipartUpload>
         <DaysAfterInitiation>3</DaysAfterInitiation>
      </AbortIncompleteMultipartUpload>
   </Rule>
</LifecycleConfiguration>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/s3-2006-03-01/getbucketlifecycleconfiguration.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/s3-2006-03-01/getbucketlifecycleconfiguration.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3-2006-03-01/getbucketlifecycleconfiguration.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/s3-2006-03-01/getbucketlifecycleconfiguration.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3-2006-03-01/getbucketlifecycleconfiguration.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/s3-2006-03-01/getbucketlifecycleconfiguration.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/s3-2006-03-01/getbucketlifecycleconfiguration.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/s3-2006-03-01/getbucketlifecycleconfiguration.md)

- [AWS SDK for Python](../../../goto/boto3/s3-2006-03-01/getbucketlifecycleconfiguration.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3-2006-03-01/getbucketlifecycleconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetBucketLifecycle

GetBucketLocation

All content copied from https://docs.aws.amazon.com/.
