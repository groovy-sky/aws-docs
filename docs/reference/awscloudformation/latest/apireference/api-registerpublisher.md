# RegisterPublisher

Registers your account as a publisher of public extensions in the CloudFormation registry.
Public extensions are available for use by all CloudFormation users. This publisher ID applies to
your account in all AWS Regions.

For information about requirements for registering as a public extension publisher, see
[Prerequisite: Registering your account to publish CloudFormation extensions](../../../../services/cloudformation-cli/latest/userguide/publish-extension-publish-extension-prereqs.md) in the
_AWS CloudFormation Command Line Interface (CLI) User Guide_.

## Request Parameters

For information about the parameters that are common to all actions, see [Common Parameters](commonparameters.md).

**AcceptTermsAndConditions**

Whether you accept the [Terms and Conditions](https://cloudformation-registry-documents.s3.amazonaws.com/Terms_and_Conditions_for_AWS_CloudFormation_Registry_Publishers.pdf) for publishing extensions in the CloudFormation registry. You must
accept the terms and conditions in order to register to publish public extensions to the
CloudFormation registry.

The default is `false`.

Type: Boolean

Required: No

**ConnectionArn**

If you are using a Bitbucket or GitHub account for identity verification, the Amazon
Resource Name (ARN) for your connection to that account.

For more information, see [Prerequisite: Registering your account to publish CloudFormation extensions](../../../../services/cloudformation-cli/latest/userguide/publish-extension-publish-extension-prereqs.md) in the
_AWS CloudFormation Command Line Interface (CLI) User Guide_.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

Pattern: `arn:aws(-[\w]+)*:.+:.+:[0-9]{12}:.+`

Required: No

## Response Elements

The following element is returned by the service.

**PublisherId**

The ID assigned this account by CloudFormation for publishing extensions.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 40.

Pattern: `[0-9a-zA-Z]{12,40}`

## Errors

For information about the errors that are common to all actions, see [Common Error Types](commonerrors.md).

**CFNRegistry**

An error occurred during a CloudFormation registry operation.

**Message**

A message with details about the error that occurred.

HTTP Status Code: 400

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/cloudformation-2010-05-15/registerpublisher.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/cloudformation-2010-05-15/registerpublisher.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/cloudformation-2010-05-15/registerpublisher.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/cloudformation-2010-05-15/registerpublisher.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/cloudformation-2010-05-15/registerpublisher.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/cloudformation-2010-05-15/registerpublisher.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/cloudformation-2010-05-15/registerpublisher.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/cloudformation-2010-05-15/registerpublisher.md)

- [AWS SDK for Python](../../../../services/goto/boto3/cloudformation-2010-05-15/registerpublisher.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/cloudformation-2010-05-15/registerpublisher.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

RecordHandlerProgress

RegisterType
