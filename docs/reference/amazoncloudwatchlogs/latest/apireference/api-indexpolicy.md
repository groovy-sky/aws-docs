---
title: "IndexPolicy"
---

# IndexPolicy

This structure contains information about one field index policy in this account.

## Contents

**lastUpdateTime**

The date and time that this index policy was most recently updated.

Type: Long

Valid Range: Minimum value of 0.

Required: No

**logGroupIdentifier**

The ARN of the log group that this index policy applies to.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 2048.

Pattern: `[\w#+=/:,.@-]*`

Required: No

**policyDocument**

The policy document for this index policy, in JSON format.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 51200.

Required: No

**policyName**

The name of this policy. Responses about log group-level field index policies don't have
this field, because those policies don't have names.

Type: String

Required: No

**source**

This field indicates whether this is an account-level index policy or an index policy that
applies only to a single log group.

Type: String

Valid Values: `ACCOUNT | LOG_GROUP`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/logs-2014-03-28/IndexPolicy)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/logs-2014-03-28/IndexPolicy)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/logs-2014-03-28/IndexPolicy)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ImportStatistics

InputLogEvent

All content copied from https://docs.aws.amazon.com/.
