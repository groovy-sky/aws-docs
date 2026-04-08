# DescribeAvailabilityZones

Describes the Availability Zones, Local Zones, and Wavelength Zones that are available to
you.

For more information about Availability Zones, Local Zones, and Wavelength Zones, see
[Regions and zones](../../../../services/ec2/latest/userguide/using-regions-availability-zones.md)
in the _Amazon EC2 User Guide_.

###### Note

The order of the elements in the response, including those within nested
structures, might vary. Applications should not assume the elements appear in a
particular order.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**AllAvailabilityZones**

Include all Availability Zones, Local Zones, and Wavelength Zones regardless of your
opt-in status.

If you do not use this parameter, the results include only the zones for the Regions where you have chosen the option to opt in.

Type: Boolean

Required: No

**DryRun**

Checks whether you have the required permissions for the action, without actually making the request,
and provides an error response. If you have the required permissions, the error response is `DryRunOperation`.
Otherwise, it is `UnauthorizedOperation`.

Type: Boolean

Required: No

**Filter.N**

The filters.

- `group-long-name` \- The long name of the zone group for the Availability Zone (for
example, `US West (Oregon) 1`), the Local Zone (for example, for Zone group `us-west-2-lax-1`, it is `US West (Los Angeles)`,
or the Wavelength Zone (for example, for Zone group `us-east-1-wl1`, it is `US East (Verizon)`.

- `group-name` \- The name of the zone group for the Availability Zone (for
example, `us-east-1-zg-1`), the Local Zone (for example, `us-west-2-lax-1`),
or the Wavelength Zone (for example, `us-east-1-wl1`).

- `message` \- The Zone message.

- `opt-in-status` \- The opt-in status ( `opted-in` \|
`not-opted-in` \| `opt-in-not-required`).

- `parent-zone-id` \- The ID of the zone that handles some of the Local Zone
and Wavelength Zone control plane operations, such as API calls.

- `parent-zone-name` \- The ID of the zone that handles some of the Local Zone
and Wavelength Zone control plane operations, such as API calls.

- `region-name` \- The name of the Region for the Zone (for example,
`us-east-1`).

- `state` \- The state of the Availability Zone, the Local Zone, or the
Wavelength Zone ( `available` \| `unavailable` \|
`constrained`).

- `zone-id` \- The ID of the Availability Zone (for example,
`use1-az1`), the Local Zone (for example, `usw2-lax1-az1`), or the
Wavelength Zone (for example, `us-east-1-wl1-bos-wlz-1`).

- `zone-name` \- The name of the Availability Zone (for example,
`us-east-1a`), the Local Zone (for example, `us-west-2-lax-1a`), or
the Wavelength Zone (for example, `us-east-1-wl1-bos-wlz-1`).

- `zone-type` \- The type of zone ( `availability-zone` \|
`local-zone` \| `wavelength-zone`).

Type: Array of [Filter](api-filter.md) objects

Required: No

**ZoneId.N**

The IDs of the Availability Zones, Local Zones, and Wavelength Zones.

Type: Array of strings

Required: No

**ZoneName.N**

The names of the Availability Zones, Local Zones, and Wavelength Zones.

Type: Array of strings

Required: No

## Response Elements

The following elements are returned by the service.

**availabilityZoneInfo**

Information about the Availability Zones, Local Zones, and Wavelength Zones.

Type: Array of [AvailabilityZone](api-availabilityzone.md) objects

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example request describes the zones in the current Region that are enabled for
your account.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DescribeAvailabilityZones
&AUTHPARAMS
```

#### Sample Response

```

<DescribeAvailabilityZonesResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>e23c5a54-a29c-43ee-8b55-0c13c26e9e01</requestId>
    <availabilityZoneInfo>
        <item>
            <groupName>us-west-2-zg-1</groupName>
            <optInStatus>opt-in-not-required</optInStatus>
            <zoneName>us-west-2a</zoneName>
            <zoneId>usw2-az1</zoneId>
            <zoneState>available</zoneState>
            <zoneType>availability-zone</zoneType>
            <regionName>us-west-2</regionName>
            <messageSet/>
            <networkBorderGroup>us-west-2</networkBorderGroup>
            <groupLongName>US West (Oregon) 1</groupLongName>
            <geographySet>
                <item>
                    <name>United States of America</name>
                </item>
            </geographySet>
            <subGeographySet>
                <item>
                    <name>Oregon</name>
                </item>
            </subGeographySet>
        </item>
        <item>
            <groupName>us-west-2-zg-1</groupName>
            <optInStatus>opt-in-not-required</optInStatus>
            <zoneName>us-west-2b</zoneName>
            <zoneId>usw2-az2</zoneId>
            <zoneState>available</zoneState>
            <zoneType>availability-zone</zoneType>
            <regionName>us-west-2</regionName>
            <messageSet/>
            <networkBorderGroup>us-west-2</networkBorderGroup>
            <groupLongName>US West (Oregon) 1</groupLongName>
            <geographySet>
                <item>
                    <name>United States of America</name>
                </item>
            </geographySet>
            <subGeographySet>
                <item>
                    <name>Oregon</name>
                </item>
            </subGeographySet>
        </item>
        <item>
            <groupName>us-west-2-zg-1</groupName>
            <optInStatus>opt-in-not-required</optInStatus>
            <zoneName>us-west-2c</zoneName>
            <zoneId>usw2-az3</zoneId>
            <zoneState>available</zoneState>
            <zoneType>availability-zone</zoneType>
            <regionName>us-west-2</regionName>
            <messageSet/>
            <networkBorderGroup>us-west-2</networkBorderGroup>
            <groupLongName>US West (Oregon) 1</groupLongName>
            <geographySet>
                <item>
                    <name>United States of America</name>
                </item>
            </geographySet>
            <subGeographySet>
                <item>
                    <name>Oregon</name>
                </item>
            </subGeographySet>
        </item>
        <item>
            <groupName>us-west-2-zg-1</groupName>
            <optInStatus>opt-in-not-required</optInStatus>
            <zoneName>us-west-2d</zoneName>
            <zoneId>usw2-az4</zoneId>
            <zoneState>available</zoneState>
            <zoneType>availability-zone</zoneType>
            <regionName>us-west-2</regionName>
            <messageSet/>
            <networkBorderGroup>us-west-2</networkBorderGroup>
            <groupLongName>US West (Oregon) 1</groupLongName>
            <geographySet>
                <item>
                    <name>United States of America</name>
                </item>
            </geographySet>
            <subGeographySet>
                <item>
                    <name>Oregon</name>
                </item>
            </subGeographySet>
        </item>
    </availabilityZoneInfo>
</DescribeAvailabilityZonesResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/describeavailabilityzones.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/describeavailabilityzones.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/describeavailabilityzones.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/describeavailabilityzones.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/describeavailabilityzones.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/describeavailabilityzones.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/describeavailabilityzones.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/describeavailabilityzones.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/describeavailabilityzones.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/describeavailabilityzones.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribeAggregateIdFormat

DescribeAwsNetworkPerformanceMetricSubscriptions
