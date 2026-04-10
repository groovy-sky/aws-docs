This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataBrew::Job AllowedStatistics

Configuration of statistics that are allowed to be run on columns that
contain detected entities. When undefined, no statistics will be computed
on columns that contain detected entities.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Statistics" : [ String, ... ]
}

```

### YAML

```yaml

  Statistics:
    - String

```

## Properties

`Statistics`

One or more column statistics to allow for columns that contain detected entities.

_Required_: Yes

_Type_: Array of String

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::DataBrew::Job

ColumnSelector

All content copied from https://docs.aws.amazon.com/.
