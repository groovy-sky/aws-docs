# Application Access

By default, WorkSpaces Applications enables the applications that you specify in your image to launch other applications and executable files on the image builder and fleet instance. This ensures that applications with dependencies on other applications (for example, an application that launches the browser to navigate to a product website) function as expected. Make sure that you configure your administrative controls, security groups, and other security software to grant users the minimum permissions required to access resources and transfer data between their local computers and fleet instances.

You can use application control software, such as [Microsoft AppLocker](https://docs.microsoft.com/en-us/windows/security/threat-protection/windows-defender-application-control/applocker/applocker-overview), and policies to control which applications and files your users can run. Application control software and policies help you control the executable files, scripts, Windows installer files, dynamic-link libraries, and application packages that your users can run on WorkSpaces Applications image builders and fleet instances.

###### Note

The WorkSpaces Applications agent software relies on the Windows command prompt and Windows Powershell to provision streaming instances. If you choose to prevent users from launching the Windows command prompt or Windows Powershell, the policies must not apply to the Windows NT AUTHORITY\\SYSTEM or users in the Administrators group.

Rule typeActionWindows user or groupName/PathConditionDescriptionExecutableAllowNT AUTHORITY\\System\*PathRequired for the WorkSpaces Applications agent softwareExecutableAllowBUILTIN\\Administrators\*PathRequired for the WorkSpaces Applications agent softwareExecutableAllowEveryone%PROGRAMFILES%\\nodejs\\\*PathRequired for the WorkSpaces Applications agent softwareExecutableAllowEveryone%PROGRAMFILES%\\NICE\\\*PathRequired for the WorkSpaces Applications agent softwareExecutableAllowEveryone%PROGRAMFILES%\\Amazon\\\*PathRequired for the WorkSpaces Applications agent softwareExecutableAllowEveryone%PROGRAMFILES%\\< `default-browser` >\\\*PathRequired for the WorkSpaces Applications agent software when persistent storage solutions, such as Google Drive or Microsoft OneDrive for Business, are used. This exception is not required when WorkSpaces Applications home folders are used.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Administrator Controls

Identity and Access Management

All content copied from https://docs.aws.amazon.com/.
