This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Location::RouteCalculator

Specifies a route calculator resource in your AWS account.

You can send requests to a route calculator resource to estimate travel time,
distance, and get directions. A route calculator sources traffic and road network data
from your chosen data provider.

###### Note

If your application is tracking or routing assets you use in your business, such
as delivery vehicles or employees, you must not use Esri as your geolocation
provider. See section 82 of the [AWS\
service terms](https://aws.amazon.com/service-terms) for more details.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Location::RouteCalculator",
  "Properties" : {
      "CalculatorName" : String,
      "DataSource" : String,
      "Description" : String,
      "PricingPlan" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Location::RouteCalculator
Properties:
  CalculatorName: String
  DataSource: String
  Description: String
  PricingPlan: String
  Tags:
    - Tag

```

## Properties

`CalculatorName`

The name of the route calculator resource.

Requirements:

- Can use alphanumeric characters (A–Z, a–z, 0–9) , hyphens (-), periods (.),
and underscores (\_).

- Must be a unique Route calculator resource name.

- No spaces allowed. For example, `ExampleRouteCalculator`.

_Required_: Yes

_Type_: String

_Pattern_: `^[-._\w]+$`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DataSource`

Specifies the data provider of traffic and road network data.

###### Note

This field is case-sensitive. Enter the valid values as shown. For example,
entering `HERE` returns an error.

Valid values include:

- `Esri` – For additional information about [Esri](https://docs.aws.amazon.com/location/previous/developerguide/esri.html)'s coverage in your region of interest, see [Esri details on street networks and traffic coverage](https://doc.arcgis.com/en/arcgis-online/reference/network-coverage.htm).

Route calculators that use Esri as a data source only calculate routes that
are shorter than 400 km.

- `Grab` – Grab provides routing functionality for Southeast Asia.
For additional information about [GrabMaps](https://docs.aws.amazon.com/location/previous/developerguide/grab.html)' coverage,
see [GrabMaps\
countries and areas covered](https://docs.aws.amazon.com/location/previous/developerguide/grab.html#grab-coverage-area).

- `Here` – For additional information about [HERE\
Technologies](https://docs.aws.amazon.com/location/previous/developerguide/HERE.html)' coverage in your region of interest, see [HERE car routing coverage](https://developer.here.com/documentation/routing-api/dev_guide/topics/coverage/car-routing.html) and [HERE truck routing coverage](https://developer.here.com/documentation/routing-api/dev_guide/topics/coverage/truck-routing.html).

For additional information , see [Data\
providers](https://docs.aws.amazon.com/location/previous/developerguide/what-is-data-provider.html) on the _Amazon Location Service Developer Guide_.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`Description`

The optional description for the route calculator resource.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`PricingPlan`

No longer used. If included, the only allowed value is
`RequestBasedUsage`.

_Allowed Values_: `RequestBasedUsage`

_Required_: No

_Type_: String

_Allowed values_: `RequestBasedUsage`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Tags`

Property description not available.

_Required_: No

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-location-routecalculator-tag.html)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the `RouteCalculator` name.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) for the route calculator resource. Use the ARN when you
specify a resource across all AWS.

- Format example:
`arn:aws:geo:region:account-id:route-calculator/ExampleCalculator`

`CalculatorArn`

Synonym for `Arn`. The Amazon Resource Name (ARN) for the route calculator resource. Use the ARN when you
specify a resource across all AWS.

- Format example:
`arn:aws:geo:region:account-id:route-calculator/ExampleCalculator`

`CreateTime`

The timestamp for when the route calculator resource was created in [ISO 8601](https://www.iso.org/iso-8601-date-and-time-format.html) format:
`YYYY-MM-DDThh:mm:ss.sssZ`.

`UpdateTime`

The timestamp for when the route calculator resource was last updated in [ISO 8601](https://www.iso.org/iso-8601-date-and-time-format.html) format:
`YYYY-MM-DDThh:mm:ss.sssZ`.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

Tag
