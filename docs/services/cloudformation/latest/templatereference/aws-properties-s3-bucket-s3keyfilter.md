This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::S3::Bucket S3KeyFilter

A container for object key name prefix and suffix filtering rules. For more information
about object key name filtering, see [Configuring event\
notifications using object key name filtering](../../../s3/latest/userguide/notification-how-to-filtering.md) in the _Amazon S3 User_
_Guide_.

###### Note

The same type of filter rule cannot be used more than once. For example, you cannot
specify two prefix rules.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Rules" : [ FilterRule, ... ]
}

```

### YAML

```yaml

  Rules:
    - FilterRule

```

## Properties

`Rules`

A list of containers for the key-value pair that defines the criteria for the filter rule.

_Required_: Yes

_Type_: Array of [FilterRule](aws-properties-s3-bucket-filterrule.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Rule

S3TablesDestination

All content copied from https://docs.aws.amazon.com/.
