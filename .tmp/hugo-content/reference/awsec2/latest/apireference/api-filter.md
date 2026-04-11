# Filter

A filter name and value pair that is used to return a more specific list of results from a describe operation.
Filters can be used to match a set of resources by specific criteria, such as tags, attributes, or IDs.

If you specify multiple filters, the filters are joined with an `AND`, and the request returns only
results that match all of the specified filters.

The filters supported by a describe operation are documented with the describe operation. For example:

- [DescribeAvailabilityZones](api-describeavailabilityzones.md)

- [DescribeImages](api-describeimages.md)

- [DescribeInstances](api-describeinstances.md)

- [DescribeKeyPairs](api-describekeypairs.md)

- [DescribeSecurityGroups](api-describesecuritygroups.md)

- [DescribeSnapshots](api-describesnapshots.md)

- [DescribeSubnets](api-describesubnets.md)

- [DescribeTags](api-describetags.md)

- [DescribeVolumes](api-describevolumes.md)

- [DescribeVpcs](api-describevpcs.md)

For more information, see [List and filter using the CLI and API](../../../../services/ec2/latest/userguide/using-filtering.md#Filtering_Resources_CLI) in the _Amazon EC2 User Guide_.

## Contents

**Name**

The name of the filter. Filter names are case-sensitive.

Type: String

Required: No

**Value.N**

The filter values. Filter values are case-sensitive. If you specify multiple values for a
filter, the values are joined with an `OR`, and the request returns all results
that match any of the specified values.

Type: Array of strings

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/filter.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/filter.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/filter.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

FederatedAuthenticationRequest

FilterPortRange
