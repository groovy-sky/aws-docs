# HookTarget

The `HookTarget` data type.

## Contents

**Action**

The action that invoked the Hook.

Type: String

Valid Values: `CREATE | UPDATE | DELETE | IMPORT`

Required: Yes

**TargetId**

The unique identifier of the Hook invocation target.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

Pattern: `[^_]*`

Required: Yes

**TargetType**

The target type.

Type: String

Valid Values: `RESOURCE`

Required: Yes

**TargetTypeName**

The target name, for example, `AWS::S3::Bucket`.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 256.

Pattern: `^[a-zA-Z0-9]{2,64}::[a-zA-Z0-9]{2,64}::[a-zA-Z0-9]{2,64}$`

Required: Yes

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/cloudformation-2010-05-15/hooktarget.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/cloudformation-2010-05-15/hooktarget.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/cloudformation-2010-05-15/hooktarget.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

HookResultSummary

LiveResourceDrift

All content copied from https://docs.aws.amazon.com/.
