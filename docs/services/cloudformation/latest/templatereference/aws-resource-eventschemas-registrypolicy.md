This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::EventSchemas::RegistryPolicy

Use the `AWS::EventSchemas::RegistryPolicy` resource to specify
resource-based policies for an EventBridge Schema Registry.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::EventSchemas::RegistryPolicy",
  "Properties" : {
      "Policy" : Json,
      "RegistryName" : String,
      "RevisionId" : String
    }
}

```

### YAML

```yaml

Type: AWS::EventSchemas::RegistryPolicy
Properties:
  Policy: Json
  RegistryName: String
  RevisionId: String

```

## Properties

`Policy`

A resource-based policy.

_Required_: Yes

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RegistryName`

The name of the registry.

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`RevisionId`

The revision ID of the policy.

_Required_: No

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you provide the logical ID of this resource to the `Ref` intrinsic
function, `Ref` the name of the registry.

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Id`

The ID of the policy.

## Examples

#### YAML

```yaml

Resources:
  RegistryPolicy:
    Type: AWS::EventSchemas::RegistryPolicy
    Properties:
      RegistryName: registryName
      Policy:
        Version: 2012-10-17
        Statement:
          Sid: 1
          Effect: Allow
          Principal:
            AWS: arn:aws:iam::012345678901:user/TestAccountForRegistryPolicy
          Action:
            - schemas:DescribeRegistry
            - schemas:CreateSchema
          Resource: registryArn
```

#### YAML

```yaml

Resources:
  RegistryPolicy:
    Type: 'AWS::EventSchemas::RegistryPolicy'
    Properties:
      RegistryName: 'MyRegistry'
      Policy:
        Version: '2012-10-17'
        Statement:
          - Sid: 'Test'
            Effect: 'Allow'
            Action:
              - 'schemas:*'
            Principal:
              AWS:
                - '109876543210'
            Resource:
              - 'arn:aws:schemas:us-east-1:012345678901:registry/MyRegistry'
              - 'arn:aws:schemas:us-east-1:012345678901:schema/MyRegistry*'
```

#### JSON

```json

{
  "Resources": {
    "RegistryPolicy": {
      "Type": "AWS::EventSchemas::RegistryPolicy",
      "Properties": {
        "RegistryName": "MyRegistry",
        "Policy": {
          "Version": "2012-10-17",
          "Statement": [
            {
              "Sid": "Test",
              "Effect": "Allow",
              "Action": [
                "schemas:*"
              ],
              "Principal": {
                "AWS": [
                  "109876543210"
                ]
              },
              "Resource": [
                "arn:aws:schemas:us-east-1:012345678901:registry/MyRegistry",
                "arn:aws:schemas:us-east-1:012345678901:schema/MyRegistry*"
              ]
            }
          ]
        }
      }
    }
  }
}
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TagsEntry

AWS::EventSchemas::Schema
