This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::FinSpace::Environment

The `AWS::FinSpace::Environment` resource represents an Amazon FinSpace
environment.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::FinSpace::Environment",
  "Properties" : {
      "Description" : String,
      "FederationMode" : String,
      "FederationParameters" : FederationParameters,
      "KmsKeyId" : String,
      "Name" : String,
      "SuperuserParameters" : SuperuserParameters,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::FinSpace::Environment
Properties:
  Description: String
  FederationMode: String
  FederationParameters:
    FederationParameters
  KmsKeyId: String
  Name: String
  SuperuserParameters:
    SuperuserParameters
  Tags:
    - Tag

```

## Properties

`Description`

The description of the FinSpace environment.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z0-9. ]{1,1000}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FederationMode`

The authentication mode for the environment.

_Required_: No

_Type_: String

_Allowed values_: `LOCAL | FEDERATED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`FederationParameters`

Configuration information when authentication mode is FEDERATED.

_Required_: No

_Type_: [FederationParameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-finspace-environment-federationparameters.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`KmsKeyId`

The KMS key id used to encrypt in the FinSpace environment.

_Required_: No

_Type_: String

_Pattern_: `^[a-zA-Z-0-9-:\/]*{1,1000}$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Name`

The name of the FinSpace environment.

_Required_: Yes

_Type_: String

_Pattern_: `^[a-zA-Z0-9]+[a-zA-Z0-9-]*[a-zA-Z0-9]{1,255}$`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`SuperuserParameters`

Configuration information for the superuser.

_Required_: No

_Type_: [SuperuserParameters](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-finspace-environment-superuserparameters.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Tags`

Property description not available.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-finspace-environment-tag.html)

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource name. For example:

`{ "Ref": "myEnvironment" }`

For the Amazon FinSpace environment group `myEnvironment`, `Ref` returns the name of the environment.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AwsAccountId`

The ID of the AWS account in which the FinSpace environment is created.

`DedicatedServiceAccountId`

The AWS account ID of the dedicated service account associated with your FinSpace
environment.

`EnvironmentArn`

The Amazon Resource Name (ARN) of your FinSpace environment.

`EnvironmentId`

The identifier of the FinSpace environment.

`EnvironmentUrl`

The sign-in url for the web application of your FinSpace environment.

`SageMakerStudioDomainUrl`

The url of the integrated FinSpace notebook environment in your web application.

`Status`

The current status of creation of the FinSpace environment.

## Examples

### Creating environments

The following examples create new FinSpace environments.

#### YAML

```yaml

Resources:
  FinSpaceEnvironment:
    Type: 'AWS::FinSpace::Environment'
    Properties:
        Name: MyEnvironment
        KmsKeyId: arn:aws:kms:us-east-1:123456789012:key/44efed01-30d0-4b39-80e7-165d5ed34524
        FederationMode: LOCAL

```

#### JSON

```json

{
   "Resources": {
      "FinSpaceEnvironment": {
         "Type": "AWS::FinSpace::Environment",
         "Properties": {
            "Name": "MyEnvironment",
            "KmsKeyId": "arn:aws:kms:us-east-1:123456789012:key/44efed01-30d0-4b39-80e7-165d5ed34524",
            "FederationMode": "LOCAL"
         }
      }
   }
}

```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Amazon FinSpace schemas

AttributeMapItems
