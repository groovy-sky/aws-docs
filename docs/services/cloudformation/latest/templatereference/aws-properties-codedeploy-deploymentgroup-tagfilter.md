This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CodeDeploy::DeploymentGroup TagFilter

`TagFilter` is a property type of the [AWS::CodeDeploy::DeploymentGroup](../userguide/aws-resource-codedeploy-deploymentgroup.md) resource that specifies which on-premises
instances to associate with the deployment group. To register on-premise instances with
AWS CodeDeploy, see [Configure Existing On-Premises\
Instances by Using AWS CodeDeploy](../../../codedeploy/latest/userguide/instances-on-premises.md) in the _AWS CodeDeploy User Guide_.

For more information about using tags and tag groups to help manage your Amazon EC2 instances and on-premises instances, see [Tagging Instances for Deployment\
Groups in AWS CodeDeploy](../../../codedeploy/latest/userguide/instances-tagging.md) in the _AWS CodeDeploy User_
_Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Key" : String,
  "Type" : String,
  "Value" : String
}

```

### YAML

```yaml

  Key: String
  Type: String
  Value: String

```

## Properties

`Key`

The on-premises instance tag filter key.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The on-premises instance tag filter type:

- KEY\_ONLY: Key only.

- VALUE\_ONLY: Value only.

- KEY\_AND\_VALUE: Key and value.

_Required_: No

_Type_: String

_Allowed values_: `KEY_ONLY | VALUE_ONLY | KEY_AND_VALUE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Value`

The on-premises instance tag filter value.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

TargetGroupInfo

All content copied from https://docs.aws.amazon.com/.
