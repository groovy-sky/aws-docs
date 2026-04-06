# Getting started with AWS IoT ExpressLink

The following sections will guide you through the required steps to connect your AWS IoT ExpressLink
Evaluation Kit to the cloud to send and receive data directly with your AWS account. For module
specific guides, see the Getting Started Guide for your specific AWS IoT ExpressLink hardware module
in the [AWS Partner Device \
Catalog](https://partners.amazonaws.com/qualified-devices).

If your ExpressLink module manufacturer supplies an Evaluation Kit, you will want to follow the
specific steps provided there. Also, the general steps you follow for a manufacturer-specific
Evaluation Kit are provided in section "11.2.1 Run the Quick Connect demo application" of the
[AWS IoT \
ExpressLink Programmer's Guide](https://docs.aws.amazon.com/iot-expresslink/latest/programmersguide/elpg-provisioning.html).

If you have questions or issues that are not answered here, please visit the
[AWS re:Post \
for AWS IoT ExpressLink](https://repost.aws/tags/TADqOo0ODORl2pC69DWwUFug/aws-iot-expresslink) page.

###### Topics

- [Set up your host machine](#elgsg-host-app-set-up)

- [AWS account set up and console login](https://docs.aws.amazon.com/iot-expresslink/latest/gettingstartedguide/elgsg-set-up.html)

- [Register an AWS IoT ExpressLink module](https://docs.aws.amazon.com/iot-expresslink/latest/gettingstartedguide/elgsg-register-account.html)

- [Connect and interact with AWS Cloud.](https://docs.aws.amazon.com/iot-expresslink/latest/gettingstartedguide/elgsg-interact-with-cloud.html)

- [Perform a firmware OTA update](https://docs.aws.amazon.com/iot-expresslink/latest/gettingstartedguide/elgsg-setup-ota-update.html)

## Set up your host machine

AWS IoT ExpressLink evaluation kits can be connected to a host machine serial interface using a
terminal application.

**Prerequisites**: To establish a serial interface
connection between your host machine and the evaluation kit, you must install the
corresponding USB to UART bridge Virtual Communication Port drivers. Refer to your
hardware module's _Getting Started Guide_ for any specific
requirements.

1. Open a terminal application for your host machine (for example, TeraTerm for Windows,
    CoolTerm for Mac) and select the port corresponding to the evaluation kit.

2. Configure the terminal application with the following settings:

```nohighlight

Baudrate:     115,200
Bits:          8
Parity:        None
Stop:          1
Flow control:  None
Local Echo:    Yes
End of Line:   Line Feed
```

3. To check your connection, enter the following command from the terminal:

```nohighlight

AT
```

If you receive the answer 'OK', then you've successfully connected the evaluation
    kit to your host machine.

Keep the terminal window open. You'll use the terminal later in this procedure.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

AWS account set up and console login
