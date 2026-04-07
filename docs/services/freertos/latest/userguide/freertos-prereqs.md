# First steps

###### Important

This page refers to the Amazon-FreeRTOS repository which is deprecated. We recommend
that you [start here](freertos-getting-started-modular.md) when you
create a new project. If you already have an existing FreeRTOS project based on the now
deprecated Amazon-FreeRTOS repository, see the [Amazon-FreeRTOS Github Repository Migration Guide](github-repo-migration.md).

To get started using FreeRTOS with AWS IoT, you must have an AWS account, a user with permissions to
access AWS IoT and FreeRTOS cloud services.
You also must download FreeRTOS and configure your board's FreeRTOS demo project to work with AWS IoT. The
following sections walk you through these requirements.

###### Note

- If you're using the Espressif ESP32-DevKitC, ESP-WROVER-KIT, or the ESP32-WROOM-32SE, skip these steps
and go to [Getting started with the Espressif ESP32-DevKitC and the ESP-WROVER-KIT](getting-started-espressif.md).

- If you're using the Nordic nRF52840-DK, skip these steps and go to [Getting started with the Nordic nRF52840-DK](https://docs.aws.amazon.com/freertos/latest/userguide/getting_started_nordic.html).

1. [Setting up your AWS account and permissions](#freertos-account-and-permissions)

2. [Registering your MCU board with AWS IoT](#get-started-freertos-thing)

3. [Downloading FreeRTOS](#freertos-download)

4. [Configuring the FreeRTOS demos](#freertos-configure)

## Setting up your AWS account and permissions

### Sign up for an AWS account

If you do not have an AWS account, complete the following steps to create one.

###### To sign up for an AWS account

1. Open [https://portal.aws.amazon.com/billing/signup](https://portal.aws.amazon.com/billing/signup).

2. Follow the online instructions.

Part of the sign-up procedure involves receiving a phone call or text message and entering
    a verification code on the phone keypad.

When you sign up for an AWS account, an _AWS account root user_ is created. The root user has access to all AWS services
    and resources in the account. As a security best practice, assign administrative access to a user, and use only the root user to perform [tasks that require root user access](../../../iam/latest/userguide/id-root-user.md#root-user-tasks).

AWS sends you a confirmation email after the sign-up process is
complete. At any time, you can view your current account activity and manage your account by
going to [https://aws.amazon.com/](https://aws.amazon.com/) and choosing **My**
**Account**.

### Create a user with administrative access

After you sign up for an AWS account, secure your AWS account root user, enable AWS IAM Identity Center, and create an administrative user so that you
don't use the root user for everyday tasks.

###### Secure your AWS account root user

1. Sign in to the [AWS Management Console](https://console.aws.amazon.com/) as the account owner by choosing **Root user** and entering your AWS account email address. On the next page, enter your password.

For help signing in by using root user, see [Signing in as the root user](https://docs.aws.amazon.com/signin/latest/userguide/console-sign-in-tutorials.html#introduction-to-root-user-sign-in-tutorial) in the _AWS Sign-In User Guide_.

2. Turn on multi-factor authentication (MFA) for your root user.

For instructions, see [Enable a virtual MFA device for your AWS account root user (console)](https://docs.aws.amazon.com/IAM/latest/UserGuide/enable-virt-mfa-for-root.html) in the _IAM User Guide_.

###### Create a user with administrative access

1. Enable IAM Identity Center.

For instructions, see [Enabling\
    AWS IAM Identity Center](https://docs.aws.amazon.com/singlesignon/latest/userguide/get-set-up-for-idc.html) in the
    _AWS IAM Identity Center User Guide_.

2. In IAM Identity Center, grant administrative access to a user.

For a tutorial about using the IAM Identity Center directory as your identity source, see [Configure user access with the default IAM Identity Center directory](https://docs.aws.amazon.com/singlesignon/latest/userguide/quick-start-default-idc.html) in the
    _AWS IAM Identity Center User Guide_.

###### Sign in as the user with administrative access

- To sign in with your IAM Identity Center user, use the sign-in URL that was sent to your email address when you created the IAM Identity Center user.

For help signing in using an IAM Identity Center user, see [Signing in to the AWS access portal](https://docs.aws.amazon.com/signin/latest/userguide/iam-id-center-sign-in-tutorial.html) in the _AWS Sign-In User Guide_.

###### Assign access to additional users

1. In IAM Identity Center, create a permission set that follows the best practice of applying least-privilege permissions.

For instructions, see [Create a permission set](https://docs.aws.amazon.com/singlesignon/latest/userguide/get-started-create-a-permission-set.html) in the _AWS IAM Identity Center User Guide_.

2. Assign users to a group, and then assign single sign-on access to the group.

For instructions, see [Add groups](https://docs.aws.amazon.com/singlesignon/latest/userguide/addgroups.html) in the _AWS IAM Identity Center User Guide_.

To provide access, add permissions to your users, groups, or roles:

- Users and groups in AWS IAM Identity Center:

Create a permission set. Follow the instructions in [Create a permission set](https://docs.aws.amazon.com/singlesignon/latest/userguide/howtocreatepermissionset.html) in the _AWS IAM Identity Center User Guide_.

- Users managed in IAM through an identity provider:

Create a role for identity federation. Follow the instructions in [Create a role for a third-party identity provider (federation)](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_create_for-idp.html)
in the _IAM User Guide_.

- IAM users:

- Create a role that your user can assume. Follow the instructions in [Create a role for an IAM user](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_roles_create_for-user.html) in the _IAM User Guide_.

- (Not recommended) Attach a policy directly to a user or add a user to a user group. Follow the instructions in [Adding permissions to a user (console)](https://docs.aws.amazon.com/IAM/latest/UserGuide/id_users_change-permissions.html#users_change_permissions-add-console) in the _IAM User Guide_.

## Registering your MCU board with AWS IoT

Your board must be registered with AWS IoT to communicate with the AWS Cloud. To register your board with AWS IoT, you
must have:

**An AWS IoT policy**

The AWS IoT policy grants your device permissions to access AWS IoT resources. It is stored on the AWS Cloud.

**An AWS IoT thing**

An AWS IoT thing allows you to manage your devices in AWS IoT. It is stored on the AWS Cloud.

**A private key and X.509 certificate**

The private key and certificate allow your device to authenticate with AWS IoT.

To register your board, follow the procedures below.

###### To create an AWS IoT policy

1. To create an IAM policy, you must know your AWS Region and AWS account number.

To find your AWS account number, open the [AWS Management \
    Console](https://console.aws.amazon.com/), locate and expand the menu beneath your account name in the upper-right corner, and choose
    **My Account**. Your account ID is displayed under **Account Settings**.

To find the AWS region for your AWS account, use the AWS Command Line Interface. To install the AWS CLI, follow the
    instructions in the [AWS Command Line Interface User \
    Guide](https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-install.html). After you install the AWS CLI, open a command prompt window and enter the following command:

```bash

aws iot describe-endpoint --endpoint-type=iot:Data-ATS
```

The output should look like this:

```json

{
       "endpointAddress": "xxxxxxxxxxxxxx-ats.iot.us-west-2.amazonaws.com"
}
```

In this example, the region is `us-west-2`.

###### Note

We recommend using ATS endpoints as seen in the example.

2. Browse to the [AWS IoT console](https://console.aws.amazon.com/iotv2).

3. In the navigation pane, choose **Secure**, choose **Policies**, and
    then choose **Create**.

4. Enter a name to identify your policy.

5. In the **Add statements** section, choose **Advanced mode**. Copy and
    paste the following JSON into the policy editor window. Replace `aws-region` and
    `aws-account` with your AWS Region and account ID.
JSON

```json

{
       "Version":"2012-10-17",
       "Statement": [
           {
               "Effect": "Allow",
               "Action": "iot:Connect",
               "Resource": "arn:aws:iot:us-east-1:123456789012:*"
           },
           {
               "Effect": "Allow",
               "Action": "iot:Publish",
               "Resource": "arn:aws:iot:us-east-1:123456789012:*"
           },
           {
               "Effect": "Allow",
               "Action": "iot:Subscribe",
               "Resource": "arn:aws:iot:us-east-1:123456789012:*"
           },
           {
               "Effect": "Allow",
               "Action": "iot:Receive",
               "Resource": "arn:aws:iot:us-east-1:123456789012:*"
           }
       ]
}

```

This policy grants the following permissions:

**`iot:Connect`**

Grants your device the permission to connect to the AWS IoT message broker with any client ID.

**`iot:Publish`**

Grants your device the permission to publish an MQTT message on any MQTT topic.

**`iot:Subscribe`**

Grants your device the permission to subscribe to any MQTT topic filter.

**`iot:Receive`**

Grants your device the permission to receive messages from the
AWS IoT message broker on any MQTT topic.

6. Choose **Create**.

###### To create an IoT thing, private key, and certificate for your device

01. Browse to the [AWS IoT console](https://console.aws.amazon.com/iotv2).

02. In the navigation pane, choose **Manage**, and then choose **Things**.

03. If you do not have any IoT things registered in your account, the **You don't have any things**
    **yet** page is displayed. If you see this page, choose **Register a thing**. Otherwise,
     choose **Create**.

04. On the **Creating AWS IoT things** page, choose **Create a single thing**.

05. On the **Add your device to the thing registry** page, enter a name for your thing, and then
     choose **Next**.

06. On the **Add a certificate for your thing** page, under
     **One-click certificate creation**, choose **Create certificate**.

07. Download your private key and certificate by choosing the **Download** links for each.

08. Choose **Activate** to activate your certificate. Certificates must be activated prior to
     use.

09. Choose **Attach a policy** to attach a policy to your certificate that grants your device
     access to AWS IoT operations.

10. Choose the policy you just created and choose **Register thing**.

After your board is registered with AWS IoT, you can continue to [Downloading FreeRTOS](#freertos-download).

## Downloading FreeRTOS

You can download FreeRTOS from the [FreeRTOS \
GitHub repository](https://github.com/freertos/freertos).

After you download FreeRTOS, you can continue to [Configuring the FreeRTOS demos](#freertos-configure).

## Configuring the FreeRTOS demos

You must edit some configuration files in your FreeRTOS directory before you can compile and run any demos on your
board.

###### To configure your AWS IoT endpoint

You must provide FreeRTOS with your AWS IoT endpoint so the application running on your board can send requests
to the correct endpoint.

1. Browse to the [AWS IoT console](https://console.aws.amazon.com/iotv2).

2. In the left navigation pane, choose **Settings**.

Your AWS IoT endpoint is displayed in **Device data endpoint**. It should look like
    `1234567890123-ats.iot.us-east-1.amazonaws.com`.
    Make a note of this endpoint.

3. In the navigation pane, choose **Manage**, and then choose **Things**.

Your device should have an AWS IoT thing name. Make a note of this name.

4. Open `demos/include/aws_clientcredential.h`.

5. Specify values for the following constants:

- `#define clientcredentialMQTT_BROKER_ENDPOINT "Your AWS IoT endpoint";`

- `#define clientcredentialIOT_THING_NAME "The AWS IoT thing name of your board"`

###### To configure your Wi-Fi

If your board is connecting to the internet across a Wi-Fi connection, you must provide FreeRTOS with Wi-Fi
credentials to connect to the network. If your board does not support Wi-Fi, you can skip these steps.

1. `demos/include/aws_clientcredential.h`.

2. Specify values for the following `#define` constants:

- `#define clientcredentialWIFI_SSID "The SSID for your Wi-Fi
                          network"`

- `#define clientcredentialWIFI_PASSWORD "The password for your Wi-Fi
                          network"`

- `#define clientcredentialWIFI_SECURITY` `The security type of your Wi-Fi
                          network`

Valid security types are:

- `eWiFiSecurityOpen` (Open, no security)

- `eWiFiSecurityWEP` (WEP security)

- `eWiFiSecurityWPA` (WPA security)

- `eWiFiSecurityWPA2` (WPA2 security)

###### To format your AWS IoT credentials

FreeRTOS must have the AWS IoT certificate and private keys associated with your registered thing and its permissions
policies to successfully communicate with AWS IoT on behalf of your device.

###### Note

To configure your AWS IoT credentials, you must have the private key and certificate that you downloaded from
the AWS IoT console when you registered your device. After you have registered your device as an AWS IoT thing,
you can retrieve device certificates from the AWS IoT console, but you cannot retrieve private keys.

FreeRTOS is a C language project, and the certificate and private key must be specially formatted to be added to
the project.

1. In a browser window, open
    `tools/certificate_configuration/CertificateConfigurator.html`.

2. Under **Certificate PEM file**, choose the
    `ID-certificate.pem.crt` that you downloaded from the AWS IoT
    console.

3. Under **Private Key PEM file**, choose the
    `ID-private.pem.key` that you downloaded from the AWS IoT
    console.

4. Choose **Generate and save aws\_clientcredential\_keys.h**, and then save the file in
    `demos/include`. This overwrites the existing file in the directory.

###### Note

The certificate and private key are hard-coded for demonstration purposes only. Production-level
applications should store these files in a secure location.

After you configure FreeRTOS, you can continue to the Getting Started guide for your board to set up your platform's
hardware and its software development environment, and then compile and run the demo on your board. For board-specific
instructions, see the [Board-specific getting started guides](https://docs.aws.amazon.com/freertos/latest/userguide/getting-started-guides.html). The demo
application that is used in the Getting Started tutorial is the coreMQTT Mutual Authentication demo, which is located
at `demos/coreMQTT/mqtt_demo_mutual_auth.c`.

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

FreeRTOS demo application

Troubleshooting
