# PutBucketVersioning

###### Note

This operation sets the versioning state
for
S3 on Outposts
buckets
only. To set the versioning state for an S3 bucket, see [PutBucketVersioning](api-putbucketversioning.md) in the _Amazon S3 API Reference_.

Sets the versioning state for an S3 on Outposts bucket. With
S3
Versioning,
you can save multiple distinct copies of your
objects
and recover from unintended user actions and application failures.

You can set the versioning state to one of the following:

- **Enabled** \- Enables versioning for the objects in
the bucket. All objects added to the bucket receive a unique version ID.

- **Suspended** \- Suspends versioning for the objects
in the bucket. All objects added to the bucket receive the version ID
`null`.

If you've never set versioning on your bucket, it has no versioning state. In that case,
a [GetBucketVersioning](api-control-getbucketversioning.md) request does not return a versioning state value.

When you enable S3 Versioning, for each object in your bucket, you have a current
version and zero or more noncurrent versions. You can configure your bucket S3 Lifecycle
rules to expire noncurrent versions after a specified time period. For more information,
see [Creating and managing\
a lifecycle configuration for your S3 on Outposts bucket](../userguide/s3outpostslifecyclemanaging.md) in the _Amazon S3_
_User Guide_.

If you have an object expiration lifecycle configuration in your non-versioned bucket
and you want to maintain the same permanent delete behavior when you enable versioning, you
must add a noncurrent expiration policy. The noncurrent expiration lifecycle configuration
will manage the deletes of the noncurrent object versions in the version-enabled bucket.
For more information, see [Versioning](../userguide/versioning.md) in the _Amazon S3_
_User Guide_.

All Amazon S3 on Outposts REST API requests for this action require an additional parameter of `x-amz-outpost-id` to be passed with the request. In addition, you must use an S3 on Outposts endpoint hostname prefix instead of `s3-control`. For an example of the request syntax for Amazon S3 on Outposts that uses the S3 on Outposts endpoint hostname prefix and the `x-amz-outpost-id` derived by using the access point ARN, see the [Examples](api-control-putbucketversioning.md#API_control_PutBucketVersioning_Examples) section.

The following operations are related to `PutBucketVersioning` for
S3 on Outposts.

- [GetBucketVersioning](api-control-getbucketversioning.md)

- [PutBucketLifecycleConfiguration](api-control-putbucketlifecycleconfiguration.md)

- [GetBucketLifecycleConfiguration](api-control-getbucketlifecycleconfiguration.md)

## Request Syntax

```nohighlight

PUT /v20180820/bucket/name/versioning HTTP/1.1
Host: Bucket.s3-control.amazonaws.com
x-amz-account-id: AccountId
x-amz-mfa: MFA
<?xml version="1.0" encoding="UTF-8"?>
<VersioningConfiguration xmlns="http://awss3control.amazonaws.com/doc/2018-08-20/">
   <MfaDelete>string</MfaDelete>
   <Status>string</Status>
</VersioningConfiguration>
```

## URI Request Parameters

The request uses the following URI parameters.

**[name](#API_control_PutBucketVersioning_RequestSyntax)**

The S3 on Outposts bucket to set the versioning state for.

Length Constraints: Minimum length of 3. Maximum length of 255.

Required: Yes

**[x-amz-account-id](#API_control_PutBucketVersioning_RequestSyntax)**

The AWS account ID of the S3 on Outposts bucket.

Length Constraints: Maximum length of 64.

Pattern: `^\d{12}$`

Required: Yes

**[x-amz-mfa](#API_control_PutBucketVersioning_RequestSyntax)**

The concatenation of the authentication device's serial number, a space, and the value
that is displayed on your authentication device.

## Request Body

The request accepts the following data in XML format.

**[VersioningConfiguration](#API_control_PutBucketVersioning_RequestSyntax)**

Root level tag for the VersioningConfiguration parameters.

Required: Yes

**[MFADelete](#API_control_PutBucketVersioning_RequestSyntax)**

Specifies whether MFA delete is enabled or disabled in the bucket versioning
configuration for the S3 on Outposts bucket.

Type: String

Valid Values: `Enabled | Disabled`

Required: No

**[Status](#API_control_PutBucketVersioning_RequestSyntax)**

Sets the versioning state of the S3 on Outposts bucket.

Type: String

Valid Values: `Enabled | Suspended`

Required: No

## Response Syntax

```

HTTP/1.1 200

```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

## Examples

### Sample PutBucketVersioning request on an Amazon S3 on Outposts bucket

This request sets the versioning state on an S3 on Outposts bucket
that's
named `example-outpost-bucket`.

```HTTP

            PUT /v20180820/bucket/example-outpost-bucket/?versioning HTTP/1.1
            Host:s3-outposts.region-code.amazonaws.com
            x-amz-account-id: example-account-id
            x-amz-outpost-id: op-01ac5d28a6a232904
            Content-Length: 0
            Date: Wed, 25 May  2022 12:00:00 GMT
            Content-MD5: q6yJDlIkcBaGGfb3QLY69A==
            Authorization: authorization string
            Content-Length: 214

           <VersioningConfiguration xmlns="http://s3.amazonaws.com/doc/2006-03-01/">
            <Status>Enabled</Status>
           </VersioningConfiguration>

```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/s3control-2018-08-20/putbucketversioning.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/s3control-2018-08-20/putbucketversioning.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3control-2018-08-20/putbucketversioning.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/s3control-2018-08-20/putbucketversioning.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3control-2018-08-20/putbucketversioning.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/s3control-2018-08-20/putbucketversioning.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/s3control-2018-08-20/putbucketversioning.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/s3control-2018-08-20/putbucketversioning.md)

- [AWS SDK for Python](../../../goto/boto3/s3control-2018-08-20/putbucketversioning.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3control-2018-08-20/putbucketversioning.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PutBucketTagging

PutJobTagging

All content copied from https://docs.aws.amazon.com/.
