This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Backup::RestoreTestingSelection ProtectedResourceConditions

The conditions that you define for resources in
your restore testing plan using tags.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "StringEquals" : [ KeyValue, ... ],
  "StringNotEquals" : [ KeyValue, ... ]
}

```

### YAML

```yaml

  StringEquals:
    - KeyValue
  StringNotEquals:
    - KeyValue

```

## Properties

`StringEquals`

Filters the values of your tagged resources for only
those resources that you tagged with the same value.
Also called "exact matching."

_Required_: No

_Type_: Array of [KeyValue](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-backup-restoretestingselection-keyvalue.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StringNotEquals`

Filters the values of your tagged resources for only
those resources that you tagged that do not have the same value.
Also called "negated matching."

_Required_: No

_Type_: Array of [KeyValue](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-backup-restoretestingselection-keyvalue.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

KeyValue

AWS::Backup::TieringConfiguration
