This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::WorkSpacesWeb::SessionLogger LogConfiguration

The configuration of the log.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "S3" : S3LogConfiguration
}

```

### YAML

```yaml

  S3:
    S3LogConfiguration

```

## Properties

`S3`

The configuration for delivering the logs to S3.

_Required_: No

_Type_: [S3LogConfiguration](aws-properties-workspacesweb-sessionlogger-s3logconfiguration.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

EventFilter

S3LogConfiguration

All content copied from https://docs.aws.amazon.com/.
