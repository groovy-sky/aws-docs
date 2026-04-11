This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DevOpsGuru::ResourceCollection CloudFormationCollectionFilter

Information about AWS CloudFormation stacks. You can use up to 1000
stacks to specify which AWS resources in your account to analyze. For more
information, see [Stacks](../userguide/stacks.md) in the
_AWS CloudFormation User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "StackNames" : [ String, ... ]
}

```

### YAML

```yaml

  StackNames:
    - String

```

## Properties

`StackNames`

An array of CloudFormation stack names.

_Required_: No

_Type_: Array of String

_Minimum_: `1 | 1`

_Maximum_: `128 | 1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::DevOpsGuru::ResourceCollection

ResourceCollectionFilter

All content copied from https://docs.aws.amazon.com/.
