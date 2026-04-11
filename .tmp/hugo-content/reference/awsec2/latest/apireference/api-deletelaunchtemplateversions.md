# DeleteLaunchTemplateVersions

Deletes one or more versions of a launch template.

You can't delete the default version of a launch template; you must first assign a
different version as the default. If the default version is the only version for the
launch template, you must delete the entire launch template using [DeleteLaunchTemplate](api-deletelaunchtemplate.md).

You can delete up to 200 launch template versions in a single request. To delete more
than 200 versions in a single request, use [DeleteLaunchTemplate](api-deletelaunchtemplate.md), which
deletes the launch template and all of its versions.

For more information, see [Delete a launch template version](../../../../services/ec2/latest/userguide/delete-launch-template.md#delete-launch-template-version) in the
_Amazon EC2 User Guide_.

## Request Parameters

The following parameters are for this specific action. For more information about required and optional parameters that are common to all actions, see [Common Query Parameters](commonparameters.md).

**DryRun**

Checks whether you have the required permissions for the action, without actually
making the request, and provides an error response. If you have the required
permissions, the error response is `DryRunOperation`. Otherwise, it is
`UnauthorizedOperation`.

Type: Boolean

Required: No

**LaunchTemplateId**

The ID of the launch template.

You must specify either the launch template ID or the launch template name, but not
both.

Type: String

Required: No

**LaunchTemplateName**

The name of the launch template.

You must specify either the launch template ID or the launch template name, but not
both.

Type: String

Length Constraints: Minimum length of 3. Maximum length of 128.

Pattern: `[a-zA-Z0-9\(\)\.\-/_]+`

Required: No

**LaunchTemplateVersion.N**

The version numbers of one or more launch template versions to delete. You can specify
up to 200 launch template version numbers.

Type: Array of strings

Required: Yes

## Response Elements

The following elements are returned by the service.

**requestId**

The ID of the request.

Type: String

**successfullyDeletedLaunchTemplateVersionSet**

Information about the launch template versions that were successfully deleted.

Type: Array of [DeleteLaunchTemplateVersionsResponseSuccessItem](api-deletelaunchtemplateversionsresponsesuccessitem.md) objects

**unsuccessfullyDeletedLaunchTemplateVersionSet**

Information about the launch template versions that could not be deleted.

Type: Array of [DeleteLaunchTemplateVersionsResponseErrorItem](api-deletelaunchtemplateversionsresponseerroritem.md) objects

## Errors

For information about the errors that are common to all actions, see [Common client error codes](errors-overview.md#CommonErrors).

## Examples

### Example

This example deletes version 3 of launch template
`lt-0a20c965061f64abc`.

#### Sample Request

```

https://ec2.amazonaws.com/?Action=DeleteLaunchTemplateVersions
&LaunchTemplateId=lt-0a20c965061f64abc
&LaunchTemplateVersion.1=3
&AUTHPARAMS
```

#### Sample Response

```

<DeleteLaunchTemplateVersionsResponse xmlns="http://ec2.amazonaws.com/doc/2016-11-15/">
    <requestId>67fc746a-1b3f-467e-8583-7061cexample</requestId>
    <unsuccessfullyDeletedLaunchTemplateVersionSet/>
    <successfullyDeletedLaunchTemplateVersionSet>
        <item>
            <launchTemplateId>lt-0a20c965061f64abc</launchTemplateId>
            <launchTemplateName>MyLaunchTemplate</launchTemplateName>
            <versionNumber>3</versionNumber>
        </item>
    </successfullyDeletedLaunchTemplateVersionSet>
</DeleteLaunchTemplateVersionsResponse>
```

## See Also

For more information about using this API in one of the language-specific AWS SDKs, see the following:

- [AWS Command Line Interface V2](../../../../services/goto/cli2/ec2-2016-11-15/deletelaunchtemplateversions.md)

- [AWS SDK for .NET V4](../../../goto/dotnetsdkv4/ec2-2016-11-15/deletelaunchtemplateversions.md)

- [AWS SDK for C++](../../../goto/sdkforcpp/ec2-2016-11-15/deletelaunchtemplateversions.md)

- [AWS SDK for Go v2](../../../goto/sdkforgov2/ec2-2016-11-15/deletelaunchtemplateversions.md)

- [AWS SDK for Java V2](../../../goto/sdkforjavav2/ec2-2016-11-15/deletelaunchtemplateversions.md)

- [AWS SDK for JavaScript V3](../../../goto/sdkforjavascriptv3/ec2-2016-11-15/deletelaunchtemplateversions.md)

- [AWS SDK for Kotlin](../../../goto/sdkforkotlin/ec2-2016-11-15/deletelaunchtemplateversions.md)

- [AWS SDK for PHP V3](../../../goto/sdkforphpv3/ec2-2016-11-15/deletelaunchtemplateversions.md)

- [AWS SDK for Python](../../../../services/goto/boto3/ec2-2016-11-15/deletelaunchtemplateversions.md)

- [AWS SDK for Ruby V3](../../../goto/sdkforrubyv3/ec2-2016-11-15/deletelaunchtemplateversions.md)

[Document Conventions](../../../../general/general/latest/gr/docconventions.md)

DeleteLaunchTemplate

DeleteLocalGatewayRoute
