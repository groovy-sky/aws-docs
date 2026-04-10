This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ARCRegionSwitch::Plan Route53ResourceRecordSet

The Amazon Route 53 record set.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "RecordSetIdentifier" : String,
  "Region" : String
}

```

### YAML

```yaml

  RecordSetIdentifier: String
  Region: String

```

## Properties

`RecordSetIdentifier`

The Amazon Route 53 record set identifier.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Region`

The Amazon Route 53 record set Region.

_Required_: No

_Type_: String

_Pattern_: `^[a-z]{2}-[a-z-]+-\d+$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Route53HealthCheckConfiguration

S3ReportOutputConfiguration

All content copied from https://docs.aws.amazon.com/.
