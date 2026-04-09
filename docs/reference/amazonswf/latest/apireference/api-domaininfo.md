# DomainInfo

Contains general information about a domain.

## Contents

**name**

The name of the domain. This name is unique within the account.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

Required: Yes

**status**

The status of the domain:

- `REGISTERED` – The domain is properly registered and available. You can use this domain
for registering types and creating new workflow executions.

- `DEPRECATED` – The domain was deprecated using [DeprecateDomain](api-deprecatedomain.md), but is
still in use. You should not create new workflow executions in this domain.

Type: String

Valid Values: `REGISTERED | DEPRECATED`

Required: Yes

**arn**

The ARN of the domain.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1600.

Required: No

**description**

The description of the domain provided through [RegisterDomain](api-registerdomain.md).

Type: String

Length Constraints: Maximum length of 1024.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/swf-2012-01-25/domaininfo.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/swf-2012-01-25/domaininfo.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/swf-2012-01-25/domaininfo.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DomainConfiguration

ExecutionTimeFilter

All content copied from https://docs.aws.amazon.com/.
