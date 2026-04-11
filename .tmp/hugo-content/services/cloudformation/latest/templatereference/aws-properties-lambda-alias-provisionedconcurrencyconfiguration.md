This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lambda::Alias ProvisionedConcurrencyConfiguration

A provisioned concurrency configuration for a function's alias.

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

The amount of provisioned concurrency to allocate for the alias.

_Required_: Yes

_Type_: Integer

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

### Provisioned Concurrency

An alias with 20 provisioned concurrency.

#### YAML

```yaml

  alias:
    Type: AWS::Lambda::Alias
    Properties:
      FunctionName: !Ref function
      FunctionVersion: !GetAtt newVersion.Version
      Name: BLUE
      ProvisionedConcurrencyConfig:
        ProvisionedConcurrentExecutions: 20
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AliasRoutingConfiguration

VersionWeight

All content copied from https://docs.aws.amazon.com/.
