//
// Copyright Coinbase, Inc. All Rights Reserved.
//
// SPDX-License-Identifier: Apache-2.0
//

package dealer

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/stretchr/testify/require"

	"github.com/sonrhq/sonr/crypto/core/curves"
	tt "github.com/sonrhq/sonr/crypto/internal"
	v1 "github.com/sonrhq/sonr/crypto/sharing/v1"
)

type proofParamsTest struct {
	bits                      uint
	p, q, f, alpha, n, h1, h2 *big.Int
}

func TestProofParamsDistinct(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping TestNewProofParams")
	}
	proofParams1, err := NewProofParams()
	require.NoError(t, err)
	proofParams2, err := NewProofParams()
	require.NoError(t, err)
	// Ensure two fresh params are distinct
	require.NotEqual(t, proofParams1.N, proofParams2.N)
	require.NotEqual(t, proofParams1.H1, proofParams2.H1)
	require.NotEqual(t, proofParams1.H2, proofParams2.H2)
	require.NotEqual(t, proofParams1.H1, proofParams2.H2)
	require.NotEqual(t, proofParams2.H1, proofParams1.H2)
}

func TestNewProofParamsKnownInputs(t *testing.T) {
	testValues := []*proofParamsTest{
		// Small values
		{
			bits:  32,
			p:     tt.B10("3532369103"),
			q:     tt.B10("3649146287"),
			n:     tt.B10("12890131596525970561"),
			f:     tt.B10("8088109122034248022"),
			alpha: tt.B10("1662308981529268576"),
			h1:    tt.B10("7740598546761717939"),
			h2:    tt.B10("682354612026984561"),
		},
		{
			bits:  32,
			p:     tt.B10("3944080607"),
			q:     tt.B10("3428458883"),
			n:     tt.B10("13522118192337181981"),
			f:     tt.B10("5840701024459274822"),
			alpha: tt.B10("2412272517404271893"),
			h1:    tt.B10("4853557737722970495"),
			h2:    tt.B10("5994934451042525431"),
		},
		// Moderate values
		{
			bits:  128,
			p:     tt.B10("322329737381194646436245717594125402007"),
			q:     tt.B10("282826037495283584562306655618616030003"),
			n:     tt.B10("91163242390418667918364564787170632477479160976682613402974201249506748416021"),
			f:     tt.B10("17638430582866531982508662722842743174558403736326834628201276897756690386088"),
			alpha: tt.B10("34615543197795812347714957897096525833286000039984371937454209042089659280135"),
			h1:    tt.B10("40851724117810430297136359930266015275847955981156240132756117915444116044667"),
			h2:    tt.B10("2951900472097897505645026011582677410773252477374130422613343770046117166202"),
		},
		{
			bits:  128,
			p:     tt.B10("292157587181890702665335334652613040839"),
			q:     tt.B10("325151157188330058469323448521916671923"),
			n:     tt.B10("94995377553542186895398706765950144779729213880991236718991546482239763663397"),
			f:     tt.B10("26654717808934280883230015589496263972208518550642662982505593055882359111187"),
			alpha: tt.B10("7966520803638388496765803768628435924825241054572312029050259003983212755370"),
			h1:    tt.B10("5179570387906113412737648546624738048023020835338868529502169468678133964537"),
			h2:    tt.B10("10547660360737303449442765975692388640388451365836709448045978219828751728319"),
		},
		{
			bits:  256,
			p:     tt.B10("107033239041727130475722739799997676073864393791537660670383998236210681119623"),
			q:     tt.B10("101044016059784163555459318277136957252800092661593566159166914417570906646479"),
			n:     tt.B10("10815068324662993508164204692909269429257853772524581783499643160896147777579932560873002543907262462663453338979819981987639157192530671167315407970757417"),
			f:     tt.B10("1953061770018232252742159833594632380255058093321781494393026400131699755086780416725368511941904170619021125367199106612073013879151950545295494471917586"),
			alpha: tt.B10("3834737493713501721754032502075239025227523335638258845915381809661597402119775003441676291390662996252213319663808209050925460873587874764050296101685895"),
			h1:    tt.B10("4320863187178906616455838434320546185913714725291410243809642735069293363508952325120395870462907728649344143321941759468514069154751317138801024570168772"),
			h2:    tt.B10("120718613217412490133304046154402631667794560287874694847992615611850656731360431202354938411320852540513514674759337083384626860820954644419583411766351"),
		},
		{
			bits:  384,
			p:     tt.B10("32084004412756729129951801667785515880675761456063995183947188458613139063207470311815946953832968900352953230728387"),
			q:     tt.B10("29916944379760549540437351937971471314912245460177593896040174973102477981114163011201660299005185209119438597828359"),
			n:     tt.B10("959855375496435098250249215425472550633883273892366636418539818992965634591339629957717390484127329139909742919598777040561784144823812817912999056501521038131720434324215292734547756922118558391569327009178732840170816605974926933"),
			f:     tt.B10("385587253104646973526979645949729332346264096947680952768033869130436857081349140581602832593968949445974766461768488933252048270498136280086338674644694131509548710522098707848381318033383109609217157024017362912090271855148305753"),
			alpha: tt.B10("802185911525259396858796733147029250385269666421140433965082754581471367564078646151444020878683709745451191116608505247291434446513232310868194940814603767895499634946931851453036837571701496215847350658468841482052820517809333229"),
			h1:    tt.B10("178008912521451997980828179602129531520630311485835136079559640359862691552613136589163589576614000919533806943298596513924271315778839665428154562258548640808571695748963797268761082739205874075388473490030003603923689503067183040"),
			h2:    tt.B10("745214268481251149009481704337976102510789515926575612825292375296159794448793242061611805039939202276919677151010882286558923931416568852289471804803844816398728119129991210671518959645714106367269729755312158228329249816980475270"),
		},
		// Large values
		{
			bits:  512,
			p:     tt.B10("13402506741879909705247508951754644667693815954510447147853401275166715648653624953727875501571997058765483588107170836999606735577380807635904979215471223"),
			q:     tt.B10("11653902873241241454023189143953837349357783829238806459834283425558693905220509553841347340456113574524451192031086644221307515208152789001650663503660563"),
			n:     tt.B10("156191511827829389348931232517809021193737031236087728136517838199389943369965856590107829262785004698503828473513808786810233768538839579197933368379417834150601329312673673331620361860375096296875273339782921082363787492351296298975304355190651921368035999591115305805841234398665476692295221538383486478549"),
			f:     tt.B10("105921637008007596242687136291215115379726727456238235988373686146222434936223500056214726834331554510844058841770826956376123791662340879537052405342520755293551902415159117602408529783890097619872151969618325933434785364935566958887498479829791790980528973576902755163741033898337074607062877295575097398027"),
			alpha: tt.B10("25789967854521552702486904432614912173720577957067512132821791246814877842721436035270615425020444434024971512294898663888264691831827542714695372062387891168110143236989908534429568576918984768687539132492516591672303697698170624423203504297159536671946076838229266401524840183833222277329757013745584001546"),
			h1:    tt.B10("109148140769505087324733385900286462326732849630583465544259040792765915286334491396382542664592507525818592725480821200175041450931363165947024379763315328559232391113072007292588635875180031961190886952268563630792514713201393588092419634985147285618899976278717726255619641744498785998181329854711823095546"),
			h2:    tt.B10("94378737602960275292210693440145810570305300627429873229499389511900935499225116793780120613621237391570231775304193964237372206306035564013470561242055886111069484714672545735756099242908140870434477147472839262127337437981240832226964726355298825863677601914170172556735529230548373855322923244573123279503"),
		},
		{
			bits:  768,
			p:     tt.B10("1172772268222869135746475976449967443220728362530972350209215332326944127327313132623800901521121085140613636265094266218948641934493008389679978199334246856112769202086824359372598961726505267139083231588791834137172543234789315807"),
			q:     tt.B10("1314405892055491780007828981914086343593182251638676771033374221066923695698195814335147429260730862492680087129273760043532686847298160063261905415386263791411919171601989373598015912958550608858030660357193476090089828897530026363"),
			n:     tt.B10("1541498779391422781877621899609182891994518423831078971398095064093240181958395952087587523810155282992062102677958880273569778308399535365070591300847802232397362358553902180950726216134823595994312857807899016362827564748344426602306284583250792792955584830704699574352786228055898787779684334887796554135523740299232752418576979598495344283249523357041872944823359212536062188372536232252883102629174357505712906918947289385294555564404297090523111007442619941"),
			f:     tt.B10("677989603002668106538546698801906687540953878286266137841947147184500778764869778659833934984587317604513819283010285362738864682953146933651161066202621461396232147445569683175461296649543802796602231833030736121285930269157674257535891093504924867936400869024264116704856719772498384123109010763539447471532256330132326434718941018313194910829655497567041642775909179031155007288364436570606906729509686118354750765498089348483611766753776448842378876138519492"),
			alpha: tt.B10("114288744075004384976717863018507283999460985427750993906297463893507876719107609356624383942527236695629514992610734269119471173538706034473306129332325931042872977327331297527020478557239451077438243523906598820487675455784275310274645270668758013883553399018822156980090134267245285532188867755106999667490195000880243418087891029246273568999920979861679097466175735879692809171018736759443882860496755786381993875691404126612248980033959651714747792276147737"),
			h1:    tt.B10("620967314571677993621097689227639136272103014757170315573975110856397012072491377414649496358698230838195501545740286798426409144254723458159053919007894986391912508160206658806451241744234372321843332002338225864899335573077301759675298903687495806772704923690781848994603501089797921531725638938738179420863221589774213217232543640290387805947665372254697970073247897667521511659457860642449029569029517623767487058427274486859786874744942405492665199950785432"),
			h2:    tt.B10("76914751581970807907441569885723431665256967585278657498014680460706690630690511580918136412833880806036782416407924930702533867597105318751032916063690482766856596872705517306464852735714057981896630822244930789132183799641349424925415175723846703728891058389586627867213281285663436964887344556696099523277630280786499528890816178997513077407863541104914946885418721934064536344132808255638281693284725391856052235328538580328974443057077286389922807376719663"),
		},
		{
			bits:  1024,
			p:     tt.B10("143755035639380470881675053832887329127071994369082228734251178758720295757727746866602803885188916682941316933470494309842642978838563803093231719031787157346353241144419080311914813130902891433725104776371931366053979438257736721584736553640620915684339223777632919031284192033212292380442862701004222839839"),
			q:     tt.B10("157610148541756483881391781006962391307552138517291043623605820440922385872835081567152281545287738290719615978200227103612218411158996749276828547434997453805807400563395441995925455411524662188622149635847181006072308124274294119590441349155209514553371838079002517661151936444244814496603165994642070193603"),
			n:     tt.B10("22657252520748253292205422817162431301953923432914829530688424232913850279325496327198502914522231560238552529734156383924448818535517634061008476071362010781638360092704508943571866960229942049437914690556866055765377519627454975682400206932320319743805083072214857842762721537739950074695623974079312071498296625705376593890814889314744719469735809152488403143751157723139035869185892099006348653635981206799193781030834368833947197930944812082594326193527332208252230115672713914945889734620959932802893197325106135662762752470236627025599443912886530954179753873735786171937758916890000958846322096261981191349917"),
			f:     tt.B10("1837638079389833117482811954367448162039150688934715357008176167854778196053567600507723116000842565891056468100639254152368756271384552650513939509902426167341409789748254644062944719126858575456567236001716746439508326036437929236231071381297886734438184468608864990027398302196877407296053037503604568299961010754385824604920003026632557650533262397423669648421709810685765745708396444919011990421325895569050175267461506858799263552873376629134365398033677037055272530963036099599403027301535753296115083718072540480783858873095117364364266062755097447466435769729951908277760553710182836153846164530411975604918"),
			alpha: tt.B10("10180338097046484720218938053574123628365802528878792921163521672678754788791478828324990427578268448710957792590062710921525369303666382188088935948462358688112248543323147863535957259946701818603149864941901611000864295172043472766639565354776537620683789487713327494120406803876018608041033830876741801587506282628346530619373455471179490747306225596783791589245228576159393496935812211359325406330371083545331242665519420884314024434338105230632567436938752428628410605413507478781802047209792610815229258139499934699515945155157568537597419579199189161997145564214140538430818340200949175908048368010772725324603"),
			h1:    tt.B10("1656240578726656472982971149349519787334711917407365099302234533544813508569815627980683589661064462432530865415359056057808102091523404507258432558908205534552293741530636994509001244876063100184310152841421199340736107615594714300348773466097241418448144156166001659986260228183904703220818795694087610910042320395937374686902326940538583574324286121691672644406892835183587590587202614663009833189341162830624857807052409882165022116236155557756122995333309543614144793788263137215939462312172467168512721258271559466365919027756907792017876708485439503402393043941731729574262940831871878730655264809864217559480"),
			h2:    tt.B10("14402453257845659067057298000258011503319998232626269706984076352133770609685418836115347557347426169382037951633347064435045381846434457869015364415080444373680323708145947051971271955375310253866278892655687643826235776329111596605765760603463133947524568629075859958729387193155836149415340890049121874989262064208424476603584205725570597892845954158554057663626043501240089826176762119712071675190787646116090239837347676973220240416596698697171855421900633747224108875362350155229834913561333497452922143478131464246300744179904539036740021278776277717870280483008305379221666473338482907287695159341884476431429"),
		},
	}

	for _, test := range testValues {
		pIdx := 0
		safePrimes := []*big.Int{test.p, test.q}
		randVars := []*big.Int{test.f, test.alpha}
		f := func(bits uint) (*big.Int, error) {
			r := safePrimes[pIdx]
			pIdx = (pIdx + 1) % 2
			return r, nil
		}
		rIdx := 0
		g := func(nn *big.Int) (*big.Int, error) {
			r := randVars[rIdx]
			rIdx = (rIdx + 1) % 2
			return r, nil
		}

		params, err := genProofParams(f, g, test.bits)
		require.NoError(t, err)
		require.Equal(t, params.N, test.n)
		require.Equal(t, params.H1, test.h1)
		require.Equal(t, params.H2, test.h2)
	}
}

func TestSameResult(t *testing.T) {
	test := proofParamsTest{
		bits:  32,
		p:     tt.B10("3532369103"),
		q:     tt.B10("3649146287"),
		n:     tt.B10("12890131596525970561"),
		f:     tt.B10("8088109122034248022"),
		alpha: tt.B10("1662308981529268576"),
		h1:    tt.B10("7740598546761717939"),
		h2:    tt.B10("682354612026984561"),
	}

	idx := 0
	safePrimes := []*big.Int{test.p, test.q}
	f := func(bits uint) (*big.Int, error) {
		r := safePrimes[idx]
		idx = (idx + 1) % 2
		return r, nil
	}
	randVars := []*big.Int{test.f, test.alpha}
	rIdx := 0
	g := func(nn *big.Int) (*big.Int, error) {
		r := randVars[rIdx]
		rIdx = (rIdx + 1) % 2
		return r, nil
	}
	params1, err := genProofParams(f, g, 32)
	require.NoError(t, err)
	params2, err := genProofParams(f, g, 32)
	require.NoError(t, err)
	require.Equal(t, params1.N, params2.N)
	require.Equal(t, params1.H1, params2.H1)
	require.Equal(t, params1.H2, params2.H2)
}

func TestNewDealerShares(t *testing.T) {
	curve := btcec.S256()
	for _, secretIsNil := range []bool{true, false} {
		t.Run(fmt.Sprintf("NewDealerShare should not fail if bool(ikm == nil) is %t", secretIsNil), func(t *testing.T) {
			var ikm *big.Int
			var err error
			if secretIsNil {
				ikm, err = NewSecret(curve)
				require.NoError(t, err)
			}
			pk, sharesMap, err := NewDealerShares(curve, 2, 3, ikm)
			if err != nil {
				t.Errorf("NewDealerShares failed: %v", err)
				t.FailNow()
			}

			if pk == nil {
				t.Errorf("NewDealerShares public key is nil")
				t.FailNow()
			}

			if secretIsNil {
				derivedPublicKey, err := DerivePublicKey(curve, ikm)
				require.NoError(t, err)
				require.Equal(t, pk.X, derivedPublicKey.X)
				require.Equal(t, pk.Y, derivedPublicKey.Y)
			}

			if len(sharesMap) != 3 {
				t.Errorf("NewDealerShares didn't produce enough shares")
				t.FailNow()
			}

			for _, s := range sharesMap {
				if s.ShamirShare == nil {
					t.Errorf("NewDealerShares didn't produce valid sharesMap")
					t.FailNow()
				}
				if s.Point == nil {
					t.Errorf("NewDealerShares didn't produce valid public sharesMap")
					t.FailNow()
				}
				x, y := curve.ScalarMult(curve.Gx, curve.Gy, s.ShamirShare.Value.Bytes())
				require.Equal(t, x, s.Point.X)
				require.Equal(t, y, s.Point.Y)
			}

			n := curves.NewField(curve.N)
			combiner, err := v1.NewShamir(2, 3, n)
			require.NoError(t, err)

			sShareArray := make([]*v1.ShamirShare, len(sharesMap))
			for i, s := range sharesMap {
				// Skip sharesMap[0] which is always nil
				if i == 0 && s != nil {
					t.Errorf("expected sharesMap[0] to be nil")
					t.FailNow()
				} else {
					sShareArray[i-1] = s.ShamirShare
				}
			}
			sk, err := combiner.Combine(sShareArray...)
			if err != nil {
				t.Errorf("Shares could not be recombined")
			}
			pkx, pky := curve.ScalarBaseMult(sk)
			require.NotNil(t, pkx, pky)
			require.Equal(t, pkx, pk.X) // nolint
			require.Equal(t, pky, pk.Y) // nolint

		})
	}
}

func TestPreparePublicShares(t *testing.T) {
	curve := btcec.S256()
	pk, sharesMap, err := NewDealerShares(curve, 2, 3, nil)
	if err != nil {
		t.Errorf("NewDealerShares failed: %v", err)
	}

	if pk == nil {
		t.Errorf("NewDealerShares public key is nil")
	}

	if len(sharesMap) != 3 {
		t.Errorf("NewDealerShares didn't produce enough sharesMap")
	}

	publicShares, err := PreparePublicShares(sharesMap)
	if err != nil {
		t.Errorf("PreparePublicShares failed: %v", err)
	}

	if len(publicShares) != len(sharesMap) {
		t.Errorf("len(publicShares) != len(sharesMap): %d != %d", len(publicShares), len(sharesMap))
	}
	for i := range publicShares {
		require.Equal(t, publicShares[i].Point.X, sharesMap[i].Point.X)
		require.Equal(t, publicShares[i].Point.Y, sharesMap[i].Point.Y)
	}
}
