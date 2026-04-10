This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::CloudWatch::Dashboard

The `AWS::CloudWatch::Dashboard` resource specifies an Amazon CloudWatch dashboard. A dashboard is a
customizable home page in the CloudWatch console that you can use to monitor your AWS resources in a single view.

All dashboards in your account are global, not region-specific.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::CloudWatch::Dashboard",
  "Properties" : {
      "DashboardBody" : String,
      "DashboardName" : String
    }
}

```

### YAML

```yaml

Type: AWS::CloudWatch::Dashboard
Properties:
  DashboardBody: String
  DashboardName: String

```

## Properties

`DashboardBody`

The detailed information about the dashboard in JSON format, including the widgets to include and their location
on the dashboard. This parameter is required.

For more information about the syntax,
see [Dashboard Body Structure and Syntax](../../../../reference/amazoncloudwatch/latest/apireference/cloudwatch-dashboard-body-structure.md).

_Required_: Yes

_Type_: String

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`DashboardName`

The name of the dashboard. The name must be between 1 and 255 characters. If you do not specify a name, one will be generated automatically.

_Required_: No

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the value of `DashboardName`.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

## See also

- [Amazon CloudWatch Template Snippets](../userguide/quickref-cloudwatch.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Tag

AWS::CloudWatch::InsightRule

All content copied from https://docs.aws.amazon.com/.
