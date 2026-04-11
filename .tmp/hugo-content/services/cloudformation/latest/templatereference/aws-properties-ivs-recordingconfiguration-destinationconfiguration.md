This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::IVS::RecordingConfiguration DestinationConfiguration

The DestinationConfiguration property type describes the location where recorded
videos will be stored. Each member represents a type of destination configuration. For
recording, you define one and only one type of destination configuration.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "S3" : S3DestinationConfiguration
}

```

### YAML

```yaml

  S3:
    S3DestinationConfiguration

```

## Properties

`S3`

An S3 destination configuration where recorded videos will be stored. See the [S3DestinationConfiguration](../userguide/aws-properties-ivs-recordingconfiguration-s3destinationconfiguration.md) property type for more information.

_Required_: No

_Type_: [S3DestinationConfiguration](aws-properties-ivs-recordingconfiguration-s3destinationconfiguration.md)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::IVS::RecordingConfiguration

RenditionConfiguration

All content copied from https://docs.aws.amazon.com/.
