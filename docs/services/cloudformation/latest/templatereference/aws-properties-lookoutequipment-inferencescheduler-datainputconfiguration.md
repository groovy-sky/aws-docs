This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::LookoutEquipment::InferenceScheduler DataInputConfiguration

The `DataInputConfiguration` property type specifies Property description not available. for an [AWS::LookoutEquipment::InferenceScheduler](aws-resource-lookoutequipment-inferencescheduler.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "InferenceInputNameConfiguration" : InputNameConfiguration,
  "InputTimeZoneOffset" : String,
  "S3InputConfiguration" : S3InputConfiguration
}

```

### YAML

```yaml

  InferenceInputNameConfiguration:
    InputNameConfiguration
  InputTimeZoneOffset: String
  S3InputConfiguration:
    S3InputConfiguration

```

## Properties

`InferenceInputNameConfiguration`

Property description not available.

_Required_: No

_Type_: [InputNameConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-lookoutequipment-inferencescheduler-inputnameconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`InputTimeZoneOffset`

Property description not available.

_Required_: No

_Type_: String

_Pattern_: `^(\+|\-)[0-9]{2}\:[0-9]{2}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3InputConfiguration`

Property description not available.

_Required_: Yes

_Type_: [S3InputConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-lookoutequipment-inferencescheduler-s3inputconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::LookoutEquipment::InferenceScheduler

DataOutputConfiguration
