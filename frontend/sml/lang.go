//line lang.y:2
package sml

import __yyfmt__ "fmt"

//line lang.y:3
import (
	"github.com/marcopeereboom/gck/ast"
	"math/big"
)

var d *yylexer // being set so we don't have to type assert all the time

//line lang.y:14
type yySymType struct {
	yys        int
	integer    int
	number     *big.Rat
	identifier string
	node       ast.Node
}

const INTEGER = 57346
const IDENTIFIER = 57347
const VAR = 57348
const CONST = 57349
const NUMBER = 57350
const WHILE = 57351
const IF = 57352
const ELSE = 57353
const EOL = 57354
const ASSIGN = 57355
const LE = 57356
const GE = 57357
const NE = 57358
const EQ = 57359
const LT = 57360
const GT = 57361
const UMINUS = 57362

var yyToknames = []string{
	"INTEGER",
	"IDENTIFIER",
	"VAR",
	"CONST",
	"NUMBER",
	"WHILE",
	"IF",
	"ELSE",
	"EOL",
	"ASSIGN",
	"LE",
	"GE",
	"NE",
	"EQ",
	"LT",
	"GT",
	"'+'",
	"'-'",
	"'*'",
	"'/'",
	"UMINUS",
}
var yyStatenames = []string{}

const yyEofCode = 1
const yyErrCode = 2
const yyMaxDepth = 200

//line lang.y:104

//line yacctab:1
var yyExca = []int{
	-1, 1,
	1, -1,
	-2, 0,
}

const yyNprod = 33
const yyPrivate = 57344

var yyTokenNames []string
var yyStates []string

const yyLast = 126

var yyAct = []int{

	9, 8, 5, 57, 17, 42, 43, 44, 45, 40,
	41, 20, 21, 22, 23, 16, 25, 27, 29, 29,
	38, 22, 23, 33, 34, 35, 36, 37, 24, 39,
	59, 17, 48, 47, 10, 12, 2, 1, 11, 15,
	16, 6, 58, 51, 52, 53, 54, 55, 56, 3,
	7, 13, 18, 0, 32, 4, 17, 49, 14, 0,
	60, 61, 10, 12, 28, 0, 11, 15, 16, 20,
	21, 22, 23, 10, 26, 0, 0, 11, 38, 13,
	0, 31, 18, 4, 17, 0, 14, 10, 26, 0,
	13, 11, 0, 0, 0, 46, 0, 14, 20, 21,
	22, 23, 0, 50, 13, 20, 21, 22, 23, 0,
	19, 30, 42, 43, 44, 45, 40, 41, 20, 21,
	22, 23, 20, 21, 22, 23,
}
var yyPact = []int{

	58, -1000, 58, -1000, -1000, 85, -1000, -1000, -1000, -1000,
	-1000, -1000, 15, 69, 69, 83, 83, 58, -1000, -1000,
	69, 69, 69, 69, 69, -1000, -1000, 49, -22, 98,
	83, -22, 30, -1, -1, -1000, -1000, 78, -1000, -1000,
	69, 69, 69, 69, 69, 69, -26, -9, 19, -1000,
	-1000, 102, 102, 102, 102, 102, 102, -1000, -1000, 5,
	-1000, -1000,
}
var yyPgo = []int{

	0, 49, 36, 2, 64, 50, 1, 42, 0, 41,
	37,
}
var yyR1 = []int{

	0, 10, 1, 1, 1, 1, 1, 1, 2, 2,
	8, 9, 5, 6, 7, 7, 7, 4, 4, 4,
	4, 4, 4, 4, 3, 3, 3, 3, 3, 3,
	3, 3, 3,
}
var yyR2 = []int{

	0, 1, 1, 2, 1, 1, 1, 1, 1, 2,
	3, 4, 3, 4, 0, 2, 2, 3, 3, 3,
	3, 3, 3, 3, 1, 1, 1, 2, 3, 3,
	3, 3, 3,
}
var yyChk = []int{

	-1000, -10, -2, -1, 25, -3, -9, -5, -6, -8,
	4, 8, 5, 21, 28, 9, 10, 26, -1, 25,
	20, 21, 22, 23, 13, -3, 5, -3, -4, -3,
	28, -4, -2, -3, -3, -3, -3, -3, 29, -8,
	18, 19, 14, 15, 16, 17, -4, -3, -8, 27,
	25, -3, -3, -3, -3, -3, -3, 29, -7, 11,
	-8, -6,
}
var yyDef = []int{

	0, -2, 1, 8, 2, 0, 4, 5, 6, 7,
	24, 25, 26, 0, 0, 0, 0, 0, 9, 3,
	0, 0, 0, 0, 0, 27, 26, 0, 0, 0,
	0, 0, 0, 28, 29, 30, 31, 0, 32, 12,
	0, 0, 0, 0, 0, 0, 0, 0, 14, 10,
	11, 17, 18, 19, 20, 21, 22, 23, 13, 0,
	15, 16,
}
var yyTok1 = []int{

	1, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	28, 29, 22, 20, 3, 21, 3, 23, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 25,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 3, 3, 3, 3, 3, 3, 3,
	3, 3, 3, 26, 3, 27,
}
var yyTok2 = []int{

	2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	12, 13, 14, 15, 16, 17, 18, 19, 24,
}
var yyTok3 = []int{
	0,
}

//line yaccpar:1

/*	parser for yacc output	*/

var yyDebug = 0

type yyLexer interface {
	Lex(lval *yySymType) int
	Error(s string)
}

const yyFlag = -1000

func yyTokname(c int) string {
	// 4 is TOKSTART above
	if c >= 4 && c-4 < len(yyToknames) {
		if yyToknames[c-4] != "" {
			return yyToknames[c-4]
		}
	}
	return __yyfmt__.Sprintf("tok-%v", c)
}

func yyStatname(s int) string {
	if s >= 0 && s < len(yyStatenames) {
		if yyStatenames[s] != "" {
			return yyStatenames[s]
		}
	}
	return __yyfmt__.Sprintf("state-%v", s)
}

func yylex1(lex yyLexer, lval *yySymType) int {
	c := 0
	char := lex.Lex(lval)
	if char <= 0 {
		c = yyTok1[0]
		goto out
	}
	if char < len(yyTok1) {
		c = yyTok1[char]
		goto out
	}
	if char >= yyPrivate {
		if char < yyPrivate+len(yyTok2) {
			c = yyTok2[char-yyPrivate]
			goto out
		}
	}
	for i := 0; i < len(yyTok3); i += 2 {
		c = yyTok3[i+0]
		if c == char {
			c = yyTok3[i+1]
			goto out
		}
	}

out:
	if c == 0 {
		c = yyTok2[1] /* unknown char */
	}
	if yyDebug >= 3 {
		__yyfmt__.Printf("lex %s(%d)\n", yyTokname(c), uint(char))
	}
	return c
}

func yyParse(yylex yyLexer) int {
	var yyn int
	var yylval yySymType
	var yyVAL yySymType
	yyS := make([]yySymType, yyMaxDepth)

	Nerrs := 0   /* number of errors */
	Errflag := 0 /* error recovery flag */
	yystate := 0
	yychar := -1
	yyp := -1
	goto yystack

ret0:
	return 0

ret1:
	return 1

yystack:
	/* put a state and value onto the stack */
	if yyDebug >= 4 {
		__yyfmt__.Printf("char %v in %v\n", yyTokname(yychar), yyStatname(yystate))
	}

	yyp++
	if yyp >= len(yyS) {
		nyys := make([]yySymType, len(yyS)*2)
		copy(nyys, yyS)
		yyS = nyys
	}
	yyS[yyp] = yyVAL
	yyS[yyp].yys = yystate

yynewstate:
	yyn = yyPact[yystate]
	if yyn <= yyFlag {
		goto yydefault /* simple state */
	}
	if yychar < 0 {
		yychar = yylex1(yylex, &yylval)
	}
	yyn += yychar
	if yyn < 0 || yyn >= yyLast {
		goto yydefault
	}
	yyn = yyAct[yyn]
	if yyChk[yyn] == yychar { /* valid shift */
		yychar = -1
		yyVAL = yylval
		yystate = yyn
		if Errflag > 0 {
			Errflag--
		}
		goto yystack
	}

yydefault:
	/* default state action */
	yyn = yyDef[yystate]
	if yyn == -2 {
		if yychar < 0 {
			yychar = yylex1(yylex, &yylval)
		}

		/* look through exception table */
		xi := 0
		for {
			if yyExca[xi+0] == -1 && yyExca[xi+1] == yystate {
				break
			}
			xi += 2
		}
		for xi += 2; ; xi += 2 {
			yyn = yyExca[xi+0]
			if yyn < 0 || yyn == yychar {
				break
			}
		}
		yyn = yyExca[xi+1]
		if yyn < 0 {
			goto ret0
		}
	}
	if yyn == 0 {
		/* error ... attempt to resume parsing */
		switch Errflag {
		case 0: /* brand new error */
			yylex.Error("syntax error")
			Nerrs++
			if yyDebug >= 1 {
				__yyfmt__.Printf("%s", yyStatname(yystate))
				__yyfmt__.Printf(" saw %s\n", yyTokname(yychar))
			}
			fallthrough

		case 1, 2: /* incompletely recovered error ... try again */
			Errflag = 3

			/* find a state where "error" is a legal shift action */
			for yyp >= 0 {
				yyn = yyPact[yyS[yyp].yys] + yyErrCode
				if yyn >= 0 && yyn < yyLast {
					yystate = yyAct[yyn] /* simulate a shift of "error" */
					if yyChk[yystate] == yyErrCode {
						goto yystack
					}
				}

				/* the current p has no shift on "error", pop stack */
				if yyDebug >= 2 {
					__yyfmt__.Printf("error recovery pops state %d\n", yyS[yyp].yys)
				}
				yyp--
			}
			/* there is no state on the stack with an error shift ... abort */
			goto ret1

		case 3: /* no shift yet; clobber input char */
			if yyDebug >= 2 {
				__yyfmt__.Printf("error recovery discards %s\n", yyTokname(yychar))
			}
			if yychar == yyEofCode {
				goto ret1
			}
			yychar = -1
			goto yynewstate /* try again in the same state */
		}
	}

	/* reduction by production yyn */
	if yyDebug >= 2 {
		__yyfmt__.Printf("reduce %v in:\n\t%v\n", yyn, yyStatname(yystate))
	}

	yynt := yyn
	yypt := yyp
	_ = yypt // guard against "declared and not used"

	yyp -= yyR2[yyn]
	yyVAL = yyS[yyp+1]

	/* consult goto table to find next state */
	yyn = yyR1[yyn]
	yyg := yyPgo[yyn]
	yyj := yyg + yyS[yyp].yys + 1

	if yyj >= yyLast {
		yystate = yyAct[yyg]
	} else {
		yystate = yyAct[yyj]
		if yyChk[yystate] != -yyn {
			yystate = yyAct[yyg]
		}
	}
	// dummy call; replaced with literal code
	switch yynt {

	case 1:
		//line lang.y:46
		{
			d.tree = yyS[yypt-0].node
		}
	case 2:
		//line lang.y:50
		{
			yyVAL.node = ast.NewOperand(d.d(), ast.Eos)
		}
	case 3:
		//line lang.y:51
		{
			yyVAL.node = yyS[yypt-1].node
		}
	case 4:
		//line lang.y:52
		{
			yyVAL.node = yyS[yypt-0].node
		}
	case 5:
		//line lang.y:53
		{
			yyVAL.node = yyS[yypt-0].node
		}
	case 6:
		//line lang.y:54
		{
			yyVAL.node = yyS[yypt-0].node
		}
	case 7:
		//line lang.y:55
		{
			yyVAL.node = yyS[yypt-0].node
		}
	case 8:
		//line lang.y:59
		{
			yyVAL.node = yyS[yypt-0].node
		}
	case 9:
		//line lang.y:60
		{
			yyVAL.node = ast.NewOperand(d.d(), ast.Eos, yyS[yypt-1].node, yyS[yypt-0].node)
		}
	case 10:
		//line lang.y:64
		{
			yyVAL.node = yyS[yypt-1].node
		}
	case 11:
		//line lang.y:68
		{
			yyVAL.node = ast.NewOperand(d.d(), ast.Assign, ast.NewIdentifier(nil, yyS[yypt-3].identifier), yyS[yypt-1].node)
		}
	case 12:
		//line lang.y:72
		{
			yyVAL.node = ast.NewOperand(d.d(), ast.While, yyS[yypt-1].node, yyS[yypt-0].node)
		}
	case 13:
		//line lang.y:75
		{
			yyVAL.node = ast.NewOperand(d.d(), ast.If, yyS[yypt-2].node, yyS[yypt-1].node, yyS[yypt-0].node)
		}
	case 14:
		//line lang.y:78
		{
			yyVAL.node = ast.NewOperand(d.d(), ast.Eos)
		}
	case 15:
		//line lang.y:79
		{
			yyVAL.node = yyS[yypt-0].node
		}
	case 16:
		//line lang.y:80
		{
			yyVAL.node = yyS[yypt-0].node
		}
	case 17:
		//line lang.y:84
		{
			yyVAL.node = ast.NewOperand(d.d(), ast.Lt, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 18:
		//line lang.y:85
		{
			yyVAL.node = ast.NewOperand(d.d(), ast.Gt, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 19:
		//line lang.y:86
		{
			yyVAL.node = ast.NewOperand(d.d(), ast.Le, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 20:
		//line lang.y:87
		{
			yyVAL.node = ast.NewOperand(d.d(), ast.Ge, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 21:
		//line lang.y:88
		{
			yyVAL.node = ast.NewOperand(d.d(), ast.Ne, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 22:
		//line lang.y:89
		{
			yyVAL.node = ast.NewOperand(d.d(), ast.Eq, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 23:
		//line lang.y:90
		{
			yyVAL.node = yyS[yypt-1].node
		}
	case 24:
		//line lang.y:94
		{
			yyVAL.node = ast.NewInteger(d.d(), yyS[yypt-0].integer)
		}
	case 25:
		//line lang.y:95
		{
			yyVAL.node = ast.NewNumber(d.d(), yyS[yypt-0].number)
		}
	case 26:
		//line lang.y:96
		{
			yyVAL.node = ast.NewIdentifier(d.d(), yyS[yypt-0].identifier)
		}
	case 27:
		//line lang.y:97
		{
			yyVAL.node = ast.NewOperand(d.d(), ast.Uminus, yyS[yypt-0].node)
		}
	case 28:
		//line lang.y:98
		{
			yyVAL.node = ast.NewOperand(d.d(), ast.Add, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 29:
		//line lang.y:99
		{
			yyVAL.node = ast.NewOperand(d.d(), ast.Sub, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 30:
		//line lang.y:100
		{
			yyVAL.node = ast.NewOperand(d.d(), ast.Mul, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 31:
		//line lang.y:101
		{
			yyVAL.node = ast.NewOperand(d.d(), ast.Div, yyS[yypt-2].node, yyS[yypt-0].node)
		}
	case 32:
		//line lang.y:102
		{
			yyVAL.node = yyS[yypt-1].node
		}
	}
	goto yystack /* stack new state and value */
}
