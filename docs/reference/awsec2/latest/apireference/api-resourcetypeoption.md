# ResourceTypeOption

The options that affect the scope of the response.

## Contents

**OptionName**

The name of the option.

- For `ec2:Instance`:

Specify `state-name` \- The current state of the EC2 instance.

- For `ec2:LaunchTemplate`:

Specify `version-depth` \- The number of launch template versions to check,
starting from the most recent version.

Type: String

Valid Values: `state-name | version-depth`

Required: No

**OptionValue.N**

A value for the specified option.

- For `state-name`:

- Valid values: `pending` \| `running` \| `shutting-down` \|
`terminated` \| `stopping` \| `stopped`

- Default: All states

- For `version-depth`:

- Valid values: Integers between `1` and `10000`

- Default: `10`

Type: Array of strings

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/resourcetypeoption.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/resourcetypeoption.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/resourcetypeoption.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

ResourceStatementRequest

ResourceTypeRequest
