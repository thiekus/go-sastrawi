package sastrawi

import "testing"

type testItem struct {
	value    string
	expected string
}

func TestStemmer(t *testing.T) {
	testItems := []testItem{
		testItem{value: "mei", expected: "mei"},
		testItem{value: "bui", expected: "bui"},
		testItem{value: "nilai", expected: "nilai"},
		testItem{value: "hancurlah", expected: "hancur"},
		testItem{value: "benarkah", expected: "benar"},
		testItem{value: "apatah", expected: "apa"},
		testItem{value: "siapapun", expected: "siapa"},
		testItem{value: "jubahku", expected: "jubah"},
		testItem{value: "bajumu", expected: "baju"},
		testItem{value: "celananya", expected: "celana"},
		testItem{value: "hantui", expected: "hantu"},
		testItem{value: "belikan", expected: "beli"},
		testItem{value: "jualan", expected: "jual"},
		testItem{value: "bukumukah", expected: "buku"},
		testItem{value: "miliknyalah", expected: "milik"},
		testItem{value: "kulitkupun", expected: "kulit"},
		testItem{value: "berikanku", expected: "beri"},
		testItem{value: "sakitimu", expected: "sakit"},
		testItem{value: "beriannya", expected: "beri"},
		testItem{value: "kasihilah", expected: "kasih"},
		testItem{value: "dibuang", expected: "buang"},
		testItem{value: "kesakitan", expected: "sakit"},
		testItem{value: "sesuap", expected: "suap"},
		testItem{value: "beradu", expected: "adu"},
		testItem{value: "berambut", expected: "rambut"},
		testItem{value: "bersuara", expected: "suara"},
		testItem{value: "berdaerah", expected: "daerah"},
		testItem{value: "belajar", expected: "ajar"},
		testItem{value: "bekerja", expected: "kerja"},
		testItem{value: "beternak", expected: "ternak"},
		testItem{value: "terasing", expected: "asing"},
		testItem{value: "teraup", expected: "raup"},
		testItem{value: "tergerak", expected: "gerak"},
		testItem{value: "terpuruk", expected: "puruk"},
		testItem{value: "teterbang", expected: "terbang"},
		testItem{value: "melipat", expected: "lipat"},
		testItem{value: "meringkas", expected: "ringkas"},
		testItem{value: "mewarnai", expected: "warna"},
		testItem{value: "meyakinkan", expected: "yakin"},
		testItem{value: "membangun", expected: "bangun"},
		testItem{value: "memfitnah", expected: "fitnah"},
		testItem{value: "memvonis", expected: "vonis"},
		testItem{value: "memperbaru", expected: "baru"},
		testItem{value: "mempelajar", expected: "ajar"},
		testItem{value: "meminum", expected: "minum"},
		testItem{value: "memukul", expected: "pukul"},
		testItem{value: "mencinta", expected: "cinta"},
		testItem{value: "mendua", expected: "dua"},
		testItem{value: "menjauh", expected: "jauh"},
		testItem{value: "menziarah", expected: "ziarah"},
		testItem{value: "menuklir", expected: "nuklir"},
		testItem{value: "menangkap", expected: "tangkap"},
		testItem{value: "menggila", expected: "gila"},
		testItem{value: "menghajar", expected: "hajar"},
		testItem{value: "mengqasar", expected: "qasar"},
		testItem{value: "mengudara", expected: "udara"},
		testItem{value: "mengupas", expected: "kupas"},
		testItem{value: "menyuarakan", expected: "suara"},
		testItem{value: "mempopulerkan", expected: "populer"},
		testItem{value: "pewarna", expected: "warna"},
		testItem{value: "peyoga", expected: "yoga"},
		testItem{value: "peradilan", expected: "adil"},
		testItem{value: "perumahan", expected: "rumah"},
		testItem{value: "permuka", expected: "muka"},
		testItem{value: "perdaerah", expected: "daerah"},
		testItem{value: "pembangun", expected: "bangun"},
		testItem{value: "pemfitnah", expected: "fitnah"},
		testItem{value: "pemvonis", expected: "vonis"},
		testItem{value: "peminum", expected: "minum"},
		testItem{value: "pemukul", expected: "pukul"},
		testItem{value: "pencinta", expected: "cinta"},
		testItem{value: "pendua", expected: "dua"},
		testItem{value: "penjauh", expected: "jauh"},
		testItem{value: "penziarah", expected: "ziarah"},
		testItem{value: "penuklir", expected: "nuklir"},
		testItem{value: "penangkap", expected: "tangkap"},
		testItem{value: "penggila", expected: "gila"},
		testItem{value: "penghajar", expected: "hajar"},
		testItem{value: "pengqasar", expected: "qasar"},
		testItem{value: "pengudara", expected: "udara"},
		testItem{value: "pengupas", expected: "kupas"},
		testItem{value: "penyuara", expected: "suara"},
		testItem{value: "pelajar", expected: "ajar"},
		testItem{value: "pelabuh", expected: "labuh"},
		testItem{value: "petarung", expected: "tarung"},
		testItem{value: "terpercaya", expected: "percaya"},
		testItem{value: "pekerja", expected: "kerja"},
		testItem{value: "peserta", expected: "serta"},
		testItem{value: "mempengaruhi", expected: "pengaruh"},
		testItem{value: "mengkritik", expected: "kritik"},
		testItem{value: "bersekolah", expected: "sekolah"},
		testItem{value: "bertahan", expected: "tahan"},
		testItem{value: "mencapai", expected: "capai"},
		testItem{value: "dimulai", expected: "mulai"},
		testItem{value: "petani", expected: "tani"},
		testItem{value: "terabai", expected: "abai"},
		testItem{value: "mensyaratkan", expected: "syarat"},
		testItem{value: "mensyukuri", expected: "syukur"},
		testItem{value: "mengebom", expected: "bom"},
		testItem{value: "mempromosikan", expected: "promosi"},
		testItem{value: "memproteksi", expected: "proteksi"},
		testItem{value: "memprediksi", expected: "prediksi"},
		testItem{value: "pengkajian", expected: "kaji"},
		testItem{value: "pengebom", expected: "bom"},
		testItem{value: "bersembunyi", expected: "sembunyi"},
		testItem{value: "bersembunyilah", expected: "sembunyi"},
		testItem{value: "pelanggan", expected: "langgan"},
		testItem{value: "pelaku", expected: "laku"},
		testItem{value: "pelangganmukah", expected: "langgan"},
		testItem{value: "pelakunyalah", expected: "laku"},
		testItem{value: "perbaikan", expected: "baik"},
		testItem{value: "kebaikannya", expected: "baik"},
		testItem{value: "bisikan", expected: "bisik"},
		testItem{value: "menerangi", expected: "terang"},
		testItem{value: "berimanlah", expected: "iman"},
		testItem{value: "memuaskan", expected: "puas"},
		testItem{value: "berpelanggan", expected: "langgan"},
		testItem{value: "bermakanan", expected: "makan"},
		testItem{value: "menyala", expected: "nyala"},
		testItem{value: "menyanyikan", expected: "nyanyi"},
		testItem{value: "menyatakannya", expected: "nyata"},
		testItem{value: "penyanyi", expected: "nyanyi"},
		testItem{value: "penyawaan", expected: "nyawa"},
		testItem{value: "rerata", expected: "rata"},
		testItem{value: "lelembut", expected: "lembut"},
		testItem{value: "lemigas", expected: "ligas"},
		testItem{value: "kinerja", expected: "kerja"},
		testItem{value: "bertebaran", expected: "tebar"},
		testItem{value: "terasingkan", expected: "asing"},
		testItem{value: "membangunkan", expected: "bangun"},
		testItem{value: "mencintai", expected: "cinta"},
		testItem{value: "menduakan", expected: "dua"},
		testItem{value: "menjauhi", expected: "jauh"},
		testItem{value: "menggilai", expected: "gila"},
		testItem{value: "pembangunan", expected: "bangun"},
		testItem{value: "marwan", expected: "marwan"},
		testItem{value: "subarkah", expected: "subarkah"},
		testItem{value: "memberdayakan", expected: "daya"},
		testItem{value: "persemakmuran", expected: "makmur"},
		testItem{value: "keberuntunganmu", expected: "untung"},
		testItem{value: "kesepersepuluhnya", expected: "sepuluh"},
		testItem{value: "Perekonomian", expected: "ekonomi"},
		testItem{value: "menahan", expected: "tahan"},
		testItem{value: "peranan", expected: "peran"},
		testItem{value: "memberikan", expected: "beri"},
		testItem{value: "medannya", expected: "medan"},
		testItem{value: "idealis", expected: "ideal"},
		testItem{value: "idealisme", expected: "ideal"},
		testItem{value: "finalisasi", expected: "final"},
		testItem{value: "mentaati", expected: "taat"},
		testItem{value: "melewati", expected: "lewat"},
		testItem{value: "menganga", expected: "nganga"},
		testItem{value: "kupukul", expected: "pukul"},
		testItem{value: "kauhajar", expected: "hajar"},
		testItem{value: "kuasa-Mu", expected: "kuasa"},
		testItem{value: "nikmat-Ku", expected: "nikmat"},
		testItem{value: "allah-lah", expected: "allah"},
	}

	dictionary := NewDictionary("hancur", "benar", "apa", "siapa", "jubah",
		"baju", "beli", "celana", "hantu", "jual", "buku", "milik", "kulit",
		"sakit", "kasih", "buang", "suap", "nilai", "beri", "rambut", "adu",
		"suara", "daerah", "ajar", "kerja", "ternak", "asing", "raup", "gerak",
		"puruk", "terbang", "lipat", "ringkas", "warna", "yakin", "bangun",
		"fitnah", "vonis", "baru", "ajar", "tangkap", "kupas", "minum", "pukul",
		"cinta", "dua", "jauh", "ziarah", "nuklir", "gila", "hajar", "qasar",
		"udara", "populer", "warna", "yoga", "adil", "rumah", "muka", "labuh",
		"tarung", "tebar", "indah", "daya", "untung", "sepuluh", "ekonomi",
		"makmur", "telah", "serta", "percaya", "pengaruh", "kritik", "seko",
		"sekolah", "tahan", "capa", "capai", "mula", "mulai", "petan", "tani",
		"aba", "abai", "balas", "balik", "peran", "medan", "syukur", "syarat",
		"bom", "promosi", "proteksi", "prediksi", "kaji", "sembunyi", "langgan",
		"laku", "baik", "terang", "iman", "bisik", "taat", "puas", "makan",
		"nyala", "nyanyi", "nyata", "nyawa", "rata", "lembut", "ligas",
		"budaya", "karya", "ideal", "final", "taat", "tiru", "sepak", "kuasa",
		"malaikat", "nikmat", "lewat", "nganga", "allah",
	)

	stemmer := NewStemmer(dictionary)

	for _, item := range testItems {
		result := stemmer.Stem(item.value)
		if result != item.expected {
			t.Errorf("%s, expected: %s, result: %s", item.value, item.expected, result)
		}
	}
}
