# DescribeLaunchTemplateVersions

Describes one or more versions of a specified launch template. You can describe all
versions, individual versions, or a range of versions. You can also describe all the
latest versions or all the default versions of all the launch templates in your
account.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually
making the request, and provides an error response. If you have the required
permissions, the error response is `DryRunOperation`. Otherwise, it is
`UnauthorizedOperation`.

Type: Boolean

Required: No

**Filter.N**

One or more filters.

- `create-time` \- The time the launch template version was
created.

- `ebs-optimized` \- A boolean that indicates whether the instance is
optimized for Amazon EBS I/O.

- `http-endpoint` \- Indicates whether the HTTP metadata endpoint on
your instances is enabled ( `enabled` \| `disabled`).

- `http-protocol-ipv4` \- Indicates whether the IPv4 endpoint for the
instance metadata service is enabled ( `enabled` \|
`disabled`).

- `host-resource-group-arn` \- The ARN of the host resource group in
which to launch the instances.

- `http-tokens` \- The state of token usage for your instance metadata
requests ( `optional` \| `required`).

- `iam-instance-profile` \- The ARN of the IAM instance
profile.

- `image-id` \- The ID of the AMI.

- `instance-type` \- The instance type.

- `is-default-version` \- A boolean that indicates whether the launch
template version is the default version.

- `kernel-id` \- The kernel ID.

- `license-configuration-arn` \- The ARN of the license
configuration.

- `network-card-index` \- The index of the network card.

- `ram-disk-id` \- The RAM disk ID.

Type: Array of [Filter](api-filter.md) objects

Required: No

**LaunchTemplateId**

The ID of the launch template.

To describe one or more versions of a specified launch template, you must specify
either the launch template ID or the launch template name, but not both.

To describe all the latest or default launch template versions in your account, you
must omit this parameter.

Type: String

Required: No

**LaunchTemplateName**

The name of the launch template.

To describe one or more versions of a specified launch template, you must specify
either the launch template name or the launch template ID, but not both.

To describe all the latest or default launch template versions in your account, you
must omit this parameter.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 128.

Pattern: `[a-zA-Z0-9\(\)\.\-/_]+`

Required: No

**LaunchTemplateVersion.N**

One or more versions of the launch template. Valid values depend on whether you are
describing a specified launch template (by ID or name) or all launch templates in your
account.

To describe one or more versions of a specified launch template, valid values are
`$Latest`, `$Default`, and numbers.

To describe all launch templates in your account that are defined as the latest
version, the valid value is `$Latest`. To describe all launch templates in
your account that are defined as the default version, the valid value is
`$Default`. You can specify `$Latest` and
`$Default` in the same request. You cannot specify numbers.

Type: Array of strings

Required: No

**MaxResults**

The maximum number of results to return in a single call. To retrieve the remaining
results, make another call with the returned `NextToken` value. This value
can be between 1 and 200.

Type: Integer

Required: No

**MaxVersion**

The version number up to which to describe launch template versions.

Type: String

Required: No

**MinVersion**

The version number after which to describe launch template versions.

Type: String

Required: No

**NextToken**

The token to request the next page of results.

Type: String

Required: No

**ResolveAlias**

If `true`, and if a Systems Manager parameter is specified for
`ImageId`, the AMI ID is displayed in the response for
`imageId`.

If `false`, and if a Systems Manager parameter is specified for
`ImageId`, the parameter is displayed in the response for
`imageId`.

For more information, see [Use a Systems Manager parameter instead of an AMI ID](../../../../services/ec2/latest/userguide/create-launch-template.md#use-an-ssm-parameter-instead-of-an-ami-id) in the
_Amazon EC2 User Guide_.

Default: `false`

Type: Boolean

Required: No

## Response Elements

The following elements are returned by the service.

**launchTemplateVersionSet**

Information about the launch template versions.

Type: Array of [LaunchTemplateVersion](api-launchtemplateversion.md) objects

**nextToken**

The token to use to retrieve the next page of results. This value is `null`
when there are no more results to return.

Type: String

**requestId**

The ID of the request.

Type: String

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example 1

This example describes all versions of launch template
`lt-0a20c965061f64abc` up to version 3.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DescribeLaunchTemplateVersions
&LaunchTemplateId=lt-0a20c965061f64abc
&MaxVersion=3
&AUTHPARAMS
```

#### Sample Response

```

<DescribeLaunchTemplateVersionsResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>65cadec1-b364-4354-8ca8-4176dexample</requestId>
    <launchTemplateVersionSet>
        <item>
            <createTime>2017-10-31T11:38:52.000Z</createTime>
            <createdBy>arn:aws:iam::123456789012:root</createdBy>
            <defaultVersion>true</defaultVersion>
            <launchTemplateData>
                <imageId>ami-8c1be5f6</imageId>
                <instanceType>t2.micro</instanceType>
            </launchTemplateData>
            <launchTemplateId>lt-0a20c965061f64abc</launchTemplateId>
            <launchTemplateName>MyLaunchTemplate</launchTemplateName>
            <versionDescription>FirstVersion</versionDescription>
            <versionNumber>1</versionNumber>
        </item>
        <item>
            <createTime>2017-10-31T11:52:03.000Z</createTime>
            <createdBy>arn:aws:iam::123456789012:root</createdBy>
            <defaultVersion>false</defaultVersion>
            <launchTemplateData>
                <imageId>ami-12345678</imageId>
            </launchTemplateData>
            <launchTemplateId>lt-0a20c965061f64abc</launchTemplateId>
            <launchTemplateName>MyLaunchTemplate</launchTemplateName>
            <versionDescription>AMIOnlyv1</versionDescription>
            <versionNumber>2</versionNumber>
        </item>
        <item>
            <createTime>2017-10-31T11:55:15.000Z</createTime>
            <createdBy>arn:aws:iam::123456789012:root</createdBy>
            <defaultVersion>false</defaultVersion>
            <launchTemplateData>
                <imageId>ami-aabbccdd</imageId>
            </launchTemplateData>
            <launchTemplateId>lt-0a20c965061f64abc</launchTemplateId>
            <launchTemplateName>MyLaunchTemplate</launchTemplateName>
            <versionDescription>AMIOnlyv2</versionDescription>
            <versionNumber>3</versionNumber>
        </item>
    </launchTemplateVersionSet>
</DescribeLaunchTemplateVersionsResponse>
```

### Example 2

This example describes all the latest versions of the launch templates in your
account.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DescribeLaunchTemplateVersions
&LaunchTemplateVersion.1=$Latest
&AUTHPARAMS
```

#### Sample Response

```

<DescribeLaunchTemplateVersionsResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>65cadec1-b364-4354-8ca8-4176dexample</requestId>
    <launchTemplateVersionSet>
        <item>
            <createTime>2020-01-31T11:38:52.000Z</createTime>
            <createdBy>arn:aws:iam::123456789012:root</createdBy>
            <defaultVersion>true</defaultVersion>
            <launchTemplateData>
                <imageId>ami-8c1be5f6</imageId>
                <instanceType>t2.micro</instanceType>
            </launchTemplateData>
            <launchTemplateId>lt-1111111111EXAMPLE</launchTemplateId>
            <launchTemplateName>MyLaunchTemplate1</launchTemplateName>
            <versionDescription>FirstTemplate</versionDescription>
            <versionNumber>1</versionNumber>
        </item>
        <item>
            <createTime>2020-02-14T11:52:03.000Z</createTime>
            <createdBy>arn:aws:iam::123456789012:root</createdBy>
            <defaultVersion>false</defaultVersion>
            <launchTemplateData>
                <imageId>ami-12345678</imageId>
                <instanceType>t2.micro</instanceType>
            </launchTemplateData>
            <launchTemplateId>lt-2222222222EXAMPLE</launchTemplateId>
            <launchTemplateName>MyLaunchTemplate2</launchTemplateName>
            <versionDescription>ThirdVersion</versionDescription>
            <versionNumber>3</versionNumber>
        </item>
        <item>
            <createTime>2020-03-03T11:55:15.000Z</createTime>
            <createdBy>arn:aws:iam::123456789012:root</createdBy>
            <defaultVersion>false</defaultVersion>
            <launchTemplateData>
                <imageId>ami-aabbccdd</imageId>
                <instanceType>t3.small</instanceType>
            </launchTemplateData>
            <launchTemplateId>lt-3333333333EXAMPLE</launchTemplateId>
            <launchTemplateName>MyLaunchTemplate3</launchTemplateName>
            <versionDescription>AMIOnlyv2</versionDescription>
            <versionNumber>2</versionNumber>
        </item>
    </launchTemplateVersionSet>
</DescribeLaunchTemplateVersionsResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](https://docs.aws.amazon.com/goto/cli2/ec2-2016-11-15/DescribeLaunchTemplateVersions)

- [AWS SDK for .NET V4](https://docs.aws.amazon.com/goto/DotNetSDKV4/ec2-2016-11-15/DescribeLaunchTemplateVersions)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/describelaunchtemplateversions.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/describelaunchtemplateversions.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/describelaunchtemplateversions.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/describelaunchtemplateversions.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/describelaunchtemplateversions.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/describelaunchtemplateversions.md)

- [AWS SDK for Python](https://docs.aws.amazon.com/goto/boto3/ec2-2016-11-15/DescribeLaunchTemplateVersions)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/describelaunchtemplateversions.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DescribeLaunchTemplates

DescribeLocalGatewayRouteTables
