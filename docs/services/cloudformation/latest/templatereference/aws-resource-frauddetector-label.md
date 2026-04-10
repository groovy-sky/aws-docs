This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::FraudDetector::Label

Creates or updates label. A label classifies an event as fraudulent or legitimate. Labels are associated with event types and used to train supervised machine learning models in Amazon Fraud Detector.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::FraudDetector::Label",
  "Properties" : {
      "Description" : String,
      "Name" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::FraudDetector::Label
Properties:
  Description: String
  Name: String
  Tags:
    - Tag

```

## Properties

`Description`

The label description.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The label name.

Pattern: `^[0-9a-z_-]+$`

_Required_: Yes

_Type_: String

_Pattern_: `^[0-9a-z_-]+$`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

An array of key-value pairs to apply to this resource.

For more information, see [Tag](../userguide/aws-properties-resource-tags.md).

_Required_: No

_Type_: Array of [Tag](aws-properties-frauddetector-label-tag.md)

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the primary identifier for the resource, which is the ARN.

Example: `{"Ref": "arn:aws:frauddetector:us-west-2:123123123123:outcome/outcome_name"}`

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The ARN of the label.

`CreatedTime`

Timestamp of when label was created.

`LastUpdatedTime`

Timestamp of when label was last updated.

## See also

- [PutLabel](../../../frauddetector/latest/api/api-putlabel.md) in the _Amazon Fraud Detector API Reference_

- [Create a label](../../../frauddetector/latest/ug/create-a-label.md) in the _Amazon Fraud Detector User Guide_

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

Tag

All content copied from https://docs.aws.amazon.com/.
