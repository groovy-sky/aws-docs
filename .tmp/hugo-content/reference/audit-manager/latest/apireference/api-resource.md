# Resource

A system asset that's evaluated in an Audit Manager assessment.

## Contents

**arn**

The Amazon Resource Name (ARN) for the resource.

Type: String

Length Constraints: Minimum length of 20. Maximum length of 2048.

Pattern: `^arn:.*`

Required: No

**complianceCheck**

The evaluation status for a resource that was assessed when collecting compliance check
evidence.

- Audit Manager classes the resource as non-compliant if Security Hub CSPM reports a
_Fail_ result, or if AWS Config reports a
_Non-compliant_ result.

- Audit Manager classes the resource as compliant if Security Hub CSPM reports a
_Pass_ result, or if AWS Config reports a
_Compliant_ result.

- If a compliance check isn't available or applicable, then no compliance evaluation can be made
for that resource. This is the case if a resource assessment uses AWS Config or Security Hub CSPM as the underlying data source type, but those services
aren't enabled. This is also the case if the resource assessment uses an underlying
data source type that doesn't support compliance checks (such as manual evidence,
AWS API calls, or CloudTrail).

Type: String

Length Constraints: Minimum length of 0. Maximum length of 2048.

Pattern: `.*`

Required: No

**value**

The value of the resource.

Type: String

Length Constraints: Minimum length of 0. Maximum length of 2048.

Pattern: `.*`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/auditmanager-2017-07-25/resource.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/auditmanager-2017-07-25/resource.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/auditmanager-2017-07-25/resource.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Notification

Role

All content copied from https://docs.aws.amazon.com/.
