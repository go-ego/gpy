package gpy_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/go-ego/gpy"
	"github.com/go-ego/gpy/phrase"
	"github.com/go-ego/gse"
)

var hans50 = `红牛公司（Red Bull GmbH）是一家创办于奥地利的运动饮料公司。
红牛公司的总部位于奥地利的滨湖富施尔。2013年，“红牛”运动饮料在全球的销量
达到了53.87亿罐，比2012年增长3.1%。`

var hans500 = strings.Replace(strings.Replace(`
的、一、是、在、不、了、有、和、人、这、中、大、为、上、个、国、我、以、要、他、
时、来、用、们、生、到、作、地、于、出、就、分、对、成、会、可、主、发、年、动、
同、工、也、能、下、过、子、说、产、种、面、而、方、后、多、定、行、学、法、所、
民、得、经、十、三、之、进、着、等、部、度、家、电、力、里、如、水、化、高、自、
二、理、起、小、物、现、实、加、量、都、两、体、制、机、当、使、点、从、业、本、
去、把、性、好、应、开、它、合、还、因、由、其、些、然、前、外、天、政、四、日、
那、社、义、事、平、形、相、全、表、间、样、与、关、各、重、新、线、内、数、正、
心、反、你、明、看、原、又、么、利、比、或、但、质、气、第、向、道、命、此、变、
条、只、没、结、解、问、意、建、月、公、无、系、军、很、情、者、最、立、代、想、
已、通、并、提、直、题、党、程、展、五、果、料、象、员、革、位、入、常、文、总、
次、品、式、活、设、及、管、特、件、长、求、老、头、基、资、边、流、路、级、少、
图、山、统、接、知、较、将、组、见、计、别、她、手、角、期、根、论、运、农、指、
几、九、区、强、放、决、西、被、干、做、必、战、先、回、则、任、取、据、处、队、
南、给、色、光、门、即、保、治、北、造、百、规、热、领、七、海、口、东、导、器、
压、志、世、金、增、争、济、阶、油、思、术、极、交、受、联、什、认、六、共、权、
收、证、改、清、己、美、再、采、转、更、单、风、切、打、白、教、速、花、带、安、
场、身、车、例、真、务、具、万、每、目、至、达、走、积、示、议、声、报、斗、完、
类、八、离、华、名、确、才、科、张、信、马、节、话、米、整、空、元、况、今、集、
温、传、土、许、步、群、广、石、记、需、段、研、界、拉、林、律、叫、且、究、观、
越、织、装、影、算、低、持、音、众、书、布、复、容、儿、须、际、商、非、验、连、
断、深、难、近、矿、千、周、委、素、技、备、半、办、青、省、列、习、响、约、支、
般、史、感、劳、便、团、往、酸、历、市、克、何、除、消、构、府、称、太、准、精、
值、号、率、族、维、划、选、标、写、存、候、毛、亲、快、效、斯、院、查、江、型、
眼、王、按、格、养、易、置、派、层、片、始、却、专、状、育、厂、京、识、适、属、
圆、包、火、住、调、满、县、局、照、参、红、细、引、听、该、铁、价、严、龙、飞
`, "、", "", -1), "\n", "", -1)

func benchmarkPinyin(b *testing.B, s string, args gpy.Args) {
	b.StopTimer()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		gpy.Pinyin(s, args)
	}
}

func benchmarkPinyinString(b *testing.B, s string, args gpy.Args) {
	b.StopTimer()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		py := gpy.Pinyin(s, args)
		gpy.ToString(py)
	}
}

func benchmarkPyString(b *testing.B, s string, args gpy.Args) {
	b.StopTimer()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		gpy.Py(s, args)
	}
}

func benchmarkHanPinyin(b *testing.B, s string, args gpy.Args) {
	b.StopTimer()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		gpy.HanPinyin(s, args)
	}
}

func benchmarkLazyPinyin(b *testing.B, s string, args gpy.Args) {
	b.StopTimer()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		gpy.LazyPinyin(s, args)
	}
}

func benchmark_PhrasePinyin(b *testing.B, s string) {
	b.StopTimer()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		phrase.Paragraph(s)
	}
}

var seg = gse.New()

func benchmark_PhrasePinyin_Seg(b *testing.B, s string) {
	b.StopTimer()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		phrase.Paragraph(s, seg)
	}
}

func BenchmarkPinyinOne(b *testing.B) {
	args := gpy.NewArgs()
	benchmarkPinyin(b, "中", args)
}

func Benchmark_Pinyin_500(b *testing.B) {
	args := gpy.NewArgs()
	benchmarkPinyin(b, hans500, args)
}

func Benchmark_Pinyin_Sting_500(b *testing.B) {
	args := gpy.NewArgs()
	benchmarkPinyinString(b, hans500, args)
}

func Benchmark_Py_Sting_500(b *testing.B) {
	args := gpy.NewArgs()
	benchmarkPyString(b, hans500, args)
}

func BenchmarkHanPinyinOne(b *testing.B) {
	args := gpy.NewArgs()
	benchmarkHanPinyin(b, "中", args)
}

func Benchmark_HanPinyin_500(b *testing.B) {
	args := gpy.NewArgs()
	benchmarkHanPinyin(b, hans500, args)
}

func init() {
	fmt.Println(hans500)
	phrase.LoadGseDict()
}

func Benchmark_PhrasePinyin_50(b *testing.B) {
	benchmark_PhrasePinyin(b, hans50)
}

func Benchmark_PhrasePinyin_Seg_50(b *testing.B) {
	benchmark_PhrasePinyin_Seg(b, hans50)
}

func BenchmarkLazyPinyinOne(b *testing.B) {
	args := gpy.NewArgs()
	benchmarkLazyPinyin(b, "中", args)
}

func Benchmark_LazyPinyin_500(b *testing.B) {
	args := gpy.NewArgs()
	benchmarkLazyPinyin(b, hans500, args)
}
