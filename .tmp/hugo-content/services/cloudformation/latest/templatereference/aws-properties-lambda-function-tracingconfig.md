This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lambda::Function TracingConfig

The function's [AWS X-Ray](../../../lambda/latest/dg/services-xray.md) tracing configuration.
To sample and record incoming requests, set `Mode` to `Active`.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Mode" : String
}

```

### YAML

```yaml

  Mode: String

```

## Properties

`Mode`

The tracing mode.

_Required_: No

_Type_: String

_Allowed values_: `Active | PassThrough`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

### Tracing Configuration

Enable active tracing on a function.

#### YAML

```yaml

      TracingConfig:
        Mode: Active
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TenancyConfig

VpcConfig

All content copied from https://docs.aws.amazon.com/.
