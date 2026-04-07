This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Input RouterSettings

The `RouterSettings` property type specifies Property description not available. for an [AWS::MediaLive::Input](aws-resource-medialive-input.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Destinations" : [ RouterDestinationSettings, ... ],
  "EncryptionType" : String,
  "SecretArn" : String
}

```

### YAML

```yaml

  Destinations:
    - RouterDestinationSettings
  EncryptionType: String
  SecretArn: String

```

## Properties

`Destinations`

Property description not available.

_Required_: No

_Type_: Array of [RouterDestinationSettings](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-medialive-input-routerdestinationsettings.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`EncryptionType`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`SecretArn`

Property description not available.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

RouterDestinationSettings

Smpte2110ReceiverGroup
