This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lambda::CapacityProvider CapacityProviderPermissionsConfig

Configuration that specifies the permissions required for the capacity provider to manage compute resources.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "CapacityProviderOperatorRoleArn" : String
}

```

### YAML

```yaml

  CapacityProviderOperatorRoleArn: String

```

## Properties

`CapacityProviderOperatorRoleArn`

The ARN of the IAM role that the capacity provider uses to manage compute instances and other AWS resources.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:(aws[a-zA-Z-]*)?:iam::\d{12}:role/?[a-zA-Z_0-9+=,.@\-_/]+$`

_Minimum_: `0`

_Maximum_: `10000`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Lambda::CapacityProvider

CapacityProviderScalingConfig

All content copied from https://docs.aws.amazon.com/.
