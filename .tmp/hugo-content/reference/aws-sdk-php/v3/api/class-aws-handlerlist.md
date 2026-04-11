Menu

- [Aws](namespace-aws.md)

## HandlerList        in package    - [Aws](package-aws.md)       implements  Countable

Builds a single handler function from zero or more middleware functions and
a handler. The handler function is then used to send command objects and
return a promise that is resolved with an AWS result object.

The "front" of the list is invoked before the "end" of the list. You can add
middleware to the front of the list using one of the "prepend" method, and
the end of the list using one of the "append" method. The last function
invoked in a handler list is the handler (a function that does not accept a
next handler but rather is responsible for returning a promise that is
fulfilled with an Aws\\ResultInterface object).

Handlers are ordered using a "step" that describes the step at which the
SDK is when sending a command. The available steps are:

- init: The command is being initialized, allowing you to do things like add
default options.
- validate: The command is being validated before it is serialized
- build: The command is being serialized into an HTTP request. A middleware
in this step MUST serialize an HTTP request and populate the "@request"
parameter of a command with the request such that it is available to
subsequent middleware.
- sign: The request is being signed and prepared to be sent over the wire.

Middleware can be registered with a name to allow you to easily add a
middleware before or after another middleware by name. This also allows you
to remove a middleware by name (in addition to removing by instance).

### Table of Contents  [header link](class-aws-handlerlist-toc.md)

#### Interfaces  [header link](class-aws-handlerlist-toc-interfaces.md)

Countable

#### Constants  [header link](class-aws-handlerlist-toc-constants.md)

[ATTEMPT](class-aws-handlerlist-constant-attempt.md)
= 'attempt' [BUILD](class-aws-handlerlist-constant-build.md)
= 'build' [INIT](class-aws-handlerlist-constant-init.md)
= 'init' [SIGN](class-aws-handlerlist-constant-sign.md)
= 'sign' [VALIDATE](class-aws-handlerlist-constant-validate.md)
= 'validate'

#### Methods  [header link](class-aws-handlerlist-toc-methods.md)

[\_\_construct()](class-aws-handlerlist-method-construct.md)
: mixed [\_\_toString()](class-aws-handlerlist-method-tostring.md)
: string Dumps a string representation of the list.[after()](class-aws-handlerlist-method-after.md)
: mixed Add a middleware after the given middleware by name.[appendAttempt()](class-aws-handlerlist-method-appendattempt.md)
: mixed Append a middleware to the attempt step.[appendBuild()](class-aws-handlerlist-method-appendbuild.md)
: mixed Append a middleware to the build step.[appendInit()](class-aws-handlerlist-method-appendinit.md)
: mixed Append a middleware to the init step.[appendSign()](class-aws-handlerlist-method-appendsign.md)
: mixed Append a middleware to the sign step.[appendValidate()](class-aws-handlerlist-method-appendvalidate.md)
: mixed Append a middleware to the validate step.[before()](class-aws-handlerlist-method-before.md)
: mixed Add a middleware before the given middleware by name.[count()](class-aws-handlerlist-method-count.md)
: int [hasHandler()](class-aws-handlerlist-method-hashandler.md)
: bool Returns true if the builder has a handler.[hasMiddleware()](class-aws-handlerlist-method-hasmiddleware.md)
: bool Checks if a middleware exists. The middleware
should have been added with a name in order to
use this method.[interpose()](class-aws-handlerlist-method-interpose.md)
: mixed Interpose a function between each middleware (e.g., allowing for a trace
through the middleware layers).[prependAttempt()](class-aws-handlerlist-method-prependattempt.md)
: mixed Prepend a middleware to the attempt step.[prependBuild()](class-aws-handlerlist-method-prependbuild.md)
: mixed Prepend a middleware to the build step.[prependInit()](class-aws-handlerlist-method-prependinit.md)
: mixed Prepend a middleware to the init step.[prependSign()](class-aws-handlerlist-method-prependsign.md)
: mixed Prepend a middleware to the sign step.[prependValidate()](class-aws-handlerlist-method-prependvalidate.md)
: mixed Prepend a middleware to the validate step.[remove()](class-aws-handlerlist-method-remove.md)
: mixed Remove a middleware by name or by instance from the list.[resolve()](class-aws-handlerlist-method-resolve.md)
: callable Compose the middleware and handler into a single callable function.[setHandler()](class-aws-handlerlist-method-sethandler.md)
: mixed Set the HTTP handler that actually returns a response.

### Constants  [header link](class-aws-handlerlist-constants.md)

#### ATTEMPT  [header link](class-aws-handlerlist-constant-attempt.md)

`
    public
        mixed
    ATTEMPT
    = 'attempt'
`

#### BUILD  [header link](class-aws-handlerlist-constant-build.md)

`
    public
        mixed
    BUILD
    = 'build'
`

#### INIT  [header link](class-aws-handlerlist-constant-init.md)

`
    public
        mixed
    INIT
    = 'init'
`

#### SIGN  [header link](class-aws-handlerlist-constant-sign.md)

`
    public
        mixed
    SIGN
    = 'sign'
`

#### VALIDATE  [header link](class-aws-handlerlist-constant-validate.md)

`
    public
        mixed
    VALIDATE
    = 'validate'
`

### Methods  [header link](class-aws-handlerlist-methods.md)

#### \_\_construct()  [header link](class-aws-handlerlist-method-construct.md)

`
    public
                    __construct([callable $handler = null ]) : mixed`

##### Parameters

$handler
: callable
= null

HTTP handler.

#### \_\_toString()  [header link](class-aws-handlerlist-method-tostring.md)

Dumps a string representation of the list.

`
    public
                    __toString() : string`

##### Return values

string

#### after()  [header link](class-aws-handlerlist-method-after.md)

Add a middleware after the given middleware by name.

`
    public
                    after(string|callable $findName, string $withName, callable $middleware) : mixed`

##### Parameters

$findName
: string\|callable

Add after this

$withName
: string

Optional name to give the middleware

$middleware
: callable

Middleware to add.

#### appendAttempt()  [header link](class-aws-handlerlist-method-appendattempt.md)

Append a middleware to the attempt step.

`
    public
                    appendAttempt(callable $middleware[, string $name = null ]) : mixed`

##### Parameters

$middleware
: callable

Middleware function to add.

$name
: string
= null

Name of the middleware.

#### appendBuild()  [header link](class-aws-handlerlist-method-appendbuild.md)

Append a middleware to the build step.

`
    public
                    appendBuild(callable $middleware[, string $name = null ]) : mixed`

##### Parameters

$middleware
: callable

Middleware function to add.

$name
: string
= null

Name of the middleware.

#### appendInit()  [header link](class-aws-handlerlist-method-appendinit.md)

Append a middleware to the init step.

`
    public
                    appendInit(callable $middleware[, string $name = null ]) : mixed`

##### Parameters

$middleware
: callable

Middleware function to add.

$name
: string
= null

Name of the middleware.

#### appendSign()  [header link](class-aws-handlerlist-method-appendsign.md)

Append a middleware to the sign step.

`
    public
                    appendSign(callable $middleware[, string $name = null ]) : mixed`

##### Parameters

$middleware
: callable

Middleware function to add.

$name
: string
= null

Name of the middleware.

#### appendValidate()  [header link](class-aws-handlerlist-method-appendvalidate.md)

Append a middleware to the validate step.

`
    public
                    appendValidate(callable $middleware[, string $name = null ]) : mixed`

##### Parameters

$middleware
: callable

Middleware function to add.

$name
: string
= null

Name of the middleware.

#### before()  [header link](class-aws-handlerlist-method-before.md)

Add a middleware before the given middleware by name.

`
    public
                    before(string|callable $findName, string $withName, callable $middleware) : mixed`

##### Parameters

$findName
: string\|callable

Add before this

$withName
: string

Optional name to give the middleware

$middleware
: callable

Middleware to add.

#### count()  [header link](class-aws-handlerlist-method-count.md)

`
    public
                    count() : int`

##### Return values

int

#### hasHandler()  [header link](class-aws-handlerlist-method-hashandler.md)

Returns true if the builder has a handler.

`
    public
                    hasHandler() : bool`

##### Return values

bool

#### hasMiddleware()  [header link](class-aws-handlerlist-method-hasmiddleware.md)

Checks if a middleware exists. The middleware
should have been added with a name in order to
use this method.

`
    public
                    hasMiddleware(string $name) : bool`

##### Parameters

$name
: string

##### Return values

bool

#### interpose()  [header link](class-aws-handlerlist-method-interpose.md)

Interpose a function between each middleware (e.g., allowing for a trace
through the middleware layers).

`
    public
                    interpose([callable|null $fn = null ]) : mixed`

The interpose function is a function that accepts a "step" argument as a
string and a "name" argument string. This function must then return a
function that accepts the next handler in the list. This function must
then return a function that accepts a CommandInterface and optional
RequestInterface and returns a promise that is fulfilled with an
Aws\\ResultInterface or rejected with an Aws\\Exception\\AwsException
object.

##### Parameters

$fn
: callable\|null
= null

Pass null to remove any previously set function

#### prependAttempt()  [header link](class-aws-handlerlist-method-prependattempt.md)

Prepend a middleware to the attempt step.

`
    public
                    prependAttempt(callable $middleware[, string $name = null ]) : mixed`

##### Parameters

$middleware
: callable

Middleware function to add.

$name
: string
= null

Name of the middleware.

#### prependBuild()  [header link](class-aws-handlerlist-method-prependbuild.md)

Prepend a middleware to the build step.

`
    public
                    prependBuild(callable $middleware[, string $name = null ]) : mixed`

##### Parameters

$middleware
: callable

Middleware function to add.

$name
: string
= null

Name of the middleware.

#### prependInit()  [header link](class-aws-handlerlist-method-prependinit.md)

Prepend a middleware to the init step.

`
    public
                    prependInit(callable $middleware[, string $name = null ]) : mixed`

##### Parameters

$middleware
: callable

Middleware function to add.

$name
: string
= null

Name of the middleware.

#### prependSign()  [header link](class-aws-handlerlist-method-prependsign.md)

Prepend a middleware to the sign step.

`
    public
                    prependSign(callable $middleware[, string $name = null ]) : mixed`

##### Parameters

$middleware
: callable

Middleware function to add.

$name
: string
= null

Name of the middleware.

#### prependValidate()  [header link](class-aws-handlerlist-method-prependvalidate.md)

Prepend a middleware to the validate step.

`
    public
                    prependValidate(callable $middleware[, string $name = null ]) : mixed`

##### Parameters

$middleware
: callable

Middleware function to add.

$name
: string
= null

Name of the middleware.

#### remove()  [header link](class-aws-handlerlist-method-remove.md)

Remove a middleware by name or by instance from the list.

`
    public
                    remove(string|callable $nameOrInstance) : mixed`

##### Parameters

$nameOrInstance
: string\|callable

Middleware to remove.

#### resolve()  [header link](class-aws-handlerlist-method-resolve.md)

Compose the middleware and handler into a single callable function.

`
    public
                    resolve() : callable`

##### Return values

callable

#### setHandler()  [header link](class-aws-handlerlist-method-sethandler.md)

Set the HTTP handler that actually returns a response.

`
    public
                    setHandler(callable $handler) : mixed`

##### Parameters

$handler
: callable

Function that accepts a request and array of
options and returns a Promise.

<\-\- modeled\_exceptions -->

×

**On this page**

- Table Of Contents
  - [Constants](class-aws-handlerlist-toc-constants.md)
  - [Methods](class-aws-handlerlist-toc-methods.md)
- Constants
  - [ATTEMPT](class-aws-handlerlist-constant-attempt.md)
  - [BUILD](class-aws-handlerlist-constant-build.md)
  - [INIT](class-aws-handlerlist-constant-init.md)
  - [SIGN](class-aws-handlerlist-constant-sign.md)
  - [VALIDATE](class-aws-handlerlist-constant-validate.md)
- Methods
  - [\_\_construct()](class-aws-handlerlist-method-construct.md)
  - [\_\_toString()](class-aws-handlerlist-method-tostring.md)
  - [after()](class-aws-handlerlist-method-after.md)
  - [appendAttempt()](class-aws-handlerlist-method-appendattempt.md)
  - [appendBuild()](class-aws-handlerlist-method-appendbuild.md)
  - [appendInit()](class-aws-handlerlist-method-appendinit.md)
  - [appendSign()](class-aws-handlerlist-method-appendsign.md)
  - [appendValidate()](class-aws-handlerlist-method-appendvalidate.md)
  - [before()](class-aws-handlerlist-method-before.md)
  - [count()](class-aws-handlerlist-method-count.md)
  - [hasHandler()](class-aws-handlerlist-method-hashandler.md)
  - [hasMiddleware()](class-aws-handlerlist-method-hasmiddleware.md)
  - [interpose()](class-aws-handlerlist-method-interpose.md)
  - [prependAttempt()](class-aws-handlerlist-method-prependattempt.md)
  - [prependBuild()](class-aws-handlerlist-method-prependbuild.md)
  - [prependInit()](class-aws-handlerlist-method-prependinit.md)
  - [prependSign()](class-aws-handlerlist-method-prependsign.md)
  - [prependValidate()](class-aws-handlerlist-method-prependvalidate.md)
  - [remove()](class-aws-handlerlist-method-remove.md)
  - [resolve()](class-aws-handlerlist-method-resolve.md)
  - [setHandler()](class-aws-handlerlist-method-sethandler.md)

[Back To Top](class-aws-handlerlist-top.md)
