This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::MediaConnect::Flow NdiSourceSettings

The settings for the NDI® source. This includes the exact name of the upstream NDI sender that you want to connect to your source.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "SourceName" : String
}

```

### YAML

```yaml

  SourceName: String

```

## Properties

`SourceName`

The exact name of an existing NDI sender that's registered with your discovery server. If included, the format of this name must be `MACHINENAME (ProgramName)`.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

NdiDiscoveryServerConfig

SecretsManagerEncryptionKeyConfiguration

All content copied from https://docs.aws.amazon.com/.
