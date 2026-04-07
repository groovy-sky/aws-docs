# AccountPolicy

A structure that contains information about one CloudWatch Logs account policy.

## Contents

**accountId**

The AWS account ID that the policy applies to.

Type: String

Length Constraints: Fixed length of 12.

Pattern: `^\d{12}$`

Required: No

**lastUpdatedTime**

The date and time that this policy was most recently updated.

Type: Long

Valid Range: Minimum value of 0.

Required: No

**policyDocument**

The policy document for this account policy.

The JSON specified in `policyDocument` can be up to 30,720 characters.

Type: String

Required: No

**policyName**

The name of the account policy.

Type: String

Required: No

**policyType**

The type of policy for this account policy.

Type: String

Valid Values: `DATA_PROTECTION_POLICY | SUBSCRIPTION_FILTER_POLICY | FIELD_INDEX_POLICY | TRANSFORMER_POLICY | METRIC_EXTRACTION_POLICY`

Required: No

**scope**

The scope of the account policy.

Type: String

Valid Values: `ALL`

Required: No

**selectionCriteria**

The log group selection criteria that is used for this policy.

Type: String

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/logs-2014-03-28/AccountPolicy)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/logs-2014-03-28/AccountPolicy)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/logs-2014-03-28/AccountPolicy)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Data Types

AddKeyEntry
