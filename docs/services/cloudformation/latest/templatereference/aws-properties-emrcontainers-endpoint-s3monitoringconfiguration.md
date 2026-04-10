This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EMRContainers::Endpoint S3MonitoringConfiguration

Amazon S3 configuration for monitoring log publishing. You can configure your jobs to
send log information to Amazon S3.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "LogUri" : String
}

```

### YAML

```yaml

  LogUri: String

```

## Properties

`LogUri`

Amazon S3 destination URI for log publishing.

_Required_: Yes

_Type_: String

_Pattern_: `[\u0020-\uD7FF\uE000-\uFFFD\uD800\uDBFF-\uDC00\uDFFF\r\n\t]*`

_Minimum_: `1`

_Maximum_: `10280`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MonitoringConfiguration

Tag

All content copied from https://docs.aws.amazon.com/.
