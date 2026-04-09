# Supported Features

The following table compares the features that are supported by the different access
types.

FeatureBrowser-Based AccessClient-Based Access for WindowsClient-Based Access for macOSNotesEnterprise Deployement Tool✗✓✗For more information, see [Tutorial: Install the Amazon WorkSpaces Applications Client And Customize the Client Experience for Your Users](install-client-configure-settings.md).HIPAA/PCI compliance✓✓✓For more information, see [Compliance](https://aws.amazon.com/appstream2/faqs).Active Directory authentication✓✓✓For more information, see [Using Active Directory with WorkSpaces Applications](active-directory.md).MFA (multi-factor authentication)✓✓✓For WorkSpaces Applications, MFA is supported via SAML 2.0.Smart card (CAC and PIV readers)✗✓✗For more information, see [Smart Card Redirection](feature-support-usb-devices-qualified.md#feature-support-USB-devices-qualified-smart-cards-support).Certificates for access control (OS-based)✓✓✓For WorkSpaces Applications, certificate authentication is supported via SAML
2.0.Certificate-Based Authentication✓✓✓For more information, see [Certificate-Based Authentication](certificate-based-authentication.md).Client customizationAvailable with exceptionsAvailable with exceptionsAvailable with exceptionsWorkSpaces Applications supports web-based branding and custom URLs. For more
information, see [Add Your Custom Branding to Amazon WorkSpaces Applications](branding.md).Desktop view connection mode✓✓✓Classic application mode✓✓✓Native application mode✗✓✗

USB redirection

✗✓✗Supported on WorkSpaces Applications client accessing Windows-based fleets. For more
information, see [USB Redirection](feature-support-usb-devices-qualified.md#feature-support-USB-devices-USB-redirection).

Audio input (for web conferences and calls)

✓✓✓Not supported on Linux. WorkSpaces Applications supports USB microphones.

Video input (conferencing applications)

✓✓✓

Storage redirection

✗✓✗For more information, see [Enable File System Redirection for Your WorkSpaces Applications Users](enable-file-system-redirection.md).

USB/local printer redirection

Available with exceptions✓✓WorkSpaces Applications indirect printing on browser. Full redirection not supported
on Linux-based stacks.Clipboard redirection✓✓✓Drawing tablets✓✓✗YubiKey support✗Available with exceptions✗Supported on WorkSpaces Applications client. For more information, see [Qualify USB Devices for Use with Streaming Applications](qualify-usb-devices.md).Monitor supportWeb access dual-monitor support. For more information, see [Dual-Monitor Support](dual-monitor-support-web-access-admin.md).✓✓For more information, see [Multiple Monitors](feature-support-multiple-monitors.md).User Pool✓✓✓For more information, see [Amazon WorkSpaces Applications User Pools](user-pool.md).Connect to an App Block Builder✓✓✗For more information, see [App Block Builder](app-block-builder.md).Access Multi-Stack Application✓✓✓For more information, see [Attribute-Based Application Entitlements Using a Third-Party SAML 2.0 Identity Provider](application-entitlements-saml.md).

The next topics provide information about how to configure user access to WorkSpaces Applications for application streaming.

For guidance that you can provide your users to help them get started with application streaming, see [Guidance for WorkSpaces Applications Users](user-guidance.md).

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Provide User Access

Provide Access Through a Web Browser

All content copied from https://docs.aws.amazon.com/.
