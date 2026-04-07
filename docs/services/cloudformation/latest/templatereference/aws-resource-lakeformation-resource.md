This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::LakeFormation::Resource

The `AWS::LakeFormation::Resource` represents the data ( buckets and folders) that is being registered with AWS Lake Formation.
During a stack operation, AWS CloudFormation calls the AWS Lake Formation [`RegisterResource`](https://docs.aws.amazon.com/lake-formation/latest/dg/aws-lake-formation-api-credential-vending.html#aws-lake-formation-api-credential-vending-RegisterResource) API operation to register the resource.
To remove a `Resource` type, AWS CloudFormation calls the AWS Lake Formation [`DeregisterResource`](https://docs.aws.amazon.com/lake-formation/latest/dg/aws-lake-formation-api-credential-vending.html#aws-lake-formation-api-credential-vending-DeregisterResource) API operation.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::LakeFormation::Resource",
  "Properties" : {
      "HybridAccessEnabled" : Boolean,
      "ResourceArn" : String,
      "RoleArn" : String,
      "UseServiceLinkedRole" : Boolean,
      "WithFederation" : Boolean
    }
}

```

### YAML

```yaml

Type: AWS::LakeFormation::Resource
Properties:
  HybridAccessEnabled: Boolean
  ResourceArn: String
  RoleArn: String
  UseServiceLinkedRole: Boolean
  WithFederation: Boolean

```

## Properties

`HybridAccessEnabled`

Indicates whether the data access of tables pointing to the location can be managed by both Lake Formation permissions as well as Amazon S3 bucket policies.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceArn`

The Amazon Resource Name (ARN) of the resource.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`RoleArn`

The IAM role that registered a resource.

_Required_: No

_Type_: String

_Pattern_: `arn:aws:iam::[0-9]*:role/.*`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`UseServiceLinkedRole`

Designates a trusted caller, an IAM principal, by registering this caller with the Data Catalog.

_Required_: Yes

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`WithFederation`

Allows Lake Formation to assume a role to access tables in a federated database.

_Required_: No

_Type_: Boolean

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Fn::GetAtt

`Id`

An identifier for the catalog resource.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TableWithColumnsResource

AWS::LakeFormation::Tag
