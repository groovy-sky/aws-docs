This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::APS::Workspace LimitsPerLabelSetEntry

This structure contains the limits that apply to time series that match one label
set.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MaxSeries" : Integer
}

```

### YAML

```yaml

  MaxSeries: Integer

```

## Properties

`MaxSeries`

The maximum number of active series that can be ingested that match this label set.

Setting this to 0 causes no label set limit to be enforced, but it does cause Amazon Managed Service for Prometheus to vend label set metrics to CloudWatch

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

LimitsPerLabelSet

LoggingConfiguration

All content copied from https://docs.aws.amazon.com/.
