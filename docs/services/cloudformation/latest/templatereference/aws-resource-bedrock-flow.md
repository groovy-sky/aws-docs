This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::Flow

Creates a prompt flow that you can use to send an input through various steps to yield
an output. You define a flow by configuring nodes, each of which corresponds to a step
of the flow, and creating connections between the nodes to create paths to different
outputs. You can define the flow in one of the following ways:

- Define a [FlowDefinition](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-flowdefinition.html) in the `Definition` property.

- Provide the definition in the `DefinitionString` property as a
JSON-formatted string matching the [FlowDefinition](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-flowdefinition.html) property.

- Provide an Amazon S3 location in the `DefinitionS3Location`
property that matches the [FlowDefinition](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-flowdefinition.html).

If you use the `DefinitionString` or `DefinitionS3Location`
property, you can use the `DefinitionSubstitutions` property to define
key-value pairs to replace at runtime.

For more information, see [How it works](https://docs.aws.amazon.com/bedrock/latest/userguide/flows-how-it-works.html) and
[Create\
a prompt flow in Amazon Bedrock](https://docs.aws.amazon.com/bedrock/latest/userguide/flows-create.html) in the Amazon Bedrock User Guide.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Bedrock::Flow",
  "Properties" : {
      "CustomerEncryptionKeyArn" : String,
      "Definition" : FlowDefinition,
      "DefinitionS3Location" : S3Location,
      "DefinitionString" : String,
      "DefinitionSubstitutions" : {Key: Value, ...},
      "Description" : String,
      "ExecutionRoleArn" : String,
      "Name" : String,
      "Tags" : {Key: Value, ...},
      "TestAliasTags" : {Key: Value, ...}
    }
}

```

### YAML

```yaml

Type: AWS::Bedrock::Flow
Properties:
  CustomerEncryptionKeyArn: String
  Definition:
    FlowDefinition
  DefinitionS3Location:
    S3Location
  DefinitionString:
    String
  DefinitionSubstitutions:
    Key: Value
  Description: String
  ExecutionRoleArn: String
  Name: String
  Tags:
    Key: Value
  TestAliasTags:
    Key: Value

```

## Properties

`CustomerEncryptionKeyArn`

The Amazon Resource Name (ARN) of the KMS key that the flow is encrypted with.

_Required_: No

_Type_: String

_Pattern_: `^arn:aws(|-cn|-us-gov):kms:[a-zA-Z0-9-]*:[0-9]{12}:key/[a-zA-Z0-9-]{36}$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Definition`

The definition of the nodes and connections between the nodes in the flow.

_Required_: No

_Type_: [FlowDefinition](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrock-flow-flowdefinition.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefinitionS3Location`

The Amazon S3 location of the flow definition.

_Required_: No

_Type_: [S3Location](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-bedrock-flow-s3location.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefinitionString`

The definition of the flow as a JSON-formatted string. The string must match the
format in [FlowDefinition](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-bedrock-flow-flowdefinition.html).

_Required_: No

_Type_: String

_Maximum_: `512000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DefinitionSubstitutions`

A map that specifies the mappings for placeholder variables in the prompt flow
definition. This enables the customer to inject values obtained at runtime. Variables
can be template parameter names, resource logical IDs, resource attributes, or a
variable in a key-value map. Only supported with the `DefinitionString` and
`DefinitionS3Location` fields.

Substitutions must follow the syntax: `${key_name}` or
`${variable_1,variable_2,...}`.

_Required_: No

_Type_: Object

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

A description of the flow.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExecutionRoleArn`

The Amazon Resource Name (ARN) of the service role with permissions to create a flow. For more information, see [Create a service row for flows](https://docs.aws.amazon.com/bedrock/latest/userguide/flows-permissions.html) in the Amazon Bedrock User Guide.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws(-[^:]+)?:iam::([0-9]{12})?:role/(service-role/)?.+$`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the flow.

_Required_: Yes

_Type_: String

_Pattern_: `^([0-9a-zA-Z][_-]?){1,100}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Metadata that you can assign to a resource as key-value pairs. For more information,
see the following resources:

- [Tag naming\
limits and requirements](https://docs.aws.amazon.com/tag-editor/latest/userguide/tagging.html#tag-conventions)

- [Tagging\
best practices](https://docs.aws.amazon.com/tag-editor/latest/userguide/tagging.html#tag-best-practices)

_Required_: No

_Type_: Object of String

_Pattern_: `^[a-zA-Z0-9\s._:/=+@-]*$`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TestAliasTags`

Property description not available.

_Required_: No

_Type_: Object of String

_Pattern_: `^[a-zA-Z0-9\s._:/=+@-]*$`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Number (ARN) of the flow.

For example, `{ "Ref": "myFlow" }` could return the value
`""arn:aws:bedrock:us-east-1:123456789012:flow/FLOW123456""`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

`Arn`

The Amazon Resource Name (ARN) of the flow.

`CreatedAt`

The time at which the flow was created.

`Id`

The unique identifier of the flow.

`Status`

The status of the flow. The following statuses are possible:

- NotPrepared – The flow has been created or updated, but hasn't been prepared. If you just created the flow, you can't test it. If you updated the flow, the `DRAFT` version won't contain the latest changes for testing. Send a [PrepareFlow](https://docs.aws.amazon.com/bedrock/latest/APIReference/API_agent_PrepareFlow.html) request to package the latest changes into the `DRAFT` version.

- Preparing – The flow is being prepared so that the `DRAFT` version contains the latest changes for testing.

- Prepared – The flow is prepared and the `DRAFT` version contains the latest changes for testing.

- Failed – The last API operation that you invoked on the flow failed. Send a [GetFlow](https://docs.aws.amazon.com/bedrock/latest/APIReference/API_agent_GetFlow.html) request and check the error message in the `validations` field.

`UpdatedAt`

The time at which the flow was last updated.

`Validations`

Property description not available.

`Version`

The latest version of the flow.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

WebSourceConfiguration

AgentFlowNodeConfiguration
