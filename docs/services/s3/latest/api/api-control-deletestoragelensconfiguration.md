# DeleteStorageLensConfiguration

###### Note

This operation is not supported by directory buckets.

Deletes the Amazon S3 Storage Lens configuration. For more information about S3 Storage Lens, see [Assessing your storage\
activity and usage with Amazon S3 Storage Lens](../dev/storage-lens.md) in the
_Amazon S3 User Guide_.

###### Note

To use this action, you must have permission to perform the
`s3:DeleteStorageLensConfiguration` action. For more information, see
[Setting permissions to\
use Amazon S3 Storage Lens](../dev/storage-lens-iam-permissions.md) in the _Amazon S3 User Guide_.

## Request Syntax

```nohighlight

DELETE /v20180820/storagelens/storagelensid HTTP/1.1
Host: s3-control.amazonaws.com
x-amz-account-id: AccountId

```

## URI Request Parameters

The request uses the following URI parameters.

**[storagelensid](#API_control_DeleteStorageLensConfiguration_RequestSyntax)**

The ID of the S3 Storage Lens configuration.

Length Constraints: Minimum length of 1. Maximum length of 64.

Pattern: `[a-zA-Z0-9\-\_\.]+`

Required: Yes

**[x-amz-account-id](#API_control_DeleteStorageLensConfiguration_RequestSyntax)**

The account ID of the requester.

Length Constraints: Maximum length of 64.

Pattern: `^\d{12}$`

Required: Yes

## Request Body

The request does not have a request body.

## Response Syntax

```

HTTP/1.1 200

```

## Response Elements

If the action is successful, the service sends back an HTTP 200 response with an empty HTTP body.

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../goto/cli2/s3control-2018-08-20/deletestoragelensconfiguration.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/s3control-2018-08-20/deletestoragelensconfiguration.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3control-2018-08-20/deletestoragelensconfiguration.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/s3control-2018-08-20/deletestoragelensconfiguration.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3control-2018-08-20/deletestoragelensconfiguration.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/s3control-2018-08-20/deletestoragelensconfiguration.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/s3control-2018-08-20/deletestoragelensconfiguration.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/s3control-2018-08-20/deletestoragelensconfiguration.md)

- [AWS SDK for Python](../../../goto/boto3/s3control-2018-08-20/deletestoragelensconfiguration.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3control-2018-08-20/deletestoragelensconfiguration.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeletePublicAccessBlock

DeleteStorageLensConfigurationTagging

All content copied from https://docs.aws.amazon.com/.
