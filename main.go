package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

var db *sql.DB
var err error

type Member struct {
	MemberID        string `json:"MemberID"`
	CardID          string `json:"CardID"`
	LastName        string `json:"LastName"`
	FirstName       string `json:"FirstName"`
	Sex             string `json:"Sex"`
	Title           string `json:"Title"`
	TitleOfCourtesy string `json:"TitleOfCourtesy"`
	BirthDate       string `json:"BirthDate"`
	HireDate        string `json:"HireDate"`
	Address         string `json:"Address"`
	City            string `json:"City"`
	Region          string `json:"Region"`
	PostalCode      string `json:"PostalCode"`
	Country         string `json:"Country"`
	HomePhone       string `json:"HomePhone"`
	Extension       string `json:"Extension"`
	Photo           string `json:"Photo"`
	Notes           string `json:"Notes"`
	ReportsTo       string `json:"ReportsTo"`
	AdmisionFee     string `json:"AdmisionFee"`
}

func creatMember(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		MemberID := r.FormValue("MemberID")
		CardID := r.FormValue("CardID")
		LastName := r.FormValue("LastName")
		FirstName := r.FormValue("FirstName")
		Sex := r.FormValue("Sex")
		Title := r.FormValue("Title")
		TitleOfCourtesy := r.FormValue("TitleOfCourtesy")
		BirthDate := r.FormValue("BirthDate")
		HireDate := r.FormValue("HireDate")
		Address := r.FormValue("Address")
		City := r.FormValue("City")
		Region := r.FormValue("Region")
		PostalCode := r.FormValue("PostalCode")
		Country := r.FormValue("Country")
		HomePhone := r.FormValue("HomePhone")
		Extension := r.FormValue("Extension")
		Photo := r.FormValue("Photo")
		Notes := r.FormValue("Notes")
		ReportsTo := r.FormValue("ReportsTo")
		AdmisionFee := r.FormValue("AdmisionFee")

		stmt, err := db.Prepare("INSERT INTO member (MemberID,CardID,LastName,FirstName,Sex,Title,TitleOfCourtesy,BirthDate,HireDate,Address,City,Region,PostalCode,Country,HomePhone,Extension,Photo,Notes,ReportsTo,AdmisionFee) VALUES(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")

		_, err = stmt.Exec(MemberID, CardID, LastName, FirstName, Sex, Title, TitleOfCourtesy, BirthDate, HireDate, Address, City, Region, PostalCode, Country, HomePhone, Extension, Photo, Notes, ReportsTo, AdmisionFee)

		if err != nil {
			fmt.Fprintf(w, "data duplicate")
		} else {
			fmt.Fprintf(w, "Date create ")
		}
	}
}

func getMembers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var Members []Member

	sql := `SELECT
				MemberID,
				IFNULL(CardID,'')CardID,
				IFNULL(LastName,'')LastName,
				IFNULL(FirstName,'')FirstName,
				IFNULL(Sex,'')Sex,
				IFNULL(Title,'')Title,
				IFNULL(TitleOfCourtesy,'')TitleOfCourtesy,
				IFNULL(BirthDate,'')BirthDate,
				IFNULL(HireDate,'')HireDate,
				IFNULL(Address,'')Address,
				IFNULL(City,'')City,
				IFNULL(Region,'')Region,
				IFNULL(PostalCode,'')PostalCode,
				IFNULL(Country,'')Country,
				IFNULL(HomePhone,'')HomePhone,
				IFNULL(Extension,'')Extension,
				IFNULL(Photo,'')Photo,
				IFNULL(Notes,'')Notes,
				IFNULL(ReportsTo,'')ReportsTo,
				IFNULL(AdmisionFee,'')AdmisionFee
			FROM member`

	result, err := db.Query(sql)

	defer result.Close()

	if err != nil {
		panic(err.Error())
	}

	for result.Next() {

		var member Member
		err := result.Scan(&member.MemberID, &member.CardID, &member.LastName, &member.FirstName, &member.Sex, &member.Title, &member.TitleOfCourtesy, &member.BirthDate, &member.HireDate, &member.Address, &member.City, &member.Region, &member.PostalCode, &member.Country, &member.HomePhone, &member.Extension, &member.Photo, &member.Notes, &member.ReportsTo, &member.AdmisionFee)

		if err != nil {
			panic(err.Error())
		}
		Members = append(Members, member)
	}

	json.NewEncoder(w).Encode(Members)
}

func getMember(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var Members []Member
	params := mux.Vars(r)

	sql := `SELECT
				MemberID,
				IFNULL(CardID,'')CardID,
				IFNULL(LastName,'')LastName,
				IFNULL(FirstName,'')FirstName,
				IFNULL(Sex,'')Sex,
				IFNULL(Title,'')Title,
				IFNULL(TitleOfCourtesy,'')TitleOfCourtesy,
				IFNULL(BirthDate,'')BirthDate,
				IFNULL(HireDate,'')HireDate,
				IFNULL(Address,'')Address,
				IFNULL(City,'')City,
				IFNULL(Region,'')Region,
				IFNULL(PostalCode,'')PostalCode,
				IFNULL(Country,'')Country,
				IFNULL(HomePhone,'')HomePhone,
				IFNULL(Extension,'')Extension,
				IFNULL(Photo,'')Photo,
				IFNULL(Notes,'')Notes,
				IFNULL(ReportsTo,'')ReportsTo,
				IFNULL(AdmisionFee,'')AdmisionFee
			FROM member WHERE MemberID = ?`

	result, err := db.Query(sql, params["MemberID"])

	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	var member Member

	for result.Next() {

		err := result.Scan(&member.MemberID, &member.CardID, &member.LastName, &member.FirstName, &member.Sex, &member.Title, &member.TitleOfCourtesy, &member.BirthDate, &member.HireDate, &member.Address, &member.City, &member.Region, &member.PostalCode, &member.Country, &member.HomePhone, &member.Extension, &member.Photo, &member.Notes, &member.ReportsTo, &member.AdmisionFee)

		if err != nil {
			panic(err.Error())
		}

		Members = append(Members, member)
	}

	json.NewEncoder(w).Encode(Members)
}

func deleteMember(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r)
	stmt, err := db.Prepare("DELETE FROM member WHERE MemberID = ?")

	if err != nil {
		panic(err.Error())
	}
	_, err = stmt.Exec(params["MemberID"])

	if err != nil {
		panic(err.Error())
	}

	fmt.Fprintf(w, "member with MemberID = %s was deleted", params["MemberID"])
}

func updateMember(w http.ResponseWriter, r *http.Request) {

	if r.Method == "PUT" {

		params := mux.Vars(r)

		newLastName := r.FormValue("LastName")
		newFirstName := r.FormValue("FirstName")
		newTitle := r.FormValue("Title")
		newTitleOfCourtesy := r.FormValue("TitleOfCourtesy")
		newAddress := r.FormValue("Address")
		newCity := r.FormValue("City")
		newRegion := r.FormValue("Region")
		newCountry := r.FormValue("Country")
		newPostalCode := r.FormValue("PostalCode")
		newHomePhone := r.FormValue("HomePhone")
		newExtension := r.FormValue("Extension")
		newPhoto := r.FormValue("Photo")
		newNotes := r.FormValue("Notes")
		newReportsTo := r.FormValue("ReportsTo")
		newAdmisionFee := r.FormValue("AdmisionFee")

		stmt, err := db.Prepare("UPDATE member SET LastName = ?,FirstName = ?,Title = ?,TitleOfCourtesy = ?,Address = ?,City = ?,Region = ?,Country = ?,PostalCode = ?,HomePhone = ?,Extension = ?,Photo = ?,Notes = ?,ReportsTo = ?,AdmisionFee = ? WHERE MemberID = ?")

		_, err = stmt.Exec(newLastName, newFirstName, newTitle, newTitleOfCourtesy, newAddress, newCity, newRegion, newCountry, newPostalCode, newHomePhone, newExtension, newPhoto, newNotes, newReportsTo, newAdmisionFee, params["MemberID"])

		if err != nil {
			panic(err.Error())
		}

		fmt.Fprintf(w, "member with MemberID = %s was updated", params["MemberID"])
	}
}

func getPost(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var Members []Member

	MemberID := r.FormValue("MemberID")
	FirstName := r.FormValue("FirstName")

	sql := `SELECT
			MemberID,
			IFNULL(CardID,'')CardID,
			IFNULL(LastName,'')LastName,
			IFNULL(FirstName,'')FirstName,
			IFNULL(Sex,'')Sex,
			IFNULL(Title,'')Title,
			IFNULL(TitleOfCourtesy,'')TitleOfCourtesy,
			IFNULL(BirthDate,'')BirthDate,
			IFNULL(HireDate,'')HireDate,
			IFNULL(Address,'')Address,
			IFNULL(City,'')City,
			IFNULL(Region,'')Region,
			IFNULL(PostalCode,'')PostalCode,
			IFNULL(Country,'')Country,
			IFNULL(HomePhone,'')HomePhone,
			IFNULL(Extension,'')Extension,
			IFNULL(Photo,'')Photo,
			IFNULL(Notes,'')Notes,
			IFNULL(ReportsTo,'')ReportsTo,
			IFNULL(AdmisionFee,'')AdmisionFee
		FROM member WHERE MemberID = ? AND FirstName = ?`

	result, err := db.Query(sql, MemberID, FirstName)

	if err != nil {
		panic(err.Error())
	}

	defer result.Close()

	var member Member

	for result.Next() {

		err := result.Scan(&member.MemberID, &member.CardID, &member.LastName, &member.FirstName, &member.Sex, &member.Title, &member.TitleOfCourtesy, &member.BirthDate, &member.HireDate, &member.Address, &member.City, &member.Region, &member.PostalCode, &member.Country, &member.HomePhone, &member.Extension, &member.Photo, &member.Notes, &member.ReportsTo, &member.AdmisionFee)

		if err != nil {
			panic(err.Error())
		}

		Members = append(Members, member)
	}

	json.NewEncoder(w).Encode(Members)

}

func main() {

	db, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/db_testing")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	r := mux.NewRouter()

	r.HandleFunc("/member", getMembers).Methods("GET")
	r.HandleFunc("/member", creatMember).Methods("POST")
	r.HandleFunc("/member/{MemberID}", getMember).Methods("GET")
	r.HandleFunc("/member/{MemberID}", deleteMember).Methods("DELETE")
	r.HandleFunc("/member/{MemberID}", updateMember).Methods("PUT")
	r.HandleFunc("/member/", getPost).Methods("POST", "GET")

	log.Fatal(http.ListenAndServe(":8080", r))
}
