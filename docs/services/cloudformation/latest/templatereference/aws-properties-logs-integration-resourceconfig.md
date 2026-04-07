This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Logs::Integration ResourceConfig

This structure contains configuration details about an integration between CloudWatch
Logs and another entity.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "OpenSearchResourceConfig" : OpenSearchResourceConfig
}

```

### YAML

```yaml

  OpenSearchResourceConfig:
    OpenSearchResourceConfig

```

## Properties

`OpenSearchResourceConfig`

This structure contains configuration details about an integration between CloudWatch
Logs and OpenSearch Service.

_Required_: No

_Type_: [OpenSearchResourceConfig](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-logs-integration-opensearchresourceconfig.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

OpenSearchResourceConfig

AWS::Logs::LogAnomalyDetector
