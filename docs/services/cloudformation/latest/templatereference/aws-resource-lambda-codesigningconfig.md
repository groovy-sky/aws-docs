This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Lambda::CodeSigningConfig

Details about a [Code signing configuration](https://docs.aws.amazon.com/lambda/latest/dg/configuration-codesigning.html).

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Lambda::CodeSigningConfig",
  "Properties" : {
      "AllowedPublishers" : AllowedPublishers,
      "CodeSigningPolicies" : CodeSigningPolicies,
      "Description" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Lambda::CodeSigningConfig
Properties:
  AllowedPublishers:
    AllowedPublishers
  CodeSigningPolicies:
    CodeSigningPolicies
  Description: String
  Tags:
    - Tag

```

## Properties

`AllowedPublishers`

List of allowed publishers.

_Required_: Yes

_Type_: [AllowedPublishers](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-lambda-codesigningconfig-allowedpublishers.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`CodeSigningPolicies`

The code signing policy controls the validation failure action for signature mismatch or expiry.

_Required_: No

_Type_: [CodeSigningPolicies](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-lambda-codesigningconfig-codesigningpolicies.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

Code signing configuration description.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `256`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A list of tags to add to the code signing configuration.

###### Note

You must have the `lambda:TagResource`, `lambda:UntagResource`,
and `lambda:ListTags` permissions for your [IAM principal](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_terms-and-concepts.html) to manage the CloudFormation stack. If you
don't have these permissions, there might be unexpected behavior with stack-level tags
propagating to the resource during resource creation and update.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-lambda-codesigningconfig-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource name.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`CodeSigningConfigArn`

The Amazon Resource Name (ARN) of the code signing configuration.

`CodeSigningConfigId`

The code signing configuration ID.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

TargetTrackingScalingPolicy

AllowedPublishers
