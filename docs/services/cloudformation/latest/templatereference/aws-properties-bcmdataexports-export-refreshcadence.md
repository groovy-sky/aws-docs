This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::BCMDataExports::Export RefreshCadence

The cadence for AWS to update the data export in your S3 bucket.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Frequency" : String
}

```

### YAML

```yaml

  Frequency: String

```

## Properties

`Frequency`

The frequency that data exports are updated. The export refreshes each time the source
data updates, up to three times daily.

_Required_: Yes

_Type_: String

_Allowed values_: `SYNCHRONOUS`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Export

ResourceTag

All content copied from https://docs.aws.amazon.com/.
