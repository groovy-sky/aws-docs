This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Synthetics::Canary Dependency

A structure that contains information about a dependency for a canary.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Reference" : String,
  "Type" : String
}

```

### YAML

```yaml

  Reference: String
  Type: String

```

## Properties

`Reference`

The dependency reference. For Lambda layers, this is the ARN of the Lambda layer. For more information
about Lambda ARN format, see [Lambda](../../../lambda/latest/api/api-layer.md).

_Required_: Yes

_Type_: String

_Pattern_: `arn:[a-zA-Z0-9-]+:lambda:[a-zA-Z0-9-]+:\d{12}:layer:[a-zA-Z0-9-_]+:[0-9]+`

_Minimum_: `1`

_Maximum_: `140`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of dependency. Valid value is `LambdaLayer`.

_Required_: No

_Type_: String

_Allowed values_: `LambdaLayer`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Code

RetryConfig

All content copied from https://docs.aws.amazon.com/.
