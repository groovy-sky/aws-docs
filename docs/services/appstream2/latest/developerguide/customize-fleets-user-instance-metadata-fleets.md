# User and Instance Metadata for Amazon WorkSpaces Applications Fleets

WorkSpaces Applications fleet instances have user and instance metadata available through Windows
environment variables. You can use the following environment variables in your
applications and scripts to modify your environment based on the fleet instance
details.

Environment VariableContextDescriptionAppStream\_Stack\_NameUserThe name of the stack from which the streaming session
started.AppStream\_User\_Access\_ModeUserThe access mode used to manage user access to the stream. The
available values are `custom`,
`userpool`, or `
                            saml`.AppStream\_Session\_Reservation\_DateTimeUserThe date and time when the user's streaming session started.AppStream\_UserNameUserThe user name associated with the user.AppStream\_Session\_IDUserThe session identifier for the user's streaming session.APPSTREAM\_SESSION\_CONTEXTMachineContains the parameters passed to your streaming application when a
session is started. For more information, see [Session Context in Amazon WorkSpaces Applications](managing-stacks-fleets-session-context.md).

###### Note

This environment variable is only available after the first
application launch.

AppStream\_Image\_ArnMachineThe ARN of the image that was used to create the streaming
instance.AppStream\_Instance\_TypeMachineThe streaming instance's type. For example,
`stream.standard.medium`.AppStream\_Resource\_TypeMachineThe type of WorkSpaces Applications resource. The value is either `fleet
                            ` or `image-builder`.AppStream\_Resource\_NameMachineThe fleet's name.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Use the WorkSpaces Applications Template User Account to Change the Default Internet Explorer Home Page

Update a Fleet

All content copied from https://docs.aws.amazon.com/.
