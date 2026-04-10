This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lambda::Alias VersionWeight

The [traffic-shifting](../../../lambda/latest/dg/lambda-traffic-shifting-using-aliases.md) configuration of a Lambda function alias.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "FunctionVersion" : String,
  "FunctionWeight" : Number
}

```

### YAML

```yaml

  FunctionVersion: String
  FunctionWeight: Number

```

## Properties

`FunctionVersion`

The qualifier of the second version.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FunctionWeight`

The percentage of traffic that the alias routes to the second version.

_Required_: Yes

_Type_: Number

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Examples

### Routing Configuration

An alias that routes half of incoming requests to a second version.

#### YAML

```yaml

  alias:
    Type: AWS::Lambda::Alias
    Properties:
      FunctionName: !Ref function
      FunctionVersion: !GetAtt newVersion.Version
      Name: BLUE
      RoutingConfig:
        AdditionalVersionWeights:
          - FunctionVersion: !GetAtt version.Version
            FunctionWeight: 0.5
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

ProvisionedConcurrencyConfiguration

AWS::Lambda::CapacityProvider

All content copied from https://docs.aws.amazon.com/.
