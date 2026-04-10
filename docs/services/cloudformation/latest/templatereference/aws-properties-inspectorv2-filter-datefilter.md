This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::InspectorV2::Filter DateFilter

Contains details on the time range used to filter findings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "EndInclusive" : Integer,
  "StartInclusive" : Integer
}

```

### YAML

```yaml

  EndInclusive: Integer
  StartInclusive: Integer

```

## Properties

`EndInclusive`

A timestamp representing the end of the time period filtered on.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`StartInclusive`

A timestamp representing the start of the time period filtered on.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::InspectorV2::Filter

FilterCriteria

All content copied from https://docs.aws.amazon.com/.
