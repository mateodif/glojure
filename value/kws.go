package value

var (
	KWTag       = NewKeyword("tag")
	KWFile      = NewKeyword("file")
	KWLine      = NewKeyword("line")
	KWColumn    = NewKeyword("column")
	KWEndLine   = NewKeyword("end-line")
	KWEndColumn = NewKeyword("end-column")

	KWMethods       = NewKeyword("methods")
	KWIsVariadic    = NewKeyword("variadic?")
	KWMaxFixedArity = NewKeyword("max-fixed-arity")
	KWLocal         = NewKeyword("local")
	KWName          = NewKeyword("name")
	KWFixedArity    = NewKeyword("fixed-arity")
	KWBody          = NewKeyword("body")
	KWParams        = NewKeyword("params")
	KWOp            = NewKeyword("op")
	KWForm          = NewKeyword("form")
	KWConst         = NewKeyword("const")
	KWDef           = NewKeyword("def")
	KWSetBang       = NewKeyword("set!")
	KWMaybeClass    = NewKeyword("maybe-class")
	KWWithMeta      = NewKeyword("with-meta")
	KWArgs          = NewKeyword("args")
	KWBindings      = NewKeyword("bindings")
	KWCase          = NewKeyword("case")
	KWClass         = NewKeyword("class")
	KWDefault       = NewKeyword("default")
	KWDo            = NewKeyword("do")
	KWElse          = NewKeyword("else")
	KWException     = NewKeyword("exception")
	KWExpr          = NewKeyword("expr")
	KWExprs         = NewKeyword("exprs")
	KWField         = NewKeyword("field")
	KWFinally       = NewKeyword("finally")
	KWFn            = NewKeyword("fn")
	KWHostCall      = NewKeyword("host-call")
	KWHostInterop   = NewKeyword("host-interop")
	KWIf            = NewKeyword("if")
	KWInit          = NewKeyword("init")
	KWInvoke        = NewKeyword("invoke")
	KWItems         = NewKeyword("items")
	KWKeys          = NewKeyword("keys")
	KWLet           = NewKeyword("let")
	KWLoop          = NewKeyword("loop")
	KWMOrF          = NewKeyword("m-or-f")
	KWMap           = NewKeyword("map")
	KWMaybeHostForm = NewKeyword("maybe-host-form")
	KWMeta          = NewKeyword("meta")
	KWMethod        = NewKeyword("method")
	KWNew           = NewKeyword("new")
	KWNodes         = NewKeyword("nodes")
	KWQuote         = NewKeyword("quote")
	KWRecur         = NewKeyword("recur")
	KWRet           = NewKeyword("ret")
	KWSet           = NewKeyword("set")
	KWStatements    = NewKeyword("statements")
	KWTarget        = NewKeyword("target")
	KWTest          = NewKeyword("test")
	KWTests         = NewKeyword("tests")
	KWTheVar        = NewKeyword("the-var")
	KWThen          = NewKeyword("then")
	KWThrow         = NewKeyword("throw")
	KWTry           = NewKeyword("try")
	KWVal           = NewKeyword("val")
	KWVals          = NewKeyword("vals")
	KWVar           = NewKeyword("var")
	KWVector        = NewKeyword("vector")

	KWMacro   = NewKeyword("macro")
	KWPrivate = NewKeyword("private")
	KWDynamic = NewKeyword("dynamic")
	KWNS      = NewKeyword("ns")
)