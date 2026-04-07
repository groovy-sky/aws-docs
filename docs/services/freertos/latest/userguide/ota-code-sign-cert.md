# Create a code-signing certificate

To digitally sign firmware images, you need a code-signing certificate and
private key. For testing purposes, you can create a self‐signed certificate
and private key. For production environments, purchase a certificate through a
well‐known certificate authority (CA).

Different platforms require different types of code-signing certificates. The
following sections describe how to create code-signing certificates for different
FreeRTOS-qualified platforms.

###### Topics

- [Creating a code-signing certificate for the Texas Instruments CC3220SF-LAUNCHXL](https://docs.aws.amazon.com/freertos/latest/userguide/ota-code-sign-cert-ti.html)

- [Creating a code-signing certificate for the Espressif ESP32](https://docs.aws.amazon.com/freertos/latest/userguide/ota-code-sign-cert-esp.html)

- [Creating a code-signing certificate for the Nordic nrf52840-dk](https://docs.aws.amazon.com/freertos/latest/userguide/ota-code-sign-cert-nordic.html)

- [Creating a code-signing certificate for the FreeRTOS Windows simulator](https://docs.aws.amazon.com/freertos/latest/userguide/ota-code-sign-cert-win.html)

- [Creating a code-signing certificate for custom hardware](https://docs.aws.amazon.com/freertos/latest/userguide/ota-code-sign-cert-other.html)

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Create an OTA user policy

Creating a code-signing certificate for the Texas Instruments CC3220SF-LAUNCHXL
