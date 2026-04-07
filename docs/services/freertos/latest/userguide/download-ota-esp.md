# Download, build, flash, and run the FreeRTOS OTA demo on the Espressif ESP32

###### Important

This reference integration is hosted on the Amazon-FreeRTOS repository which is deprecated.
We recommend that you [start here](https://docs.aws.amazon.com/freertos/latest/userguide/freertos-getting-started-modular.html) when
you create a new project. If you already have an existing FreeRTOS project based on the now
deprecated Amazon-FreeRTOS repository, see the [Amazon-FreeRTOS Github Repository Migration Guide](https://docs.aws.amazon.com/freertos/latest/userguide/github-repo-migration.html).

1. Download the FreeRTOS source from [GitHub](https://github.com/aws/amazon-freertos). See
    the [README.md](https://github.com/aws/amazon-freertos/blob/main/README.md) file for
    instructions. Create a project in your IDE that includes all required sources and libraries.

2. Follow the instructions in [Getting Started with \
    Espressif](https://docs.aws.amazon.com/freertos/latest/userguide/getting_started_espressif.html) to set up the required GCC-based toolchain.

3. Open
    `freertos/vendors/vendor/boards/board/aws_demos/config_files/aws_demo_config.h`,
    comment out `#define CONFIG_CORE_MQTT_MUTUAL_AUTH_DEMO_ENABLED`, and define
    `CONFIG_OTA_MQTT_UPDATE_DEMO_ENABLED` or `CONFIG_OTA_HTTP_UPDATE_DEMO_ENABLED`.

4. Build the demo project by running `make` in the
    `vendors/espressif/boards/esp32/aws_demos` directory. You can flash the demo program
    and verify its output by running `make flash monitor`, as described in [Getting Started with \
    Espressif](https://docs.aws.amazon.com/freertos/latest/userguide/getting_started_espressif.html).

5. Before running the OTA Update demo:

- Open
`freertos/vendors/vendor/boards/board/aws_demos/config_files/aws_demo_config.h`,
comment out `#define CONFIG_CORE_MQTT_MUTUAL_AUTH_DEMO_ENABLED`, and define
`CONFIG_OTA_MQTT_UPDATE_DEMO_ENABLED` or `CONFIG_OTA_HTTP_UPDATE_DEMO_ENABLED`.

- Open `vendors/vendor/boards/board/aws_demos/config_files/ota_demo_config.h`,
and copy your SHA-256/ECDSA code-signing certificate in:

```c

#define otapalconfigCODE_SIGNING_CERTIFICATE [] = "your-certificate-key";
```

[Document Conventions](https://docs.aws.amazon.com/general/latest/gr/docconventions.html)

Microchip Curiosity PIC32MZEF

Renesas RX65N
