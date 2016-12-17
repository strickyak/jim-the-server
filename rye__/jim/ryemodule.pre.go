// +build prego

package jim

import . "github.com/strickyak/rye/rye__"

// cwp: github.com/strickyak/jim-the-server
// path: github.com/strickyak/jim-the-server/jim
// thispkg: github.com/strickyak/jim-the-server/rye__/jim
// modname: jim
// internal: None

import "fmt"
import "io"
import "os"
import "reflect"
import "runtime"
import "time"
import "bytes"
import i_http "net/http" // <parse.Timport object at 0x7fbae19bc310>
// DIR ['__class__', '__delattr__', '__dict__', '__doc__', '__format__', '__getattribute__', '__hash__', '__init__', '__module__', '__new__', '__reduce__', '__reduce_ex__', '__repr__', '__setattr__', '__sizeof__', '__str__', '__subclasshook__', '__weakref__', 'alias', 'fromWhere', 'gloss', 'imported', 'line', 'pkg', 'visit', 'where'] // VARS {'imported': ['go', 'net', 'http'], 'gloss': 'from', 'fromWhere': 'go', 'alias': 'http', 'pkg': 'net/http', 'line': 21, 'where': 730}
import i_os "os" // <parse.Timport object at 0x7fbae19bc390>
// DIR ['__class__', '__delattr__', '__dict__', '__doc__', '__format__', '__getattribute__', '__hash__', '__init__', '__module__', '__new__', '__reduce__', '__reduce_ex__', '__repr__', '__setattr__', '__sizeof__', '__str__', '__subclasshook__', '__weakref__', 'alias', 'fromWhere', 'gloss', 'imported', 'line', 'pkg', 'visit', 'where'] // VARS {'imported': ['go', 'os'], 'gloss': 'from', 'fromWhere': 'go', 'alias': 'os', 'pkg': 'os', 'line': 21, 'where': 730}
import i_time "time" // <parse.Timport object at 0x7fbae19bc3d0>
// DIR ['__class__', '__delattr__', '__dict__', '__doc__', '__format__', '__getattribute__', '__hash__', '__init__', '__module__', '__new__', '__reduce__', '__reduce_ex__', '__repr__', '__setattr__', '__sizeof__', '__str__', '__subclasshook__', '__weakref__', 'alias', 'fromWhere', 'gloss', 'imported', 'line', 'pkg', 'visit', 'where'] // VARS {'imported': ['go', 'time'], 'gloss': 'from', 'fromWhere': 'go', 'alias': 'time', 'pkg': 'time', 'line': 21, 'where': 730}
var _ = i_http.ErrWriteAfterFlush // net/http
var _ = i_os.Stdout               // os
var _ = i_time.UTC                // time
var _ = fmt.Sprintf
var _ = io.EOF
var _ = os.Stderr
var _ = reflect.ValueOf
var _ = runtime.Stack
var _ = time.Sleep
var _ = bytes.Split
var _ = MkInt

var eval_module_once bool

func Eval_Module() M {
	if eval_module_once == false {
		eval_module_once = true
		_ = inner_eval_module()
	}
	return ModuleObj
}
func inner_eval_module() M {
	// @ 730 @ 21 @
	//Vimport: ['go', 'net', 'http'] http go
	//Vimport: ['go', 'net', 'http'] http go
	// $ 730 $ 21 $
	// @ 730 @ 21 @
	//Vimport: ['go', 'os'] os go
	//Vimport: ['go', 'os'] os go
	// $ 730 $ 21 $
	// @ 730 @ 21 @
	//Vimport: ['go', 'time'] time go
	//Vimport: ['go', 'time'] time go
	// $ 730 $ 21 $
	// @ 765 @ 23 @
	// $ 765 $ 23 $
	return None
}

//(begin tail)
///////////////////////////////
// name: loggingWrapper
var specNest_nesting_105 = CallSpec{Name: "loggingWrapper__nesting_105", Args: []string{"w", "r"}, Defaults: []M{MissingM, MissingM}, Star: "", StarStar: ""}

type pNest_nesting_105 struct {
	PNewCallable
	fn func(a_w M, a_r M) M
}

func (o *pNest_nesting_105) Contents() interface{} {
	return o.fn
}
func (o pNest_nesting_105) Call2(a0 M, a1 M) M {
	return o.fn(a0, a1)
}

func (o pNest_nesting_105) CallV(a1 []M, a2 []M, kv1 []KV, kv2 map[string]M) M {
	argv, star, starstar := NewSpecCall(o.CallSpec, a1, a2, kv1, kv2)
	_, _, _ = argv, star, starstar
	return o.fn(argv[0], argv[1])
}

//(tail)
///////////////////////////////
// name: makeLoggingWrapper
var specNest_nesting_104 = CallSpec{Name: "makeLoggingWrapper__nesting_104", Args: []string{"pattern", "webDir", "handler"}, Defaults: []M{MissingM, MissingM, MissingM}, Star: "", StarStar: ""}

type pNest_nesting_104 struct {
	PNewCallable
	fn func(a_pattern M, a_webDir M, a_handler M) M
}

func (o *pNest_nesting_104) Contents() interface{} {
	return o.fn
}
func (o pNest_nesting_104) Call3(a0 M, a1 M, a2 M) M {
	return o.fn(a0, a1, a2)
}

func (o pNest_nesting_104) CallV(a1 []M, a2 []M, kv1 []KV, kv2 map[string]M) M {
	argv, star, starstar := NewSpecCall(o.CallSpec, a1, a2, kv1, kv2)
	_, _, _ = argv, star, starstar
	return o.fn(argv[0], argv[1], argv[2])
}

//(tail)
// zip(p.argsPlus, typPlus): [('args', None)]
// typPlus: [None]
///////////////////////////////

func G_1_main(a_args M) M {
	var v_bindHostColonPort M = None
	_ = v_bindHostColonPort
	var v_command M = None
	_ = v_command
	var v_confFile M = None
	_ = v_confFile
	var v_detuple_101 M = None
	_ = v_detuple_101
	var v_handler M = None
	_ = v_handler
	var v_line M = None
	_ = v_line
	var v_makeLoggingWrapper M = None
	_ = v_makeLoggingWrapper
	var v_pattern M = None
	_ = v_pattern
	var v_webDir M = None
	_ = v_webDir
	var v_words M = None
	_ = v_words
	// @ 783 @ 24 @ main
	// Vvar: local var args -> 'a_args'
	// @@@@@@ Creating var "detuple_101" in scope @@@@@@
	v_detuple_101 = /*Yvar.str*/ a_args // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <class 'codegen.Yvar'>
	// Vvar: local var detuple_101 -> 'v_detuple_101'
	len_detuple_101 := /*Yvar.str*/ v_detuple_101.Len()
	if len_detuple_101 != 2 {
		panic(fmt.Sprintf("Assigning object of length %d to %d variables, in destructuring assignment.", len_detuple_101, 2))
	}
	// Vvar: local var detuple_101 -> 'v_detuple_101'
	// @@@@@@ Creating var "bindHostColonPort" in scope @@@@@@
	v_bindHostColonPort = /*Yvar.str*/ v_detuple_101.GetItem( /*Yint.str*/ litI_cfcd208495d565ef66e7dff9f98764da) // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
	// Vvar: local var detuple_101 -> 'v_detuple_101'
	// @@@@@@ Creating var "confFile" in scope @@@@@@
	v_confFile = /*Yvar.str*/ v_detuple_101.GetItem( /*Yint.str*/ litI_c4ca4238a0b923820dcc509a6f75849b) // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
	// $ 783 $ 24 $
	// @ 820 @ 25 @ main
	fmt.Fprintln(M(MkGo( /*Yimport.str*/ i_os.Stderr)).Contents().(io.Writer), "# Jim the Server")
	// $ 820 $ 25 $
	// @ 861 @ 27 @ main
	// Vcall: fn: <parse.Tfield object at 0x7fbae19bc810>
	// Vcall: args: [<parse.Tlit object at 0x7fbae19bc850>]
	// Vcall: names: ['']
	// Vcall: star: None
	// Vcall: starstar: None
	// Vcall: fn: <parse.Tfield object at 0x7fbae19bc790>
	// Vcall: args: []
	// Vcall: names: []
	// Vcall: star: None
	// Vcall: starstar: None
	// Vcall: fn: <parse.Tvar object at 0x7fbae19bc6d0>
	// Vcall: args: [<parse.Tvar object at 0x7fbae19bc710>]
	// Vcall: names: ['']
	// Vcall: star: None
	// Vcall: starstar: None
	// Making Global Yvar from 'open'
	// Vvar: local var confFile -> 'v_confFile'

	for_returning__102 := func() M { // around FOR
		var nexter__102 Nexter = /*General*/ /*invoker*/ f_INVOKE_1_split( /*General*/ /*invoker*/ F_INVOKE_0_read(CALL_1(M( /*Yvar.str*/ G_open) /*Yvar.str*/, v_confFile)) /*Ystr.str*/, litS_68b329da9893e34099c7d8ad5cb9c940).Iter()
		enougher__102, canEnough__102 := nexter__102.(Enougher)
		if canEnough__102 {
			defer enougher__102.Enough()
		}
		// else case without Enougher will be faster.
		for {
			ndx___102, more___102 := nexter__102.Next()
			if !more___102 {
				break
			}
			// BEGIN FOR

			// @@@@@@ Creating var "line" in scope @@@@@@
			v_line = ndx___102 // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
			// @ 912 @ 28 @ main
			// Vcall: fn: <parse.Tfield object at 0x7fbae19bc950>
			// Vcall: args: [<parse.Tlit object at 0x7fbae19bc990>]
			// Vcall: names: ['']
			// Vcall: star: None
			// Vcall: starstar: None
			// Vvar: local var line -> 'v_line'
			// @@@@@@ Creating var "command" in scope @@@@@@
			v_command = /*General*/ /*invoker*/ f_INVOKE_1_split( /*Yvar.str*/ v_line /*Ystr.str*/, litS_01abfc750a0c942167651c40d088531d).GetItem( /*Yint.str*/ litI_cfcd208495d565ef66e7dff9f98764da) // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
			// $ 912 $ 28 $
			// @ 975 @ 29 @ main
			// Vcall: fn: <parse.Tfield object at 0x7fbae19bcb50>
			// Vcall: args: []
			// Vcall: names: []
			// Vcall: star: None
			// Vcall: starstar: None
			// Vvar: local var command -> 'v_command'
			// @@@@@@ Creating var "words" in scope @@@@@@
			v_words = /*General*/ /*invoker*/ f_INVOKE_0_split( /*Yvar.str*/ v_command) // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
			// $ 975 $ 29 $
			// @ 1004 @ 31 @ main
			// Vvar: local var words -> 'v_words'
			if /*AsBool*/ /*Yvar.str*/ v_words.Bool() {
				// @ 1020 @ 32 @ main
				// Vvar: local var words -> 'v_words'
				// @@@@@@ Creating var "webDir" in scope @@@@@@
				v_webDir = /*Yvar.str*/ v_words.GetItem( /*Yint.str*/ litI_cfcd208495d565ef66e7dff9f98764da) // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
				// $ 1020 $ 32 $
				// @ 1044 @ 33 @ main
				// Vvar: local var words -> 'v_words'

				for_returning__103 := func() M { // around FOR
					var nexter__103 Nexter = /*Yvar.str*/ v_words.GetItemSlice( /*Yint.str*/ litI_c4ca4238a0b923820dcc509a6f75849b, None, None).Iter()
					enougher__103, canEnough__103 := nexter__103.(Enougher)
					if canEnough__103 {
						defer enougher__103.Enough()
					}
					// else case without Enougher will be faster.
					for {
						ndx___103, more___103 := nexter__103.Next()
						if !more___103 {
							break
						}
						// BEGIN FOR

						// @@@@@@ Creating var "pattern" in scope @@@@@@
						v_pattern = ndx___103 // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
						// @ 1078 @ 34 @ main
						// @@@@@@ Creating var "makeLoggingWrapper" in scope @@@@@@
						v_makeLoggingWrapper = None // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
						// zip(p.argsPlus, typPlus): [('pattern', None), ('webDir', None), ('handler', None)]
						// typPlus: [None, None, None]
						///////////////////////////////

						fn_nesting_104 := func(a_pattern M, a_webDir M, a_handler M) M {
							var v_loggingWrapper M = None
							_ = v_loggingWrapper
							// @ 1138 @ 35 @ makeLoggingWrapper
							// @@@@@@ Creating var "loggingWrapper" in scope @@@@@@
							v_loggingWrapper = None // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
							// zip(p.argsPlus, typPlus): [('w', None), ('r', None)]
							// typPlus: [None, None]
							///////////////////////////////

							fn_nesting_105 := func(a_w M, a_r M) M {
								// @ 1176 @ 36 @ loggingWrapper
								// Vcall: fn: <parse.Tfield object at 0x7fbae1956050>
								// Vcall: args: [<parse.Tfield object at 0x7fbae19560d0>]
								// Vcall: names: ['']
								// Vcall: star: None
								// Vcall: starstar: None
								// Vcall: fn: <parse.Tfield object at 0x7fbae19bcf90>
								// Vcall: args: []
								// Vcall: names: []
								// Vcall: star: None
								// Vcall: starstar: None
								// Vvar: local var webDir -> 'a_webDir'
								// Vvar: local var pattern -> 'a_pattern'
								// Vvar: local var r -> 'a_r'
								fmt.Fprintln(M(MkGo( /*Yimport.str*/ i_os.Stderr)).Contents().(io.Writer) /*ForceString*/, ( /*DoMod*/ /*Ystr.str*/ litS_3b6431eaaa10ce8061f0a1a462a6b4d3.Mod(MkTupleV( /*General*/ /*invoker*/ f_INVOKE_1_Format(MkGo(i_time.Now).Call(), MkGo( /*Yimport.str*/ i_time.Stamp)) /*Yvar.str*/, a_webDir /*Yvar.str*/, a_pattern, f_GET_RequestURI( /*Yvar.str*/ a_r)))).String())
								// $ 1176 $ 36 $
								// @ 1308 @ 38 @ loggingWrapper
								// Vcall: fn: <parse.Tfield object at 0x7fbae1956390>
								// Vcall: args: [<parse.Tvar object at 0x7fbae19563d0>, <parse.Tvar object at 0x7fbae1956410>]
								// Vcall: names: ['', '']
								// Vcall: star: None
								// Vcall: starstar: None
								// Vvar: local var handler -> 'a_handler'
								// Vvar: local var w -> 'a_w'
								// Vvar: local var r -> 'a_r'
								_ = /*General*/ /*invoker*/ f_INVOKE_2_ServeHTTP( /*Yvar.str*/ a_handler /*Yvar.str*/, a_w /*Yvar.str*/, a_r) // Assign void: = type: <type 'str'> repr: '/*General*/ /*invoker*/ f_INVOKE_2_ServeHTTP(/*Yvar.str*/a_handler, /*Yvar.str*/a_w, /*Yvar.str*/a_r) '
								// $ 1308 $ 38 $

								return None
							}

							v_loggingWrapper = MForge(&pNest_nesting_105{PNewCallable{CallSpec: &specNest_nesting_105}, fn_nesting_105}) // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
							// $ 1138 $ 35 $
							// @ 1342 @ 39 @ makeLoggingWrapper
							// Vvar: local var loggingWrapper -> 'v_loggingWrapper'
							return /*Yvar.str*/ v_loggingWrapper
							// $ 1342 $ 39 $

							return None
						}

						v_makeLoggingWrapper = MForge(&pNest_nesting_104{PNewCallable{CallSpec: &specNest_nesting_104}, fn_nesting_104}) // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
						// $ 1078 $ 34 $
						// @ 1373 @ 41 @ main
						// Vcall: fn: <parse.Tfield object at 0x7fbae1956710>
						// Vcall: args: [<parse.Tcall object at 0x7fbae1956850>]
						// Vcall: names: ['']
						// Vcall: star: None
						// Vcall: starstar: None
						// Vcall: fn: <parse.Tvar object at 0x7fbae1956750>
						// Vcall: args: [<parse.Tfield object at 0x7fbae19567d0>, <parse.Tvar object at 0x7fbae1956810>]
						// Vcall: names: ['', '']
						// Vcall: star: None
						// Vcall: starstar: None
						// Vvar: local var webDir -> 'v_webDir'
						// @@@@@@ Creating var "handler" in scope @@@@@@
						v_handler = MkGo(i_http.FileServer).Call(GoCast(GoElemType(new( /*Yimport.str*/ i_http.Dir)) /*Yvar.str*/, v_webDir)) // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
						// $ 1373 $ 41 $
						// @ 1434 @ 42 @ main
						// Vcall: fn: <parse.Tfield object at 0x7fbae1956950>
						// Vcall: args: [<parse.Tvar object at 0x7fbae1956990>, <parse.Tcall object at 0x7fbae1956ad0>]
						// Vcall: names: ['', '']
						// Vcall: star: None
						// Vcall: starstar: None
						// Vvar: local var pattern -> 'v_pattern'
						// Vcall: fn: <parse.Tvar object at 0x7fbae19569d0>
						// Vcall: args: [<parse.Tvar object at 0x7fbae1956a10>, <parse.Tvar object at 0x7fbae1956a50>, <parse.Tvar object at 0x7fbae1956a90>]
						// Vcall: names: ['', '', '']
						// Vcall: star: None
						// Vcall: starstar: None
						// Vvar: local var makeLoggingWrapper -> 'v_makeLoggingWrapper'
						// Vvar: local var pattern -> 'v_pattern'
						// Vvar: local var webDir -> 'v_webDir'
						// Vvar: local var handler -> 'v_handler'
						_ = MkGo(i_http.HandleFunc).Call( /*Yvar.str*/ v_pattern, CALL_3(M( /*Yvar.str*/ v_makeLoggingWrapper) /*Yvar.str*/, v_pattern /*Yvar.str*/, v_webDir /*Yvar.str*/, v_handler)) // Assign void: = type: <type 'str'> repr: 'MkGo(i_http.HandleFunc).Call(/*Yvar.str*/v_pattern, CALL_3( M(/*Yvar.str*/v_makeLoggingWrapper), /*Yvar.str*/v_pattern, /*Yvar.str*/v_webDir, /*Yvar.str*/v_handler )) '
						// $ 1434 $ 42 $
						// @ 1513 @ 43 @ main
						// Vvar: local var webDir -> 'v_webDir'
						// Vvar: local var pattern -> 'v_pattern'
						fmt.Fprintln(M(MkGo( /*Yimport.str*/ i_os.Stderr)).Contents().(io.Writer) /*ForceString*/, ( /*DoMod*/ /*Ystr.str*/ litS_d7ba898028d447a7ef280df1d6365b74.Mod(MkTupleV( /*Yvar.str*/ v_webDir /*Yvar.str*/, v_pattern))).String())
						// $ 1513 $ 43 $

						// END FOR
					}
					return MissingM
				}() // around FOR
				if for_returning__103 != MissingM {
					return for_returning__103
				}

				// $ 1044 $ 33 $
			}
			// $ 1004 $ 31 $

			// END FOR
		}
		return MissingM
	}() // around FOR
	if for_returning__102 != MissingM {
		return for_returning__102
	}

	// $ 861 $ 27 $
	// @ 1573 @ 45 @ main
	// Vcall: fn: <parse.Tfield object at 0x7fbae195c150>
	// Vcall: args: [<parse.Tfield object at 0x7fbae195c1d0>]
	// Vcall: names: ['']
	// Vcall: star: None
	// Vcall: starstar: None
	// Vcall: fn: <parse.Tfield object at 0x7fbae195c0d0>
	// Vcall: args: []
	// Vcall: names: []
	// Vcall: star: None
	// Vcall: starstar: None
	fmt.Fprintln(M(MkGo( /*Yimport.str*/ i_os.Stderr)).Contents().(io.Writer) /*ForceString*/, ( /*DoMod*/ /*Ystr.str*/ litS_a7b59768def62f6c9ad89722a73531a1.Mod( /*General*/ /*invoker*/ f_INVOKE_1_Format(MkGo(i_time.Now).Call(), MkGo( /*Yimport.str*/ i_time.Stamp)))).String())
	// $ 1573 $ 45 $
	// @ 1652 @ 46 @ main
	// Vcall: fn: <parse.Tfield object at 0x7fbae195c350>
	// Vcall: args: [<parse.Tvar object at 0x7fbae195c390>, <parse.Traw object at 0x7fbae195c3d0>]
	// Vcall: names: ['', '']
	// Vcall: star: None
	// Vcall: starstar: None
	// Vvar: local var bindHostColonPort -> 'v_bindHostColonPort'
	_ = MkGo(i_http.ListenAndServe).Call( /*Yvar.str*/ v_bindHostColonPort, None) // Assign void: = type: <type 'str'> repr: 'MkGo(i_http.ListenAndServe).Call(/*Yvar.str*/v_bindHostColonPort, None) '
	// $ 1652 $ 46 $

	return None
}

///////////////////////////////
// name: main

var specFunc_main = CallSpec{Name: "main", Args: []string{"args"}, Defaults: []M{MissingM}, Star: "", StarStar: ""}

type pFunc_main struct{ PNewCallable }

func (o *pFunc_main) Contents() interface{} {
	return G_main
}
func (o pFunc_main) Call1(a0 M) M {
	return G_1_main(a0)
}

func (o pFunc_main) CallV(a1 []M, a2 []M, kv1 []KV, kv2 map[string]M) M {
	argv, star, starstar := NewSpecCall(o.CallSpec, a1, a2, kv1, kv2)
	_, _, _ = argv, star, starstar
	return G_1_main(argv[0])
}

//(end tail)

var G___name__ M // M
var G_http M     // *PModule
var G_main M     // *pFunc_main
var G_os M       // *PModule
var G_time M     // *PModule

func init /*New_Module*/ () {
	G___name__ = MkStr("jim")
	G_http = None
	G_main = MForge(&pFunc_main{PNewCallable{CallSpec: &specFunc_main}})
	G_os = None
	G_time = None
}

var ModuleMap = map[string]*M{
	"__name__": &G___name__,
	"http":     &G_http,
	"main":     &G_main,
	"os":       &G_os,
	"time":     &G_time,
}

var ModuleObj = MakeModuleObject(ModuleMap, "github.com/strickyak/jim-the-server/jim")

var litI_c4ca4238a0b923820dcc509a6f75849b = MkInt(1)
var litI_cfcd208495d565ef66e7dff9f98764da = MkInt(0)
var litS_01abfc750a0c942167651c40d088531d = MkStr("#")
var litS_3b6431eaaa10ce8061f0a1a462a6b4d3 = MkStr("%s (%s == %s) %q")
var litS_68b329da9893e34099c7d8ad5cb9c940 = MkStr("\x0a")
var litS_9d222d8f9b2a29c89c3ecedd1f5470cc = MkStr("# Jim the Server")
var litS_a7b59768def62f6c9ad89722a73531a1 = MkStr("# Starting Server at %s")
var litS_d7ba898028d447a7ef280df1d6365b74 = MkStr("# %q handles %q")

func f_INVOKE_0_split(fn M) M {
	if fn.X == nil {
		if len(fn.S) == 0 {
			panic("cannot INVOKE on int")
		}
		fn = M{X: MkBStr(fn.S).Self}
	}
	switch x := fn.X.(type) {
	case i_INVOKE_0_split:
		return x.M_0_split()
	case I_GET_split:
		tmp := x.GET_split()
		return CALL_0(tmp)

	case *PGo:
		return x.Invoke("split")
	}
	panic(fmt.Sprintf("Cannot invoke 'split' with 0 arguments on %v", fn))
}

type i_INVOKE_0_split interface {
	M_0_split() M
}

func f_INVOKE_1_Format(fn M, a_0 M) M {
	if fn.X == nil {
		if len(fn.S) == 0 {
			panic("cannot INVOKE on int")
		}
		fn = M{X: MkBStr(fn.S).Self}
	}
	switch x := fn.X.(type) {
	case i_INVOKE_1_Format:
		return x.M_1_Format(a_0)
	case i_GET_Format:
		tmp := x.GET_Format()
		return CALL_1(tmp, a_0)

	case *PGo:
		return x.Invoke("Format", a_0)
	}
	panic(fmt.Sprintf("Cannot invoke 'Format' with 1 arguments on %v", fn))
}

type i_INVOKE_1_Format interface {
	M_1_Format(a_0 M) M
}

func f_INVOKE_1_split(fn M, a_0 M) M {
	if fn.X == nil {
		if len(fn.S) == 0 {
			panic("cannot INVOKE on int")
		}
		fn = M{X: MkBStr(fn.S).Self}
	}
	switch x := fn.X.(type) {
	case i_INVOKE_1_split:
		return x.M_1_split(a_0)
	case I_GET_split:
		tmp := x.GET_split()
		return CALL_1(tmp, a_0)

	case *PGo:
		return x.Invoke("split", a_0)
	}
	panic(fmt.Sprintf("Cannot invoke 'split' with 1 arguments on %v", fn))
}

type i_INVOKE_1_split interface {
	M_1_split(a_0 M) M
}

func f_INVOKE_2_ServeHTTP(fn M, a_0 M, a_1 M) M {
	if fn.X == nil {
		if len(fn.S) == 0 {
			panic("cannot INVOKE on int")
		}
		fn = M{X: MkBStr(fn.S).Self}
	}
	switch x := fn.X.(type) {
	case i_INVOKE_2_ServeHTTP:
		return x.M_2_ServeHTTP(a_0, a_1)
	case i_GET_ServeHTTP:
		tmp := x.GET_ServeHTTP()
		return CALL_2(tmp, a_0, a_1)

	case *PGo:
		return x.Invoke("ServeHTTP", a_0, a_1)
	}
	panic(fmt.Sprintf("Cannot invoke 'ServeHTTP' with 2 arguments on %v", fn))
}

type i_INVOKE_2_ServeHTTP interface {
	M_2_ServeHTTP(a_0 M, a_1 M) M
}

type i_GET_Format interface {
	GET_Format() M
}

func f_GET_Format(h M) M {
	if h.X == nil {
		panic("cannot GET Field on int or str")
	}
	switch x := h.X.(type) {
	case i_GET_Format:
		return x.GET_Format()
	}
	return h.FetchField("Format")
}

type i_GET_RequestURI interface {
	GET_RequestURI() M
}

func f_GET_RequestURI(h M) M {
	if h.X == nil {
		panic("cannot GET Field on int or str")
	}
	switch x := h.X.(type) {
	case i_GET_RequestURI:
		return x.GET_RequestURI()
	}
	return h.FetchField("RequestURI")
}

type i_GET_ServeHTTP interface {
	GET_ServeHTTP() M
}

func f_GET_ServeHTTP(h M) M {
	if h.X == nil {
		panic("cannot GET Field on int or str")
	}
	switch x := h.X.(type) {
	case i_GET_ServeHTTP:
		return x.GET_ServeHTTP()
	}
	return h.FetchField("ServeHTTP")
}

// self.signatures.items: {}

//ydefs// main => Yfunc:(jim.main) [[ <codegen.Yfunc object at 0x7fbae195c710> ]]
//
//
//
