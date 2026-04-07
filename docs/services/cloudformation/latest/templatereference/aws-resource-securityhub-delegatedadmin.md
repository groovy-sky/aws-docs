This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SecurityHub::DelegatedAdmin

The `AWS::SecurityHub::DelegatedAdmin` resource designates the delegated AWS Security Hub CSPM
administrator account for an organization. You must enable the integration between Security Hub CSPM and
AWS Organizations before you can designate a delegated Security Hub CSPM administrator. Only the management account for
an organization can designate the delegated Security Hub CSPM administrator account. For more information, see
[Designating the delegated Security Hub CSPM administrator](https://docs.aws.amazon.com/securityhub/latest/userguide/designate-orgs-admin-account.html#designate-admin-instructions) in
the _AWS Security Hub CSPM User Guide_.

To change the delegated administrator account, remove the current delegated administrator account, and then
designate the new account.

To designate multiple delegated administrators in different organizations and AWS Regions, we
recommend using [AWS CloudFormation mappings](../userguide/mappings-section-structure.md).

Tags aren't supported for this resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SecurityHub::DelegatedAdmin",
  "Properties" : {
      "AdminAccountId" : String
    }
}

```

### YAML

```yaml

Type: AWS::SecurityHub::DelegatedAdmin
Properties:
  AdminAccountId: String

```

## Properties

`AdminAccountId`

The AWS account identifier of the account to designate as the Security Hub CSPM administrator
account.

_Required_: Yes

_Type_: String

_Pattern_: `^[0-9]{12}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the ID of the delegated Security Hub CSPM administrator account. The format is
`accountID/Region`. For example, `123456789012/us-west-2`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`DelegatedAdminIdentifier`

The ID of the delegated Security Hub CSPM administrator account, in the format of `accountID/Region`.

`Status`

Whether the delegated Security Hub CSPM administrator is set for the organization.

## Examples

### Designating the delegated Security Hub CSPM administrator

The following example designates the specified AWS account as the delegated Security Hub CSPM administrator for an organization.

#### JSON

```json

{
	"Description": "Example template to create the delegated Security Hub administrator",
	"Resources": {
		"SecurityHubDelegatedAdmin": {
			"Type": "AWS::SecurityHub::DelegatedAdmin",
			"Properties": {
				"AdminAccountId": "123456789012"
			}
		}
	}
}
```

#### YAML

```yaml

Description: Example template to create the delegated Security Hub administrator
Resources:
  SecurityHubDelegatedAdmin:
    Type: 'AWS::SecurityHub::DelegatedAdmin'
    Properties:
      AdminAccountId: '123456789012'
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

ServiceNowProviderConfiguration

AWS::SecurityHub::FindingAggregator
