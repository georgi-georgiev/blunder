package blunder

//    Problem type definitions MAY extend the problem details object with
//    additional members.

//    For example, our "out of credit" problem above defines two such
//    extensions -- "balance" and "accounts" to convey additional, problem-
//    specific information.

//    Clients consuming problem details MUST ignore any such extensions
//    that they don't recognize; this allows problem types to evolve and
//    include additional information in the future.

//    Note that because extensions are effectively put into a namespace by
//    the problem type, it is not possible to define new "standard" members
//    without defining a new media type.

// 	  When an HTTP API needs to define a response that indicates an error
//    condition, it might be appropriate to do so by defining a new problem
//    type.

//    Before doing so, it's important to understand what they are good for,
//    and what's better left to other mechanisms.

//    Problem details are not a debugging tool for the underlying
//    implementation; rather, they are a way to expose greater detail about
//    the HTTP interface itself.  Designers of new problem types need to
//    carefully consider the Security Considerations (Section 5), in
//    particular, the risk of exposing attack vectors by exposing
//    implementation internals through error messages.

//    Likewise, truly generic problems -- i.e., conditions that could
//    potentially apply to any resource on the Web -- are usually better
//    expressed as plain status codes.  For example, a "write access
//    disallowed" problem is probably unnecessary, since a 403 Forbidden
//    status code in response to a PUT request is self-explanatory.

//    Finally, an application might have a more appropriate way to carry an
//    error in a format that it already defines.  Problem details are
//    intended to avoid the necessity of establishing new "fault" or
//    "error" document formats, not to replace existing domain-specific
//    formats.

//    That said, it is possible to add support for problem details to
//    existing HTTP APIs using HTTP content negotiation (e.g., using the
//    Accept request header to indicate a preference for this format; see
//    [RFC7231], Section 5.3.2).

func Extend(typeURI string, title string, status int) {

}
