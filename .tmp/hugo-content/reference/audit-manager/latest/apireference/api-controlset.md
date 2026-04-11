# ControlSet

A set of controls in AWS Audit Manager.

## Contents

**controls**

The list of controls within the control set.

Type: Array of [Control](api-control.md) objects

Array Members: Minimum number of 1 item.

Required: No

**id**

The identifier of the control set in the assessment. This is the control set name in a
plain string format.

Type: String

Length Constraints: Fixed length of 36.

Pattern: `^[a-f0-9]{8}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{4}-[a-f0-9]{12}$`

Required: No

**name**

The name of the control set.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 300.

Pattern: `^[^\\\_]*$`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/auditmanager-2017-07-25/controlset.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/auditmanager-2017-07-25/controlset.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/auditmanager-2017-07-25/controlset.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ControlMetadata

CreateAssessmentFrameworkControl

All content copied from https://docs.aws.amazon.com/.
