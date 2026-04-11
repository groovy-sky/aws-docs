# Annotation

The `Annotation` data type.

A `GetHookResult` call returns detailed information and remediation guidance from
Control Tower, Guard, Lambda, or custom Hooks for a Hook invocation result.

## Contents

**AnnotationName**

An identifier for the evaluation logic that was used when invoking the Hook. For Control Tower,
this is the control ID. For Guard, this is the rule ID. For Lambda and custom Hooks,
this is a user-defined identifier.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1024.

Required: No

**RemediationLink**

A URL that you can access for additional remediation guidance.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 5120.

Required: No

**RemediationMessage**

Suggests what to change if your Hook returns a `FAILED` status. For example,
"Block public access to the bucket".

Type: String

Length Constraints: Minimum length of 1. Maximum length of 16384.

Required: No

**SeverityLevel**

The relative risk associated with any violations of this type.

Type: String

Valid Values: `INFORMATIONAL | LOW | MEDIUM | HIGH | CRITICAL`

Required: No

**Status**

The status of the Hook invocation from the downstream service.

Type: String

Valid Values: `PASSED | FAILED | SKIPPED`

Required: No

**StatusMessage**

The explanation for the specific status assigned to this Hook invocation. For example,
"Bucket does not block public access".

Type: String

Length Constraints: Minimum length of 1. Maximum length of 16384.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/cloudformation-2010-05-15/annotation.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/cloudformation-2010-05-15/annotation.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/cloudformation-2010-05-15/annotation.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AccountLimit

AutoDeployment

All content copied from https://docs.aws.amazon.com/.
