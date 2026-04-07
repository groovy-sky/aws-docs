This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Connect::IntegrationAssociation

Specifies the association of an AWS resource such as Lex bot (both v1
and v2) and Lambda function with an Amazon Connect instance.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Connect::IntegrationAssociation",
  "Properties" : {
      "InstanceId" : String,
      "IntegrationArn" : String,
      "IntegrationType" : String
    }
}

```

### YAML

```yaml

Type: AWS::Connect::IntegrationAssociation
Properties:
  InstanceId: String
  IntegrationArn: String
  IntegrationType: String

```

## Properties

`InstanceId`

The Amazon Resource Name (ARN) of the instance.

_Minimum_: `1`

_Maximum_: `100`

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws[-a-z0-9]*:connect:[-a-z0-9]*:[0-9]{12}:instance/[-a-zA-Z0-9]*$`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`IntegrationArn`

ARN of the integration being associated with the instance.

_Minimum_: `1`

_Maximum_: `140`

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `140`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`IntegrationType`

Specifies the integration type to be associated with the instance.

_Allowed Values_: `LEX_BOT` \|
`LAMBDA_FUNCTION`

_Required_: Yes

_Type_: String

_Allowed values_: `LEX_BOT | LAMBDA_FUNCTION | APPLICATION | CASES_DOMAIN`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`IntegrationAssociationId`

Identifier of the association with an Amazon Connect instance.

## Examples

- [Specify a Lex V1 Bot integration](#aws-resource-connect-integrationassociation--examples--Specify_a_Lex_V1_Bot_integration)

- [Specify a Lex V2 Bot integration](#aws-resource-connect-integrationassociation--examples--Specify_a_Lex_V2_Bot_integration)

- [Specify a Lambda Function integration](#aws-resource-connect-integrationassociation--examples--Specify_a_Lambda_Function_integration)

### Specify a Lex V1 Bot integration

The following example specifies a Lex V1 Bot integration for an Amazon Connect instance.

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Description: Specifies a Lex V1 Bot integration for an Amazon Connect instance
Resources:
  IntegrationAssociation:
    Type: AWS::Connect::IntegrationAssociation
    Properties:
      InstanceId: arn:aws:connect:region-name:aws-account-id:instance/instance-id
      IntegrationType: LEX_BOT
      IntegrationArn: arn:aws:lex:region-name:aws-account-id:bot/bot-name
```

### Specify a Lex V2 Bot integration

The following example specifies a Lex V2 Bot integration for an Amazon Connect instance.

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Description: Specifies a Lex V2 Bot integration for an Amazon Connect instance
Resources:
  IntegrationAssociation:
    Type: AWS::Connect::IntegrationAssociation
    Properties:
      InstanceId: arn:aws:connect:region-name:aws-account-id:instance/instance-id
      IntegrationType: LEX_BOT
      IntegrationArn: arn:aws:lex:region-name:aws-account-id:bot-alias/bot-id/alias-id
```

### Specify a Lambda Function integration

The following example specifies a Lambda Function for an Amazon Connect
instance.

#### YAML

```yaml

AWSTemplateFormatVersion: 2010-09-09
Description: Specifies a Lambda Function integration for an Amazon Connect instance
Resources:
  IntegrationAssociation:
    Type: AWS::Connect::IntegrationAssociation
    Properties:
      InstanceId: arn:aws:connect:region-name:aws-account-id:instance/instance-id
      IntegrationType: LAMBDA_FUNCTION
      IntegrationArn: arn:aws:lambda:region-name:aws-account-id:function:function-arn
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

S3Config

AWS::Connect::Notification
