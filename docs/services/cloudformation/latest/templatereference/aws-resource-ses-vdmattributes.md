This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::SES::VdmAttributes

The Virtual Deliverability Manager (VDM) attributes that apply to your Amazon SES account.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::SES::VdmAttributes",
  "Properties" : {
      "DashboardAttributes" : DashboardAttributes,
      "GuardianAttributes" : GuardianAttributes
    }
}

```

### YAML

```yaml

Type: AWS::SES::VdmAttributes
Properties:
  DashboardAttributes:
    DashboardAttributes
  GuardianAttributes:
    GuardianAttributes

```

## Properties

`DashboardAttributes`

Specifies additional settings for your VDM configuration as applicable to the
Dashboard.

_Required_: No

_Type_: [DashboardAttributes](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ses-vdmattributes-dashboardattributes.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`GuardianAttributes`

Specifies additional settings for your VDM configuration as applicable to the
Guardian.

_Required_: No

_Type_: [GuardianAttributes](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-ses-vdmattributes-guardianattributes.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the resource name.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

`VdmAttributesResourceId`

Property description not available.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

DashboardAttributes
