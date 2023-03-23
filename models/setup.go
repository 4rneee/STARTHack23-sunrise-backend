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
