# CreateAssessmentFrameworkControlSet

A `controlSet` entity that represents a collection of controls in AWS Audit Manager. This doesn't contain the control set ID.

## Contents

**name**

The name of the control set.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 300.

Pattern: `^[^\\\_]*$`

Required: Yes

**controls**

The list of controls within the control set. This doesn't contain the control set ID.

Type: Array of [CreateAssessmentFrameworkControl](api-createassessmentframeworkcontrol.md) objects

Array Members: Minimum number of 1 item.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/auditmanager-2017-07-25/createassessmentframeworkcontrolset.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/auditmanager-2017-07-25/createassessmentframeworkcontrolset.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/auditmanager-2017-07-25/createassessmentframeworkcontrolset.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CreateAssessmentFrameworkControl

CreateControlMappingSource

All content copied from https://docs.aws.amazon.com/.
