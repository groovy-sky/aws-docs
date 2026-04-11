# UpdateAssessmentFrameworkControlSet

A `controlSet` entity that represents a collection of controls in AWS Audit Manager. This doesn't contain the control set ID.

## Contents

**controls**

The list of controls that are contained within the control set.

Type: Array of [CreateAssessmentFrameworkControl](api-createassessmentframeworkcontrol.md) objects

Array Members: Minimum number of 1 item.

Required: Yes

**name**

The name of the control set.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 300.

Pattern: `^[^\\\_]*$`

Required: Yes

**id**

The unique identifier for the control set.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 300.

Pattern: `^[^\\\_]*$`

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/auditmanager-2017-07-25/updateassessmentframeworkcontrolset.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/auditmanager-2017-07-25/updateassessmentframeworkcontrolset.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/auditmanager-2017-07-25/updateassessmentframeworkcontrolset.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

SourceKeyword

URL

All content copied from https://docs.aws.amazon.com/.
