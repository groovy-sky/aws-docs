This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::MLTransform InputRecordTables

A list of AWS Glue table definitions used by the transform.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "GlueTables" : [ GlueTables, ... ]
}

```

### YAML

```yaml

  GlueTables:
    - GlueTables

```

## Properties

`GlueTables`

The database and table in the AWS Glue Data Catalog that is used for input or output data.

_Required_: No

_Type_: [Array](aws-properties-glue-mltransform-gluetables.md) of [GlueTables](aws-properties-glue-mltransform-gluetables.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

GlueTables

MLUserDataEncryption

All content copied from https://docs.aws.amazon.com/.
