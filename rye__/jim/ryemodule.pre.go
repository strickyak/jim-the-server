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
import i_http "net/http" // <parse.Timport object at 0x2b1270cf1310>
// DIR ['__class__', '__delattr__', '__dict__', '__doc__', '__format__', '__getattribute__', '__hash__', '__init__', '__module__', '__new__', '__reduce__', '__reduce_ex__', '__repr__', '__setattr__', '__sizeof__', '__str__', '__subclasshook__', '__weakref__', 'alias', 'fromWhere', 'gloss', 'imported', 'line', 'pkg', 'visit', 'where'] // VARS {'imported': ['go', 'net', 'http'], 'gloss': 'from', 'fromWhere': 'go', 'alias': 'http', 'pkg': 'net/http', 'line': 21, 'where': 730}
import i_os "os" // <parse.Timport object at 0x2b1270cf1450>
// DIR ['__class__', '__delattr__', '__dict__', '__doc__', '__format__', '__getattribute__', '__hash__', '__init__', '__module__', '__new__', '__reduce__', '__reduce_ex__', '__repr__', '__setattr__', '__sizeof__', '__str__', '__subclasshook__', '__weakref__', 'alias', 'fromWhere', 'gloss', 'imported', 'line', 'pkg', 'visit', 'where'] // VARS {'imported': ['go', 'os'], 'gloss': 'from', 'fromWhere': 'go', 'alias': 'os', 'pkg': 'os', 'line': 21, 'where': 730}
import i_filepath "path/filepath" // <parse.Timport object at 0x2b1270cf1490>
// DIR ['__class__', '__delattr__', '__dict__', '__doc__', '__format__', '__getattribute__', '__hash__', '__init__', '__module__', '__new__', '__reduce__', '__reduce_ex__', '__repr__', '__setattr__', '__sizeof__', '__str__', '__subclasshook__', '__weakref__', 'alias', 'fromWhere', 'gloss', 'imported', 'line', 'pkg', 'visit', 'where'] // VARS {'imported': ['go', 'path', 'filepath'], 'gloss': 'from', 'fromWhere': 'go', 'alias': 'filepath', 'pkg': 'path/filepath', 'line': 21, 'where': 730}
import i_syscall "syscall" // <parse.Timport object at 0x2b1270cf14d0>
// DIR ['__class__', '__delattr__', '__dict__', '__doc__', '__format__', '__getattribute__', '__hash__', '__init__', '__module__', '__new__', '__reduce__', '__reduce_ex__', '__repr__', '__setattr__', '__sizeof__', '__str__', '__subclasshook__', '__weakref__', 'alias', 'fromWhere', 'gloss', 'imported', 'line', 'pkg', 'visit', 'where'] // VARS {'imported': ['go', 'syscall'], 'gloss': 'from', 'fromWhere': 'go', 'alias': 'syscall', 'pkg': 'syscall', 'line': 21, 'where': 730}
import i_time "time" // <parse.Timport object at 0x2b1270cf1510>
// DIR ['__class__', '__delattr__', '__dict__', '__doc__', '__format__', '__getattribute__', '__hash__', '__init__', '__module__', '__new__', '__reduce__', '__reduce_ex__', '__repr__', '__setattr__', '__sizeof__', '__str__', '__subclasshook__', '__weakref__', 'alias', 'fromWhere', 'gloss', 'imported', 'line', 'pkg', 'visit', 'where'] // VARS {'imported': ['go', 'time'], 'gloss': 'from', 'fromWhere': 'go', 'alias': 'time', 'pkg': 'time', 'line': 21, 'where': 730}
var _ = i_filepath.SkipDir          // path/filepath
var _ = i_http.ErrWriteAfterFlush   // net/http
var _ = i_os.Stdout                 // os
var _ = i_syscall.SocketDisableIPv6 // syscall
var _ = i_time.UTC                  // time
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
	//Vimport: ['go', 'path', 'filepath'] filepath go
	//Vimport: ['go', 'path', 'filepath'] filepath go
	// $ 730 $ 21 $
	// @ 730 @ 21 @
	//Vimport: ['go', 'syscall'] syscall go
	//Vimport: ['go', 'syscall'] syscall go
	// $ 730 $ 21 $
	// @ 730 @ 21 @
	//Vimport: ['go', 'time'] time go
	//Vimport: ['go', 'time'] time go
	// $ 730 $ 21 $
	// @ 838 @ 24 @
	// $ 838 $ 24 $
	// @ 1221 @ 34 @
	// $ 1221 $ 34 $
	return None
}

//(begin tail)
// zip(p.argsPlus, typPlus): [('p', None)]
// typPlus: [None]
///////////////////////////////

func G_1_CheckPublicPermissions(a_p M) M {
	var v_d M = None
	_ = v_d
	// @ 871 @ 25 @ CheckPublicPermissions
	// Vcall: fn: <parse.Tfield object at 0x2b1270cf15d0>
	// Vcall: args: [<parse.Tvar object at 0x2b1270cf1610>]
	// Vcall: names: ['']
	// Vcall: star: None
	// Vcall: starstar: None
	// Vvar: local var p -> 'a_p'
	a_p = MkGo(i_filepath.Clean).Call( /*Yvar.str*/ a_p) // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
	// $ 871 $ 25 $
	// @ 895 @ 26 @ CheckPublicPermissions
	// Vcall: fn: <parse.Tfield object at 0x2b1270cf17d0>
	// Vcall: args: []
	// Vcall: names: []
	// Vcall: star: None
	// Vcall: starstar: None
	// Vcall: fn: <parse.Tfield object at 0x2b1270cf1710>
	// Vcall: args: [<parse.Tvar object at 0x2b1270cf1750>]
	// Vcall: names: ['']
	// Vcall: star: None
	// Vcall: starstar: None
	// Vvar: local var p -> 'a_p'
	if /*DoNE*/ /*General*/ /*invoker*/ f_INVOKE_0_Mode(MkGo(i_os.Stat).Call( /*Yvar.str*/ a_p)).BitAnd( /*Yint.str*/ litI_95b09698fda1f64af16708ffb859eab9).NE( /*Yint.str*/ litI_95b09698fda1f64af16708ffb859eab9) {
		// @ 938 @ 27 @ CheckPublicPermissions
		// Vcall: fn: <parse.Tfield object at 0x2b1270cf1a90>
		// Vcall: args: []
		// Vcall: names: []
		// Vcall: star: None
		// Vcall: starstar: None
		// Vcall: fn: <parse.Tfield object at 0x2b1270cf19d0>
		// Vcall: args: [<parse.Tvar object at 0x2b1270cf1a10>]
		// Vcall: names: ['']
		// Vcall: star: None
		// Vcall: starstar: None
		// Vvar: local var p -> 'a_p'
		// Vvar: local var p -> 'a_p'
		panic(M(( /*DoMod*/ /*Ystr.str*/ litS_1e41c88a380b6897e691627e99ec07fc.Mod(MkTupleV( /*General*/ /*invoker*/ f_INVOKE_0_Mode(MkGo(i_os.Stat).Call( /*Yvar.str*/ a_p)) /*Yvar.str*/, a_p)))))
		// $ 938 $ 27 $
	}
	// $ 895 $ 26 $
	// @ 1006 @ 28 @ CheckPublicPermissions
	// Vcall: fn: <parse.Tfield object at 0x2b1270cf1d10>
	// Vcall: args: [<parse.Tvar object at 0x2b1270cf1d50>]
	// Vcall: names: ['']
	// Vcall: star: None
	// Vcall: starstar: None
	// Vvar: local var p -> 'a_p'
	// @@@@@@ Creating var "d" in scope @@@@@@
	v_d = MkGo(i_filepath.Dir).Call( /*Yvar.str*/ a_p) // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
	// $ 1006 $ 28 $
	// @ 1028 @ 29 @ CheckPublicPermissions
	for {
		// Vvar: local var d -> 'v_d'
		// Ybool::__str__  '(/*DoNE*//*Yvar.str*/v_d.NE(/*Ystr.str*/litS_5058f1af8388633f609cadb75a75dc9d))' None
		var andand_102 M = /*Ybool.str*/ MkBool(( /*DoNE*/ /*Yvar.str*/ v_d.NE( /*Ystr.str*/ litS_5058f1af8388633f609cadb75a75dc9d)))
		if andand_102.Bool() {
			// Vvar: local var d -> 'v_d'
			// Ybool::__str__  '(/*DoNE*//*Yvar.str*/v_d.NE(/*Ystr.str*/litS_6666cd76f96956469e7be39d750cc7d9))' None
			andand_102 = /*Ybool.str*/ MkBool(( /*DoNE*/ /*Yvar.str*/ v_d.NE( /*Ystr.str*/ litS_6666cd76f96956469e7be39d750cc7d9)))
		}
		var andand_101 M = andand_102
		if andand_101.Bool() {
			// Vvar: local var d -> 'v_d'
			// Ybool::__str__  '(/*DoNE*//*Yvar.str*/v_d.NE(/*Ystr.str*/litS_d41d8cd98f00b204e9800998ecf8427e))' None
			andand_101 = /*Ybool.str*/ MkBool(( /*DoNE*/ /*Yvar.str*/ v_d.NE( /*Ystr.str*/ litS_d41d8cd98f00b204e9800998ecf8427e)))
		}
		if !( /*AsBool*/ andand_101.Bool()) {
			break
		}
		// @ 1073 @ 30 @ CheckPublicPermissions
		// Vcall: fn: <parse.Tfield object at 0x2b1270d0d210>
		// Vcall: args: []
		// Vcall: names: []
		// Vcall: star: None
		// Vcall: starstar: None
		// Vcall: fn: <parse.Tfield object at 0x2b1270d0d150>
		// Vcall: args: [<parse.Tvar object at 0x2b1270d0d190>]
		// Vcall: names: ['']
		// Vcall: star: None
		// Vcall: starstar: None
		// Vvar: local var d -> 'v_d'
		if /*DoNE*/ /*General*/ /*invoker*/ f_INVOKE_0_Mode(MkGo(i_os.Stat).Call( /*Yvar.str*/ v_d)).BitAnd( /*Yint.str*/ litI_d39934ce111a864abf40391f3da9cdf5).NE( /*Yint.str*/ litI_d39934ce111a864abf40391f3da9cdf5) {
			// @ 1118 @ 31 @ CheckPublicPermissions
			// Vcall: fn: <parse.Tfield object at 0x2b1270d0d4d0>
			// Vcall: args: []
			// Vcall: names: []
			// Vcall: star: None
			// Vcall: starstar: None
			// Vcall: fn: <parse.Tfield object at 0x2b1270d0d410>
			// Vcall: args: [<parse.Tvar object at 0x2b1270d0d450>]
			// Vcall: names: ['']
			// Vcall: star: None
			// Vcall: starstar: None
			// Vvar: local var d -> 'v_d'
			// Vvar: local var d -> 'v_d'
			panic(M(( /*DoMod*/ /*Ystr.str*/ litS_01a99182d678d59178e9adbb72ff87ae.Mod(MkTupleV( /*General*/ /*invoker*/ f_INVOKE_0_Mode(MkGo(i_os.Stat).Call( /*Yvar.str*/ v_d)) /*Yvar.str*/, v_d)))))
			// $ 1118 $ 31 $
		}
		// $ 1073 $ 30 $
		// @ 1200 @ 32 @ CheckPublicPermissions
		// Vcall: fn: <parse.Tfield object at 0x2b1270d0d750>
		// Vcall: args: [<parse.Tvar object at 0x2b1270d0d790>]
		// Vcall: names: ['']
		// Vcall: star: None
		// Vcall: starstar: None
		// Vvar: local var d -> 'v_d'
		v_d = MkGo(i_filepath.Dir).Call( /*Yvar.str*/ v_d) // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
		// $ 1200 $ 32 $
	}
	// $ 1028 $ 29 $

	return None
}

///////////////////////////////
// name: CheckPublicPermissions

var specFunc_CheckPublicPermissions = CallSpec{Name: "CheckPublicPermissions", Args: []string{"p"}, Defaults: []M{MissingM}, Star: "", StarStar: ""}

type pFunc_CheckPublicPermissions struct{ PNewCallable }

func (o *pFunc_CheckPublicPermissions) Contents() interface{} {
	return G_CheckPublicPermissions
}
func (o pFunc_CheckPublicPermissions) Call1(a0 M) M {
	return G_1_CheckPublicPermissions(a0)
}

func (o pFunc_CheckPublicPermissions) CallV(a1 []M, a2 []M, kv1 []KV, kv2 map[string]M) M {
	argv, star, starstar := NewSpecCall(o.CallSpec, a1, a2, kv1, kv2)
	_, _, _ = argv, star, starstar
	return G_1_CheckPublicPermissions(argv[0])
}

//(tail)
///////////////////////////////
// name: wrapper
var specNest_nesting_107 = CallSpec{Name: "wrapper__nesting_107", Args: []string{"w", "r"}, Defaults: []M{MissingM, MissingM}, Star: "", StarStar: ""}

type pNest_nesting_107 struct {
	PNewCallable
	fn func(a_w M, a_r M) M
}

func (o *pNest_nesting_107) Contents() interface{} {
	return o.fn
}
func (o pNest_nesting_107) Call2(a0 M, a1 M) M {
	return o.fn(a0, a1)
}

func (o pNest_nesting_107) CallV(a1 []M, a2 []M, kv1 []KV, kv2 map[string]M) M {
	argv, star, starstar := NewSpecCall(o.CallSpec, a1, a2, kv1, kv2)
	_, _, _ = argv, star, starstar
	return o.fn(argv[0], argv[1])
}

//(tail)
///////////////////////////////
// name: makeWrapper
var specNest_nesting_106 = CallSpec{Name: "makeWrapper__nesting_106", Args: []string{"pattern", "webDir", "handler"}, Defaults: []M{MissingM, MissingM, MissingM}, Star: "", StarStar: ""}

type pNest_nesting_106 struct {
	PNewCallable
	fn func(a_pattern M, a_webDir M, a_handler M) M
}

func (o *pNest_nesting_106) Contents() interface{} {
	return o.fn
}
func (o pNest_nesting_106) Call3(a0 M, a1 M, a2 M) M {
	return o.fn(a0, a1, a2)
}

func (o pNest_nesting_106) CallV(a1 []M, a2 []M, kv1 []KV, kv2 map[string]M) M {
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
	var v_detuple_103 M = None
	_ = v_detuple_103
	var v_handler M = None
	_ = v_handler
	var v_line M = None
	_ = v_line
	var v_makeWrapper M = None
	_ = v_makeWrapper
	var v_pattern M = None
	_ = v_pattern
	var v_webDir M = None
	_ = v_webDir
	var v_words M = None
	_ = v_words
	// @ 1255 @ 36 @ main
	// Vvar: local var args -> 'a_args'
	// @@@@@@ Creating var "detuple_103" in scope @@@@@@
	v_detuple_103 = /*Yvar.str*/ a_args // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <class 'codegen.Yvar'>
	// Vvar: local var detuple_103 -> 'v_detuple_103'
	len_detuple_103 := /*Yvar.str*/ v_detuple_103.Len()
	if len_detuple_103 != 2 {
		panic(fmt.Sprintf("Assigning object of length %d to %d variables, in destructuring assignment.", len_detuple_103, 2))
	}
	// Vvar: local var detuple_103 -> 'v_detuple_103'
	// @@@@@@ Creating var "bindHostColonPort" in scope @@@@@@
	v_bindHostColonPort = /*Yvar.str*/ v_detuple_103.GetItem( /*Yint.str*/ litI_cfcd208495d565ef66e7dff9f98764da) // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
	// Vvar: local var detuple_103 -> 'v_detuple_103'
	// @@@@@@ Creating var "confFile" in scope @@@@@@
	v_confFile = /*Yvar.str*/ v_detuple_103.GetItem( /*Yint.str*/ litI_c4ca4238a0b923820dcc509a6f75849b) // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
	// $ 1255 $ 36 $
	// @ 1292 @ 37 @ main
	fmt.Fprintln(M(MkGo( /*Yimport.str*/ i_os.Stderr)).Contents().(io.Writer), "# Jim the Server")
	// $ 1292 $ 37 $
	// @ 1333 @ 39 @ main
	// Vcall: fn: <parse.Tfield object at 0x2b1270d0dd50>
	// Vcall: args: [<parse.Tlit object at 0x2b1270d0dd90>]
	// Vcall: names: ['']
	// Vcall: star: None
	// Vcall: starstar: None
	// Vcall: fn: <parse.Tfield object at 0x2b1270d0dcd0>
	// Vcall: args: []
	// Vcall: names: []
	// Vcall: star: None
	// Vcall: starstar: None
	// Vcall: fn: <parse.Tvar object at 0x2b1270d0dc10>
	// Vcall: args: [<parse.Tvar object at 0x2b1270d0dc50>]
	// Vcall: names: ['']
	// Vcall: star: None
	// Vcall: starstar: None
	// Making Global Yvar from 'open'
	// Vvar: local var confFile -> 'v_confFile'

	for_returning__104 := func() M { // around FOR
		var nexter__104 Nexter = /*General*/ /*invoker*/ f_INVOKE_1_split( /*General*/ /*invoker*/ F_INVOKE_0_read(CALL_1(M( /*Yvar.str*/ G_open) /*Yvar.str*/, v_confFile)) /*Ystr.str*/, litS_68b329da9893e34099c7d8ad5cb9c940).Iter()
		enougher__104, canEnough__104 := nexter__104.(Enougher)
		if canEnough__104 {
			defer enougher__104.Enough()
		}
		// else case without Enougher will be faster.
		for {
			ndx___104, more___104 := nexter__104.Next()
			if !more___104 {
				break
			}
			// BEGIN FOR

			// @@@@@@ Creating var "line" in scope @@@@@@
			v_line = ndx___104 // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
			// @ 1384 @ 40 @ main
			// Vcall: fn: <parse.Tfield object at 0x2b1270d0de90>
			// Vcall: args: [<parse.Tlit object at 0x2b1270d0ded0>]
			// Vcall: names: ['']
			// Vcall: star: None
			// Vcall: starstar: None
			// Vvar: local var line -> 'v_line'
			// @@@@@@ Creating var "command" in scope @@@@@@
			v_command = /*General*/ /*invoker*/ f_INVOKE_1_split( /*Yvar.str*/ v_line /*Ystr.str*/, litS_01abfc750a0c942167651c40d088531d).GetItem( /*Yint.str*/ litI_cfcd208495d565ef66e7dff9f98764da) // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
			// $ 1384 $ 40 $
			// @ 1447 @ 41 @ main
			// Vcall: fn: <parse.Tfield object at 0x2b1270d130d0>
			// Vcall: args: []
			// Vcall: names: []
			// Vcall: star: None
			// Vcall: starstar: None
			// Vvar: local var command -> 'v_command'
			// @@@@@@ Creating var "words" in scope @@@@@@
			v_words = /*General*/ /*invoker*/ f_INVOKE_0_split( /*Yvar.str*/ v_command) // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
			// $ 1447 $ 41 $
			// @ 1476 @ 43 @ main
			// Vvar: local var words -> 'v_words'
			if /*AsBool*/ /*Yvar.str*/ v_words.Bool() {
				// @ 1492 @ 44 @ main
				// Vvar: local var words -> 'v_words'
				// @@@@@@ Creating var "webDir" in scope @@@@@@
				v_webDir = /*Yvar.str*/ v_words.GetItem( /*Yint.str*/ litI_cfcd208495d565ef66e7dff9f98764da) // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
				// $ 1492 $ 44 $
				// @ 1516 @ 45 @ main
				// Vvar: local var words -> 'v_words'

				for_returning__105 := func() M { // around FOR
					var nexter__105 Nexter = /*Yvar.str*/ v_words.GetItemSlice( /*Yint.str*/ litI_c4ca4238a0b923820dcc509a6f75849b, None, None).Iter()
					enougher__105, canEnough__105 := nexter__105.(Enougher)
					if canEnough__105 {
						defer enougher__105.Enough()
					}
					// else case without Enougher will be faster.
					for {
						ndx___105, more___105 := nexter__105.Next()
						if !more___105 {
							break
						}
						// BEGIN FOR

						// @@@@@@ Creating var "pattern" in scope @@@@@@
						v_pattern = ndx___105 // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
						// @ 1550 @ 46 @ main
						// @@@@@@ Creating var "makeWrapper" in scope @@@@@@
						v_makeWrapper = None // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
						// zip(p.argsPlus, typPlus): [('pattern', None), ('webDir', None), ('handler', None)]
						// typPlus: [None, None, None]
						///////////////////////////////

						fn_nesting_106 := func(a_pattern M, a_webDir M, a_handler M) M {
							var v_wrapper M = None
							_ = v_wrapper
							// @ 1603 @ 47 @ makeWrapper
							// @@@@@@ Creating var "wrapper" in scope @@@@@@
							v_wrapper = None // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
							// zip(p.argsPlus, typPlus): [('w', None), ('r', None)]
							// typPlus: [None, None]
							///////////////////////////////

							fn_nesting_107 := func(a_w M, a_r M) M {
								var v_e M = None
								_ = v_e
								// @ 1634 @ 48 @ wrapper
								// Vcall: fn: <parse.Tfield object at 0x2b1270d13590>
								// Vcall: args: [<parse.Tfield object at 0x2b1270d13610>]
								// Vcall: names: ['']
								// Vcall: star: None
								// Vcall: starstar: None
								// Vcall: fn: <parse.Tfield object at 0x2b1270d13510>
								// Vcall: args: []
								// Vcall: names: []
								// Vcall: star: None
								// Vcall: starstar: None
								// Vvar: local var webDir -> 'a_webDir'
								// Vvar: local var pattern -> 'a_pattern'
								// Vvar: local var r -> 'a_r'
								fmt.Fprintln(M(MkGo( /*Yimport.str*/ i_os.Stderr)).Contents().(io.Writer) /*ForceString*/, ( /*DoMod*/ /*Ystr.str*/ litS_3b6431eaaa10ce8061f0a1a462a6b4d3.Mod(MkTupleV( /*General*/ /*invoker*/ f_INVOKE_1_Format(MkGo(i_time.Now).Call(), MkGo( /*Yimport.str*/ i_time.Stamp)) /*Yvar.str*/, a_webDir /*Yvar.str*/, a_pattern, f_GET_RequestURI( /*Yvar.str*/ a_r)))).String())
								// $ 1634 $ 48 $
								// @ 1766 @ 50 @ wrapper

								// BEGIN OUTER EXCEPT try_108
								try_108_try := func() (try_108_z M) {
									defer func() {
										r := recover()
										if r != nil {
											PrintStackFYIUnlessEOFBecauseExcept(r)
											try_108_z = func() M {
												// BEGIN EXCEPT

												// @@@@@@ Creating var "e" in scope @@@@@@
												v_e = MkRecovered(r) // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
												// @ 1882 @ 53 @ wrapper
												// Vcall: fn: <parse.Tfield object at 0x2b1270d13c10>
												// Vcall: args: [<parse.Tvar object at 0x2b1270d13c50>, <parse.Tvar object at 0x2b1270d13c90>]
												// Vcall: names: ['', '']
												// Vcall: star: None
												// Vcall: starstar: None
												// Vvar: local var w -> 'a_w'
												// Vvar: local var r -> 'a_r'
												_ = MkGo(i_http.NotFound).Call( /*Yvar.str*/ a_w /*Yvar.str*/, a_r) // Assign void: = type: <type 'str'> repr: 'MkGo(i_http.NotFound).Call(/*Yvar.str*/a_w, /*Yvar.str*/a_r) '
												// $ 1882 $ 53 $
												// @ 1916 @ 54 @ wrapper
												// Vvar: local var e -> 'v_e'
												// Vvar: local var webDir -> 'a_webDir'
												// Vvar: local var pattern -> 'a_pattern'
												// Vvar: local var r -> 'a_r'
												fmt.Fprintln(M(MkGo( /*Yimport.str*/ i_os.Stderr)).Contents().(io.Writer) /*ForceString*/, ( /*DoMod*/ /*Ystr.str*/ litS_7e8a11cf9c12e875994ef4481ac49e3d.Mod(MkTupleV( /*Yvar.str*/ v_e /*Yvar.str*/, a_webDir /*Yvar.str*/, a_pattern, f_GET_RequestURI( /*Yvar.str*/ a_r)))).String())
												// $ 1916 $ 54 $
												// @ 2012 @ 55 @ wrapper
												return None
												// $ 2012 $ 55 $

												return MissingM
												// END EXCEPT
											}()
											return
										}
									}()

									// BEGIN TRY try_108
									// @ 1785 @ 51 @ wrapper
									// Vcall: fn: <parse.Tvar object at 0x2b1270d13890>
									// Vcall: args: [<parse.Tcall object at 0x2b1270d13a50>]
									// Vcall: names: ['']
									// Vcall: star: None
									// Vcall: starstar: None
									// Making Global Yvar from 'CheckPublicPermissions'
									// Vcall: fn: <parse.Tfield object at 0x2b1270d13910>
									// Vcall: args: [<parse.Tvar object at 0x2b1270d13950>, <parse.Tfield object at 0x2b1270d13a10>]
									// Vcall: names: ['', '']
									// Vcall: star: None
									// Vcall: starstar: None
									// Vvar: local var webDir -> 'a_webDir'
									// Vvar: local var r -> 'a_r'
									_ = G_1_CheckPublicPermissions(MkGo(i_filepath.Join).Call( /*Yvar.str*/ a_webDir, f_GET_Path(f_GET_URL( /*Yvar.str*/ a_r)))) // Assign void: = type: <type 'str'> repr: 'G_1_CheckPublicPermissions(MkGo(i_filepath.Join).Call(/*Yvar.str*/a_webDir,  f_GET_Path( f_GET_URL(/*Yvar.str*/a_r) ) ) ) '
									// $ 1785 $ 51 $
									// END TRY try_108

									return MissingM
								}()
								if try_108_try != MissingM {
									return try_108_try
								}
								// END OUTER EXCEPT try_108

								// $ 1766 $ 50 $
								// @ 2031 @ 56 @ wrapper
								// Vcall: fn: <parse.Tfield object at 0x2b1270d1b1d0>
								// Vcall: args: [<parse.Tvar object at 0x2b1270d1b210>, <parse.Tvar object at 0x2b1270d1b250>]
								// Vcall: names: ['', '']
								// Vcall: star: None
								// Vcall: starstar: None
								// Vvar: local var handler -> 'a_handler'
								// Vvar: local var w -> 'a_w'
								// Vvar: local var r -> 'a_r'
								_ = /*General*/ /*invoker*/ f_INVOKE_2_ServeHTTP( /*Yvar.str*/ a_handler /*Yvar.str*/, a_w /*Yvar.str*/, a_r) // Assign void: = type: <type 'str'> repr: '/*General*/ /*invoker*/ f_INVOKE_2_ServeHTTP(/*Yvar.str*/a_handler, /*Yvar.str*/a_w, /*Yvar.str*/a_r) '
								// $ 2031 $ 56 $

								return None
							}

							v_wrapper = MForge(&pNest_nesting_107{PNewCallable{CallSpec: &specNest_nesting_107}, fn_nesting_107}) // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
							// $ 1603 $ 47 $
							// @ 2065 @ 57 @ makeWrapper
							// Vvar: local var wrapper -> 'v_wrapper'
							return /*Yvar.str*/ v_wrapper
							// $ 2065 $ 57 $

							return None
						}

						v_makeWrapper = MForge(&pNest_nesting_106{PNewCallable{CallSpec: &specNest_nesting_106}, fn_nesting_106}) // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
						// $ 1550 $ 46 $
						// @ 2089 @ 59 @ main
						// Vcall: fn: <parse.Tfield object at 0x2b1270d1b550>
						// Vcall: args: [<parse.Tcall object at 0x2b1270d1b690>]
						// Vcall: names: ['']
						// Vcall: star: None
						// Vcall: starstar: None
						// Vcall: fn: <parse.Tvar object at 0x2b1270d1b590>
						// Vcall: args: [<parse.Tfield object at 0x2b1270d1b610>, <parse.Tvar object at 0x2b1270d1b650>]
						// Vcall: names: ['', '']
						// Vcall: star: None
						// Vcall: starstar: None
						// Vvar: local var webDir -> 'v_webDir'
						// @@@@@@ Creating var "handler" in scope @@@@@@
						v_handler = MkGo(i_http.FileServer).Call(GoCast(GoElemType(new( /*Yimport.str*/ i_http.Dir)) /*Yvar.str*/, v_webDir)) // Assign <class 'parse.Tvar'> lhs <type 'str'> = rhs <type 'str'>
						// $ 2089 $ 59 $
						// @ 2150 @ 60 @ main
						// Vcall: fn: <parse.Tfield object at 0x2b1270d1b790>
						// Vcall: args: [<parse.Tvar object at 0x2b1270d1b7d0>, <parse.Tcall object at 0x2b1270d1b910>]
						// Vcall: names: ['', '']
						// Vcall: star: None
						// Vcall: starstar: None
						// Vvar: local var pattern -> 'v_pattern'
						// Vcall: fn: <parse.Tvar object at 0x2b1270d1b810>
						// Vcall: args: [<parse.Tvar object at 0x2b1270d1b850>, <parse.Tvar object at 0x2b1270d1b890>, <parse.Tvar object at 0x2b1270d1b8d0>]
						// Vcall: names: ['', '', '']
						// Vcall: star: None
						// Vcall: starstar: None
						// Vvar: local var makeWrapper -> 'v_makeWrapper'
						// Vvar: local var pattern -> 'v_pattern'
						// Vvar: local var webDir -> 'v_webDir'
						// Vvar: local var handler -> 'v_handler'
						_ = MkGo(i_http.HandleFunc).Call( /*Yvar.str*/ v_pattern, CALL_3(M( /*Yvar.str*/ v_makeWrapper) /*Yvar.str*/, v_pattern /*Yvar.str*/, v_webDir /*Yvar.str*/, v_handler)) // Assign void: = type: <type 'str'> repr: 'MkGo(i_http.HandleFunc).Call(/*Yvar.str*/v_pattern, CALL_3( M(/*Yvar.str*/v_makeWrapper), /*Yvar.str*/v_pattern, /*Yvar.str*/v_webDir, /*Yvar.str*/v_handler )) '
						// $ 2150 $ 60 $
						// @ 2222 @ 61 @ main
						// Vvar: local var webDir -> 'v_webDir'
						// Vvar: local var pattern -> 'v_pattern'
						fmt.Fprintln(M(MkGo( /*Yimport.str*/ i_os.Stderr)).Contents().(io.Writer) /*ForceString*/, ( /*DoMod*/ /*Ystr.str*/ litS_d7ba898028d447a7ef280df1d6365b74.Mod(MkTupleV( /*Yvar.str*/ v_webDir /*Yvar.str*/, v_pattern))).String())
						// $ 2222 $ 61 $

						// END FOR
					}
					return MissingM
				}() // around FOR
				if for_returning__105 != MissingM {
					return for_returning__105
				}

				// $ 1516 $ 45 $
			}
			// $ 1476 $ 43 $

			// END FOR
		}
		return MissingM
	}() // around FOR
	if for_returning__104 != MissingM {
		return for_returning__104
	}

	// $ 1333 $ 39 $
	// @ 2282 @ 63 @ main
	// Vcall: fn: <parse.Tfield object at 0x2b1270d1bf50>
	// Vcall: args: [<parse.Tfield object at 0x2b1270d1bfd0>]
	// Vcall: names: ['']
	// Vcall: star: None
	// Vcall: starstar: None
	// Vcall: fn: <parse.Tfield object at 0x2b1270d1bed0>
	// Vcall: args: []
	// Vcall: names: []
	// Vcall: star: None
	// Vcall: starstar: None
	fmt.Fprintln(M(MkGo( /*Yimport.str*/ i_os.Stderr)).Contents().(io.Writer) /*ForceString*/, ( /*DoMod*/ /*Ystr.str*/ litS_a7b59768def62f6c9ad89722a73531a1.Mod( /*General*/ /*invoker*/ f_INVOKE_1_Format(MkGo(i_time.Now).Call(), MkGo( /*Yimport.str*/ i_time.Stamp)))).String())
	// $ 2282 $ 63 $
	// @ 2362 @ 65 @ main
	// Vcall: fn: <parse.Tfield object at 0x2b1270d21190>
	// Vcall: args: [<parse.Tvar object at 0x2b1270d211d0>, <parse.Traw object at 0x2b1270d21210>]
	// Vcall: names: ['', '']
	// Vcall: star: None
	// Vcall: starstar: None
	// Vvar: local var bindHostColonPort -> 'v_bindHostColonPort'
	_ = MkGo(i_http.ListenAndServe).Call( /*Yvar.str*/ v_bindHostColonPort, None) // Assign void: = type: <type 'str'> repr: 'MkGo(i_http.ListenAndServe).Call(/*Yvar.str*/v_bindHostColonPort, None) '
	// $ 2362 $ 65 $

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

var G_CheckPublicPermissions M // *pFunc_CheckPublicPermissions
var G___name__ M               // M
var G_filepath M               // *PModule
var G_http M                   // *PModule
var G_main M                   // *pFunc_main
var G_os M                     // *PModule
var G_syscall M                // *PModule
var G_time M                   // *PModule

func init /*New_Module*/ () {
	G_CheckPublicPermissions = MForge(&pFunc_CheckPublicPermissions{PNewCallable{CallSpec: &specFunc_CheckPublicPermissions}})
	G___name__ = MkStr("jim")
	G_filepath = None
	G_http = None
	G_main = MForge(&pFunc_main{PNewCallable{CallSpec: &specFunc_main}})
	G_os = None
	G_syscall = None
	G_time = None
}

var ModuleMap = map[string]*M{
	"CheckPublicPermissions": &G_CheckPublicPermissions,
	"__name__":               &G___name__,
	"filepath":               &G_filepath,
	"http":                   &G_http,
	"main":                   &G_main,
	"os":                     &G_os,
	"syscall":                &G_syscall,
	"time":                   &G_time,
}

var ModuleObj = MakeModuleObject(ModuleMap, "github.com/strickyak/jim-the-server/jim")

var litI_95b09698fda1f64af16708ffb859eab9 = MkInt(0004)
var litI_c4ca4238a0b923820dcc509a6f75849b = MkInt(1)
var litI_cfcd208495d565ef66e7dff9f98764da = MkInt(0)
var litI_d39934ce111a864abf40391f3da9cdf5 = MkInt(0005)
var litS_01a99182d678d59178e9adbb72ff87ae = MkStr("Dir not public readable & executable: 0%o %q")
var litS_01abfc750a0c942167651c40d088531d = MkStr("#")
var litS_1e41c88a380b6897e691627e99ec07fc = MkStr("File not public readable: 0%o %q")
var litS_3b6431eaaa10ce8061f0a1a462a6b4d3 = MkStr("%s (%s == %s) %q")
var litS_5058f1af8388633f609cadb75a75dc9d = MkStr(".")
var litS_6666cd76f96956469e7be39d750cc7d9 = MkStr("/")
var litS_68b329da9893e34099c7d8ad5cb9c940 = MkStr("\x0a")
var litS_7e8a11cf9c12e875994ef4481ac49e3d = MkStr("ERROR %q (%s == %s) %q")
var litS_9d222d8f9b2a29c89c3ecedd1f5470cc = MkStr("# Jim the Server")
var litS_a7b59768def62f6c9ad89722a73531a1 = MkStr("# Starting Server at %s")
var litS_d41d8cd98f00b204e9800998ecf8427e = MkStr("")
var litS_d7ba898028d447a7ef280df1d6365b74 = MkStr("# %q handles %q")

func f_INVOKE_0_Mode(fn M) M {
	if fn.X == nil {
		if len(fn.S) == 0 {
			panic("cannot INVOKE on int")
		}
		fn = M{X: MkBStr(fn.S).Self}
	}
	switch x := fn.X.(type) {
	case i_INVOKE_0_Mode:
		return x.M_0_Mode()
	case i_GET_Mode:
		tmp := x.GET_Mode()
		return CALL_0(tmp)

	case *PGo:
		return x.Invoke("Mode")
	}
	panic(fmt.Sprintf("Cannot invoke 'Mode' with 0 arguments on %v", fn))
}

type i_INVOKE_0_Mode interface {
	M_0_Mode() M
}

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

type i_GET_Mode interface {
	GET_Mode() M
}

func f_GET_Mode(h M) M {
	if h.X == nil {
		panic("cannot GET Field on int or str")
	}
	switch x := h.X.(type) {
	case i_GET_Mode:
		return x.GET_Mode()
	}
	return h.FetchField("Mode")
}

type i_GET_Path interface {
	GET_Path() M
}

func f_GET_Path(h M) M {
	if h.X == nil {
		panic("cannot GET Field on int or str")
	}
	switch x := h.X.(type) {
	case i_GET_Path:
		return x.GET_Path()
	}
	return h.FetchField("Path")
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

type i_GET_URL interface {
	GET_URL() M
}

func f_GET_URL(h M) M {
	if h.X == nil {
		panic("cannot GET Field on int or str")
	}
	switch x := h.X.(type) {
	case i_GET_URL:
		return x.GET_URL()
	}
	return h.FetchField("URL")
}

// self.signatures.items: {}

//ydefs// CheckPublicPermissions => Yfunc:(jim.CheckPublicPermissions) [[ <codegen.Yfunc object at 0x2b1270d21590> ]]
//
//ydefs// main => Yfunc:(jim.main) [[ <codegen.Yfunc object at 0x2b1270d215d0> ]]
//
//
//
