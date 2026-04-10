This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SecurityHub::PolicyAssociation

The `AWS::SecurityHub::PolicyAssociation` resource specifies associations for a configuration policy or
a self-managed configuration. You can associate a AWS Security Hub CSPM configuration policy or self-managed configuration with
the organization root, organizational units (OUs), or AWS accounts. After a successful association, the
configuration policy takes effect in the specified targets. For more information, see
[Creating and associating \
Security Hub CSPM configuration policies](../../../securityhub/latest/userguide/create-associate-policy.md) in the _AWS Security Hub CSPM User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SecurityHub::PolicyAssociation",
  "Properties" : {
      "ConfigurationPolicyId" : String,
      "TargetId" : String,
      "TargetType" : String
    }
}

```

### YAML

```yaml

Type: AWS::SecurityHub::PolicyAssociation
Properties:
  ConfigurationPolicyId: String
  TargetId: String
  TargetType: String

```

## Properties

`ConfigurationPolicyId`

The universally unique identifier (UUID) of the configuration policy. A self-managed configuration has no UUID. The
identifier of a self-managed configuration is `SELF_MANAGED_SECURITY_HUB`.

_Required_: Yes

_Type_: String

_Pattern_: `^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$|^SELF_MANAGED_SECURITY_HUB$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`TargetId`

The identifier of the target account, organizational unit, or the root.

_Required_: Yes

_Type_: String

_Pattern_: `.*\S.*`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`TargetType`

Specifies whether the target is an AWS account, organizational unit, or the root.

_Required_: Yes

_Type_: String

_Allowed values_: `ACCOUNT | ORGANIZATIONAL_UNIT | ROOT`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the association identifier, formatted as `TargetType/TargetId`. For example,
`ACCOUNT/123456789012`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AssociationIdentifier`

The association identifier, formatted as `TargetType/TargetId`. For example,
`ACCOUNT/123456789012`.

`AssociationStatus`

The current status of the association between the specified target and the configuration.

`AssociationStatusMessage`

The explanation for a `FAILED` value for `AssociationStatus`.

`AssociationType`

Indicates whether the association between the specified target and the configuration was directly applied by the
AWS Security Hub CSPM delegated administrator or inherited from a parent.

`UpdatedAt`

The date and time, in UTC and ISO 8601 format, that the configuration policy association was last updated.

## Examples

### Association a configuration policy or self-managed configuration

The following example associates the specified Security Hub CSPM configuration policy with the specified account.

#### JSON

```json

{
	"Description": "Example template to associate a Security Hub configuration policy or self-managed configuration",
	"Resources": {
		"SecurityHubPolicyAssociation": {
			"Type": "AWS::SecurityHub::PolicyAssociation",
			"Properties": {
				"TargetType": "ACCOUNT",
				"TargetId": "123456789012",
				"ConfigurationPolicyId": "a1b2c3d4-5678-90ab-cdef-EXAMPLE11111"
			}
		}
	}
}
```

#### YAML

```yaml

Description: Example template to associate a SecurityHub configuration policy or self-managed configuration
Resources:
  SecurityHubPolicyAssociation:
    Type: "AWS::SecurityHub::PolicyAssociation"
    Properties:
      TargetType: "ACCOUNT"
      TargetId: "123456789012"
      ConfigurationPolicyId: "a1b2c3d4-5678-90ab-cdef-EXAMPLE11111"
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::SecurityHub::OrganizationConfiguration

AWS::SecurityHub::ProductSubscription

All content copied from https://docs.aws.amazon.com/.
