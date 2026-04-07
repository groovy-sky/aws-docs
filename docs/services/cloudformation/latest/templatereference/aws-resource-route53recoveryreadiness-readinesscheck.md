This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Route53RecoveryReadiness::ReadinessCheck

Creates a readiness check in Amazon Route 53 Application Recovery Controller. A readiness check continually monitors a resource set in your application,
such as a set of Amazon Aurora instances, that Route 53 ARC is auditing recovery readiness for. The audits run once every minute on every resource that's associated with
a readiness check.

Every resource type has a set of rules associated with it that Route 53 ARC uses to audit resources for readiness. For more information, see
[Readiness rules descriptions](https://docs.aws.amazon.com/r53recovery/latest/dg/recovery-readiness.rules-resources.html)
in the Amazon Route 53 Application Recovery Controller Developer Guide.

Route 53 ARC Readiness supports us-east-1 and us-west-2 AWS Regions only.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Route53RecoveryReadiness::ReadinessCheck",
  "Properties" : {
      "ReadinessCheckName" : String,
      "ResourceSetName" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Route53RecoveryReadiness::ReadinessCheck
Properties:
  ReadinessCheckName: String
  ResourceSetName: String
  Tags:
    - Tag

```

## Properties

`ReadinessCheckName`

The name of the readiness check to create.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9_]+`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`ResourceSetName`

The name of the resource set to check.

_Required_: No

_Type_: String

_Pattern_: `[a-zA-Z0-9_]+`

_Minimum_: `1`

_Maximum_: `64`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

A collection of tags associated with a resource.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-route53recoveryreadiness-readinesscheck-tag.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the `ReadinessCheckName`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`ReadinessCheckArn`

The Amazon Resource Name (ARN) of the readiness check.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Tag
