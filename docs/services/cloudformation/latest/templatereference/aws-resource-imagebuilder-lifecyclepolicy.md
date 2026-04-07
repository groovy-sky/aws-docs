This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::ImageBuilder::LifecyclePolicy

Create a lifecycle policy resource.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::ImageBuilder::LifecyclePolicy",
  "Properties" : {
      "Description" : String,
      "ExecutionRole" : String,
      "Name" : String,
      "PolicyDetails" : [ PolicyDetail, ... ],
      "ResourceSelection" : ResourceSelection,
      "ResourceType" : String,
      "Status" : String,
      "Tags" : {Key: Value, ...}
    }
}

```

### YAML

```yaml

Type: AWS::ImageBuilder::LifecyclePolicy
Properties:
  Description: String
  ExecutionRole: String
  Name: String
  PolicyDetails:
    - PolicyDetail
  ResourceSelection:
    ResourceSelection
  ResourceType: String
  Status: String
  Tags:
    Key: Value

```

## Properties

`Description`

Optional description for the lifecycle policy.

_Required_: No

_Type_: String

_Minimum_: `1`

_Maximum_: `1024`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ExecutionRole`

The name or Amazon Resource Name (ARN) for the IAM role you create that grants
Image Builder access to run lifecycle actions.

_Required_: Yes

_Type_: String

_Pattern_: `^(?:arn:aws(?:-[a-z]+)*:iam::[0-9]{12}:role/)?[a-zA-Z_0-9+=,.@\-_/]+$`

_Minimum_: `1`

_Maximum_: `2048`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Name`

The name of the lifecycle policy to create.

_Required_: Yes

_Type_: String

_Pattern_: `^[-_A-Za-z-0-9][-_A-Za-z0-9 ]{1,126}[-_A-Za-z-0-9]$`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`PolicyDetails`

Configuration details for the lifecycle policy rules.

_Required_: Yes

_Type_: Array of [PolicyDetail](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-imagebuilder-lifecyclepolicy-policydetail.html)

_Minimum_: `1`

_Maximum_: `3`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceSelection`

Selection criteria for the resources that the lifecycle policy applies to.

_Required_: Yes

_Type_: [ResourceSelection](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-imagebuilder-lifecyclepolicy-resourceselection.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`ResourceType`

The type of Image Builder resource that the lifecycle policy applies to.

_Required_: Yes

_Type_: String

_Allowed values_: `AMI_IMAGE | CONTAINER_IMAGE`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Status`

Indicates whether the lifecycle policy resource is enabled.

_Required_: No

_Type_: String

_Allowed values_: `DISABLED | ENABLED`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Tags to apply to the lifecycle policy resource.

_Required_: No

_Type_: Object of String

_Pattern_: `.{1,}`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

### Fn::GetAtt

`Arn`

The Amazon Resource Name (ARN) of the lifecycle policy resource.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

S3Logs

Action
