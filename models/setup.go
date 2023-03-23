package models

import (
	"golang.org/x/crypto/bcrypt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open("sqlite3", "sunrise.db")

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.LogMode(true)

	database.AutoMigrate(&User{}, &Poll{}, &PollAnswer{}, &Stream{}, &Comment{}, &Clip{})

	DB = database

    AddMockUsers()
    AddMockStreams()
}

func pw(password string) []byte {
    p, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return p
}

func AddMockUsers() {
    mockUsers := []User{
        {ID: 1, Name: "Test User", Email: "T@t.com", Points: 210, Password: pw("123"), StreamID: 1},
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
        {ID: 1, Name: "Fleury Vegas - Golden Knights", Description: "Stanley Cup Final", Thumbnail: "https://cdn.britannica.com/50/219150-050-0032E44D/Marc-Andre-Fleury-Vegas-Golden-Knights-Stanley-Cup-Final-2018.jpg"},
        {ID: 2, Name: "Bears - Bulls", Description: "Deutsche Bundesliga", Thumbnail: "https://www.scb.ch/fileadmin/_processed_/3/f/csm_SIM13266__DSC8739_Fotocredit_Tom_HILLER_2023-lpr_cc7d931c90.jpg"},
        {ID: 3, Name: "Belgien - Schweiz", Description: "U18 EM", Thumbnail: "https://phothockey.ch/wp-content/uploads/2023/02/IMG_5320-392x272.jpg"},
        {ID: 4, Name: "PSU - CCM", Description: "Bundescup - Rückspiel", Thumbnail: "https://static.wixstatic.com/media/725be2_4e93b8d0075446879b2985dcb155f770~mv2.jpg/v1/fill/w_1960,h_908,al_c,q_85,usm_0.66_1.00_0.01,enc_auto/725be2_4e93b8d0075446879b2985dcb155f770~mv2.jpg"},
        {ID: 5, Name: "The blue promise", Description: "Documentation on a historic woner", Thumbnail: "https://www.evz.ch/typo3temp/assets/_processed_/4/d/csm_ff813cd10cbaa0943b4182ad4f591edd23f65050-fp-600-460-8-42_4144586341.jpg"},


    }

    for _, stream := range mockStreams {
        DB.Create(&stream)
    }
}

func AddMockFriends() {

}
