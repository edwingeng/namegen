package namegen

var (
	greekMaleFirstNames = []string{
		"Abacus", "Acacius", "Achilios", "Achilles", "Achilleus", "Adonis", "Aegis", "Aeneas", "Aesop", "Aimilios", "Ajax", "Alec",
		"Alex", "Alexander", "Alexandre", "Alexandros", "Alexei", "Alexios", "Alexis", "Alipio", "Alpheus", "Alvertos", "Amphion",
		"Anastasios", "Anatole", "Anatolios", "Andonios", "Andreas", "Andrew", "Androcles", "Anthanasios", "Antonios", "Apollo",
		"Archimedes", "Ares", "Argo", "Argus", "Aristedes", "Aristotle", "Artemas", "Atlas", "Atticus", "Augustin", "Augustinos",
		"Avram", "Azarius", "Bacchus", "Balthasar", "Balthazar", "Basil", "Bastian", "Cadmus", "Calix", "Calixto", "Calypso",
		"Carolos", "Castor", "Christian", "Christiano", "Christion", "Christopher", "Christos", "Christy", "Claudios", "Cleanth",
		"Cleon", "Clete", "Cletus", "Colin", "Collins", "Constantinos", "Cosmo", "Costa", "Cronus", "Cymbeline", "Cyprian",
		"Cyrano", "Cyril", "Daedalus", "Daimon", "Damian", "Daphnis", "Darius", "Deacon", "Delias", "Demetri", "Demetrios",
		"Demetrius", "Demitrius", "Demos", "Dhimitrios", "Dimitros", "Dimos", "Diogenes", "Dion", "Dionysius", "Dorian", "Draco",
		"Drew", "Eladio", "Elias", "Endymion", "Erasmus", "Erastus", "Eros", "Eryx", "Eugen", "Eugene", "Eugenios", "Eustace",
		"Ezio", "Flavian", "Galen", "Galenus", "George", "Giles", "Gregor", "Gregorios", "Gregory", "Halcyon", "Hali", "Hector",
		"Helios", "Heraklees", "Herakles", "Hercules", "Heremias", "Hermes", "Hero", "Hesperos", "Hieremias", "Hieronymos",
		"Hilarion", "Hilary", "Hippolytos", "Homer", "Homeros", "Icarus", "Ignatios", "Ilya", "Indigo", "Isidore", "Isidorios",
		"Isiforos", "Janus", "Jase", "Jason", "Jeno", "Jerome", "Jonas", "Judas", "Kastas", "Konstantinos", "Kostas", "Kostis",
		"Kristopher", "Lafcadio", "Lazarus", "Leander", "Leandros", "Leon", "Leonidas", "Leonides", "Leontios", "Leviticus", "Lex",
		"Lexus", "Linus", "Loucas", "Loukas", "Lucais", "Lucas", "Lukas", "Luke", "Lukus", "Lyric", "Lysander", "Lysandros",
		"Macarius", "Markos", "Mattathias", "Matthaios", "Matthaiso", "Matthias", "Maximos", "Mimis", "Mitros", "Mitsos",
		"Moisis", "Moris", "Morpheus", "Moyses", "Myron", "Napoleon", "Narcissus", "Neander", "Nectarios", "Nemo", "Nereus",
		"Nestor", "Nicholas", "Nico", "Nicodemus", "Nicolas", "Nicomedes", "Nikasios", "Nike", "Nikolas", "Nikos", "Nikostratos",
		"Nile", "Oceanus", "Odysseus", "Oedipus", "Olimpio", "Olympos", "Orestes", "Orion", "Orpheus", "Osias", "Ozias", "Pan",
		"Panos", "Parthenios", "Pelagios", "Pello", "Pericles", "Perseus", "Peter", "Philadelphia", "Philander", "Philemon",
		"Philip", "Phillip", "Philo", "Philomon", "Phoebus", "Phoenix", "Phyllon", "Piers", "Plato", "Pluto", "Pollux",
		"Porfirio", "Poseidon", "Priam", "Ptolemy", "Rhodes", "Rihardos", "Rodion", "Romanos", "Rouvin", "Samouel", "Sandros",
		"Sebastian", "Semon", "Sisyphus", "Socrates", "Solon", "Sosthenes", "Stacey", "Stamatios", "Stamos", "Stavros",
		"Stephanos", "Stephen", "Symeon", "Takis", "Tarantino", "Terentilo", "Terentino", "Thanos", "Thao", "Theo", "Theodore",
		"Theodosios", "Theon", "Theophilos", "Theophilus", "Theron", "Theseus", "Tigris", "Tim", "Timaeus", "Timeo", "Timon",
		"Timotheos", "Timothy", "Timun", "Titos", "Tobias", "Tycho", "Tygrys", "Vasili", "Vasos", "Venedict", "Venedictos",
		"Veniamin", "Vernados", "Xan", "Xander", "Xanthos", "Xenophon", "Xenos", "Xylon", "Yannis", "Yeranos", "Zale", "Zan",
		"Zander", "Zandy", "Zeno", "Zenobios", "Zenon", "Zephyr", "Zeus",
	}

	greekFemaleFirstNames = []string{
		"Acacia", "Acantha", "Adelpha", "Agapi", "Agatha", "Agathe", "Agnes", "Aimilia", "Alcie", "Alcina", "Alethea",
		"Alex", "Alexa", "Alexandra", "Alexandria", "Alexandrina", "Alexina", "Alexis", "Alike", "Aliz", "Alizka",
		"Alpha", "Alsie", "Althea", "Amara", "Amarantha", "Amaryllis", "Ambrosia", "Aminta", "Anastacia", "Anastasha",
		"Anastasia", "Anatola", "Andromeda", "Anemone", "Ange", "Angela", "Angele", "Angeliki", "Angelina", "Aniceta",
		"Annis", "Annys", "Anthea", "Antigone", "Antimony", "Antinea", "Aphrodite", "Apollonia", "Arcadia", "Arcangela",
		"Arete", "Aretha", "Ariadne", "Arianna", "Arista", "Arkadina", "Artemis", "Artemisia", "Asta", "Aster",
		"Asteria", "Astraea", "Astraia", "Atalanta", "Athena", "Basilia", "Beranice", "Beraniece", "Berenice",
		"Berenike", "Bernice", "Beryl", "Beta", "Bronte", "Bryonia", "Calandra", "Calantha", "Calista", "Calixta",
		"Calla", "Callie", "Calliope", "Callista", "Calypso", "Carissa", "Cassandra", "Cassia", "Cassiane",
		"Cassiopeia", "Cat", "Cate", "Catherine", "Celena", "Chara", "Charis", "Charmian", "Chloe", "Chloris",
		"Christina", "Cinda", "Cipriana", "Circe", "Clematis", "Cleopatra", "Cleora", "Cliantha", "Clio", "Collins",
		"Cora", "Corinna", "Corisande", "Cosima", "Cressida", "Crisanta", "Cyane", "Cybele", "Cynara", "Cynthia",
		"Cytherea", "Damara", "Damaris", "Damia", "Damiana", "Daphne", "Daria", "Darian", "Delia", "Delta", "Demeter",
		"Demetria", "Demi", "Desdemona", "Despina", "Diamanta", "Diandra", "Diantha", "Dido", "Dina", "Dione", "Dionne",
		"Dora", "Dorcas", "Dorian", "Dorinda", "Doris", "Dorothea", "Dorothy", "Dree", "Echo", "Effie", "Effy",
		"Eirene", "Eirini", "Elara", "Electa", "Electra", "Elektra", "Elena", "Eleni", "Eleusine", "Elexis", "Eliana",
		"Eliane", "Elidi", "Eloisia", "Eos", "Epiphany", "Ereni", "Eudora", "Eugenia", "Eugenie", "Eulala", "Eulalia",
		"Eunice", "Euphemia", "Eurydice", "Eustacia", "Evadne", "Evangeline", "Evanthe", "Evathia", "Fantasia", "Fedora",
		"Filomena", "Gaia", "Galatea", "Galen", "Halcyon", "Hali", "Harmonia", "Harmony", "Hebe", "Hecuba", "Helen",
		"Helena", "Helia", "Hera", "Hermia", "Hermione", "Hero", "Hestia", "Hilary", "Hillary", "Hyacinth", "Hyacinthe",
		"Hyacynthe", "Ianthe", "Ilena", "Ilene", "Iliana", "Indigo", "Ino", "Io", "Ioanna", "Iola", "Iolanda", "Iolande",
		"Iolanthe", "Ione", "Ionia", "Ionna", "Iphigenia", "Irene", "Irina", "Iris", "Isadora", "Isaura", "Ismene",
		"Jocasta", "Junia", "Justina", "Kacia", "Kalliope", "Kallista", "Karissa", "Kasiani", "Kassandra", "Kassia",
		"Kassiani", "Katerina", "Katharine", "Katherine", "Katie", "Khloe", "Kosma", "Kosta", "Kostantina", "Kristiana",
		"Kynthia", "Lalage", "Lamia", "Larisa", "Larissa", "Leda", "Lenore", "Leora", "Letha", "Lethe", "Lex", "Lexia",
		"Lexis", "Lexus", "Libra", "Lici", "Liliah", "Lilika", "Lilis", "Lois", "Lotus", "Lydia", "Lyra", "Lyric",
		"Magdalen", "Magdalena", "Mago", "Mahail", "Mahaila", "Mahalia", "Maia", "Makis", "Malina", "Malva", "Margalo",
		"Margaret", "Maryam", "Maya", "Medea", "Medora", "Meg", "Melania", "Melanie", "Melantha", "Melany", "Melia",
		"Melina", "Melissa", "Melita", "Melody", "Melora", "Muse", "Mya", "Myra", "Myrtle", "Nani", "Narcissa", "Narda",
		"Natasa", "Nemea", "Neola", "Neoma", "Neri", "Nerida", "Nerine", "Nerissa", "Nickelle", "Nicola", "Nicole",
		"Nicolina", "Nicoline", "Nidia", "Nike", "Niki", "Nikola", "Nikoleta", "Nikolia", "Niobe", "Nitsa", "Nyssa",
		"Nyx", "Obelia", "Oceana", "Olympia", "Omega", "Ophelia", "Orion", "Orphea", "Pallas", "Pandora", "Panthea",
		"Parmenia", "Parthenia", "Pasha", "Peg", "Peggy", "Pelagia", "Penelope", "Penthia", "Peri", "Perrine",
		"Persephone", "Persis", "Pesha", "Peta", "Petal", "Petra", "Petrina", "Petrini", "Petronella", "Petronelle",
		"Phaedra", "Phedora", "Pheobe", "Phila", "Philadelphia", "Philippa", "Philomela", "Philomena", "Phoebe",
		"Phoenix", "Phyllida", "Phyllis", "Pinelopi", "Pipitsa", "Popi", "Praxis", "Psyche", "Raemonia", "Rena", "Reta",
		"Reveka", "Rhea", "Rheta", "Rheya", "Rhoda", "Roxane", "Rue", "Sapphira", "Sappho", "Selene", "Selia",
		"Sibley", "Sibyl", "Sirena", "Sofi", "Sofia", "Sophia", "Sophie", "Sophoon", "Sophronia", "Stasia", "Stavra",
		"Stefania", "Stephanie", "Sybella", "Sybil", "Tana", "Tancy", "Tansy", "Tasia", "Tasoula", "Tassia", "Tempe",
		"Tessa", "Thais", "Thalassa", "Thalia", "Thea", "Theia", "Thekla", "Themis", "Theo", "Theodosia", "Theone",
		"Theora", "Theresa", "Thesally", "Thetis", "Thisbe", "Tiffany", "Timothea", "Tina", "Titania", "Topaz", "Toula",
		"Typhaine", "Urania", "Ursa", "Varya", "Vasilia", "Vasiliki", "Venedicta", "Vernada", "Vernamina", "Veronike",
		"Veronique", "Violante", "Xanthe", "Xanthipe", "Xantho", "Xena", "Xenia", "Xenobia", "Yalena", "Yannia",
		"Yolanda", "Zelena", "Zelenia", "Zena", "Zenaida", "Zenobia", "Zephyr", "Zephyra", "Zephyrine", "Zeta", "Zita",
		"Zoe", "Zoei", "Zoey", "Zoie", "Zoila", "Zooey", "Zosma",
	}

	greekLastNames = []string{
		"Agne", "Agnes", "Agnes", "Alanis", "Alexopoulos", "Ambrosia", "Anagnos", "Anastas", "Anastos", "Andreadis", "Andreas",
		"Andris", "Angelis", "Angelopoulos", "Anthes", "Anthis", "Antoni", "Antoniou", "Antonopoulos", "Apostolos", "Apostolou",
		"Argyros", "Artino", "Arvanitis", "Asker", "Athan", "Athanas", "Athanasiou", "Athans", "Athas", "Bakas", "Bakas", "Balaban",
		"Ballas", "Balli", "Ballis", "Ballis", "Banis", "Barba", "Barbas", "Barberis", "Barlas", "Baros", "Bella", "Biros", "Booras",
		"Boosalis", "Bouras", "Buros", "Callas", "Callas", "Callis", "Caras", "Caras", "Carras", "Castellanos", "Chaconas", "Chontos",
		"Chris", "Christakos", "Christodoulou", "Christopoulos", "Christos", "Christou", "Chronis", "Collias", "Comis", "Condos",
		"Constantinides", "Constantinou", "Contos", "Cosmos", "Cosse", "Costas", "Delis", "Dellis", "Demetriou", "Demo", "Demopoulos",
		"Demos", "Dimitriou", "Doukas", "Drakos", "Drivas", "Drivas", "Dukas", "Economides", "Economos", "Economou", "Eliades", "Elias",
		"Eliopoulos", "Floros", "Fotopoulos", "Fotos", "Frangos", "Gabris", "Galanis", "Galanos", "Galatas", "Ganas", "Ganas", "Ganis",
		"Garis", "Gekas", "Georgas", "Georgiades", "Georgiadis", "Georgiou", "Georgopoulos", "German", "Gerou", "Gianakos",
		"Giannopoulos", "Gianopoulos", "Gikas", "Glaros", "Goga", "Gogola", "Golias", "Gonce", "Goulas", "Gounaris", "Grivas",
		"Halkias", "Hallas", "Hanas", "Harris", "Hatzis", "Hero", "Herod", "Hondros", "Ioannou", "Kairis", "Kakos", "Kalfas", "Kalivas",
		"Kallas", "Kallis", "Kamber", "Kanas", "Kanelos", "Kappas", "Kara", "Karahalios", "Karalis", "Karalis", "Karas", "Karras",
		"Karras", "Katsaros", "Kazan", "Kefalas", "Kellis", "Kokes", "Kollias", "Kondos", "Kontos", "Korba", "Kormos", "Kosko", "Kosta",
		"Kostas", "Kostopoulos", "Kotas", "Kotas", "Kouris", "Kritikos", "Ladas", "Lagana", "Lagos", "Lambros", "Lampros", "Laskaris",
		"Lasko", "Lazos", "Lekas", "Leos", "Leva", "Leventis", "Liakos", "Lias", "Lillis", "Linard", "Livas", "Logo", "Lois", "Loris",
		"Loukas", "Macris", "Maheras", "Makos", "Makris", "Mallas", "Mangas", "Maniatis", "Manikas", "Manis", "Manolis", "Manos",
		"Maragos", "Maras", "Marinos", "Maris", "Markos", "Martis", "Mate", "Mates", "Matis", "Mattas", "Mattas", "Matthias", "Mavros",
		"Mega", "Melis", "Mellas", "Mellis", "Mena", "Metaxas", "Metro", "Michaelides", "Michel", "Mihal", "Mikos", "Milas", "Milonas",
		"Mina", "Minas", "Minga", "Mires", "Miron", "Misko", "Mola", "Moraitis", "Moros", "Mula", "Mundis", "Myron", "Nanos", "Nanos",
		"Nasso", "Nicolaides", "Nicolaou", "Nicoli", "Nicolo", "Nikas", "Nikitas", "Nikolas", "Nini", "Nino", "Pagonis", "Pagonis",
		"Palas", "Pallas", "Panagakos", "Panagopoulos", "Panagos", "Panas", "Panas", "Panos", "Papa", "Papadakis", "Papadopoulos",
		"Papageorgiou", "Papas", "Pappas", "Patera", "Patras", "Paules", "Paulos", "Pavlis", "Penna", "Pepi", "Pepi", "Pepi", "Peri",
		"Peris", "Perris", "Perro", "Petrakis", "Petralia", "Petras", "Petrides", "Petro", "Petropoulos", "Petros", "Petrou", "Pipes",
		"Polites", "Politis", "Politis", "Poulos", "Primo", "Psomas", "Pulos", "Rallis", "Raptis", "Regas", "Rella", "Remes", "Remis",
		"Rigas", "Rines", "Roda", "Rodi", "Rodia", "Rodis", "Rodis", "Rokos", "Romanos", "Ronda", "Rondo", "Rosi", "Rosso", "Rota",
		"Roussos", "Rubis", "Rubis", "Sagona", "Salis", "Sallas", "Sallis", "Samaras", "Sanna", "Sarantos", "Sarkis", "Sarris", "Sava",
		"Savas", "Sica", "Sideris", "Simos", "Simos", "Simos", "Siska", "Sisko", "Soter", "Sotir", "Sotiropoulos", "Soulis", "Spanos",
		"Speros", "Spiros", "Stamas", "Stamatis", "Stamos", "Stanis", "Stathakis", "Stathis", "Stathopoulos", "Stathos",
		"Stavropoulos", "Stavros", "Stavrou", "Stratis", "Stratos", "Takes", "Tassi", "Tasso", "Tavoularis", "Teresi", "Terzi",
		"Thanos", "Theodorou", "Thoma", "Tocci", "Toles", "Tomaras", "Toto", "Trakas", "Tripi", "Valis", "Vallas", "Valli", "Vallis",
		"Varela", "Vasco", "Vasil", "Vasilakis", "Vasiliou", "Vassos", "Velis", "Vidales", "Vitalis", "Vlachos", "Vlahakis",
		"Vlahos", "Xenakis", "Xenos", "Zenon", "Zervas", "Zervos", "Zika", "Zografos",
	}
)
