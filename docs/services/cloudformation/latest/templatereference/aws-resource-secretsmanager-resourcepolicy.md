This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SecretsManager::ResourcePolicy

Attaches a resource-based permission policy to a secret. A resource-based policy is
optional. If a secret already has a resource policy attached, you must first remove it before attaching a new policy using this CloudFormation resource. You can remove the policy using the [console](../../../secretsmanager/latest/userguide/auth-and-access-resource-policies.md), [CLI](../../../cli/latest/reference/secretsmanager/delete-resource-policy.md), or [API](../../../../reference/secretsmanager/latest/apireference/api-deleteresourcepolicy.md). For more information, see [Authentication and access control\
for Secrets Manager](../../../secretsmanager/latest/userguide/auth-and-access.md).

**Required permissions:** `secretsmanager:PutResourcePolicy`, `secretsmanager:GetResourcePolicy`. For more information, see [IAM policy actions for Secrets Manager](../../../service-authorization/latest/reference/list-awssecretsmanager.md#awssecretsmanager-actions-as-permissions) and [Authentication and access control\
in Secrets Manager](../../../secretsmanager/latest/userguide/auth-and-access.md).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SecretsManager::ResourcePolicy",
  "Properties" : {
      "BlockPublicPolicy" : Boolean,
      "ResourcePolicy" : Json,
      "SecretId" : String
    }
}

```

### YAML

```yaml

Type: AWS::SecretsManager::ResourcePolicy
Properties:
  BlockPublicPolicy: Boolean
  ResourcePolicy: Json
  SecretId: String

```

## Properties

`BlockPublicPolicy`

Specifies whether to block resource-based policies that allow broad access to the secret.
By default, Secrets Manager blocks policies that allow broad access, for example those that
use a wildcard for the principal.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourcePolicy`

A JSON-formatted string for an AWS
resource-based policy. For example policies, see [Permissions \
policy examples](../../../secretsmanager/latest/userguide/auth-and-access-examples.md).

_Required_: Yes

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SecretId`

The ARN or name of the secret to attach the resource-based policy.

For an ARN, we recommend that you specify a complete ARN rather
than a partial ARN.

_Required_: Yes

_Type_: String

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of an `AWS::SecretsManager::ResourcePolicy`
resource to the intrinsic `Ref` function, the function returns the ARN of the
configured secret, such as:

`arn:aws:secretsmanager:us-west-2:123456789012:secret:my-path/my-secret-name-1a2b3c`

This enables you to reference a secret you created in one part of the stack template from
within the definition of another resource later, in the same template. You would typically use
this with the [AWS::SecretsManager::SecretTargetAttachment](../userguide/aws-resource-secretsmanager-secrettargetattachment.md) resource type.

For more information about using the Ref function, see [Ref](../userguide/intrinsic-function-reference-ref.md).

### Fn::GetAtt

## Examples

### Attaching a resource-based policy to an RDS database instance secret

The following example shows how to attach a resource-based policy to a secret. The
JSON request string input and response output displays as formatted with white space and
line breaks for better readability. Submit your input as a single line JSON string.

#### JSON

```json

{
  "MySecret": {
    "Type": "AWS::SecretsManager::Secret",
    "Properties": {
      "Description": "This is a secret that I want to attach a resource-based policy to"
    }
  },
  "MySecretResourcePolicy": {
    "Type": "AWS::SecretsManager::ResourcePolicy",
    "Properties": {
      "BlockPublicPolicy": "True",
      "SecretId": {
        "Ref": "MySecret"
      },
      "ResourcePolicy": {
        "Version": "2012-10-17",
        "Statement": [
        {
          "Resource": "*",
          "Action": "secretsmanager:DeleteSecret",
          "Effect": "Deny",
          "Principal": {
            "AWS": {
              "Fn::Sub": "arn:aws:iam::${AWS::AccountId}:root"
            }
          }
        }
        ]
      }
    }
  }
}

```

#### YAML

```yaml

---
MySecret:
  Type: AWS::SecretsManager::Secret
  Properties:
        Description: This is a secret that I want to attach a resource-based policy to
MySecretResourcePolicy:
  Type: AWS::SecretsManager::ResourcePolicy
  Properties:
    BlockPublicPolicy: True
    SecretId:
      Ref: MySecret
    ResourcePolicy:
      Version: '2012-10-17'
      Statement:
      - Resource: "*"
        Action: secretsmanager:DeleteSecret
        Effect: Deny
        Principal:
          AWS:
            Fn::Sub: arn:aws:iam::${AWS::AccountId}:root

```

## See also

- [AWS::SecretsManager::Secret](../userguide/aws-resource-secretsmanager-secret.md)

- [AWS::SecretsManager::RotationSchedule](../userguide/aws-resource-secretsmanager-rotationschedule.md)

- [AWS::SecretsManager::SecretTargetAttachment](../userguide/aws-resource-secretsmanager-secrettargetattachment.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS Secrets Manager

AWS::SecretsManager::RotationSchedule

All content copied from https://docs.aws.amazon.com/.
