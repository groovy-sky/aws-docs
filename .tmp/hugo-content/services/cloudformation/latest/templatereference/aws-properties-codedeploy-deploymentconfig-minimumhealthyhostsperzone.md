This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeDeploy::DeploymentConfig MinimumHealthyHostsPerZone

Information about the minimum number of healthy instances per Availability
Zone.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : String,
  "Value" : Integer
}

```

### YAML

```yaml

  Type: String
  Value: Integer

```

## Properties

`Type`

The `type` associated with the `MinimumHealthyHostsPerZone`
option.

_Required_: Yes

_Type_: String

_Allowed values_: `HOST_COUNT | FLEET_PERCENT`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Value`

The `value` associated with the `MinimumHealthyHostsPerZone`
option.

_Required_: Yes

_Type_: Integer

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

MinimumHealthyHosts

TimeBasedCanary

All content copied from https://docs.aws.amazon.com/.
