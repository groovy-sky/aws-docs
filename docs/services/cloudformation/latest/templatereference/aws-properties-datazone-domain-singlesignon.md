This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::DataZone::Domain SingleSignOn

The single sign-on details in Amazon DataZone.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "IdcInstanceArn" : String,
  "Type" : String,
  "UserAssignment" : String
}

```

### YAML

```yaml

  IdcInstanceArn: String
  Type: String
  UserAssignment: String

```

## Properties

`IdcInstanceArn`

The ARN of the IDC instance.

_Required_: No

_Type_: String

_Pattern_: `arn:(aws|aws-us-gov|aws-cn|aws-iso|aws-iso-b):sso:::instance/(sso)?ins-[a-zA-Z0-9-.]{16}`

_Minimum_: `10`

_Maximum_: `1224`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Type`

The type of single sign-on in Amazon DataZone.

_Required_: No

_Type_: String

_Allowed values_: `IAM_IDC | DISABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserAssignment`

The single sign-on user assignment in Amazon DataZone.

_Required_: No

_Type_: String

_Allowed values_: `AUTOMATIC | MANUAL`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::DataZone::Domain

Tag
