# Configuring utilities for the `APPSYNC_JS` runtime

AWS AppSync provides two libraries that aid in the development of resolvers with the `APPSYNC_JS`
runtime:

- `@aws-appsync/eslint-plugin` \- Catches and fixes problems quickly during
development.

- `@aws-appsync/utils` \- Provides type validation and autocompletion in code editors.

## Configuring the eslint plugin

[ESLint](https://eslint.org/) is a tool that statically analyzes your code to quickly
find problems. You can run ESLint as part of your continuous integration pipeline.
`@aws-appsync/eslint-plugin` is an ESLint plugin that catches invalid syntax in your code when
leveraging the `APPSYNC_JS` runtime. The plugin allows you to quickly get feedback about your
code during development without having to push your changes to the cloud.

`@aws-appsync/eslint-plugin` provides two rule sets that you can use during development.

**"plugin:@aws-appsync/base"** configures a base set of rules that you can
leverage in your project:

RuleDescriptionno-asyncAsync processes and promises are not supported.no-awaitAsync processes and promises are not supported.no-classesClasses are not supported.no-for`for` is not supported (except for `for-in` and `for-of`,
which are supported)no-continue`continue` is not supported.no-generatorsGenerators are not supported.no-yield`yield` is not supported.no-labelsLabels are not supported.no-this`this` keyword is not supported.no-tryTry/catch structure is not supported.no-whileWhile loops are not supported.no-disallowed-unary-operators`++`, `--`, and `~` unary operators are not
allowed.no-disallowed-binary-operatorsThe `instanceof` operator is not allowed.no-promiseAsync processes and promises are not supported.

**"plugin:@aws-appsync/recommended"** provides some additional rules but
also requires you to add TypeScript configurations to your project.

RuleDescriptionno-recursionRecursive function calls are not allowed.no-disallowed-methodsSome methods are not allowed. See the [reference](resolver-util-reference-js.md) for a full
set of supported built-in functions.no-function-passingPassing functions as function arguments to functions is not allowed.no-function-reassignFunctions cannot be reassigned.no-function-returnFunctions cannot be the return value of functions.

To add the plugin to your project, follow the installation and usage steps at [Getting Started\
with ESLint](https://eslint.org/docs/latest/user-guide/getting-started). Then, install the [plugin](https://www.npmjs.com/package/@aws-appsync/eslint-plugin) in your project using your
project package manager (e.g., npm, yarn, or pnpm):

```Sh

$ npm install @aws-appsync/eslint-plugin
```

In your `.eslintrc.{js,yml,json}` file, add **"plugin:@aws-appsync/base"** or **"plugin:@aws-appsync/recommended"** to the `extends` property. The snippet below is
a basic sample `.eslintrc` configuration for JavaScript:

```TypeScript

{
  "extends": ["plugin:@aws-appsync/base"]
}
```

To use the **"plugin:@aws-appsync/recommended"** rule set, install the
required dependency:

```Sh

$ npm install -D @typescript-eslint/parser
```

Then, create an `.eslintrc.js` file:

```TypeScript

{
  "parser": "@typescript-eslint/parser",
  "parserOptions": {
    "ecmaVersion": 2018,
    "project": "./tsconfig.json"
  },
  "extends": ["plugin:@aws-appsync/recommended"]
}
```

[Document Conventions](../../../../general/latest/gr/docconventions.md)

Example pipeline resolver with Amazon DynamoDB

Bundling, TypeScript, and source maps for the APPSYNC\_JS runtime

All content copied from https://docs.aws.amazon.com/.
