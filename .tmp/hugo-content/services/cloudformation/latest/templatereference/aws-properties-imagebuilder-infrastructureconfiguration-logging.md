This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ImageBuilder::InfrastructureConfiguration Logging

Logging configuration defines where Image Builder uploads your logs.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "S3Logs" : S3Logs
}

```

### YAML

```yaml

  S3Logs:
    S3Logs

```

## Properties

`S3Logs`

The Amazon S3 logging configuration.

_Required_: No

_Type_: [S3Logs](aws-properties-imagebuilder-infrastructureconfiguration-s3logs.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

InstanceMetadataOptions

Placement

All content copied from https://docs.aws.amazon.com/.
