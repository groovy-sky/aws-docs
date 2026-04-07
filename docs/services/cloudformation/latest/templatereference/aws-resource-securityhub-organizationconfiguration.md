This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SecurityHub::OrganizationConfiguration

The `AWS::SecurityHub::OrganizationConfiguration` resource specifies the way that your AWS
organization is configured in AWS Security Hub CSPM. Specifically, you can use this resource to specify the configuration type
for your organization and whether to automatically Security Hub CSPM and security standards in new member accounts. For more information,
see [Managing administrator \
and member accounts](https://docs.aws.amazon.com/securityhub/latest/userguide/securityhub-accounts.html) in the _AWS Security Hub CSPM User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SecurityHub::OrganizationConfiguration",
  "Properties" : {
      "AutoEnable" : Boolean,
      "AutoEnableStandards" : String,
      "ConfigurationType" : String
    }
}

```

### YAML

```yaml

Type: AWS::SecurityHub::OrganizationConfiguration
Properties:
  AutoEnable: Boolean
  AutoEnableStandards: String
  ConfigurationType: String

```

## Properties

`AutoEnable`

Whether to automatically enable Security Hub CSPM in new member accounts when they join the organization.

If set to `true`, then Security Hub CSPM is automatically enabled in new accounts. If set to `false`,
then Security Hub CSPM isn't enabled in new accounts automatically. The default value is `false`.

If the `ConfigurationType` of your organization is set to `CENTRAL`, then this field is set
to `false` and can't be changed in the home Region and linked Regions. However, in that case, the delegated administrator can create a configuration
policy in which Security Hub CSPM is enabled and associate the policy with new organization accounts.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`AutoEnableStandards`

Whether to automatically enable Security Hub CSPM [default standards](https://docs.aws.amazon.com/securityhub/latest/userguide/securityhub-standards-enable-disable.html)
in new member accounts when they join the organization.

The default value of this parameter is equal to `DEFAULT`.

If equal to `DEFAULT`, then Security Hub CSPM default standards are automatically enabled for new member
accounts. If equal to `NONE`, then default standards are not automatically enabled for new member
accounts.

If the `ConfigurationType` of your organization is set to `CENTRAL`, then this field is set
to `NONE` and can't be changed in the home Region and linked Regions. However, in that case, the delegated administrator can create a configuration
policy in which specific security standards are enabled and associate the policy with new organization accounts.

_Required_: No

_Type_: String

_Allowed values_: `DEFAULT | NONE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ConfigurationType`

Indicates whether the organization uses local or central configuration.

If you use local configuration, the
Security Hub CSPM delegated administrator can set `AutoEnable` to `true` and
`AutoEnableStandards` to `DEFAULT`. This automatically enables Security Hub CSPM and
default security standards in new organization accounts. These new account settings must be set separately in
each AWS Region, and settings may be different in each Region.

If you use central configuration, the delegated administrator can create configuration policies. Configuration
policies can be used to configure Security Hub CSPM, security standards, and security controls in multiple
accounts and Regions. If you want new organization accounts to use a specific configuration, you can create a
configuration policy and associate it with the root or specific organizational units (OUs). New accounts will
inherit the policy from the root or their assigned OU.

_Required_: No

_Type_: String

_Allowed values_: `CENTRAL | LOCAL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the organization configuration identifier, formatted as `AccountId/Region/securityhub-organization-configuration`.
For example, `123456789012/us-east-1/securityhub-organization-configuration`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`MemberAccountLimitReached`

Whether the maximum number of allowed member accounts are already associated with the
Security Hub CSPM administrator account.

`OrganizationConfigurationIdentifier`

The organization configuration identifier, formatted as `AccountId/Region/securityhub-organization-configuration`.
For example, `123456789012/us-east-1/securityhub-organization-configuration`.

`Status`

Describes whether central configuration could be enabled as the `ConfigurationType` for the
organization. If your `ConfigurationType` is local configuration, then the value of `Status`
is always `ENABLED`.

`StatusMessage`

Provides an explanation if the value of `Status` is equal to `FAILED` when `ConfigurationType`
is equal to `CENTRAL`.

## Examples

### Configuring your organization in Security Hub CSPM

The following example configures organization settings in Security Hub CSPM.

#### JSON

```json

{
	"Description": "Example template to configure an organization in Security Hub",
	"Resources": {
		"SecurityHubOrganizationConfiguration": {
			"Type": "AWS::SecurityHub::OrganizationConfiguration",
			"Properties": {
				"AutoEnable": false,
				"AutoEnableStandards": "NONE",
				"ConfigurationType": "CENTRAL"
			}
		}
	}
}
```

#### YAML

```yaml

Description: Example template to configure an organization in Security Hub
Resources:
  SecurityHubOrganizationConfiguration:
    Type: 'AWS::SecurityHub::OrganizationConfiguration'
    Properties:
      AutoEnable: false
      AutoEnableStandards: "NONE"
      ConfigurationType: "CENTRAL"
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

StringFilter

AWS::SecurityHub::PolicyAssociation
