# TagResource

Creates a new user-defined tag or updates an existing tag. Each tag is a label consisting of a key and value that is applied to your resource. Tags can help you organize, track costs for, and control access to your resources. You can add up to 50 AWS resource tags for each S3 resource.

###### Note

This operation is only supported for the following Amazon S3 resource:

- [General purpose buckets](../userguide/buckets-tagging.md)

- [Access Points for directory buckets](../userguide/access-points-db-tagging.md)

- [Access Points for general purpose buckets](../userguide/access-points-tagging.md)

- [Directory buckets](../userguide/directory-buckets-tagging.md)

- [S3 Storage Lens groups](../userguide/storage-lens-groups.md)

- [S3 Access Grants instances, registered locations, or grants](../userguide/access-grants-tagging.md).

Permissions

For general purpose buckets, access points for general purpose buckets, Storage Lens groups, and S3 Access Grants, you must have the `s3:TagResource` permission to use this operation.

Directory bucket permissions

For directory buckets, you must have the `s3express:TagResource` permission to use this operation. For more information about directory buckets policies and permissions, see [Identity and Access Management (IAM) for S3 Express One Zone](../userguide/s3-express-permissions.md) in the _Amazon S3 User Guide_.

HTTP Host header syntax

**Directory buckets** \- The HTTP Host header syntax is `s3express-control.region.amazonaws.com`.

For information about S3 Tagging errors, see [List of Amazon S3 Tagging error codes](errorresponses.md#S3TaggingErrorCodeList).

## Request Syntax

```nohighlight

POST /v20180820/tags/resourceArn+ HTTP/1.1
Host: s3-control.amazonaws.com
x-amz-account-id: AccountId
<?xml version="1.0" encoding="UTF-8"?>
<TagResourceRequest xmlns="http://awss3control.amazonaws.com/doc/2018-08-20/">
   <Tags>
      <Tag>
         <Key>string</Key>
         <Value>string</Value>
      </Tag>
   </Tags>
</TagResourceRequest>
```

## URI Request Parameters

The request uses the following URI parameters.

**[resourceArn](#API_control_TagResource_RequestSyntax)**

The Amazon Resource Name (ARN) of the S3 resource that you're applying tags to. The tagged resource can be a directory bucket, S3 Storage Lens group or S3 Access Grants instance, registered location, or grant.

Length Constraints: Maximum length of 1011.

Pattern: `arn:[^:]+:s3(express)?:[^:].*`

Required: Yes

**[x-amz-account-id](#API_control_TagResource_RequestSyntax)**

The AWS account ID that created the S3 resource that you're trying to add tags to or the requester's account ID.

Length Constraints: Maximum length of 64.

Pattern: `^\d{12}$`

Required: Yes

## Request Body

The request accepts the following data in XML format.

**[TagResourceRequest](#API_control_TagResource_RequestSyntax)**

Root level tag for the TagResourceRequest parameters.

Required: Yes

**[Tags](#API_control_TagResource_RequestSyntax)**

The AWS resource tags that you want to add to the specified S3 resource.

Type: Array of [Tag](api-control-tag.md) data types

Array Members: Minimum number of 0 items. Maximum number of 50 items.

Required: Yes

## Response Syntax

```

HTTP/1.1 204

```

## Response Elements

If the action is successful, the service sends back an HTTP 204 response with an empty HTTP body.

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/s3control-2018-08-20/tagresource.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/s3control-2018-08-20/tagresource.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3control-2018-08-20/tagresource.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/s3control-2018-08-20/tagresource.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3control-2018-08-20/tagresource.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/s3control-2018-08-20/tagresource.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/s3control-2018-08-20/tagresource.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/s3control-2018-08-20/tagresource.md)

- [AWS SDK for Python](../../../goto/boto3/s3control-2018-08-20/tagresource.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3control-2018-08-20/tagresource.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SubmitMultiRegionAccessPointRoutes

UntagResource

All content copied from https://docs.aws.amazon.com/.
