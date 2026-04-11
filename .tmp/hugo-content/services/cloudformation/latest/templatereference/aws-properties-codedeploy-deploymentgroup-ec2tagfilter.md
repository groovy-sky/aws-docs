This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeDeploy::DeploymentGroup EC2TagFilter

Information about an Amazon EC2 tag filter.

For more information about using tags and tag groups to help manage your Amazon EC2 instances and on-premises instances, see [Tagging Instances for Deployment\
Groups in AWS CodeDeploy](../../../codedeploy/latest/userguide/instances-tagging.md) in the _AWS CodeDeploy User_
_Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Key" : String,
  "Type" : String,
  "Value" : String
}

```

### YAML

```yaml

  Key: String
  Type: String
  Value: String

```

## Properties

`Key`

The tag filter key.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The tag filter type:

- `KEY_ONLY`: Key only.

- `VALUE_ONLY`: Value only.

- `KEY_AND_VALUE`: Key and value.

_Required_: No

_Type_: String

_Allowed values_: `KEY_ONLY | VALUE_ONLY | KEY_AND_VALUE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The tag filter value.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DeploymentStyle

EC2TagSet

All content copied from https://docs.aws.amazon.com/.
