This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lambda::EventSourceMapping ScalingConfig

(Amazon SQS only) The scaling configuration for the event source. To remove the configuration, pass an empty value.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "MaximumConcurrency" : Integer
}

```

### YAML

```yaml

  MaximumConcurrency: Integer

```

## Properties

`MaximumConcurrency`

Limits the number of concurrent instances that the Amazon SQS event source can invoke.

_Required_: No

_Type_: Integer

_Minimum_: `2`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ProvisionedPollerConfig

SchemaRegistryAccessConfig
