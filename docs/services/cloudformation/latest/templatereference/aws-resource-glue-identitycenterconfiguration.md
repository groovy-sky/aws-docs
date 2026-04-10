This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Glue::IdentityCenterConfiguration

Creates a new AWS Glue Identity Center configuration to enable integration between AWS Glue and AWS IAM
Identity Center for authentication and authorization.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Glue::IdentityCenterConfiguration",
  "Properties" : {
      "InstanceArn" : String,
      "Scopes" : [ String, ... ],
      "UserBackgroundSessionsEnabled" : Boolean
    }
}

```

### YAML

```yaml

Type: AWS::Glue::IdentityCenterConfiguration
Properties:
  InstanceArn: String
  Scopes:
    - String
  UserBackgroundSessionsEnabled: Boolean

```

## Properties

`InstanceArn`

The Amazon Resource Name (ARN) of the Identity Center instance associated with the AWS Glue configuration.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:(aws|aws-us-gov|aws-cn|aws-iso|aws-iso-b):sso:::instance/(sso)?ins-[a-zA-Z0-9-.]{16}$`

_Minimum_: `10`

_Maximum_: `1224`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Scopes`

A list of Identity Center scopes that define the permissions and access levels for the AWS Glue configuration.

_Required_: No

_Type_: Array of String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UserBackgroundSessionsEnabled`

Indicates whether users can run background sessions when using Identity Center authentication with AWS Glue services.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`AccountId`

Property description not available.

`ApplicationArn`

The Amazon Resource Name (ARN) of the Identity Center application associated with the AWS Glue configuration.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS::Glue::DevEndpoint

AWS::Glue::Integration

All content copied from https://docs.aws.amazon.com/.
