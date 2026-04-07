This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SecurityHub::ConfigurationPolicy

The `AWS::SecurityHub::ConfigurationPolicy` resource creates a central configuration policy with the defined settings.
Only the AWS Security Hub CSPM delegated administrator can create this resource in the home Region. For more information, see
[Central configuration in Security Hub CSPM](https://docs.aws.amazon.com/securityhub/latest/userguide/central-configuration-intro.html) in
the _AWS Security Hub CSPM User Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SecurityHub::ConfigurationPolicy",
  "Properties" : {
      "ConfigurationPolicy" : Policy,
      "Description" : String,
      "Name" : String,
      "Tags" : {Key: Value, ...}
    }
}

```

### YAML

```yaml

Type: AWS::SecurityHub::ConfigurationPolicy
Properties:
  ConfigurationPolicy:
    Policy
  Description: String
  Name: String
  Tags:
    Key: Value

```

## Properties

`ConfigurationPolicy`

An object that defines how AWS Security Hub CSPM is configured. It includes whether Security Hub CSPM is
enabled or disabled, a list of enabled security standards, a list of enabled or disabled security controls, and a list of custom parameter values for specified controls.
If you provide a list of security controls that are enabled in the configuration policy, Security Hub CSPM
disables all other controls (including newly released controls). If you provide a list of security controls that
are disabled in the configuration policy, Security Hub CSPM enables all other controls (including newly
released controls).

_Required_: Yes

_Type_: [Policy](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-securityhub-configurationpolicy-policy.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The description of the configuration policy.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `512`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the configuration policy. Alphanumeric characters and the following ASCII characters are permitted:
`-, ., !, *, /`.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `128`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

User-defined tags associated with a configuration policy. For more information, see
[Tagging AWS Security Hub CSPM resources](https://docs.aws.amazon.com/securityhub/latest/userguide/tagging-resources.html)
in the _Security Hub CSPM user guide_.

_Required_: No

_Type_: Object of String

_Pattern_: `^(?!aws:)[a-zA-Z+-=._:/]{1,128}$`

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the Amazon Resource Name (ARN) of the configuration policy. For example, `arn:aws:securityhub:us-east-1:123456789012:configuration-policy/a1b2c3d4-5678-90ab-cdef-EXAMPLE11111`.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The ARN of the configuration policy.

`CreatedAt`

Property description not available.

`Id`

The universally unique identifier (UUID) of the configuration policy. A self-managed configuration has no UUID. The
identifier of a self-managed configuration is `SELF_MANAGED_SECURITY_HUB`.

`ServiceEnabled`

Indicates whether the service that the configuration policy applies to is enabled in the policy.

`UpdatedAt`

The date and time, in UTC and ISO 8601 format, that the configuration policy was last updated.

## Examples

### Creating a Security Hub CSPM central configuration policy

The following example creates a configuration policy with the specified settings. Only the delegated
Security Hub CSPM administrator can create a configuration policy from the home Region.

#### JSON

```json

{
	"Description": "Example template to create a SecurityHub configuration policy",
	"Resources": {
		"SecurityHubConfigurationPolicy": {
			"Type": "AWS::SecurityHub::ConfigurationPolicy",
			"Properties": {
				"Tags": {
					"key1": "value1"
				},
				"Name": "SecurityHubConfigurationPolicyExample",
				"Description": "Example template to create SecurityHub Configuration Policy",
				"ConfigurationPolicy" : {
      				"SecurityHub": {
        				"ServiceEnabled": true,
        				"EnabledStandardIdentifiers": [
        					"arn:aws:securityhub:us-west-2::standards/aws-foundational-security-best-practices/v/1.0.0"
        				],
        				"SecurityControlsConfiguration": {
          					"EnabledSecurityControlIdentifiers": [
          						"APIGateway.1",
          						"IAM.7",
          						"RDS.14",
          						"CloudFront.5",
          						"EC2.18","S3.11",
          						"CloudFront.6"
          					],
          					"SecurityControlCustomParameters": [
            					{
				              		"SecurityControlId": "APIGateway.1",
				              		"Parameters": {
				                		"loggingLevel": {
				                			"ValueType": "CUSTOM",
				                			"Value": {
				                				"Enum": "ERROR"
				                			}
				                		}
              				  		}
              					}
          					]
        				}
      				}
				}
			}
		}
	}
}
```

#### YAML

```yaml

Description: Example template to create a Security Hub configuration policy
Resources:
  SecurityHubConfigurationPolicy:
    Type: "AWS::SecurityHub::ConfigurationPolicy"
    Properties:
      Tags:
        key1: value1
      Name: "SecurityHubConfigurationPolicyExample"
      Description: "Example template to create SecurityHub Configuration Policy"
      ConfigurationPolicy:
        SecurityHub:
          ServiceEnabled: true
          EnabledStandardIdentifiers:
            - !Sub "arn:${AWS::Partition}:securityhub:${AWS::Region}::standards/aws-foundational-security-best-practices/v/1.0.0"
          SecurityControlsConfiguration:
            EnabledSecurityControlIdentifiers:
              - "APIGateway.1"
              - "IAM.7"
              - "RDS.14"
              - "CloudFront.5"
              - "EC2.18"
              - "S3.11"
              - "CloudFront.6"
            SecurityControlCustomParameters:
              - SecurityControlId: "APIGateway.1"
                Parameters:
                  loggingLevel:
                    ValueType: "CUSTOM"
                    Value:
                      Enum: "ERROR"
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

StringFilter

ParameterConfiguration
