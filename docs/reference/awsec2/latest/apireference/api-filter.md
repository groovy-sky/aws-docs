# Filter

A filter name and value pair that is used to return a more specific list of results from a describe operation.
Filters can be used to match a set of resources by specific criteria, such as tags, attributes, or IDs.

If you specify multiple filters, the filters are joined with an `AND`, and the request returns only
results that match all of the specified filters.

The filters supported by a describe operation are documented with the describe operation. For example:

- [DescribeAvailabilityZones](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeAvailabilityZones.html)

- [DescribeImages](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeImages.html)

- [DescribeInstances](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeInstances.html)

- [DescribeKeyPairs](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeKeyPairs.html)

- [DescribeSecurityGroups](api-describesecuritygroups.md)

- [DescribeSnapshots](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeSnapshots.html)

- [DescribeSubnets](api-describesubnets.md)

- [DescribeTags](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeTags.html)

- [DescribeVolumes](https://docs.aws.amazon.com/AWSEC2/latest/APIReference/API_DescribeVolumes.html)

- [DescribeVpcs](api-describevpcs.md)

For more information, see [List and filter using the CLI and API](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/Using_Filtering.html#Filtering_Resources_CLI) in the _Amazon EC2 User Guide_.

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

- [AWS SDK for C++](https://docs.aws.amazon.com/goto/SdkForCpp/ec2-2016-11-15/Filter)

- [AWS SDK for Java V2](https://docs.aws.amazon.com/goto/SdkForJavaV2/ec2-2016-11-15/Filter)

- [AWS SDK for Ruby V3](https://docs.aws.amazon.com/goto/SdkForRubyV3/ec2-2016-11-15/Filter)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

FederatedAuthenticationRequest

FilterPortRange
