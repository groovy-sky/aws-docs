This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lambda::Version ProvisionedConcurrencyConfiguration

A [provisioned concurrency](../../../lambda/latest/dg/configuration-concurrency.md) configuration for a function's version.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ProvisionedConcurrentExecutions" : Integer
}

```

### YAML

```yaml

  ProvisionedConcurrentExecutions: Integer

```

## Properties

`ProvisionedConcurrentExecutions`

The amount of provisioned concurrency to allocate for the version.

_Required_: Yes

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Examples

### Provisioned Concurrency Configuration

Allocate 20 provisioned concurrency for a version.

#### YAML

```yaml

      ProvisionedConcurrencyConfig:
        ProvisionedConcurrentExecutions: 20
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

FunctionScalingConfig

RuntimePolicy

All content copied from https://docs.aws.amazon.com/.
