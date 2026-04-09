# GetBucketVersioning

###### Note

This operation returns the versioning state
for
S3 on Outposts
buckets
only. To return the versioning state for an S3 bucket, see [GetBucketVersioning](api-getbucketversioning.md) in the _Amazon S3 API Reference_.

Returns the versioning state for an S3 on Outposts bucket. With
S3
Versioning,
you can save multiple distinct copies of your
objects
and recover from unintended user actions and application failures.

If you've never set versioning on your bucket, it has no versioning state. In that case,
the `GetBucketVersioning` request does not return a versioning state
value.

For more information about versioning, see [Versioning](../userguide/versioning.md) in the _Amazon S3_
_User Guide_.

All Amazon S3 on Outposts REST API requests for this action require an additional parameter of `x-amz-outpost-id` to be passed with the request. In addition, you must use an S3 on Outposts endpoint hostname prefix instead of `s3-control`. For an example of the request syntax for Amazon S3 on Outposts that uses the S3 on Outposts endpoint hostname prefix and the `x-amz-outpost-id` derived by using the access point ARN, see the [Examples](api-control-getbucketversioning.md#API_control_GetBucketVersioning_Examples) section.

The following operations are related to `GetBucketVersioning` for
S3 on Outposts.

- [PutBucketVersioning](api-control-putbucketversioning.md)

- [PutBucketLifecycleConfiguration](api-control-putbucketlifecycleconfiguration.md)

- [GetBucketLifecycleConfiguration](api-control-getbucketlifecycleconfiguration.md)

## Request Syntax

```nohighlight

GET /v20180820/bucket/name/versioning HTTP/1.1
Host: Bucket.s3-control.amazonaws.com
x-amz-account-id: AccountId

```

## URI Request Parameters

The request uses the following URI parameters.

**[name](#API_control_GetBucketVersioning_RequestSyntax)**

The S3 on Outposts bucket to return the versioning state for.

Length Constraints: Minimum length of 3. Maximum length of 255.

Required: Yes

**[x-amz-account-id](#API_control_GetBucketVersioning_RequestSyntax)**

The AWS account ID of the S3 on Outposts bucket.

Length Constraints: Maximum length of 64.

Pattern: `^\d{12}$`

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```nohighlight

HTTP/1.1 200
<?xml version="1.0" encoding="UTF-8"?>
<GetBucketVersioningResult>
   <Status>string</Status>
   <MfaDelete>string</MfaDelete>
</GetBucketVersioningResult>
```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response.

The following data is returned in XML format by the service.

**[GetBucketVersioningResult](#API_control_GetBucketVersioning_ResponseSyntax)**

Root level tag for the GetBucketVersioningResult parameters.

Required: Yes

**[MFADelete](#API_control_GetBucketVersioning_ResponseSyntax)**

Specifies whether MFA delete is enabled in the bucket versioning configuration. This
element is returned only if the bucket has been configured with MFA delete. If MFA delete
has never been configured for the bucket, this element is not returned.

Type: String

Valid Values: `Enabled | Disabled`

**[Status](#API_control_GetBucketVersioning_ResponseSyntax)**

The versioning state of the S3 on Outposts bucket.

Type: String

Valid Values: `Enabled | Suspended`

## Examples

### Sample GetBucketVersioning request on an S3 on Outposts bucket

This request returns the versioning state for an S3 on Outposts bucket
that's
named `example-outpost-bucket`.

```HTTP

            GET /v20180820/bucket/example-outpost-bucket/?versioning HTTP/1.1
            Host:s3-outposts.region-code.amazonaws.com
            x-amz-account-id: example-account-id
            x-amz-outpost-id: op-01ac5d28a6a232904
            x-amz-date: Wed, 25 May  2022 00:14:21 GMT
            Authorization: signatureValue

```

### Sample GetBucketVersioning response on a versioning-enabled S3 on Outposts bucket

If you enabled versioning on a bucket, the response is:

```

     <VersioningConfiguration xmlns="http://awss3control.amazonaws.com/doc/2018-08-20/">
        <Status>Enabled</Status>
     </VersioningConfiguration>

```

### Sample GetBucketVersioning response on a versioning-suspended bucket

If you suspended versioning on a bucket, the response is:

```

     <VersioningConfiguration xmlns="http://awss3control.amazonaws.com/doc/2018-08-20/">
        <Status>Suspended</Status>
     </VersioningConfiguration>

```

### Sample GetBucketVersioning response if you have never enabled versioning.

If you have
never enabled versioning on a bucket, the response is:

```

     <VersioningConfiguration xmlns="http://awss3control.amazonaws.com/doc/2018-08-20/">

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/s3control-2018-08-20/getbucketversioning.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/s3control-2018-08-20/getbucketversioning.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3control-2018-08-20/getbucketversioning.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/s3control-2018-08-20/getbucketversioning.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3control-2018-08-20/getbucketversioning.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/s3control-2018-08-20/getbucketversioning.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/s3control-2018-08-20/getbucketversioning.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/s3control-2018-08-20/getbucketversioning.md)

- [AWS SDK for Python](../../../goto/boto3/s3control-2018-08-20/getbucketversioning.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3control-2018-08-20/getbucketversioning.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GetBucketTagging

GetDataAccess

All content copied from https://docs.aws.amazon.com/.
