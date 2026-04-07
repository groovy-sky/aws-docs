This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Wisdom::MessageTemplate GroupingConfiguration

The configuration information of the grouping of Amazon Q in Connect users.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Criteria" : String,
  "Values" : [ String, ... ]
}

```

### YAML

```yaml

  Criteria: String
  Values:
    - String

```

## Properties

`Criteria`

The criteria used for grouping Amazon Q in Connect users.

The following is the list of supported criteria values.

- `RoutingProfileArn`: Grouping the users by their [Amazon Connect routing profile ARN](https://docs.aws.amazon.com/connect/latest/APIReference/API_RoutingProfile.html). User should
have [SearchRoutingProfile](https://docs.aws.amazon.com/connect/latest/APIReference/API_SearchRoutingProfiles.html) and [DescribeRoutingProfile](https://docs.aws.amazon.com/connect/latest/APIReference/API_DescribeRoutingProfile.html) permissions when
setting criteria to this value.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Values`

The list of values that define different groups of Amazon Q in Connect users.

- When setting `criteria` to `RoutingProfileArn`, you need to provide a list of ARNs of
[Amazon Connect routing\
profiles](https://docs.aws.amazon.com/connect/latest/APIReference/API_RoutingProfile.html) as values of this parameter.

_Required_: Yes

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

EmailMessageTemplateHeader

MessageTemplateAttachment
