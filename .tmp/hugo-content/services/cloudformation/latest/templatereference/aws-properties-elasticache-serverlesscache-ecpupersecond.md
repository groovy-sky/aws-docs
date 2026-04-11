This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElastiCache::ServerlessCache ECPUPerSecond

The configuration for the number of ElastiCache Processing Units (ECPU) the cache can consume per second.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Maximum" : Integer,
  "Minimum" : Integer
}

```

### YAML

```yaml

  Maximum: Integer
  Minimum: Integer

```

## Properties

`Maximum`

The configuration for the maximum number of ECPUs the cache can consume per second.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Minimum`

The configuration for the minimum number of ECPUs the cache should be able consume per second.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

DataStorage

Endpoint

All content copied from https://docs.aws.amazon.com/.
