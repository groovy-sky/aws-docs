# CreateBucket

###### Note

This action creates an Amazon S3 bucket. To create an Amazon S3 on Outposts bucket, see [`CreateBucket`](https://docs.aws.amazon.com/AmazonS3/latest/API/API_control_CreateBucket.html).

Creates a new S3 bucket. To create a bucket, you must set up Amazon S3 and have a valid AWS Access Key
ID to authenticate requests. Anonymous requests are never allowed to create buckets. By creating the
bucket, you become the bucket owner.

There are two types of buckets: general purpose buckets and directory buckets. For more information about
these bucket types, see [Creating, configuring, and working with Amazon S3\
buckets](../userguide/creating-buckets-s3.md) in the _Amazon S3 User Guide_.

General purpose buckets exist in a global namespace, which means that each bucket name must be unique
across all AWS accounts in all the AWS Regions within a partition. A partition is a grouping of
Regions. AWS currently has four partitions: `aws` (Standard Regions), `aws-cn`
(China Regions), `aws-us-gov` (AWS GovCloud (US)), and `aws-eusc`
(European Sovereign Cloud). When you create a general purpose bucket, you can choose to create a bucket in
the shared global namespace or you can choose to create a bucket in your account regional namespace.
Your account regional namespace is a subdivision of the global namespace that only your account can
create buckets in. For more information on account regional namespaces, see [Namespaces for general purpose buckets](../userguide/gpbucketnamespaces.md).

###### Note

- **General purpose buckets** \- If you send your
`CreateBucket` request to the `s3.amazonaws.com` global endpoint, the
request goes to the `us-east-1` Region. So the signature calculations in Signature
Version 4 must use `us-east-1` as the Region, even if the location constraint in the
request specifies another Region where the bucket is to be created. If you create a bucket in a
Region other than US East (N. Virginia), your application must be able to handle 307 redirect. For
more information, see [Virtual hosting of buckets](https://docs.aws.amazon.com/AmazonS3/latest/dev/VirtualHosting.html) in the _Amazon S3 User Guide_.

- **Directory buckets** \- For directory buckets, you must make requests for this API operation to the Regional endpoint. These endpoints support path-style requests in the format `https://s3express-control.region-code.amazonaws.com/bucket-name
                 `. Virtual-hosted-style requests aren't supported.
For more information about endpoints in Availability Zones, see [Regional and Zonal endpoints for directory buckets in Availability Zones](../userguide/endpoint-directory-buckets-az.md) in the
_Amazon S3 User Guide_. For more information about endpoints in Local Zones, see [Concepts for directory buckets in Local Zones](../userguide/s3-lzs-for-directory-buckets.md) in the
_Amazon S3 User Guide_.

Permissions

- **General purpose bucket permissions** \- In addition to the
`s3:CreateBucket` permission, the following permissions are required in a policy
when your `CreateBucket` request includes specific headers:

- **Access control lists (ACLs)** \- In your
`CreateBucket` request, if you specify an access control list (ACL) and set
it to `public-read`, `public-read-write`,
`authenticated-read`, or if you explicitly specify any other custom ACLs,
both `s3:CreateBucket` and `s3:PutBucketAcl` permissions are
required. In your `CreateBucket` request, if you set the ACL to
`private`, or if you don't specify any ACLs, only the
`s3:CreateBucket` permission is required.

- **Object Lock** \- In your
`CreateBucket` request, if you set
`x-amz-bucket-object-lock-enabled` to true, the
`s3:PutBucketObjectLockConfiguration` and `s3:PutBucketVersioning`
permissions are required.

- **S3 Object Ownership** \- If your
`CreateBucket` request includes the `x-amz-object-ownership`
header, then the `s3:PutBucketOwnershipControls` permission is required.

###### Important

To set an ACL on a bucket as part of a `CreateBucket` request, you must
explicitly set S3 Object Ownership for the bucket to a different value than the default,
`BucketOwnerEnforced`. Additionally, if your desired bucket ACL grants
public access, you must first create the bucket (without the bucket ACL) and then
explicitly disable Block Public Access on the bucket before using
`PutBucketAcl` to set the ACL. If you try to create a bucket with a public
ACL, the request will fail.

For the majority of modern use cases in S3, we recommend that you keep all Block
Public Access settings enabled and keep ACLs disabled. If you would like to share data
with users outside of your account, you can use bucket policies as needed. For more
information, see [Controlling ownership of\
objects and disabling ACLs for your bucket](../userguide/about-object-ownership.md) and [Blocking\
public access to your Amazon S3 storage](../userguide/access-control-block-public-access.md) in the
_Amazon S3 User Guide_.

- **S3 Block Public Access** \- If your specific use
case requires granting public access to your S3 resources, you can disable Block Public
Access. Specifically, you can create a new bucket with Block Public Access enabled, then
separately call the [`DeletePublicAccessBlock`](api-deletepublicaccessblock.md) API. To use this operation, you must have the
`s3:PutBucketPublicAccessBlock` permission. For more information about S3
Block Public Access, see [Blocking public\
access to your Amazon S3 storage](../userguide/access-control-block-public-access.md) in the _Amazon S3 User Guide_.

- **Directory bucket permissions** \- You must have the
`s3express:CreateBucket` permission in an IAM identity-based policy instead of a bucket policy.
Cross-account access to this API operation isn't supported. This operation can only be performed by the AWS account that owns the resource. For more information about directory bucket policies and permissions, see [AWS Identity and Access Management (IAM) for S3 Express One Zone](../userguide/s3-express-security-iam.md) in the _Amazon S3 User Guide_.

###### Important

The permissions for ACLs, Object Lock, S3 Object Ownership, and S3 Block Public Access
are not supported for directory buckets. For directory buckets, all Block Public Access
settings are enabled at the bucket level and S3 Object Ownership is set to Bucket owner
enforced (ACLs disabled). These settings can't be modified.

For more information about permissions for creating and working with directory buckets,
see [Directory buckets](../userguide/directory-buckets-overview.md)
in the _Amazon S3 User Guide_. For more information about supported S3
features for directory buckets, see [Features of\
S3 Express One Zone](https://docs.aws.amazon.com/AmazonS3/latest/userguide/s3-express-one-zone.html#s3-express-features) in the _Amazon S3 User Guide_.

HTTP Host header syntax

**Directory buckets** \- The HTTP Host header syntax is `s3express-control.region-code.amazonaws.com`.

The following operations are related to `CreateBucket`:

- [PutObject](api-putobject.md)

- [DeleteBucket](https://docs.aws.amazon.com/AmazonS3/latest/API/API_DeleteBucket.html)

###### Important

You must URL encode any signed header values that contain spaces. For example, if your header value is `my  file.txt`, containing two spaces after `my`, you must URL encode this value to `my%20%20file.txt`.

## Request Syntax

```nohighlight

PUT / HTTP/1.1
Host: Bucket.s3.amazonaws.com
x-amz-acl: ACL
x-amz-grant-full-control: GrantFullControl
x-amz-grant-read: GrantRead
x-amz-grant-read-acp: GrantReadACP
x-amz-grant-write: GrantWrite
x-amz-grant-write-acp: GrantWriteACP
x-amz-bucket-object-lock-enabled: ObjectLockEnabledForBucket
x-amz-object-ownership: ObjectOwnership
x-amz-bucket-namespace: BucketNamespace
<?xml version="1.0" encoding="UTF-8"?>
<CreateBucketConfiguration xmlns="http://s3.amazonaws.com/doc/2006-03-01/">
   <LocationConstraint>string</LocationConstraint>
   <Location>
      <Name>string</Name>
      <Type>string</Type>
   </Location>
   <Bucket>
      <DataRedundancy>string</DataRedundancy>
      <Type>string</Type>
   </Bucket>
   <Tags>
      <Tag>
         <Key>string</Key>
         <Value>string</Value>
      </Tag>
   </Tags>
</CreateBucketConfiguration>
```

## URI Request Parameters

The request uses the following URI parameters.

**[Bucket](#API_CreateBucket_RequestSyntax)**

The name of the bucket to create.

**General purpose buckets** \- For information about bucket naming
restrictions, see [Bucket naming\
rules](../userguide/bucketnamingrules.md) in the _Amazon S3 User Guide_.

**Directory buckets** \- When you use this operation with a directory bucket, you must use path-style requests in the format `https://s3express-control.region-code.amazonaws.com/bucket-name
                  `. Virtual-hosted-style requests aren't supported. Directory bucket names must be unique in the chosen Zone (Availability Zone or Local Zone). Bucket names must also follow the format `
                     bucket-base-name--zone-id--x-s3` (for example, `
                     DOC-EXAMPLE-BUCKET--usw2-az1--x-s3`). For information about bucket naming restrictions, see [Directory bucket naming rules](../userguide/directory-bucket-naming-rules.md) in the _Amazon S3 User Guide_

Required: Yes

**[x-amz-acl](#API_CreateBucket_RequestSyntax)**

The canned ACL to apply to the bucket.

###### Note

This functionality is not supported for directory buckets.

Valid Values: `private | public-read | public-read-write | authenticated-read`

**[x-amz-bucket-namespace](#API_CreateBucket_RequestSyntax)**

Specifies the namespace where you want to create your general purpose bucket. When you create a
general purpose bucket, you can choose to create a bucket in the shared global namespace or you can choose to
create a bucket in your account regional namespace. Your account regional namespace is a subdivision of
the global namespace that only your account can create buckets in. For more information on bucket
namespaces, see [Namespaces for general purpose buckets](../userguide/gpbucketnamespaces.md).

General purpose buckets in your account regional namespace must follow a specific naming convention. These
buckets consist of a bucket name prefix that you create, and a suffix that contains your 12-digit AWS
Account ID, the AWS Region code, and ends with `-an`. Bucket names must follow the format
`bucket-name-prefix-accountId-region-an` (for example,
`amzn-s3-demo-bucket-111122223333-us-west-2-an`). For information about bucket naming
restrictions, see [Account regional namespace naming rules](../userguide/bucketnamingrules.md#account-regional-naming-rules)
in the _Amazon S3 User Guide_.

###### Note

This functionality is not supported for directory buckets.

Valid Values: `account-regional | global`

**[x-amz-bucket-object-lock-enabled](#API_CreateBucket_RequestSyntax)**

Specifies whether you want S3 Object Lock to be enabled for the new bucket.

###### Note

This functionality is not supported for directory buckets.

**[x-amz-grant-full-control](#API_CreateBucket_RequestSyntax)**

Allows grantee the read, write, read ACP, and write ACP permissions on the bucket.

###### Note

This functionality is not supported for directory buckets.

**[x-amz-grant-read](#API_CreateBucket_RequestSyntax)**

Allows grantee to list the objects in the bucket.

###### Note

This functionality is not supported for directory buckets.

**[x-amz-grant-read-acp](#API_CreateBucket_RequestSyntax)**

Allows grantee to read the bucket ACL.

###### Note

This functionality is not supported for directory buckets.

**[x-amz-grant-write](#API_CreateBucket_RequestSyntax)**

Allows grantee to create new objects in the bucket.

For the bucket and object owners of existing objects, also allows deletions and overwrites of those
objects.

###### Note

This functionality is not supported for directory buckets.

**[x-amz-grant-write-acp](#API_CreateBucket_RequestSyntax)**

Allows grantee to write the ACL for the applicable bucket.

###### Note

This functionality is not supported for directory buckets.

**[x-amz-object-ownership](#API_CreateBucket_RequestSyntax)**

The container element for object ownership for a bucket's ownership controls.

`BucketOwnerPreferred` \- Objects uploaded to the bucket change ownership to the bucket
owner if the objects are uploaded with the `bucket-owner-full-control` canned ACL.

`ObjectWriter` \- The uploading account will own the object if the object is uploaded with
the `bucket-owner-full-control` canned ACL.

`BucketOwnerEnforced` \- Access control lists (ACLs) are disabled and no longer affect
permissions. The bucket owner automatically owns and has full control over every object in the bucket.
The bucket only accepts PUT requests that don't specify an ACL or specify bucket owner full control ACLs
(such as the predefined `bucket-owner-full-control` canned ACL or a custom ACL in XML format
that grants the same permissions).

By default, `ObjectOwnership` is set to `BucketOwnerEnforced` and ACLs are
disabled. We recommend keeping ACLs disabled, except in uncommon use cases where you must control access
for each object individually. For more information about S3 Object Ownership, see [Controlling\
ownership of objects and disabling ACLs for your bucket](../userguide/about-object-ownership.md) in the
_Amazon S3 User Guide_.

###### Note

This functionality is not supported for directory buckets. Directory buckets use the bucket owner enforced setting for S3 Object Ownership.

Valid Values: `BucketOwnerPreferred | ObjectWriter | BucketOwnerEnforced`

## Request Body

The request accepts the following data in XML format.

**[CreateBucketConfiguration](#API_CreateBucket_RequestSyntax)**

Root level tag for the CreateBucketConfiguration parameters.

Required: Yes

**[Bucket](#API_CreateBucket_RequestSyntax)**

Specifies the information about the bucket that will be created.

###### Note

This functionality is only supported by directory buckets.

Type: [BucketInfo](https://docs.aws.amazon.com/AmazonS3/latest/API/API_BucketInfo.html) data type

Required: No

**[Location](#API_CreateBucket_RequestSyntax)**

Specifies the location where the bucket will be created.

**Directory buckets** \- The location type is Availability Zone or Local Zone. To
use the Local Zone location type, your account must be enabled
for Local Zones. Otherwise, you get an HTTP `403 Forbidden` error with the error code
`AccessDenied`. To learn more,
see [Enable\
accounts for Local Zones](../userguide/opt-in-directory-bucket-lz.md) in the _Amazon S3 User Guide_.

###### Note

This functionality is only supported by directory buckets.

Type: [LocationInfo](https://docs.aws.amazon.com/AmazonS3/latest/API/API_LocationInfo.html) data type

Required: No

**[LocationConstraint](#API_CreateBucket_RequestSyntax)**

Specifies the Region where the bucket will be created. You might choose a Region to optimize
latency, minimize costs, or address regulatory requirements. For example, if you reside in Europe, you
will probably find it advantageous to create buckets in the Europe (Ireland) Region.

If you don't specify a Region, the bucket is created in the US East (N. Virginia) Region (us-east-1)
by default. Configurations using the value `EU` will create a bucket in
`eu-west-1`.

For a list of the valid values for all of the AWS Regions, see [Regions and Endpoints](https://docs.aws.amazon.com/general/latest/gr/rande.html#s3_region).

###### Note

This functionality is not supported for directory buckets.

Type: String

Valid Values: `af-south-1 | ap-east-1 | ap-northeast-1 | ap-northeast-2 | ap-northeast-3 | ap-south-1 | ap-south-2 | ap-southeast-1 | ap-southeast-2 | ap-southeast-3 | ap-southeast-4 | ap-southeast-5 | ca-central-1 | cn-north-1 | cn-northwest-1 | EU | eu-central-1 | eu-central-2 | eu-north-1 | eu-south-1 | eu-south-2 | eu-west-1 | eu-west-2 | eu-west-3 | il-central-1 | me-central-1 | me-south-1 | sa-east-1 | us-east-2 | us-gov-east-1 | us-gov-west-1 | us-west-1 | us-west-2`

Required: No

**[Tags](#API_CreateBucket_RequestSyntax)**

An array of tags that you can apply to the bucket that you're creating. Tags are key-value pairs of metadata used to categorize and organize your buckets, track costs, and control access.

You must have the `s3:TagResource` permission to create a general
purpose bucket with tags or the `s3express:TagResource` permission to create a directory bucket with tags.

When creating buckets with tags, note that tag-based conditions using `aws:ResourceTag` and `s3:BucketTag` condition keys are applicable only after ABAC is enabled on the bucket. To learn more, see [Enabling ABAC in general purpose buckets](https://docs.aws.amazon.com/AmazonS3/latest/userguide/buckets-tagging-enable-abac.html).

Type: Array of [Tag](https://docs.aws.amazon.com/AmazonS3/latest/API/API_Tag.html) data types

Required: No

## Response Syntax

```nohighlight

HTTP/1.1 200
Location: Location
x-amz-bucket-arn: BucketArn

```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The response returns the following HTTP headers.

**[Location](#API_CreateBucket_ResponseSyntax)**

A forward slash followed by the name of the bucket.

**[x-amz-bucket-arn](#API_CreateBucket_ResponseSyntax)**

The Amazon Resource Name (ARN) of the S3 bucket. ARNs uniquely identify AWS resources across all
of AWS.

###### Note

This parameter is only supported for S3 directory buckets. For more information, see [Using tags with\
directory buckets](../userguide/directory-buckets-tagging.md).

Length Constraints: Minimum length of 1. Maximum length of 128.

Pattern: `arn:[^:]+:(s3|s3express):.*`

## Errors

**BucketAlreadyExists**

The requested bucket name is not available. The bucket namespace is shared by all users of the
system. Select a different name and try again.

HTTP Status Code: 409

**BucketAlreadyOwnedByYou**

The bucket you tried to create already exists, and you own it. Amazon S3 returns this error in all AWS
Regions except in the North Virginia Region. For legacy compatibility, if you re-create an existing
bucket that you already own in the North Virginia Region, Amazon S3 returns 200 OK and resets the bucket
access control lists (ACLs).

HTTP Status Code: 409

## Examples

### Sample Request for general purpose buckets

This request creates a bucket named `amzn-s3-demo-bucket`.

```

            PUT / HTTP/1.1
            Host: amzn-s3-demo-bucket.s3.<Region>.amazonaws.com
            Content-Length: 0
            Date: Wed, 01 Mar  2006 12:00:00 GMT
            Authorization: authorization string

```

### Sample Response for general purpose buckets

This example illustrates one usage of CreateBucket.

```

            HTTP/1.1 200 OK
            x-amz-id-2: YgIPIfBiKa2bj0KMg95r/0zo3emzU4dzsD4rcKCHQUAdQkf3ShJTOOpXUueF6QKo
            x-amz-request-id: 236A8905248E5A01
            Date: Wed, 01 Mar  2006 12:00:00 GMT

            Location: /amzn-s3-demo-bucket
            Content-Length: 0
            Connection: close
            Server: AmazonS3

```

### Sample Request for general purpose buckets: Setting the Region of a bucket

The following request sets the Region for the bucket to Europe.

```

            PUT / HTTP/1.1
            Host: amzn-s3-demo-bucket.s3.amazonaws.com
            Date: Wed, 12 Oct 2009 17:50:00 GMT
            Authorization: authorization string
            Content-Type: text/plain
            Content-Length: 124

            <CreateBucketConfiguration xmlns="http://s3.amazonaws.com/doc/2006-03-01/">
             <LocationConstraint>EU</LocationConstraint>
            </CreateBucketConfiguration >

```

### Sample Request for general purpose buckets: Creating a bucket and applying the ObjectWriter setting for S3 Object Ownership.

This request creates a bucket and applies the `ObjectWriter` setting for Object
Ownership.

```

            PUT / HTTP/1.1
            Host: amzn-s3-demo-bucket.s3.<Region>.amazonaws.com
            Content-Length: 0
            x-amz-object-ownership: ObjectWriter
            Date: Tue, 30 Nov  2021 12:00:00 GMT
            Authorization: authorization string

```

### Sample Response for general purpose buckets

This example illustrates one usage of CreateBucket.

```

            HTTP/1.1 200 OK
            x-amz-id-2: YgIPIfBiKa2bj0KMg95r/0zo3emzU4dzsD4rcKCHQUAdQkf3ShJTOOpXUueF6QKo
            x-amz-request-id: 236A8905248E5A01
            Date: Tue, 30 Nov  2021 12:00:00 GMT

            Location: /amzn-s3-demo-bucket
            Content-Length: 0
            Connection: close
            Server: AmazonS3

```

### Sample Request for general purpose buckets: Creating a bucket and configuring access permissions explicitly

This request creates a bucket named `amzn-s3-demo-bucket` and grants WRITE permission
to the AWS account identified by canonical id.

```

            PUT HTTP/1.1
            Host: amzn-s3-demo-bucket.s3.<Region>.amazonaws.com
            x-amz-date: Sat, 07 Apr 2012 00:54:40 GMT
            Authorization: authorization string
            x-amz-grant-write: id="79a59df900b949e55d96a1e698fbacedfd6e09d98eacf8f8d5218e7cd47ef2be", id="52b113e7a2f25102679df27bb0746532b8776c37c9fc7ca4ac94015369746b9f"

```

### Sample Response for general purpose buckets

This example illustrates one usage of CreateBucket.

```

           HTTP/1.1 200 OK

```

### Sample Request for general purpose buckets: Creating a bucket and configuring access permission using a canned ACL

This request creates a bucket named `amzn-s3-demo-bucket` and sets the ACL to
private.

```

            PUT / HTTP/1.1
            Host: amzn-s3-demo-bucket.s3.<Region>.amazonaws.com
            Content-Length: 0
            x-amz-acl: private
            Date: Wed, 01 Mar  2006 12:00:00 GMT
            Authorization: authorization string

```

### Sample Response for general purpose buckets

This example illustrates one usage of CreateBucket.

```

            HTTP/1.1 200 OK
            x-amz-id-2: YgIPIfBiKa2bj0KMg95r/0zo3emzU4dzsD4rcKCHQUAdQkf3ShJTOOpXUueF6QKo
            x-amz-request-id: 236A8905248E5A01
            Date: Wed, 01 Mar  2006 12:00:00 GMT

            Location: /amzn-s3-demo-bucket
            Content-Length: 0
            Connection: close
            Server: AmazonS3

```

### Sample Request for general purpose buckets: Creating a bucket in your account regional namespace

This request creates the general purpose bucket `amzn-s3-demo-bucket-012345678910-us-west-2-an`
in the account regional namespace for AWS Account 012345678910 in the US West (Oregon) Region
`(us-west-2)` Region.

```

            PUT / HTTP/1.1
            Host: amzn-s3-demo-bucket-012345678910-us-west-2-an.s3.us-west-2.amazonaws.com
            Content-Length: 0
            x-amz-bucket-namespace: account-regional
            Date: Wed, 01 Mar  2006 12:00:00 GMT
            Authorization: authorization string
            <CreateBucketConfiguration xmlns="http://s3.amazonaws.com/doc/2006-03-01/">
             <LocationConstraint>us-west-2</LocationConstraint>
            </CreateBucketConfiguration >

```

### Sample Response for general purpose buckets

This example illustrates one usage of CreateBucket.

```

            HTTP/1.1 200 OK
            x-amz-id-2: YgIPIfBiKa2bj0KMg95r/0zo3emzU4dzsD4rcKCHQUAdQkf3ShJTOOpXUueF6QKo
            x-amz-request-id: 236A8905248E5A01
            Date: Wed, 01 Mar  2006 12:00:00 GMT

            Location: /amzn-s3-demo-bucket-012345678910-us-west-2-an
            Content-Length: 0
            Connection: close
            Server: AmazonS3

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/s3-2006-03-01/CreateBucket)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/s3-2006-03-01/CreateBucket)

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/s3-2006-03-01/CreateBucket)

- [AWS SDK for Go v2](https://docs.aws.amazon.com/goto/SdkForGoV2/s3-2006-03-01/CreateBucket)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/s3-2006-03-01/CreateBucket)

- [AWS SDK for JavaScript V3](https://docs.aws.amazon.com/goto/SdkForJavaScriptV3/s3-2006-03-01/CreateBucket)

- [AWS SDK for Kotlin](https://docs.aws.amazon.com/goto/SdkForKotlin/s3-2006-03-01/CreateBucket)

- [AWS SDK for PHP V3](https://docs.aws.amazon.com/goto/SdkForPHPV3/s3-2006-03-01/CreateBucket)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/s3-2006-03-01/CreateBucket)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/s3-2006-03-01/CreateBucket)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

CopyObject

CreateBucketMetadataConfiguration
