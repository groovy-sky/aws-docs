This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DevOpsAgent::AgentSpace IamAuthConfiguration

IAM-based authentication configuration for the DevOps Agent web app.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CreatedAt" : String,
  "OperatorAppRoleArn" : String,
  "UpdatedAt" : String
}

```

### YAML

```yaml

  CreatedAt: String
  OperatorAppRoleArn: String
  UpdatedAt: String

```

## Properties

`CreatedAt`

The timestamp when the IAM authentication configuration was created.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`OperatorAppRoleArn`

The ARN of the IAM role that grants access to the DevOps Agent web app.

_Required_: Conditional

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UpdatedAt`

The timestamp when the IAM authentication configuration was last updated.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::DevOpsAgent::AgentSpace

IdcAuthConfiguration

All content copied from https://docs.aws.amazon.com/.
