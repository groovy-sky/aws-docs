# ControlMetadata

The metadata that's associated with the standard control or custom control.

## Contents

**arn**

The Amazon Resource Name (ARN) of the control.

Type: String

Length Constraints: Minimum length of 20. Maximum length of 2048.

Pattern: `^arn:.*:auditmanager:.*`

Required: No

**controlSources**

The data source that determines where Audit Manager collects evidence from for the
control.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 100.

Pattern: `^[a-zA-Z_0-9-\s.,]+$`

Required: No

**createdAt**

The time when the control was created.

Type: Timestamp

Required: No

**id**

The unique identifier for the control.

Type: String

Length Constraints: Fixed length of 36.

Pattern: `^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$`

Required: No

**lastUpdatedAt**

The time when the control was most recently updated.

Type: Timestamp

Required: No

**name**

The name of the control.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 300.

Pattern: `^[^\\]*$`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/auditmanager-2017-07-25/controlmetadata.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/auditmanager-2017-07-25/controlmetadata.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/auditmanager-2017-07-25/controlmetadata.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ControlMappingSource

ControlSet

All content copied from https://docs.aws.amazon.com/.
