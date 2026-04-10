This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ElastiCache::ServerlessCache DataStorage

The data storage limit.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Maximum" : Integer,
  "Minimum" : Integer,
  "Unit" : String
}

```

### YAML

```yaml

  Maximum: Integer
  Minimum: Integer
  Unit: String

```

## Properties

`Maximum`

The upper limit for data storage the cache is set to use.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Minimum`

The lower limit for data storage the cache is set to use.

_Required_: No

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Unit`

The unit that the storage is measured in, in GB.

_Required_: Yes

_Type_: String

_Allowed values_: `GB`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

CacheUsageLimits

ECPUPerSecond

All content copied from https://docs.aws.amazon.com/.
