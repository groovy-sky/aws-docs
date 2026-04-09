# Application

Describes an application in the application catalog.

## Contents

**AppBlockArn**

The app block ARN of the application.

Type: String

Pattern: `^arn:aws(?:\-cn|\-iso\-b|\-iso|\-us\-gov)?:[A-Za-z0-9][A-Za-z0-9_/.-]{0,62}:[A-Za-z0-9_/.-]{0,63}:[A-Za-z0-9_/.-]{0,63}:[A-Za-z0-9][A-Za-z0-9:_/+=,@.\\-]{0,1023}$`

Required: No

**Arn**

The ARN of the application.

Type: String

Pattern: `^arn:aws(?:\-cn|\-iso\-b|\-iso|\-us\-gov)?:[A-Za-z0-9][A-Za-z0-9_/.-]{0,62}:[A-Za-z0-9_/.-]{0,63}:[A-Za-z0-9_/.-]{0,63}:[A-Za-z0-9][A-Za-z0-9:_/+=,@.\\-]{0,1023}$`

Required: No

**CreatedTime**

The time at which the application was created within the app block.

Type: Timestamp

Required: No

**Description**

The description of the application.

Type: String

Length Constraints: Minimum length of 1.

Required: No

**DisplayName**

The application name to display.

Type: String

Length Constraints: Minimum length of 1.

Required: No

**Enabled**

If there is a problem, the application can be disabled after image creation.

Type: Boolean

Required: No

**IconS3Location**

The S3 location of the application icon.

Type: [S3Location](api-s3location.md) object

Required: No

**IconURL**

The URL for the application icon. This URL might be time-limited.

Type: String

Length Constraints: Minimum length of 1.

Required: No

**InstanceFamilies**

The instance families for the application.

Type: Array of strings

Length Constraints: Minimum length of 1.

Required: No

**LaunchParameters**

The arguments that are passed to the application at launch.

Type: String

Length Constraints: Minimum length of 1.

Required: No

**LaunchPath**

The path to the application executable in the instance.

Type: String

Length Constraints: Minimum length of 1.

Required: No

**Metadata**

Additional attributes that describe the application.

Type: String to string map

Key Length Constraints: Minimum length of 1.

Value Length Constraints: Minimum length of 1.

Required: No

**Name**

The name of the application.

Type: String

Length Constraints: Minimum length of 1.

Required: No

**Platforms**

The platforms on which the application can run.

Type: Array of strings

Array Members: Maximum number of 4 items.

Valid Values: `WINDOWS | WINDOWS_SERVER_2016 | WINDOWS_SERVER_2019 | WINDOWS_SERVER_2022 | WINDOWS_SERVER_2025 | AMAZON_LINUX2 | RHEL8 | ROCKY_LINUX8 | UBUNTU_PRO_2404`

Required: No

**WorkingDirectory**

The working directory for the application.

Type: String

Length Constraints: Minimum length of 1.

Required: No

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS SDK for C++](../../../goto/sdkforcpp/appstream-2016-12-01/application.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/appstream-2016-12-01/application.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/appstream-2016-12-01/application.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

AppBlockBuilderStateChangeReason

ApplicationConfig

All content copied from https://docs.aws.amazon.com/.
