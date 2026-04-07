This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::FleetMetric

Use the `AWS::IoT::FleetMetric` resource to declare a fleet metric.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IoT::FleetMetric",
  "Properties" : {
      "AggregationField" : String,
      "AggregationType" : AggregationType,
      "Description" : String,
      "IndexName" : String,
      "MetricName" : String,
      "Period" : Integer,
      "QueryString" : String,
      "QueryVersion" : String,
      "Tags" : [ Tag, ... ],
      "Unit" : String
    }
}

```

### YAML

```yaml

Type: AWS::IoT::FleetMetric
Properties:
  AggregationField: String
  AggregationType:
    AggregationType
  Description: String
  IndexName: String
  MetricName: String
  Period: Integer
  QueryString:
    String
  QueryVersion: String
  Tags:
    - Tag
  Unit: String

```

## Properties

`AggregationField`

The field to aggregate.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AggregationType`

The type of the aggregation query.

_Required_: No

_Type_: [AggregationType](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iot-fleetmetric-aggregationtype.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The fleet metric description.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IndexName`

The name of the index to search.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`MetricName`

The name of the fleet metric to create.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Period`

The time in seconds between fleet metric emissions. Range \[60(1 min), 86400(1 day)\] and
must be multiple of 60.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`QueryString`

The search query string.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`QueryVersion`

The query version.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Metadata which can be used to manage the fleet metric.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-iot-fleetmetric-tag.html)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Unit`

Used to support unit transformation such as milliseconds to seconds. Must be a unit
supported by CW metric. Default to null.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the fleet metric name.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CreationDate`

The time the fleet metric was created.

`LastModifiedDate`

The time the fleet metric was last modified.

`MetricArn`

The Amazon Resource Name (ARN) of the fleet metric.

`Version`

The fleet metric version.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ConfigurationDetails

AggregationType
