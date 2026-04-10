This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MSK::Configuration LatestRevision

Describes a configuration revision.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CreationTime" : String,
  "Description" : String,
  "Revision" : Integer
}

```

### YAML

```yaml

  CreationTime: String
  Description: String
  Revision: Integer

```

## Properties

`CreationTime`

The time when the configuration revision was created.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of the configuration revision.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Revision`

The revision number.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::MSK::Configuration

AWS::MSK::Replicator

All content copied from https://docs.aws.amazon.com/.
