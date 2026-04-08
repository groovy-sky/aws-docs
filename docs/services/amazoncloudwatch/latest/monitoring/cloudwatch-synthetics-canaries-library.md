# Synthetics runtime versions

When you create or update a canary, you choose a Synthetics runtime version for the
canary.
A Synthetics runtime is a combination of the Synthetics code that calls your script handler,
and the Lambda layers of bundled dependencies.

CloudWatch Synthetics currently supports runtimes that use Node.js, Python, or Java languages.
The frameworks supported are Puppeteer, Playwright, and Selenium.

We recommend that you always use the most recent runtime version for your canaries, to
be
able to use the latest features and updates made to the Synthetics library.

**Please note**: whenever you run a canary to use the new
version of the Synthetics runtime, all Synthetics library functions that your canary uses
are also automatically moved to the same version of NodeJS that the Synthetics runtime
supports.

###### Topics

- [Runtime versions using Java](cloudwatch-synthetics-library-java.md)

- [Runtime versions using Node.js and Playwright](cloudwatch-synthetics-library-nodejs-playwright.md)

- [Runtime versions using Node.js and Puppeteer](cloudwatch-synthetics-library-nodejs-puppeteer.md)

- [Runtime versions using Python and Selenium Webdriver](cloudwatch-synthetics-library-python-selenium.md)

- [Runtime versions using Node.js](cloudwatch-synthetics-library-nodejs.md)

- [Runtime versions support policy](cloudwatch-synthetics-runtime-support-policy.md)

- [Runtime versions update](cloudwatch-synthetics-runtime-version-update.md)

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Using the CloudWatch Synthetics Recorder for Google Chrome

Runtime versions using Java

All content copied from https://docs.aws.amazon.com/.
