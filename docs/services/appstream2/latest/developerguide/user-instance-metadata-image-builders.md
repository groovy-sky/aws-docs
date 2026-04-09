# Instance Metadata for WorkSpaces Applications Image Builders

WorkSpaces Applications image builder instances have instance metadata available through Windows
environment variables. You can use the following environment variables in your applications
and scripts to modify your environment based on the image builder instance details.

Environment VariableContextDescriptionAppStream\_Image\_ArnMachineThe ARN of the image that was used to create the streaming
instance.AppStream\_Instance\_TypeMachineThe instance type of the streaming instance. For example,
`stream.standard.medium`.AppStream\_Resource\_TypeMachineThe type of WorkSpaces Applications resource. The value is either `fleet
                        ` or `imagebuilder`.AppStream\_Resource\_NameMachineThe name of the image builder.

On Linux image builders, environment variables are exported through the script at
**/etc/profile.d/appstream\_system\_vars.sh**. To access the environment
variables, you can explicitly source this file in your application.

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Image Builder Actions

Install AMD Driver on Graphics Design Instances

All content copied from https://docs.aws.amazon.com/.
