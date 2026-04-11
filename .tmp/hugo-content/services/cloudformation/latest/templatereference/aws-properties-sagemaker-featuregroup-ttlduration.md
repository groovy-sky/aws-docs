This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SageMaker::FeatureGroup TtlDuration

Time to live duration, where the record is hard deleted after the expiration time is
reached; `ExpiresAt` = `EventTime` \+ `TtlDuration`. For
information on HardDelete, see the [DeleteRecord](../../../../reference/sagemaker/latest/apireference/api-feature-store-deleterecord.md) API in the Amazon SageMaker API Reference guide.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Unit" : String,
  "Value" : Integer
}

```

### YAML

```yaml

  Unit: String
  Value: Integer

```

## Properties

`Unit`

`TtlDuration` time unit.

_Required_: No

_Type_: String

_Allowed values_: `Seconds | Minutes | Hours | Days | Weeks`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

`TtlDuration` time value.

_Required_: No

_Type_: Integer

_Minimum_: `1`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ThroughputConfig

AWS::SageMaker::Image

All content copied from https://docs.aws.amazon.com/.
