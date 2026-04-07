This is the new _CloudFormation Template Reference Guide_.
Please update your bookmarks and links. For help getting started with CloudFormation, see the
[AWS CloudFormation User Guide](../userguide/welcome.md).

# AWS::Location::PlaceIndex

Specifies a place index resource in your AWS account. Use a place index
resource to geocode addresses and other text queries by using the
`SearchPlaceIndexForText` operation, and reverse geocode coordinates by
using the `SearchPlaceIndexForPosition` operation, and enable autosuggestions
by using the `SearchPlaceIndexForSuggestions` operation.

###### Note

If your application is tracking or routing assets you use in your business, such
as delivery vehicles or employees, you must not use Esri as your geolocation
provider. See section 82 of the [AWS service terms](https://aws.amazon.com/service-terms) for more details.

## Syntax

To declare this entity in your CloudFormation template, use the following syntax:

### JSON

```json

{
  "Type" : "AWS::Location::PlaceIndex",
  "Properties" : {
      "DataSource" : String,
      "DataSourceConfiguration" : DataSourceConfiguration,
      "Description" : String,
      "IndexName" : String,
      "PricingPlan" : String,
      "Tags" : [ Tag, ... ]
    }
}

```

### YAML

```yaml

Type: AWS::Location::PlaceIndex
Properties:
  DataSource: String
  DataSourceConfiguration:
    DataSourceConfiguration
  Description: String
  IndexName: String
  PricingPlan: String
  Tags:
    - Tag

```

## Properties

`DataSource`

Specifies the geospatial data provider for the new place index.

###### Note

This field is case-sensitive. Enter the valid values as shown. For example,
entering `HERE` returns an error.

Valid values include:

- `Esri` – For additional information about [Esri](https://docs.aws.amazon.com/location/previous/developerguide/esri.html)'s coverage in your region of interest, see [Esri details on geocoding coverage](https://developers.arcgis.com/rest/geocode/api-reference/geocode-coverage.htm).

- `Grab` – Grab provides place index functionality for Southeast
Asia. For additional information about [GrabMaps](https://docs.aws.amazon.com/location/previous/developerguide/grab.html)' coverage,
see [GrabMaps\
countries and areas covered](https://docs.aws.amazon.com/location/previous/developerguide/grab.html#grab-coverage-area).

- `Here` – For additional information about [HERE\
Technologies](https://docs.aws.amazon.com/location/previous/developerguide/HERE.html)' coverage in your region of interest, see [HERE details on goecoding coverage](https://developer.here.com/documentation/geocoder/dev_guide/topics/coverage-geocoder.html).

###### Important

If you specify HERE Technologies ( `Here`) as the data provider,
you may not [store results](https://docs.aws.amazon.com/location-places/latest/APIReference/API_DataSourceConfiguration.html) for locations in Japan. For more information, see
the [AWS service\
terms](https://aws.amazon.com/service-terms) for Amazon Location Service.

For additional information , see [Data\
providers](https://docs.aws.amazon.com/location/previous/developerguide/what-is-data-provider.html) on the _Amazon Location Service developer guide_.

_Required_: Yes

_Type_: String

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

`DataSourceConfiguration`

Specifies the data storage option requesting Places.

_Required_: No

_Type_: [DataSourceConfiguration](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-location-placeindex-datasourceconfiguration.html)

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`Description`

The optional description for the place index resource.

_Required_: No

_Type_: String

_Minimum_: `0`

_Maximum_: `1000`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

`IndexName`

The name of the place index resource.

Requirements:

- Contain only alphanumeric characters (A–Z, a–z, 0–9), hyphens (-), periods
(.), and underscores (\_).

- Must be a unique place index resource name.

- No spaces allowed. For example, `ExamplePlaceIndex`.

_Required_: Yes

_Type_: String

_Pattern_: `^[-._\w]+$`

_Minimum_: `1`

_Maximum_: `100`

_Update requires_: [Replacement](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-replacement)

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

_Type_: Array of [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/TemplateReference/aws-properties-location-placeindex-tag.html)

_Minimum_: `0`

_Maximum_: `200`

_Update requires_: [No interruption](../userguide/using-cfn-updating-stacks-update-behaviors.md#update-no-interrupt)

## Return values

### Ref

When you pass the logical ID of this resource to the intrinsic `Ref` function, `Ref` returns the `PlaceIndex` name.

For more information about using the `Ref` function, see [`Ref`](intrinsic-function-reference-ref.md).

### Fn::GetAtt

The `Fn::GetAtt` intrinsic function returns a value for a specified attribute of this type. The following are the available attributes and sample return values.

For more information about using the `Fn::GetAtt` intrinsic function, see [`Fn::GetAtt`](intrinsic-function-reference-getatt.md).

`Arn`

The Amazon Resource Name (ARN) for the place index resource. Used to specify a
resource across AWS.

- Format example:
`arn:aws:geo:region:account-id:place-index/ExamplePlaceIndex`

`CreateTime`

The timestamp for when the place index resource was created in [ISO 8601](https://www.iso.org/iso-8601-date-and-time-format.html)
format: `YYYY-MM-DDThh:mm:ss.sssZ`.

`IndexArn`

Synonym for `Arn`. The Amazon Resource Name (ARN) for the place index
resource. Used to specify a resource across AWS.

- Format example:
`arn:aws:geo:region:account-id:place-index/ExamplePlaceIndex`

`UpdateTime`

The timestamp for when the place index resource was last updated in [ISO 8601](https://www.iso.org/iso-8601-date-and-time-format.html)
format: `YYYY-MM-DDThh:mm:ss.sssZ`.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Tag

DataSourceConfiguration
