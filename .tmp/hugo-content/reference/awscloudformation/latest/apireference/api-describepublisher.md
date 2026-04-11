# DescribePublisher

Returns information about a CloudFormation extension publisher.

If you don't supply a `PublisherId`, and you have registered as an extension
publisher, `DescribePublisher` returns information about your own publisher
account.

For more information about registering as a publisher, see:

- [RegisterPublisher](api-registerpublisher.md)

- [Publishing extensions\
to make them available for public use](../../../../services/cloudformation-cli/latest/userguide/publish-extension.md) in the
_AWS CloudFormation Command Line Interface (CLI) User Guide_

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**PublisherId**

The ID of the extension publisher.

If you don't supply a `PublisherId`, and you have registered as an extension
publisher, `DescribePublisher` returns information about your own publisher
account.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 40.

Pattern: `[0-9a-zA-Z]{12,40}`

Required: No

## Response Elements

The following elements are returned by the service.

**IdentityProvider**

The type of account used as the identity provider when registering this publisher with
CloudFormation.

Type: String

Valid Values: `AWS_Marketplace | GitHub | Bitbucket`

**PublisherId**

The ID of the extension publisher.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 40.

Pattern: `[0-9a-zA-Z]{12,40}`

**PublisherProfile**

The URL to the publisher's profile with the identity provider.

Type: String

Length Constraints: Maximum length of 1024.

Pattern: `(http:|https:)+[^\s]+[\w]`

**PublisherStatus**

Whether the publisher is verified. Currently, all registered publishers are
verified.

Type: String

Valid Values: `VERIFIED | UNVERIFIED`

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**CFNRegistry**

An error occurred during a CloudFormation registry operation.

**Message**

A message with details about the error that occurred.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/cloudformation-2010-05-15/describepublisher.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/cloudformation-2010-05-15/describepublisher.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/cloudformation-2010-05-15/describepublisher.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/cloudformation-2010-05-15/describepublisher.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/cloudformation-2010-05-15/describepublisher.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/cloudformation-2010-05-15/describepublisher.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/cloudformation-2010-05-15/describepublisher.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/cloudformation-2010-05-15/describepublisher.md)

- [AWS SDK for Python](../../../../services/goto/boto3/cloudformation-2010-05-15/describepublisher.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/cloudformation-2010-05-15/describepublisher.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DescribeOrganizationsAccess

DescribeResourceScan

All content copied from https://docs.aws.amazon.com/.
