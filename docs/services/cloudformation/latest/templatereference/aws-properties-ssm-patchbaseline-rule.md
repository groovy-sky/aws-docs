This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SSM::PatchBaseline Rule

The `Rule` property type specifies an approval rule for a Systems Manager patch baseline.

The `PatchRules` property of the [RuleGroup](../userguide/aws-properties-ssm-patchbaseline-rulegroup.md) property type contains a list of `Rule` property
types.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "ApproveAfterDays" : Integer,
  "ApproveUntilDate" : String,
  "ComplianceLevel" : String,
  "EnableNonSecurity" : Boolean,
  "PatchFilterGroup" : PatchFilterGroup
}

```

### YAML

```yaml

  ApproveAfterDays: Integer
  ApproveUntilDate: String
  ComplianceLevel: String
  EnableNonSecurity: Boolean
  PatchFilterGroup:
    PatchFilterGroup

```

## Properties

`ApproveAfterDays`

The number of days after the release date of each patch matched by the rule that the patch
is marked as approved in the patch baseline. For example, a value of `7` means that
patches are approved seven days after they are released.

Patch Manager evaluates patch release dates using Coordinated Universal Time (UTC). If a
patch is released at `2025-11-09T18:00:00Z` and `ApproveAfterDays` is set to
`7`, the patch will be approved after `2025-11-16T18:00:00Z`.

This parameter is marked as `Required: No`, but your request must include a value
for either `ApproveAfterDays` or `ApproveUntilDate`.

Not supported for Debian Server or Ubuntu Server.

###### Important

Use caution when setting this value for Windows Server patch baselines. Because patch
updates that are replaced by later updates are removed, setting too broad a value for this
parameter can result in crucial patches not being installed. For more information, see the
**Windows Server** tab in the topic [How security\
patches are selected](../../../systems-manager/latest/userguide/patch-manager-selecting-patches.md) in the _AWS Systems Manager User Guide_.

_Required_: No

_Type_: Integer

_Minimum_: `0`

_Maximum_: `360`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ApproveUntilDate`

The cutoff date for auto approval of released patches. Any patches released on or before
this date are installed automatically.

Enter dates in the format `YYYY-MM-DD`. For example,
`2025-11-16`.

Patch Manager evaluates patch release dates using Coordinated Universal Time (UTC). If you
enter the date `2025-11-16`, patches released between
`2025-11-16T00:00:00Z` and `2025-11-16T23:59:59Z` will be included in the
approval.

This parameter is marked as `Required: No`, but your request must include a value
for either `ApproveUntilDate` or `ApproveAfterDays`.

Not supported for Debian Server or Ubuntu Server.

###### Important

Use caution when setting this value for Windows Server patch baselines. Because patch
updates that are replaced by later updates are removed, setting too broad a value for this
parameter can result in crucial patches not being installed. For more information, see the
**Windows Server** tab in the topic [How security\
patches are selected](../../../systems-manager/latest/userguide/patch-manager-selecting-patches.md) in the _AWS Systems Manager User Guide_.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `10`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ComplianceLevel`

A compliance severity level for all approved patches in a patch baseline. Valid
compliance severity levels include the following: `UNSPECIFIED`,
`CRITICAL`, `HIGH`, `MEDIUM`, `LOW`, and
`INFORMATIONAL`.

_Required_: No

_Type_: String

_Allowed values_: `CRITICAL | HIGH | INFORMATIONAL | LOW | MEDIUM | UNSPECIFIED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`EnableNonSecurity`

For managed nodes identified by the approval rule filters, enables a patch baseline to apply
non-security updates available in the specified repository. The default value is
`false`. Applies to Linux managed nodes only.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PatchFilterGroup`

The patch filter group that defines the criteria for the rule.

_Required_: No

_Type_: [PatchFilterGroup](aws-properties-ssm-patchbaseline-patchfiltergroup.md)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## See also

- [PatchRule](../../../../reference/systems-manager/latest/apireference/api-patchrule.md) in the _AWS Systems Manager API_
_Reference_.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

PatchSource

RuleGroup

All content copied from https://docs.aws.amazon.com/.
