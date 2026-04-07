This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lambda::Function TracingConfig

The function's [AWS X-Ray](https://docs.aws.amazon.com/lambda/latest/dg/services-xray.html) tracing configuration.
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

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TenancyConfig

VpcConfig
