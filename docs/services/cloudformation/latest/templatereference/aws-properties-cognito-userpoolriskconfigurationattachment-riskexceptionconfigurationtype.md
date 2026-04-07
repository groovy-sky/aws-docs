This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Cognito::UserPoolRiskConfigurationAttachment RiskExceptionConfigurationType

Exceptions to the risk evaluation configuration, including always-allow and
always-block IP address ranges.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "BlockedIPRangeList" : [ String, ... ],
  "SkippedIPRangeList" : [ String, ... ]
}

```

### YAML

```yaml

  BlockedIPRangeList:
    - String
  SkippedIPRangeList:
    - String

```

## Properties

`BlockedIPRangeList`

An always-block IP address list. Overrides the risk decision and always blocks
authentication requests. This parameter is displayed and set in CIDR notation.

_Required_: No

_Type_: Array of String

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SkippedIPRangeList`

An always-allow IP address list. Risk detection isn't performed on the IP addresses in
this range list. This parameter is displayed and set in CIDR notation.

_Required_: No

_Type_: Array of String

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

NotifyEmailType

AWS::Cognito::UserPoolUICustomizationAttachment
