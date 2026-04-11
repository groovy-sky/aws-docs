# VpcConnector

Describes an AWS App Runner VPC connector resource. A VPC connector describes the Amazon Virtual Private Cloud (Amazon VPC) that an App Runner service is
associated with, and the subnets and security group that are used.

Multiple revisions of a connector might have the same `Name` and different `Revision` values.

###### Note

At this time, App Runner supports only one revision per name.

## Contents

**CreatedAt**

The time when the VPC connector was created. It's in Unix time stamp format.

Type: Timestamp

Required: No

**DeletedAt**

The time when the VPC connector was deleted. It's in Unix time stamp format.

Type: Timestamp

Required: No

**SecurityGroups**

A list of IDs of security groups that App Runner uses for access to AWS resources under the specified subnets. If not specified, App Runner uses the default
security group of the Amazon VPC. The default security group allows all outbound traffic.

Type: Array of strings

Length Constraints: Minimum length of 0. Maximum length of 51200.

Pattern: `.*`

Required: No

**Status**

The current state of the VPC connector. If the status of a connector revision is `INACTIVE`, it was deleted and can't be
used. Inactive connector revisions are permanently removed some time after they are deleted.

Type: String

Valid Values: `ACTIVE | INACTIVE`

Required: No

**Subnets**

A list of IDs of subnets that App Runner uses for your service. All IDs are of subnets of a single Amazon VPC.

Type: Array of strings

Length Constraints: Minimum length of 0. Maximum length of 51200.

Pattern: `.*`

Required: No

**VpcConnectorArn**

The Amazon Resource Name (ARN) of this VPC connector.

Type: String

Length Constraints: Minimum length of 1. Maximum length of 1011.

Pattern: `arn:aws(-[\w]+)*:[a-z0-9-\\.]{0,63}:[a-z0-9-\\.]{0,63}:[0-9]{12}:(\w|\/|-){1,1011}`

Required: No

**VpcConnectorName**

The customer-provided VPC connector name.

Type: String

Length Constraints: Minimum length of 4. Maximum length of 40.

Pattern: `[A-Za-z0-9][A-Za-z0-9\-_]{3,39}`

Required: No

**VpcConnectorRevision**

The revision of this VPC connector. It's unique among all the active connectors ( `"Status": "ACTIVE"`) that share the same
`Name`.

###### Note

At this time, App Runner supports only one revision per name.

Type: Integer

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../../reference/goto/sdkforcpp/apprunner-2020-05-15/vpcconnector.md)

- [AWS SDK for Java V2](../../../../reference/goto/sdkforjavav2/apprunner-2020-05-15/vpcconnector.md)

- [AWS SDK for Ruby V3](../../../../reference/goto/sdkforrubyv3/apprunner-2020-05-15/vpcconnector.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

TraceConfiguration

VpcDNSTarget

All content copied from https://docs.aws.amazon.com/.
