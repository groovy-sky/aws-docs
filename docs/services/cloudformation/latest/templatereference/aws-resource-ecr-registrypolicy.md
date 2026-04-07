This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECR::RegistryPolicy

The `AWS::ECR::RegistryPolicy` resource creates or updates the permissions
policy for a private registry.

A private registry policy is used to specify permissions for another AWS account and is used when configuring cross-account replication. For
more information, see [Registry\
permissions](https://docs.aws.amazon.com/AmazonECR/latest/userguide/registry-permissions.html) in the _Amazon Elastic Container Registry User_
_Guide_.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ECR::RegistryPolicy",
  "Properties" : {
      "PolicyText" : Json
    }
}

```

### YAML

```yaml

Type: AWS::ECR::RegistryPolicy
Properties:
  PolicyText: Json

```

## Properties

`PolicyText`

The JSON policy text for your registry.

_Required_: Yes

_Type_: Json

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`RegistryId`

The account ID of the private registry the policy is associated with.

## Examples

### Specify a registry policy for a private registry

The following example specifies a private registry policy in
`us-west-2` that grants permission for account
`210987654321` to create repositories and replicate their
contents to your private registry.

#### JSON

```json

"TestRegistryPolicy": {
   "Type": "AWS::ECR::RegistryPolicy",
   "Properties": {
      "PolicyText": {
         "Version":"2012-10-17",
         "Statement":[
            {
               "Sid":"ReplicationAccessCrossAccount",
               "Effect":"Allow",
               "Principal":{
                  "AWS":"arn:aws:iam::210987654321:root"
               },
               "Action":[
                  "ecr:CreateRepository",
                  "ecr:ReplicateImage"
               ],
               "Resource": "arn:aws:ecr:us-west-2:123456789012:repository/*"
             }
          ]
       }
    }
}
```

#### YAML

```yaml

Resources:
  TestRegistryPolicy:
    Type: 'AWS::ECR::RegistryPolicy'
    Properties:
      PolicyText:
        Version: 2012-10-17
        Statement:
          - Sid: UpdatedRegistryPolicy
            Effect: Allow
            Principal:
              AWS: 'arn:aws:iam::210987654321:root'
            Action:
              - 'ecr:CreateRepository'
              - 'ecr:ReplicateImage'
            Resource: 'arn:aws:ecr:us-west-2:123456789012:repository/*'
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::ECR::PullTimeUpdateExclusion

AWS::ECR::RegistryScanningConfiguration
