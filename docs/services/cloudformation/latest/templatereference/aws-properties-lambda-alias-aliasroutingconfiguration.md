This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lambda::Alias AliasRoutingConfiguration

The [traffic-shifting](../../../lambda/latest/dg/lambda-traffic-shifting-using-aliases.md) configuration of a Lambda function alias.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "AdditionalVersionWeights" : [ VersionWeight, ... ]
}

```

### YAML

```yaml

  AdditionalVersionWeights:
    - VersionWeight

```

## Properties

`AdditionalVersionWeights`

The second version, and the percentage of traffic that's routed to it.

_Required_: No

_Type_: Array of [VersionWeight](aws-properties-lambda-alias-versionweight.md)

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

AWS::Lambda::Alias

ProvisionedConcurrencyConfiguration

All content copied from https://docs.aws.amazon.com/.
