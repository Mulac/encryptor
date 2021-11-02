package encryptor

import (
	"fmt"
	"testing"
)

var testData = []struct {
	decrypted string
	encrypted string
	key       int
}{
	{"a", "b", 1},
	{"z", "b", 2},
	{"a", "z", -1},
	{"z", "x", -2},
	{"hello", "khoor", 3},
	{"hello", "uryyb", 13},
	{"hELlo", "uRYyb", 13},
	{
		"This above all: to thine own self be true.",
		"Wklv deryh doo: wr wklqh rzq vhoi eh wuxh.",
		3,
	},
}

func TestCaesarCipherEncrypt(t *testing.T) {
	encryptor := &caesarCipher{}
	for _, test := range testData {
		t.Run(fmt.Sprintf("Encrypting %s with key %d", test.decrypted, test.key), func(t *testing.T) {
			actual, err := encryptor.Encrypt(test.decrypted, Key(test.key))
			if err != nil {
				t.Fatalf("FAIL: failed to encrypt '%s' with key '%d', err: %v",
					test.decrypted, test.key, err)
			}
			if actual != test.encrypted {
				t.Errorf("Got '%s' when encrypting '%s' with key '%d', expected '%s'",
					actual, test.decrypted, test.key, test.encrypted)
			}
		})
	}
}

func TestCaesarCipherDecrypt(t *testing.T) {
	encryptor := &caesarCipher{}
	for _, test := range testData {
		t.Run(fmt.Sprintf("Encrypting %s with key %d", test.decrypted, test.key), func(t *testing.T) {
			actual, err := encryptor.Decrypt(test.encrypted, Key(test.key))
			if err != nil {
				t.Fatalf("FAIL: failed to decrypt '%s' with key '%d', err: %v",
					test.encrypted, test.key, err)
			}
			if actual != test.decrypted {
				t.Errorf("Got '%s' when decrypting '%s' with key '%d', expected '%s'",
					actual, test.encrypted, test.key, test.decrypted)
			}
		})
	}
}

func BenchmarkCaesarCipherEncrypt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		new(caesarCipher).Encrypt(testText, Key(22))
	}
}

var testText string = `
	Lorem ipsum dolor sit amet. Est officiis dolor vel dolores nobis non iure soluta cum illo dignissimos eum laudantium rerum ut voluptates sunt est fuga enim! Ut quis dolore qui consequatur vitae velit sequi et dolorem voluptatem debitis deleniti. Eum quia iure eos itaque sequi est velit corporis ea quos culpa et explicabo beatae quo odit ducimus et praesentium nihil. Cum voluptatum unde sit amet repellat et nisi sapiente. Sed impedit tempore sit labore optio a voluptas asperiores est quibusdam internos rem veniam repudiandae. Eum mollitia cupiditate quo aperiam voluptate et nesciunt tenetur. Et exercitationem perspiciatis aut accusantium consequatur qui aliquid soluta est iusto velit cum beatae excepturi. Non distinctio modi ut aliquam dicta ut rerum delectus qui sunt perspiciatis ut dolorum internos.
	Et mollitia voluptatem non accusantium enim ea dolorem facere aut enim quisquam et nulla autem qui exercitationem cumque? Ut deserunt error sed eaque quibusdam rem quis adipisci quo voluptatibus consequatur aut eaque maiores ut nihil aperiam. Hic molestiae internos At officia alias est animi labore ut commodi nemo vel doloremque voluptatem? Qui harum magnam ut illum iusto aut repellat labore rem inventore soluta et consequuntur itaque sed dolorem illo. Vel vitae ullam vel excepturi debitis sed commodi pariatur aut officia molestiae et doloremque facere. Ut nulla dolores eos aliquid nostrum ab eveniet eveniet est magnam quis et enim quae qui quia deserunt. Aut placeat commodi aut aperiam suscipit sit laudantium quia qui tempore pariatur aut dolor sint. Aut iusto sint nam suscipit deleniti nam amet voluptates a internos totam qui sint optio et veniam quasi cum iure reiciendis. Quo provident sunt qui nesciunt nostrum et quas recusandae aut sint aperiam 33 nihil reiciendis. Et quod facere aut perspiciatis voluptas ut quos nostrum sit consequatur exercitationem. Et unde quod quia illum rem molestiae laborum. Et quis tempora eum quis provident doloremque magni! In voluptatibus cupiditate est ipsa possimus et itaque galisum sit expedita dolores qui facere omnis At corporis quos. Ut internos nulla ea quam illo At repudiandae quia eos rerum porro sed adipisci alias?
	Aut molestias molestiae et cupiditate eligendi qui dolorum consequatur id soluta facilis in rerum internos. Quo officia vero et tempora repellendus qui voluptatem internos aut ratione repellendus id harum molestiae qui veritatis debitis. Ex magnam velit sit rerum nobis sit porro odio aut fuga architecto non quisquam laudantium ut laudantium optio. Est maxime vero id quisquam laudantium rem asperiores beatae sit quia dicta a quaerat corporis et culpa porro. 33 impedit unde ab harum delectus est voluptates esse ab quia facilis aut beatae molestiae et voluptas odit sed cumque Quis. Et nobis omnis qui accusantium voluptas ut placeat deserunt qui optio deserunt est quod delectus 33 tenetur accusamus! Ea laboriosam quibusdam hic incidunt laborum et repellat maxime non molestiae numquam. Et animi fugiat est vero quisquam qui quis debitis ut voluptatibus repellat in reprehenderit quod qui earum dolorem. Et blanditiis voluptas nam voluptate omnis vel fuga voluptatem! Eum laborum perspiciatis est ipsa totam in officia nostrum et voluptatum nisi ut placeat consequatur. Sit quia neque in totam amet nam nemo omnis aut debitis ipsa eum dignissimos ipsa. Ut exercitationem eligendi ut labore nostrum ex nulla consequatur. Est minima similique et dolorem praesentium non magni quis ut asperiores expedita vel quas dolore ad consequatur natus. Ex magni deserunt ut nesciunt quidem sit neque minus odit debitis qui quod dolorum.
	Eum nobis enim non minima odio et voluptatibus sint. Non numquam maiores qui sint maiores quo illum quas aut quae voluptatem ad laudantium temporibus ut sunt suscipit. Aut voluptate dicta sed omnis excepturi est earum vero ut necessitatibus sint! Et eveniet provident et provident ducimus ut dolore minima. Ut laboriosam vitae non voluptates blanditiis qui exercitationem iure At illum Quis aut exercitationem explicabo 33 magni alias! Qui esse architecto id atque aspernatur sed voluptatem ipsa qui sapiente voluptas. Id sequi perferendis qui odio tempore eum sint asperiores qui placeat fuga. Aut officiis excepturi ut ratione ipsam id rerum debitis sed suscipit laborum et suscipit consequuntur aut excepturi necessitatibus in rerum dolores.
	Qui aliquid quidem aut enim reiciendis quo quaerat odit qui natus consectetur eos asperiores voluptatem! Ut dolore sed aperiam repudiandae sit ullam vitae ut repellat quia et possimus placeat. Aut omnis consectetur est maxime voluptatem vel esse pariatur in modi nam voluptas laudantium? Ut nesciunt quibusdam est magni laborum 33 quaerat unde qui numquam dolore aut dolorum eaque. In magni id dolore iste sed eligendi obcaecati qui obcaecati atque. Est necessitatibus ipsam et quis voluptas et facilis incidunt eos quam excepturi id nisi expedita. Et eveniet fuga vel illo voluptatum ea autem tenetur et neque architecto nam distinctio quia 33 fugit quae. Non quia fugit eos delectus iure ea reprehenderit nihil ut totam tenetur ea fugiat facilis et minus natus. Eum ipsum architecto et delectus internos et veritatis nihil et internos quia rem provident aliquam ut placeat aliquid. Sit repellendus dolorum qui iusto voluptatibus et ipsam internos non expedita commodi minima nihil. Id assumenda quod ut iste repellat non aliquid voluptas qui corporis nulla hic maxime galisum aut eligendi vero qui harum saepe.
	Et numquam sunt qui sint maiores et quia maxime aut facere porro sed quas dolor et fuga saepe aut deleniti odio. Id libero dolor a ullam accusamus sit sunt quia ut dicta reiciendis. Id quia aliquid est libero voluptatem eos aperiam fuga sed voluptas aliquid ut animi reiciendis non adipisci iure non iusto voluptatem. At galisum dolore est enim fuga et consequatur harum nam totam fuga ex ducimus omnis ut internos dolor ab maiores deleniti. Et perspiciatis dignissimos et reprehenderit facere nam internos aliquid qui eaque fugiat? Ab Quis sint hic consequatur aliquid et consectetur molestiae est illum nulla ut architecto vero est enim error est pariatur officia. Ex mollitia voluptates ut aspernatur veniam et corporis similique qui corrupti totam aut ipsam reprehenderit sit aliquid voluptates 33 minus voluptate. Et quia excepturi ab incidunt libero est molestias molestiae ad laboriosam unde sed ipsa sint?
	In galisum cumque nam autem distinctio sed alias aliquid quo reprehenderit? Vel quia veniam est ducimus corrupti in molestias consequatur eos quia officiis nam quod aspernatur id obcaecati impedit rem tenetur autem? Vel repudiandae eius non perferendis accusantium cum inventore nihil 33 itaque architecto est illo maxime sed nisi voluptates? Et maxime odio cum suscipit fuga ut voluptas neque aut dicta aliquid nam atque eaque quia sint vel deserunt blanditiis? Rem reprehenderit tempora quo laborum quia aut rerum ipsa qui quasi obcaecati est magnam sint et ipsa eligendi? Sed rerum temporibus et dolores officia et nemo eius. Aut aperiam ducimus et dicta minus est accusamus sint ut unde mollitia eum mollitia iste. Est reprehenderit explicabo ut veritatis quas et corporis doloribus ut officia reiciendis in tenetur reprehenderit aut galisum libero ab aperiam repellendus.
	Sed delectus ipsum ad corrupti placeat rem nemo quae quo natus deleniti aut quam molestiae ab quibusdam aliquam in sequi distinctio. At incidunt odio ea consectetur tempora est quod vero sed error illum et officiis rerum. Sit aliquam voluptas ea modi porro est voluptatem illo. Est ratione voluptatem qui eligendi dolorem nam voluptatem sunt eum necessitatibus saepe et cumque deleniti sed molestiae facere? Ut enim culpa et nihil voluptas et voluptatem voluptate et pariatur voluptates. Est reprehenderit vitae sed velit labore et sint culpa vel dignissimos voluptatem est nostrum similique qui totam perspiciatis. Id deserunt placeat dolores velit ea nihil explicabo? Ut repellendus pariatur et iusto quibusdam aut atque galisum non iusto quis non veniam sapiente sit quam laudantium. Est quia voluptas eos incidunt aliquam qui consequuntur officia et molestiae quisquam. Ad facilis recusandae non velit quia et ducimus aliquid vel quis internos in dolore alias id enim sunt. Et tempora perferendis aut galisum totam ad enim internos qui rerum laudantium hic voluptatibus architecto eum nesciunt nostrum ut velit odio!
	Ex dicta impedit id odio aperiam ut iure possimus ex adipisci consequatur et molestiae quia ut nobis laboriosam eos voluptas tempore. Ut vitae voluptas et molestiae internos ea neque omnis nam expedita veniam. Ut perspiciatis optio in rerum alias et dicta temporibus aut unde internos vel voluptatibus quaerat. Aut aperiam ipsum et adipisci dolores qui laudantium eligendi sed harum praesentium At Quis alias. Qui cumque consequatur et perspiciatis voluptatibus ea porro voluptas et dolorem quis qui autem porro qui doloremque dicta est autem aperiam! Et esse earum eum veniam Quis At cupiditate voluptate ut vero rerum et quia fugit est velit voluptas. In asperiores explicabo est nobis laborum qui voluptatem voluptatem qui laborum facere 33 repellendus ratione et reprehenderit reiciendis ut officiis fugiat! Et asperiores nisi non perferendis doloremque et nobis perferendis non error velit ab numquam amet et dolor quam et omnis veniam.
	Qui eaque galisum eos quia accusantium ut consequatur quis. At dignissimos obcaecati ad minus natus ut esse harum sit temporibus quia similique facilis hic suscipit harum. Sed doloribus maxime qui esse pariatur ad molestias galisum quo dolorem dolorum ut maxime autem non fugiat soluta! Ut mollitia porro aut incidunt iure et labore optio non impedit maxime est tempora cumque aut labore ipsa est ullam corrupti. Galisum iure vel repellat laboriosam qui itaque enim aut vitae porro eum distinctio aliquid aut suscipit unde aut modi odit. Aut nobis ducimus est consequatur dolores aut officiis corrupti qui veniam voluptatem qui esse incidunt rem vero accusantium ut commodi maiores? Et rerum veniam aut inventore assumenda id dolorum impedit qui quia dolor eos facilis galisum! Est corporis quia cum voluptate voluptatem et tempora tempore sed vero maxime qui perferendis quia. Sed voluptatem dolores et suscipit provident aut vitae accusantium ad deleniti nihil. Ea voluptatem accusantium ea galisum excepturi et exercitationem perferendis rem obcaecati saepe et labore ducimus? Aut possimus fuga hic sint nobis vel facere necessitatibus id sint accusantium est recusandae atque!
`
