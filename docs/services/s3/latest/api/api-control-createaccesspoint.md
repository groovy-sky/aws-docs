# CreateAccessPoint

Creates an access point and associates it to a specified bucket. For more information, see
[Managing\
access to shared datasets with access points](../userguide/access-points.md) or [Managing access to\
shared datasets in directory buckets with access points](../userguide/access-points-directory-buckets.md) in the
_Amazon S3 User Guide_.

To create an access point and attach it to a volume on an Amazon FSx file system, see [CreateAndAttachS3AccessPoint](../../../../reference/fsx/latest/apireference/api-createandattachs3accesspoint.md) in the _Amazon FSx API_
_Reference_.

###### Note

S3 on Outposts only supports VPC-style access points.

For more information, see [Accessing Amazon S3 on Outposts using\
virtual private cloud (VPC) only access points](../userguide/s3onoutposts.md) in the
_Amazon S3 User Guide_.

All Amazon S3 on Outposts REST API requests for this action require an additional parameter of `x-amz-outpost-id` to be passed with the request. In addition, you must use an S3 on Outposts endpoint hostname prefix instead of `s3-control`. For an example of the request syntax for Amazon S3 on Outposts that uses the S3 on Outposts endpoint hostname prefix and the `x-amz-outpost-id` derived by using the access point ARN, see the [Examples](api-control-createaccesspoint.md#API_control_CreateAccessPoint_Examples) section.

The following actions are related to `CreateAccessPoint`:

- [GetAccessPoint](api-control-getaccesspoint.md)

- [DeleteAccessPoint](api-control-deleteaccesspoint.md)

- [ListAccessPoints](api-control-listaccesspoints.md)

- [ListAccessPointsForDirectoryBuckets](api-control-listaccesspointsfordirectorybuckets.md)

## Request Syntax

```nohighlight

PUT /v20180820/accesspoint/name HTTP/1.1
Host: s3-control.amazonaws.com
x-amz-account-id: AccountId
<?xml version="1.0" encoding="UTF-8"?>
<CreateAccessPointRequest xmlns="http://awss3control.amazonaws.com/doc/2018-08-20/">
   <Bucket>string</Bucket>
   <VpcConfiguration>
      <VpcId>string</VpcId>
   </VpcConfiguration>
   <PublicAccessBlockConfiguration>
      <BlockPublicAcls>boolean</BlockPublicAcls>
      <BlockPublicPolicy>boolean</BlockPublicPolicy>
      <IgnorePublicAcls>boolean</IgnorePublicAcls>
      <RestrictPublicBuckets>boolean</RestrictPublicBuckets>
   </PublicAccessBlockConfiguration>
   <BucketAccountId>string</BucketAccountId>
   <Scope>
      <Permissions>
         <Permission>string</Permission>
      </Permissions>
      <Prefixes>
         <Prefix>string</Prefix>
      </Prefixes>
   </Scope>
   <Tags>
      <Tag>
         <Key>string</Key>
         <Value>string</Value>
      </Tag>
   </Tags>
</CreateAccessPointRequest>
```

## URI Request Parameters

The request uses the following URI parameters.

**[name](#API_control_CreateAccessPoint_RequestSyntax)**

The name you want to assign to this access point.

For directory buckets, the access point name must consist of a base name that you provide and
suffix that includes the `ZoneID` (AWS Availability Zone or Local Zone) of your bucket location,
followed by `--xa-s3`. For more information, see [Managing access to shared datasets in directory buckets with\
access points](../userguide/access-points-directory-buckets.md) in the _Amazon S3 User Guide_.

Length Constraints: Minimum length of 3. Maximum length of 255.

Required: Yes

**[x-amz-account-id](#API_control_CreateAccessPoint_RequestSyntax)**

The AWS account ID for the account that owns the specified access point.

Length Constraints: Maximum length of 64.

Pattern: `^\d{12}$`

Required: Yes

## Request Body

The request accepts the following data in XML format.

**[CreateAccessPointRequest](#API_control_CreateAccessPoint_RequestSyntax)**

Root level tag for the CreateAccessPointRequest parameters.

Required: Yes

**[Bucket](#API_control_CreateAccessPoint_RequestSyntax)**

The name of the bucket that you want to associate this access point with.

For using this parameter with Amazon S3 on Outposts with the REST API, you must specify the name and the x-amz-outpost-id as well.

For using this parameter with S3 on Outposts with the AWS SDK and CLI, you must specify the ARN of the bucket accessed in the format `arn:aws:s3-outposts:<Region>:<account-id>:outpost/<outpost-id>/bucket/<my-bucket-name>`. For example, to access the bucket `reports` through Outpost `my-outpost` owned by account `123456789012` in Region `us-west-2`, use the URL encoding of `arn:aws:s3-outposts:us-west-2:123456789012:outpost/my-outpost/bucket/reports`. The value must be URL encoded.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 255.

Required: Yes

**[BucketAccountId](#API_control_CreateAccessPoint_RequestSyntax)**

The AWS account ID associated with the S3 bucket associated with this access point.

For same account access point when your bucket and access point belong to the same account owner, the `BucketAccountId` is not required.
For cross-account access point when your bucket and access point are not in the same account, the `BucketAccountId` is required.

Type: String

Length Constraints: Maximum length of 64.

Pattern: `^\d{12}$`

Required: No

**[PublicAccessBlockConfiguration](#API_control_CreateAccessPoint_RequestSyntax)**

The `PublicAccessBlock` configuration that you want to apply to the access point.

Type: [PublicAccessBlockConfiguration](api-control-publicaccessblockconfiguration.md) data type

Required: No

**[Scope](#API_control_CreateAccessPoint_RequestSyntax)**

For directory buckets, you can filter access control to specific prefixes, API
operations, or a combination of both. For more information, see [Managing access to shared datasets in directory buckets with\
access points](../userguide/access-points-directory-buckets.md) in the _Amazon S3 User Guide_.

###### Note

Scope is only supported for access points attached to directory buckets.

Type: [Scope](api-control-scope.md) data type

Required: No

**[Tags](#API_control_CreateAccessPoint_RequestSyntax)**

An array of tags that you can apply to an access point. Tags are key-value pairs of metadata used to control access to your access points. For more information about tags, see [Using tags with Amazon S3](../userguide/tagging.md). For information about tagging access points, see [Using tags for attribute-based access control (ABAC)](../userguide/tagging.md#using-tags-for-abac).

Type: Array of [Tag](api-control-tag.md) data types

Array Members: Minimum number of 0 items. Maximum number of 50 items.

Required: No

**[VpcConfiguration](#API_control_CreateAccessPoint_RequestSyntax)**

If you include this field, Amazon S3 restricts access to this access point to requests from the
specified virtual private cloud (VPC).

###### Note

This is required for creating an access point for Amazon S3 on Outposts buckets.

Type: [VpcConfiguration](api-control-vpcconfiguration.md) data type

Required: No

## Response Syntax

```nohighlight

HTTP/1.1 200
<?xml version="1.0" encoding="UTF-8"?>
<CreateAccessPointResult>
   <AccessPointArn>string</AccessPointArn>
   <Alias>string</Alias>
</CreateAccessPointResult>
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in XML format by the service.

**[CreateAccessPointResult](#API_control_CreateAccessPoint_ResponseSyntax)**

Root level tag for the CreateAccessPointResult parameters.

Required: Yes

**[AccessPointArn](#API_control_CreateAccessPoint_ResponseSyntax)**

The ARN of the access point.

###### Note

This is only supported by Amazon S3 on Outposts.

Type: String

Length Constraints: Minimum length of 4. Maximum length of 128.

**[Alias](#API_control_CreateAccessPoint_ResponseSyntax)**

The name or alias of the access point.

Type: String

Length Constraints: Maximum length of 63.

Pattern: `^[0-9a-z\\-]{63}`

## Examples

### Sample request for creating an access point for an Amazon S3 on Outposts bucket

This request creates an access point for S3 on Outposts bucket.

```HTTP

            PUT /v20180820/accesspoint/example-access-point HTTP/1.1
            Host:s3-outposts.<Region>.amazonaws.com
            x-amz-account-id: example-account-id
            x-amz-outpost-id: op-01ac5d28a6a232904
            <?xml version="1.0" encoding="UTF-8"?>
               <CreateAccessPointRequest xmlns="http://awss3control.amazonaws.com/doc/2018-08-20/">
                  <Bucket>example-outpost-bucket </Bucket>
               </CreateAccessPointRequest>

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/s3control-2018-08-20/createaccesspoint.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/s3control-2018-08-20/createaccesspoint.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3control-2018-08-20/createaccesspoint.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/s3control-2018-08-20/createaccesspoint.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3control-2018-08-20/createaccesspoint.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/s3control-2018-08-20/createaccesspoint.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/s3control-2018-08-20/createaccesspoint.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/s3control-2018-08-20/createaccesspoint.md)

- [AWS SDK for Python](../../../goto/boto3/s3control-2018-08-20/createaccesspoint.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3control-2018-08-20/createaccesspoint.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CreateAccessGrantsLocation

CreateAccessPointForObjectLambda

All content copied from https://docs.aws.amazon.com/.
