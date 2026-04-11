This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::LookoutEquipment::InferenceScheduler DataOutputConfiguration

The `DataOutputConfiguration` property type specifies Property description not available. for an [AWS::LookoutEquipment::InferenceScheduler](aws-resource-lookoutequipment-inferencescheduler.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "KmsKeyId" : String,
  "S3OutputConfiguration" : S3OutputConfiguration
}

```

### YAML

```yaml

  KmsKeyId: String
  S3OutputConfiguration:
    S3OutputConfiguration

```

## Properties

`KmsKeyId`

Property description not available.

_Required_: No

_Type_: String

_Pattern_: `^[A-Za-z0-9][A-Za-z0-9:_/+=,@.-]{0,2048}$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`S3OutputConfiguration`

Property description not available.

_Required_: Yes

_Type_: [S3OutputConfiguration](aws-properties-lookoutequipment-inferencescheduler-s3outputconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DataInputConfiguration

InputNameConfiguration

All content copied from https://docs.aws.amazon.com/.
