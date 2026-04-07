This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ECR::PullTimeUpdateExclusion

The `AWS::ECR::PullTimeUpdateExclusion` resource Property description not available. for ECR.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ECR::PullTimeUpdateExclusion",
  "Properties" : {
      "PrincipalArn" : String
    }
}

```

### YAML

```yaml

Type: AWS::ECR::PullTimeUpdateExclusion
Properties:
  PrincipalArn: String

```

## Properties

`PrincipalArn`

The ARN of the IAM principal to remove from the pull time update exclusion list.

_Required_: Yes

_Type_: String

_Pattern_: `^arn:aws(-[a-z]+)*:iam::[0-9]{12}:(role|user)/[\w+=,.@-]+(/[\w+=,.@-]+)*$`

_Maximum_: `200`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS::ECR::PullThroughCacheRule

AWS::ECR::RegistryPolicy
