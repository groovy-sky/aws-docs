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

_Type_: [S3LogConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-workspacesweb-sessionlogger-s3logconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

EventFilter

S3LogConfiguration
