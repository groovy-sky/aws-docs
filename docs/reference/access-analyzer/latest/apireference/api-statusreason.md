# StatusReason

Provides more details about the current status of the analyzer. For example, if the
creation for the analyzer fails, a `Failed` status is returned. For an analyzer
with organization as the type, this failure can be due to an issue with creating the
service-linked roles required in the member accounts of the AWS organization.

## Contents

**code**

The reason code for the current status of the analyzer.

Type: String

Valid Values: `AWS_SERVICE_ACCESS_DISABLED | DELEGATED_ADMINISTRATOR_DEREGISTERED | ORGANIZATION_DELETED | SERVICE_LINKED_ROLE_CREATION_FAILED`

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/accessanalyzer-2019-11-01/statusreason.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/accessanalyzer-2019-11-01/statusreason.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/accessanalyzer-2019-11-01/statusreason.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

SqsQueueConfiguration

Substring
