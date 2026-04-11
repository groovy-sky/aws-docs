This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DevOpsAgent::AgentSpace

The `AWS::DevOpsAgent::AgentSpace` resource specifies an Agent Space for the AWS
DevOps Agent Service.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::DevOpsAgent::AgentSpace",
  "Properties" : {
      "Description" : String,
      "KmsKeyArn" : String,
      "Name" : String,
      "OperatorApp" : OperatorApp,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::DevOpsAgent::AgentSpace
Properties:
  Description: String
  KmsKeyArn: String
  Name: String
  OperatorApp:
    OperatorApp
  Tags:
    - Tag

```

## Properties

`Description`

The description of the Agent Space.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`KmsKeyArn`

Property description not available.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the Agent Space.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `255`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OperatorApp`

Configuration for the connection to the DevOps Agent web app.

_Required_: No

_Type_: [OperatorApp](aws-properties-devopsagent-agentspace-operatorapp.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Property description not available.

_Required_: No

_Type_: Array of [Tag](aws-properties-devopsagent-agentspace-tag.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the AgentSpaceId.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AgentSpaceId`

The unique identifier of the Agent Space.

`Arn`

The Amazon Resource Name (ARN) of the Agent Space.

`CreatedAt`

The timestamp when the resource was created.

`OperatorApp.Iam.CreatedAt`

The timestamp when the IAM authentication configuration was created.

`OperatorApp.Iam.UpdatedAt`

The timestamp when the IAM authentication configuration was last updated.

`OperatorApp.Idc.CreatedAt`

The timestamp when the IAM Identity Center authentication configuration was created.

`OperatorApp.Idc.IdcApplicationArn`

The ARN of the IAM Identity Center application created for the DevOps Agent web app.

`OperatorApp.Idc.UpdatedAt`

The timestamp when the IAM Identity Center authentication configuration was last updated.

`UpdatedAt`

The timestamp when the resource was last updated.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS DevOps Agent

IamAuthConfiguration

All content copied from https://docs.aws.amazon.com/.
