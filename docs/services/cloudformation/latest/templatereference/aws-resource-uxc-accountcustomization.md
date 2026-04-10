This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::UXC::AccountCustomization

The `AWS::UXC::AccountCustomization` resource specifies account-level UX customization settings for the console, including account color, visible services, and visible regions.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::UXC::AccountCustomization",
  "Properties" : {
      "AccountColor" : String,
      "VisibleRegions" : [ String, ... ],
      "VisibleServices" : [ String, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::UXC::AccountCustomization
Properties:
  AccountColor: String
  VisibleRegions:
    - String
  VisibleServices:
    - String

```

## Properties

`AccountColor`

The account color preference to set. Set to `none` to reset to the default (no color).

_Required_: No

_Type_: String

_Allowed values_: `none | pink | purple | darkBlue | lightBlue | teal | green | yellow | orange | red`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VisibleRegions`

The list of AWS Region codes to make visible in the AWS Management Console. Set to `null` to reset to the default, which makes all Regions visible. For a list of valid Region codes, see [AWS Regions](../../../global-infrastructure/latest/regions/aws-regions.md).

_Required_: No

_Type_: Array of String

_Maximum_: `100`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`VisibleServices`

The list of AWS service identifiers to make visible in the AWS Management Console. Set to `null` to reset to the default, which makes all services visible. For valid service identifiers, call [ListServices](../../../../reference/awsconsolehelpdocs/latest/apireference/api-listservices.md).

_Required_: No

_Type_: Array of String

_Minimum_: `1`

_Maximum_: `64 | 500`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the account ID. For example: `123456789012`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`AccountId`

The 12-digit account ID that this customization belongs to. For example: `123456789012`.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AWS User Experience Customization

Next

All content copied from https://docs.aws.amazon.com/.
