package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"golang.org/x/crypto/bcrypt"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open("sqlite3", "sunrise.db")

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.LogMode(true)

	database.AutoMigrate(&User{}, &Poll{}, &PollAnswer{}, &Stream{}, &Comment{}, &Clip{}, &Tweet{})

	DB = database

	AddMockUsers()
	AddMockFriends()
	AddMockStreams()
	AddMockPolls()
	AddMockComments()
}

func pw(password string) []byte {
	p, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return p
}

func AddMockUsers() {
	mockUsers := []User{
		{ID: 1, Name: "Test User", Email: "A", Points: 210, Password: pw("A"), StreamID: 1},
		{ID: 2, Name: "Noah", Email: "noah@tum.de", Points: 0, Password: pw("noah"), StreamID: 3},
		{ID: 3, Name: "Arne", Email: "A@F.com", Points: 0, Password: pw("123"), StreamID: 0},
		{ID: 4, Name: "Harry", Email: "H@Z.com", Points: 85, Password: pw("harry"), StreamID: 1},
		{ID: 5, Name: "Peyman", Email: "P@E.com", Points: 90, Password: pw("pucpm"), StreamID: 3},
		{ID: 6, Name: "Karola", Email: "K@k.com", Points: 0, Password: pw("k"), StreamID: 1},
		{ID: 7, Name: "Alice", Email: "alice@example.com", Points: 50, Password: pw("alice"), StreamID: 2},
		{ID: 8, Name: "Bob", Email: "bob@example.com", Points: 30, Password: pw("bob"), StreamID: 1},
		{ID: 9, Name: "Charlie", Email: "charlie@example.com", Points: 80, Password: pw("charlie"), StreamID: 3},
		{ID: 10, Name: "Dave", Email: "dave@example.com", Points: 20, Password: pw("dave"), StreamID: 1},
		{ID: 11, Name: "Eve", Email: "eve@example.com", Points: 100, Password: pw("eve"), StreamID: 2},
		{ID: 12, Name: "Frank", Email: "frank@example.com", Points: 75, Password: pw("frank"), StreamID: 1},
		{ID: 13, Name: "Grace", Email: "grace@example.com", Points: 65, Password: pw("grace"), StreamID: 3},
		{ID: 14, Name: "Heidi", Email: "heidi@example.com", Points: 90, Password: pw("heidi"), StreamID: 1},
		{ID: 15, Name: "Isaac", Email: "isaac@example.com", Points: 55, Password: pw("isaac"), StreamID: 2},
		{ID: 16, Name: "Jasmine Smith", Email: "jasmine.smith@example.com", Points: 100, Password: pw("password16"), StreamID: 2},
		{ID: 17, Name: "Daniel Johnson", Email: "daniel.johnson@example.com", Points: 50, Password: pw("password17"), StreamID: 3},
		{ID: 18, Name: "Emily Wilson", Email: "emily.wilson@example.com", Points: 20, Password: pw("password18"), StreamID: 1},
		{ID: 19, Name: "Ethan Brown", Email: "ethan.brown@example.com", Points: 75, Password: pw("password19"), StreamID: 2},
		{ID: 20, Name: "Olivia Davis", Email: "olivia.davis@example.com", Points: 30, Password: pw("password20"), StreamID: 3},
		{ID: 21, Name: "Sophia Rodriguez", Email: "sophia.rodriguez@example.com", Points: 85, Password: pw("password21"), StreamID: 1},
		{ID: 22, Name: "Sophie Thompson", Email: "sophie.thompson@example.com", Points: 120, Password: pw("password22"), StreamID: 2},
		{ID: 23, Name: "Thomas Parker", Email: "thomas.parker@example.com", Points: 85, Password: pw("password23"), StreamID: 3},
		{ID: 24, Name: "Ella Mitchell", Email: "ella.mitchell@example.com", Points: 50, Password: pw("password24"), StreamID: 1},
		{ID: 25, Name: "Christopher Turner", Email: "christopher.turner@example.com", Points: 10, Password: pw("password25"), StreamID: 2},
		{ID: 26, Name: "Ava Roberts", Email: "ava.roberts@example.com", Points: 70, Password: pw("password26"), StreamID: 3},
		{ID: 27, Name: "Jacob Campbell", Email: "jacob.campbell@example.com", Points: 55, Password: pw("password27"), StreamID: 1},
		{ID: 28, Name: "Grace Phillips", Email: "grace.phillips@example.com", Points: 30, Password: pw("password28"), StreamID: 2},
		{ID: 29, Name: "Andrew Anderson", Email: "andrew.anderson@example.com", Points: 65, Password: pw("password29"), StreamID: 3},
	}

	for _, user := range mockUsers {
		DB.Create(&user)
	}
}

func AddMockStreams() {
	mockStreams := []Stream{
		{ID: 1, Name: "Washington Capitals - Vegas Golden Knights", Description: "Stanley Cup Final", Thumbnail: "https://cdn.britannica.com/50/219150-050-0032E44D/Marc-Andre-Fleury-Vegas-Golden-Knights-Stanley-Cup-Final-2018.jpg"},
		{ID: 2, Name: "Bears - Bulls", Description: "Deutsche Bundesliga", Thumbnail: "https://www.scb.ch/fileadmin/_processed_/3/f/csm_SIM13266__DSC8739_Fotocredit_Tom_HILLER_2023-lpr_cc7d931c90.jpg"},
		{ID: 3, Name: "Belgien - Schweiz", Description: "U18 EM", Thumbnail: "https://phothockey.ch/wp-content/uploads/2023/02/IMG_5320-392x272.jpg"},
		{ID: 4, Name: "PSU - CCM", Description: "Bundescup - Rückspiel", Thumbnail: "https://static.wixstatic.com/media/725be2_4e93b8d0075446879b2985dcb155f770~mv2.jpg/v1/fill/w_1960,h_908,al_c,q_85,usm_0.66_1.00_0.01,enc_auto/725be2_4e93b8d0075446879b2985dcb155f770~mv2.jpg"},
		{ID: 5, Name: "The blue promise", Description: "Documentation on a historic woner", Thumbnail: "https://www.evz.ch/typo3temp/assets/_processed_/4/d/csm_ff813cd10cbaa0943b4182ad4f591edd23f65050-fp-600-460-8-42_4144586341.jpg"},
		{ID: 6, Name: "Canada - Russia", Description: "World Cup recap", Thumbnail: "https://image.stern.de/31606106/t/rR/v2/w1440/r1/-/eishockey-kanada-russland.jpg"},
		{ID: 7, Name: "Season recap", Description: "NHL Cup 2022 recap", Thumbnail: "https://www.hockeyweb.de/index.php?rex_media_type=hw_article_image&rex_media_file=120820102.jpg"},
        {ID: 8, Name: "Great Britain - Canada (U18)", Description: "U18 Women World Cup", Thumbnail: "https://blob.iihf.com/iihf-media/iihfmvc/media/2022/ww18ii/gallery/18c.jpg"},
        {ID: 9, Name: "4 decades after 'Miracle'", Description: "Old school and its highlights", Thumbnail: "https://newscdn2.weigelbroadcasting.com/NI5FB-1582385302-158182-blog-200222003834-03-miracle-on-ice-reunion-restricted-exlarge-169.jpg"},


	}

	for _, stream := range mockStreams {
		DB.Create(&stream)
	}
}

func AddMockFriends() {
	DB.Exec("INSERT INTO friends VALUES (1, 2)")
	DB.Exec("INSERT INTO friends VALUES (2, 1)")

	DB.Exec("INSERT INTO friends VALUES (2, 3)")
	DB.Exec("INSERT INTO friends VALUES (3, 2)")

	DB.Exec("INSERT INTO friends VALUES (3, 4)")
	DB.Exec("INSERT INTO friends VALUES (4, 3)")

	DB.Exec("INSERT INTO friends VALUES (3, 5)")
	DB.Exec("INSERT INTO friends VALUES (5, 3)")

	DB.Exec("INSERT INTO friends VALUES (2, 6)")
	DB.Exec("INSERT INTO friends VALUES (6, 2)")

	DB.Exec("INSERT INTO friends VALUES (3, 6)")
	DB.Exec("INSERT INTO friends VALUES (6, 3)")

	DB.Exec("INSERT INTO friends VALUES (4, 5)")
	DB.Exec("INSERT INTO friends VALUES (5, 4)")
}

func AddMockPolls() {
	mockPolls := []Poll{
		{Question: "Who will win?", StreamID: 1},
		{Question: "How many goals will be scored in the first half?", StreamID: 1},
		{Question: "Which team will score the next goal?", StreamID: 2},
	}

	mockPollAnswers := []PollAnswer{
		{PollID: 0, Answer: "Washington Capitals", Votes: 645},
		{PollID: 0, Answer: "Vegas Golden Knights", Votes: 559},
		{PollID: 1, Answer: "0", Votes: 31},
		{PollID: 1, Answer: "1", Votes: 55},
		{PollID: 1, Answer: "2", Votes: 79},
		{PollID: 1, Answer: "3", Votes: 75},
		{PollID: 1, Answer: "4", Votes: 33},
		{PollID: 1, Answer: "5 or more", Votes: 24},
	}

	for _, poll := range mockPolls {
		DB.Create(&poll)
	}

	for _, pollAnswer := range mockPollAnswers {
		DB.Create(&pollAnswer)
	}
}

func AddMockComments() {
	mockComments := []Comment{
		{Content: "LETS GO VEGAS!!", StreamID: 1, UserID: 4},
		{Content: "This game is sooo exciting. I hope Washington wins.", StreamID: 1, UserID: 6},
		{Content: "Hi", StreamID: 1, UserID: 5},
		{Content: "Greetings from Zurich", StreamID: 1, UserID: 2},
		{Content: "Watching this while cooking", StreamID: 2, UserID: 3},
		{Content: "Hi", StreamID: 4, UserID: 5},
		{Content: "🎉🎉🎉", StreamID: 2, UserID: 4},
	}

	for _, comment := range mockComments {
		DB.Create(&comment)
	}
}
