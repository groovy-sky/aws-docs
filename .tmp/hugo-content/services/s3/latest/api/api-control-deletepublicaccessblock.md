# DeletePublicAccessBlock

###### Note

This operation is not supported by directory buckets.

Removes the `PublicAccessBlock` configuration for an AWS account. This
operation might be restricted when the account is managed by organization-level Block
Public Access policies. You’ll get an Access Denied (403) error when the account is managed
by organization-level Block Public Access policies. Organization-level policies override
account-level settings, preventing direct account-level modifications. For more
information, see [Using Amazon S3 block\
public access](../dev/access-control-block-public-access.md).

Related actions include:

- [GetPublicAccessBlock](api-control-getpublicaccessblock.md)

- [PutPublicAccessBlock](api-control-putpublicaccessblock.md)

## Request Syntax

```nohighlight

DELETE /v20180820/configuration/publicAccessBlock HTTP/1.1
Host: s3-control.amazonaws.com
x-amz-account-id: AccountId

```

## URI Request Parameters

The request uses the following URI parameters.

**[x-amz-account-id](#API_control_DeletePublicAccessBlock_RequestSyntax)**

The account ID for the AWS account whose `PublicAccessBlock` configuration
you want to remove.

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

- [AWS Command Line Interface V2](../../../goto/cli2/s3control-2018-08-20/deletepublicaccessblock.md)

- [AWS SDK for .NET V4](../../../../reference/goto/dotnetsdkv4/s3control-2018-08-20/deletepublicaccessblock.md)

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/s3control-2018-08-20/deletepublicaccessblock.md)

- [AWS SDK for Go v2](../../../../reference/goto/sdkforgov2/s3control-2018-08-20/deletepublicaccessblock.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/s3control-2018-08-20/deletepublicaccessblock.md)

- [AWS SDK for JavaScript V3](../../../../reference/goto/sdkforjavascriptv3/s3control-2018-08-20/deletepublicaccessblock.md)

- [AWS SDK for Kotlin](../../../../reference/goto/sdkforkotlin/s3control-2018-08-20/deletepublicaccessblock.md)

- [AWS SDK for PHP V3](../../../../reference/goto/sdkforphpv3/s3control-2018-08-20/deletepublicaccessblock.md)

- [AWS SDK for Python](../../../goto/boto3/s3control-2018-08-20/deletepublicaccessblock.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/s3control-2018-08-20/deletepublicaccessblock.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeleteMultiRegionAccessPoint

DeleteStorageLensConfiguration

All content copied from https://docs.aws.amazon.com/.
