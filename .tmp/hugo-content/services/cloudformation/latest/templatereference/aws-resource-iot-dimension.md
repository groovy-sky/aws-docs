This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IoT::Dimension

Use the `AWS::IoT::Dimension` to limit the scope of a metric used in a
security profile for AWS IoT Device Defender. For example, using a `TOPIC_FILTER`
dimension, you can narrow down the scope of the metric to only MQTT topics where the name
matches the pattern specified in the dimension. For API reference, see [CreateDimension](../../../../reference/iot/latest/apireference/api-createdimension.md) and for general information, see [Scoping metrics in\
security profiles using dimensions](../../../iot/latest/developerguide/scoping-security-behavior.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::IoT::Dimension",
  "Properties" : {
      "Name" : String,
      "StringValues" : [ String, ... ],
      "Tags" : [ Tag, ... ],
      "Type" : String
    }
}

```

### YAML

```yaml

Type: AWS::IoT::Dimension
Properties:
  Name: String
  StringValues:
    - String
  Tags:
    - Tag
  Type: String

```

## Properties

`Name`

A unique identifier for the dimension.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9:_-]+`

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`StringValues`

Specifies the value or list of values for the dimension. For `TOPIC_FILTER`
dimensions, this is a pattern used to match the MQTT topic (for example, "admin/#").

_Required_: Yes

_Type_: Array of String

_Minimum_: `1 | 1`

_Maximum_: `256 | 5`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Metadata that can be used to manage the dimension.

_Required_: No

_Type_: Array of [Tag](aws-properties-iot-dimension-tag.md)

_Maximum_: `50`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

Specifies the type of dimension. Supported types: `TOPIC_FILTER.`

_Required_: Yes

_Type_: String

_Allowed values_: `TOPIC_FILTER`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the dimension name.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) of the dimension.

## Examples

#### JSON

```json

{
  "AWSTemplateFormatVersion": "2010-09-09",
  "Description": "Amazon Web Services IoT Dimension Sample Template",
  "Resources": {
    "TopicFilterForAuthMessagesDimension": {
      "Type": "AWS::IoT::Dimension",
      "Properties": {
        "Name": "TopicFilterForAuthMessages",
        "Type": "TOPIC_FILTER",
        "StringValues": [
          "device/+/auth"
        ],
        "Tags": [
          {
            "Key": "Application",
            "Value": "SmartHome"
          }
        ]
      }
    }
  }
}
```

#### YAML

```yaml

AWSTemplateFormatVersion: '2010-09-09'
Description: Amazon Web Services IoT Dimension Sample Template
Resources:
  TopicFilterForAuthMessagesDimension:
    Type: 'AWS::IoT::Dimension'
    Properties:
      Name: TopicFilterForAuthMessages
      Type: TOPIC_FILTER
      StringValues:
        - device/+/auth
      Tags:
        - Key: Application
          Value: SmartHome
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Tag

All content copied from https://docs.aws.amazon.com/.
