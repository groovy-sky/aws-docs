# GetBucketTagging

###### Note

This action gets an Amazon S3 on Outposts bucket's tags. To get an S3 bucket tags, see
[GetBucketTagging](api-getbuckettagging.md) in the _Amazon S3 API Reference_.

Returns the tag set associated with the Outposts bucket. For more information, see
[Using\
Amazon S3 on Outposts](../userguide/s3onoutposts.md) in the _Amazon S3 User Guide_.

To use this action, you must have permission to perform the
`GetBucketTagging` action. By default, the bucket owner has this permission
and can grant this permission to others.

`GetBucketTagging` has the following special error:

- Error code: `NoSuchTagSetError`

- Description: There is no tag set associated with the bucket.

All Amazon S3 on Outposts REST API requests for this action require an additional parameter of `x-amz-outpost-id` to be passed with the request. In addition, you must use an S3 on Outposts endpoint hostname prefix instead of `s3-control`. For an example of the request syntax for Amazon S3 on Outposts that uses the S3 on Outposts endpoint hostname prefix and the `x-amz-outpost-id` derived by using the access point ARN, see the [Examples](api-control-getbuckettagging.md#API_control_GetBucketTagging_Examples) section.

The following actions are related to `GetBucketTagging`:

- [PutBucketTagging](api-control-putbuckettagging.md)

- [DeleteBucketTagging](api-control-deletebuckettagging.md)

## Request Syntax

```nohighlight

GET /v20180820/bucket/name/tagging HTTP/1.1
Host: Bucket.s3-control.amazonaws.com
x-amz-account-id: AccountId

```

## URI Request Parameters

The request uses the following URI parameters.

**[name](#API_control_GetBucketTagging_RequestSyntax)**

Specifies the bucket.

For using this parameter with Amazon S3 on Outposts with the REST API, you must specify the name and the x-amz-outpost-id as well.

For using this parameter with S3 on Outposts with the AWS SDK and CLI, you must specify the ARN of the bucket accessed in the format `arn:aws:s3-outposts:<Region>:<account-id>:outpost/<outpost-id>/bucket/<my-bucket-name>`. For example, to access the bucket `reports` through Outpost `my-outpost` owned by account `123456789012` in Region `us-west-2`, use the URL encoding of `arn:aws:s3-outposts:us-west-2:123456789012:outpost/my-outpost/bucket/reports`. The value must be URL encoded.

Length Constraints: Minimum length of 3. Maximum length of 255.

Required: Yes

**[x-amz-account-id](#API_control_GetBucketTagging_RequestSyntax)**

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
<GetBucketTaggingResult>
   <TagSet>
      <S3Tag>
         <Key>string</Key>
         <Value>string</Value>
      </S3Tag>
   </TagSet>
</GetBucketTaggingResult>
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in XML format by the service.

**[GetBucketTaggingResult](#API_control_GetBucketTagging_ResponseSyntax)**

Root level tag for the GetBucketTaggingResult parameters.

Required: Yes

**[TagSet](#API_control_GetBucketTagging_ResponseSyntax)**

The tags set of the Outposts bucket.

Type: Array of [S3Tag](api-control-s3tag.md) data types

## Examples

### Amazon S3 on Outposts request example for getting a tag set for an Outposts bucket

The following request gets the tag set of the specified Outposts bucket
`example-outpost-bucket`.

```HTTP

            GET /v20180820/bucket/example-outpost-bucket/tagging HTTP/1.1
            Host: s3-outposts.<Region>.amazonaws.com
            x-amz-date: Wed, 28 Oct 2020 22:32:00 GMT
            x-amz-account-id: example-account-id
            x-amz-outpost-id: op-01ac5d28a6a232904
            Authorization: authorization string

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/s3control-2018-08-20/getbuckettagging.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/s3control-2018-08-20/getbuckettagging.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3control-2018-08-20/getbuckettagging.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/s3control-2018-08-20/getbuckettagging.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3control-2018-08-20/getbuckettagging.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/s3control-2018-08-20/getbuckettagging.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/s3control-2018-08-20/getbuckettagging.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/s3control-2018-08-20/getbuckettagging.md)

- [AWS SDK for Python](../../../goto/boto3/s3control-2018-08-20/getbuckettagging.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3control-2018-08-20/getbuckettagging.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetBucketReplication

GetBucketVersioning

All content copied from https://docs.aws.amazon.com/.
