This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::PCS::ComputeNodeGroup CustomLaunchTemplate

An Amazon EC2 launch template AWS PCS uses to launch compute
nodes.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "TemplateId" : String,
  "Version" : String
}

```

### YAML

```yaml

  TemplateId: String
  Version: String

```

## Properties

`TemplateId`

The ID of the EC2 launch template to use to provision instances.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Version`

The version of the EC2 launch template to use to provision instances.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::PCS::ComputeNodeGroup

ErrorInfo

All content copied from https://docs.aws.amazon.com/.
