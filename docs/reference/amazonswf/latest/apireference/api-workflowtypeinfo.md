# WorkflowTypeInfo

Contains information about a workflow type.

## Contents

**creationDate**

The date when this type was registered.

Type: Timestamp

Required: Yes

**status**

The current status of the workflow type.

Type: String

Valid Values: `REGISTERED | DEPRECATED`

Required: Yes

**workflowType**

The workflow type this information is about.

Type: [WorkflowType](api-workflowtype.md) object

Required: Yes

**deprecationDate**

If the type is in deprecated state, then it is set to the date when the type was deprecated.

Type: Timestamp

Required: No

**description**

The description of the type registered through [RegisterWorkflowType](api-registerworkflowtype.md).

Type: String

Length Constraints: Maximum length of 1024.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/swf-2012-01-25/workflowtypeinfo.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/swf-2012-01-25/workflowtypeinfo.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/swf-2012-01-25/workflowtypeinfo.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

WorkflowTypeFilter

Common Parameters

All content copied from https://docs.aws.amazon.com/.
