# External credentials

External credentials is a generic authentication plugin that you can use to connect to any
external SAML based identity provider. To use the plugin, you pass an executable file that
returns a SAML response.

## Authentication type

**Connection string name****Parameter type****Default value****Connection string example**AuthenticationTypeRequired`IAM Credentials``AuthenticationType=External Credentials;`

## Executable path

The path to the executable that has the logic of your custom SAML-based credential
provider. The output of the executable must be the parsed SAML response from the
identity provider.

**Connection string name****Parameter type****Default value****Connection string example**ExecutablePathRequired`none`ExecutablePath=C:\\Users\ `user_name`\ `external_credential.exe`

## Argument list

The list of arguments that you want to pass to the executable.

**Connection string name****Parameter type****Default value****Connection string example**ArgumentListOptional`none``ArgumentList=arg1
                                arg2
                                arg3`

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Default credentials

Instance profile

All content copied from https://docs.aws.amazon.com/.
