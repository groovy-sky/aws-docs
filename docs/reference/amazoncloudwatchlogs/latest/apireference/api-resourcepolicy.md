# ResourcePolicy

A policy enabling one or more entities to put logs to a log group in this
account.

## Contents

**lastUpdatedTime**

Timestamp showing when this policy was last updated, expressed as the number of
milliseconds after `Jan 1, 1970 00:00:00 UTC`.

Type: Long

Valid Range: Minimum value of 0.

Required: No

**policyDocument**

The details of the policy.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 51200.

Required: No

**policyName**

The name of the resource policy.

Type: String

Required: No

**policyScope**

Specifies scope of the resource policy. Valid values are ACCOUNT or RESOURCE.

Type: String

Valid Values: `ACCOUNT | RESOURCE`

Required: No

**resourceArn**

The ARN of the CloudWatch Logs resource to which the resource policy is attached. Only
populated for resource-scoped policies.

Type: String

Required: No

**revisionId**

The revision ID of the resource policy. Only populated for resource-scoped
policies.

Type: String

Length Constraints: Minimum length of 1.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/logs-2014-03-28/resourcepolicy.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/logs-2014-03-28/resourcepolicy.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/logs-2014-03-28/resourcepolicy.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ResourceConfig

ResultField

All content copied from https://docs.aws.amazon.com/.
