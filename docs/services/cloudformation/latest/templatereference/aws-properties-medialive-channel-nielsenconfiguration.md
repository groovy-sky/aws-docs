This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaLive::Channel NielsenConfiguration

The settings to configure Nielsen watermarks.

The parent of this entity is EncoderSettings.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "DistributorId" : String,
  "NielsenPcmToId3Tagging" : String
}

```

### YAML

```yaml

  DistributorId: String
  NielsenPcmToId3Tagging: String

```

## Properties

`DistributorId`

Enter the Distributor ID assigned to your organization by Nielsen.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`NielsenPcmToId3Tagging`

Enables Nielsen PCM to ID3 tagging

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NielsenCBET

NielsenNaesIiNw

All content copied from https://docs.aws.amazon.com/.
