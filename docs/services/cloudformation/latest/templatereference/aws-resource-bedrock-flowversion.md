This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Bedrock::FlowVersion

Creates a version of the flow that you can deploy. For more information, see [Deploy a flow in Amazon Bedrock](https://docs.aws.amazon.com/bedrock/latest/userguide/flows-deploy.html) in the Amazon Bedrock User Guide.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Bedrock::FlowVersion",
  "Properties" : {
      "Description" : String,
      "FlowArn" : String
    }
}

```

### YAML

```yaml

Type: AWS::Bedrock::FlowVersion
Properties:
  Description: String
  FlowArn: String

```

## Properties

`Description`

The description of the flow version.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `200`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`FlowArn`

The Amazon Resource Name (ARN) of the flow that the version belongs to.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws(-[^:]+)?:bedrock:[a-z0-9-]{1,20}:[0-9]{12}:flow/[0-9a-zA-Z]{10}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Number (ARN) of the flow the version
of the flow, separated by a pipe ( `|`).

For example, `{ "Ref": "myFlowVersion" }` could return the value
`"arn:aws:bedrock:us-east-1:123456789012:flow/FLOW12345|1"`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

`CreatedAt`

The time at the version was created.

`CustomerEncryptionKeyArn`

The Amazon Resource Name (ARN) of the KMS key that the flow version is encrypted
with.

`ExecutionRoleArn`

The Amazon Resource Name (ARN) of the service role with permissions to create a flow. For more information, see [Create a service row for flows](https://docs.aws.amazon.com/bedrock/latest/userguide/flows-permissions.html) in the Amazon Bedrock User Guide.

`FlowId`

The unique identifier of the flow.

`Name`

The name of the flow.

`Status`

The status of the flow.

`Version`

The version of the flow.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

FlowAliasRoutingConfigurationListItem

AgentFlowNodeConfiguration
