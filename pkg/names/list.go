package names

import "math/rand"

func GetRandomName() string {
	n := rand.Intn(372)
	return FourLetterNames[n]
}

var FourLetterNames = [372]string{
	"abel",
	"adam",
	"adan",
	"addy",
	"aden",
	"ajax",
	"alan",
	"aldo",
	"alec",
	"alex",
	"alto",
	"alva",
	"alvy",
	"amal",
	"ames",
	"amir",
	"amos",
	"andy",
	"ares",
	"arie",
	"arlo",
	"aron",
	"aven",
	"axel",
	"ayan",
	"aziz",
	"bart",
	"bear",
	"beau",
	"beck",
	"bert",
	"beto",
	"bill",
	"bing",
	"bird",
	"blue",
	"bode",
	"bolt",
	"bond",
	"bose",
	"boss",
	"boyd",
	"brad",
	"bram",
	"bran",
	"bret",
	"brio",
	"buck",
	"budd",
	"burk",
	"burl",
	"burr",
	"burt",
	"bush",
	"cade",
	"cage",
	"cain",
	"cale",
	"carl",
	"cary",
	"case",
	"cash",
	"cato",
	"chad",
	"chet",
	"chip",
	"clay",
	"clem",
	"cobb",
	"coby",
	"codi",
	"cody",
	"coen",
	"cole",
	"colt",
	"cory",
	"cree",
	"crew",
	"cris",
	"cruz",
	"curt",
	"dale",
	"dane",
	"dash",
	"dave",
	"davy",
	"dean",
	"deon",
	"desi",
	"dick",
	"dino",
	"dion",
	"dior",
	"doug",
	"drew",
	"duke",
	"earl",
	"eddy",
	"eden",
	"edge",
	"elby",
	"elio",
	"elmo",
	"elon",
	"emil",
	"emir",
	"emry",
	"emyr",
	"enzo",
	"epic",
	"eric",
	"erik",
	"erle",
	"eros",
	"esau",
	"evan",
	"ewan",
	"ezra",
	"finn",
	"ford",
	"fred",
	"gabe",
	"gael",
	"gage",
	"gary",
	"gene",
	"gibb",
	"glen",
	"gray",
	"grey",
	"guss",
	"hale",
	"hank",
	"hans",
	"hawk",
	"hays",
	"herb",
	"holt",
	"hoyt",
	"huck",
	"huey",
	"hugh",
	"hugo",
	"hunt",
	"hyde",
	"iago",
	"igor",
	"iker",
	"indy",
	"isom",
	"ivan",
	"iver",
	"izzy",
	"jace",
	"jack",
	"jair",
	"jake",
	"jase",
	"jaxx",
	"jazz",
	"jean",
	"jeff",
	"jess",
	"jett",
	"joah",
	"jody",
	"joel",
	"joey",
	"john",
	"jose",
	"josh",
	"juan",
	"judd",
	"jude",
	"kace",
	"kade",
	"kale",
	"kane",
	"karl",
	"kase",
	"kash",
	"kato",
	"kent",
	"kerr",
	"khan",
	"kian",
	"kiel",
	"kiev",
	"king",
	"kipp",
	"kirk",
	"knox",
	"kobe",
	"koda",
	"kody",
	"kole",
	"krew",
	"kurt",
	"kyle",
	"kylo",
	"kyng",
	"kyro",
	"lake",
	"lane",
	"lars",
	"leaf",
	"leif",
	"leon",
	"levi",
	"levy",
	"liam",
	"lian",
	"liev",
	"link",
	"loki",
	"luca",
	"ludo",
	"luis",
	"luka",
	"luke",
	"lyle",
	"lynx",
	"mace",
	"mack",
	"mako",
	"marc",
	"mark",
	"mars",
	"matt",
	"mell",
	"merl",
	"mica",
	"mick",
	"mike",
	"milo",
	"mitt",
	"mont",
	"moss",
	"musa",
	"mylo",
	"nash",
	"nate",
	"navy",
	"neil",
	"neon",
	"nero",
	"newt",
	"nial",
	"nick",
	"nico",
	"niko",
	"nile",
	"nilo",
	"noah",
	"noel",
	"noon",
	"nova",
	"nyle",
	"odin",
	"odis",
	"olaf",
	"omar",
	"omni",
	"onyx",
	"opus",
	"orlo",
	"otho",
	"otis",
	"otto",
	"owen",
	"paco",
	"park",
	"paul",
	"penn",
	"pete",
	"phil",
	"pike",
	"poet",
	"polk",
	"quil",
	"quin",
	"rafe",
	"rain",
	"rand",
	"raul",
	"reed",
	"reef",
	"reid",
	"reif",
	"remi",
	"remy",
	"rene",
	"rhys",
	"rice",
	"rich",
	"rick",
	"rico",
	"riot",
	"roan",
	"robb",
	"rock",
	"rolf",
	"rome",
	"rook",
	"rory",
	"ross",
	"rowe",
	"rudy",
	"rush",
	"russ",
	"ryan",
	"sage",
	"saul",
	"sean",
	"seth",
	"shaw",
	"shel",
	"shep",
	"skip",
	"snow",
	"solo",
	"stan",
	"sven",
	"taft",
	"tate",
	"teal",
	"thad",
	"theo",
	"thom",
	"thor",
	"tito",
	"toby",
	"todd",
	"tony",
	"tosh",
	"trey",
	"troy",
	"tuan",
	"tuck",
	"uley",
	"utah",
	"vern",
	"vick",
	"vito",
	"vlad",
	"voss",
	"wade",
	"walt",
	"webb",
	"west",
	"will",
	"whit",
	"witt",
	"wolf",
	"wynn",
	"xavi",
	"xyan",
	"yale",
	"york",
	"yuri",
	"zach",
	"zack",
	"zaid",
	"zain",
	"zale",
	"zane",
	"zayd",
	"zayn",
	"zeke",
	"zeus",
	"zion",
	"zuma",
}
